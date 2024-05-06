# Flask

## Flask与django的区别

- session：django 中sessoion存在数据库中，flask中seesion存在浏览器中。
- 请求：django请求是逐一封装和传递，flask请求是利用上下文管理实现的。

## 依赖wsgi的werkzeug

```python
from werkzeug.serving import run_simple


def func(environ, start_response):
    print("请求来了")
    pass


if __name__ == '__main__':
    run_simple("127.0.0.1", 5000, func)
```

![简单使用](./assets/简单使用.png)

## 常用扩展包

扩展列表：http://flask.pocoo.org/extensions/

Flask中文文档（https://dormousehole.readthedocs.io/en/latest/）

Flask英文文档（https://flask.palletsprojects.com/en/2.1.x/）

| 扩展包           | 介绍                          |
| ---------------- | ----------------------------- |
| Flask-SQLalchemy | 操作数据库                    |
| Flask-script     | 插入脚本                      |
| Flask-migrate    | 管理迁移数据库                |
| Flask-Session    | Session存储方式指定           |
| Flask-WTF        | 表单                          |
| Flask-Mail       | 邮件                          |
| Flask-Bable      | 提供国际化和本地化支持，翻译  |
| Flask-Login      | 认证用户状态                  |
| Flask-OpenID     | 认证                          |
| Flask-RESTful    | 开发REST API的工具            |
| Flask-Bootstrap  | 集成前端Twitter Bootstrap框架 |
| Flask-Moment     | 本地化日期和时间              |
| Flask-Admin      | 简单而可扩展的管理接口的框架  |

## pip

``` bash
# 升级pip自身
pip install --upgrade pip

# 查看已经通过pip安装的包
pip list

# 显示当前已经通过pip安装的包及版本号
pip freeze

# 将pip列表和版本冻结并保存到requirement.txt
pip freeze > requirements.txt

# 安装已冻结pip列表
pip install -r requirement.txt

# 安装本地的安装包(whl文件或tar.gz压缩包路劲，或包所在URL)
pip install 目录|文件名
```

**Powershell启动脚本执行**

- 以管理员身份运行 Windows PowerShell
- 输入 Set-ExecutionPolicy RemoteSigned
- 输入Y 确认开启

**模板设定**

设定—>flask—>勾选Flask集成

设定—>flask—>模板语言—>选择jinjia2

## 运行flask

flask run命令运行的开发服务器默认会监听 http://127.0.0.1:5000/  地址 （按Crtl+C退出），并开启多线程⽀持，Flask默认使用5000端口。

```bash
# 启动flask项目
$ flask run

# 指定IP地址运行
flask run --host=0.0.0.0

# 指定端口地址运行
flask run --port=8000

# 安装python-dotenv 
pip install python-dotenv
```

如果想不需要每次启动项目都指定地址IP，可以内容保存到当前目录下的 <span style="color:red"> .flaskenv </span> 或者是 <span style="color:red"> .env </span> 文件中（需要安装 <span style="color:red">  python-dotenv</span> ）。

```bash
FLASK_DEBUG=True		# 启用开发模式
FLASK_RUN_PORT=5050			# 设置运行的端口
FLASK_RUN_HOST=0.0.0.0		# 设置监听的 ip 
FLASK_APP=app.py

production生产环境必须重新运行才会更新 开发环境会直接变换
```

开发环境（development enviroment）是指在本地编写和测试程序时的计算机环境，而生产环境（production enviroment）与开发环境相对，它指的是⽹站部署上线供用户访问时的服务器环境

根据运行环境的不同，Flask程序、扩展以及其他程序会改变相应的行为和设置。为了区分程序运行环境，Flask提供了一个FLASK_ENV环境变 量用来设置环境，**默认为 production（生产）**。在开发时，我们可以将其设 为development（开发），这会开启所有⽀持开发的特性。为了方便管理， 我们将把环境变量 FLASK_ENV 的值写入 `.flaskenv` 文件中：

## Flask Shell 

有许多操作需要在Python Shell（即Python交互式解释器）里执行。在开发Flask程序时，并不会直接使用python命令启动Python Shell，而是使用flask shell命令： 

```bash
(helloflask) $ flask shell
Python 3.6.5 |Anaconda, Inc.| (default, Mar 29 2018, 13:32:41) [MSC v.1900 64 bit (AMD64)] on win32
App: app [development]
    
Instance: C:\Users\Administrator\Desktop\helloflask\instance
>>> app
<Flask 'app'>
>>> app.name
'app'
>>>
```

和其他 flask 命令相同，执行这个命令前我们要确保程序实例可以被正常找到。

使用flask shell命令打开的Python Shell自动包含程序上下文，上下文（context）可以理解为环境。为了正常运行程序，一些操作相关的状态和数据需要被临时保存下来，这些状态和数据被统称为上下文。在Flask中，上下文有两种，分别为程序上下文和请求上下文。

# 路由

## app.route 的参数

可以自由修改传入 `app.route` 装饰器里的 URL 规则字符串，来修改访问的网站地址，但要注意以斜线 `/` 作为开头。比如：

```python
@app.route('/home')
def home():
    return 'Welcome to flask !'
```

我们之所以把传入 `app.route` 装饰器的参数称为 URL 规则，是因为我们也可以在 URL 里定义变量部分。比如下面这个视图函数会处理所有类似 `/user/<name>` 的请求：

```python
# 下面的方式，我们也可以在视图函数里获取到这个变量值

@app.route('/user/<user>')
def user_page(user):
    return 'User: %s' % user

"""
http://localhost:5000/user/greyli

http://localhost:5000/user/peter

http://localhost:5000/user/
"""
```

## 绑定多个路由 

一个视图函数也可以绑定多个 URL，通过附加多个装饰器实现，比如：

```python
@app.route('/index')
@app.route('/user')
def user():
    return 'Welcome to flask !'
"""
现在无论是访问 http://localhost:5000/home还是
http://localhost:5000/index
都可以看到返回值。
"""
```

## 路由源码流程

![路由源码流程](./assets/路由源码流程.png)

## 动态路由

Flask内置的URL变量转换器 

| 转换器 | 说 明                                                        |
| ------ | ------------------------------------------------------------ |
| string | 不包含斜线的字符串（默认值）                                 |
| int    | 整型                                                         |
| float  | 浮点数                                                       |
| path   | 包含斜线的字符串。static路由的URL规则中的filename变量就使用了这个转换器 |
| any    | 匹配一系列给定值中的一个元素                                 |
| uuid   | UUID字符串                                                   |

## 自定义转换器

如果遇到需要匹配提取 `/sms_codes/18512345678` 中的手机号数据，Flask内置的转换器就无法满足需求，此时需要自定义转换器,自定义转换器主要做3步:

```python
# 1. 创建转换器类，保存匹配时的正则表达式
from werkzeug.routing import BaseConverter

class MobileConverter(BaseConverter):
    """
    手机号格式
    """
    regex = r'1[3-9]\d{9}'  # 注意regex名字固定

# 2. 将自定义的转换器告知Flask应用
app = Flask(__name__)

# 将自定义转换器添加到转换器字典中，并指定转换器使用时名字为: mobile
app.url_map.converters['mobile'] = MobileConverter

# 3. 在使用转换器的地方定义使用
@app.route('/sms_codes/<mobile:mob_num>')
def send_sms_code(mob_num):
    return 'send sms code to {}'.format(mob_num)
```

## CBV源码流程

![CBV源码](./assets/CBV源码.png)

# 蓝图

在一个Flask 应用项目中，如果业务视图过多，可否将以某种方式划分出的业务单元单独维护，将每个单元用到的视图、静态文件、模板文件等独立分开？例如从业务角度上，可将整个应用划分为用户模块单元、商品模块单元、订单模块单元，如何分别开发这些不同单元，并最终整合到一个项目应用中？

在Flask中，使用蓝图Blueprint来分模块组织管理。蓝图实际可以理解为是一个存储一组视图方法的容器对象，其具有如下特点：

- 一个应用可以具有多个Blueprint
- 可以将一个Blueprint注册到任何一个未使用的URL下比如 “/user”、“/goods”
- Blueprint可以单独具有自己的模板、静态文件或者其它的通用操作方法，它并不是必须要实现应用的视图和函数的
- 在一个应用初始化时，就应该要注册需要使用的Blueprint

但是一个Blueprint并不是一个完整的应用，它不能独立于应用运行，而必须要注册到某一个应用中。

## 使用方式

使用蓝图可以分为三个步骤

```python
# 1. 创建一个蓝图对象
auth_bp = Blueprint('auth', __name__)

# 2. 在这个蓝图对象上进行操作,注册路由,指定静态文件夹,注册模版过滤器
@auth_bp.route('/login')
def login():
    return 'user_login'

# 3.在应用对象上注册这个蓝图对象
app.register_blueprint(auth_bp)
```

## 单文件蓝图

可以将创建蓝图对象与定义视图放到一个文件中 。

**目录（包）蓝图** 

对于一个打算包含多个文件的蓝图，通常将创建蓝图对象放到Python包的`__init__.py`文件中

```shell
--- project # 工程目录
  |------ main.py # 启动文件
  |------ auth  #用户蓝图
  |  |--- __init__.py  # 此处创建蓝图对象
  |  |--- passport.py  
  |  |--- profile.py
  |  |--- ...
  |
  |------ goods # 商品蓝图
  |  |--- __init__.py
  |  |--- ...
  |...
```

##  指定蓝图的url前缀

在应用中注册蓝图时使用`url_prefix`参数指定

```python
admin_bp = Blueprint('admin', __name__)


@admin_bp.route('/user')
def user_profile():
    return '用户后台页面'


# 给路由添加前缀
app.register_blueprint(admin_bp, url_prefix='/admin')

```

## 蓝图内部静态文件

和应用对象不同，蓝图对象创建时不会默认注册静态目录的路由。需要我们在创建时指定 static_folder 参数。

下面的示例将蓝图所在目录下的static_admin目录设置为静态目录

```python
admin_bp = Blueprint("admin", __name__, static_folder='static_admin')
app.register_blueprint(admin_bp, url_prefix='/admin')
```

现在就可以使用`/admin/static_admin/`访问`static_admin`目录下的静态文件了。

也可通过`static_url_path`改变访问路径

```python
admin = Blueprint("admin", __name__, static_folder='static_admin', static_url_path='/lib')
app.register_blueprint(admin,url_prefix='/admin')

```

## 蓝图内部模板目录

蓝图对象默认的模板目录为系统的模版目录，可以在创建蓝图对象时使用 template_folder 关键字参数设置模板目录

```python
admin_bp = Blueprint('admin', __name__, template_folder='my_templates')
```

URL是指向⽹络上资源的地址。在Flask中，我们需要让请求的URL匹配对应的视图函数，视图函数返回值就是URL对应的资源。 

## 路由匹配

为了便于将请求分发到对应的视图函数，程序实例中存储了一个路由表（app.url_map），其中定义了URL规则和视图函数的映射关系。当请求发来后，Flask会根据请求报文中的URL（path部分）来尝试与这个表中的 所有URL规则进行匹配，调用匹配成功的视图函数。如果没有找到匹配的 URL规则，Flask会自动返回 404错误响应（Not Found，表示资源未找到）。 

使用flask routes命令可以查看程序中定义的所有路由，这 个列表由app.url_map解析得到：

```bash
>>flask routes
Endpoint  Methods  Rule
--------  -------  -----------------------
index     GET      /
static    GET      /static/<path:filename>

```

在输出的文本中，我们可以看到每个路由对应的端点（Endpoint）、 HTTP方法（Methods）和URL规则（Rule），其中static端点是Flask添加的 特殊路由，用来访问静态文件。

## 静态文件配置

一个完整的⽹站当然不能只返回用户一句“Hello，World！”，我们需要模板（template）和静态文件（static file）来生成更加丰富的⽹页。**模板** 即包含程序页面的HTML文件，**静态文件** 则是需要在HTML文件中加载的 CSS 和 JavaScript 文件，以及图⽚、字体文件等资源文件。

默认情况下，模板文件存放在项目根目录中的 templates 文件夹中，静态文件存放在 static 文 件夹下，这两个文件夹需要和包含程序实例的模块处于同一个目录下，对应的项目结构示例如下所示：

```shell
helloflask/
	- templates/ 
	- static/ 
	- app.py
```

## 初始化参数

Flask 程序实例在创建的时候，需要默认传入当前 Flask 程序所指定的包(模块)，接下来就来详细查看一下 Flask 应用程序在创建的时候一些需要我们关注的参数：

- import_name
  - Flask程序所在的包(模块)，传 `__name__` 就可以
  - 其可以决定 Flask 在访问静态文件时查找的路径
- static_folder
  - 静态文件存储的文件夹，可以不传，默认为 `static` 
- template_folder
  - 模板文件存储的文件夹，可以不传，默认为 `templates` 

```python
# 初始化 Flask 应用程序，并指定当前程序所处于的包名
from flask import current_app, Flask, render_template

app = Flask(__name__,
            static_folder='static',  # 静态文件所在的文件夹名，默认为 static
            template_folder='templates',  # 模板文件所在的文件夹名，默认为 templates
		)

```

## 程序配置加载

