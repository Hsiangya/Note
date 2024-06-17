# sudo journalctl -u kubelet -n 100 --no-pager简介

## 组件

![image-20240607163933895](./assets/image-20240607163933895.png)

### master

![image-20240614010201848](./assets/image-20240614010201848.png)

- api-server：接口服务，基于REST风格开放k8s接口的服务
- kube-controller-manager：控制器管理器，管理各个类型的控制器，针对k8s中的各种资源进行管理
- cloud-controller-manager：云控制器管理器，第三方平台提供的控制器API对接管理功能
- kube-scheduler：调度器，负责将Pod基于一定算法，将其调用到更合适的节点（服务器）上
- etcd：理解为k8s的数据库，键值类型存储的分布式数据库，提供了基于Raft算法实现自主的集群高可用
  - 老版本：基于内存
  - 新版本：持久话存储

### node

![image-20240614010223672](./assets/image-20240614010223672.png)

- kubelet：负责Pod的生命周期、存储、网络（container runtime interface）
- kube-proxy：网络代理，负责Service的服务发现、负载均衡（4层负载）
- Pod：一个Pod中最少需要有1个容器
- container-runtime：容器运行时环境，docker、containerd、CRI-O（container runtime interface）

### 附加组件

- kube-dns：域名管理
- ingress Controller：外部服务访问控制器
- Heapster/Prometheus：提供资源监控
- Dashboard：提供GUI界面
- Federation：提供跨可用取的集群
- Fluentd-elasticsearch：日志搜索

**分层架构:**

![image-20240614010518075](./assets/image-20240614010518075.png)

- 生态系统：构建的K8s生态系统，如GUI等
- 接口层：所有生态系统都需要调用k8s的接口层
- 管理层：系统度量（如基础设施，容器和网络的度量）、自动化（如自动扩展、动态Provision等）以及策略管理（RBAC、Quota、PSP、NetworkPOlicy等）
- 应用层：部署（无状态应用、有状态应用、批处理任务、集群应用等）和路由（服务发现、DNS解析等）
- 核心层：最核心的功能，对外提供API构建高层的应用，对内提供插件式应用的执行环境

## 服务的分类

- 无状态应用：不会对本地环境产生依赖，如不会存储数据到本地磁盘
  - 优点：对客户端透明，无依赖关系，可以高校实现扩容、迁移
  - 缺点：不能存放数据，需要额外的数据服务支撑
- 有状态应用：会对本地环境产生依赖，如会存储数据到本地磁盘
  - 优点：可以独立存储数据，实现数据管理
  - 缺点：集群环境下需要实现主从、数据同步、备份、水平扩容复杂

## 资源和对象

k8s中所有内容都被抽象为"资源"，如Pod、Service、NOde等都是资源。"对象"就是资源的示例，是持久化的实体。

如某个具体的POd、某个具体的Node，kubernetes使用这些实体去表示整个集群的状态。

对象的创建、删除、修改是通过kubernetes API，也就是API server组件提供的API接口，这些是RESTful风格的Api，与k8s万物接对象理念相符。命令行工具kubectl实际上也是调用kubernetes api

k8s中的资源类别有很多种，kubectl可以通过配置文件来创建

规约：spec是规约、规格的意思，spec是必须的，他描述了对象的期望状态----希望对向所具有的特征。当创建kubernetes对象时，必须提供对象的规约，用来描述该对象的期望状态，以及关于对象的一些基本信息

状态：表示对象的实际状态，该属性由k8s自己维护，k8s会通过一系列的控制器对对应对象进行管理，让对象尽可能的让实际状态与期望状态重合。

资源清单：

![image-20240614012826540](./assets/image-20240614012826540.png)

![image-20240614012707412](./assets/image-20240614012707412.png)

![image-20240614012840603](./assets/image-20240614012840603.png)

## 资源的分类

![image-20240614014816311](./assets/image-20240614014816311.png)

### 元数据类型

对于资源的元数据的描述，每一个资源都可以使用元数据的数据：

