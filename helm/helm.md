# Helm

## 概念

Helm是查找、分享和使用软件构建kubernetes的最优方式，一个kubernetes包管理器

Helm管理名为chart的kubernetes包的工具，Helm可以做以下的事情：

- 从头开始创建新的chart
- 将chart打包成归档(tgz)文件
- 与存储chart的仓库进行交互
- 在现有的kubernetes集群中安装和卸载chart
- 管理与Helm一起安装的chart的发布周期

对于Helm有三个重要的概念：

1. chart创建kubernetes应用程序所必须的一组信息

   > 将所需要的yaml等等键都管理好了 在一个目录下

2. config包含了可以合并到打包的chart中的配置信息

3. release是一个与特定配置相结合的chart的运行实例

   > release是chart运行的实例



Helm客户端：Helm客户端是终端用户的命令行客户端，负责以下内容

- 本地chart开发
- 管理仓库
- 管理发布
- 与Helm库建立接口
  - 发送安装的chart
  - 发送升级或卸载现有发布的请求

Helm库：Helm库提供执行所有Helm操作的逻辑。与kubernetes API服务交互并提供以下功能：

- 结合chart和配置来构建版本
- 将chart安装到kubernetes中，并提供后续发布对象
- 与kubernetes交互升级和卸载chart

独立的Helm库封装了Helm逻辑一边不同的客户端可以使用它

## 安装Helm

- https://github.com/helm/helm

```bash
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
```

## 常用命令

```bash
# 列出、增加、更新、删除chart仓库
helm repo 
helm repo list # 查看当前repo 列表

# 使用关键词搜索chart
helm search
helm search repo redis

# 拉取远程仓库中的chart到本地
helm pull

# 在本地创建信的chart
helm create

# 管理chart依赖
helm dependency

# 安装chart
helm install

# 列出所有release
helm list
helm list -n ingress-nginx

# 检查chart 配置是否有误
helm lint

# 打包本地chart
helm package

#  回滚release到历史版本
helm rollback

# 卸载release
helm uninstall

# 升级release
helm upgrade
```

## chart目录结构

- `chart.yaml`：
- `charts`：该目录保存其他依赖的chart（子chart）
- `templates`:chart配置模板，用于渲染最终的kubernetes YAML
  - `NOTES.txt`：用户运行helm install时候的提示信息
  - `_helpers.tpl`：用于创建模板时的提示信息
  - `deployment.yaml`：kubernetes deployment配置
  - `ingress.yaml`:kubernetes ingress配置
  - `service.yaml`:kubernetes service 配置
  - `serviceaccount.yaml`：kubernetes serviceaccount配置
  - `tests`：
    - `test-connection.yaml`
- `values.yaml`：定义chart模板中的自定义配置的默认值

## chart 安装redis

```bash
# 搜索redis 源
sudo helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo list
helm search repo redis
helm show readme bitnami/redis

# 拉取到本地
sudo helm pull bitnami/redis
sudo tar -xvf redis-19.5.5.tgz

sudo vim values.yaml
# 修改storageClass
# 设置redis的password
# 修改集群架构architecture，默认是主从（replication，3个节点），可以修改为standalone单击模式
# 修改示例存储大小persistence.size
# 修改service.nodePorts.redis 向外暴露端口，

# 创建一个redis的命名空间
kubectl create namespace redis
helm install redis_release ./redis -n redis

# 获取资源
kubectl get all -n redis
kubectl get pvc -n redis
```

## chart升级与回滚

- 利用chart升级：`helm upgrade [release] [chart] [flags]`

  > `helm upgrade redis ./redis -n redis`，制定命名空间 

- 回滚：`helm rollback <release> [revision] [flags]`

- 查看历史：`helm history <name>`

- 回退到上一个版本：`helm rollback <name>`

- 回退到制定版本：`helm rollback <name> <revision>`



# harbor

```bash
# 添加存储仓库
helm repo add harbor https://helm.goharbor.io
helm repo list

# 安装
helm install my-release harbor/harbor
# 卸载
helm uninstall my-release

# 创建一个namespace
kubectl create namespace harbor

# 加载配置
helm show values harbor/harbor > harbor-values.yaml


```