在 Flask 程序运行的时候，可以给 Flask 设置相关配置，比如：配置数据库连接地址等等，设置 Flask 配置有以下四种方式：

```python
# 1.使用app对象加载
app.secret_key = '123456'

# 2.从配置对象中加载(常用)
app.config.form_object()

# 3.从配置文件中加载
app.config.form_pyfile()

# 1.从环境变量中加载(了解)
app.config.from_envvar()
```

下面演示如何从对象中加载配置文件

```python
class TestConfig:
    SECRET_KEY = "测试秘钥"
    

# app.secret_key = '使用app对象进行配置秘钥'
# app.config['SECRET_KEY'] = '使用配置文件进行配置秘钥'
# app.config.from_object(TestConfig)
app.config.from_pyfile('config.py')
```

## Flask与MVC架构 

MVC架构最初是用来设计桌面程序的，后来也被用于Web程序，应用 了这种架构的 Web 框架有 Django、Ruby on Rails等。

在 MVC 架构中，程序被分为三个组件：数据处理（Model）、用户界面（View）、交互逻辑 （Controller）。 严格来说，Flask并不是MVC架构的框架，因为它没有内置数据模型⽀持。为了方便表述，在本教程中，使用了 app.route（）装饰器的函数仍被称为视图函数，同时会使用“<函数名>视图”（比如index视图）的形式来代指某个视图函数。 

如果想要使用 Flask 来编写一个MVC架构的程序，那么视图函数可以作为控制器（Controller），视图（View）则是将要学习的使用 Jinja2 渲染的 HTML 模板，而模型（Model）可以使用其他库来实现，用 SQLAlchemy 来创建数据库模型。

flask 是 MVT 结构

## 自定义Flask命令

除了 Flask 内置的 flask run 等命令，我们也可以自定义命令。在虚拟环境安装Flask后，包含许多内置命令的flask脚本就可以使用了。在前面我们 已经接触了很多flask命令，比如运行服务器的 flask run，启动 shell 的 flask shell。

通过创建任意一个函数，并为其添加 app.cli.command() 装饰器，我 们就可以注册一个 flask 命令。 

```python
import click


@app.cli.command() 
def hello(): 
	click.echo('Hello, flask!')

```

函数的名称即为命令名称，这里注册的命令即hello，你可以使用flask hello命令来触发函数。作为替代，也可以在 `app.cli.command()` 装饰器 中传入参数来设置命令名称，比如 `app.cli.command('say-hello')` 会把命令 名称设置为say-hello，完整的命令即 flask say-hello 。 

借助 click 模块的 echo() 函数，我们可以在命令行界面输出字符。命令函数的文档字符串则会作为帮助信息显示（flask hello--help）。在命令 行下执行flask hello命令就会触发这个hello（）函数： 

```
$ flask hello 
Hello, flask!
```

在命令下执行flask--help可以查看Flask提供的命令帮助文档，我们自定 义的hello命令也会出现在输出的命令列表中，如下所示： 

```bash
$ flask
Usage: flask [OPTIONS] COMMAND [ARGS]...

  A general utility script for Flask applications.

  Provides commands from Flask, extensions, and the application. Loads the
  application defined in the FLASK_APP environment variable, or from a
  wsgi.py file. Setting the FLASK_ENV environment variable to 'development'
  will enable debug mode.

    > set FLASK_APP=hello.py
    > set FLASK_ENV=development
    > flask run

Options:
  --version  Show the flask version
  --help     Show this message and exit.

Commands:
  hello   自己定义的指令
  routes  Show the routes for the app.
  run     Runs a development server.
  shell   Runs a shell in the app context.
```

关于自定义命令更多的设置和功能请参考Click的官方文档 ：<https://click.palletsprojects.com/en/6.x/>

# 请求与响应



##  用户访问一个URL的整个流程

1. 浏览器便生成对应的 HTTP 请求，经由互联⽹发送到对应的Web服务器。
2. Web 服务器接收请求，通过 WSGI 将 HTTP 格式的请求数据转换成我们的 Flask 程序能够使用的 Python 数据。在程序中，Flask 根据请求的 URL 执行对应的视图函数，获取返回值生成响应。
3. 响应依次经 过 WSGI 转换生成 HTTP 响应，再经由 Web 服务器传递，最终被发出请求的客户端接收。
4. 浏览器渲染响应中包含的 HTML 和 CSS 代码，并执行 JavaScript 代码，最终把解析后的页面呈现在用户浏览器的窗口中。

## HTTP方法

HTTP通过方法来区分不同的请求类型。下表是常见的几种HTTP方法类型

| 方 法   | 说 明          |
| ------- | -------------- |
| GET     | 获取资源       |
| POST    | 传输数据       |
| PUT     | 传输文件       |
| DELETE  | 删除资源       |
| HEAD    | 获得报文首部   |
| OPTIONS | 询问支持的方法 |

## Request对象 

请求报文中的信息都可以通过request对象提供的属性和方法获取，其中常用的部分如下表所示：

| 属性    | 说明                           | 类型           |
| :------ | :----------------------------- | :------------- |
| data    | 记录请求的数据，并转换为字符串 | *              |
| form    | 记录请求中的表单数据           | MultiDict      |
| args    | 记录请求中的查询参数           | MultiDict      |
| cookies | 记录请求中的cookie信息         | Dict           |
| headers | 记录请求中的报文头             | EnvironHeaders |
| method  | 记录请求使用的HTTP方法         | GET/POST       |
| url     | 记录请求的URL地址              | string         |
| files   | 记录请求上传的文件             | *              |

获取请求URL中的查询字符串 

```python
from flask import Flask, request
app = Flask(__name__)


@app.route('/')
def index():
    # 答应请求头
    print(request.headers)
    # 查看请求拥有的方法
    print(dir(request))
    # 打印请求方法
    print(request.method)
    # 获取请求参数
    print(request.args.get('name'))
    return 'hello flask !'
```

## HTTP方法

每一个路由除了包含URL规则外，还设置了监听的HTTP方法。GET是最常用的 HTTP方法，所以视图函数默认监听的方法类型就是 GET，HEAD、 OPTIONS 方法的请求由 Flask 处理，而像 DELETE、PUT 等方法一般不会在程序中实现，在构建 Web API 时才会用到这些方法。 

我们可以在 app.route() 装饰器中使用 methods 参数传入一个包含监听 的 HTTP 方法的可迭代对象。

```python
@app.route('/hello', methods=['GET', 'POST'])
def hello(): 
	return '<h1>Hello, Flask!</h1>'

```

当某个请求的方法不符合要求时，请求将无法被正常处理。通过定义方法列表，我们可以为同一个 URL 规则定义多个视图函数， 分别处理不同 HTTP 方法的请求。

## 响应报文

响应报文主要由协议版本、状态码（status code）、原因短语（reason phrase）、响应首部和响应主体组成。

| 组成说明                                   | 响应报文内容                                                 |
| ------------------------------------------ | ------------------------------------------------------------ |
| 报文首部：状态行（协议、状态码、原因短语） | HTTP/1.0 200 OK                                              |
| 报文首部：各种首部字段                     | Content-Type: text/html; charset=utf-8<br/>Content-Length: 31<br/>Server: Werkzeug/0.15.4 Python/3.6.5<br/>Date: Wed, 03 Jul 2019 14:06:49 GMT |
| 空行                                       |                                                              |
| 报文主体                                   | <hl>Hello, flask!</h1>                                       |

- 响应报文的首部

  包含一些关于响应和服务器的信息，这些内容由Flask 生成

- 响应报文中的主体内容

  视图函数中返回的内容，浏览器接收到响应后，会把返回的响应主体解析并显示在浏览器窗口上。

## HTTP状态码

HTTP状态码用来表示请求处理的结果，常见的HTTP状态码和相应的原因短语如下表：

| 类 型        | 状态码 | 原因短语（用于解释状态码） | 说 明                                                        |
| ------------ | ------ | -------------------------- | ------------------------------------------------------------ |
| 成功         | 200    | OK                         | 请求被正常处理                                               |
|              | 201    | Created                    | 请求被处理，并创建了一个新资源                               |
|              | 204    | No Content                 | 请求处理成功，但无内容返回                                   |
| 重定向       | 301    | Moved Permanently          | 永久重定向                                                   |
|              | 302    | Found                      | 临时性重定向                                                 |
|              | 304    | Not Modified               | 请求的资源未被修改，重定向到缓存的资源                       |
| 客户端错误   | 400    | Bad Request                | 表示请求无效，即请求报文中存在错误                           |
|              | 401    | Unauthorized               | 类似403,表示请求的资源需要获取授权信息， 在浏览器中会弹出认证弹窗 |
|              | 403    | Forbidden                  | 表示请求的资源被服务器拒绝访问                               |
|              | 404    | Not Found                  | 表示服务器上无法找到请求的资源或URL无效                      |
| 服务器端错误 | 500    | Internal Server Error      | 服务器内部发生错误                                           |

## 在Flask中生成响应 

响应在 Flask 中使用 Response 对象表示，响应报文中的大部分内容由服务器处理，大多数情况下，我们只负责返回主体内容。 

视图函数可以返回最多由三个元素组成的元组：响应主体、状态码、首部字段。其中首部字段可以为字典，或是两元素元组组成的列表。Flask 会调用 make_response() 方法将视图函数返回值转换为响应对象。

### 普通响应

普通的响应可以只包含主体内容： 

```python
@app.route('/')
def hello(): 
    return '<h1>Hello, Flask!</h1>'

```

默认的状态码为200，下面指定了不同的状态码

### 元组方式响应

可以返回一个元组，这样的元组必须是 **(response, status, headers)** 的形式，且至少包含一个元素。 status 值会覆盖状态代码， headers 可以是一个列表或字典，作为额外的消息标头值。

```python
@app.route('/') 
def hello(): 
	... 
	return '<h1>Hello, Flask!</h1>', 201

```

有时你会想附加或修改某个首部字段。比如，要生成状态码为 3XX 的 

### make_response方式

如果视图函数单纯返回"****"的字符串的话，flask会自动进行一些封装生成response，让他变成浏览器可以读取的格式。使用Flask提供的make_response 方法来自定义自己的response对象

```python
from flask import make_response

@app.route('/')
def hello():
    resp = make_response('make response测试')
    resp.headers["code"] = "Python"
    resp.status = "404 not found"
    return resp
```

### 错误响应 

大多数情况下，Flask会自动处理常见的错误响应。如果想手动返回错误响应，更方便的方法是使用Flask提供的 abort() 函数。 

在 abort() 函数中传入状态码即可返回对应的错误响应

```python
from flask import Flask, abort


@app.route('/not_found')
def not_found():
    abort(404)

```

 abort() 函数前不需要使用 return 语句，但一旦 abort() 函数被调用，  abort() 函数之后的代码将不会被执行。 

## 响应格式 

常用的数据格式有纯文本、HTML、XML和JSON

### JSON 

F直接从Flask中导入 json 对象，然后调用 dumps() 方法将字典、列表或元组序列化（serialize）为 JSON 字符串，并设置正确的 MIME 类型，返回  JSON 响应 

```python
from flask import Flask, make_response, json 

... 

@app.route('/foo') 
def foo(): 
	data = { '姓名':'张三', '性别':'男' }
	response = make_response(json.dumps(data)) 
	response.mimetype = 'application/json' 
	return response

```

我们一般并不直接使用 json 模块的 dumps() 、load() 等方法， 因为Flask通过包装这些方法提供了更方便的jsonify() 函数。借助 jsonify() 函数，我们仅需要传入数据或参数，它会对我们传入的参数进行序列化，转换成JSON字符串作为响应的主体，然后生成一个响应对象， 并且设置正确的 MIME 类型。

jsonify() 函数接收多种形式的参数。你既可以传入普通参数，也可以传入关键字参数。如果你想要更直观一点，也可以像使用 dumps() 方 法一样传入字典、列表或元组，比如： 

```python
from flask import jsonify 
@app.route('/foo')
def foo(): 
	return jsonify({name: 'zhangsan', gender: 'male'})

```

会生成下面的JSON字符串： 

```python
'{"gender": "male", "name": "zhangsan"}'
```

### HTML 

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>给李四的信</title>
</head>
<body>
<h1>给李四的信</h1>
<p>亲爱的李四: </p>
<p>你好</p>
<p><strong>生活就像是强奸，如果不能反抗，那就好好享受！</strong></p>
<p>祝：工作顺利</p>
<p>x年x月x日</p>
<p>你的朋友：张三</p>
</body>
</html>

```

HTML（https://www.w3.org/html/ ）指Hypertext Markup Language（超文本标记语言），是最常用的数据格式，也是Flask返回响应的默认数据类 型。从我们在本书一开始的最⼩程序中的视图函数返回的字符串，到我们 后面会学习的HTML模板，都是HTML。当数据类型为HTML时，浏览器会 自动根据HTML标签以及样式类定义渲染对应的样式。 这也是默认的数据类型。 

## 网页重定向

视图函数的名字是自由定义的，和 URL 规则无关。和定义其他函数或变量一样，只需要让它表达出所要处理页面的含义即可。使用Flask提供的 url_for() 函数获取 URL，当路由中定义的 URL 规则被修改时，这个函数总会返回正确的 URL 。

### url_for

除此之外，它还有一个重要的作用：作为代表某个路由的端点（endpoint），同时用来生成 URL。对于程序内的 URL，为了避免手写，Flask 提供了一个 `url_for` 函数来生成 URL，它接受的第一个参数就是端点值，默认为视图函数的名称。

```python
from flask import url_for