- Horizontal Pod Autoscaler（HPA）：Pod自动扩容

  > 控制器每隔30s（可以通过horizontal-pod-autoscaler-sync-period修改）查询metrics的资源使用情况
  >
  > 支持三种metrics类型
  >
  > - 预定义metrics（比如Pod的CPU）以利用率的方式计算
  > - 自定义Pod metrics，以原始值（raw value）的方式计算
  > - 自定义object metrics
  >
  > 支持两种metrics查询方式：Heapster和自定义REST API
  >
  > 支持多metrics

- PodTemplate：关于Pod的定义，但是背包汗在其他的kubernetes对象中（如Deployment、StatefulSet、DaemonSet等控制器）‘控制器通过Pod Template信息来创建Pod

- LimitRange：对器群内的Request和LImits的配置做一个全局的统一的限制，相当于批量设置了某一个范围内（某个命名空间）的Pod的资源使用限制。

### 集群级

作用于集群之上，集群下的所有资源都可以共享使用

- Namespace
- Node：不像其他的资源，Node本质上不是kubernetes来创建的，kuberneter只是管理Node上的资源。虽然可以通过Mainfest创建一个NOde对象，但kubernetes也只是取检查是否真的是有这么一个Node，如果检查时白，也不会往上调度Pod
- ClusterRole：集群角色管理
- ClusterRoleBinding：角色与资源进行绑定，只能往集群资源级别上进行绑定

### 命名空间级

#### 工作负载型

作用在命名空间上，通常只能在该命名空间内使用

**Pod：**kubernetes中最小的可部署单元。

一个Pod(容器组)包含了一个应用程序容器（某些情况下是多个容器）、存储资源、一个唯一的网络IP地址、以及一些确定容器该如何运行的选项。Pod容器组代表了kubernetes中的一个独立的应用程序运行实例，该示例可能有单个容器或者几个紧耦合在一起的容器组成

docker是kubernetes Pod中使用最广泛的容器引擎kubernetes Pod同时也支持其他类型的容器引擎。

![image-20240614021049037](./assets/image-20240614021049037.png)![image-20240614022114140](./assets/image-20240614022114140.png)

一个pod可以被复制成多份，每一份可被称之为一个副本，这些副本除了一些描述的信息（pod名字、uid等）不一样外，其他信息都是一样的。



**控制器：**pod的控制器通常包含一个名为replicas的属性。replicas属性则指定了特定Pod的副本的数量，当当前集群中该pod的数量与该属性制定的值不一致时，k8s会采取一些策略取使得当前状态满足配置的要求

适用无状态服务：ReplicationController(RC)、ReplicaSet(RS)、Deployment。

> RC主要的作用就是用来保持容器应用的副本数始终保持在用户定义的副本数，如果有容器异常退出，会自动创建新的Pod来替代quechua的容器也会自动挥手，v1.11版本废弃
>
> kubernetes建议使用RS替代RC进行部署，RS跟RC没有本质的不同，只是名字不一样，并且RS支持集合是的selector，可以通过selector来选择对哪些Pod生效
>
> deployment针对RS的更高层次的封装，提供了更丰富的部署相关的功能：
>
> - 自动创建Replica Set/Pod
> - 滚动升级/回滚
> - 平滑扩容和缩容
> - 暂停和恢复

适用有状态服务：StatefulSet

![image-20240614025111246](./assets/image-20240614025111246.png)

> 稳定的持久存储、稳定的网络标识、有序部署、有序或栈、有序收缩、有序删除

- headless service：用于定义网络标识（DNS domain） domain name service。pod name-->访问路径（域名）自动绑定到ip上

  ![image-20240614024833043](./assets/image-20240614024833043.png)

- volumeClaimTemplate：

守护进程：Daemon

> 保证每个Node上都运行一个容器副本，常用来部署一些集群的日志、监控或者其他系统管理应用。
>
> - 日志收集，比如fluentd，logstash
> - 系统键桶，比如Prometheus Node Exporter，collectd，New Relic agent，Ganglia gmond等
> - 系统程序：比如kube-proxy，kube-dns，glusterd，ceph等

![image-20240614025843195](./assets/image-20240614025843195.png)

