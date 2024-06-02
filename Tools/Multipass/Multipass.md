# Multipass

MUltipass是一个轻量级的虚拟机管理工具，可以用来在本地快速创建和管理虚拟机，相比于VirtualBox或者VMware这样的虚拟机管理工具，Multipass更加轻量快速，而且它还提供了一些命令工具来方便我们管理虚拟机，官方地址（[https://Multipass.run/](https://multipass.run/)）

## 安装

```bash
# 安装
sudo snap install multipass
```

## 常用命令

常用命令，可以使用`multipass hlep`来查看

```bash
# 查看帮助
multipass help
multipass help <command>

# 创建一个名字叫k3s的虚拟机
multipass launch --name k3s
multipass launch --name k3s --cpus 2 --memory 8G --disk 10G


# 在虚拟机中执行命令
multipass exec k3s -- ls -l

# 进入虚拟机并执行shell
multipass shell k3s

# 查看虚拟机的信息
multipass info k3s

# 停止虚拟机
multipass stop k3s

# 启动虚拟机
multipass start k3s

# 删除虚拟机
multipass delete k3s

# 清理虚拟机
multipass purge

# 查看虚拟机列表
multipass list

# 挂载目录（将本地的~/kubernetes/master目录挂载到虚拟机中的~/master目录）
multipass mount ~/kubernetes/master master:~/master
```

## 配置网络

```bash
# 进入虚拟ubuntu
multipass shell systemname
```