@app.route('/test_url_for')
def test_url_for():
    return url_for('user', user='test_url_for')

```

- 路由的端点即视图函数的名称 test_url_for，调用 url_for('test_url_for') 即可 获取对应的 URL，即 `/test_url_for`。 
- 在 app.route() 装饰器中使用 endpoint 参数可以自定义端点值，不过我们通常不需要这样做。 
- 如果URL含有动态部分，那么我们需要在 url_for() 函数里传入相应的参数 

### redirect

在 Web 程序中，我们经常需要进行重定向。比如，当某个用户在没有经过认证的情况下访问需要登录后才能访问的资源，程序通常会重定向到登录页面。 

对于重定向这一类特殊响应，Flask 提供了一些辅助函数。除了像前面 那样手动生成302响应，我们可以使用 Flask提供的 redirect() 函数来生成 重定向响应，重定向的目标 URL 作为第一个参数。前面的例⼦可以简化 为：

```python
from flask import Flask, redirect 


@app.route('/baidu')
def bai_du():
    return redirect('http://www.baidu.com')

```

# Cookie 与Session

## Cookie 介绍

Cookie技术通过在请求和响应报文中添加Cookie数据来保存客户端的状态信息。

有时需要保持下来用户浏览的状态，比如用户是否登录过，浏览过哪些商品等，实现状态保持主要有两种方式：

- 在客户端存储信息使用 `Cookie` 
- 在服务器端存储信息使用 `Session` 

Cookie：指某些网站为了辨别用户身份、进行会话跟踪而储存在用户本地的数据（通常经过加密）。

- 复数形式Cookies。
- Cookie是由服务器端生成，发送给客户端浏览器，浏览器会将 Cookie 的 key/value 保存，下次请求同一网站时就发送该 Cookie 给服务器（前提是浏览器设置为启用 cookie ）。
- Cookie 的 key/value 可以由服务器端自己定义。

应用：

- 最典型的应用是判定注册用户是否已经登录网站，用户可能会得到提示，是否在下一次进入此网站时保留用户信息以便简化登录手续，这些都是 Cookie 的功用。
- 网站的广告推送，经常遇到访问某个网站时，会弹出小窗口，展示我们曾经在购物网站上看过的商品信息。
- 购物车，用户可能会在一段时间内在同一家网站的不同页面中选择不同的商品，这些信息都会写入Cookie，以便在最后付款时提取信息。

提示：

- Cookie 是存储在浏览器中的一段纯文本信息，建议不要存储敏感信息如密码，因为电脑上的浏览器可能被其它人使用
- Cookie 基于域名安全，不同域名的 Cookie 是不能互相访问的
  - 如访问 python.org 时向浏览器中写了 Cookie 信息，使用同一浏览器访问 baidu.com 时，无法访问到 python.org 写的 Cookie 信息
  - 浏览器的同源策略
- 当浏览器请求某网站时，会将本网站下所有 Cookie 信息提交给服务器，所以在 request 中可以读取 Cookie 信息

## 设置 cookie

通过操作 response 对象的 set_cookie 方法我们可以设置自定义cookie:

```python
from flask import Flask, abort, request, make_response

app = Flask(__name__)

@app.route('/set_cookie')
def set_cookie():
    response = make_response('设置 cookie')
    response.set_cookie('username', 'tom')
    return response
```

| **属 性** | **说 明**                                                    |
| --------- | ------------------------------------------------------------ |
| key       | cookie的键（名称）                                           |
| value     | cookie的值                                                   |
| max_age   | cookie被保存的时间数，单位为秒；默认在用户会话结束（即关闭浏览器）时过期 |
| expires   | 具体的过期时间，一个datetime对象或UNIX时间戳                 |
| path      | 限制cookie只在给定的路径可用，默认为整个域名                 |
| domain    | 设置cookie可用的域名                                         |
| secure    | 如果设为True,只有通过HTTPS才可以使用                         |
| httponly  | 如果设为True,禁止客户端JavaScript获取cookie                  |

设置过期时间

```python
@app.route('/set_cookie')
def set_cookie():
    response = make_response('设置 cookie')
    response.set_cookie('username', 'tom', max_age=3600)
    return response

```

## 获取cookie

可以使用 request 对象 cookies 字段的 get 方法来获取我们所需要的 cookie 

```python
request.cookies.get('username')
```

```python
@app.route('/get_cookie')
def get_cookie():
    username = request.cookies.get('username')
    return "cookie 中的名字为：" + username

```

## 删除cookie

```python
# 删除cookie
response.delete_cookie('username')

# 删除cookike
response.set_cookie('username', '', expires=0)
```



共有三种方法可以删除一个 cookie : 
(1) 可以通过在浏览器中设置来清除 cookie. 

(2) 使用 response 的 set_cookie 进行清除

(3)使用Response的 delete_cookie方法.

```python
@app.route('/del_cookie')
def del_cookie():
    response = make_response('删除cookie')
    # response.set_cookie('username', '', expires=0)
    response.delete_cookie('username')
    return response

```

## session

session 是基于 cookie 实现， 保存在服务端的键值对, 同时在浏览器中的 cookie 中也对应一相同的随机字符串，用来再次请求的时候验证；

- 对于敏感、重要的信息，建议存储在服务器端，不能存储在浏览器中，如用户名、余额、等级、验证码等信息
- 在服务器端进行状态保持的方案就是 `Session` 
- **Session依赖于Cookie**  

## session设置与删除

flask的session是通过加密之后放到了cookie中。所以有加密就有密钥用于解密，所以，只要用到了flask的session模块就一定要配置“SECRET_KEY”这个全局宏。一般设置为24位的字符。

```python
# 设置用户的session
session['username'] = 'tony'

# 获取用户的session
session_name = session.get('username')

# 删除用户session
session.pop('username', None)

# 删除用户session
del session['username']

# 删除session中的所有值
session.clear()
```





```python
from flask import Flask, session

app = Flask(__name__)
# 设置secret_key
app.secret_key = 'secret'


@app.route('/set_session')
def index1():
    session['username'] = 'tony'
    return session['username']


@app.route('/get_session')
def get_session():
    # session_name = session['username']
    session_name = session.get('username')
    return session_name


@app.route('/del_session/')
def delete():
    session.pop('username', None)
    session['username'] = False
    print(session.get('username'))
    return "删除会话"


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5050, debug=True)


```

## session过期时间

如果没有指定 session 的过期时间，那么默认是浏览器关闭后就自动结束。

```python
# 有效期延长至一个月
session.permanent = True

# 从配置中设置有效期为7天，并加载，这个值的数据类型是datetime.timedelay
from datetime import timedelta # 导入
app.config['PERMANENT_SESSION_LIFETIME'] = timedelta(days=7) # 设置时间
session.permanent = True # 开启延时
```

```python
@app.route('/del_cookie')
def del_cookie():
response = make_response('删除cookie')
# response.set_cookie('username', '', expires=0)
response.delete_cookie('username')
return response
```

# 请求钩子

在客户端和服务器交互的过程中，有些准备工作或扫尾工作需要处理，比如：

- 在请求开始时，建立数据库连接；
- 在请求开始时，根据需求进行权限校验；
- 在请求结束时，指定数据的交互格式；

为了让每个视图函数避免编写重复功能的代码，Flask提供了通用设施的功能，即请求钩子。

请求钩子是通过装饰器的形式实现，Flask支持如下四种请求钩子：

## 方法 

| 钩 子                | 说 明                                            |
| -------------------- | ------------------------------------------------ |
| before_first_request | 在处理第一个请求前执行                           |
| before_request       | 在每次请求前执行                                 |
| after_request        | 如果没有未处理的异常抛出，会在每个请求结束后运行 |
| teardown_request     | 在每次请求后执行，如果有相关错误抛出             |

这些钩⼦使用起来和 app.route() 装饰器基本相同，每个钩⼦可以注册任意多个处理函数，函数名并不是必须和钩⼦名称相同，下面是一个基本示例： 

```python
from flask import Flask
from flask import abort

app = Flask(__name__)


# 在第一次请求之前调用，可以在此方法内部做一些初始化操作
@app.before_first_request
def before_first_request():
    print("before_first_request")


# 在每一次请求之前调用，这时候已经有请求了，可能在这个方法里面做请求的校验
# 如果请求的校验不成功，可以直接在此方法中进行响应，直接return之后那么就不会执行视图函数
@app.before_request
def before_request():
    print("before_request")
    # if 请求不符合条件:
    #     return "laowang"


# 在执行完视图函数之后会调用，并且会把视图函数所生成的响应传入,可以在此方法中对响应做最后一步统一的处理
@app.after_request
def after_request(response):
    print("after_request")
    response.headers["Content-Type"] = "application/json"
    return response


# 请每一次请求之后都会调用，会接受一个参数，参数是服务器出现的错误信息
@app.teardown_request
def teardown_request(e):
    print("teardown_request")


@app.route('/')
def index():
    return 'index'


```

## 常见应用场景 

+ before_first_request

  在玩具程序中，运行程序前我们需要进行一些程序的初始化操作，比如创建数据库表，添加管理员用户。这些工作可以放 到使用before_first_request装饰器注册的函数中。 

+ before_request

  比如⽹站上要记录用户最后在线的时间，可以通过用户最后发送的请求时间来实现。为了避免在每个视图函数都添加更新在线 时间的代码，我们可以仅在使用before_request钩⼦注册的函数中调用这段代码。

+ after_request

  我们经常在视图函数中进行数据库操作，比如更新、 插入等，之后需要将更改提交到数据库中。提交更改的代码就可以放到 after_request 钩⼦注册的函数中。

在Flask程序中，客户端发出的请求触发相应的视图函数，获取返回值会作为响应的主体，最后生成完整的响应，即响应报文。

## HTTP 异常主动抛出

- abort 方法
  - 抛出一个给定状态代码的 HTTPException 或者 指定响应，例如想要用一个页面未找到异常来终止请求，可以调用 abort(404)。
- 参数：
  - code – HTTP的错误状态码

只能抛出 HTTP 协议的错误状态码

```python
# abort(404)
abort(500)

```

## 捕获错误 

- errorhandler 装饰器
  - 注册一个错误处理程序，当程序抛出指定错误状态码的时候，就会调用该装饰器所装饰的方法
- 参数：
  - code_or_exception – HTTP的错误状态码或指定异常
- 例如统一处理状态码为500的错误给用户友好的提示：

```python
# -*- coding: utf-8 -*-
from flask import Flask

app = Flask(__name__)


@app.route('/')
def index():
    return 'index'


@app.errorhandler(500)
def internal_server_error(e):
    return '服务器搬家了'


@app.errorhandler(404)
def internal_server_error(e):
    return '404 错误'


```

- 捕获指定异常

```python
@app.errorhandler(ZeroDivisionError)
def zero_division_error(e):
    return '除数不能为0'

```

## flask 上下文

Flask中有两种上下文，请求上下文和应用上下文

- 请求上下文：保存了客户端和服务器交互的数据
- 应用上下文：flask 应用程序运行过程中，保存的一些配置信息，比如程序名、数据库连接、应用信息等

| **变量名**  | 上下文类别 | 说 明                                                        |
| ----------- | ---------- | ------------------------------------------------------------ |
| current_app | 程序上下文 | 指向处理请求的当前程序实例                                   |
| g           | 程序上下文 | 替代Python的全局变量用法，确保仅在当前请求中可用。用于存储全局数据，每次请求都会重设 |
| request     | 请求上下文 | 封装客户端发出的请求报文数据                                 |
| session     | 请求上下文 | 用于记住请求之间的数据，通过签名的Cookie实现                 |

### 请求上下文(request context)

在 flask 中，可以直接在视图函数中使用 **request** 这个对象进行获取相关数据，而 **request** 就是请求上下文的对象，保存了当前本次请求的相关数据，请求上下文对象有：request、session

- request
  - 封装了HTTP请求的内容，针对的是 http 请求。举例：user = request.args.get('user')，获取的是get请求的参数。
- session
  - 用来记录请求会话中的信息，针对的是用户信息。举例：session['name'] = user.id，可以记录用户信息。还可以通过 session.get('name') 获取用户信息。

### 应用上下文(application context)

它的字面意思是 应用上下文，但它不是一直存在的，它只是 request context 中的一个对 app 的代理(人)，所谓local proxy。它的作用主要是帮助 request 获取当前的应用，它是伴 request 而生，随 request 而灭的。

应用上下文对象有：current_app，g

### current_app

应用程序上下文,用于存储应用程序中的变量，可以通过 current_app.name 打印当前app的名称，也可以在current_app中存储一些变量

`current_app` 就是当前运行的flask app，在代码不方便直接操作flask的app对象时，可以操作`current_app`就等价于操作flask app对象

**示例**

```python
# -*- coding: utf-8 -*-
from flask import Flask, current_app

app1 = Flask(__name__)
app2 = Flask(__name__)

