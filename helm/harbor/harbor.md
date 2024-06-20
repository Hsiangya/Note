# 安装

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

