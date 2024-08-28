# 边角料

## 创建证书私钥

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

## Nerdctl

github：`https://github.com/moby/buildkit`

> - nerdctl分精简版与完整版
> - 完整版 `lib` 目录下有现成 `buildkit.service` 文件，需要注意默认的路径是 `/usr/local/bin/buildkitd`

```bash
# 下载并安装nerdctl、
wget https://github.com/containerd/nerdctl/releases/download/v1.7.6/nerdctl-full-1.7.6-linux-amd64.tar.gz
tar -zxvf nerdctl-full-1.7.6-linux-amd64.tar.gz
tar Cxzvvf /usr/local nerdctl-full-1.7.6-linux-amd64.tar.gz

# 有现成的service文件
cp lib/systemd/system/buildkit.service /lib/systemd/system/
systemctl enable buildkit.service --now
sudo mv nerdctl /usr/local/bin/
nerdctl --version
```

- 验证

```bash
echo 'FROM alpine:latest
CMD echo "Hello from custom image"' > Dockerfile

nerdctl build -t alpine:test .
```

## 将docker服务迁移到k8s

- 原始docker命令

```bash
docker run --name bill -dp 10000:80 \
-w /app \
-v "/data/beancount:/data/beancount" \
-v "/data/beancount/icons:/app/public/icons" \
-v "/data/beancount/config:/app/config" \
-v "/data/beancount/logs:/app/logs" \
xdbin/beancount-gs:latest \
sh -c "cp -rn /app/public/default_icons/* /app/public/icons && ./beancount-gs -p 80"
```

- 编辑k8s配置文件

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: hsiangya
---
# ... (previous PersistentVolumeClaim and Deployment configurations remain unchanged)
---
apiVersion: v1
kind: Service
metadata:
  name: beancount-gs
  namespace: hsiangya
spec:
  selector:
    app: beancount-gs
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: beancount-ingress
  namespace: hsiangya
  annotations:
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
    nginx.ingress.kubernetes.io/force-ssl-redirect: "true"
    nginx.ingress.kubernetes.io/auth-type: basic
    nginx.ingress.kubernetes.io/auth-secret: basic-auth
    nginx.ingress.kubernetes.io/auth-realm: "Authentication Required"
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - hsiangya.top
    secretName: hsiangya-tls
  rules:
  - host: hsiangya.top
    http:
      paths:
      - path: /web
        pathType: Prefix
        backend:
          service:
            name: beancount-gs
            port: 
              number: 80
      - path: /api
        pathType: Prefix
        backend:
          service:
            name: beancount-gs
            port: 
              number: 80
```

- 启动服务

```bash
# 创建一个基础验证
sudo apt install apache2-utils
htpasswd -c auth hsiangya
kubectl create secret generic basic-auth --from-file=auth -n hsiangya

kubectl create secret tls hsiangya-tls -n hsiangya \
  --key /opt/certs/hsiangya.top.key \
  --cert /opt/certs/hsiangya.top.pem
kubectl apply -f bill-deploy.yaml
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

### NodePort单体手动构建PVC

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

### 单体自动构建PVC

- 编辑配置文件：`config.yaml`

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mysql-config
  namespace: dev 
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

- 编辑PVC文件：`pvc.yaml`:

```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  namespace: dev
  name: dev-mysql-pvc
spec:
  storageClassName: managed-nfs-storage  # 指定 StorageClass
  resources:
    requests:
      storage: 10Gi  # 设置 pvc 存储资源大小
  accessModes:
  - ReadWriteOnce
```

- 编辑deploy文件：`deploy,yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: dev  # 修改为dev命名空间
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
          claimName: dev-mysql-pvc  # 修改为新的PVC名称
      - name: config
        configMap:
          name: mysql-config
      - name: localtime
        hostPath:
          type: File
          path: /etc/localtime
```

- 编辑service配置：`service.yaml`:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql
  namespace: dev
  labels:
    app: mysql
spec:
  type: NodePort
  ports:
  - name: mysql
    port: 3306
    targetPort: 3306
    # nodePort: 30306  # 你可以指定一个特定的端口，或者让Kubernetes自动分配
  selector:
    app: mysql