任务/定时任务：Job、CronJob

> Job：一次性任务，运行玩就销毁
>
> CronJob：周期性执行的任务

#### 服务发现

service：集群内网络通信，跨节点通信，负载均衡（四层负载）

ingress：实现k8s内部服务暴露给外网访问的服务

![image-20240614030620859](./assets/image-20240614030620859.png)

![image-20240614031132164](./assets/image-20240614031132164.png)

#### 配置与存储

volume：数据卷，共享Pod中容器使用的数据，用来持久化的数据，比如数据库数据

CSI：container server interface，将任意存储系统暴露给容器化应用的程序

#### 特殊类型配置

- ConfigMap：key、value类型的配置，可以将配置加载到pod中，可以将pod中的配置暴露出来，方便修改
- Secret：与ConfigMap作用相同，加密了
- DownwardAPI：提供了两种方式用于将pod信息注入到容器内，环境变量、volume挂载

#### 其他

Role：定义一组权限

RoleBinding：角色绑定

## 安装

### 环境准备

```bash
# node节点安装依赖
sudo apt update
sudo apt install openssh-server
sudo systemctl status ssh

# 生成ssh密钥，默认路径/home/hsiangya/.ssh/id_rsa  # 将密钥传递node节点
ssh-keygen -t rsa -b 4096
ssh-copy-id -i ~/.ssh/id_rsa.pub hsiangya@192.168.1.101

# 每台机器都关闭空间交换
sudo swapoff -a

# 关闭SElinux
sudo apt-get install policycoreutils
sudo setenforce 0
sudo setenforce 1
sestatus
sudo  # 永久关闭
SELINUX=disabled

# 关闭swap，找到包含swap的行 注释掉
sudo vim /etc/fstab

# 设置主机名
sudo hostnamectl set-hostname <hostname>
# 在master 添加hosts
sudo cat >> /etc/hosts <<EOF
192.168.1.7 master
192.168.1.101 node1
192.168.1.102 node2
192.168.1.104 node4
EOF

sudo cat >> /etc/hosts <<EOF
116.198.217.180 master
EOF


# 时间同步1. 
sudo apt-get update
sudo apt-get install ntpdate
sudo ntpdate time.windows.com
# 同步时间2. 阿里云
sudo apt install ntp
sudo vim /etc/ntp.conf
server ntp1.aliyun.com iburst
server ntp2.aliyun.com iburst
server ntp3.aliyun.com iburst
server ntp4.aliyun.com iburst
sudo systemctl enable ntp
sudo systemctl start ntp
ntpdate -u ntp1.aliyun.com
```

### 防火墙端口

```bash
# 开放主节点端口
sudo ufw allow 6443
sudo ufw allow 2379:2380/tcp
sudo ufw allow 10250
sudo ufw allow 10251
sudo ufw allow 10252
sudo ufw reload

# 开放工作节点端口
sudo ufw allow 10250

# 如果使用 NodePort 范围
sudo ufw allow 30000:32767/tcp
sudo ufw allow 30000:32767/udp

# 应用设置
sudo ufw enable
```

### containerd

https://github.com/containerd/containerd/blob/main/docs/getting-started.md

```bash
# 下载并解压到/usr/local
sudo wget https://github.com/containerd/containerd/releases/download/v1.7.18/containerd-1.7.18-linux-amd64.tar.gz
sudo tar Cxzvf /usr/local containerd-1.7.18-linux-amd64.tar.gz

# 下载自动启动脚本到指定位置
sudo curl https://raw.githubusercontent.com/containerd/containerd/main/containerd.service -o /etc/systemd/system/containerd.service
# 修改sandbox image
sudo containerd config default > /etc/containerd/config.toml
sudo sed -i 's#sandbox_image = "registry.k8s.io/pause:3.8"#sandbox_image = "registry.aliyuncs.com/google_containerd/pause"#g' /etc/containerd/config.toml
sudo systemctl daemon-reload
sudo systemctl enable containerd.service
sudo systemctl start containerd.service
sudo systemctl status containerd.service
containerd --version

# 安装runc
sudo wget https://github.com/opencontainers/runc/releases/download/v1.1.13/libseccomp-2.5.5.tar.gz
sudo tar xf libseccomp-2.5.5.tar.gz
cd libseccomp-2.5.5
sudo apt install gperf -y
./configure
make && make install
find / -name "libseccomp.so"
sudo wget https://github.com/opencontainers/runc/releases/download/v1.1.13/runc.amd64
sudo install -m 755 runc.amd64 /usr/local/sbin/runc
runc --version

# 安装CNI插件
sudo wget https://github.com/containernetworking/plugins/releases/download/v1.5.0/cni-plugins-linux-amd64-v1.5.0.tgz
sudo mkdir -p /opt/cni/bin
sudo tar Cxzvf /opt/cni/bin cni-plugins-linux-amd64-v1.5.0.tgz
ls /opt/cni/bin
```

