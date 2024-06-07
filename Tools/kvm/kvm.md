### 安装KVM并配置网桥模式

```bash
# 确认是否支持虚拟化，大于0就支持
egrep -c '(vmx|svm)' /proc/cpuinfo

# 安装
sudo apt update
sudo apt install bridge-utils
sudo apt-get install qemu-kvm libvirt-clients libvirt-daemon-system
sudo apt-get install virt-manager

# 创建桥接名称
sudo brctl addbr <name>
sudo brctl show
sudo brctl addif <name> <本地网络接口名(enp3s0)> 
udo brctl show

# 删除物理往开IP，把网卡接口配置到网桥上
# 应用配置
sudo ip addr del dev <本地网络接口名(enp3s0)> <ip>/24
udo ip addr add <ip>/24 dev <name>
sudo ip link set up <name>
sudo route add default gw 192.168.1.1   
iptables -A FORWARD -p all -i br0 -j ACCEPT  #开启转发

# 启动虚拟机
virt-manager

# 设置开机自启
sudo systemctl enable libvirtd
sudo systemctl start libvirtd

# 列出所有虚拟机的名称
sudo virsh list --all
# 取消某个虚拟机自启动
sudo virsh autostart --disable <vm-name>
# 查看自启虚拟机
sudo virsh list --autostart

# 编辑虚拟机中的配置文件
virsh edit <name>
# 修改网络配置文件
<interface type='bridge'>
	...
  <source bridge='br0'/>
	...
</interface>
# 重启虚拟机
virsh shutdown <name>
virsh start <name>

# 修改虚拟机网络配置
sudo vim /etc/netplan/01-netcfg.yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    enp1s0:
      dhcp4: no
  bridges:
    br0:
      interfaces: [enp1s0]
      dhcp4: no
      addresses: [192.168.1.100/24]
      nameservers:
        addresses: [8.8.8.8, 8.8.4.4]
      routes:
        - to: default
          via: 192.168.1.1
          
# 重启网络配置
sudo netplan apply
```

- 删除网卡恢复网络

```bash
sudo ip link set <name> down
sudo brctl delif <name> enp0s25
sudo ip link set enp0s25 down
sudo ip link set up enp0s25 #重启物理网卡
```