```

- 部署mysql

```bash
kubectl apply -f config.yaml
kubectl apply -f pvc.yaml
kubectl apply -f deploy,yaml
kubectl apply -f service.yaml
kubectl get po -n dev
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
# 编辑配置文件
cd harbor
vim values.yaml
```

- 创建相关资源

```bash
kubectl create namespace harbor
# 创建ssl证书配置
kubectl create secret tls harbor-tls-secret --cert=/opt/certs/harbor.hsiangya.top.pem --key=/opt/certs/hsiangya.key -n harbor
```

- 修改PVC配置

```yaml
persistence:
  persistentVolumeClaim:
    registry:
      storageClass: "managed-nfs-storage"
    chartmuseum:
      storageClass: "managed-nfs-storage"
    jobservice:
      storageClass: "managed-nfs-storage"
    database:
      storageClass: "managed-nfs-storage"
    redis:
      storageClass: "managed-nfs-storage"
```

- 修改证书域名配置

```yaml
expose:
  type: ingress
  tls:
    enabled: true
    certSource: secret
    secret:
      secretName: harbor-tls-secret  # 使用刚才创建的 Secret 名称
  ingress:
    hosts:
      core: harbor.hsiangya.top
    annotations:
      kubernetes.io/ingress.class: nginx  
```



- 运行

```bash
helm install harbor . -f values.yaml -n harbor 
helm upgrade harbor . -f values.yaml -n harbor
helm uninstall harbor -n harbor
```

# Jenkins

## 镜像下载保存

- 通过containerd下载镜像并导出

```bash
sudo ctr images pull docker.io/jenkins/jenkins:2.430-jdk21
sudo ctr images list | grep jenkins
sudo ctr images export jenkins2.430-jdk21.tar docker.io/jenkins/jenkins:2.430-jdk21
```

- 从文件加载镜像

```bash
docker load -i jenkins2.430-jdk21.tar
```

## docker 安装

- 生成ssh

  > -C：注释说明，通常以邮箱格式，也可以为随机字符串
  >
  > 生成ssh密钥会提示输入密码：jenkins

```bash
mkdir -p /opt/jenkins/ssh
ssh-keygen -t rsa -b 4096 -C "jenkins@example.com" -f /opt/jenkins/ssh/id_rsa
chmod 700 /opt/jenkins/ssh
chmod 600 /opt/jenkins/ssh/id_rsa*
ssh-copy-id -i /opt/jenkins/ssh/id_rsa.pub user@ipaddress

# 指定私钥路径进行连接
ssh -i /opt/jenkins/ssh/id_rsa user@ipaddress
```

- 运行

```bash
sudo chown -R 1000:1000 /opt/jenkins/data
sudo chown -R 1000:1000 /opt/jenkins/ssh


docker run -d --name jenkins \
  --restart always \
  -p 8379:8080 -p 50000:50000 \
  -v /opt/jenkins/data:/var/jenkins_home \
  -v /opt/jenkins/ssh:/var/jenkins_home/.ssh \
  -u 1000:1000 \
  jenkins/jenkins:2.430-jdk21
```

## 通过k8s配置文件安装

- 权限配置文件：`permission.yaml`

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: jenkins
  namespace: jenkins
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: jenkins
rules:
  # 对 deployments 和 replicasets 的权限
  - apiGroups: ["apps"]
    resources: ["deployments", "replicasets"]
    verbs: ["create", "delete", "get", "list", "watch", "patch", "update"]
  # 对 ingresses 的权限 (注意将 extensions 改为 networking.k8s.io 对应新的 API 组)
  - apiGroups: ["networking.k8s.io"]
    resources: ["ingresses"]
    verbs: ["create", "delete", "get", "list", "watch", "patch", "update"]
  # 对 services 的权限
  - apiGroups: [""]
    resources: ["services"]
    verbs: ["create", "delete", "get", "list", "watch", "patch", "update"]
  # 对 pods 和 pods 相关操作的权限
  - apiGroups: [""]
    resources: ["pods", "pods/exec", "pods/log", "pods/portforward"]
    verbs: ["create", "delete", "get", "list", "patch", "update", "watch"]
  # 对 events 的权限
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["get", "list", "watch"]
  # 对 secrets 的权限
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: jenkins
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: jenkins
subjects:
  - kind: ServiceAccount
    name: jenkins
    namespace: jenkins
```

