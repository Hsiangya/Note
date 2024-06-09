#  介绍

## APIview与View

![View与APIView](./assets/View与APIView.png)

## Requset对象的封装

![DRFReuqest对象](./assets/DRFReuqest对象.png)

## Request封装中的面向对象

![Request对象封装](./assets/Request对象封装.png)

## 匿名用户

![匿名用户源码](./assets/匿名用户源码.png)

## dispatch源码

![dispatch源码](./assets/dispatch源码.png)

# 认证组件

## 使用方式

![认证快速应用](./assets/认证快速应用.png)

## 全局配置

![认证组件全局与局部配置](./assets/认证组件全局与局部配置.png)

## 源码流程

![认证组件源码流程](./assets/认证组件源码流程.png)

### 多认证组件

实现多个认证组件，由于匿名用户的配置，认证没通过不报错，因此可以自定义一个认证组件，放在末尾，主动抛出异常

```python
from rest_framework.authentication import BaseAuthentication
from rest_framework.exceptions import AuthenticationFailed

from api import models


class QueryParamsAuthentication(BaseAuthentication):
    """参数中携带token"""
    def authenticate(self, request):
        # 获取token
        token = request.query_params.get("token")
        if not token:
            return
		
        # 获取用户对象
        user_object = models.UserInfo.objects.filter(token=token).first()
        if user_object:
            # request.user = 用户对象; request.auth = token
            return user_object, token  

    def authenticate_header(self, request):
        # return 'Basic realm="API"'
        # 若不返回，会抛出403异常
        return "API"


class HeaderAuthentication(BaseAuthentication):
    """请求头中携带token"""
    def authenticate(self, request):
        token = request.META.get("HTTP_AUTHORIZATION")
        if not token:
            return
        user_object = models.UserInfo.objects.filter(token=token).first()
        if user_object:
            return user_object, token  # request.user = 用户对象; request.auth = token

        return

    def authenticate_header(self, request):
        # return 'Basic realm="API"'
        return "API"


class NoAuthentication(BaseAuthentication):
    """没有权限"""
    def authenticate(self, request):
        raise AuthenticationFailed({"status": False, 'msg': "认证失败"})
	"""不定义该方法认证失败时会返回状态码403"""
    def authenticate_header(self, request):
        return "API"

```

## 状态码一致

![状态码一致](./assets/状态码一致.png)

# 权限组件

## 局部和全局使用

![全局与局部使用](./assets/全局与局部使用.png)

## 源码流程

![权限源码流程](./assets/权限源码流程.png)

## 权限或逻辑扩展

![权限逻辑改写和扩展](./assets/权限逻辑改写和扩展.png)

# 限流组件

## 限流逻辑

![限流逻辑](./assets/限流逻辑.png)

## 使用

![限流类的使用](./assets/限流类的使用.png)

## 源码

![限流组件源码](./assets/限流组件源码.png)

# 版本

## 源码

![版本组件源码](./assets/版本组件源码.png)

# 解析器

## 源码

![解析器类源码](./assets/解析器类源码.png)

## 文件解析器

![文件解析器](./assets/文件解析器.png)

# 元类

## 类的创建

![类的创建方式](./assets/类的创建方式.png)

## 元类创建

![元类创建](./assets/元类创建.png)

## 元类的继承

![元类的继承](./assets/元类的继承.png)

## call

![call](./assets/call.png)

# 序列化器

目的：将数据库获取到的QuerySet

## 继承关系

![image-20230327220119668](./assets/image-20230327220119668.png)

## 简单使用

![序列化器简单使用](./assets/序列化器简单使用.png)

## ModelSerializer

![modelSerializer的使用](./assets/modelSerializer的使用.png)

## 自定义返回字段

![自定义返回字段](./assets/自定义返回字段.png)![image-20230329102131529](./assets/image-20230329102131529.png)

## 序列化嵌套和继承

![序列化嵌套和继承](./assets/序列化嵌套和继承.png)

## 源码--创建

字段创建时，会维护一个`_creation_counter_`用于决定后续处理时字段的顺序

![字段对象的创建排序](./assets/字段对象的创建排序.png)

## 源码--字段方法

![to_representation](./assets/to_representation.png)

## 源码--类的创建

![序列化类的创建](./assets/序列化类的创建.png)

## 源码-- 实例化对象

![序列化类实例化对象](./assets/序列化类实例化对象.png)

## 序列化过程

![序列化过程](./assets/序列化过程.png)

## List序列化过程

![List序列化的过程](./assets/List序列化的过程.png)

# 数据校验

## 基本校验

![校验器基本校验](./assets/校验器基本校验.png)

## 正则校验

