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

