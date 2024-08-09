#   证书

- 使用 OpenSSL 生成私钥和 CSR

```bash
# 创建配置文件 openssl-san.cnf
[ req ]
default_bits       = 4096
prompt             = no
default_md         = sha256
req_extensions     = req_ext
distinguished_name = dn

[ dn ]
C  = CN
ST = State
L  = City
O  = Organization
OU = Organizational Unit
CN = hsiangya.top

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = hsiangya.top
# DNS.2 = harbor.hsiangya.top


# 生成私钥
openssl genrsa -out hsiangya.key 4096

# 生成证书签名请求 (CSR)
openssl req -new -key hsiangya.key -out hsiangya.csr -config openssl-san.cnf
```

- 使用 Keytool 生成私钥和 CSR

```bash
# 生成密钥库和密钥对
keytool -genkeypair -alias harboralias -keyalg RSA -keysize 2048 -keystore harbor.hsiangya.top.jks

```

# Mysql

## helm安装

- 运行命令安装


```bash
helm install mysql-cluster \
oci://registry-1.docker.io/bitnamicharts/mysql \
--set global.storageClass=managed-nfs-storage \
--set architecture=replication \
--set secondary.replicaCount=2

# 卸载
helm uninstall mysql-cluster
```

> mysql-cluster：自己定义的名称
>
> storageClass的名称：`managed-nfs-storage` 根据实际名称填写，可以使用`kubectl get sc 查看`
>
> architecture：架构`standalone`或`replication`
>
> secondary.replicaCount：副本数

- 创建成功后，会输出如下信息

```bash
# 连接mysql的root密码保存到secret中
  echo Username: root
  MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace default mysql-cluster -o jsonpath="{.data.mysql-root-password}" | base64 -d)

To connect to your database:
# 创建临时连接客户端，可以连接mysql集群
  1. Run a pod that you can use as a client:

      kubectl run mysql-cluster-client --rm --tty -i --restart='Never' --image  docker.io/bitnami/mysql:8.4.1-debian-12-r3 --namespace default --env MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD --command -- bash

  2. To connect to primary service (read/write):
# 通过service命令访问primary节点
      mysql -h mysql-cluster-primary.default.svc.cluster.local -uroot -p"$MYSQL_ROOT_PASSWORD"

  3. To connect to secondary service (read-only):
# 访问从节点，只读
      mysql -h mysql-cluster-secondary.default.svc.cluster.local -uroot -p"$MYSQL_ROOT_PASSWORD"
```

- 查看运行状态

```bash
helm list
kubectl get pod
```

- 根据之前提示的信息，测试连接集群

```bash
# 获取密码
MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace default mysql-cluster -o jsonpath="{.data.mysql-root-password}" | base64 -d)

# 进入到pod中
kubectl run mysql-cluster-client --rm --tty -i --restart='Never' --image  docker.io/bitnami/mysql:8.4.1-debian-12-r3 --namespace default --env MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD --command -- bash

# 连接mysql cli
mysql -h mysql-cluster-primary.default.svc.cluster.local -uroot -p"$MYSQL_ROOT_PASSWORD"
# 查看主从复制状态
show slave status\G;
```

## 配置文件安装

### NodePort单体

- 创建命名空间：`namespace.yaml`

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: grpc-k8s
```

- 创建ConfigMap存储Mysql配置文件：`mysql-config.yaml`

```yaml
apiVersion: v1
kind: ConfigMap # configMap
metadata:
  name: mysql-config
  namespace: grpc-k8s  # 指定命名空间
  labels:
    app: mysql
data:
  my.cnf: |-
    [client]
    default-character-set=utf8mb4
    [mysql]
    default-character-set=utf8mb4
    [mysqld]
    max_connections = 2000
    secure_file_priv=/var/lib/mysql
    sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```

- 构建pvc：`mysql-storage.yaml`

  > 这里由于环境已经通过nfs构建了名为managed-nfs-storage的stroageClass，以nfs形式挂在，会自动进行PV的申领，所以不需要单独编写pv配置文件

```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  namespace: grpc-k8s
  name: grpc-k8s-mysql-pvc
spec:
  storageClassName: managed-nfs-storage  # 指定 StorageClass
  resources:
    requests:
      storage: 10Gi  # 设置 pvc 存储资源大小
  accessModes:
  - ReadWriteOnce
```

- 编写deployment文件：`mysql-deploy.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: grpc-k8s
  name: mysql
  labels:
    app: mysql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
      - name: mysql
        image: mysql:8.0.19
        ports:
        - containerPort: 3306
        env:
        - name: MYSQL_ROOT_PASSWORD
          value: "password"
        resources:
          limits:
            cpu: 2000m
            memory: 512Mi
          requests:
            cpu: 2000m
            memory: 512Mi
        livenessProbe:
          initialDelaySeconds: 30
          periodSeconds: 10
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 3
          exec:
            command: ["mysqladmin", "-uroot", "-p${MYSQL_ROOT_PASSWORD}", "ping"]
        readinessProbe:
          initialDelaySeconds: 10
          periodSeconds: 10
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 3
          exec:
            command: ["mysqladmin", "-uroot", "-p${MYSQL_ROOT_PASSWORD}", "ping"]
        volumeMounts:
        - name: data
          mountPath: /var/lib/mysql
        - name: config
          mountPath: /etc/mysql/conf.d
        - name: localtime
          readOnly: true
          mountPath: /etc/localtime
      volumes:
      - name: data
        persistentVolumeClaim:
          claimName: grpc-k8s-mysql-pvc # 改为适当的PVC名称，确保它配置了managed-nfs-storage
      - name: config
        configMap:
          name: mysql-config
      - name: localtime
        hostPath:
          type: File
          path: /etc/localtime
