```bash
# node节点安装依赖
sudo apt update
sudo apt install openssh-server
sudo systemctl status ssh

# 生成ssh密钥，默认路径/home/hsiangya/.ssh/id_rsa  # 将密钥传递node节点
ssh-keygen -t rsa -b 4096
ssh-copy-id -i ~/.ssh/id_rsa.pub username@ipaddress
```