### 先决条件

> 官方文档 目前未知影响

```bash
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# sysctl params required by setup, params persist across reboots
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
vm.swappiness = 0
EOF

# Apply sysctl params without reboot
sudo sysctl --system

lsmod | grep br_netfilter
lsmod | grep overlay

sysctl net.bridge.bridge-nf-call-iptables net.bridge.bridge-nf-call-ip6tables net.ipv4.ip_forward

# 配置systemdcgroup 驱动程序
sudo vim /etc/containerd/config.toml

[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
  ...
  [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    SystemdCgroup = true
```



### kubeadm、kubelet、kubectl

```bash
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl gpg
# sudo mkdir -p -m 755 /etc/apt/keyrings
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.29/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.29/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
systemctl enable kubelet
```

### 构建集群

```bash
# 查看镜像,并提前拉取
kubeadm config images list
kubeadm config images pull

sudo kubeadm init \
--apiserver-advertise-address=192.168.1.7 \
--kubernetes-version v1.29.3 \
--service-cidr=10.96.0.0/12 \
--pod-network-cidr=192.168.0.0/16 

sudo kubeadm init \
--apiserver-advertise-address=116.198.217.180 \
--image-repository registry.aliyuncs.com/google_containers \
--kubernetes-version v1.29.3 \
--service-cidr=10.96.0.0/12 \
--pod-network-cidr=192.168.0.0/16 

# 根据输出信息执行对应指令
sudo mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 根据输出信息执行对应指令将从节点加入主节点
kubeadm join 192.168.1.7:6443 --token l6s3dm.0ndg4vs7x5m85ovw \
	--discovery-token-ca-cert-hash sha256:cfafc4ef737b6931d596c9912faad7bf9580bcd350cbef57e5d7ab6c7131cd87 
```

### 网络插件

https://kubernetes.io/docs/concepts/cluster-administration/addons/

```bash
# 安装
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.28.0/manifests/tigera-operator.yaml
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.28.0/manifests/custom-resources.yaml
watch kubectl get pods -n calico-system

# 安装完成后查看节点状态
kubectl get node -o wide
```

```bash
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 198.13.42.52:6443 --token 985810.tz905wcz5hvdxmwh \
	--discovery-token-ca-cert-hash sha256:d97ab8d3009cc56032d1bfaff1c5a7e8f1f620cb15de6ae341f7565baa0087a7 

```