- 构建PVC文件：`jenkins-pvc.yaml`

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: jenkins-pvc
  namespace: jenkins
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 30Gi
  storageClassName: managed-nfs-storage
```

- deployment配置：`jenkins.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: jenkins
  namespace: jenkins
spec:
  selector:
    matchLabels:
      app: jenkins
  template:
    metadata:
      labels:
        app: jenkins
    spec:
      serviceAccount: jenkins
      initContainers:
        - name: fix-permissions
          image: busybox
          command: ["sh", "-c", "chown -R 1000:1000 /var/jenkins_home"]
          securityContext:
            privileged: true
          volumeMounts:
            - name: jenkinshome
              mountPath: /var/jenkins_home
      containers:
        - name: jenkins
          image: jenkins/jenkins
          imagePullPolicy: IfNotPresent
          env:
          - name: JAVA_OPTS
            value: -Dhudson.model.DownloadService.noSignatureCheck=true
          ports:
            - containerPort: 8080
              name: web
              protocol: TCP
            - containerPort: 50000
              name: agent
              protocol: TCP
          resources:
            limits:
              cpu: 1500m
              memory: 4096Mi
            requests:
              cpu: 1500m
              memory: 2048Mi
          readinessProbe:
            httpGet:
              path: /login
              port: 8080
            initialDelaySeconds: 60
            timeoutSeconds: 5
            failureThreshold: 12
          volumeMounts:
            - name: jenkinshome
              mountPath: /var/jenkins_home
      volumes:
        - name: jenkinshome
          persistentVolumeClaim:
            claimName: jenkins-pvc
```

- 配置service：`jenkins-service.yaml`

```yaml
apiVersion: v1
kind: Service
metadata:
  name: jenkins
  namespace: jenkins
  labels:
    app: jenkins
spec:
  selector:
    app: jenkins
  ports:
    - name: web
      port: 8080
      targetPort: web
    - name: agent
      port: 50000
      targetPort: agent
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: jenkins
  namespace: jenkins
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - jenkins.hsiangya.top
    secretName: jenkins-tls  # 指向TLS secret
  rules:
  - host: jenkins.hsiangya.top
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: jenkins
            port:
              number: 8080
```

- 部署

```bash
# 创建证书配置
kubectl create secret tls jenkins-tls \
	--cert=/opt/certs/jenkins.hsiangya.top.pem \
	--key=/opt/certs/jenkins.key \
	-n jenkins

# 部署
kubectl apply -f jenkins-permission.yaml
kubectl apply -f jenkins-pvc.yaml
kubectl apply -f jenkins.yaml
kubectl apply -f jenkins-service.yaml
```

## 基础设置

- 汉化
  1. Manage Jenkins --> Manage Plugins ---> Locale插件
  2. Manage Jenkins --> Configure System --> Locale选项输入zh_CN

- 安装git插件
  1. git
  2. gitlab
  3. Publish Over SSH
  4. Monitoring：监控Jenkins所消耗的系统资源
  5. Maven Integration
  6. ansible

- 配置信息
  1. Manage Jenkins --> Tools：配置git 等相关信息与路径
  2. Manage Jenkins --> system：配置ssh等相关信息

## 部署java项目

- 插件：Pipeline Maven Integration
- 配置相关maven、gitlab、git、ssh等

```bash
# 安装maven环境挂在到容器中
wget https://downloads.apache.org/maven/maven-3/3.8.8/binaries/apache-maven-3.8.8-bin.tar.gz
mkdir -p /opt/jenkins/maven
tar -xvzf apache-maven-3.8.8-bin.tar.gz -C /opt/jenkins/maven
```

- 打包shell

```shell
# 参数
version=$1  # 传入的版本号
app_name="project"  # 项目名称
workspace="${WORKSPACE}"  # Jenkins 工作空间，或本地路径
build_number="${BUILD_NUMBER:-manual_build}"  # Jenkins 构建号，或手动构建标识
backup_dir="${workspace}/bak/${build_number}"  # 备份目录
target_dir="${workspace}/target"  # 打包输出目录
jar_name="${app_name}.jar"  # 打包生成的 JAR 文件名

# 步骤 1: 执行 Maven 打包
echo "Executing Maven build..."
mvn clean install -Dmaven.test.skip=true -e

# 检查构建是否成功
if [ $? -ne 0 ]; then
    echo "Build failed, aborting."
    exit 1
