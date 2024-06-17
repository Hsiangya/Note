### linux配置clash

```bash
wget -O config.yaml  订阅地址 
get -O Country.mmdb https://www.sub-speeder.com/client-download/Country.mmdb
./clash -d .

# 配置环境变量
export http_proxy="http://127.0.0.1:7890"
export https_proxy="http://127.0.0.1:7890"
export ftp_proxy="http://127.0.0.1:7890"
export no_proxy="localhost,127.0.0.1,::1"
```