# 以 secret_key 对象为例
# 为了方便在各个视图中使用，将创建的 secret_key 变量保存到flask app中，
# 后续可以在视图中使用 current_app.secret_key 获取
app1.secret_key = 'app1 秘钥'
app2.secret_key = 'app2 秘钥'


@app1.route('/')
def route11():
    # current_app 是当前请求的对象，是app对象的一个代理对象，只能读取数据不能修改
    return current_app.secret_key


@app2.route('/')
def route12():
    return current_app.secret_key
```

运行

```shell
export FLASK_APP=FLASK_APP=08context.py:app1
flask run

```

- 访问`/` 显示`app1 秘钥`

```shell
export FLASK_APP=current_app_demo:app2
flask run

```

- 访问`/` 显示`app2 秘钥`

### 变量 g

g 作为 flask 程序全局的一个临时变量，充当中间媒介的作用，我们可以通过它在一次请求调用的多个函数间传递一些数据。每次请求都会重设这个变量。

大的区别是，session 对象是可以跨 request 的，只要 session 还未失效，不同的 request 的请求会获取到同一个session，但是 g 对象不是，g 对象不需要管过期时间，请求一次就 g 对象就改变了一次，或者重新赋值了一次，不同的请求，会有不同的全局变量。

```python
from flask import Flask, g

app = Flask(__name__)

def db_query():
    user_id = g.user_id
    user_name = g.user_name
    print('user_id={} user_name={}'.format(user_id, user_name))

@app.route('/')
def get_user_profile():
    g.user_id = 123
    g.user_name = 'info'
    db_query()
    return 'hello world'

```

## 请求钩子综合案例

### 需求

（无前端页面）

- 构建认证机制
- 对于特定视图可以提供强制要求用户登录的限制
- 对于所有视图，无论是否强制要求用户登录，都可以在视图中尝试获取用户认证后的身份信息

```python
from flask import Flask, abort, g

app = Flask(__name__)

@app.before_request
def authentication():
    """
    利用before_request请求钩子，在进入所有视图前先尝试判断用户身份
    :return:
    """
    # TODO 此处利用鉴权机制（如cookie、session、jwt等）鉴别用户身份信息
    # if 已登录用户，用户有身份信息
    g.user_id = 123
    # else 未登录用户，用户无身份信息
    # g.user_id = None

def login_required(func):
    def wrapper(*args, **kwargs):
        if g.user_id is not None:
            return func(*args, **kwargs)
        else:
            abort(401)

    return wrapper

@app.route('/')
def index():
    return 'home page user_id={}'.format(g.user_id)

@app.route('/profile')
@login_required
def get_user_profile():
    return 'user profile page user_id={}'.format(g.user_id)