```bash
deb http://archive.ubuntu.com/ubuntu jammy main restricted
# deb-src http://archive.ubuntu.com/ubuntu jammy main restricted

deb http://ubuntu.mirror.constant.com jammy main restricted
# deb-src http://ubuntu.mirror.constant.com jammy main restricted

deb http://archive.ubuntu.com/ubuntu jammy-updates main restricted
# deb-src http://archive.ubuntu.com/ubuntu jammy-updates main restricted

deb http://ubuntu.mirror.constant.com jammy-updates main restricted
# deb-src http://ubuntu.mirror.constant.com jammy-updates main restricted

deb http://archive.ubuntu.com/ubuntu jammy universe
# deb-src http://archive.ubuntu.com/ubuntu jammy universe
deb http://archive.ubuntu.com/ubuntu jammy-updates universe
# deb-src http://archive.ubuntu.com/ubuntu jammy-updates universe

deb http://ubuntu.mirror.constant.com jammy universe
# deb-src http://ubuntu.mirror.constant.com jammy universe
deb http://ubuntu.mirror.constant.com jammy-updates universe
# deb-src http://ubuntu.mirror.constant.com jammy-updates universe

deb http://archive.ubuntu.com/ubuntu jammy multiverse
# deb-src http://archive.ubuntu.com/ubuntu jammy multiverse
deb http://archive.ubuntu.com/ubuntu jammy-updates multiverse
# deb-src http://archive.ubuntu.com/ubuntu jammy-updates multiverse

deb http://ubuntu.mirror.constant.com jammy multiverse
# deb-src http://ubuntu.mirror.constant.com jammy multiverse
deb http://ubuntu.mirror.constant.com jammy-updates multiverse
# deb-src http://ubuntu.mirror.constant.com jammy-updates multiverse

deb http://archive.ubuntu.com/ubuntu jammy-backports main restricted universe multiverse
# deb-src http://archive.ubuntu.com/ubuntu jammy-backports main restricted universe multiverse

deb http://ubuntu.mirror.constant.com jammy-backports main restricted universe multiverse
# deb-src http://ubuntu.mirror.constant.com jammy-backports main restricted universe multiverse

deb http://archive.ubuntu.com/ubuntu jammy-security main restricted
# deb-src http://archive.ubuntu.com/ubuntu jammy-security main restricted
deb http://archive.ubuntu.com/ubuntu jammy-security universe
# deb-src http://archive.ubuntu.com/ubuntu jammy-security universe
deb http://archive.ubuntu.com/ubuntu jammy-security multiverse
# deb-src http://archive.ubuntu.com/ubuntu jammy-security multiverse

deb http://ubuntu.mirror.constant.com jammy-security main restricted
# deb-src http://ubuntu.mirror.constant.com jammy-security main restricted
deb http://ubuntu.mirror.constant.com jammy-security universe
# deb-src http://ubuntu.mirror.constant.com jammy-security universe
deb http://ubuntu.mirror.constant.com jammy-security multiverse
# deb-src http://ubuntu.mirror.constant.com jammy-security multiverse

```

### 基于阿里云

前置准备

```bash
# 每台机器都关闭空间交换
sudo swapoff -a

# 关闭SElinux
sudo apt-get install policycoreutils
sudo setenforce 0
sudo setenforce 1
sestatus
sed -ri 's/SELINUX=enforcing/SELINUX=disabled/' /etc/selinux/config

# 关闭swap，找到包含swap的行 注释掉
sudo vim /etc/fstab
```



```bash
# 设置主机名
sudo hostnamectl set-hostname <hostname>
# 在master 添加hosts
sudo cat >> /etc/hosts <<EOF
192.168.1.7 master
192.168.1.101 node1
192.168.1.102 node2
192.168.1.104 node4
EOF

# 时间同步1. 
sudo apt-get update
sudo apt-get install ntpdate
sudo ntpdate time.windows.com
# 同步时间2. 阿里云
sudo apt install ntp
sudo vim /etc/ntp.conf
server ntp1.aliyun.com iburst
server ntp2.aliyun.com iburst
server ntp3.aliyun.com iburst
server ntp4.aliyun.com iburst
sudo systemctl enable ntp
sudo systemctl start ntp
ntpdate -u ntp1.aliyun.com
.
# 内核ip转发
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
vm.swappiness = 0
EOF
sudo sysctl --system
lsmod | grep br_netfilter
lsmod | grep overlay
sysctl net.bridge.bridge-nf-call-iptables net.bridge.bridge-nf-call-ip6tables net.ipv4.ip_forward

# 写入文件
sudo apt update
sudo apt install ipset
sudo apt install ipvsadm
cat > /etc/sysconfig/modules/ipvs.modules <<EOF
#!/bin/bash
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack
EOF
chmod 755 /etc/sysconfig/modules/ipvs.modules && bash /etc/sysconfig/modules/ipvs.modules && lsmod | grep -e ip_vs -e nf_conntrack


```

