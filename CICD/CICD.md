# 介绍

## 结构图

- 整体结构图：

![image-20240612214722483](./assets/image-20240612214722483.png)

- 代码从提交到进入生产环境整个过程：

![image-20240612222711751](./assets/image-20240612222711751.png)

![image-20240622110921019](./assets/image-20240622110921019.png)

# gitlab部署

gitlab主要针对group和project两个维度进行代码和文档管理，其中group是群组，project是工程项目，一个group可以管理多个project，一个project中可以包含多个branch，意为每个项目中有多个分支，分支之间相互独立，不同分支可以进行归并。

```bash
# 下载镜像
sudo helm repo add gitlab http://charts.gitlab.io
sudo helm repo update
sudo helm search repo gitlab
sudo helm search repo gitlab --versions

# 创建命名空间
kubectl create ns gitlab
sudo helm show values gitlab/gitlab > gitlab-values.yaml

# 创建PV目录
export ROOTDIR=/opt/k8s/gitlab
mkdir -p $ROOTDIR/data-postgresql-0-pv
mkdir -p $ROOTDIR/minio-pv
mkdir -p $ROOTDIR/prometheus-server-pv
mkdir -p $ROOTDIR/data-redis-master-0-pv
mkdir -p $ROOTDIR/repo-data-gitaly-0-pv

# 为node打上标签
kubectl label nodes master gitlab=local-pv
```

- 创建PV

```yaml
 apiVersion: v1
kind: PersistentVolume
metadata:
  name: gitlab-postgresql-0-pv
spec:
  accessModes:
    - ReadWriteOnce
  capacity:
    storage: 8Gi
  claimRef:
    apiVersion: v1
    kind: PersistentVolumeClaim
    name: data-gitlab-postgresql-0
    namespace: gitlab
  local:
    path: /opt/k8s/gitlab/data-postgresql-0-pv
  nodeAffinity:
    required:
      nodeSelectorTerms:
        - matchExpressions:
            - key: gitlab
              operator: In
              values:
                - local-pv
  persistentVolumeReclaimPolicy: Retain
  volumeMode: Filesystem

---

apiVersion: v1
kind: PersistentVolume
metadata:
  name: gitlab-minio-pv
spec:
  accessModes:
    - ReadWriteOnce
  capacity:
    storage: 10Gi
  claimRef:
    apiVersion: v1
    kind: PersistentVolumeClaim
    name: minio-pv
    namespace: gitlab
  local:
    path: /opt/k8s/gitlab/minio-pv
  nodeAffinity:
    required:
      nodeSelectorTerms:
        - matchExpressions:
            - key: gitlab
              operator: In
              values:
                - local-pv
  persistentVolumeReclaimPolicy: Retain
  volumeMode: Filesystem
---

apiVersion: v1
kind: PersistentVolume
metadata:
  name: gitlab-prometheus-server-pv
spec:
  accessModes:
    - ReadWriteOnce
  capacity:
    storage: 8Gi
  claimRef:
    apiVersion: v1
    kind: PersistentVolumeClaim
    name: gitlab-prometheus-server
    namespace: gitlab
  local:
    path: /opt/k8s/gitlab/prometheus-server-pv
  nodeAffinity:
    required:
      nodeSelectorTerms:
        - matchExpressions:
            - key: gitlab
              operator: In
              values:
                - local-pv
  persistentVolumeReclaimPolicy: Retain
  volumeMode: Filesystem

---

apiVersion: v1
kind: PersistentVolume
metadata:
  name: gitlab-redis-master-0-pv
spec:
  accessModes:
    - ReadWriteOnce
  capacity:
    storage: 8Gi
  claimRef:
    apiVersion: v1
    kind: PersistentVolumeClaim
    name: gitlab-data-redis-master-0-pv
    namespace: gitlab
  local:
    path: /opt/k8s/gitlab/data-redis-master-0-pv
  nodeAffinity:
    required:
      nodeSelectorTerms:
        - matchExpressions:
            - key: gitlab
              operator: In
              values:
                - local-pv
  persistentVolumeReclaimPolicy: Retain
  volumeMode: Filesystem

---

apiVersion: v1
kind: PersistentVolume
metadata:
  name: gitlab-repo-data-gitaly-0-pv
spec:
  accessModes:
    - ReadWriteOnce
  capacity:
    storage: 50Gi
  claimRef:
    apiVersion: v1
    kind: PersistentVolumeClaim
    name: repo-data-gitaly-0-pv
    namespace: gitlab
  local:
    path: /opt/k8s/gitlab/repo-data-gitaly-0-pv
  nodeAffinity:
    required:
      nodeSelectorTerms:
        - matchExpressions:
            - key: gitlab
              operator: In
              values:
                - local-pv
  persistentVolumeReclaimPolicy: Retain
  volumeMode: Filesystem
```

# Harbor

## 通过helm安装

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

# 加载配置
helm show values harbor/harbor > harbor-values.yaml
```

- 编辑storage class

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: local-disk
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer

# 运行
kubectl apply -f storage-class.yaml
```

- 创建PV

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: redis-pv
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  storageClassName: local-disk
  hostPath:
    path: "/mnt/data/redis"

# 创建
kubectl apply -f harbor-pv.yaml

```

- 构建

```bash
helm install harbor harbor/harbor --namespace harbor -f values.yaml

# 升级
helm upgrade harbor harbor/harbor -f values.yaml -n harbor