![正则校验](./assets/正则校验.png)

## 钩子校验

![钩子校验](./assets/钩子校验.png)

## model校验

![model校验](./assets/model校验.png)

## FK与M2M

![FK与M2M](./assets/FK与M2M.png)

# 分页

| 方法                                                       | 说明 |
| ---------------------------------------------------------- | ---- |
| `_divide_with_ceil(a, b)`                                  |      |
| `_get_displayed_page_numbers(current, final)`              |      |
| `_get_page_links(page_numbers, current, url_func)`         |      |
| `_positive_int(integer_string, strict=False, cutoff=None)` |      |
| `_reverse_ordering(ordering_tuple)`                        |      |
| BasePagination类                                           |      |
| CursorPagination(BasePagination)类                         |      |
| LimitOffsetPagination(BasePagination)类                    |      |
| PageNumberPagination(BasePagination)类                     |      |

## BasePagination

```python
display_page_controls = False
def get_paginated_response_schema(self, schema):
    return schema

def get_results(self, data):
    return data["results"]

def get_schema_fields(self, view):
    assert (
        coreapi is not None
    ), "coreapi must be installed to use `get_schema_fields()`"
    return []

def get_schema_operation_parameters(self, view):
    return []
```

## LimitOffsetPagination

### 使用方式

```Python
from rest_framework.pagination import LimitOffsetPagination
def get(self, request, *args, **kwargs):
    # 1.读取数据库中的数据，返回查询集
    queryset = models.xxxx.objects.all().order_by("-id")
    # ?max_id=1
    # ?min_id=13
    max_id = request.query_params.get("max_id")
    if max_id:
        queryset = queryset.filter(id__lt=max_id)
        
    # 2.分页处理得到分页后的->queryset
    pager = LimitOffsetPagination()
    result = pager.paginate_queryset(queryset, request, self)

    # 3.序列化
    ser = BlogSerializers(instance=result, many=True)

    # 4.获取序列化结果 or 分页返回处理
    response = pager.get_paginated_response(ser.data)
    return response
```



### 方法汇总

| 方法                              | 说明                             |
| --------------------------------- | -------------------------------- |
| paginate_queryset()               | 从查询中返回单页数据             |
| get_limit()                       | 从请求参数中获取每页数据条数     |
| get_offset()                      | 从请求参数中获取数据起始地址     |
| get_count()                       | 从queryset中获取数据总条数       |
| get_next_link()                   | 返回下一页数据的URL              |
| get_previous_link()               | 返回上一页数据的URL              |
| get_paginated_response()          | 将分页后的结果封装成Response对象 |
| get_paginated_response_schema()   |                                  |
| get_schema_fields()               |                                  |
| get_schema_operation_parameters() |                                  |
| to_html()                         |                                  |
| get_html_context()                |                                  |

### 变量参数

- limit：每页显示的最大数据量，默认未None即不分页
- offset：返回结果的起始位置，即从第几条开始返回，默认未0，即第一条数据。

```python
# 每页数据量
default_limit = api_settings.PAGE_SIZE

# 每页显示的最大数据量参数名称
limit_query_param = "limit"

# 返回结构的起始位置参数名称
offset_query_param = "offset"

# 允许最大返回的数据量
max_limit = None
template = "rest_framework/pagination/numbers.html"
```

### paginate_queryset

该从查询集数据中获取单页数据并返回

```python
def paginate_queryset(self, queryset, request, view=None):
    """
    queryset:数据查询集，
	"""
    self.limit = self.get_limit(request)
    if self.limit is None:
        return None

    self.count = self.get_count(queryset)
    self.offset = self.get_offset(request)
    self.request = request
    if self.count > self.limit and self.template is not None:
        self.display_page_controls = True

    if self.count == 0 or self.offset > self.count:
        return []
    return list(queryset[self.offset : self.offset + self.limit])
```

![LimitOffsetPagination_paginate_queryset](./assets/LimitOffsetPagination_paginate_queryset.png)

### get_paginated_response

分页后的查询结果封装成Response对象

```Python
def get_paginated_response(self, data):
    return Response(
        OrderedDict(
            [
                ("count", self.count),
                ("next", self.get_next_link()),
                ("previous", self.get_previous_link()),
                ("results", data),
            ]
        )
    )
```

![LimitOffsetPagination_get_paginated_response](./assets/LimitOffsetPagination_get_paginated_response.png)

## PageNumberPagination

### 使用方式