```

- 创建服务运行配置文件：`mysql-service.yaml`

```yaml
apiVersion: v1
kind: Service
metadata:
  name: grpc-k8s-mysql
  namespace: grpc-k8s
  labels:
    app: grpc-k8s-mysql
spec:
  type: NodePort
  ports:
  - name: mysql
    port: 3306
    targetPort: 3306
    # nodePort: 30036  手动制定端口号，如果不指定则由k8s自己创建
  selector:
    app: mysql
```

- 运行文件

```bash
kubectl apply -f namespace.yaml
kubectl apply -f mysql-config.yaml
sudo mkdir -p /opt/k8s/data/grpc_k8s/mysql/data
kubectl apply -f mysql-storage.yaml
kubectl apply -f mysql-deploy.yaml
kubectl apply -f mysql-service.yaml
kubectl get svc -n grpc-k8s # 查看服务运行情况
```

### 主从

- 创建configMap：`configMap.yaml `

  > 配置主库能够将复制日志提供给从库，并且从库拒绝任何不是通过复制进行的写操作

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mysql-replica
  labels:
    app: mysql
  namespace: dev
data:
  primary.cnf: |
    # Apply this config only on the primary.
    [mysql]
    default-character-set=utf8mb4
    [mysqld]
    log-bin
    character-set-server=utf8mb4
    [client]
    default-character-set=utf8mb4

  replica.cnf: |
    # Apply this config only on replicas.
    [mysql]
    default-character-set=utf8mb4
    [mysqld]
    super-read-only    
    character-set-server=utf8mb4
    [client]
    default-character-set=utf8mb4
```

- 编写service：`service.yaml`

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql-replica
  labels:
    app: mysql
  namespace: dev
spec:
  ports:
  - name: mysql
    port: 3306
  clusterIP: None
  selector:
    app: mysql
---
apiVersion: v1
kind: Service
metadata:
  name: mysql-read
  labels:
    app: mysql
  namespace: dev
spec:
  ports:
  - name: mysql
    port: 3306
  selector:
    app: mysql
```

- 运行文件

```bash
kubectl apply -f configMap.yaml
kubectl apply -f service.yaml
kubectl get configMap -n dev
```







# Mongo

## 单节点

- 创建配置文件：`mongo-configMap.yml`

```yml
apiVersion: v1
kind: Namespace
metadata:
  name: dev
--- 
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: dev
  name: mongodb-config
data:
  mongod.conf: |
    storage:
      dbPath: /data/db
    net:
      bindIp: 0.0.0.0
      port: 27017
    security:
      authorization: "enabled"
```

- 创建secret文件：`mongo-secret.yml`

  > - base64编码：`echo -n 'your_string_here' | base64`
  > - base64解码：`echo 'base64_encoded_string_here' | base64 --decode`

```yml
apiVersion: v1
kind: Secret
metadata:
  namespace: dev
  name: mongodb-secret
type: Opaque
data:
  mongo-root-username: xxxxx= # base64 encoded value of 'rootuser'
  mongo-root-password: xxxxxx= # base64 encoded value of 'rootpassword'
```

- 编辑部署文件：`mongo-deployment.yml`

```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: dev
  name: mongodb
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mongodb
  template:
    metadata:
      labels:
        app: mongodb
    spec:
      containers:
        - name: mongodb
          image: mongo:5.0.28
          ports:
            - containerPort: 27017
          volumeMounts:
            - name: mongodb-data
              mountPath: /data/db
            - name: mongodb-config
              mountPath: /etc/mongod.conf
              subPath: mongod.conf
          env:
            - name: MONGO_INITDB_ROOT_USERNAME
              valueFrom:
                secretKeyRef:
                  name: mongodb-secret
                  key: mongo-root-username
            - name: MONGO_INITDB_ROOT_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: mongodb-secret
                  key: mongo-root-password
      volumes:
        - name: mongodb-data
          persistentVolumeClaim:
            claimName: mongodb-pvc
        - name: mongodb-config
          configMap:
            name: mongodb-config
```

- 创建service文件：`mongo-service.yml`

```yml
apiVersion: v1
kind: Service
metadata:
  name: mongodb
  namespace: dev
spec:
  ports:
    - port: 27017
      targetPort: 27017
      # nodePort: 32017  # 不指定k8s默认自动分配
  selector:
    app: mongodb
  type: NodePort
```

- 创建pvc文件：`mongo-pvc.yml`

```yml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  namespace: dev
  name: mongodb-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
  storageClassName: managed-nfs-storage
```

- 启动服务

```bash
kubectl apply -f mongo-configMap.yml
kubectl apply -f mongo-secret.yml
kubectl apply -f mongo-deployment.yml
kubectl apply -f mongo-pvc.yml
kubectl apply -f mongo-service.yml
```

# Harbor

## 从helm安装

- 添加存储仓、拉取配置、创建基础配置

```bash
# 添加存储仓库
helm repo add harbor https://helm.goharbor.io
helm repo list
helm search repo harbor -l |  grep harbor/harbor  | head  -4

# 拉取harbor
helm pull harbor/harbor --version 1.15.0
tar zxvf harbor-1.15.0.tgz

# 创建一个namespace
kubectl create namespace harbor

# 创建ssl证书
kubectl create secret tls harbor-cert --cert=/opt/certs/harbor.hsiangya.top.pem --key=/opt/certs/hsiangya.key -n harbor

# 编辑配置我呢见
cd harbor
vim values.yaml

```

- 修改配置文件：`values.yaml`
  - persistentVolumeClaim：数据卷挂载 指定`storageClass`(通过`kubectl get storageClass`查询)
  - harborAdminPassword：配置管理员密码
  - exppos：secretName,证书配置信息与域名

- 运行

```bash
helm install harbor . -f values.yaml -n harbor 
```