# 卸载
helm uninstall harbor -n harbor
```

# SonarQube

- 配置pgsql

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: postgres-data
  namespace: sonar
spec:
  accessModes:
    - ReadWriteMany
  storageClassName: "local-disk-retain"
  resources:
    requests:
      storage: 1Gi
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres-sonar
  namespace: sonar
spec:
  replicas: 1
  selector:
    matchLabels: 
      app: postgres-sonar
  template:
    metadata:
      labels: 
        app: postgres-sonar
    spec:
      containers:
      - name: postgres-sonar
        image: postgres:14.2
        ports:
        - containerPort: 5432
        env:
        - name: POSTGRES_DB
          value: "sonarDB"
        - name: POSTGRES_USER
          value: "sonarUser"
        - name: POSTGRES_PASSWORD
          value: "123456"
        volumeMounts:
          - name: data
            mountPath: /var/lib/postgresql/data
      volumes:
        - name: data
          persistentVolumeClaim:
            claimName: postgres-data
---
apiVersion: v1
kind: Service
metadata:
  name: postgres-sonar
  namespace: sonar
  labels:
    app: postgres-sonar
spec:
  type: NodePort
  ports:
  - name: postgres-sonar
    port: 5432
    protocol: TCP
  selector:
    app: postgres-sonar
```

- 配置sonar

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: sonarqube-data
  namespace: sonar
spec:
  accessModes:
    - ReadWriteMany
  storageClassName: "local-disk-retain"
  resources:
    requests:
      storage: 1Gi
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sonarqube
  namespace: sonar
  labels:
    app: sonarqube
spec:
  replicas: 1
  selector:
    matchLabels: 
      app: sonarqube
  template:
    metadata:
      labels: 
        app: sonarqube
    spec:
      initContainers:
      - name: init-sysctl
        image: busybox:1.28.4
        command: ["sysctl", "-w", "vm.max_map_count=262144"]
        securityContext:
          privileged: true
      containers:
      - name: sonarqube
        image: sonarqube
        ports:
        - containerPort: 9000
        env:
        - name: SONARQUBE_JDBC_USERNAME
          value: "sonarUser"
        - name: SONARQUBE_JDBC_PASSWORD
          value: "123456"
        - name: SONARQUBE_JDBC_URL
          value: "jdbc:postgresql://postgres-sonar:5432/sonarDB"
        livenessProbe:
          httpGet:
            path: /sessions/new
            port: 9000
          initialDelaySeconds: 30
          periodSeconds: 30
        readinessProbe:
          httpGet:
            path: /sessions/new
            port: 9000
          initialDelaySeconds: 60
          periodSeconds: 30
          failureThreshold: 6
        volumeMounts:
        - mountPath: /opt/sonarqube/conf
          name: data
        - mountPath: /opt/sonarqube/data
          name: data
        - mountPath: /opt/sonarqube/extensions
          name: data
      volumes:
      - name: data
        persistentVolumeClaim:
          claimName: sonarqube-data
---
apiVersion: v1
kind: Service
metadata:
  name: sonarqube
  namespace: sonar
  labels:
    app: sonarqube
spec:
  type: NodePort
  ports:
  - name: sonarqube
    port: 9000
    targetPort: 9000
    protocol: TCP
  selector:
    app: sonarqube

```

- 创建sonar pv

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: sonarqube-data-pv
spec:
  capacity:
    storage: 1Gi
  volumeMode: Filesystem
  accessModes:
    - ReadWriteMany
  persistentVolumeReclaimPolicy: Retain
  storageClassName: "local-disk-retain"
  local:
    path: /opt/k8s/sonarQube/sonar-data
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
          - master
```

- 创建sonar pv

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: postgres-data-pv
spec:
  capacity:
    storage: 1Gi
  volumeMode: Filesystem
  accessModes:
    - ReadWriteMany
  persistentVolumeReclaimPolicy: Retain
  storageClassName: "local-disk-retain"
  local:
    path: /opt/k8s/sonarQube/pgsql-data
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
          - master


```

- 启动：`kubectl apply -f xxx.yaml xxx.yaml`
- sonar默认帐号密码：`admin admin`

# Jenkins

## 基础概念

### 安装

```bash
# 拉取镜像
sudo docker pull jenkins/jenkins:jdk17

# 运行jenkins
sudo docker run -d \
  -u root \
  --name jenkins \
  -p 8080:8080 \
  -p 50000:50000 \
  -v /opt/cluster/jenkins/jenkins1/jenkins_home:/var/jenkins_home \
  --restart always \
  jenkins/jenkins:jdk17

# 访问8080并查看初始密码，进入之后修改访问密码
sudo cat secrets/initialAdminPassword

```

### 插件安装

- 修改插件下载源：可以提升下载速度
- 在页面上导入jenkins插件
- 在服务器上进行插件迁移

### 配置文件说明

- config.xml：主配置
- jenkins.telemetry.Correlator.xml  
- plugins：插件目录   
- secrets：密码
- users：用户配置信息
- copy_reference_file.log        
- jobs：工作流程目录                    
- secret.key 
- updates     
-  war
- hudson.model.UpdateCenter.xml 
-  nodeMonitors.xml                  
- secret.key.not-so-secret  
- userContent

### 修改时区

- Manage Jenkins-->Script console

```bash
System.setProperty('org.apache.commons.jelly.tags.fmt.timeZone','Asia/Shanghai')
```

### 构建job