```python
from rest_framework.pagination import PageNumberPagination
def get(self, request, *args, **kwargs):
    # 1.读取数据库中的博客信息
    queryset = models.xxxx.objects.all().order_by("id")
	# 实例化PageNumberPageination类
	pager = PageNumberPagination()
	
	# 对查询集数据进行分页处理
	result = pager.paginate_queryset(queryset, request, self)
	
	# 3.序列化
	ser = BlogSerializers(instance=result, many=True)
	
	# 4.获取序列化结果 or 分页返回处理
	response = pager.get_paginated_response(ser.data)
	return response
```

### 方法汇总

| 方法                     | 说明 |
| ------------------------ | ---- |
| paginate_queryset()      |      |
| get_next_link()          |      |
| get_paginated_response() |      |
| get_page_size            |      |

```python
# 每页返回的数据量
page_size = api_settings.PAGE_SIZE

django_paginator_class = DjangoPaginator

page_query_param = "page"
page_query_description = _("A page number within the paginated result set.")
page_size_query_param = None
page_size_query_description = _("Number of results to return per page.")

max_page_size = None

last_page_strings = ("last",)

template = "rest_framework/pagination/numbers.html"
invalid_page_message = _("Invalid page.")
```

### get_paginated_response

返回数据

```Python
def get_paginated_response(self, data):
    return Response(
        OrderedDict(
            [
                ("count", self.page.paginator.count),
                ("next", self.get_next_link()),
                ("previous", self.get_previous_link()),
                ("results", data),
            ]
        )
    )
```

![PageNumberPagination_get_paginated_response](./assets/PageNumberPagination_get_paginated_response.png)

## _positive_int

返回int类型的数据，用于判端返回结果时否能为0，且数据量不大于cutoff

```python
def _positive_int(integer_string, strict=False, cutoff=None):
    ret = int(integer_string)
    if ret < 0 or (ret == 0 and strict):
        raise ValueError()
    if cutoff:
        return min(ret, cutoff)
    return ret
```

# 视图

**视图继承类概览：**

![视图继承](./assets/视图继承.png)

## APIView 

继承自Django View，在请求到来时，添加了免除csrf、请求封装、版本、认证、权限、限流等功能/`APIView`是DRF中 “顶层” 的视图类，在他的内部主要实现drf基础的组件的使用，例如：版本、认证、权限、限流等。

```python
# views.py

from rest_framework.views import APIView
from rest_framework.response import Response

class UserView(APIView):
    
    # 认证、权限、限流等
    
    def get(self, request):
		# 业务逻辑：查看列表
        return Response({"code": 0, 'data': "..."})

    def post(self, request):
        # 业务逻辑：新建
        return Response({'code': 0, 'data': "..."})
    
class UserDetailView(APIView):
    
	# 认证、权限、限流等
        
    def get(self, request,pk):
		# 业务逻辑：查看某个数据的详细
        return Response({"code": 0, 'data': "..."})

    def put(self, request,pk):
        # 业务逻辑：全部修改
        return Response({'code': 0, 'data': "..."})
    
    def patch(self, request,pk):
        # 业务逻辑：局部修改
        return Response({'code': 0, 'data': "..."})
    
    def delete(self, request,pk):
        # 业务逻辑：删除
        return Response({'code': 0, 'data': "..."})
```

##  GenericAPIView

`GenericAPIView` 继承APIView，在APIView的基础上又增加了一些功能。例如：`get_queryset`、`get_object`等。

实际在开发中一般不会直接继承它，他更多的是担任 `中间人`的角色，为子类提供公共功能。

最大的意义，将数据库查询、序列化类提取到类变量中，后期再提供公共的get/post/put/delete等方法，让开发者只定义类变量，自动实现增删改查。 

![GenericAPIView](./assets/GenericAPIView.png)

## Viewset

![image-20230426203751493](./assets/image-20230426203751493.png)



### CreateAPIView

```python
class CreateAPIView(mixins.CreateModelMixin, GenericAPIView):
    def post(self, request, *args, **kwargs):
        return self.create(request, *args, **kwargs)
```

### DestroyAPIView

```python
class DestroyAPIView(mixins.DestroyModelMixin, GenericAPIView):
    def delete(self, request, *args, **kwargs):
        return self.destroy(request, *args, **kwargs)
```

### GenericAPIView

| 方法                     | 说明                   |
| ------------------------ | ---------------------- |
| get_serializer()         | 获取序列化器类并实例化 |
| get_serializer_context() |                        |

#### get_serializer

```python
def get_serializer(self, *args, **kwargs):
    # 获取序列化器类并实例化
    serializer_class = self.get_serializer_class()
    kwargs.setdefault("context", self.get_serializer_context())
    return serializer_class(*args, **kwargs)
```

#### get_serializer_context

该方法中封装了request对象