fi

# 步骤 2: 创建备份目录
echo "Creating backup directory: $backup_dir"
if [ -d "$backup_dir" ]; then
    echo "Backup directory already exists."
else
    mkdir -p "$backup_dir"
    echo "Backup directory created."
fi

# 步骤 3: 备份打包生成的 JAR 文件
echo "Backing up JAR file to $backup_dir"
\cp -f "$target_dir/$jar_name" "$backup_dir/"
if [ $? -eq 0 ]; then
    echo "Backup completed successfully!"
else
    echo "Backup failed!"
    exit 1
fi

echo "Script execution completed."
```

-  备份shell

```shell
ReservedNum=5  # 保留的文件夹数
FileDir=${WORKSPACE}/bak/  # 备份目录
date=$(date "+%Y%m%d-%H%M%S")

# 切换到备份目录
cd $FileDir

# 获取当前目录下的文件夹数量
FileNum=$(ls -l | grep '^d' | wc -l)
echo "Current number of backup directories: $FileNum"
echo "Reserved number of backups: $ReservedNum"

# 当文件夹数量大于保留的数量时，删除最旧的文件夹
while [ "$FileNum" -gt "$ReservedNum" ]
do
    # 获取最旧的目录名
    OldFile=$(ls -rt | head -1)
    echo "$date Deleting Directory: $OldFile"
    
    if [ -n "$OldFile" ]; then  # 检查OldFile是否为空
        rm -rf "$FileDir/$OldFile"
    else
        echo "No directory found to delete."
        break
    fi
    
    # 更新文件夹数量
    FileNum=$(ls -l | grep '^d' | wc -l)
done
```

-  远程发送shell

```shell
date +"%Y-%m-%d %H:%M:%S" 
set -e

# 配置参数
APP_DIR="/path/to/project"
JAVA_CMD="nohup java"
LOG_FILE="$APP_DIR/run.log"
JAR_FILE="prject.jar"

# 切换到应用目录
cd $APP_DIR || { echo "Failed to change directory to $APP_DIR"; exit 1; }

# 检查并停止旧应用
stop_old_app() {
    echo "Searching for running instances of $JAR_FILE..."
    PIDS=$(ps -aux | grep $JAR_FILE | grep -v grep | awk '{print $2}')
    
    if [ -z "$PIDS" ]; then
        echo "No running instances of $JAR_FILE found."
    else
        for PID in $PIDS; do
            echo "Stopping application with PID $PID..."
            kill $PID || { echo "Failed to kill PID $PID"; exit 1; }
            echo "Application with PID $PID stopped."
        done
    fi
}

# 启动新应用
start_new_app() {
    if [ ! -f $JAR_FILE ]; then
        echo "File $JAR_FILE not found in $APP_DIR."
        exit 2
    fi

    for port in 10000 10001; do
        JAVA_OPTS="-Dserver.port=$port"
        LOG_FILE_PORT="${LOG_FILE%.*}_$port.log"
        echo "Starting application on port $port"
        $JAVA_CMD $JAVA_OPTS -jar $JAR_FILE > "$LOG_FILE_PORT" 2>&1 &
        echo "Application started on port $port with PID $!."
        echo "Logs are being written to $LOG_FILE_PORT."
    done
}

# 主逻辑
stop_old_app
start_new_app
```

## 部署静态文件

- 下载node,并将环境挂载到镜像中

```bash
wget https://nodejs.org/dist/v21.7.3/node-v21.7.3-linux-x64.tar.xz
tar -xvf node-v21.7.3-linux-x64.tar.xz
```



```shell
DATE=$(date +%Y-%m-%d-%H-%M-%S)
web_server="10.211.55.12"

get_code(){
	cd $WORKSPACE && \
	tar czf /opt/web-${DATE}.tar.gz ./*
}

scp_web_server(){
for hosts in $web_server
do
	scp /opt/web-${DATE}.tar.gz root@$hosts:/opt/
	ssh root@$hosts "mkdir -p /code/web-${DATE} && \
					tar -zxf /opt/web-${DATE}.tar.gz -C	/code/web-${DATE}
					rm -rf /code/web && \
					ln -s /code/web-${DATE} /code/web"
done
}

deploy(){
	get_code
	scp_web_server
}

deploy
```
