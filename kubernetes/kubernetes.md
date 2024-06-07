# 简介

## 结构图

![image-20240607163933895](./assets/image-20240607163933895.png)

## 安装

- 禁止swap分区

```bash
# 关闭防火墙
sudo ufw disable

# 暂时关闭swap
sudo swapoff -a
# 永久关闭，注释包含swap的行
sudo vim /etc/fstab

# 设置主机名
hostnamectl set-hostname <hostname>
```