```python
def get_serializer_context(self):
    return {"request": self.request, "format": self.format_kwarg, "view": self}
```

## mixins

### CreateModelMixin

![image-20230426212429640](./assets/image-20230426212429640.png)

### DestroyModelMixin

![image-20230426212446280](./assets/image-20230426212446280.png)

### UpdateModelMixin

![image-20230426212527092](./assets/image-20230426212527092.png)

### ListModelMixin

![image-20230426212555201](./assets/image-20230426212555201.png)

### RetrieveModelMixin

![image-20230426212617109](./assets/image-20230426212617109.png)

## 补充：权限

在之前定义权限类时，类中可以定义两个方法：`has_permission` 和 `has_object_permission` 

- `has_permission` ，在请求进入视图之前就会执行。
- `has_object_permission`，当视图中调用 `self.get_object`时就会被调用（删除、更新、查看某个对象时都会调用），一般用于检查对某个对象是否具有权限进行操作。

```python
class PermissionA(BasePermission):
    message = {"code": 1003, 'data': "无权访问"}

    def has_permission(self, request, view):
        exists = request.user.roles.filter(title="员工").exists()
        if exists:
            return True
        return False

    def has_object_permission(self, request, view, obj):
        return True
    
    
```

所以，让我们在编写视图类时，如果是直接获取间接继承了 GenericAPIView，同时内部调用 `get_object`方法，这样在权限中通过 `has_object_permission` 就可以进行权限的处理。

# 路由routers

## 使用方式

- 视图继承APIView

  ```python
  from django.urls import path
  from app01 import views
  
  urlpatterns = [
      path('api/users/', views.UserView.as_view()),  # APIView
  ]
  ```

- 视图继承 `ViewSetMixin`（GenericViewSet、ModelViewSet）

  ```python
  from django.urls import path, re_path, include
  from app01 import views
  
  urlpatterns = [
      path('api/users/', views.UserView.as_view({"get":"list","post":"create"})),
      path('api/users/<int:pk>/', views.UserView.as_view({"get":"retrieve","put":"update","patch":"partial_update","delete":"destory"})),
  ]
  ```

  对于这种形式的路由，drf中提供了更简便的方式：

  ```python
  from rest_framework import routers
  from app01 import views
  
  router = routers.SimpleRouter()
  router.register(r'api/users', views.UserView)
  
  urlpatterns = [
      # 其他URL
      # path('xxxx/', xxxx.as_view()),
  ]
  
  urlpatterns += router.urls
  ```

  

  也可以利用include，给URL加前缀：

  ```python
  from django.urls import path, include
  from rest_framework import routers
  from app01 import views
  
  router = routers.SimpleRouter()
  router.register(r'users', views.UserView)
  
  urlpatterns = [
      path('api/', include((router.urls, 'app_name'), namespace='instance_name')),
      # 其他URL
      # path('forgot-password/', ForgotPasswordFormView.as_view()),
  ]
  ```



## 额外的URL

```python
from rest_framework.viewsets import ModelViewSet
from rest_framework.decorators import action


class XXXModelSerializer(serializers.ModelSerializer):
    class Meta:
        model = models.UserInfo
        fields = "__all__"

        
class XXXView(ModelViewSet):
    queryset = models.UserInfo.objects.all()
    serializer_class = XXXModelSerializer

    # @action(detail=False, methods=['get'], url_path="yyy/(?P<xx>\d+)/xxx")
    # def get_password(self, request, xx, pk=None):
    #     print(xx)
    #     return Response("...")

    # @action(detail=True, methods=['get'], url_path="yyy/(?P<xx>\d+)/xxx")
    # def set_password(self, request, xx, pk=None):
    #     print(xx)
    #     return Response("...")

```

## action装饰器

Django默认的路由分发规则决定了视图函数只能以get、post等请求方式命名，如果想要使用自定义的方式命名，我们可以使用action去映射请求方法名与自定义方法

```python
from rest_framework.decorators import action

@action(detail=False, methods=['get'], url_path="yyy/(?P<xx>\d+)/xxx")
def get_password(self, request, xx, pk=None):
    print(xx)
    return Response("...")

@action(detail=True, methods=['get'], url_path="yyy/(?P<xx>\d+)/xxx")
def set_password(self, request, xx, pk=None):
    print(xx)
    return Response("...")
```

| 参数     | 说明                                                         |
| -------- | ------------------------------------------------------------ |
| methods  | 请求映射的方法                                               |
| url_path | 路径拼接的尾缀，默认是函数名                                 |
| detail   | - True：前缀(perfix)+pk正则分组+尾缀(url_path)<br />- False：前缀(perfix)+尾缀(url_path) |



