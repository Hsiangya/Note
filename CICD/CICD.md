# 介绍

## 结构图

- 整体结构图：

![image-20240612214722483](./assets/image-20240612214722483.png)

- 代码从提交到进入生产环境整个过程：

![image-20240612222711751](./assets/image-20240612222711751.png)

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