```

# 模板

## 模板语法

| 语法        | 介绍                                   |
| ----------- | -------------------------------------- |
| `{{ ... }}` | 用来标记变量                           |
| `{% ... %}` | 用来标记语句，比如 if 语句，for 语句等 |
| `{# ... #}` | 用来写注释                             |

模板中使用的变量需要在渲染的时候传递进去

### 使用变量

代码中传入字符串，列表，字典到模板中

```python
@app.route('/args')
def args():
    my_str = 'hello'
    my_int = 10
    my_arr = [1, 2, 3, 4, 5]
    my_dict = {
        "name": "Hsiangya",
        'age': 18
    }
    return render_template(
        '0103args.html',
        my_str=my_str,
        my_int=my_int,
        my_arr=my_arr,
        my_dict=my_dict,
    )
```

**模板中代码** 

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>模板变量</title>
</head>
<body>
<h1>我的模板变量</h1>
<div>str：{{ my_str }}</div>
<div>int：{{ my_int }}</div>
<div>arr：{{ my_arr }}</div>
<div>dict：{{ my_dict }}</div>
</body>
</html>

```

**相关运算，取值** 

```html
<h1>相关运算</h1>
<div>str + str：{{ my_str + ' world !' }}</div>
<div>int + int：{{ my_int + 100}}</div>
<div>arr[1]：{{ my_arr[1] }}</div>
<div>arr[1:]：{{ my_arr[1:] }}</div>
<div>dict['name']：{{ my_dict['name'] }}</div>
<div>my_dict.items()：{{ my_dict.items() }}</div>

```

### 条件语句

在渲染过程中使用实际值替换占位符，只是Jinja2在模板文件中支持的诸多强大操作之一。 模板也支持在`{%...％}`块内使用控制语句。 *temp_demo3.html*模板的下一个版本添加了一个条件语句：

```jinja2
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    {% if username %}
        <title>个人主页 - {{ username }}</title>
    {% else %}
        <title>欢迎来到个人主页!</title>
    {% endif %}
</head>
<body>
<h1>Hello, {{ username }}!</h1>
</body>
</html>

```

```python
@app.route('/if')
def demo_if():
    # username = request.args.get('username')
    # if not username:
    #     username = 'Hsiangya'
    return render_template(
        '0202if.html',
        # username=username,
    )

```

### 循环语句

for循环中的参数，在循环内部,你可以使用一个叫做 loop 的特殊变量来获得关于 for 循环的一些信息

| 变量           | 描述                                  |
| -------------- | ------------------------------------- |
| loop.index     | 当前循环迭代的次数（从 1 开始）       |
| loop.index0    | 当前循环迭代的次数（从 0 开始）       |
| loop.revindex  | 到循环结束需要迭代的次数（从 1 开始） |
| loop.revindex0 | 到循环结束需要迭代的次数（从 0 开始） |
| loop.first     | 如果是第一次迭代，为 True 。          |
| loop.last      | 如果是最后一次迭代，为 True 。        |
| loop.length    | 序列中的项目数。                      |
| loop.cycle     | 在一串序列间期取值的辅助函数。        |

```python
@app.route('/loop')
def demo_loop():
    my_array = ['苹果', '橘子', '西瓜']
    return render_template('0203loop.html',
                           arr=my_array)


```

Jinja2提供了 `for` 控制结构来应对这类问题：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>循环</title>
</head>
<body>
<h3>列表取值</h3>
<div>我爱吃 {{ arr[0] }}!</div>
<div>我爱吃 {{ arr[1] }}!</div>
<div>我爱吃 {{ arr[2] }}!</div>
<h3>循环遍历</h3>
{% for i in arr %}
    <div>我爱吃 {{ i }}!</div>
{% endfor %}
<h3>循环特殊变量</h3>
{% for post in arr %}
    <div>{{ loop.index }}, {{ post }}</div>
{% endfor %}
</body>
</html>

```

### 案例

```python
@app.route('/edit')
def demo_edit():
    student = {
        'name': 'hsiangya',
        'level': '二年级',
        'class': '2班',
        'math': 75,
        'chinese': 75,
        'english': 60,
    }
    level = [
        {'id': 1, 'name': '一年级'},
        {'id': 2, 'name': '二年级'},
        {'id': 3, 'name': '三年级'},
        {'id': 4, 'name': '四年级'},
        {'id': 5, 'name': '五年级'},
        {'id': 6, 'name': '六年级'},
    ]
    _class = [
        {'id': 1, 'name': '1班'},
        {'id': 2, 'name': '2班'},
        {'id': 3, 'name': '3班'},
    ]
    return render_template(
        '0204demo_edit.html',
        student=student,
        level=level,
        _class=_class
    )

```



```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>循环</title>
    <link rel="stylesheet" href="/static/css/layui.css">
</head>
<body>
<form class="layui-form" action="">
    <div class="layui-form-item">
        <label class="layui-form-label">用户名</label>
        <div class="layui-input-block">
            <input type="text" name="name" required
                   lay-verify="required" placeholder="请输入用户名"
                   autocomplete="off" class="layui-input"
            >
        </div>
    </div>
    <div class="layui-form-item">
        <label class="layui-form-label">年级</label>
        <div class="layui-input-block">
            <select name="level" lay-verify="required">
                <option value="1">一年级</option>
                <option value="2">二年级</option>
                <option value="3">三年级</option>
            </select>
        </div>
    </div>

    <div class="layui-form-item">
        <label class="layui-form-label">班级</label>
        <div class="layui-input-block">
            <input type="radio" name="class" value="1" title="1班">
            <input type="radio" name="class" value="2" title="2班">
            <input type="radio" name="class" value="3" title="3班">
        </div>
    </div>
    <div class="layui-form-item">
        <label class="layui-form-label">数学</label>
        <div class="layui-input-block">
            <input type="text" name="math" required
                   lay-verify="required" placeholder="请输入数学成绩"
                   autocomplete="off" class="layui-input"
            >
        </div>
    </div>
    <div class="layui-form-item">
        <label class="layui-form-label">语文</label>
        <div class="layui-input-block">
            <input type="text" name="math" required
                   lay-verify="required" placeholder="请输入语文成绩"
                   autocomplete="off" class="layui-input"
            >
        </div>
    </div>
    <div class="layui-form-item">
        <label class="layui-form-label">英语</label>
        <div class="layui-input-block">
            <input type="text" name="math" required
                   lay-verify="required" placeholder="请输入英语成绩"
                   autocomplete="off" class="layui-input"
            >
        </div>
    </div>
    <div class="layui-form-item">
        <div class="layui-input-block">
            <button class="layui-btn" lay-submit lay-filter="formDemo">立即提交</button>
            <button type="reset" class="layui-btn layui-btn-primary">重置</button>
        </div>
    </div>
</form>

<script src="/static/layui.js"></script>
<script>
  layui.use(function() {
    var form = layui.form;

    // 监听提交
    form.on('submit(formDemo)', function(data) {
      layer.msg(JSON.stringify(data.field));
      return false;
    });
  });
</script>
</body>
</html>
```

## 过滤器

### 使用方式



```python
# 过滤器的使用方式为：变量名 | 过滤器。
{{variable | filter_name(*args)}}

# 如果没有任何参数传给过滤器,则可以把括号省略掉
{{variable | filter_name}}
```

### 链式调用

在 jinja2 中，过滤器是可以支持链式调用的，示例如下：

```python
{{ "hello world" | reverse | upper }}
```

### 字符串操作

| 方法       | 简介                                     |
| ---------- | ---------------------------------------- |
| safe       | 禁用转义                                 |
| truncate   | 字符串截断                               |
| striptags  | 渲染之前把值中所有的HTML标签都删掉       |
| capitalize | 把变量值的首字母转成大写，其余字母转小写 |
| lower      | 把值转成小写                             |
| upper      | 把值转成大写                             |
| title      | 把值中的每个单词的首字母都转成大写       |
| reverse    | 字符串反转                               |
| format     | 格式化输出                               |

```python
# 禁止转义
<p>{{ '<b>hello</b>' | safe }}</p>

# 字符串截断
<p>{{ 'hello every one' | truncate(9)}}</p>

# 去掉所有HTML表亲啊
<p>{{ '<em>hello</em>' | striptags }}</p>

# 把变量值的首字母转成大写，其余字母转小写
<p>{{ 'hello' | capitalize }}</p>

# 把值转成小写
<p>{{ 'HELLO' | lower }}</p>

# 把值转成大写
<p>{{ 'hello' | upper }}</p>

# 把值中的每个单词的首字母都转成大写
<p>{{ 'hello' | title }}</p>

# 字符串反转
<p>{{ 'olleh' | reverse }}</p>

# 格式化输出
<p>{{ '%s is %d' | format('name',17) }}</p>
```

### 列表操作

| 方法   | 简介           |
| ------ | -------------- |
| first  | 取第一个元素   |
| last   | 取最后一个元素 |
| length | 获取列表长度   |
| sum    | 列表求和       |
| sort   | 列表排序       |

```html
# 取第一个元素
<p>{{ [1,2,3,4,5,6] | first }}</p>

# 取最后一个元素
<p>{{ [1,2,3,4,5,6] | last }}</p>

# 获取列表长度
<p>{{ [1,2,3,4,5,6] | length }}</p>

# 列表求和
<p>{{ [1,2,3,4,5,6] | sum }}</p>

# 列表排序
<p>{{ [6,2,3,1,5,4] | sort }}</p>
```

## 宏

定义宏

```html
{#定义宏，相当于定义一个函数，在使用的时候直接调用该宏，传入不同的参数就可以了#}
{% macro input(label="", type="text", name="", value="") %}
<label>{{ label }}</label><input type="{{ type }}" name="{{ name }}" value="{{ value }}">
{% endmacro %}

```

使用宏

```html
<form>
    {{ input("用户名：", name="username") }}<br/>
    {{ input("身份证号：", name="idcard") }}<br/>
    {{ input("密码：", type="password", name="password") }}<br/>
    {{ input("确认密码：", type="password", name="password2") }}<br/>
    {{ input(type="submit", value="注册") }}
</form>

```

对宏(macro)的理解：

1. 把它看作 Jinja2 中的一个函数，它会返回一个模板或者 HTML 字符串
2. 为了避免反复地编写同样的模板代码，出现代码冗余，可以把他们写成函数以进行重用
3. 需要在多处重复使用的模板代码片段可以写入单独的文件，再包含在所有模板中，以避免重复

把宏单独抽取出来，封装成html文件，其它模板中导入使用，文件名可以自定义macro.html

```html
{% macro function(type='text', name='name', value='张三') %}
<input type="{{type}}" name="{{name}}" value="{{value}}" class="form-control">
{% endmacro %}
```

在其它模板文件中先导入，再调用

```python
{% import 'macro.html' as func %}
{{ func.function() }}
```

## 模板的继承

Jinja2 有一个模板继承特性，专门解决这个问题。从本质上来讲，就是将所有模板中相同的部分转移到一个基础模板中，然后再从它继承过来。

定义一个名为 `base.html` 的基本模板，其中包含一个简单的导航栏

```html
<!DOCTYPE html>
<html lang="en">
<head>
    {% if username %}
        <title>个人主页 - {{ username }}</title>
    {% else %}
        <title>欢迎来到个人主页</title>
    {% endif %}
</head>
<body>

{% block content %}{% endblock %}

</body>
</html>

```

在这个模板中，我使用 `block` 控制语句来定义派生模板可以插入代码的位置。 *block* 被赋予一个唯一的名称，派生的模板可以在提供其内容时进行引用。

通过从基础模板 *base.html* 继承 HTML 元素，我现在可以简化模板 *index.html* 了：

```python
{% extends "base.html" %}

{% block content %}
    {% for i in username %}
            <h1>Hello, {{ i }}!</h1>
    {% endfor %}
{% endblock %}

```

自从基础模板 *base.html* 接手页面的布局之后，我就可以从 *index.html* 中删除所有这方面的元素，只留下内容部

分。 `extends` 语句用来建立了两个模板之间的继承关系，这样 Jinja2 才知道当要求呈现 `index.html` 时，需要将其嵌入到 `base.html` 中。 而两个模板中匹配的`block`语句和其名称 `content` ，让Jinja2知道如何将这两个模板合并成在一起。 现在，扩展应用程序的页面就变得极其方便了，我可以创建从同一个基础模板*base.html*继承的派生模板，这就是我让应用程序的所有页面拥有统一外观布局而不用重复编写代码的秘诀。

模板继承是为了重用模板中的公共内容。一般Web开发中，继承主要使用在网站的顶部菜单、底部。这些内容可以定义在父模板中，子模板直接继承，而不需要重复书写。

- 标签定义的内容

```python
{% block top %} {% endblock %}
```

- 相当于在父模板中挖个坑，当子模板继承父模板时，可以进行填充。
- 子模板使用 extends 指令声明这个模板继承自哪个模板
- 父模板中定义的块在子模板中被重新定义，在子模板中调用父模板的内容可以使用 super()

### 父模板

- base.html

```jinja2
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    {% if title %}
        <title>个人主页 - {{ username }}</title>
    {% else %}
        <title>欢迎来到个人主页</title>
    {% endif %}
    <style>
        div {
            border: 2px solid;
            height: 100px;

        }

        .nav {
            height: 40px;
        }

        .top {
            background-color: pink;
        }

        .content {
            background-color: pink;
        }

        .bottom {
            background-color: pink;
        }
    </style>
</head>
<body>
{% block top %}
    <div class="top">base 顶部菜单</div>
{% endblock top %}
<br>
{% block content %}
    <dic class="content">base 内容部分</dic>
{% endblock content %}
<br>
{% block bottom %}
    <div class="bottom">base 底部</div>
{% endblock bottom %}
</body>
</html>

```

### 子模板

- extends指令声明这个模板继承自哪

```python
{% extends 'base.html' %}
{% block content %}
 需要填充的内容
{% endblock content %}

```

- 模板继承使用时注意点：
  - 不支持多继承
  - 为了便于阅读，在子模板中使用 extends 时，尽量写在模板的第一行。
  - 不能在一个模板文件中定义多个相同名字的 block 标签。
  - 当在页面中使用多个 block 标签时，建议给结束标签起个名字，当多个 block 嵌套时，阅读性更好。

### 包含

Jinja2 模板中，除了宏和继承，还支持一种代码重用的功能，叫包含(Include)。它的功能是将另一个模板整个加载到当前模板中，并直接渲染。

- templates/header.html 的使用

```python
{% include 'header.html' %}
```

文件：header.html

```
<div class="nav">
    头部部分
</div>
```

包含在使用时，如果包含的模板文件不存在时，程序会抛出 **TemplateNotFound** 异常，可以加上 `ignore missing` 关键字。如果包含的模板文件不存在，会忽略这条include语句。

- include 的使用加上关键字ignore missing

```python
{% include 'header.html' ignore missing %}
```

**小结** 

- 宏(Macro)、继承(Block)、包含(include)均能实现代码的复用。
- 继承(Block)的本质是代码替换，一般用来实现多个页面中重复不变的区域。
- 宏(Macro)的功能类似函数，可以传入参数，需要定义、调用。
- 包含(include)是直接将目标模板文件整个渲染出来。

在模板中，可能会遇到以下情况：

- 多个模板具有完全相同的顶部和底部内容
- 多个模板中具有相同的模板代码内容，但是内容中部分值不一样
- 多个模板中具有完全相同的 html 代码块内容

## 生成静态文件 URL

在 HTML 文件里，引入这些静态文件需要给出资源所在的 URL。为了更加灵活，这些文件的 URL 可以通过 Flask 提供的 `url_for()` 函数来生成。

**由 url_for 生成静态url**

templates/static.html：引入 CSS 文件

```html
<link rel="stylesheet" href="{{ url_for('static', filename='style.css') }}" type="text/css">

```

最后要为对应的元素设置 `class` 属性值，以便和对应的 CSS 定义关联起来：

*templates/static.html：添加 class 属性*

```html
<img alt="user_ico" class="user_ico" src="{{ url_for('static', filename='a.jpg') }}">
```

## 自定义错误页面

错误处理函数和视图函数很相似，返回值将会作为响应的主体。并在其中为最常见的404和500错误创建了模板文件，表示404 页面的 404.html 模板内容。

### 404页面模板 

错误处理函数需要附加app.errorhandler()装饰器，并传入错误状态 码作为参数。当发生错误时，对应的错误处理函数会被调用，它的返回值会作为错误响应的主体。

```python
from flask import Flask, render_template, flash

app = Flask(__name__)
app.secret_key = '123456'


@app.errorhandler(404)
def error404(e):
    flash('请重新刷新页面试试')
    flash('刷新也没用？检查一下代码哪里错了')
    flash('代码没错？重启pycharm试试')
    flash('重启pycharm没有？重启电脑试试')
    flash('重启电脑没用？重装系统试试')
    return render_template('404.html'), 404


```

### HTTP 异常类的常用属性

Werkzeug 内置的 HTTP 异常类的常用属性，错误处理函数接收异常对象作为参数，内置的异常对象提供了下列常用属性

| 属 性           | 说 明                                                        |
| --------------- | ------------------------------------------------------------ |
| **code**        | 状态码                                                       |
| **name**        | 原因短语                                                     |
| **description** | 错误描述，另外使用 get_description() 方法还可以获取 HTML 格式的错误 描述代码 |

## 自定义过滤器

如果内置的过滤器不能满足你的需要，还可以添加自定义过滤器。使用 app.template_filter（）装饰器可以注册自定义过滤器

- 通过Flask应用对象的 **add_template_filter** 方法
- 通过装饰器来实现自定义过滤器

**重要：自定义的过滤器名称如果和内置的过滤器重名，会覆盖内置的过滤器。**

### 需求：添加列表反转的过滤器

#### 方式一

通过调用应用程序实例的 add_template_filter 方法实现自定义过滤器。该方法第一个参数是函数名，第二个参数是自定义的过滤器名称：

```python
@app.template_filter('sort_reverse')
def sort_reverse(li):
    # 通过原列表创建一个新列表
    temp_li = list(li)
    # 将新列表进行返转
    temp_li.reverse()
    return temp_li
```

#### 方式二

用装饰器来实现自定义过滤器。装饰器传入的参数是自定义的过滤器名称。

```python
def sort_reverse(li):
    # 通过原列表创建一个新列表
    temp_li = list(li)
    # 将新列表进行返转
    temp_li.reverse()
    return temp_li

app.add_template_filter(do_listreverse, 'lireverse')

```

在 html 中使用该自定义过滤器

```html
<br/> my_array 原内容：{{ [3, 4, 2, 1, 7, 9]  }}
<br/> my_array 反转：{{ [3, 4, 2, 1, 7, 9]  | sort_reverse }}
```

运行结果

```
my_array 原内容：[3, 4, 2, 1, 7, 9] 
my_array 反转：[9, 7, 1, 2, 4, 3]
```

## Jinja2常用内置过滤器

| 过滤器                                                       | 说 明                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| default (value, default_value=u", boolean=False)             | 设置默认值，默认值作为参数传入，别名为d                      |
| escape(s)                                                    | 转义HTML文本，别名为e                                        |
| first (seq)                                                  | 返回序列的第一个元素                                         |
| last(seq)                                                    | 返回序列的最后一个元素                                       |
| length(object)                                               | 返回变量的长度                                               |
| random(seq)                                                  | 返回序列中的随机元素                                         |
| safe( value)                                                 | 将变量值标记为安全，避免转义                                 |
| trim(value)                                                  | 清除变量值前后的空格                                         |
| max(value, case_sensitive=False, attribute=None)             | 返回序列中的最大值                                           |
| min(value, case_sensitive=False, attribute=None)             | 返回序列中的最小值                                           |
| unique(value, case_sensitive=F alse, attribute=None)         | 返回序列中的不重复的值                                       |
| striptags(value)                                             | 清除变量值内的HTML标签                                       |
| urlize (value, trim_url_limit=None, nofbllow=False, target=None, rel=None) | 将URL文本转换为可单击的HTML链接                              |
| wordcount (s)                                                | 计算单词数量                                                 |
| tojson(value, indent=None)                                   | 将变量值转换为JSON格式                                       |
| truncate(s, length=255, killwords=False,end-...', leeway=None) | 截断字符串，常用于显示文章摘要，length参数设置截断的长度, killwords参数设置是否截断单词，end参数设置结尾的符号 |

这里只列出了一部分常用的过滤器，完整的列表请 [访问]( http://jinja.pocoo.org/docs/2.10/templates/#builtin-filters) 查看。

# 数据库操作(ORM)

## Flask-SQLAlchemy

扩展 Flask-SQLAlchemy 集成了SQLAlchemy，它简化了连接数据库服务器、管理数据库操作会话等各类工作，让Flask中的数据处理体验变得更加轻松。

```bash
# 安装依赖
pip install flask-sqlalchemy 
```

```python
from flask import Flask 

# 导入
from flask_sqlalchemy import SQLAlchemy 
app = Flask(__name__) 

# 实例化SQLAlchemy对象
db = SQLAlchemy(app) 
```

## 连接数据库服务器  

URI代表统一资源标识符，是用来标示资源的一组字符串。URL是它 的⼦集。在大多数情况下，这两者可以交替使用。 下表是一些常用的DBMS及其数据库URI格式

| DBMS             | URI                                                          |
| ---------------- | ------------------------------------------------------------ |
| PostgreSQL       | postgresql://username:password@host/databasename             |
| MySQL            | mysql+pymysql://username:password@host/databasename          |
| Oracle           | oracle ://username:password@host:port/sidname                |
| SQLite (UNIX)    | sqlite:////absolute/path/to/foo.db                           |
| SQLite (Windows) | sqlite:///absolute\\path\\to\\foo.db 或 r'sqlite:///absolute\path\to\foo.db', |
| SQlite (内存型)  | sqlite:///或 sqlite:///:memory:                              |

## 字段类型 

| **字**	**段** | **说**	**明**                             |
| ---------------- | -------------------------------------------- |
| Integer          | 整数                                         |
| String           | 字符串，可选参数length可以用来设置最大长度   |
| Text             | 较长的Unicode文本                            |
| Date             | 日期，存储Python的datetime.date对象          |
| Time             | 时间，存储Python的datetime.time对象          |
| DateTime         | 时间和日期，存储Python的datetime对象         |
| Interval         | 时间间隔，存储Python的datetime.timedelta对象 |
| Float            | 浮点数                                       |
| Boolean          | 布尔值                                       |
| PickleType       | 存储Pickle列化的Python对象                   |
| LargeBinary      | 存储任意二进制数据                           |

```python
class Students(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(length=20))
    math = db.Column(db.Integer)
    chinese = db.Column(db.Integer)
    english = db.Column(db.Integer)。 
```

## 常用的字段参数 

| 参数名      | 说	明                                          |
| ----------- | ------------------------------------------------- |
| primary key | 如果设为True,该字段为主键                         |
| unique      | 如果设为True,该字段不允许出现重复值               |
| index       | 如果设为True，为该字段创建索引，以提高査询效率    |
| nullable    | 确定字段值可否为空，值为True或False，默认值为True |
| default     | 为字段设置默认值                                  |

在实例化字段类时，通过把参数primary_key设为True可以将其定义为主键。在我们定义的Note类中，id字段即表的主键（primary key）。主键是每一条记录（行）独一无二的标识，也是模型类中必须定义的字段，一 般命名为id或pk。 

## 创建表  create_all()

这通过对db 对象调用 create_all() 方法实现： 

数据库和表一旦创建后，之后对模型的改动不会自动作用到实际的表中。如果要使改动生效，最简单的方式是调用 db.drop_all() 方法删除数据库和表，然后再调用 db.create_all() 方法创建

```python
import click 


@app.cli.command()
def init_db():
    """创建数据库"""
    db.create_all()
    click.echo('初始化数据库。')： 
```

## 数据库操作 

Flask-SQLAlchemy 自动帮我们创建会话，可以通过 db.session 属性获取

Flask- SQLAlchemy 会自动为模型类生成一个 `__repr__()` 方法。 当在 Python Shell 中调用模型的对象时，`__repr__()` 方法会返回一条类 似“<模型类名主键值>”的字符串，比如 <Note 2>。

### insert

添加一条新记录到数据库主要分为三步：

1. 创建Python对象（实例化模型类）作为一条记录。 

2. 添加新创建的记录到数据库会话。 

3. 提交数据库会话。 

| 方法                 | 介绍             |
| -------------------- | ---------------- |
| db.session.add()     | 一次添加一条记录 |
| db.session.add_all() | 一次添加多条记录 |

```python
@app.cli.command()
def create():
    """新增数据"""
    # 使用类对象创建一个实例对象（一条数据）
    stu1 = Students(name='张三', math=60, chinese=60, english=60)
    stu2 = Students(name='李四', math=60, chinese=60, english=60)
    # 添加一条数据
    db.session.add(stu1)
    db.session.add(stu2)
    # 提交数据
    db.session.commit()
    print(stu1.id)
。 
```

### Read 

使用模型类提供的query属性附加调用各种过滤方法及查询方法 

一般来说，一个完整的查询遵循下面的模式： 

```
<模型类>.query.<过滤方法>.<查询方法>
过滤方法为可选，查询方法为必选
```

![image-20211023163512195](.\assets\image-20211023163512195.png)

从某个模型类出发，通过在query属性对应的Query对象上附加的过滤 方法和查询函数对模型类对应的表中的记录进行各种筛选和调整，最终返 回包含对应数据库记录数据的模型类实例，对返回的实例调用属性即可获 取对应的字段数据。 

#### 常用的查询方法 

| 查询方法（触发器）     | 说	明                                                     |
| ---------------------- | ------------------------------------------------------------ |
| all()                  | 返回包含所有查询记录的列表                                   |
| first()                | 返回查询的第一条记录，如果未找到，则返回None                 |
| one()                  | 返回第一条记录，旦仅允许有一条记录。如果记录数量大于1或小于1，则拋出错误 |
| get(ident)             | 传入主键值作为参数，返回指定主键值的记录，如果未找到，则返回None |
| count()                | 返回查询结果的数量                                           |
| one_or_none()          | 类似one()，如果结果数量不为1，返回None                       |
| first_or_404()         | 返回查询的第一条记录，如果未找到，则返回404错误响应          |
| get_or_404(ident)      | 传人主键值作为参数，返回指定主键值的记录，如果未找到，则返回404错误响应 |
| paginate()             | 返回一个Pagination对象，可以对记录进行分页处理               |
| with_parent( instance) | 传人模型类实例作为参数，返回和这个实例相关联的对象，后面会详细介绍 |

first_or_404()、get_or_404() 以及 paginate() 方法是 Flask-SQLAlchemy 附加的查询方法。 

```python
# 查询所有的记录
<模型类>.query.all()

# 返回第一条记录
<模型类>.query.first()

# 返回指定主键值（id字段）的记录
<模型类>.query.get(2)

# 返回记录的数量
<模型类>.query.count()
```

#### 常用的过滤方法

SQLAlchemy还提供了许多过滤方法，使用这些过滤方法可以获取更 精确的查询，比如获取指定字段值的记录。 

| 查询过滤器     | 说	明                                                     |
| -------------- | ------------------------------------------------------------ |
| filter_by()    | 使用指定规则过滤记录（以关键字表达式的形式），返回**新产生的查洵对象** |
| filter()       | 使用指定的规则过滤记录，返回新产生的查询对象                 |
| order_by()     | 根据指定条件对记录进行排序，返回新产生的查询对象             |
| limit(limit)   | 使用指定的值限制原查询返回的记录数量，返回新产生的查询对象   |
| group_by()     | 根据指定条件对记录进行分组，返回新产生的查询对象             |
| offset(offset) | 使用指定的值偏移原查询的结果，返回新产生的查询对象           |

```python
# 指定name字段为Hsiangya的记录
<模型类>.query.filter(Students.name == 'Hsiangya').first()

# 指定name字段不为Hsiangya的记录
<模型类>.query.filter(Students.name != 'Hsiangya').first()

# LIKE过滤，指定name字段中含有“正”的记录
filter(Students.name.like('%正%'))

# IN过滤，指定name字段在列表中的记录
filter(Students.name.in_(['张三', '李四'])

# NOT IN 过滤，指定name字段不在列表中的记录
filter(~Students.name.in_(['张三', '李四'])

# AND过滤，指定对应字段名
from sqlalchemy import and_
# 方式1，使用and_()
filter(and_(Students.name == 'Hsiangya', Students.chinese == 100)) 
# 方式2 filter()中加入多个表达式，使用逗号分隔
filter(Students.name == 'Hsiangya', Students.chinese == 100) 
# 方式3 叠加调用多个 filte()/ filte_by() 方法
filter(Students.name == 'Hsiangya').filter(Students.chinese == 100)

# OR过滤，
from sqlalchemy import or_
filter(or_(Students.name == 'Hsiangya', Students.chinese == 60)

# filter_by 过滤
filter_by(name='Hsiangya').all()

# 查询指定字典
query.with_entities(Students.name, Students.math).filter_by(name='Hsiangya').all()
```

### Update 

更新一条记录非常简单，直接赋值给模型类的字段属性就可以改变字 段值，然后调用commit()方法提交会话即可 

```python
"""更新数据"""
stu4 = Students.query.get(4)
stu4.chinese = 99
db.session.commit()
```

插入新的记录或要将现有的记录添加到会话中时才需要使用 add() 方法，单纯要更新现有的记录时只需要直接为属性赋新值，然后提交会话。

### Delete 

删除记录和添加记录很相似，不过要把add()方法换成delete()方法，最后都需要调用commit()方法提交修改。下面的示例删除了id（主 键）为2的记录： 

```python
"""删除数据"""
stu4 = Students.query.get(4)
db.session.delete(stu4)
db.session.commit()

```

## 在视图函数里操作数据库

### Read 

在视图函数中查询数据库记录并传入模板 

```python
@app.route('/')
def index():
    students = Students.query.all()
    return render_template('index.html', stus=students)

```

在模板中，我们迭代这个 stus 列表，调用 stus对象的body属性（note.body）获取body字段的值。通过length过滤器获取笔记的数量。

```
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>学生信息管理系统</title>
    <style>
        * {margin: 0;padding: 0;}
        table {margin: 100px auto;}
        .footer {width: 500px;margin: 0 auto;}
        .footer a {display: inline-block;width: 20%;text-decoration: none;color: #cccccc;}
        .footer a:hover {color: orange;}
    </style>
</head>
<body>
<div class="table">
    <table border="1" align="center" width="500" cellpadding="2" cellspacing="0">
        <tr>
            <td>姓名</td>
            <td>数学</td>
            <td>语文</td>
            <td>英语</td>
            <td>总分</td>
        </tr>
        {% for stu in stus %}
            <tr>
                <td>{{ stu.name }}</td>
                <td>{{ stu.math }}</td>
                <td>{{ stu.chinese }}</td>
                <td>{{ stu.english }}</td>
                <td>{{ stu.total }}</td>
            </tr>
        {% endfor %}
    </table>
    <div class="footer">
        <span><a href="{{ url_for('insert') }}">新增</a></span>
        <span><a href="{{ url_for('delete') }}">删除</a></span>
        <span><a href="#">修改</a></span> <span>
        <a href="{{ url_for('index') }}">查询</a></span>
    </div>
</div>
</body>
</html>
```

### 分页查询

参数定义：

- `page` 查询的页数
- `per_page` 每页的条数
- `max_per_page` 每页最大条数，有值时，`per_page` 受它影响
- `error_out`  错误是否报错
  - 当值为 True 时，下列情况会报错
  - 当 page 为 1 时，找不到任何数据
  - page 小于 1，或者 per_page 为负数
  - page 或 per_page 不是整数


该方法返回一个分页对象 `Pagination` 

| 方法                                                         | 简介                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| has_next                                                     | 如果下一页存在，返回 True                                    |
| has_prev                                                     | 如果上一页存在，返回 True                                    |
| items                                                        | 当前页的数据列表                                             |
| next_num                                                     | 下一页的页码                                                 |
| page                                                         | 当前页码                                                     |
| pages                                                        | 总页数                                                       |
| per_page                                                     | 每页的条数                                                   |
| prev_num                                                     | 上一页的页码                                                 |
| query                                                        | 用于创建此分页对象的无限查询对象。                           |
| total                                                        | 总条数                                                       |
| iter_pages(left_edge=2, left_current=2, right_current=5, right_edge=2) | 迭代分页中的页码，四个参数，分别控制了省略号左右两侧各显示多少页码，在模板中可以这样渲染 |

```python
# 分页查询 第一页，每页20条数据。 默认第一页。
paginate_obj = User.query.paginate(page=1, per_page=20, error_out=False)  
# 参数：error_out 设为True表示页数不是int或超过总页数时,会报错,并返回404状态码。 默认True


# 获取总页数
total_page = paginate_obj.pages
```

### Create 

创建Srudents类模型

```python
from flask_wtf import FlaskForm
from wtforms import StringField,  SubmitField, IntegerField
from wtforms.validators import DataRequired

class Students(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(length=20))
    math = db.Column(db.Integer, default=0)
    chinese = db.Column(db.Integer)
    english = db.Column(db.Integer)

    @property
    def total(self):
        return self.math + self.chinese + self.english

    def __repr__(self):
        return '<Students %s>' % self.name


```

创建一个 insert 视图，这个视图负责渲染创建笔记的模板，并处理表单的提交

```python
@app.route('/insert', methods=['GET', 'POST'])
def insert():
    form = InsertFrom(request.form)
    # 前端返回True时处理代码，否则返回insert.html
    if form.validate_on_submit():
        # 构造stu对象
        stu = Students(
            # 从表单中获取字段数据
            name=request.form.get('name'),
            math=request.form.get('math'),
            chinese=request.form.get('chinese'),
            english=request.form.get('english'),
        )
        db.session.add(stu)
        # 提交会话
        db.session.commit()
        # 完成后重定向到index
        return redirect(url_for('index'))
    return render_template('insert.html', form=form)
```

这个过程接收用户通过表单提交的数据并保存到数据库中，最后我们使用flash()函数发送提示消息并重定向到index视图。 

> insert.html

```jinja2
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>新增数据</title>
    <style>
        .form {
            width: 500px;
            margin: 100px auto;
        }
    </style>
</head>
<body>
<div class="form">
    <form method="POST" action="{{ url_for('insert') }}">
        {{ form.csrf_token() }}
        <table>
            <tr>
                <td>{{ form.name.label }}</td>
                <td>{{ form.name }}</td>
            </tr>
            <tr>
                <td>{{ form.math.label }}</td>
                <td>{{ form.math }}</td>
            </tr>
            <tr>
                <td>{{ form.chinese.label }}</td>
                <td>{{ form.chinese }}</td>
            </tr>
            <tr>
                <td>{{ form.english.label }}</td>
                <td>{{ form.english }}</td>
            </tr>
        </table>
        {{ form.submit }}
    </form>
</div>
</body>
</html>
```

### Delete 

在程序中，删除的实现也非常简单，不过这里经常会有一个误区。大 多数人通常会考虑在笔记内容下添加一个删除链接： 

```
<a href="{{ url_for('delete', id=stu.id) }}">删除</a>

```

这个链接指向用来删除笔记的delete_note视图： 

```python
@app.route('/delete')
def delete():
    # 获取请求中的查询参数
    id_ = request.args.get('id')
    if id_:
        stu = Students.query.get(id_)
        db.session.delete(stu)
        db.session.commit()
        return redirect(url_for('delete'))

    students = Students.query.all()
    return render_template('delete.html', stus= students)
```

### Update 

略

## 更新数据库表 重新生成表 

数据库表不会随着模型的修改而自动更新，重新调用create_all（）方法并不会起到更新表或重新创建表的作用。 如果你并不在意表中的数据，最简单的方法是使用drop_all（）方法删除表 以及其中的数据，然后再使用create_all（）方法重新创建： 

```python
db.drop_all()
db.create_all()
```

**这会清除数据库里的原有数据，请勿在生产环境下使用**

为了方便开发，我们修改initdb命令函数的内容，为其增加一个--drop 选项来⽀持删除表和数据库后进行重建

> ⽀持删除表后重建 

```python
@app.cli.command() 
@click.option('--drop', is_flag=True, help='Create after drop.') 
def initdb(drop): 
    """Initialize the database.""" 
    if drop: 
        click.confirm('This operation will delete the database, do you want to continue?', abort=True) 
        db.drop_all() 
        click.echo('Drop tables.')
    db.create_all()
    click.echo('Initialized database.')

```

在这个命令函数前，我们使用click提供的option装饰器为命令添加了一 个--drop选项，将is_flag参数设为True可以将这个选项声明为布尔值标志 （boolean flag）。--drop选项的值作为drop参数传入命令函数，如果提供了 这个选项，那么drop的值将是True，否则为False。因为添加--drop选项会直 接清空数据库内容，如果需要，也可以通过click.confirm（）函数添加一个 确认提示，这样只有输入y或yes才会继续执行操作。 

现在，执行下面的命令会重建数据库和表： 

```
$ flask initdb --drop
```

当使用SQLite时，直接删除data.db文件和调用 drop_all() 方法效果相同，而且更直接，不容易出错。 

## 数据迁移

```bash
# 安装
pip install flask-migrate

# 初始化
flask db init

# 提交版本说明
flask db migrate -m "init db"

# 更新数据库
flask db upgrade

# 数据库版本降级
flask db history
flask db downgrade 版本号



# 安装
pip install alembic

# 建立一个 Alembic 脚本环境
alembic init alembic
https://zhuanlan.zhihu.com/p/90106173
```

## 定义关系

使用relationship()方法来建立表模型与表模型之间的关系

```python
# 1. 多侧创建外键，并指向一侧主键
class Article(db.Model):
    ...
	author_id = db.Column(db.Integer, db.ForeignKey('author.id'))

# 2. 一侧建立关系，articles 为集合属性
class Author(db.Model): 
    ...
	articles = db.relationship('Article')
```

传入ForeignKey类的参数author.id，其中author指的是Author模型对应的表名称，而id指的是字段名，即“表名.字段名”。模型类对应的表名由Flask-SQLlchemy生成，默认为类名称的⼩写形式，多个单词通过下划线分隔，你也可以显式地通过`__tablename__` 属性自己指定。 

### 属性

| 属性           | 介绍                                                        |
| -------------- | ----------------------------------------------------------- |
| back_populates | 显示定义反向引用，两侧都需要定义                            |
| backref        | 隐式定义反向引用，外键另一次定义                            |
| lazy           | 指定如何加载相关记录                                        |
| uselist        | 指定是否使用列表的形式加载记录，设为False则使用标量(scalar) |
| cascade        | 设置级联操作                                                |
| order_by       | 指定加载相关记录时的排序方式                                |
| secondary      | 多对多关系中 指定二者关系表                                 |
| primaryjoin    | 指定多对多关系中的一级联结条件                              |
| secondaryjoin  | 指定多对多关系中的二级联结条件                              |

### back_populates

显示定义，需要一与多两侧都添加关系

```python
class Author(db.Model):
	...
    articles = db.relationship('Article', back_populates='author')


class Article(db.Model):
    ...
    # 1. 定义外键
    author_id = db.Column(db.Integer, db.ForeignKey('author.id'))
    author = db.relationship('Author', back_populates='articles')

```

###  backref

backref只需要在一侧添加关系

```python
class Author(db.Model):
	...
    # 4. 直接添加反向引用
    articles = db.relationship('Article', backref='author')


class Article(db.Model):
    # 1. 定义外键
    author_id = db.Column(db.Integer, db.ForeignKey('author.id'))
```

### lazy

当关系属性被调用时，关系函数会加载相应的记录，lazy参数可选值

| 关系加载方式 | 说	明                                                     |
| ------------ | ------------------------------------------------------------ |
| select       | 在必要时一次性加载记录，返回包含记录的列表（默认值），等同于lazy=Tme |
| joined       | 和父査询一样加载记录，但使用联结，等同于lazy=False           |
| immediate    | 一旦父查询加载就加载                                         |
| subquery     | 类似于joined,不过将使用子查询                                |
| dynamic      | 不直接加载记录，而是返回一个包含相关记录的qiiery对象，以便再继续附加查询函 数对结果进行过滤 |

dynamic选项仅用于集合关系属性，不可用于多对一、一对一或是在关 系函数中将uselist参数设为False的情况。 

许多教程和示例使用dynamic来动态加载所有集合关系属性对应的记 录，这是应该避免的行为。使用dynamic加载方式意味着每次操作关系都会 执行一次SQL查询，这会造成潜在的性能问题。大多数情况下我们只需要 使用默认值（select），只有在调用关系属性会返回大量记录，并且总是需要对关系属性返回的结果附加额外的查询时才需要使用动态加载 （lazy='dynamic'）。 

### uselist

一对一就是多对一和一对多的一个特例,只需在relationship加上一个参数uselist=False替换多的一端就是一对一

一对多转一对一

```python
class Parent(Base):
    __tablename__ = 'parent'
    id = Column(Integer, primary_key=True)
    child = relationship("Child", uselist=False, backref="parent")

class Child(Base):
    __tablename__ = 'child'
    id = Column(Integer, primary_key=True)
    parent_id = Column(Integer, ForeignKey('parent.id'))
```

多对一转一对一

```python
class Parent(Base):
    __tablename__ = 'parent'
    id = Column(Integer, primary_key=True)
    child_id = Column(Integer, ForeignKey('child.id'))
    child = relationship("Child", backref=backref("parent", uselist=False))

class Child(Base):
    __tablename__ = 'child'
    id = Column(Integer, primary_key=True)
```

### secondary

在 SQLAlchemy中，要想表示多对多关系，除了关系两侧的模型外，我们还 需要创建一个关联表（association table）。关联表不存储数据，只用来存 储关系两侧模型的外键对应关系，如代码所示。 

```python
association_table = db.Table(
    'association',
    db.Column('student_id', db.Integer, db.ForeignKey('student.id')),
    db.Column('teacher_id', db.Integer, db.ForeignKey('teacher.id'))
)

class Student(db.Model):
	...
    teachers = db.relationship('Teacher',
                               secondary=association_table,
                               back_populates='students')


class Teacher(db.Model):
	...
    students = db.relationship('Student',
                               secondary=association_table,
                               back_populates='teachers') 
```

需要为某一个学生添加老师时，对关系属性使用 append()方法即可。想要解除关系，那么可以使用remove()方 法。 

关联表由SQLAlchemy接管，它会帮我们管理这个表：我们只需要像 往常一样通过操作关系属性来建立或解除关系，SQLAlchemy会自动在关 联表中创建或删除对应的关联表记录，而不用手动操作关联表。

同样的，在多对多关系中我们也只需要在关系的一侧操作关系。当为 学生A的teachers添加了老师B后，调用老师B的students属性时返回的学生 记录也会包含学生A，反之亦同。

## 建立关系

建立关系有两种方式，第一种方式是为外键字段赋值 

```python
# 建立关系 方式1 为外键赋值
article1.author_id = 1 
db.session.commit()

# 建立关系 方式2 调用append()方法
author1.articles.append(article1) author1.articles.append(article2)
db.session.commit()

# 解除关系 方法1 使用remove()方法
articles.remove(spam) 
db.session.commit() 
author1.articles 

# 接触关系 方式2 使用pop()方法
articles.pop(spam) 
db.session.commit() 
author1.articles 
```

## 多对多案例

在项目开发过程中，会遇到很多数据之间多对多关系的情况，比如：

- 学生网上选课(学生和课程)
- 老师与其授课的班级(老师和班级)
- 用户与其收藏的新闻(用户和新闻)
- 等等...

所以在开发过程中需要使用 ORM 模型将表与表的多对多关联关系使用代码描述出来。多对多关系描述有一个唯一的点就是：**需要添加一张单独的表去记录两张表之间的对应关系**

### 需求分析

- 学生可以网上选课，学生有多个，课程也有多个
- 学生有：张三、李四、王五
- 课程有：物理、化学、生物
- 选修关系有：
  - 张三选修了化学和生物
  - 李四选修了化学
  - 王五选修了物理、化学和生物
- 需求：
  1. 查询某个学生选修了哪些课程
  2. 查询某个课程都有哪些学生选择

### 思路分析

- 可以通过分析得出
  - 用一张表来保存所有的学生数据
  - 用一张表来保存所有的课程数据
- 具体表及测试数据可以如下：

### 学生表(Student)

| 主键(id) | 学生名(name) |
| :------- | :----------- |
| 1        | 张三         |
| 2        | 李四         |
| 3        | 王五         |

### 选修课表(Course)

| 主键(id) | 课程名(name) |
| :------- | :----------- |
| 1        | 物理         |
| 2        | 化学         |
| 3        | 生物         |

### 数据关联关系表(Student_Course)

| 主键(student.id) | 主键(course.id) |
| :--------------- | :-------------- |
| 1                | 2               |
| 1                | 3               |
| 2                | 2               |
| 3                | 1               |
| 3                | 2               |
| 3                | 3               |

### 结果

- 查询某个学生选修了哪些课程，例如：查询王五选修了哪些课程
  - 取出王五的 id 去 Student_Course 表中查询 student.id 值为 3 的所有数据
  - 查询出来有3条数据，然后将这3条数据里面的 course.id 取值并查询 Course 表即可获得结果
- 查询某个课程都有哪些学生选择，例如：查询生物课程都有哪些学生选修
  - 取出生物课程的 id 去 Student_Course 表中查询 course.id 值为 3 的所有数据
  - 查询出来有2条数据，然后将这2条数据里面的 student.id 取值并查询 Student 表即可获得结果

### 代码演练

- 定义模型及表

```python
tb_student_course = db.Table('tb_student_course',
                             db.Column('student_id', db.Integer, db.ForeignKey('students.id')),
                             db.Column('course_id', db.Integer, db.ForeignKey('courses.id'))
                             )


class Student(db.Model):
    __tablename__ = "students"
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(64), unique=True)

    courses = db.relationship('Course', secondary=tb_student_course,
                              backref='student',
                              lazy='dynamic')


class Course(db.Model):
    __tablename__ = "courses"
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(64), unique=True)
```

- 添加测试数据

```python
if __name__ == '__main__':
    db.drop_all()
    db.create_all()

    # 添加测试数据

    stu1 = Student(name='张三')
    stu2 = Student(name='李四')
    stu3 = Student(name='王五')

    cou1 = Course(name='物理')
    cou2 = Course(name='化学')
    cou3 = Course(name='生物')

    stu1.courses = [cou2, cou3]
    stu2.courses = [cou2]
    stu3.courses = [cou1, cou2, cou3]

    db.session.add_all([stu1, stu2, stu2])
    db.session.add_all([cou1, cou2, cou3])

    db.session.commit()

    app.run(debug=True)
```

# Restful 规范

## 目的

REST即表述性状态传递（英文：Representational State Transfer，简称REST），RESTful API 统一约束客户端和服务器之间的接口。简化和分离系统架构，使每个模块独立！

+ 请求中使用URI定位资源
+ 用HTTP Verbs[动词]（GET、POST、PUT、DELETE）描述操作（具体表现形式）
+ 数据传递（默认）采用：`Content-Type: application/json; charset=utf-8` 

## Rest

### 无状态（Stateless）

REST要求状态要么被放入资源状态中，要么保存在客户端上。服务器端不能保持除了单次请求之外的，任何与其通信的客户端的通信状态。从客户端的每个请求要包含服务器所需要的所有信息。这样做的最直接的理由就是可伸缩性—— 如果服务器需要保持客户端状态，那么大量的客户端交互会严重影响服务器的内存可用空间（footprint）。

### 缓存（Cachable）

服务器返回信息必须被标记是否可以缓存，如果缓存，客户端可能会重用之前的信息发送请求。

### 客户-服务器（Client-Server）

客户端无需关注数据存储，服务器端无需关注用户界面，提高了前后端可移植性。

### 分层系统（Layered System）

客户端不关心直接连接到最终服务器还是连接到中间服务器。中间服务器可以通过启用负载平衡和提供共享缓存来提高系统可扩展性。分层系统也可以执行安全策略。

### 支持按需代码（Code on Demand，可选）

服务器可以通过传输逻辑来临时扩展或定制客户端的功能。

## Request

| 请求方法         | 说明                                 |
| ---------------- | ------------------------------------ |
| GET（SELECT）    | 获取资源                             |
| POST（CREATE）   | 创建资源                             |
| PUT（UPDATE）    | 更新资源                             |
| DELETE（DELETE） | 删除资源                             |
| HEAD             | 获取资源的元数据                     |
| OPTIONS          | 获取资源哪些属性是客户端可变的信息。 |

## 过滤

如果记录数量很多，服务器不可能都将它们返回给用户。API应该提供参数，过滤返回结果。

下面是一些常见的参数。

```
?limit=10					：指定返回记录的数量
?offset=10					：指定返回记录的开始位置。
?page=2&per_page=100		：指定第几页，以及每页的记录数。
?sortby=name&order=asc		：指定返回结果按照哪个属性排序，以及排序顺序。
?category_type_id=1			：指定筛选条件
```

参数的设计允许存在冗余，即允许API路径和URL参数偶尔有重复。比如，GET /category/ID/todo 与 GET /todo ?category_id=ID 的含义是相同的。

## Respone

### 状态码

| 状态码                    | 说明                                                       |
| ------------------------- | ---------------------------------------------------------- |
| 200 OK                    | 服务器成功返回请求的数据                                   |
| 201 CREATED               | 新建或修改数据成功                                         |
| 202 Accepted              | 表示一个请求已经进入后台排队（异步任务）                   |
| 204 NO CONTENT            | 删除数据成功                                               |
| 400 INVALID REQUEST       | 请求有错误，服务器没有进行新建或修改数据的操作（幂等操作） |
| 401 Unauthorized          | 没有权限（令牌、用户名、密码错误）                         |
| 403 Forbidden             | 得到授权（与401错误相对），但是访问是被禁止的              |
| 404 NOT FOUND             | 请求记录不存在，服务器没有进行操作（幂等操作）             |
| 406 Not Acceptable        | 请求的格式不符合（比如用户请求JSON格式，但是只有XML格式）  |
| 500 INTERNAL SERVER ERROR | 服务器发生错误，无法判断发出的请求是否成功                 |

### 返回结果

针对不同操作，服务器向用户返回的结果应该符合以下规范。

```python
GET 	/collection				：返回资源对象的列表（数组）
GET 	/collection/resource	：返回单个资源对象
POST 	/collection				：返回新生成的资源对象
PUT 	/collection/resource	：返回完整的资源对象
PATCH 	/collection/resource	：返回完整的资源对象
DELETE 	/collection/resource	：返回一个空文档
```

### 格式

前后端交互字段全部使用小驼峰方式

```json
{
  "code": "200", 					// HTTP响应码(好多javascript框架并不会获取http状态码，所以包装到body中便于使用)
  "status": "success/fail/error", 	// 见下述表格
  "content/data/result": []/{}, 	// 多条记录使用JSON数组，单条记录使用JSON对象
  "message": []     				// 状态为error或fail时，对应的错误信息
}
```

**status说明** 

| 状态    | 说明                        |
| ------- | --------------------------- |
| fail    | 返回码为 500-599            |
| error   | 返回码为 400-499            |
| success | 其他状态码（1xx、2xx、3xx） |

# login_manager

## 插件结构

```
──flask_login	
    |-- __about__.py
    |-- __init__.py
    |-- _compat.py
    |-- config.py
    |-- login_manager.py
    |-- mixins.py
    |-- signals.py
    `-- utils.py
```

## Login_Manager

LoginManager 是一个类，有多个方法和属性；此类初始化的对象用于保存用于登录的设置。实例：class：LoginManager 不绑定到特定应用程序，因此可以在代码的主体中创建一个，然后将其绑定到您的应用程序app中 工厂函数。

```python
# 初始化一个LoginManager类对象
login_manager = LoginManager()

# init__app() – 配置该对象
login_manager.init_app(app)

# login-view ：验证失败跳转的界面
login_manager.login_view = "index"  # 跳转到/首页目录

# 用户重定向到登录页面时闪出的消息
login_manager.login_message ='Please login first!'

# 用户需要重新进行身份验证时要重定向到的视图的名称。
login_manager.refresh_view = "/"

# user_loader() 回调函数
@login_manager.user_loader
def load_user(userid):
    return User.get(userid)

# 保存当前成功登录的用户
login_user(user)

# 退出登录
@app.route('/logout')
@login_required
def logout():
    logout_user()
    return 'Logged out successfully!'
```

## 方法源码

### user_loader

自定义回调函数，user_loader源代码:

```python
    def user_loader(self, callback):
        '''
        This sets the callback for reloading a user from the session. The
        function you set should take a user ID (a ``unicode``) and return a
        user object, or ``None`` if the user does not exist.

        :param callback: The callback for retrieving a user object.
        :type callback: callable
        '''
        self.user_callback = callback
        return callback
```

这将设置从会话重新加载用户的回调。 您设置的函数应该使用用户ID（“unicode”）并返回用户对象，如果用户不存在则返回“None”。

在执行该段代码之后，注册了load_user()这个自定义的callback。

### reload_user()

1、首先获取user id，如果获取不到有效的id，就将user设为anonymous user
2、获取到id后，再通过@login_manager.user_loader装饰的函数获取到user对象，如果没有获取到有效的user对象，就认为是anonymous user
3、最后将user保存于request context中（无论是正常的用户还是anonymous用户）

```python
    def reload_user(self, user=None):
        '''
        This set the ctx.user with the user object loaded by your customized
        user_loader callback function, which should retrieved the user object
        with the user_id got from session.

        Syntax example:
        from flask_login import LoginManager
        @login_manager.user_loader
        def any_valid_func_name(user_id):
            # get your user object using the given user_id,
            # if you use SQLAlchemy, for example:
            user_obj = User.query.get(int(user_id))
            return user_obj

        Reason to let YOU define this self.user_callback:
            Because we won't know how/where you will load you user object.
        '''
        ctx = _request_ctx_stack.top

        if user is None:
            user_id = session.get('user_id')
            if user_id is None:
                ctx.user = self.anonymous_user()
            else:
                if self.user_callback is None:
                    raise Exception(
                        "No user_loader has been installed for this "
                        "LoginManager. Refer to"
                        "https://flask-login.readthedocs.io/"
                        "en/latest/#how-it-works for more info.")
                user = self.user_callback(user_id)
                if user is None:
                    ctx.user = self.anonymous_user()
                else:
                    ctx.user = user
        else:
            ctx.user = user
```

## utils

### Login_required

current_user.is_authenticated，判断当前用户是否认证，如果没有认证的话就执行unauthorized()，
unauthorized()会重定向到login_view参数设置的路由函数中去，所以在实例化LoginManager后要设置login_view属性，当用户没有登陆时，会自动重定向到登陆界面，没有设置会返回401错误， 用户登陆后，这个装饰器就直接返回func(*args, **kwargs)，相当于没有包装一样，装饰器起到包装接口的作用。源码

```python
def login_required(func):
    '''
    If you decorate a view with this, it will ensure that the current user is
    logged in and authenticated before calling the actual view. (If they are
    not, it calls the :attr:`LoginManager.unauthorized` callback.) For
    example::

        @app.route('/post')
        @login_required
        def post():
            pass

    If there are only certain times you need to require that your user is
    logged in, you can do so with::

        if not current_user.is_authenticated:
            return current_app.login_manager.unauthorized()

    ...which is essentially the code that this function adds to your views.

    It can be convenient to globally turn off authentication when unit testing.
    To enable this, if the application configuration variable `LOGIN_DISABLED`
    is set to `True`, this decorator will be ignored.

    .. Note ::

        Per `W3 guidelines for CORS preflight requests
        <http://www.w3.org/TR/cors/#cross-origin-request-with-preflight-0>`_,
        HTTP ``OPTIONS`` requests are exempt from login checks.

    :param func: The view function to decorate.
    :type func: function
    '''
    @wraps(func)
    def decorated_view(*args, **kwargs):
        if request.method in EXEMPT_METHODS:
            return func(*args, **kwargs)
        elif current_app.login_manager._login_disabled:
            return func(*args, **kwargs)
        elif not current_user.is_authenticated:
            return current_app.login_manager.unauthorized()
        return func(*args, **kwargs)
    return decorated_view


```

源码解释

1、如果使用此装饰视图，它将确保在调用实际视图之前登录并验证当前用户。如果验证不通过，那么则会调用LoginManager.unauthorized（）
2、如果request method为例外method，即在EXEMPT_METHODS中的method，可以不必鉴权
3、如果_login_disabled为True则不必鉴权

### Login_user

用户登录操作，用户在通过自定义的login视图函数，通过验证并登录成功后，需要使用Flask-Login中的login_user可以记录/保存当前成功登陆的用户

```
  login_user(user)
```

> 源码

```python
def login_user(user, remember=False, duration=None, force=False, fresh=True):
    '''
    Logs a user in. You should pass the actual user object to this. If the
    user's `is_active` property is ``False``, they will not be logged in
    unless `force` is ``True``.

    This will return ``True`` if the log in attempt succeeds, and ``False`` if
    it fails (i.e. because the user is inactive).

    :param user: The user object to log in.
    :type user: object
    :param remember: Whether to remember the user after their session expires.
        Defaults to ``False``.
    :type remember: bool
    :param duration: The amount of time before the remember cookie expires. If
        ``None`` the value set in the settings is used. Defaults to ``None``.
    :type duration: :class:`datetime.timedelta`
    :param force: If the user is inactive, setting this to ``True`` will log
        them in regardless. Defaults to ``False``.
    :type force: bool
    :param fresh: setting this to ``False`` will log in the user with a session
        marked as not "fresh". Defaults to ``True``.
    :type fresh: bool
    '''
    if not force and not user.is_active:
        return False

    user_id = getattr(user, current_app.login_manager.id_attribute)()
    session['user_id'] = user_id
    session['_fresh'] = fresh
    session['_id'] = current_app.login_manager._session_identifier_generator()

    if remember:
        session['remember'] = 'set'
        if duration is not None:
            try:
                # equal to timedelta.total_seconds() but works with Python 2.6
                session['remember_seconds'] = (duration.microseconds +
                                               (duration.seconds +
                                                duration.days * 24 * 3600) *
                                               10**6) / 10.0**6
            except AttributeError:
                raise Exception('duration must be a datetime.timedelta, '
                                'instead got: {0}'.format(duration))

    _request_ctx_stack.top.user = user
    user_logged_in.send(current_app._get_current_object(), user=_get_user())
    return True


```

### logout_user

登出功能类似，除了基本的操作外，还需要把flask-login中的登出进行操作

```python
@app.route('/logout')
@login_required
def logout():
    logout_user()
    return 'Logged out successfully!'

```

> 源码

```python
def logout_user():
    '''
    Logs a user out. (You do not need to pass the actual user.) This will
    also clean up the remember me cookie if it exists.
    '''

    user = _get_user()

    if 'user_id' in session:
        session.pop('user_id')

    if '_fresh' in session:
        session.pop('_fresh')

    cookie_name = current_app.config.get('REMEMBER_COOKIE_NAME', COOKIE_NAME)
    if cookie_name in request.cookies:
        session['remember'] = 'clear'
        if 'remember_seconds' in session:
            session.pop('remember_seconds')

    user_logged_out.send(current_app._get_current_object(), user=user)

    current_app.login_manager.reload_user()
    return True

```

## UserMixin

- is_authenticated
  当用户通过验证时，也即提供有效证明时返回 True 。（只有通过验证的用户会满足 login_required 的条件。）

- is_active
  如果这是一个活动用户且通过验证，账户也已激活，未被停用，也不符合任何你 的应用拒绝一个账号的条件，返回 True 。不活动的账号可能不会登入（当然， 是在没被强制的情况下）。

- is_anonymous
  如果是一个匿名用户，返回 True 。（真实用户应返回 False 。）

- get_id()
  返回一个能唯一识别用户的，并能用于从 user_loader 回调中加载用户的 unicode 。注意着 必须 是一个 unicode —— 如果 ID 原本是 一个 int 或其它类型，你需要把它转换为 unicode 。

要简便地实现用户类，你可以从 UserMixin 继承，它提供了对所有这些方法的默认 实现。（虽然这不是必须的。）

## 使用基础流程

Flask-Login 通过 user session，提供登录的常见任务，比如登入 (logging in)、登出 (logging out) 和当前用户 (current user)

1. login_user()：实现用户的登入，一般在登入的视图函数中调用
2. logout_user()：实现登出功能
3. current_user 属性：获取当前用户

如果需要页面是授权用户才可见，在相应视图函数前加上 @login_required 装饰器进行声明即可，@login_required 装饰器对于未登录用户访问，默认处理是重定向到LoginManager.login_view 所指定的视图

### 定义 User

登录基于用户，需要定义 User 类，Flask-Login 规定 User 类必须实现三个属性和一个方法：

is_authenticated 属性

is_active 属性

is_anonymous 属性

get_id() 方法

最简单的方法是从 UserMixin 类继承，该类提供了默认的实现。

```python
from flask_login import UserMixin

class User(UserMixin，db.Model):
    pass
```

### 应用程序的配置

创建 LoginManager 实例，然后与 app 绑定。

```python
from flask_login import LoginManager

login_manager = LoginManager()

login_manager.login_view = 'login'
login_manager.login_message_category = 'info'
login_manager.login_message = 'Access denied.'

login_manager.init_app(app)


```

### user_loader 回调函数

user session 记录的是用户 ID (user_id)，回调函数的作用就是通过 user_id 返回对应的 User 对象。user_loader 回调函数在 user_id 非法的时候不应该抛出异常，而要返回 None。没有这个回调函数的话，Flask-Login 将无法工作

```python
@login_manager.user_loader
def load_user(user_id):
    if query_user(user_id) is not None:
        curr_user = User()
        curr_user.id = user_id

        return curr_user

```

### 登入功能实现

```python
@app.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        user_id = request.form.get('userid')
        user = query_user(user_id)
        if user is not None and request.form['password'] == user['password']:

        curr_user = User()
        curr_user.id = user_id

        # 通过Flask-Login的login_user方法登录用户
        login_user(curr_user)

        return redirect(url_for('index'))

    flash('Wrong username or password!')

# GET 请求
return render_template('login.html')


```

关键就是 login_user(curr_user) 这句代码，之前要构建 User 对象，并指定 id。

### 登出功能实现

```python
@app.route('/logout')
@login_required
def logout():
    logout_user()
    return 'Logged out successfully!'

```

# 源码

## local

![threading_local](./assets/threading_local.png)

## Flask中的LocalStack

![LocalStack](./assets/LocalStack.png)

## 请求准备阶段

![Flask启动准备阶段](./assets/Flask启动准备阶段.png)
