# Ajax

## Ajax概述

**传统网页**

+ 网速慢的情况下，页面加载时间长，用户只能等待
+ 表单提交后，如果一项内容不合格，需要重新填写所有表单内容
+ 页面跳转，重新加载页面，造成资源浪费，增加用户等待时间

**Ajax 的用场景**

1. 页面上拉加载更多数据
2. 列表数据无刷新分页
3. 表单项离开焦点数据验证
4. 搜索框提示文字下拉列表

**Ajax的运行环境**

Ajax 技术  <span style="color:red;">需要运行在网站环境中才能生效</span> ，

**运行原理**

Ajax 相当于浏览器发送请求与接收响应的代理人，以实现在不影响用户浏览页面的情况下，局部更新页面数据，从而提高用户体验。

![image-20201109001851935](assets-前端/image-20201109001851935.png)

## Ajax 的实现步骤

```javascript
// 1.创建 Ajax 对象
var xhr = new XMLHttpRequest();

// 2.告诉 Ajax 请求地址以及请求方式
xhr.open('get', 'http://www.example.com');

// 3.发送请求
xhr.send();

// 4.获取服务器端给与客户端的响应数据
xhr.onload = function () {
    console.log(xhr.responseText);}
```

## 服务器端响应的数据格式

在 http 请求与响应的过程中，无论是请求参数还是响应内容，如果是对象类型，最终都会被转换为对象字符串进行传输。

```JavaScript
JSON.parse()   // 将 json 字符串转换为json对象
```

## 请求参数传递

传统网站表单提交

```html
<form method="get" action="http://www.example.com">
    <input type="text" name="username"/>
    <input type="password" name="password">
</form><!– http://www.example.com?username=zhangsan&password=123456 -->
```


+ GET 请求方式

  ```javascript
  xhr.open('get', 'http://www.example.com?name=zhangsan&age=20');
  ```

+ POST 请求方式

  ```javascript
  xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded')
  xhr.send('name=zhangsan&age=20');
  ```

## 请求报文

在 HTTP 请求和响应的过程中传递的数据块就叫报文，包括要传送的数据和一些附加信息，这些数据和信息要遵守规定好的格式。

![image-20201110141635781](assets-前端/image-20201110141635781.png)

## 请求参数的格式

1. application/x-www-form-urlencoded

   ```
   name=zhangsan&age=20&sex=男
   ```

2. application/json

   ```
   {name: 'zhangsan', age: '20', sex: '男'}
   ```

   在请求头中指定 Content-Type 属性的值是 application/json，告诉服务器端当前请求参数的格式是 json。

   ```
   JSON.stringify() // 将json对象转换为json字符串
   ```

 <span style="color:red;">注意：get 请求是不能提交 json 对象数据格式的，传统网站的表单提交也是不支持 json 对象数据格式的。
</span> 

## Ajax状态码

在创建ajax对象，配置ajax对象，发送请求，以及接收完服务器端响应数据，这个过程中的每一个步骤都会对应一个数值，这个数值就是ajax状态码。

```JavaScript
xhr.readyState // 获取Ajax状态码
```

| 数值 | 说明                                             |
| ---- | ------------------------------------------------ |
| 0    | 请求未初始化(还没有调用open())                   |
| 1    | 请求已经建立，但是还没有发送(还没有调用send())   |
| 2    | 请求已经发送                                     |
| 3    | 请求正在处理中，通常响应中已经有部分数据可以用了 |
| 4    | 响应已经完成，可以获取并使用服务器的响应了       |

## Ajax 错误处理

1. 网络畅通，服务器端能接收到请求，服务器端返回的结果不是预期结果。

   > 可以判断服务器端返回的状态码，分别进行处理。xhr.status 获取http状态码

2. 网络畅通，服务器端没有接收到请求，返回404状态码。

   > 检查请求地址是否错误。

3. 网络畅通，服务器端能接收到请求，服务器端返回500状态码。

   > 服务器端错误，找后端程序员进行沟通。

4. 网络中断，请求无法发送到服务器端。

   > 会触发xhr对象下面的onerror事件，在onerror事件处理函数中对错误进行处理。

## Ajax 封装

发送一次请求代码过多，发送多次请求代码冗余且重复。
解决方案：将请求代码封装到函数中，发请求时调用函数即可。

```javascript
function ajax(options) {
    // 存储的是默认值
    var defaults = {
        type: 'get',
        url: '',
        data: {},
        header: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        success: function () {
        },
        error: function () {
        }
    };
    // 使用options对象中的属性覆盖defaults对象中的属性
    Object.assign(defaults, options);
    // 创建ajax对象
    var xhr = new XMLHttpRequest();
    // 拼接请求参数的变量
    var params = '';
    // 循环用户传递进来的对象格式参数
    for (var attr in defaults.data) {
        // 将参数转换为字符串格式
        params += attr + '=' + defaults.data[attr] + '&';
    }
    // 将参数最后面的&截取掉
    // 将截取的结果重新赋值给params变量
    params = params.substr(0, params.length - 1); 
    // 判断请求方式 
    if (defaults.type == 'get') {
        defaults.url = defaults.url + '?' + params;
    }
    /*
    {
    name: 'zhangsan',
    age: 20
    }
    name=zhangsan&age=20
    */ 
    // 配置ajax对象
    xhr.open(defaults.type, defaults.url);
    // 如果请求方式为post
    if (defaults.type == 'post') {
        // 用户希望的向服务器端传递的请求参数的类型
        var contentType = defaults.header['Content-Type']
        // 设置请求参数格式的类型
        xhr.setRequestHeader('Content-Type', contentType);
        // 判断用户希望的请求参数格式的类型
        // 如果类型为json
        if (contentType == 'application/json') {
            // 向服务器端传递json数据格式的参数
            xhr.send(JSON.stringify(defaults.data))
        } 
        else {
            // 向服务器端传递普通类型的请求参数
            xhr.send(params);
        }
    }
    else {
        // 发送请求
        xhr.send();
    }
    // 监听xhr对象下面的onload事件
    // 当xhr对象接收完响应数据后触发
    xhr.onload = function () {
        // xhr.getResponseHeader()
        // 获取响应头中的数据
        var contentType = xhr.getResponseHeader('Content-Type');
        // 服务器端返回的数据
        var responseText = xhr.responseText;
        // 如果响应类型中包含applicaition/json
        if (contentType.includes('application/json')) {
            // 将json字符串转换为json对象
            responseText = JSON.parse(responseText)
        }
        /*  请求成功或者失败 */
        // 当 http 状态码等于 200 的时候
        if (xhr.status == 200) {
            // 请求成功 调用处理成功情况的函数
            defaults.success(responseText, xhr);
        } 
        else {
            // 请求失败 调用处理失败情况的函数
            defaults.error(responseText, xhr);
        } 
    }
}
```



```javascript
ajax({
    type: 'get',
    url: 'http://www.example.com',
    success: function (data) {
        console.log(data);
    }})
```

### $.ajax()

#### $.ajax()方法概述

作用：发送Ajax请求。

```javascript
$.ajax({
    type: 'get',
    url: 'http://www.example.com', 
    data: { name: 'zhangsan', age: '20' },
    contentType: 'application/x-www-form-urlencoded',
    beforeSend: function () {
        return false 
    },
    success: function (response) {},
    error: function (xhr) {}});
```

```javascript
{    data: 'name=zhangsan&age=20'}
```

```javascript
{    contentType: 'application/json'}
```

```javascript
JSON.stringify({name: 'zhangsan', age: '20'})
```

#### serialize方法

```html
<form id="form">
    <input type="text" name="username">
    <input type="password" name="password">
    <input type="submit" value="提交">
</form>
<script type="text/javascript">
    // 将表单内容拼接成字符串类型的参数
    var params = $('#form').serialize()
    console.log(params)
</script>
```

#### `$.get()、$.post()`方法概述

作用：$.get方法用于发送get请求，$.post方法用于发送post请求。

```javascript
$.get('http://www.example.com',
      {name: 'zhangsan', age: 30},
      function (response) {})

$.post('http://www.example.com',
       {name: 'lisi', age: 22},
       function (response) {})
```

#### 案例：省市区联动

```html
<link rel="stylesheet" href="bootstrap.min.css"><style type="text/css">    .container {        padding-top: 150px;    }</style>
```

```html
<div class="container">    <div class="form-inline">        <div class="form-group">            <select class="form-control" id="province"></select>        </div>        <div class="form-group">            <select class="form-control" id="city">                <option>请选择城市</option>            </select>        </div>        <div class="form-group">            <select class="form-control" id="area">                <option>请选择县城</option>            </select>        </div>    </div></div>
```

```html
<script src="ajax.js"></script><script>    // 获取省市区下拉框元素    var province = document.getElementById('province');    var city = document.getElementById('city');    var area = document.getElementById('area');    // 获取省份信息    ajax({        type: 'get',        url: 'http://localhost:5000/province',        success: function (data) {            // 将服务器端返回的数据和html进行拼接            // {#var html = template('provinceTpl', {province: data});#}            html = "<option>请选择省份</option>"            data.forEach(function (item) {                html += "<option value=" + item.id + ">" + item.name + "</option>"            })            // 将拼接好的html字符串显示在页面中            province.innerHTML = html;        }    });    // 为省份的下拉框添加值改变事件    province.onchange = function () {        // 获取省份id        var pid = this.value;        // 清空县城下拉框中的数据        var html = "<option>请选择县城</option>";        area.innerHTML = html;        // 根据省份id获取城市信息        ajax({            type: 'get',            url: 'http://127.0.0.1:5000/cities',            data: {                id: pid            },            success: function (data) {                var html = "<option>请选择县城</option>"                data.forEach(function (item) {                    html += '<option value="'+item.id+'">'+item.name+'</option>'                })                city.innerHTML = html;            }        })    };    // 当用户选择城市的时候    city.onchange = function () {        // 获取城市id        var cid = this.value;        // 根据城市id获取县城信息        ajax({            type: 'get',            url: 'http://localhost:5000/areas',            data: {                id: cid            },            success: function (data) {                var html = '<option>请选择县城</option>';                data.forEach(function (item) {                    html += '<option value="'+item.id+'">'+item.name+'</option>'                })                area.innerHTML = html;            }        })    }</script>
```

### 附录

restful 风格指南：http://www.ruanyifeng.com/blog/2014/05/restful_api.html

# BOM

虽然 ECMAScript 把浏览器对象模型（BOM，Browser Object Model）描述为 JavaScript 的核心，但实际上 BOM 是使用 JavaScript 开发 Web 应用程序的核心。BOM 提供了与网页无关的浏览器功能对象。多年来，BOM 是在缺乏规范的背景下发展起来的，因此既充满乐趣又问题多多。毕竟，浏览器开发商都按照自己的意愿来为它添砖加瓦。最终，浏览器实现之间共通的部分成为了事实标准，为 Web 开发提供了浏览器间互操作的基础。HTML5 规范中有一部分涵盖了 BOM 的主要内容，因为 W3C 希望将 JavaScript 在浏览器中最基础的部分标准化。

## window 对象

BOM 的核心是 window 对象，表示浏览器的实例。window 对象在浏览器中有两重身份，一个是 ECMAScript 中的 Global 对象，另一个就是浏览器窗口的 JavaScript 接口。这意味着网页中定义的所有对象、变量和函数都以 window 作为其 Global 对象，都可以访问其上定义的 parseInt()等全局方法。

### **Global** 作用域

因为 window 对象被复用为 ECMAScript 的 Global 对象，所以通过 var 声明的所有全局变量和函数都会变成 window 对象的属性和方法。比如：

```javascript
var age = 29; 
var sayAge = () => alert(this.age); 

alert(window.age); // 29
sayAge(); // 29 
window.sayAge(); // 29
```

这里，变量 age 和函数 sayAge()被定义在全局作用域中，它们自动成为了 window 对象的成员。因此，变量 age 可以通过 window.age 来访问，而函数 sayAge()也可以通过 window.sayAge()来访问。因为 sayAge()存在于全局作用域，this.age 映射到 window.age，所以就可以显示正确的结果了。

如果在这里使用 let 或 const 替代 var，则不会把变量添加给全局对象：

```javascript
let age = 29; 
const sayAge = () => alert(this.age); 

alert(window.age); // undefined 
sayAge(); // undefined 
window.sayAge(); // TypeError: window.sayAge is not a function
```

另外，访问未声明的变量会抛出错误，但是可以在 window 对象上查询是否存在可能未声明的变量。比如：

```javascript
// 这会导致抛出错误，因为 oldValue 没有声明
var newValue = oldValue; 
// 这不会抛出错误，因为这里是属性查询
// newValue 会被设置为 undefined 
var newValue = window.oldValue;
```

记住，JavaScript 中有很多对象都暴露在全局作用域中，比如 location 和 navigator（本章后面都会讨论），因而它们也是 window 对象的属性。

### 窗口关系

top 对象始终指向最上层（最外层）窗口，即浏览器窗口本身。而 parent 对象则始终指向当前窗口的父窗口。如果当前窗口是最上层窗口，则 parent 等于 top（都等于 window）。最上层的 window

如果不是通过 window.open()打开的，那么其 name 属性就不会包含值，本章后面会讨论。

还有一个 self 对象，它是终极 window 属性，始终会指向 window。实际上，self 和 window 就是同一个对象。之所以还要暴露 self，就是为了和 top、parent 保持一致。

这些属性都是 window 对象的属性，因此访问 window.parent、window.top 和 window.self 都可以。这意味着可以把访问多个窗口的 window 对象串联起来，比如 window.parent.parent。

### 窗口位置与像素比

window 对象的位置可以通过不同的属性和方法来确定。现代浏览器提供了 screenLeft 和 screenTop 属性，用于表示窗口相对于屏幕左侧和顶部的位置 ，返回值的单位是 CSS 像素。

可以使用 moveTo()和 moveBy()方法移动窗口。这两个方法都接收两个参数，其中 moveTo() 接收要移动到的新位置的绝对坐标 *x* 和 *y*；而 moveBy() 则接收相对当前位置在两个方向上移动的像素数。

比如：

```javascript
// 把窗口移动到左上角
window.moveTo(0,0);

// 把窗口向下移动 100 像素
window.moveBy(0, 100);

// 把窗口移动到坐标位置(200, 300) 
window.moveTo(200, 300);

// 把窗口向左移动 50 像素
window.moveBy(-50, 0);
```

依浏览器而定，以上方法可能会被部分或全部禁用。

**像素比**

CSS 像素是 Web 开发中使用的统一像素单位。这个单位的背后其实是一个角度：0.0213°。如果屏幕距离人眼是一臂长，则以这个角度计算的 CSS 像素大小约为 1/96 英寸。这样定义像素大小是为了在不同设备上统一标准。比如，低分辨率平板设备上 12 像素（CSS 像素）的文字应该与高清 4K 屏幕下 12 像素（CSS 像素）的文字具有相同大小。这就带来了一个问题，不同像素密度的屏幕下就会有不同的缩放系数，以便把物理像素（屏幕实际的分辨率）转换为 CSS 像素（浏览器报告的虚拟分辨率）。

举个例子，手机屏幕的物理分辨率可能是 1920×1080，但因为其像素可能非常小，所以浏览器就需要将其分辨率降为较低的逻辑分辨率，比如 640×320。这个物理像素与 CSS 像素之间的转换比率由 window.devicePixelRatio 属性提供。对于分辨率从 1920×1080 转换为 640×320 的设备，window.devicePixelRatio 的值就是 3。这样一来，12 像素（CSS 像素）的文字实际上就会用 36 像素的物理像素来显示。

window.devicePixelRatio 实际上与每英寸像素数（DPI，dots per inch）是对应的。DPI 表示单位像素密度，而 window.devicePixelRatio 表示物理像素与逻辑像素之间的缩放系数。

### 窗口大小

在不同浏览器中确定浏览器窗口大小没有想象中那么容易。所有现代浏览器都支持 4 个属性：innerWidth、innerHeight、outerWidth 和 outerHeight。outerWidth 和 outerHeight 返回浏览器窗口自身的大小（不管是在最外层 window 上使用，还是在窗格<frame>中使用）。innerWidth 和 innerHeight 返回浏览器窗口中页面视口的大小（不包含浏览器边框和工具栏）。

document.documentElement.clientWidth 和 document.documentElement.clientHeight 返回页面视口的宽度和高度。

浏览器窗口自身的精确尺寸不好确定，但可以确定页面视口的大小，如下所示：

```javascript
let pageWidth = window.innerWidth,
    pageHeight = window.innerHeight;
if (typeof pageWidth != 'number') {
    if (document.compatMode == 'CSS1Compat') {
        pageWidth = document.documentElement.clientWidth;
        pageHeight = document.documentElement.clientHeight;
    } else {
        pageWidth = document.body.clientWidth;
        pageHeight = document.body.clientHeight;
    }
}
```

这里，先将 pageWidth 和 pageHeight 的值分别设置为 window.innerWidth 和 window.innerHeight。然后，检查 pageWidth 是不是一个数值，如果不是则通过 document.compatMode 来检查页面是否处于标准模式。如果是，则使用  document.documentElement.clientWidth 和 document.documentElement.clientHeight；否则，就使用 document.body.clientWidth 和 document.body.clientHeight。

在移动设备上，window.innerWidth 和 window.innerHeight 返回视口的大小，也就是屏幕上页面可视区域的大小。Mobile Internet Explorer 支持这些属性，但在 document.documentElement.clientWidth 和 document.documentElement.clientHeight 中提供了相同的信息。在放大或缩小页面时，这些值也会相应变化。

在其他移动浏览器中，document.documentElement.clientWidth 和 document.documentElement.clientHeight 返回的布局视口的大小，即渲染页面的实际大小。布局视口是相对于可见视口的概念，可见视口只能显示整个页面的一小部分。Mobile Internet Explorer 把布局视口的信息保存在 document.body.clientWidth 和 document.body.clientHeight 中。在放大或缩小页面时，这些值也会相应变化。

因为桌面浏览器的差异，所以需要先确定用户是不是在使用移动设备，然后再决定使用哪个属性。

可以使用resizeTo()和resizeBy()方法调整窗口大小。这两个方法都接收两个参数，resizeTo() 接收新的宽度和高度值，而 resizeBy()接收宽度和高度各要缩放多少。下面看个例子：

```
// 缩放到 100×100 
window.resizeTo(100, 100); 

// 缩放到 200×150 
window.resizeBy(100, 50); 

// 缩放到 300×300 
window.resizeTo(300, 300);
```

与移动窗口的方法一样，缩放窗口的方法可能会被浏览器禁用，而且在某些浏览器中默认是禁用的。同样，缩放窗口的方法只能应用到最上层的 window 对象。

### 视口位置

浏览器窗口尺寸通常无法满足完整显示整个页面，为此用户可以通过滚动在有限的视口中查看文档。度量文档相对于视口滚动距离的属性有两对，返回相等的值：window.pageXoffset/window. scrollX 和 window.pageYoffset/window.scrollY。

可以使用 scroll()、scrollTo()和 scrollBy()方法滚动页面。这 3 个方法都接收表示相对视口距离的 *x* 和 *y* 坐标，这两个参数在前两个方法中表示要滚动到的坐标，在最后一个方法中表示滚动的距离。

```
// 相对于当前视口向下滚动 100 像素
window.scrollBy(0, 100); 

// 相对于当前视口向右滚动 40 像素
window.scrollBy(40, 0); 

// 滚动到页面左上角
window.scrollTo(0, 0);

// 滚动到距离屏幕左边及顶边各 100 像素的位置
window.scrollTo(100, 100);
```

这几个方法也都接收一个 ScrollToOptions 字典，除了提供偏移值，还可以通过 behavior 属性告诉浏览器是否平滑滚动。

```
// 正常滚动 
window.scrollTo({ 
 left: 100, 
 top: 100, 
 behavior: 'auto' 
}); 

// 平滑滚动
window.scrollTo({ 
 left: 100, 
 top: 100, 
 behavior: 'smooth' 
});
```

### 导航与打开新窗口

window.open()方法可以用于导航到指定 URL，也可以用于打开新浏览器窗口。这个方法接收 4 个参数：要加载的 URL、目标窗口、特性字符串和表示新窗口在浏览器历史记录中是否替代当前加载页面的布尔值。通常，调用这个方法时只传前 3 个参数，最后一个参数只有在不打开新窗口时才会使用。

如果 window.open()的第二个参数是一个已经存在的窗口或窗格（frame）的名字，则会在对应的窗口或窗格中打开 URL。下面是一个例子：

```
// 与<a href="http://www.wrox.com" target="topFrame"/>相同
window.open("http://www.wrox.com/", "topFrame");
```

执行这行代码的结果就如同用户点击了一个 href 属性为"http://www.wrox.com"，target 属性为"topFrame"的链接。如果有一个窗口名叫"topFrame"，则这个窗口就会打开这个 URL；否则就会打开一个新窗口并将其命名为"topFrame"。第二个参数也可以是一个特殊的窗口名，比如`_self`、`_parent`、`_top` 或_blank。

#### 1. 弹出窗口

如果 window.open()的第二个参数不是已有窗口，则会打开一个新窗口或标签页。第三个参数，即特性字符串，用于指定新窗口的配置。如果没有传第三个参数，则新窗口（或标签页）会带有所有默认的浏览器特性（工具栏、地址栏、状态栏等都是默认配置）。如果打开的不是新窗口，则忽略第三个参数。

特性字符串是一个逗号分隔的设置字符串，用于指定新窗口包含的特性。下表列出了一些选项。

| 设 置      | 值          | 说 明                                                        |
| ---------- | ----------- | ------------------------------------------------------------ |
| fullscreen | "yes"或"no" | 表示新窗口是否最大化。仅限 IE 支持                           |
| height     | 数值        | 新窗口高度。这个值不能小于 100                               |
| left       | 数值        | 新窗口的 *x* 轴坐标。这个值不能是负值                        |
| location   | "yes"或"no" | 表示是否显示地址栏。不同浏览器的默认值也不一样。在设置为"no"时，地址栏可能隐藏或禁用（取决于浏览器） |
| Menubar    | "yes"或"no" | 表示是否显示菜单栏。默认为"no"                               |
| resizable  | "yes"或"no" | 表示是否可以拖动改变新窗口大小。默认为"no"                   |
| scrollbars | "yes"或"no" | 表示是否可以在内容过长时滚动。默认为"no"                     |
| status     | "yes"或"no" | 表示是否显示状态栏。不同浏览器的默认值也不一样               |
| toolbar    | "yes"或"no" | 表示是否显示工具栏。默认为"no"                               |
| top        | 数值        | 新窗口的 *y* 轴坐标。这个值不能是负值                        |
| width      | 数值        | 新窗口的宽度。这个值不能小于 100                             |

这些设置需要以逗号分隔的名值对形式出现，其中名值对以等号连接。（特性字符串中不能包含空格。）来看下面的例子：

```
window.open("http://www.wrox.com/", "wroxWindow",  "height=400,width=400,top=10,left=10,resizable=yes"); 
```

这行代码会打开一个可缩放的新窗口，大小为 400 像素×400 像素，位于离屏幕左边及顶边各 10 像素的位置。window.open()方法返回一个对新建窗口的引用。这个对象与普通 window 对象没有区别，只是为控制新窗口提供了方便。例如，某些浏览器默认不允许缩放或移动主窗口，但可能允许缩放或移动通过 window.open() 创建的窗口。跟使用任何 window 对象一样，可以使用这个对象操纵新打开的窗口。

```
let wroxWin = window.open("http://www.wrox.com/", 
    "wroxWindow", 
    "height=400,width=400,top=10,left=10,resizable=yes"); 
    
// 缩放
wroxWin.resizeTo(500, 500); 

// 移动
wroxWin.moveTo(100, 100);
```

还可以使用 close()方法像这样关闭新打开的窗口：

```
wroxWin.close();
```

这个方法只能用于 window.open()创建的弹出窗口。虽然不可能不经用户确认就关闭主窗口，但弹出窗口可以调用 top.close()来关闭自己。关闭窗口以后，窗口的引用虽然还在，但只能用于检查其 closed 属性了：

```
wroxWin.close(); 
alert(wroxWin.closed); // true
```

新创建窗口的 window 对象有一个属性 opener，指向打开它的窗口。这个属性只在弹出窗口的最上层 window 对象（top）有定义，是指向调用 window.open()打开它的窗口或窗格的指针。例如：

```
let wroxWin = window.open("http://www.wrox.com/", 
 "wroxWindow", 
 "height=400,width=400,top=10,left=10,resizable=yes"); 
 
alert(wroxWin.opener === window); // true
```

虽然新建窗口中有指向打开它的窗口的指针，但反之则不然。窗口不会跟踪记录自己打开的新窗口，因此开发者需要自己记录。

在某些浏览器中，每个标签页会运行在独立的进程中。如果一个标签页打开了另一个，而 window 对象需要跟另一个标签页通信，那么标签便不能运行在独立的进程中。在这些浏览器中，可以将新打开的标签页的 opener 属性设置为 null，表示新打开的标签页可以运行在独立的进程中。比如：

```
let wroxWin = window.open("http://www.wrox.com/", 
 "wroxWindow", 
 "height=400,width=400,top=10,left=10,resizable=yes"); 
 
wroxWin.opener = null;
```

把 opener 设置为 null 表示新打开的标签页不需要与打开它的标签页通信，因此可以在独立进程中运行。这个连接一旦切断，就无法恢复了。

#### 2. 安全限制

弹出窗口有段时间被在线广告用滥了。很多在线广告会把弹出窗口伪装成系统对话框，诱导用户点击。因为长得像系统对话框，所以用户很难分清这些弹窗的来源。为了让用户能够区分清楚，浏览器开始对弹窗施加限制。

IE 的早期版本实现针对弹窗的多重安全限制，包括不允许创建弹窗或把弹窗移出屏幕之外，以及不允许隐藏状态栏等。从 IE7 开始，地址栏也不能隐藏了，而且弹窗默认是不能移动或缩放的。Firefox 1禁用了隐藏状态栏的功能，因此无论 window.open()的特性字符串是什么，都不会隐藏弹窗的状态栏。

Firefox 3 强制弹窗始终显示地址栏。Opera 只会在主窗口中打开新窗口，但不允许它们出现在系统对话框的位置。

此外，浏览器会在用户操作下才允许创建弹窗。在网页加载过程中调用 window.open()没有效果，而且还可能导致向用户显示错误。弹窗通常可能在鼠标点击或按下键盘中某个键的情况下才能打开。

#### 3. 弹窗屏蔽程序

所有现代浏览器都内置了屏蔽弹窗的程序，因此大多数意料之外的弹窗都会被屏蔽。在浏览器屏蔽弹窗时，可能会发生一些事。如果浏览器内置的弹窗屏蔽程序阻止了弹窗，那么 window.open()很可能会返回 null。此时，只要检查这个方法的返回值就可以知道弹窗是否被屏蔽了，比如：

```
let wroxWin = window.open("http://www.wrox.com", "_blank"); 
if (wroxWin == null){ 
 alert("The popup was blocked!"); 
}
```

在浏览器扩展或其他程序屏蔽弹窗时，window.open()通常会抛出错误。因此要准确检测弹窗是否被屏蔽，除了检测 window.open()的返回值，还要把它用 try/catch 包装起来，像这样：

```
let blocked = false;
try { 
 let wroxWin = window.open("http://www.wrox.com", "_blank"); 
 if (wroxWin == null){
  blocked = true; 
 } 
} catch (ex){ 
 blocked = true; 
} 
if (blocked){ 
 alert("The popup was blocked!"); 
}
```

无论弹窗是用什么方法屏蔽的，以上代码都可以准确判断调用 window.open()的弹窗是否被屏蔽了。

### 定时器

JavaScript 在浏览器中是单线程执行的，但允许使用定时器指定在某个时间之后或每隔一段时间就执行相应的代码。setTimeout()用于指定在一定时间后执行某些代码，而 setInterval()用于指定每隔一段时间执行某些代码。

setTimeout()方法通常接收两个参数：要执行的代码和在执行回调函数前等待的时间（毫秒）。第一个参数可以是包含 JavaScript 代码的字符串（类似于传给 eval()的字符串）或者一个函数，比如：

```JavaScript
// 在 1 秒后显示警告框
setTimeout(() => alert("Hello world!"), 1000);
```

第二个参数是要等待的毫秒数，而不是要执行代码的确切时间。JavaScript 是单线程的，所以每次只能执行一段代码。为了调度不同代码的执行，JavaScript 维护了一个任务队列。其中的任务会按照添加到队列的先后顺序执行。setTimeout()的第二个参数只是告诉 JavaScript 引擎在指定的毫秒数过后把任务添加到这个队列。如果队列是空的，则会立即执行该代码。如果队列不是空的，则代码必须等待前面的任务执行完才能执行。

调用 setTimeout()时，会返回一个表示该超时排期的数值 ID。这个超时 ID 是被排期执行代码的唯一标识符，可用于取消该任务。要取消等待中的排期任务，可以调用 clearTimeout()方法并传入超时 ID，如下面的例子所示：

```JavaScript
// 设置超时任务
let timeoutId = setTimeout(() => alert('Hello world!'), 1000);

// 取消超时任务
clearTimeout(timeoutId);
```

只要是在指定时间到达之前调用 clearTimeout()，就可以取消超时任务。在任务执行后再调用 clearTimeout()没有效果。

setInterval()与 setTimeout()的使用方法类似，只不过指定的任务会每隔指定时间就执行一次，直到取消循环定时或者页面卸载。setInterval()同样可以接收两个参数：要执行的代码（字符串或函数），以及把下一次执行定时代码的任务添加到队列要等待的时间（毫秒）。下面是一个例子：

```JavaScript
setInterval(() => alert("Hello world!"), 10000);
```

setInterval()方法也会返回一个循环定时 ID，可以用于在未来某个时间点上取消循环定时。要取消循环定时，可以调用 clearInterval()并传入定时 ID。相对于 setTimeout()而言，取消定时的能力对 setInterval()更加重要。毕竟，如果一直不管它，那么定时任务会一直执行到页面卸载。下面是一个常见的例子：

```JavaScript
let num = 0, intervalId = null;
let max = 10;
let incrementNumber = function() {
    num++;
    // 如果达到最大值，则取消所有未执行的任务
    if (num == max) {
        clearInterval(intervalId);
        alert('Done');
    }
};
intervalId = setInterval(incrementNumber, 500);
```

在这个例子中，变量 num 会每半秒递增一次，直至达到最大限制值。此时循环定时会被取消。这个模式也可以使用 setTimeout()来实现，比如：

```JavaScript
let num = 0;
let max = 10;
let incrementNumber = function() {
    num++;
    // 如果还没有达到最大值，再设置一个超时任务
    if (num < max) {
        setTimeout(incrementNumber, 500);
    } else {
        alert('Done');
    }
};

setTimeout(incrementNumber, 500);
```

注意在使用 setTimeout()时，不一定要记录超时 ID，因为它会在条件满足时自动停止，否则会自动设置另一个超时任务。这个模式是设置循环任务的推荐做法。setIntervale()在实践中很少会在生产环境下使用，因为一个任务结束和下一个任务开始之间的时间间隔是无法保证的，有些循环定时任务可能会因此而被跳过。而像前面这个例子中一样使用 setTimeout()则能确保不会出现这种情况。一般来说，最好不要使用 setInterval()。

### 系统对话框

使用 alert()、confirm()和 prompt()方法，可以让浏览器调用系统对话框向用户显示消息。这些对话框与浏览器中显示的网页无关，而且也不包含 HTML。它们的外观由操作系统或者浏览器决定，无法使用 CSS 设置。此外，这些对话框都是同步的模态对话框，即在它们显示的时候，代码会停止执行，在它们消失以后，代码才会恢复执行。

alert()方法在本书示例中经常用到。它接收一个要显示给用户的字符串。与 console.log 可以接收任意数量的参数且能一次性打印这些参数不同，alert()只接收一个参数。调用 alert()时，传入的字符串会显示在一个系统对话框中。对话框只有一个“OK”（确定）按钮。如果传给 alert()的参数不是一个原始字符串，则会调用这个值的 toString()方法将其转换为字符串。

警告框（alert）通常用于向用户显示一些他们无法控制的消息，比如报错。用户唯一的选择就是在看到警告框之后把它关闭。



第二种对话框叫确认框，通过调用 confirm()来显示。确认框跟警告框类似，都会向用户显示消息。但不同之处在于，确认框有两个按钮：“Cancel”（取消）和“OK”（确定）。用户通过单击不同的按钮表明希望接下来执行什么操作。比如，confirm("Are you sure?")会显示确认框。

要知道用户单击了 OK 按钮还是 Cancel 按钮，可以判断 confirm()方法的返回值：true 表示单击了 OK 按钮，false 表示单击了 Cancel 按钮或者通过单击某一角上的 X 图标关闭了确认框。确认框的典型用法如下所示：

```javascript
if (confirm('Are you sure?')) {
    alert('I\'m so glad you\'re sure!');
} else {
    alert('I\'m sorry to hear you\'re not sure.');
}
```

在这个例子中，第一行代码向用户显示了确认框，也就是 if 语句的条件。如果用户单击了 OK 按钮，

则会弹出警告框显示"I'm so glad you're sure!"。如果单击了 Cancel，则会显示"I'm sorry to hear you're not sure."。确认框通常用于让用户确认执行某个操作，比如删除邮件等。因为这种对话框会完全打断正在浏览网页的用户，所以应该在必要时再使用。

最后一种对话框是提示框，通过调用 prompt()方法来显示。提示框的用途是提示用户输入消息。除了 OK 和 Cancel 按钮，提示框还会显示一个文本框，让用户输入内容。prompt()方法接收两个参数：要显示给用户的文本，以及文本框的默认值（可以是空字符串）。调用 prompt("What is your name?", "Jake")会显示提示框。



如果用户单击了 OK 按钮，则 prompt()会返回文本框中的值。如果用户单击了 Cancel 按钮，或者对话框被关闭，则 prompt()会返回 null。下面是一个例子：

```javascript
let result = prompt('What is your name? ', '');
if (result !== null) {
    alert('Welcome, ' + result);
}
```

这些系统对话框可以向用户显示消息、确认操作和获取输入。由于不需要 HTML 和 CSS，所以系统对话框是 Web 应用程序最简单快捷的沟通手段。

很多浏览器针对这些系统对话框添加了特殊功能。如果网页中的脚本生成了两个或更多系统对话框，则除第一个之外所有后续的对话框上都会显示一个复选框，如果用户选中则会禁用后续的弹框，直到页面刷新。

如果用户选中了复选框并关闭了对话框，在页面刷新之前，所有系统对话框（警告框、确认框、提示框）都会被屏蔽。开发者无法获悉这些对话框是否显示了。对话框计数器会在浏览器空闲时重置，因此如果两次独立的用户操作分别产生了两个警告框，则两个警告框上都不会显示屏蔽复选框。如果一次独立的用户操作连续产生了两个警告框，则第二个警告框会显示复选框。

JavaScript 还可以显示另外两种对话框：find()和 print()。这两种对话框都是异步显示的，即控制权会立即返回给脚本。用户在浏览器菜单上选择“查找”（find）和“打印”（print）时显示的就是这两种对话框。通过在 window 对象上调用 find()和 print()可以显示它们，比如：

```javascript
// 显示打印对话框
window.print();

// 显示查找对话框
window.find();
```

这两个方法不会返回任何有关用户在对话框中执行了什么操作的信息，因此很难加以利用。此外，因为这两种对话框是异步的，所以浏览器的对话框计数器不会涉及它们，而且用户选择禁用对话框对它们也没有影响。

## **location** 对象

location 是最有用的 BOM 对象之一，提供了当前窗口中加载文档的信息，以及通常的导航功能。这个对象独特的地方在于，它既是 window 的属性，也是 document 的属性。也就是说，window.location 和 document.location 指向同一个对象。location 对象不仅保存着当前加载文档的信息，也保存着把 URL 解析为离散片段后能够通过属性访问的信息。这些解析后的属性在下表中有详细说明（location 前缀是必需的）。

假设浏览器当前加载的 URL 是 http://foouser:barpassword@www.wrox.com:80/WileyCDA/?q=javascript#contents，location 对象的内容如下表所示。

| 属 性             | 值                                                       | 说 明                                                        |
| ----------------- | -------------------------------------------------------- | ------------------------------------------------------------ |
| location.hash     | "#contents"                                              | URL 散列值（井号后跟零或多个字符），如果没有则为空字符串     |
| location.host     | "www.wrox.com:80"                                        | 服务器名及端口号                                             |
| location.hostname | "www.wrox.com"                                           | 服务器名                                                     |
| location.href     | "http://www.wrox.com:80/WileyCDA/?q=javascript#contents" | 当前加载页面的完整 URL。location 的 toString() 方法返回这个值 |
| location.pathname | "/WileyCDA/"                                             | URL 中的路径和（或）文件名                                   |
| location.port     | "80"                                                     | 请求的端口。如果 URL中没有端口，则返回空字符串               |
| location.protocol | "http:"                                                  | 页面使用的协议。通常是"http:"或"https:"                      |
| location.search   | "?q=javascript"                                          | URL 的查询字符串。这个字符串以问号开头                       |
| location.username | "foouser"                                                | 域名前指定的用户名                                           |
| location.password | "barpassword"                                            | 域名前指定的密码                                             |
| location.origin   | "http://www.wrox.com"                                    | URL 的源地址。只读                                           |

### 查询字符串

location 的多数信息都可以通过上面的属性获取。但是 URL 中的查询字符串并不容易使用。虽然location.search 返回了从问号开始直到 URL 末尾的所有内容，但没有办法逐个访问每个查询参数。下面的函数解析了查询字符串，并返回一个以每个查询参数为属性的对象：

```javascript
let getQueryStringArgs = function() {
    // 取得没有开头问号的查询字符串
    let qs = (location.search.length > 0 ? location.search.substring(1) : ''),
        // 保存数据的对象
        args = {};
    // 把每个参数添加到 args 对象
    for (let item of qs.split('&').map(kv => kv.split('='))) {
        let name = decodeURIComponent(item[0]),
            value = decodeURIComponent(item[1]);
        if (name.length) {
            args[name] = value;
        }
    }
    return args;
};
```

这个函数首先删除了查询字符串开头的问号，当然前提是 location.search 必须有内容。解析后的参数将被保存到 args 对象，这个对象以字面量形式创建。接着，先把查询字符串按照&分割成数组，每个元素的形式为 name=value。for 循环迭代这个数组，将每一个元素按照=分割成数组，这个数组第一项是参数名，第二项是参数值。参数名和参数值在使用 decodeURIComponent()解码后（这是因为查询字符串通常是被编码后的格式）分别保存在 name 和 value 变量中。最后，name 作为属性而 value作为该属性的值被添加到 args 对象。这个函数可以像下面这样使用：

```javascript
// 假设查询字符串为?q=javascript&num=10 

let args = getQueryStringArgs(); 

alert(args["q"]); // "javascript" 
alert(args["num"]); // "10"
```

现在，查询字符串中的每个参数都是返回对象的一个属性，这样使用起来就方便了。

**URLSearchParams**

URLSearchParams 提供了一组标准 API 方法，通过它们可以检查和修改查询字符串。给URLSearchParams 构造函数传入一个查询字符串，就可以创建一个实例。这个实例上暴露了 get()、set()和 delete()等方法，可以对查询字符串执行相应操作。下面来看一个例子：

```javascript
let qs = '?q=javascript&num=10';

let searchParams = new URLSearchParams(qs);

alert(searchParams.toString()); // " q=javascript&num=10"
searchParams.has('num'); // true
searchParams.get('num'); // 10

searchParams.set('page', '3');
alert(searchParams.toString()); // " q=javascript&num=10&page=3"

searchParams.delete('q');
alert(searchParams.toString()); // " num=10&page=3"
```

大多数支持 URLSearchParams 的浏览器也支持将 URLSearchParams 的实例用作可迭代对象：

```javascript
let qs = "?q=javascript&num=10"; 

let searchParams = new URLSearchParams(qs); 

for (let param of searchParams) { 
 console.log(param); 
} 
// ["q", "javascript"] 
// ["num", "10"]
```

### 操作地址

可以通过修改 location 对象修改浏览器的地址。首先，最常见的是使用 assign()方法并传入一个 URL，如下所示：

```javascript
location.assign("http://www.wrox.com");
```

这行代码会立即启动导航到新 URL 的操作，同时在浏览器历史记录中增加一条记录。如果给location.href 或 window.location 设置一个 URL，也会以同一个 URL 值调用 assign()方法。比如，下面两行代码都会执行与显式调用 assign()一样的操作：

```javascript
window.location = "http://www.wrox.com"; 
location.href = "http://www.wrox.com";
```

在这 3 种修改浏览器地址的方法中，设置 location.href 是最常见的。

修改 location 对象的属性也会修改当前加载的页面。其中，hash、search、hostname、pathname和 port 属性被设置为新值之后都会修改当前 URL，如下面的例子所示：

```javascript
// 假设当前 URL 为 http://www.wrox.com/WileyCDA/ 

// 把 URL 修改为 http://www.wrox.com/WileyCDA/#section1 
location.hash = "#section1"; 

// 把 URL 修改为 http://www.wrox.com/WileyCDA/?q=javascript 
location.search = "?q=javascript"; 

// 把 URL 修改为 http://www.somewhere.com/WileyCDA/ 
location.hostname = "www.somewhere.com"; 

// 把 URL 修改为 http://www.somewhere.com/mydir/ 
location.pathname = "mydir"; 

// 把 URL 修改为 http://www.somewhere.com:8080/WileyCDA/ 
location.port = 8080;
```

除了 hash 之外，只要修改 location 的一个属性，就会导致页面重新加载新 URL。



在以前面提到的方式修改 URL 之后，浏览器历史记录中就会增加相应的记录。当用户单击“后退”按钮时，就会导航到前一个页面。如果不希望增加历史记录，可以使用 replace()方法。这个方法接收一个 URL 参数，但重新加载后不会增加历史记录。调用 replace()之后，用户不能回到前一页。比如下面的例子：

```javascript
<!DOCTYPE html> 
<html> 
<head> 
 <title>You won't be able to get back here</title> 
</head> 
<body> 
 <p>Enjoy this page for a second, because you won't be coming back here.</p> 
 <script> 
 setTimeout(() => location.replace("http://www.wrox.com/"), 1000); 
 </script> 
</body> 
</html>
```

浏览器加载这个页面 1 秒之后会重定向到 www.wrox.com。此时，“后退”按钮是禁用状态，即不能返回这个示例页面，除非手动输入完整的 URL。

最后一个修改地址的方法是 reload()，它能重新加载当前显示的页面。调用 reload()而不传参数，页面会以最有效的方式重新加载。也就是说，如果页面自上次请求以来没有修改过，浏览器可能会从缓存中加载页面。如果想强制从服务器重新加载，可以像下面这样给 reload()传个 true：

```javascript
location.reload(); // 重新加载，可能是从缓存加载
location.reload(true); // 重新加载，从服务器加载
```

脚本中位于 reload() 调用之后的代码可能执行也可能不执行，这取决于网络延迟和系统资源等因素。为此，最好把 reload() 作为最后一行代码。

## **navigator** 对象

navigator 是由 Netscape Navigator 2 最早引入浏览器的，现在已经成为客户端标识浏览器的标准。只要浏览器启用 JavaScript，navigator 对象就一定存在。但是与其他 BOM 对象一样，每个浏览器都支持自己的属性。

navigator 对象实现了 NavigatorID 、 NavigatorLanguage 、 NavigatorOnLine 、NavigatorContentUtils 、 NavigatorStorage 、 NavigatorStorageUtils 、 Navigator、ConcurrentHardware、NavigatorPlugins 和 NavigatorUserMedia 接口定义的属性和方法。

下表列出了这些接口定义的属性和方法：

| 属性/方法                     | 说 明                                                        |
| ----------------------------- | ------------------------------------------------------------ |
| activeVrDisplays              | 返回数组，包含 ispresenting 属性为 true 的 VRDisplay 实例    |
| appCodeName                   | 即使在非 Mozilla 浏览器中也会返回"Mozilla"                   |
| appName                       | 浏览器全名                                                   |
| appVersion                    | 浏览器版本。通常与实际的浏览器版本不一致                     |
| battery                       | 返回暴露 Battery Status API 的 BatteryManager 对象           |
| buildId                       | 浏览器的构建编号                                             |
| connection                    | 返回暴露 Network Information API 的 NetworkInformation 对象  |
| cookieEnabled                 | 返回布尔值，表示是否启用了 cookie                            |
| credentials                   | 返回暴露 Credentials Management API 的 CredentialsContainer 对象 |
| deviceMemory                  | 返回单位为 GB 的设备内存容量                                 |
| doNotTrack                    | 返回用户的“不跟踪”（do-not-track）设置                       |
| geolocation                   | 返回暴露 Geolocation API 的 Geolocation 对象                 |
| getVRDisplays()               | 返回数组，包含可用的每个 VRDisplay 实例                      |
| getUserMedia()                | 返回与可用媒体设备硬件关联的流                               |
| hardwareConcurrency           | 返回设备的处理器核心数量                                     |
| javaEnabled                   | 返回布尔值，表示浏览器是否启用了 Java                        |
| language                      | 返回浏览器的主语言                                           |
| languages                     | 返回浏览器偏好的语言数组                                     |
| locks                         | 返回暴露 Web Locks API 的 LockManager 对象                   |
| mediaCapabilities             | 返回暴露 Media Capabilities API 的 MediaCapabilities 对象    |
| mediaDevices                  | 返回可用的媒体设备                                           |
| maxTouchPoints                | 返回设备触摸屏支持的最大触点数                               |
| mimeTypes                     | 返回浏览器中注册的 MIME 类型数组                             |
| onLine                        | 返回布尔值，表示浏览器是否联网                               |
| oscpu                         | 返回浏览器运行设备的操作系统和（或）CPU                      |
| permissions                   | 返回暴露 Permissions API 的 Permissions 对象                 |
| platform                      | 返回浏览器运行的系统平台                                     |
| plugins                       | 返回浏览器安装的插件数组。在 IE 中，这个数组包含页面中所有<embed>元素 |
| product                       | 返回产品名称（通常是"Gecko"）                                |
| productSub                    | 返回产品的额外信息（通常是 Gecko 的版本）                    |
| registerProtocolHandler()     | 将一个网站注册为特定协议的处理程序                           |
| requestMediaKeySystemAccess() | 返回一个期约，解决为 MediaKeySystemAccess 对象               |
| sendBeacon()                  | 异步传输一些小数据                                           |
| serviceWorker                 | 返回用来与 ServiceWorker 实例交互的 ServiceWorkerContainer   |
| share()                       | 返回当前平台的原生共享机制                                   |
| storage                       | 返回暴露 Storage API 的 StorageManager 对象                  |
| userAgent                     | 返回浏览器的用户代理字符串                                   |
| vendor                        | 返回浏览器的厂商名称                                         |
| vendorSub                     | 返回浏览器厂商的更多信息                                     |
| vibrate()                     | 触发设备振动                                                 |
| webdriver                     | 返回浏览器当前是否被自动化程序控制                           |

navigator 对象的属性通常用于确定浏览器的类型。

### 检测插件

检测浏览器是否安装了某个插件是开发中常见的需求。除 IE10 及更低版本外的浏览器，都可以通过 plugins 数组来确定。这个数组中的每一项都包含如下属性。

+ name：插件名称。
+ description：插件介绍。
+ filename：插件的文件名。
+ length：由当前插件处理的 MIME 类型数量。

通常，name 属性包含识别插件所需的必要信息，尽管不是特别准确。检测插件就是遍历浏览器中可用的插件，并逐个比较插件的名称，如下所示：

```javascript
// 插件检测，IE10 及更低版本无效
let hasPlugin = function(name) {
    name = name.toLowerCase();
    for (let plugin of window.navigator.plugins) {
        if (plugin.name.toLowerCase().indexOf(name) > -1) {
            return true;
        }
    }
    return false;
};
// 检测 Flash
alert(hasPlugin('Flash'));
// 检测 QuickTime
alert(hasPlugin('QuickTime'));
```

这个 hasPlugin() 方法接收一个参数，即待检测插件的名称。第一步是把插件名称转换为小写形式，以便于比较。然后，遍历 plugins 数组，通过 indexOf() 方法检测每个 name 属性，看传入的名称是不是存在于某个数组中。比较的字符串全部小写，可以避免大小写问题。传入的参数应该尽可能独一无二，以避免混淆。像 "Flash"、"QuickTime" 这样的字符串就可以避免混淆。这个方法可以在 Firefox、Safari、Opera 和 Chrome 中检测插件。

IE11 的 window.navigator 对象开始支持 plugins 和 mimeTypes 属性。这意味着前面定义的函数可以适用于所有较新版本的浏览器。而且，IE11 中的 ActiveXObject 也从 DOM 中隐身了，意味着不能再用它来作为检测特性的手段。

**旧版本 IE 中的插件检测**

IE10 及更低版本中检测插件的问题比较多，因为这些浏览器不支持 Netscape 式的插件。在这些 IE 中检测插件要使用专有的 ActiveXObject，并尝试实例化特定的插件。IE 中的插件是实现为 COM 对象的，由唯一的字符串标识。因此，要检测某个插件就必须知道其 COM 标识符。例如，Flash 的标识符是 "ShockwaveFlash.ShockwaveFlash"。知道了这个信息后，就可以像这样检测 IE中是否安装了 Flash：

```javascript
// 在旧版本 IE 中检测插件
function hasIEPlugin(name) {
    try {
        new ActiveXObject(name);
        return true;
    } catch (ex) {
        return false;
    }
}

// 检测 Flash 
alert(hasIEPlugin('ShockwaveFlash.ShockwaveFlash'));
// 检测 QuickTime 
alert(hasIEPlugin('QuickTime.QuickTime'));
```

在这个例子中，hasIEPlugin()函数接收一个 DOM 标识符参数。为检测插件，这个函数会使用传入的标识符创建一个新 ActiveXObject 实例。相应代码封装在一个 try/catch 语句中，因此如果创建的插件不存在则会抛出错误。如果创建成功则返回 true，如果失败则在 catch 块中返回 false。上面的例子还演示了如何检测 Flash 和 QuickTime 插件。

因为检测插件涉及两种方式，所以一般要针对特定插件写一个函数，而不是使用通常的检测函数。比如下面的例子：

```javascript
// 在所有浏览器中检测 Flash 
function hasFlash() {
    var result = hasPlugin('Flash');
    if (!result) {
        result = hasIEPlugin('ShockwaveFlash.ShockwaveFlash');
    }
    return result;
}

// 在所有浏览器中检测 QuickTime 
function hasQuickTime() {
    var result = hasPlugin('QuickTime');
    if (!result) {
        result = hasIEPlugin('QuickTime.QuickTime');
    }
    return result;
}

// 检测 Flash 
alert(hasFlash());
// 检测 QuickTime 
alert(hasQuickTime());
```

以上代码定义了两个函数 hasFlash()和 hasQuickTime()。每个函数都先尝试使用非 IE 插件检测方式，如果返回 false（对 IE可能会），则再使用 IE插件检测方式。如果 IE插件检测方式再返回 false，整个检测方法也返回 false。只要有一种方式返回 true，检测方法就会返回 true。

### 注册处理程序

现代浏览器支持 navigator 上的（在 HTML5 中定义的）registerProtocolHandler()方法。这个方法可以把一个网站注册为处理某种特定类型信息应用程序。随着在线 RSS 阅读器和电子邮件客户端的流行，可以借助这个方法将 Web 应用程序注册为像桌面软件一样的默认应用程序。

要使用 registerProtocolHandler()方法，必须传入 3 个参数：要处理的协议（如"mailto"或"ftp"）、处理该协议的 URL，以及应用名称。比如，要把一个 Web 应用程序注册为默认邮件客户端，可以这样做：

```
navigator.registerProtocolHandler("mailto", 
 "http://www.somemailclient.com?cmd=%s", 
 "Some Mail Client");
```

这个例子为"mailto"协议注册了一个处理程序，这样邮件地址就可以通过指定的 Web 应用程序打开。注意，第二个参数是负责处理请求的 URL，%s 表示原始的请求。

## **screen** 对象

window 的另一个属性 screen 对象，是为数不多的几个在编程中很少用的 JavaScript 对象。这个对象中保存的纯粹是客户端能力信息，也就是浏览器窗口外面的客户端显示器的信息，比如像素宽度和像素高度。每个浏览器都会在 screen 对象上暴露不同的属性。下表总结了这些属性。

| 属 性       | 说 明                                        |
| ----------- | -------------------------------------------- |
| availHeight | 屏幕像素高度减去系统组件高度（只读）         |
| availLeft   | 没有被系统组件占用的屏幕的最左侧像素（只读） |
| availTop    | 没有被系统组件占用的屏幕的最顶端像素（只读） |
| availWidth  | 屏幕像素宽度减去系统组件宽度（只读）         |
| colorDepth  | 表示屏幕颜色的位数；多数系统是 32（只读）    |
| height      | 屏幕像素高度                                 |
| left        | 当前屏幕左边的像素距离                       |
| pixelDepth  | 屏幕的位深（只读）                           |
| top         | 当前屏幕顶端的像素距离                       |
| width       | 屏幕像素宽度                                 |
| orientation | 返回 Screen Orientation API 中屏幕的朝向     |

## **history** 对象

history 对象表示当前窗口首次使用以来用户的导航历史记录。因为 history 是 window 的属性，所以每个 window 都有自己的 history 对象。出于安全考虑，这个对象不会暴露用户访问过的 URL，但可以通过它在不知道实际 URL 的情况下前进和后退。

### 导航

go()方法可以在用户历史记录中沿任何方向导航，可以前进也可以后退。这个方法只接收一个参数，这个参数可以是一个整数，表示前进或后退多少步。负值表示在历史记录中后退（类似点击浏览器的“后退”按钮），而正值表示在历史记录中前进（类似点击浏览器的“前进”按钮）。下面来看几个例子：

```
// 后退一页
history.go(-1); 

// 前进一页
history.go(1); 

// 前进两页
history.go(2);
```

在旧版本的一些浏览器中，go()方法的参数也可以是一个字符串，这种情况下浏览器会导航到历史中包含该字符串的第一个位置。最接近的位置可能涉及后退，也可能涉及前进。如果历史记录中没有匹配的项，则这个方法什么也不做，如下所示：

```
// 导航到最近的 wrox.com 页面
history.go("wrox.com"); 

// 导航到最近的 nczonline.net 页面
history.go("nczonline.net");
```

go()有两个简写方法：back()和 forward()。顾名思义，这两个方法模拟了浏览器的后退按钮和前进按钮：

```
// 后退一页
history.back(); 

// 前进一页
history.forward();
```

history 对象还有一个 length 属性，表示历史记录中有多个条目。这个属性反映了历史记录的数量，包括可以前进和后退的页面。对于窗口或标签页中加载的第一个页面，history.length 等于 1。通过以下方法测试这个值，可以确定用户浏览器的起点是不是你的页面：

```
if (history.length == 1){ 
 // 这是用户窗口中的第一个页面
}
```

history 对象通常被用于创建“后退”和“前进”按钮，以及确定页面是不是用户历史记录中的第一条记录。

### 历史状态管理

现代 Web 应用程序开发中最难的环节之一就是历史记录管理。用户每次点击都会触发页面刷新的时代早已过去，“后退”和“前进”按钮对用户来说就代表“帮我切换一个状态”的历史也就随之结束了。为解决这个问题，首先出现的是 hashchange 事件。HTML5 也为history 对象增加了方便的状态管理特性。

hashchange 会在页面 URL 的散列变化时被触发，开发者可以在此时执行某些操作。而状态管理 API 则可以让开发者改变浏览器 URL 而不会加载新页面。为此，可以使用 history.pushState()方法。这个方法接收 3 个参数：一个 state 对象、一个新状态的标题和一个（可选的）相对 URL。例如：

```
let stateObject = {foo:"bar"}; 

history.pushState(stateObject, "My title", "baz.html");
```

pushState()方法执行后，状态信息就会被推到历史记录中，浏览器地址栏也会改变以反映新的相对 URL。除了这些变化之外，即使 location.href 返回的是地址栏中的内容，浏览器页不会向服务器发送请求。第二个参数并未被当前实现所使用，因此既可以传一个空字符串也可以传一个短标题。第一个参数应该包含正确初始化页面状态所必需的信息。为防止滥用，这个状态的对象大小是有限制的，通常在 500KB～1MB 以内。

因为 pushState()会创建新的历史记录，所以也会相应地启用“后退”按钮。此时单击“后退”按钮，就会触发 window 对象上的 popstate 事件。popstate 事件的事件对象有一个 state 属性，其中包含通过 pushState()第一个参数传入的 state 对象：

```
window.addEventListener("popstate", (event) => { 
 let state = event.state; 
 if (state) { // 第一个页面加载时状态是 null 
 processState(state); 
 } 
});
```

基于这个状态，应该把页面重置为状态对象所表示的状态（因为浏览器不会自动为你做这些）。记住，页面初次加载时没有状态。因此点击“后退”按钮直到返回最初页面时，event.state 会为 null。

可以通过 history.state 获取当前的状态对象，也可以使用 replaceState()并传入与pushState()同样的前两个参数来更新状态。更新状态不会创建新历史记录，只会覆盖当前状态：

```
history.replaceState({newFoo: "newBar"}, "New title");
```

传给 pushState()和 replaceState()的 state 对象应该只包含可以被序列化的信息。因此，DOM 元素之类并不适合放到状态对象里保存。

## 小结

浏览器对象模型（BOM，Browser Object Model）是以 window 对象为基础的，这个对象代表了浏览器窗口和页面可见的区域。window 对象也被复用为 ECMAScript 的 Global 对象，因此所有全局变量和函数都是它的属性，而且所有原生类型的构造函数和普通函数也都从一开始就存在于这个对象之上。本章讨论了 BOM 的以下内容。

+ 要引用其他 window 对象，可以使用几个不同的窗口指针。
+ 通过 location 对象可以以编程方式操纵浏览器的导航系统。通过设置这个对象上的属性，可以改变浏览器 URL 中的某一部分或全部。
+ 使用 replace()方法可以替换浏览器历史记录中当前显示的页面，并导航到新 URL。 
+ navigator 对象提供关于浏览器的信息。提供的信息类型取决于浏览器，不过有些属性如userAgent 是所有浏览器都支持的。

BOM 中的另外两个对象也提供了一些功能。screen 对象中保存着客户端显示器的信息。这些信息通常用于评估浏览网站的设备信息。history 对象提供了操纵浏览器历史记录的能力，开发者可以确定历史记录中包含多少个条目，并以编程方式实现在历史记录中导航，而且也可以修改历史记录。

# cookie

Cookie 是直接存储在浏览器中的一小串数据。它们是 HTTP 协议的一部分，由 [RFC 6265](https://tools.ietf.org/html/rfc6265) 规范定义。

Cookie 通常是由 Web 服务器使用响应 `Set-Cookie` HTTP-header 设置的。然后浏览器使用 `Cookie` HTTP-header 将它们自动添加到（几乎）每个对相同域的请求中。

最常见的用处之一就是身份验证：

1. 登录后，服务器在响应中使用 `Set-Cookie` HTTP-header 来设置具有唯一“会话标识符（session identifier）”的 cookie。
2. 下次当请求被发送到同一个域时，浏览器会使用 `Cookie` HTTP-header 通过网络发送 cookie。
3. 所以服务器知道是谁发起了请求。

我们还可以使用 `document.cookie` 属性从浏览器访问 cookie。

关于 cookie 及其选项，有很多棘手的事情。在本章中，我们将详细介绍它们。

## 从 document.cookie 中读取

```javascript
// 在 javascript.info，我们使用谷歌分析来进行统计，
// 所以应该存在一些 cookie
console.log(document.cookie); // cookie1=value1; cookie2=value2;...
```

`document.cookie` 的值由 `name=value` 对组成，以 `;` 分隔。每一个都是独立的 cookie。

为了找到一个特定的 cookie，我们可以以 `;` 作为分隔，将 `document.cookie` 分开，然后找到对应的名字。我们可以使用正则表达式或者数组函数来实现。

我们把这个留给读者当作练习。此外，在本章的最后，你可以找到一些操作 cookie 的辅助函数。

### 写入 document.cookie

我们可以写入 `document.cookie`。但这不是一个数据属性，它是一个 [访问器（getter/setter）](https://zh.javascript.info/property-accessors)。对其的赋值操作会被特殊处理。

**对 `document.cookie` 的写入操作只会更新其中提到的 cookie，而不会涉及其他 cookie。**

例如，此调用设置了一个名称为 `user` 且值为 `John` 的 cookie：

```javascript
document.cookie = "user=John"; // 只会更新名称为 user 的 cookie
console.log(document.cookie); // 展示所有 cookie
```

如果你运行了上面这段代码，你会看到多个 cookie。这是因为 `document.cookie=` 操作不是重写整所有 cookie。它只设置代码中提到的 cookie `user`。

从技术上讲，cookie 的名称和值可以是任何字符。为了保持有效的格式，它们应该使用内建的 `encodeURIComponent` 函数对其进行转义：

```javascript
// 特殊字符（空格），需要编码
let name = "my name";
let value = "John Smith"

// 将 cookie 编码为 my%20name=John%20Smith
document.cookie = encodeURIComponent(name) + '=' + encodeURIComponent(value);

console.log(document.cookie); // ...; my%20name=John%20Smith
```

**限制**

存在一些限制：

- `encodeURIComponent` 编码后的 `name=value` 对，大小不能超过 4KB。因此，我们不能在一个 cookie 中保存大的东西。
- 每个域的 cookie 总数不得超过 20+ 左右，具体限制取决于浏览器。



Cookie 有几个选项，其中很多都很重要，应该设置它。

选项被列在 `key=value` 之后，以 `;` 分隔，像这样：

```javascript
document.cookie = "user=John; path=/; expires=Tue, 19 Jan 2038 03:14:07 GMT"
```

### path

- **`path=/mypath`**

url 路径前缀必须是绝对路径。它使得该路径下的页面可以访问该 cookie。默认为当前路径。

如果一个 cookie 带有 `path=/admin` 设置，那么该 cookie 在 `/admin` 和 `/admin/something` 下都是可见的，但是在 `/home` 或 `/adminpage` 下不可见。

通常，我们应该将 `path` 设置为根目录：`path=/`，以使 cookie 对此网站的所有页面可见。

### domain

- **`domain=site.com`**

domain 控制了可访问 cookie 的域。但是在实际中，有一些限制。我们无法设置任何域。

**无法从另一个二级域访问 cookie，因此 `other.com` 永远不会收到在 `site.com` 设置的 cookie。**

这是一项安全限制，为了允许我们将敏感数据存储在应该仅在一个站点上可用的 cookie 中。

默认情况下，cookie 只有在设置的域下才能被访问到。

请注意，默认情况下，cookie 也不会共享给子域，例如 `forum.site.com`。

```javascript
// 如果我们在 site.com 网站上设置了 cookie……
document.cookie = "user=John"

// ……在 forum.site.com 域下我们无法访问它
alert(document.cookie); // 没有 user
```

……但这是可以设置的。如果我们想允许像 `forum.site.com` 这样的子域在 `site.com` 上设置 cookie，也是可以实现的。

为此，当在 `site.com` 设置 cookie 时，我们应该明确地将 `domain` 选项设置为根域：`domain=site.com`。那么，所有子域都可以访问到这样的 cookie。

例如：

```javascript
// 在 site.com
// 使 cookie 可以被在任何子域 *.site.com 访问：
document.cookie = "user=John; domain=site.com"

// 之后

// 在 forum.site.com
alert(document.cookie); // 有 cookie user=John
```

出于历史原因，`domain=.site.com`（`site.com` 前面有一个点符号）也以相同的方式工作，允许从子域访问 cookie。这是一个旧的表示方式，如果我们需要支持非常旧的浏览器，那么应该使用它。

总结一下，通过 `domain` 选项的设置，可以实现允许在子域访问 cookie。

### expires，max-age

默认情况下，如果一个 cookie 没有设置这两个参数中的任何一个，那么在关闭浏览器之后，它就会消失。此类 cookie 被称为 "session cookie”。

为了让 cookie 在浏览器关闭后仍然存在，我们可以设置 `expires` 或 `max-age` 选项中的一个。

- **`expires=Tue, 19 Jan 2038 03:14:07 GMT`**

cookie 的过期时间定义了浏览器会自动清除该 cookie 的时间。

日期必须完全采用 GMT 时区的这种格式。我们可以使用 `date.toUTCString` 来获取它。例如，我们可以将 cookie 设置为 1 天后过期。

```javascript
// 当前时间 +1 天
let date = new Date(Date.now() + 86400e3);
date = date.toUTCString();
document.cookie = "user=John; expires=" + date;
```

如果我们将 `expires` 设置为过去的时间，则 cookie 会被删除。

- **`max-age=3600`**

它是 `expires` 的替代选项，指明了 cookie 的过期时间距离当前时间的秒数。

如果将其设置为 0 或负数，则 cookie 会被删除：

```javascript
// cookie 会在一小时后失效
document.cookie = "user=John; max-age=3600";

// 删除 cookie（让它立即过期）
document.cookie = "user=John; max-age=0";
```

### secure

- **`secure`**

Cookie 应只能被通过 HTTPS 传输。

**默认情况下，如果我们在 `http://site.com` 上设置了 cookie，那么该 cookie 也会出现在 `https://site.com` 上，反之亦然。**

也就是说，cookie 是基于域的，它们不区分协议。

使用此选项，如果一个 cookie 是通过 `https://site.com` 设置的，那么它不会在相同域的 HTTP 环境下出现，例如 `http://site.com`。所以，如果一个 cookie 包含绝不应该通过未加密的 HTTP 协议发送的敏感内容，那么就应该设置 `secure` 标识。

```javascript
// 假设我们现在在 HTTPS 环境下
// 设置 cookie secure（只在 HTTPS 环境下可访问）
document.cookie = "user=John; secure";
```

### samesite

这是另外一个关于安全的特性。它旨在防止 XSRF（跨网站请求伪造）攻击。

为了了解它是如何工作的，以及何时有用，让我们看一下 XSRF 攻击。

#### XSRF 攻击

想象一下，你登录了 `bank.com` 网站。此时：你有了来自该网站的身份验证 cookie。你的浏览器会在每次请求时将其发送到 `bank.com`，以便识别你，并执行所有敏感的财务上的操作。

现在，在另外一个窗口中浏览网页时，你不小心访问了另一个网站 `evil.com`。该网站具有向 `bank.com` 网站提交一个具有启动与黑客账户交易的字段的表单 `<form action="https://bank.com/pay">` 的 JavaScript 代码。

你每次访问 `bank.com` 时，浏览器都会发送 cookie，即使该表单是从 `evil.com` 提交过来的。因此，银行会识别你的身份，并执行真实的付款。

这就是所谓的“跨网站请求伪造（Cross-Site Request Forgery，简称 XSRF）”攻击。

当然，实际的银行会防止出现这种情况。所有由 `bank.com` 生成的表单都具有一个特殊的字段，即所谓的 “XSRF 保护 token”，恶意页面既不能生成，也不能从远程页面提取它。它可以在那里提交表单，但是无法获取数据。并且，网站 `bank.com` 会对收到的每个表单都进行这种 token 的检查。

但是，实现这种防护需要花费时间。我们需要确保每个表单都具有所需的 token 字段，并且我们还必须检查所有请求。

#### 输入 cookie samesite 选项

Cookie 的 `samesite` 选项提供了另一种防止此类攻击的方式，（理论上）不需要要求 “XSRF 保护 token”。

它有两个可能的值：

- `samesite=strict`（和没有值的 `samesite` 一样)

如果用户来自同一网站之外，那么设置了 `samesite=strict` 的 cookie 永远不会被发送。

换句话说，无论用户是通过邮件链接还是从 `evil.com` 提交表单，或者进行了任何来自其他域下的操作，cookie 都不会被发送。

如果身份验证 cookie 具有 `samesite` 选项，那么 XSRF 攻击是没有机会成功的，因为来自 `evil.com` 的提交没有 cookie。因此，`bank.com` 将无法识别用户，也就不会继续进行付款。

这种保护是相当可靠的。只有来自 `bank.com` 的操作才会发送 `samesite` cookie，例如来自 `bank.com` 的另一页面的表单提交。

虽然，这样有一些不方便。

当用户通过合法的链接访问 `bank.com` 时，例如从他们自己的笔记，他们会感到惊讶，`bank.com` 无法识别他们的身份。实际上，在这种情况下不会发送 `samesite=strict` cookie。

我们可以通过使用两个 cookie 来解决这个问题：一个 cookie 用于“一般识别”，仅用于说 “Hello, John”，另一个带有 `samesite=strict` 的 cookie 用于进行数据更改的操作。这样，从网站外部来的用户会看到欢迎信息，但是支付操作必须是从银行网站启动的，这样第二个 cookie 才能被发送。

- **`samesite=lax`**

一种更轻松的方法，该方法还可以防止 XSRF 攻击，并且不会破坏用户体验。

宽松（lax）模式，和 `strict` 模式类似，当从外部来到网站，则禁止浏览器发送 cookie，但是增加了一个例外。

如果以下两个条件均成立，则会发送含 `samesite=lax` 的 cookie：

1. HTTP 方法是“安全的”（例如 GET 方法，而不是 POST）。

   所有安全的 HTTP 方法详见 [RFC7231 规范](https://tools.ietf.org/html/rfc7231)。基本上，这些都是用于读取而不是写入数据的方法。它们不得执行任何更改数据的操作。跟随链接始终是 GET，是安全的方法。

2. 该操作执行顶级导航（更改浏览器地址栏中的 URL）。

   这通常是成立的，但是如果导航是在一个 `<iframe>` 中执行的，那么它就不是顶级的。此外，用于网络请求的 JavaScript 方法不会执行任何导航，因此它们不适合。

所以，`samesite=lax` 所做的是基本上允许最常见的“前往 URL”操作携带 cookie。例如，从笔记中打开网站链接就满足这些条件。

但是，任何更复杂的事儿，例如来自另一个网站的网络请求或表单提交都会丢失 cookie。

如果这种情况适合你，那么添加 `samesite=lax` 将不会破坏用户体验并且可以增加保护。

总体而言，`samesite` 是一个很好的选项。

但它有个缺点：

- `samesite` 会被到 2017 年左右的旧版本浏览器忽略（不兼容）。

**因此，如果我们仅依靠 `samesite` 提供保护，那么在旧版本的浏览器上将很容易受到攻击。**

但是，我们肯定可以将 `samesite` 与其他保护措施（例如 XSRF token）一起使用，例如 xsrf token，这样可以多增加一层保护，将来，当旧版本的浏览器淘汰时，我们可能就可以删除 xsrf token 这种方式了。

### httpOnly

这个选项和 JavaScript 没有关系，但是我们必须为了完整性也提一下它。

Web 服务器使用 `Set-Cookie` header 来设置 cookie。并且，它可以设置 `httpOnly` 选项。

这个选项禁止任何 JavaScript 访问 cookie。我们使用 `document.cookie` 看不到此类 cookie，也无法对此类 cookie 进行操作。

这是一种预防措施，当黑客将自己的 JavaScript 代码注入网页，并等待用户访问该页面时发起攻击，而这个选项可以防止此时的这种攻击。这应该是不可能发生的，黑客应该无法将他们的代码注入我们的网站，但是网站有可能存在 bug，使得黑客能够实现这样的操作。

通常来说，如果发生了这种情况，并且用户访问了带有黑客 JavaScript 代码的页面，黑客代码将执行并通过 `document.cookie` 获取到包含用户身份验证信息的 cookie。这就很糟糕了。

但是，如果 cookie 设置了 `httpOnly`，那么 `document.cookie` 则看不到 cookie，所以它受到了保护。

### 附录: Cookie 函数

这里有一组有关 cookie 操作的函数，比手动修改 `document.cookie` 方便得多。

有很多这种 cookie 库，所以这些函数只用于演示。虽然它们都能正常使用。

#### getCookie(name)

获取 cookie 最简短的方式是使用 [正则表达式](https://zh.javascript.info/regular-expressions)。

`getCookie(name)` 函数返回具有给定 `name` 的 cookie：

```javascript
// 返回具有给定 name 的 cookie，
// 如果没找到，则返回 undefined
function getCookie(name) {
  let matches = document.cookie.match(new RegExp(
    "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
  ));
  return matches ? decodeURIComponent(matches[1]) : undefined;
}
```

这里的 `new RegExp` 是动态生成的，以匹配 `; name=<value>`。

请注意 cookie 的值是经过编码的，所以 `getCookie` 使用了内建方法 `decodeURIComponent` 函数对其进行解码。

#### setCookie(name, value, options)

将 cookie 的 `name` 设置为具有默认值 `path=/`（可以修改以添加其他默认值）和给定值 `value`：

```javascript
function setCookie(name, value, options = {}) {

  options = {
    path: '/',
    // 如果需要，可以在这里添加其他默认值
    ...options
  };

  if (options.expires instanceof Date) {
    options.expires = options.expires.toUTCString();
  }

  let updatedCookie = encodeURIComponent(name) + "=" + encodeURIComponent(value);

  for (let optionKey in options) {
    updatedCookie += "; " + optionKey;
    let optionValue = options[optionKey];
    if (optionValue !== true) {
      updatedCookie += "=" + optionValue;
    }
  }

  document.cookie = updatedCookie;
}

// 使用范例：
setCookie('user', 'John', {secure: true, 'max-age': 3600});
```

#### deleteCookie(name)

要删除一个 cookie，我们可以给它设置一个负的过期时间来调用它：

```javascript
function deleteCookie(name) {
  setCookie(name, "", {
    'max-age': -1
  })
}
```

**更新或删除必须使用相同的路径和域**

请注意：当我们更新或删除一个 cookie 时，我们应该使用和设置 cookie 时相同的路径和域选项。

代码放在：[cookie.js](https://zh.javascript.info/article/cookie/cookie.js)。

# Storage

Web 存储对象 `localStorage` 和 `sessionStorage` 允许我们在浏览器上保存键/值对。两个存储对象都提供相同的方法和属性：

```javascript
// 存储键值对
setItem(key,value)

// 依据键获取值
getItem(key)

// 删除键及其对应的值
removeItem(key)

// 删除所有的数据
clear()

// 获取该索引下的键名
key(index)

// 获取存储的内容长度
length
```

- sessionStorage：页面刷新后数据仍然保留在浏览器中
- localStorage：浏览器完全重启后数据仍然保留在浏览器中

我们已经有了 cookie。为什么还要其他存储对象呢？

- 与 cookie 不同，Web 存储对象不会随每个请求被发送到服务器。因此，我们可以保存更多数据。大多数浏览器都允许保存至少 2MB 的数据（或更多），并且具有用于配置数据的设置。
- 还有一点和 cookie 不同，服务器无法通过 HTTP header 操纵存储对象。一切都是在 JavaScript 中完成的。
- 存储绑定到源（域/协议/端口三者）。也就是说，不同协议或子域对应不同的存储对象，它们之间无法访问彼此数据。

## localStorage 

- 在同源的所有标签页和窗口之间共享数据。
- 数据不会过期。它在浏览器重启甚至系统重启后仍然存在。

我们只需要在同一个源（域/端口/协议），URL 路径可以不同。在所有同源的窗口之间，`localStorage` 数据可以共享。因此，如果我们在一个窗口中设置了数据，则在另一个窗口中也可以看到数据变化。

### 类对象形式调用

我们还可以像使用一个普通对象那样，读取/设置键，像这样：

```javascript
// 设置 key
localStorage.test = 2;

// 获取 key
console.log(localStorage.test); // 2

// 删除 key
delete localStorage.test;
```

这是历史原因造成的，并且大多数情况下都可行，但通常不建议这样做，因为：

1. 如果键是由用户生成的，那么它可以是任何内容，例如 `length` 或 `toString`，也可以是 `localStorage` 的另一种内建方法。在这种情况下，`getItem/setItem` 可以正常工作，而类对象访问的方式则会失败：

   ```javascript
   let key = 'length';
   localStorage[key] = 5; // Error，无法对 length 进行赋值
   ```

2. 有一个 `storage` 事件，在我们更改数据时会触发。但以类对象方式访问时，不会触发该事件。

### 遍历键

正如我们所看到的，这些方法提供了“按照键获取/设置/删除”的功能。但我们如何获取所有保存的值或键呢？不幸的是，存储对象是不可迭代的。一种方法是像遍历数组那样遍历它们：

```javascript
for (let i = 0; i < localStorage.length; i++) {
    let key = localStorage.key(i);
  	console.log(`${key}: ${localStorage.getItem(key)}`);
}
```

另一个方式是使用 `for key in localStorage` 循环，就像处理常规对象一样。

它会遍历所有的键，但也会输出一些我们不需要的内建字段。

```javascript
// 不好的尝试
for (let key in localStorage) {
    console.log(key); // 显示 getItem，setItem 和其他内建的东西
}
```

因此，我们需要使用 `hasOwnProperty` 检查来过滤掉原型中的字段：

```javascript
for (let key in localStorage) {
    if (!localStorage.hasOwnProperty(key)) {
        continue; // 跳过像 "setItem"，"getItem" 等这样的键
  }
    console.log(`${key}: ${localStorage.getItem(key)}`);
}
```

使用 `Object.keys` 获取只属于“自己”的键，然后如果需要，可以遍历它们：

```javascript
for (let key of keys) {
    console.log(`${key}: ${localStorage.getItem(key)}`);
}
```

后者有效，因为 `Object.keys` 只返回属于对象的键，会忽略原型上的。

### 仅字符串

请注意，键和值都必须是字符串。如果是任何其他类型，例数字或对象，它会被自动转换为字符串。

```javascript
  localStorage.user = {name: 'John'};
  console.log(localStorage.user); // [object Object]
```

我们可以使用 `JSON` 来存储对象：

```javascript
  // sometime later
  let user = JSON.parse(localStorage.user);
  console.log(user.name); // John
```

也可以对整个存储对象进行字符串化处理，例如出于调试目的：

```javascript
  // 为 JSON.stringify 增加了格式设置选项，以使对象看起来更美观
  console.log(JSON.stringify(localStorage, null, 2));
```

## sessionStorage

`sessionStorage` 对象的使用频率比 `localStorage` 对象低得多。属性和方法是相同的，但是它有更多的限制：

sessionStorage的数据只存在于当前浏览器标签页。

- 具有相同页面的另一个标签页中将会有不同的存储。
- 但是，它在同一标签页下的 iframe 之间是共享的（假如它们来自相同的源）。

- 数据在页面刷新后仍然保留，但在关闭/重新打开浏览器标签页后不会被保留。

但是，如果你在另一个新的标签页中打开此页面，然后在新页面中再次运行上面这行代码，则会得到 `null`，表示“未找到数据”。

这是因为 `sessionStorage` 不仅绑定到源，还绑定在同一浏览器标签页。因此，`sessionStorage` 很少被使用。

## Storage 事件

当 `localStorage` 或 `sessionStorage` 中的数据更新后，[storage](https://html.spec.whatwg.org/multipage/webstorage.html#the-storageevent-interface) 事件就会触发，它具有以下属性：

- `key` —— 发生更改的数据的 `key`（如果调用的是 `.clear()` 方法，则为 `null`）。
- `oldValue` —— 旧值（如果是新增数据，则为 `null`）。
- `newValue` —— 新值（如果是删除数据，则为 `null`）。
- `url` —— 发生数据更新的文档的 url。
- `storageArea` —— 发生数据更新的 `localStorage` 或 `sessionStorage` 对象。

重要的是：该事件会在所有可访问到存储对象的 `window` 对象上触发，导致当前数据改变的 `window` 对象除外。

我们来详细解释一下。

想象一下，你有两个窗口，它们具有相同的页面。所以 `localStorage` 在它们之间是共享的。

你可以想在浏览器的两个窗口中打开此页面来测试下面的代码。

如果两个窗口都在监听 `window.onstorage` 事件，那么每个窗口都会对另一个窗口中发生的更新作出反应。

```javascript
// 在其他文档对同一存储进行更新时触发
window.onstorage = event => { // 也可以使用 window.addEventListener('storage', event => {
  if (event.key != 'now') return;
  alert(event.key + ':' + event.newValue + " at " + event.url);
};

localStorage.setItem('now', Date.now());
```

请注意，该事件还包含：`event.url` —— 发生数据更新的文档的 url。

并且，`event.storageArea` 包含存储对象 —— `sessionStorage` 和 `localStorage` 具有相同的事件，所以 `event.storageArea` 引用了被修改的对象。我们可能会想设置一些东西，以“响应”更改。

**这允许同源的不同窗口交换消息。**

现代浏览器还支持 [Broadcast channel API](https://developer.mozilla.org/zh/docs/Web/api/Broadcast_Channel_API)，这是用于同源窗口之间通信的特殊 API，它的功能更全，但被支持的情况不好。有一些库基于 `localStorage` 来 polyfill 该 API，使其可以用在任何地方。

## 总结

Web 存储对象 `localStorage` 和 `sessionStorage` 允许我们在浏览器中保存键/值对。

- `key` 和 `value` 都必须为字符串。
- 存储大小限制为 5MB+，具体取决于浏览器。
- 它们不会过期。
- 数据绑定到源（域/端口/协议）。

| `localStorage`                       | `sessionStorage`                                       |
| :----------------------------------- | :----------------------------------------------------- |
| 在同源的所有标签页和窗口之间共享数据 | 在当前浏览器标签页中可见，包括同源的 iframe            |
| 浏览器重启后数据仍然保留             | 页面刷新后数据仍然保留（但标签页关闭后数据则不再保留） |

API：

- `setItem(key, value)` —— 存储键/值对。
- `getItem(key)` —— 按照键获取值。
- `removeItem(key)` —— 删除键及其对应的值。
- `clear()` —— 删除所有数据。
- `key(index)` —— 获取该索引下的键名。
- `length` —— 存储的内容的长度。
- 使用 `Object.keys` 来获取所有的键。
- 我们将键作为对象属性来访问，在这种情况下，不会触发 `storage` 事件。

Storage 事件：

- 在调用 `setItem`，`removeItem`，`clear` 方法后触发。
- 包含有关操作的所有数据（`key/oldValue/newValue`），文档 `url` 和存储对象 `storageArea`。
- 在所有可访问到存储对象的 `window` 对象上触发，导致当前数据改变的 `window` 对象除外（对于 `sessionStorage` 是在当前标签页下，对于 `localStorage` 是在全局，即所有同源的窗口）。

# fetch

## 分派请求

fetch()只有一个必需的参数 input。多数情况下，这个参数是要获取资源的 URL。这个方法返回一个期约：

```javascript
let r = fetch('/bar'); 

console.log(r); // Promise <pending> 
```

URL 的格式（相对路径、绝对路径等）的解释与 XHR 对象一样。请求完成、资源可用时，期约会解决为一个 Response 对象。这个对象是 API 的封装，可以通过它取得相应资源。获取资源要使用这个对象的属性和方法，掌握响应的情况并将负载转换为有用的形式，如下所示：

```javascript
  fetch('bar.txt').then((response) => {
    console.log(response);
  });

// Response { type: "basic", url: ... }
```

## 读取响应

读取响应内容的最简单方式是取得纯文本格式的内容，这要用到 text()方法。这个方法返回一个期约，会解决为取得资源的完整内容：

```javascript
  fetch('bar.txt').then((response) => {
    response.text().then((data) => {
      console.log(data);
    });
  });

// bar.txt 的内容
```

内容的结构通常是打平的：

```javascript
  fetch('bar.txt')
  .then((response) => response.text())
  .then((data) => console.log(data));

// bar.txt 的内容
```

## 处理状态码和请求失败

Fetch API 支持通过 Response 的 status（状态码）和 statusText（状态文本）属性检查响应状态。成功获取响应的请求通常会产生值为 200 的状态码，如下所示：

```javascript
  // 处理状态码和请求失败
  fetch('bar.txt').then((response) => {
    console.log(response.status); // 200
    console.log(response.statusText); // OK
  });
```

请求不存在的资源通常会产生值为 404 的状态码：

```javascript
  fetch('/does-not-exist').then((response) => {
    console.log(response.status); // 404 
    console.log(response.statusText); // Not Found 
  });
```

请求的 URL 如果抛出服务器错误会产生值为 500 的状态码：

```javascript
  fetch('/throw-server-error').then((response) => {
    console.log(response.status); // 500 
    console.log(response.statusText); // Internal Server Error 
  });
```

可以显式地设置 fetch()在遇到重定向时的行为（本章后面会介绍），不过默认行为是跟随重定向并返回状态码不是 300~399 的响应。跟随重定向时，响应对象的 redirected 属性会被设置为 true，而状态码仍然是 200：

```javascript
  fetch('/permanent-redirect').then((response) => {
    // 默认行为是跟随重定向直到最终 URL
    // 这个例子会出现至少两轮网络请求
    // <origin url>/permanent-redirect -> <redirect url>
    console.log(response.status); // 200
    console.log(response.statusText); // OK
    console.log(response.redirected); // true
  });
```

在前面这几个例子中，虽然请求可能失败（如状态码为 500），但都只执行了期约的解决处理函数。事实上，只要服务器返回了响应，fetch()期约都会解决。这个行为是合理的：系统级网络协议已经成功完成消息的一次往返传输。至于真正的“成功”请求，则需要在处理响应时再定义。

通常状态码为 200 时就会被认为成功了，其他情况可以被认为未成功。为区分这两种情况，可以在状态码非 200~299 时检查 Response 对象的 ok 属性：

```javascript
fetch('/bar') 
 .then((response) => { 
 console.log(response.status); // 200 
 console.log(response.ok); // true 
 }); 

fetch('/does-not-exist') 
 .then((response) => { 
 console.log(response.status); // 404 
 console.log(response.ok); // false 
 });
```

因为服务器没有响应而导致浏览器超时，这样真正的 fetch()失败会导致期约被拒绝：

```javascript
  fetch('/hangs-forever').then((response) => {
    console.log(response);
  }, (err) => {
    console.log(err);
  });

//（浏览器超时后）
// TypeError: "NetworkError when attempting to fetch resource."
```

违反 CORS、无网络连接、HTTPS 错配及其他浏览器/网络策略问题都会导致期约被拒绝。可以通过 url 属性检查通过 fetch()发送请求时使用的完整 URL：

```javascript
  // foo.com/bar/baz 发送的请求
  console.log(window.location.href); // https://foo.com/bar/baz 

  fetch('qux').then((response) => console.log(response.url));
  // https://foo.com/bar/qux

  fetch('/qux').then((response) => console.log(response.url));
  // https://foo.com/qux

  fetch('//qux.com').then((response) => console.log(response.url));
  // https://qux.com

  fetch('https://qux.com').then((response) => console.log(response.url));
  // https://qux.com
```

## 自定义选项

只使用 URL 时，fetch()会发送 GET 请求，只包含最低限度的请求头。要进一步配置如何发送请求，需要传入可选的第二个参数 init 对象。init 对象要按照下表中的键/值进行填充。

| 键             | 值                                                           |
| -------------- | ------------------------------------------------------------ |
| body           | 指定使用请求体时请求体的内容<br />必须是 Blob、BufferSource、FormData、URLSearchParams、ReadableStream 或 String 的实例 |
| cache          | 用于控制浏览器与 HTTP缓存的交互。要跟踪缓存的重定向，请求的 redirect 属性值必须是"follow"，而且必须符合同源策略限制。必须是下列值之一<br />Default<br />+ fetch()返回命中的有效缓存。不发送请求<br />+ 命中无效（stale）缓存会发送条件式请求。如果响应已经改变，则更新缓存的值。然后 fetch()返回缓存的值<br />+ 未命中缓存会发送请求，并缓存响应。然后 fetch()返回响应 <br />no-store<br />+ 浏览器不检查缓存，直接发送请求<br />+ 不缓存响应，直接通过 fetch()返回<br />reload<br />+ 浏览器不检查缓存，直接发送请求<br />+ 缓存响应，再通过 fetch()返回<br />no-cache<br />+ 无论命中有效缓存还是无效缓存都会发送条件式请求。如果响应已经改变，则更新缓存的值。然后 fetch()返回缓存的值<br />+ 未命中缓存会发送请求，并缓存响应。然后 fetch()返回响应<br />force-cache<br />+ 无论命中有效缓存还是无效缓存都通过 fetch()返回。不发送请求<br />+ 未命中缓存会发送请求，并缓存响应。然后 fetch()返回响应<br />only-if-cached<br />+ 只在请求模式为 same-origin 时使用缓存<br />+ 无论命中有效缓存还是无效缓存都通过 fetch()返回。不发送请求<br />+ 未命中缓存返回状态码为 504（网关超时）的响应<br />默认为 default |
| credentials    | 用于指定在外发请求中如何包含 cookie。与 XMLHttpRequest 的 withCredentials 标签类似<br />必须是下列字符串值之一<br />+ omit：不发送 cookie<br />+ same-origin：只在请求 URL 与发送 fetch()请求的页面同源时发送 cookie <br />+ include：无论同源还是跨源都包含 cookie <br />在支持 Credential Management API 的浏览器中，也可以是一个 FederatedCredential 或PasswordCredential 的实例<br />默认为 same-origin |
| headers        | 用于指定请求头部<br />必须是 Headers 对象实例或包含字符串格式键/值对的常规对象<br />默认值为不包含键/值对的 Headers 对象。这不意味着请求不包含任何头部，浏览器仍然会随请求发送一些头部。虽然这些头部对 JavaScript 不可见，但浏览器的网络检查器可以观察到 |
| integrity      | 用于强制子资源完整性<br />必须是包含子资源完整性标识符的字符串<br />默认为空字符串 |
| keepalive      | 用于指示浏览器允许请求存在时间超出页面生命周期。适合报告事件或分析，比如页面在 fetch() 请求后很快卸载。设置 keepalive 标志的 fetch()请求可用于替代 Navigator.sendBeacon() <br />必须是布尔值 默认为 false |
| method         | 用于指定 HTTP 请求方法，基本上就是如下字符串值：GET、POST、PUT、PATCH、DELETE、HEAD、OPTIONS、CONNECT、TARCE。默认为 GET |
| mode           | 用于指定请求模式。这个模式决定来自跨源请求的响应是否有效，以及客户端可以读取多少响应。违反这里指定模式的请求会抛出错误，必须是下列字符串值之一<br />+ cors：允许遵守 CORS 协议的跨源请求。响应是“CORS 过滤的响应”，意思是响应中可以访问的浏览器头部是经过浏览器强制白名单过滤的<br />+ no-cors：允许不需要发送预检请求的跨源请求（HEAD、GET 和只带有满足 CORS 请求头部的POST）。响应类型是 opaque，意思是不能读取响应内容<br />+ same-origin：任何跨源请求都不允许发送<br />+ navigate：用于支持 HTML 导航，只在文档间导航时使用。基本用不到<br />在通过构造函数手动创建 Request 实例时，默认为 cors；否则，默认为 no-cors |
| redirect       | 用于指定如何处理重定向响应（状态码为 301、302、303、307 或 308）<br />必须是下列字符串值之一<br />+ follow：跟踪重定向请求，以最终非重定向 URL 的响应作为最终响应<br />+ error：重定向请求会抛出错误<br />+ manual：不跟踪重定向请求，而是返回 opaqueredirect 类型的响应，同时仍然暴露期望的重定向 URL。允许以手动方式跟踪重定向<br />默认为 follow |
| referrer       | 用于指定 HTTP 的 Referer 头部的内容，必须是下列字符串值之一<br />+ no-referrer：以 no-referrer 作为值<br />+ client/about:client：以当前 URL 或 no-referrer（取决于来源策略 referrerPolicy）作为值<br />+ <URL>：以伪造 URL 作为值。伪造 URL 的源必须与执行脚本的源匹配默认为 client/about:client |
| referrerPolicy | 用于指定 HTTP 的 Referer 头部，必须是下列字符串值之一<br />no-referrer<br />+ 请求中不包含 Referer 头部<br />no-referrer-when-downgrade<br />+ 对于从安全 HTTPS 上下文发送到 HTTP URL 的请求，不包含 Referer 头部<br />+ 对于所有其他请求，将 Referer 设置为完整 URL <br />origin<br />+ 对于所有请求，将 Referer 设置为只包含源<br />same-origin<br />+ 对于跨源请求，不包含 Referer 头部<br />+ 对于同源请求，将 Referer 设置为完整 URL <br />strict-origin<br />+ 对于从安全 HTTPS 上下文发送到 HTTP URL 的请求，不包含 Referer 头部<br />+ 对于所有其他请求，将 Referer 设置为只包含源<br />origin-when-cross-origin<br />+ 对于跨源请求，将 Referer 设置为只包含源<br />+ 对于同源请求，将 Referer 设置为完整 URL<br />strict-origin-when-cross-origin<br />+ 对于从安全 HTTPS 上下文发送到 HTTP URL 的请求，不包含 Referer 头部<br />+ 对于所有其他跨源请求，将 Referer 设置为只包含源<br />+ 对于同源请求，将 Referer 设置为完整 URL <br />unsafe-url<br />+ 对于所有请求，将 Referer 设置为完整 URL <br />默认为 no-referrer-when-downgrade |
| signal         | 用于支持通过 AbortController 中断进行中的 fetch()请求<br />必须是 AbortSignal 的实例<br />默认为未关联控制器的 AbortSignal 实例 |

## 发送 JSON 数据

可以像下面这样发送简单 JSON 字符串：

```javascript
  let payload = JSON.stringify({
    foo: 'bar',
  });

  let jsonHeaders = new Headers({
    'Content-Type': 'application/json',
  });

  fetch('/send-me-json', {
    method: 'POST', // 发送请求体时必须使用一种 HTTP 方法
    body: payload,
    headers: jsonHeaders,
  });

```

## 在请求体中发送参数

因为请求体支持任意字符串值，所以可以通过它发送请求参数：

```javascript
  let payload = 'foo=bar&baz=qux';
  let paramHeaders = new Headers({
    'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
  });

  fetch('/send-me-params', {
    method: 'POST', // 发送请求体时必须使用一种 HTTP 方法
    body: payload,
    headers: paramHeaders,
  });
```

## 发送文件

因为请求体支持 FormData 实现，所以 fetch()也可以序列化并发送文件字段中的文件：

```html
<body>
<input type="file">
<input type="button" value="上传" id="upload">
<script>
  let upload = document.querySelector('#upload');
  upload.onclick = function() {
    let imageFormData = new FormData();
    let imageInput = document.querySelector('input[type=\'file\']');

    imageFormData.append('image', imageInput.files[0]);

    fetch('/img-upload', {
      method: 'POST',
      body: imageFormData,
    });
  };
</script>
</body>
```

```javascript
    // 这个 fetch()实现可以支持多个文件：
	let imageFormData = new FormData();
    let imageInput = document.querySelector('input[type=\'file\'][multiple]');

    for (let i = 0; i < imageInput.files.length; ++i) {
      imageFormData.append('image', imageInput.files[i]);
    }

    fetch('/img-upload', {
      method: 'POST',
      body: imageFormData,
    });
```



## 加载 **Blob** 文件

Fetch API也能提供 Blob 类型的响应，而 Blob 又可以兼容多种浏览器 API。一种常见的做法是明确将图片文件加载到内存，然后将其添加到 HTML图片元素。为此，可以使用响应对象上暴露的 blob()方法。这个方法返回一个期约，解决为一个 Blob 的实例。然后，可以将这个实例传给 URL.createObjectUrl() 以生成可以添加给图片元素 src 属性的值：

```javascript
const imageElement = document.querySelector('img');

fetch('my-image.png') 
 .then((response) => response.blob()) 
 .then((blob) => { 
 imageElement.src = URL.createObjectURL(blob); 
 });
```

## 发送跨源请求

从不同的源请求资源，响应要包含 CORS 头部才能保证浏览器收到响应。没有这些头部，跨源请求会失败并抛出错误。

```javascript
fetch('//cross-origin.com'); 
// TypeError: Failed to fetch 
// No 'Access-Control-Allow-Origin' header is present on the requested resource.
```

如果代码不需要访问响应，也可以发送 no-cors 请求。此时响应的 type 属性值为 opaque，因此无法读取响应内容。这种方式适合发送探测请求或者将响应缓存起来供以后使用。

```javascript
fetch('//cross-origin.com', { method: 'no-cors' }) 
 .then((response) => console.log(response.type)); 
// opaque
```

## 中断请求

Fetch API 支持通过 AbortController/AbortSignal 对中断请求。调用 AbortController. abort()会中断所有网络传输，特别适合希望停止传输大型负载的情况。中断进行中的 fetch()请求会导致包含错误的拒绝。

```javascript
let abortController = new AbortController(); 

fetch('wikipedia.zip', { signal: abortController.signal }) 
 .catch(() => console.log('aborted!'); 
 
// 10 毫秒后中断请求
setTimeout(() => abortController.abort(), 10); 

// 已经中断
```

## Headers

Headers 对象是所有外发请求和入站响应头部的容器。每个外发的 Request 实例都包含一个空的Headers 实例，可以通过 Request.prototype.headers 访问，每个入站 Response 实例也可以通过Response.prototype.headers 访问包含着响应头部的 Headers 对象。这两个属性都是可修改属性。另外，使用 new Headers()也可以创建一个新实例。

## Headers 与 **Map** 的相似之处

Headers 对象与 Map 对象极为相似。这是合理的，因为 HTTP 头部本质上是序列化后的键/值对，它们的 JavaScript 表示则是中间接口。Headers 与 Map 类型都有 get()、set()、has() 和 delete() 等实例方法，如下面的代码所示：

```javascript
let h = new Headers(); 
let m = new Map(); 

// 设置键
h.set('foo', 'bar'); 
m.set('foo', 'bar'); 

// 检查键
console.log(h.has('foo')); // true 
console.log(m.has('foo')); // true 
console.log(h.has('qux')); // false 
console.log(m.has('qux')); // false 

// 获取值
console.log(h.get('foo')); // bar 
console.log(m.get('foo')); // bar 

// 更新值
h.set('foo', 'baz'); 
m.set('foo', 'baz');

// 取得更新的值
console.log(h.get('foo')); // baz 
console.log(m.get('foo')); // baz 

// 删除值
h.delete('foo'); 
m.delete('foo'); 

// 确定值已经删除
console.log(h.get('foo')); // undefined 
console.log(m.get('foo')); // undefined 
```

Headers 和 Map 都可以使用一个可迭代对象来初始化，比如：

```javascript
let seed = [['foo', 'bar']]; 
let h = new Headers(seed); 
let m = new Map(seed); 
console.log(h.get('foo')); // bar 
console.log(m.get('foo')); // bar 
```

而且，它们也都有相同的 keys()、values()和 entries()迭代器接口：

```javascript
let seed = [['foo', 'bar'], ['baz', 'qux']]; 
let h = new Headers(seed); 
let m = new Map(seed); 
console.log(...h.keys()); // foo, baz 
console.log(...m.keys()); // foo, baz 
console.log(...h.values()); // bar, qux 
console.log(...m.values()); // bar, qux 
console.log(...h.entries()); // ['foo', 'bar'], ['baz', 'qux'] 
console.log(...m.entries()); // ['foo', 'bar'], ['baz', 'qux']
```

##  **Headers** 独有的特性

Headers 并不是与 Map 处处都一样。在初始化 Headers 对象时，也可以使用键/值对形式的对象，而 Map 则不可以：

```javascript
let seed = {foo: 'bar'}; 

let h = new Headers(seed); 
console.log(h.get('foo')); // bar 

let m = new Map(seed); 
// TypeError: object is not iterable
```

一个 HTTP 头部字段可以有多个值，而 Headers 对象通过 append()方法支持添加多个值。在Headers 实例中还不存在的头部上调用 append()方法相当于调用 set()。后续调用会以逗号为分隔符拼接多个值：

```javascript
let h = new Headers();

h.append('foo', 'bar'); 
console.log(h.get('foo')); // "bar"

h.append('foo', 'baz'); 
console.log(h.get('foo')); // "bar, baz"
```

## 头部护卫

某些情况下，并非所有 HTTP 头部都可以被客户端修改，而 Headers 对象使用护卫来防止不被允许的修改。不同的护卫设置会改变 set()、append()和 delete()的行为。违反护卫限制会抛出TypeError。

Headers 实例会因来源不同而展现不同的行为，它们的行为由护卫来控制。JavaScript 可以决定Headers 实例的护卫设置。下表列出了不同的护卫设置和每种设置对应的行为。

| 护 卫           | 适用情形                                                     | 限 制                                                        |
| --------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| none            | 在通过构造函数创建 Headers 实例时激活                        | 无                                                           |
| request         | 在通过构造函数初始化 Request对象，且 mode 值为非 no-cors 时激活 | 不允许修改禁止修改的头部（参见 MDN 文档中的 forbidden header name 词条） |
| request-no-cors | 在通过构造函数初始化 Request对象，且 mode值为 no-cors 时激活 | 不允许修改非简单头部（参见 MDN 文档中的simple header 词条）  |
| response        | 在通过构造函数初始化 Response 对象时激活                     | 不允许修改禁止修改的响应头部（参见 MDN 文档中的 forbidden response header name 词条） |
| immutable       | 在通过 error()或 redirect()静态方法初始化 Response 对象时激活 | 不允许修改任何头部                                           |

## Request对象

顾名思义，Request 对象是获取资源请求的接口。这个接口暴露了请求的相关信息，也暴露了使用请求体的不同方式。

### 创建 **Request** 对象

可以通过构造函数初始化 Request 对象。为此需要传入一个 input 参数，一般是 URL：

```javascript
let r = new Request('https://foo.com'); 
console.log(r); 
// Request {...}
```

Request 构造函数也接收第二个参数——一个 init 对象。这个 init 对象与前面介绍的 fetch() 的 init 对象一样。没有在 init 对象中涉及的值则会使用默认值：

```javascript
// 用所有默认值创建 Request 对象
console.log(new Request('')); 
// Request { 
// bodyUsed: false 
// cache: "default" 
// credentials: "same-origin" 
// destination: "" 
// headers: Headers {}
// integrity: "" 
// keepalive: false 
// method: "GET" 
// mode: "cors" 
// redirect: "follow" 
// referrer: "about:client" 
// referrerPolicy: "" 
// signal: AbortSignal {aborted: false, onabort: null} 
// url: "<current URL>" 
// }

// 用指定的初始值创建 Request 对象
console.log(new Request('https://foo.com', { method: 'POST' }));
 
// Request { 
// bodyUsed: false 
// cache: "default" 
// credentials: "same-origin" 
// destination: "" 
// headers: Headers {} 
// integrity: "" 
// keepalive: false 
// method: "POST"
// mode: "cors" 
// redirect: "follow" 
// referrer: "about:client" 
// referrerPolicy: "" 
// signal: AbortSignal {aborted: false, onabort: null} 
// url: "https://foo.com/"
// }
```

### 克隆 **Request** 对象

Fetch API 提供了两种不太一样的方式用于创建 Request 对象的副本：使用 Request 构造函数和使用 clone()方法。

将 Request 实例作为 input 参数传给 Request 构造函数，会得到该请求的一个副本：

```javascript
let r1 = new Request('https://foo.com'); 
let r2 = new Request(r1); 

console.log(r2.url); // https://foo.com/
```

如果再传入 init 对象，则 init 对象的值会覆盖源对象中同名的值：

```javascript
let r1 = new Request('https://foo.com'); 
let r2 = new Request(r1, {method: 'POST'}); 

console.log(r1.method); // GET 
console.log(r2.method); // POST
```

这种克隆方式并不总能得到一模一样的副本。最明显的是，第一个请求的请求体会被标记为“已使用”：

```javascript
let r1 = new Request('https://foo.com', { method: 'POST', body: 'foobar' }); 
let r2 = new Request(r1); 

console.log(r1.bodyUsed); // true 
console.log(r2.bodyUsed); // false
```

如果源对象与创建的新对象不同源，则 referrer 属性会被清除。此外，如果源对象的 mode 为navigate，则会被转换为 same-origin。

第二种克隆 Request 对象的方式是使用 clone()方法，这个方法会创建一模一样的副本，任何值都不会被覆盖。与第一种方式不同，这种方法不会将任何请求的请求体标记为“已使用”：

```javascript
let r1 = new Request('https://foo.com', { method: 'POST', body: 'foobar' }); 
let r2 = r1.clone(); 

console.log(r1.url); // https://foo.com/ 
console.log(r2.url); // https://foo.com/ 

console.log(r1.bodyUsed); // false 
console.log(r2.bodyUsed); // false
```

如果请求对象的 bodyUsed 属性为 true（即请求体已被读取），那么上述任何一种方式都不能用来创建这个对象的副本。在请求体被读取之后再克隆会导致抛出 TypeError。

```javascript
let r = new Request('https://foo.com'); 
r.clone(); 
new Request(r); 
// 没有错误

r.text(); // 设置 bodyUsed 为 true 
r.clone(); 
// TypeError: Failed to execute 'clone' on 'Request': Request body is already used 

new Request(r); 
// TypeError: Failed to construct 'Request': Cannot construct a Request with a 
Request object that has already been used.
```

### 在 **fetch()** 中使用 **Request** 对象

fetch()和 Request 构造函数拥有相同的函数签名并不是巧合。在调用 fetch()时，可以传入已经创建好的 Request 实例而不是 URL。与 Request 构造函数一样，传给 fetch()的 init 对象会覆盖传入请求对象的值：

```javascript
let r = new Request('https://foo.com'); 

// 向 foo.com 发送 GET 请求
fetch(r); 

// 向 foo.com 发送 POST 请求
fetch(r, { method: 'POST' });
```

fetch()会在内部克隆传入的 Request 对象。与克隆 Request 一样，fetch()也不能拿请求体已经用过的 Request 对象来发送请求：

```javascript
let r = new Request('https://foo.com', { method: 'POST', body: 'foobar' }); 

r.text(); 

fetch(r); 
// TypeError: Cannot construct a Request with a Request object that has already been used.
```

关键在于，通过 fetch 使用 Request 会将请求体标记为已使用。也就是说，有请求体的 Request 只能在一次 fetch 中使用。（不包含请求体的请求不受此限制。）演示如下：

```javascript
let r = new Request('https://foo.com',  { method: 'POST', body: 'foobar' }); 

fetch(r); 
fetch(r); 
// TypeError: Cannot construct a Request with a Request object that has already been used.
```

要想基于包含请求体的相同 Request 对象多次调用 fetch()，必须在第一次发送 fetch()请求前调用 clone()：

```javascript
let r = new Request('https://foo.com', { method: 'POST', body: 'foobar' }); 

// 3 个都会成功
fetch(r.clone()); 
fetch(r.clone()); 
fetch(r);
```

## Response对象

顾名思义，Response 对象是获取资源响应的接口。这个接口暴露了响应的相关信息，也暴露了使用响应体的不同方式。

### 创建 **Response** 对象

可以通过构造函数初始化 Response 对象且不需要参数。此时响应实例的属性均为默认值，因为它并不代表实际的 HTTP 响应：

```javascript
let r = new Response(); 
console.log(r); 
// Response { 
// body: (...) 
// bodyUsed: false 
// headers: Headers {} 
// ok: true 
// redirected: false 
// status: 200 
// statusText: "OK" 
// type: "default" 
// url: "" 
// }
```

Response 构造函数接收一个可选的 body 参数。这个 body 可以是 null，等同于 fetch()参数 init 中的 body。还可以接收一个可选的 init 对象，这个对象可以包含下表所列的键和值。

| 键         | 值                                                           |
| ---------- | ------------------------------------------------------------ |
| headers    | 必须是 Headers 对象实例或包含字符串键/值对的常规对象实例，默认为没有键/值对的 Headers 对象 |
| status     | 表示 HTTP 响应状态码的整数，默认为 200                       |
| statusText | 表示 HTTP 响应状态的字符串，默认为空字符串                   |

可以像下面这样使用 body 和 init 来构建 Response 对象：

```javascript
let r = new Response('foobar', {
    status: 418,
    statusText: 'I\'m a teapot',
});
console.log(r); 
// Response { 
// body: (...) 
// bodyUsed: false 
// headers: Headers {} 
// ok: false 
// redirected: false 
// status: 418
// statusText: "I'm a teapot"
// type: "default" 
// url: "" 
// }
```

大多数情况下，产生 Response 对象的主要方式是调用 fetch()，它返回一个最后会解决为Response 对象的期约，这个 Response 对象代表实际的 HTTP 响应。下面的代码展示了这样得到的Response 对象：

```javascript
fetch('bar.txt').then((response) => {
    console.log(response);
});
// Response { 
// body: (...) 
// bodyUsed: false 
// headers: Headers {} 
// ok: true 
// redirected: false 
// status: 200 
// statusText: "OK" 
// type: "basic" 
// url: "https://foo.com/" 
// }
```

Response 类还有两个用于生成 Response 对象的静态方法：Response.redirect()和 Response. error()。前者接收一个 URL 和一个重定向状态码（301、302、303、307 或 308），返回重定向的 Response对象：

```javascript
console.log(Response.redirect('https://foo.com', 301)); 
// Response { 
// body: (...) 
// bodyUsed: false 
// headers: Headers {} 
// ok: false 
// redirected: false 
// status: 301 
// statusText: "" 
// type: "default" 
// url: "" 
// }
```

提供的状态码必须对应重定向，否则会抛出错误：

```javascript
Response.redirect('https://foo.com', 200); 
// RangeError: Failed to execute 'redirect' on 'Response': Invalid status code
```

另一个静态方法 Response.error()用于产生表示网络错误的 Response 对象（网络错误会导致fetch()期约被拒绝）。

```javascript
console.log(Response.error()); 
// Response { 
// body: (...) 
// bodyUsed: false 
// headers: Headers {} 
// ok: false 
// redirected: false 
// status: 0 
// statusText: "" 
// type: "error" 
// url: "" 
// }
```

### 读取响应状态信息

Response 对象包含一组只读属性，描述了请求完成后的状态，如下表所示。

| 属 性      | 值                                                           |
| ---------- | ------------------------------------------------------------ |
| headers    | 响应包含的 Headers 对象                                      |
| ok         | 布尔值，表示 HTTP 状态码的含义。200~299 的状态码返回 true，其他状态码返回 false |
| redirected | 布尔值，表示响应是否至少经过一次重定向                       |
| status     | 整数，表示响应的 HTTP 状态码                                 |
| statusText | 字符串，包含对 HTTP 状态码的正式描述。这个值派生自可选的 HTTP Reason-Phrase 字段，因此如果服务器以 Reason-Phrase 为由拒绝响应，这个字段可能是空字符串 |
| type       | 字符串，包含响应类型。可能是下列字符串值之一<br />+ basic：表示标准的同源响应<br />+ cors：表示标准的跨源响应<br />+ error：表示响应对象是通过 Response.error()创建的<br />+ opaque：表示 no-cors 的 fetch()返回的跨源响应<br />+ opaqueredirect：表示对 redirect 设置为 manual 的请求的响应 |
| url        | 包含响应 URL 的字符串。对于重定向响应，这是最终的 URL，非重定向响应就是它产生的 |

以下代码演示了返回 200、302、404 和 500 状态码的 URL 对应的响应：

```javascript
fetch('//foo.com').then(console.log); 
// Response { 
// body: (...) 
// bodyUsed: false 
// headers: Headers {} 
// ok: true 
// redirected: false 
// status: 200 
// statusText: "OK" 
// type: "basic" 
// url: "https://foo.com/" 
// }

fetch('//foo.com/redirect-me').then(console.log); 
// Response { 
// body: (...) 
// bodyUsed: false 
// headers: Headers {} 
// ok: true 
// redirected: true
// status: 200 
// statusText: "OK" 
// type: "basic" 
// url: "https://foo.com/redirected-url/" 
// }

fetch('//foo.com/does-not-exist').then(console.log); 
// Response { 
// body: (...) 
// bodyUsed: false 
// headers: Headers {} 
// ok: false 
// redirected: true 
// status: 404 
// statusText: "Not Found"
// type: "basic" 
// url: "https://foo.com/does-not-exist/" 
// }

fetch('//foo.com/throws-error').then(console.log); 
// Response { 
// body: (...) 
// bodyUsed: false 
// headers: Headers {} 
// ok: false 
// redirected: true 
// status: 500 
// statusText: "Internal Server Error"
// type: "basic" 
// url: "https://foo.com/throws-error/" 
// }
```

### 克隆 **Response** 对象

克隆 Response 对象的主要方式是使用 clone()方法，这个方法会创建一个一模一样的副本，不会覆盖任何值。这样不会将任何请求的请求体标记为已使用：

```javascript
let r1 = new Response('foobar'); 
let r2 = r1.clone(); 

console.log(r1.bodyUsed); // false 
console.log(r2.bodyUsed); // false
```

如果响应对象的 bodyUsed 属性为 true（即响应体已被读取），则不能再创建这个对象的副本。在响应体被读取之后再克隆会导致抛出 TypeError。

```javascript
let r = new Response('foobar'); 
r.clone(); 
// 没有错误

r.text(); // 设置 bodyUsed 为 true 
r.clone(); 
// TypeError: Failed to execute 'clone' on 'Response': Response body is already used 
```

有响应体的 Response 对象只能读取一次。（不包含响应体的 Response 对象不受此限制。）比如：

```javascript
let r = new Response('foobar'); 
r.text().then(console.log); // foobar 
r.text().then(console.log); 
// TypeError: Failed to execute 'text' on 'Response': body stream is locked
```

要多次读取包含响应体的同一个 Response 对象，必须在第一次读取前调用 clone()：

```javascript
let r = new Response('foobar'); 
r.clone().text().then(console.log); // foobar 
r.clone().text().then(console.log); // foobar 
r.text().then(console.log); // foobar
```

此外，通过创建带有原始响应体的 Response 实例，可以执行伪克隆操作。关键是这样不会把第一个 Response 实例标记为已读，而是会在两个响应之间共享：

```javascript
let r1 = new Response('foobar'); 
let r2 = new Response(r1.body); 

console.log(r1.bodyUsed); // false 
console.log(r2.bodyUsed); // false 

r2.text().then(console.log); // foobar 
r1.text().then(console.log); 
// TypeError: Failed to execute 'text' on 'Response': body stream is locked
```

## Request**、**Response及 **Body** 混入

Request 和 Response 都使用了 Fetch API 的 Body 混入，以实现两者承担有效载荷的能力。这个混入为两个类型提供了只读的 body 属性（实现为 ReadableStream）、只读的 bodyUsed 布尔值（表示 body 流是否已读）和一组方法，用于从流中读取内容并将结果转换为某种 JavaScript 对象类型。

通常，将 Request 和 Response 主体作为流来使用主要有两个原因。一个原因是有效载荷的大小可能会导致网络延迟，另一个原因是流 API 本身在处理有效载荷方面是有优势的。除此之外，最好是一次性获取资源主体。

Body 混入提供了 5 个方法，用于将 ReadableStream 转存到缓冲区的内存里，将缓冲区转换为某种 JavaScript 对象类型，以及通过期约来产生结果。在解决之前，期约会等待主体流报告完成及缓冲被解析。这意味着客户端必须等待响应的资源完全加载才能访问其内容。



### Body.text()

Body.text()方法返回期约，解决为将缓冲区转存得到的 UTF-8 格式字符串。下面的代码展示了在 Response 对象上使用 Body.text()：

```javascript
fetch('https://foo.com') 
 .then((response) => response.text()) 
 .then(console.log); 
// <!doctype html><html lang="en"> 
// <head> 
// <meta charset="utf-8"> 
// ...
```

以下代码展示了在 Request 对象上使用 Body.text()：

```javascript
let request = new Request('https://foo.com', { method: 'POST', body: 'barbazqux' }); 
 
request.text() 
 .then(console.log); 
 
// barbazqux
```

### Body.json()

Body.json()方法返回期约，解决为将缓冲区转存得到的 JSON。下面的代码展示了在 Response 对象上使用 Body.json()：

```javascript
fetch('https://foo.com/foo.json') 
 .then((response) => response.json()) 
 .then(console.log); 
// {"foo": "bar"}
```

以下代码展示了在 Request 对象上使用 Body.json()：

```javascript
let request = new Request('https://foo.com', { method:'POST', body: JSON.stringify({ bar: 'baz' }) }); 

request.json() 
 .then(console.log); 
 
// {bar: 'baz'}
```

### Body.formData()

浏览器可以将 FormData 对象序列化/反序列化为主体。例如，下面这个 FormData 实例：

```javascript
let myFormData = new FormData(); 
myFormData.append('foo', 'bar');
```

在通过 HTTP 传送时，WebKit 浏览器会将其序列化为下列内容：

```javascript
------WebKitFormBoundarydR9Q2kOzE6nbN7eR 
Content-Disposition: form-data; name="foo"
bar 
------WebKitFormBoundarydR9Q2kOzE6nbN7eR--
```

Body.formData()方法返回期约，解决为将缓冲区转存得到的 FormData 实例。下面的代码展示

了在 Response 对象上使用 Body.formData()：

```javascript
fetch('https://foo.com/form-data') 
 .then((response) => response.formData()) 
 .then((formData) => console.log(formData.get('foo')); 
 
// bar
```

以下代码展示了在 Request 对象上使用 Body.formData()：

```javascript
let myFormData = new FormData(); 
myFormData.append('foo', 'bar'); 

let request = new Request('https://foo.com', 
 { method:'POST', body: myFormData }); 
 
request.formData() 
 .then((formData) => console.log(formData.get('foo')); 
 
// bar
```

### Body.arrayBuffer()

有时候，可能需要以原始二进制格式查看和修改主体。为此，可以使用 Body.arrayBuffer()将主体内容转换为 ArrayBuffer 实例。Body.arrayBuffer()方法返回期约，解决为将缓冲区转存得到的 ArrayBuffer 实例。下面的代码展示了在 Response 对象上使用 Body.arrayBuffer()：

```javascript
fetch('https://foo.com') 
 .then((response) => response.arrayBuffer()) 
 .then(console.log); 
// ArrayBuffer(...) {} 

```

以下代码展示了在 Request 对象上使用 Body.arrayBuffer()：

```javascript
let request = new Request('https://foo.com', 
 { method:'POST', body: 'abcdefg' }); 
 
// 以整数形式打印二进制编码的字符串
request.arrayBuffer() 
 .then((buf) => console.log(new Int8Array(buf))); 
// Int8Array(7) [97, 98, 99, 100, 101, 102, 103]
```

### Body.blob()

有时候，可能需要以原始二进制格式使用主体，不用查看和修改。为此，可以使用 Body.blob()将主体内容转换为 Blob 实例。Body.blob()方法返回期约，解决为将缓冲区转存得到的 Blob 实例。下面的代码展示了在 Response 对象上使用 Body.blob()：

```javascript
fetch('https://foo.com') 
 .then((response) => response.blob()) 
 .then(console.log); 
// Blob(...) {size:..., type: "..."}
```

以下代码展示了在 Request 对象上使用 Body.blob()：

```javascript
let request = new Request('https://foo.com', 
 { method:'POST', body: 'abcdefg' }); 
 
request.blob() 
 .then(console.log); 
// Blob(7) {size: 7, type: "text/plain;charset=utf-8"}
```

# 模块

现代 JavaScript 开发毋庸置疑会遇到代码量大和广泛使用第三方库的问题。解决这个问题的方案通常需要把代码拆分成很多部分，然后再通过某种方式将它们连接起来。

在 ECMAScript 6 模块规范出现之前，虽然浏览器原生不支持模块的行为，但迫切需要这样的行为。ECMAScript 同样不支持模块，因此希望使用模块模式的库或代码库必须基于 JavaScript 的语法和词法特性“伪造”出类似模块的行为。

因为 JavaScript 是异步加载的解释型语言，所以得到广泛应用的各种模块实现也表现出不同的形态。这些不同的形态决定了不同的结果，但最终它们都实现了经典的模块模式。

## 理解模块模式

将代码拆分成独立的块，然后再把这些块连接起来可以通过模块模式来实现。这种模式背后的思想很简单：把逻辑分块，各自封装，相互独立，每个块自行决定对外暴露什么，同时自行决定引入执行哪些外部代码。不同的实现和特性让这些基本的概念变得有点复杂，但这个基本的思想是所有 JavaScript 模块系统的基础。

### 模块标识符

模块标识符是所有模块系统通用的概念。模块系统本质上是键/值实体，其中每个模块都有个可用于引用它的标识符。这个标识符在模拟模块的系统中可能是字符串，在原生实现的模块系统中可能是模块文件的实际路径。

有的模块系统支持明确声明模块的标识，还有的模块系统会隐式地使用文件名作为模块标识符。不管怎样，完善的模块系统一定不会存在模块标识冲突的问题，且系统中的任何模块都应该能够无歧义地引用其他模块。

将模块标识符解析为实际模块的过程要根据模块系统对标识符的实现。原生浏览器模块标识符必须提供实际 JavaScript 文件的路径。除了文件路径，Node.js 还会搜索 node_modules 目录，用标识符去匹配包含 index.js 的目录。

### 模块依赖

模块系统的核心是管理依赖。指定依赖的模块与周围的环境会达成一种契约。本地模块向模块系统声明一组外部模块（依赖），这些外部模块对于当前模块正常运行是必需的。模块系统检视这些依赖，

进而保证这些外部模块能够被加载并在本地模块运行时初始化所有依赖。每个模块都会与某个唯一的标识符关联，该标识符可用于检索模块。这个标识符通常是 JavaScript 文件的路径，但在某些模块系统中，这个标识符也可以是在模块本身内部声明的命名空间路径字符串。

### 模块加载

加载模块的概念派生自依赖契约。当一个外部模块被指定为依赖时，本地模块期望在执行它时，依赖已准备好并已初始化。

在浏览器中，加载模块涉及几个步骤。加载模块涉及执行其中的代码，但必须是在所有依赖都加载并执行之后。如果浏览器没有收到依赖模块的代码，则必须发送请求并等待网络返回。收到模块代码之后，浏览器必须确定刚收到的模块是否也有依赖。然后递归地评估并加载所有依赖，直到所有依赖模块都加载完成。只有整个依赖图都加载完成，才可以执行入口模块。



## 使用 ES6 模块

ES6 最大的一个改进就是引入了模块规范。这个规范全方位简化了之前出现的模块加载器，原生浏览器支持意味着加载器及其他预处理都不再必要。从很多方面看，ES6 模块系统是集 AMD 和 CommonJS 之大成者。

### 模块标签及定义

ECMAScript 6 模块是作为一整块 JavaScript 代码而存在的。带有 type="module"属性的 `<script>` 标签会告诉浏览器相关代码应该作为模块执行，而不是作为传统的脚本执行。模块可以嵌入在网页中，也可以作为外部文件引入：

```html
<script type="module"> 
 // 模块代码
</script> 

<script type="module" src="path/to/myModule.js"></script>
```

即使与常规 JavaScript 文件处理方式不同，JavaScript 模块文件也没有专门的内容类型。

与传统脚本不同，所有模块都会像 `<script defer>` 加载的脚本一样按顺序执行。解析到 `<script type="module">` 标签后会立即下载模块文件，但执行会延迟到文档解析完成。无论对嵌入的模块代码，还是引入的外部模块文件，都是这样。`<script type="module">` 在页面中出现的顺序就是它们执行

的顺序。与 `<script defer>` 一样，修改模块标签的位置，无论是在 `<head>` 还是在 `<body>` 中，只会影响文件什么时候加载，而不会影响模块什么时候加载。

下面演示了嵌入模块代码的执行顺序：

```html
<!-- 第二个执行 --> 
<script type="module"></script> 
<!-- 第三个执行 --> 
<script type="module"></script> 
<!-- 第一个执行 --> 
<script></script> 
另外，可以改为加载外部 JS 模块定义：
<!-- 第二个执行 --> 
<script type="module" src="module.js"></script> 
<!-- 第三个执行 --> 
<script type="module" src="module.js"></script> 
<!-- 第一个执行 --> 
<script></script>
```

也可以给模块标签添加 async 属性。这样影响就是双重的：不仅模块执行顺序不再与 `<script>` 标签在页面中的顺序绑定，模块也不会等待文档完成解析才执行。不过，入口模块仍必须等待其依赖加载完成。

与 `<script type="module">` 标签关联的 ES6 模块被认为是模块图中的入口模块。一个页面上有多少个入口模块没有限制，重复加载同一个模块也没有限制。同一个模块无论在一个页面中被加载多少次，也不管它是如何加载的，实际上都只会加载一次，如下面的代码所示：

```html
<!-- moduleA 在这个页面上只会被加载一次 --> 
<script type="module"> 
 import './moduleA.js' 
<script> 
<script type="module"> 
 import './moduleA.js' 
<script> 
<script type="module" src="./moduleA.js"></script> 
<script type="module" src="./moduleA.js"></script>
```

嵌入的模块定义代码不能使用 import 加载到其他模块。只有通过外部文件加载的模块才可以使用 import 加载。因此，嵌入模块只适合作为入口模块。

### 模块加载

ECMAScript 6 模块的独特之处在于，既可以通过浏览器原生加载，也可以与第三方加载器和构建工具一起加载。有些浏览器还没有原生支持 ES6 模块，因此可能还需要第三方工具。事实上，很多时候使用第三方工具可能会更方便。

完全支持 ECMAScript 6 模块的浏览器可以从顶级模块加载整个依赖图，且是异步完成的。浏览器会解析入口模块，确定依赖，并发送对依赖模块的请求。这些文件通过网络返回后，浏览器就会解析它们的内容，确定它们的依赖，如果这些二级依赖还没有加载，则会发送更多请求。这个异步递归加载过程会持续到整个应用程序的依赖图都解析完成。解析完依赖图，应用程序就可以正式加载模块了。

这个过程与 AMD 风格的模块加载非常相似。模块文件按需加载，且后续模块的请求会因为每个依赖模块的网络延迟而同步延迟。即，如果 moduleA 依赖 moduleB，moduleB 依赖 moduleC。浏览器在对 moduleB 的请求完成之前并不知道要请求 moduleC。这种加载方式效率很高，也不需要外部工具，但加载大型应用程序的深度依赖图可能要花费很长时间。

### 模块行为

ECMAScript 6 模块借用了 CommonJS 和 AMD 的很多优秀特性。下面简单列举一些。

+ 模块代码只在加载后执行。
+ 模块只能加载一次。
+ 模块是单例。
+ 模块可以定义公共接口，其他模块可以基于这个公共接口观察和交互。
+ 模块可以请求加载其他模块。
+ 支持循环依赖。



ES6 模块系统也增加了一些新行为。

+ ES6 模块默认在严格模式下执行。
+ ES6 模块不共享全局命名空间。
+ 模块顶级 this 的值是 undefined（常规脚本中是 window）。
+ 模块中的 var 声明不会添加到 window 对象。
+ ES6 模块是异步加载和执行的。



浏览器运行时在知道应该把某个文件当成模块时，会有条件地按照上述 ECMAScript 6 模块行为来施加限制。与 `<script type="module">` 关联或者通过 import 语句加载的 JavaScript 文件会被认定为模块。

### 模块导出

ES6 模块的公共导出系统与 CommonJS 非常相似。控制模块的哪些部分对外部可见的是 export 关键字。ES6 模块支持两种导出：命名导出和默认导出。不同的导出方式对应不同的导入方式，下一节会介绍导入。

export 关键字用于声明一个值为命名导出。导出语句必须在模块顶级，不能嵌套在某个块中：

```javascript
// 允许
export ... 
// 不允许
if (condition) { 
 export ... 
}
```

导出值对模块内部 JavaScript 的执行没有直接影响，因此 export 语句与导出值的相对位置或者 export 关键字在模块中出现的顺序没有限制。export 语句甚至可以出现在它要导出的值之前：

```javascript
// 允许
const foo = 'foo'; 
export { foo }; 
// 允许
export const foo = 'foo'; 
// 允许，但应该避免
export { foo }; 
const foo = 'foo';
```

命名导出（named export）就好像模块是被导出值的容器。行内命名导出，顾名思义，可以在同一行执行变量声明。下面展示了一个声明变量同时又导出变量的例子。外部模块可以导入这个模块，而 foo 将成为这个导入模块的一个属性：

```javascript
export const foo = 'foo';
```

变量声明跟导出可以不在一行。可以在 export 子句中执行声明并将标识符导出到模块的其他地方：

```javascript
const foo = 'foo'; 
export { foo };
```

导出时也可以提供别名，别名必须在 export 子句的大括号语法中指定。因此，声明值、导出值和为导出值提供别名不能在一行完成。在下面的例子中，导入这个模块的外部模块可以使用 myFoo 访问导出的值：

```javascript
const foo = 'foo'; 
export { foo as myFoo };
```

因为 ES6 命名导出可以将模块作为容器，所以可以在一个模块中声明多个命名导出。导出的值可以在导出语句中声明，也可以在导出之前声明：

```javascript
export const foo = 'foo'; 
export const bar = 'bar'; 
export const baz = 'baz';
```

考虑到导出多个值是常见的操作，ES6 模块也支持对导出声明分组，可以同时为部分或全部导出值指定别名：

```javascript
const foo = 'foo'; 
const bar = 'bar'; 
const baz = 'baz'; 
export { foo, bar as myBar, baz };
```

默认导出（default export）就好像模块与被导出的值是一回事。默认导出使用 default 关键字将一个值声明为默认导出，每个模块只能有一个默认导出。重复的默认导出会导致 SyntaxError。

下面的例子定义了一个默认导出，外部模块可以导入这个模块，而这个模块本身就是 foo 的值：

```javascript
const foo = 'foo'; 
export default foo;
```

另外，ES6 模块系统会识别作为别名提供的 default 关键字。此时，虽然对应的值是使用命名语法导出的，实际上则会成为默认导出：

```javascript
const foo = 'foo'; 

// 等同于 export default foo; 
export { foo as default };
```

因为命名导出和默认导出不会冲突，所以 ES6 支持在一个模块中同时定义这两种导出：

```javascript
const foo = 'foo'; 
const bar = 'bar'; 

export { bar }; 
export default foo;
```

这两个 export 语句可以组合为一行：

```javascript
const foo = 'foo'; 
const bar = 'bar'; 

export { foo as default, bar };
```

ES6 规范对不同形式的 export 语句中可以使用什么不可以使用什么规定了限制。某些形式允许声明和赋值，某些形式只允许表达式，而某些形式则只允许简单标识符。注意，有的形式使用了分号，有的则没有：

```javascript
// 命名行内导出
export const baz = 'baz'; 
export const foo = 'foo', bar = 'bar'; 
export function foo() {} 
export function* foo() {} 
export class Foo {} 

// 命名子句导出
export { foo }; 
export { foo, bar }; 
export { foo as myFoo, bar }; 

// 默认导出
export default 'foo'; 
export default 123; 
export default /[a-z]*/; 
export default { foo: 'foo' };

export { foo, bar as default }; 
export default foo 
export default function() {} 
export default function foo() {} 
export default function*() {} 
export default class {}

// 会导致错误的不同形式：

// 行内默认导出中不能出现变量声明
export default const foo = 'bar'; 

// 只有标识符可以出现在 export 子句中
export { 123 as foo } 

// 别名只能在 export 子句中出现
export const foo = 'foo' as myFoo;
```

### 模块导入

模块可以通过使用 import 关键字使用其他模块导出的值。与 export 类似，import 必须出现在模块的顶级：

```javascript
// 允许
import ... 
// 不允许
if (condition) { 
 import ... 
}
```

import 语句被提升到模块顶部。因此，与 export 关键字类似，import 语句与使用导入值的语句的相对位置并不重要。不过，还是推荐把导入语句放在模块顶部。

```javascript
// 允许
import { foo } from './fooModule.js'; 
console.log(foo); // 'foo' 

// 允许，但应该避免
console.log(foo); // 'foo' 
import { foo } from './fooModule.js';
```

模块标识符可以是相对于当前模块的相对路径，也可以是指向模块文件的绝对路径。它必须是纯字符串，不能是动态计算的结果。例如，不能是拼接的字符串。

如果在浏览器中通过标识符原生加载模块，则文件必须带有.js 扩展名，不然可能无法正确解析。不过，如果是通过构建工具或第三方模块加载器打包或解析的 ES6 模块，则可能不需要包含文件扩展名。

```javascript
// 解析为/components/bar.js 
import ... from './bar.js';

// 解析为/bar.js 
import ... from '../bar.js';

// 解析为/bar.js 
import ... from '/bar.js';
```

不是必须通过导出的成员才能导入模块。如果不需要模块的特定导出，但仍想加载和执行模块以利用其副作用，可以只通过路径加载它：

```javascript
import './foo.js';
```

导入对模块而言是只读的，实际上相当于 const 声明的变量。在使用*执行批量导入时，赋值给别名的命名导出就好像使用 Object.freeze()冻结过一样。直接修改导出的值是不可能的，但可以修改导出对象的属性。同样，也不能给导出的集合添加或删除导出的属性。要修改导出的值，必须使用有内部变量和属性访问权限的导出方法。

```javascript
import foo, * as Foo './foo.js'; 

foo = 'foo'; // 错误

Foo.foo = 'foo'; // 错误

foo.bar = 'bar'; // 允许
```

命名导出和默认导出的区别也反映在它们的导入上。命名导出可以使用*批量获取并赋值给保存导出集合的别名，而无须列出每个标识符：

```javascript
const foo = 'foo', bar = 'bar', baz = 'baz'; 
export { foo, bar, baz } 
import * as Foo from './foo.js'; 
console.log(Foo.foo); // foo 
console.log(Foo.bar); // bar 
console.log(Foo.baz); // baz
```

要指名导入，需要把标识符放在 import 子句中。使用 import 子句可以为导入的值指定别名：

```javascript
import { foo, bar, baz as myBaz } from './foo.js'; 

console.log(foo); // foo 
console.log(bar); // bar 
console.log(myBaz); // baz
```

默认导出就好像整个模块就是导出的值一样。可以使用 default 关键字并提供别名来导入。也可以不使用大括号，此时指定的标识符就是默认导出的别名：

```javascript
// 等效
import { default as foo } from './foo.js'; 
import foo from './foo.js';
```

如果模块同时导出了命名导出和默认导出，则可以在 import 语句中同时取得它们。可以依次列出特定导出的标识符来取得，也可以使用*来取得：

```javascript
import foo, { bar, baz } from './foo.js'; 
import { default as foo, bar, baz } from './foo.js'; 
import foo, * as Foo from './foo.js';
```

### 模块转移导出

模块导入的值可以直接通过管道转移到导出。此时，也可以将默认导出转换为命名导出，或者相反。如果想把一个模块的所有命名导出集中在一块，可以像下面这样在 bar.js 中使用*导出：

```javascript
export * from './foo.js';
```

这样，foo.js 中的所有命名导出都会出现在导入 bar.js 的模块中。如果 foo.js 有默认导出，则该语法会忽略它。使用此语法也要注意导出名称是否冲突。如果 foo.js 导出 baz，bar.js 也导出 baz，则最终导出的是 bar.js 中的值。这个“重写”是静默发生的：

```javascript
foo.js
export const baz = 'origin:foo'; 

bar.js
export * from './foo.js'; 
export const baz = 'origin:bar'; 

main.js
import { baz } from './bar.js'; 
console.log(baz); // origin:bar
```

此外也可以明确列出要从外部模块转移本地导出的值。该语法支持使用别名：

```javascript
export { foo, bar as myBar } from './foo.js';
```

类似地，外部模块的默认导出可以重用为当前模块的默认导出：

```javascript
export { default } from './foo.js';
```

这样不会复制导出的值，只是把导入的引用传给了原始模块。在原始模块中，导入的值仍然是可用的，与修改导入相关的限制也适用于再次导出的导入。

在重新导出时，还可以在导入模块修改命名或默认导出的角色。比如，可以像下面这样将命名导出指定为默认导出：

```javascript
export { foo as default } from './foo.js';
```

### 工作者模块

ECMAScript 6 模块与 Worker 实例完全兼容。在实例化时，可以给工作者传入一个指向模块文件的路径，与传入常规脚本文件一样。Worker 构造函数接收第二个参数，用于说明传入的是模块文件。

下面是两种类型的 Worker 的实例化行为：

```javascript
// 第二个参数默认为{ type: 'classic' } 
const scriptWorker = new Worker('scriptWorker.js'); 

const moduleWorker = new Worker('moduleWorker.js', { type: 'module' });
```

在基于模块的工作者内部，self.importScripts()方法通常用于在基于脚本的工作者中加载外

部脚本，调用它会抛出错误。这是因为模块的 import 行为包含了 importScripts()。

### 向后兼容

ECMAScript 模块的兼容是个渐进的过程，能够同时兼容支持和不支持的浏览器对早期采用者是有价值的。对于想要尽可能在浏览器中原生使用 ECMAScript 6 模块的用户，可以提供两个版本的代码：基于模块的版本与基于脚本的版本。如果嫌麻烦，可以使用第三方模块系统（如 SystemJS）或在构建时将 ES6 模块进行转译，这都是不错的方案。

第一种方案涉及在服务器上检查浏览器的用户代理，与支持模块的浏览器名单进行匹配，然后基于匹配结果决定提供哪个版本的 JavaScript 文件。这个方法不太可靠，而且比较麻烦，不推荐。更好、更优雅的方案是利用脚本的 type 属性和 nomodule 属性。

浏览器在遇到 `<script>` 标签上无法识别的 type 属性时会拒绝执行其内容。对于不支持模块的浏览器，这意味着 `<script type="module">` 不会被执行。因此，可以在 `<script type="module">` 标签旁边添加一个回退 `<script>` 标签：

```html
// 不支持模块的浏览器不会执行这里的代码
<script type="module" src="module.js"></script> 

// 不支持模块的浏览器会执行这里的代码
<script src="script.js"></script>
```

当然，这样一来支持模块的浏览器就有麻烦了。此时，前面的代码会执行两次，显然这不是我们想要的结果。为了避免这种情况，原生支持 ECMAScript 6 模块的浏览器也会识别 nomodule 属性。此属性通

知支持 ES6 模块的浏览器不执行脚本。不支持模块的浏览器无法识别该属性，从而忽略这个属性的存在。因此，下面代码会生成一个设置，在这个设置中，支持模块和不支持模块的浏览器都只会执行一段脚本：

```html
// 支持模块的浏览器会执行这段脚本
// 不支持模块的浏览器不会执行这段脚本
<script type="module" src="module.js"></script> 

// 支持模块的浏览器不会执行这段脚本
// 不支持模块的浏览器会执行这段脚本
<script nomodule src="script.js"></script>
```

## 小结

模块模式是管理复杂性的永恒工具。开发者可以通过它创建逻辑彼此独立的代码段，在这些代码段之间声明依赖，并将它们连接在一起。此外，这种模式也是经证明能够优雅扩展到任意复杂度且跨平台的方案。

多年以来，CommonJS 和 AMD 这两个分别针对服务器端环境和受延迟限制的客户端环境的模块系统长期分裂。两个系统都获得了爆炸性增强，但为它们编写的代码则在很多方面不一致，经常也会带有冗余的样板代码。而且，这两个系统都没有在浏览器中实现。缺乏兼容导致出现了相关工具，从而让在浏览器中实现模块模式成为可能。

ECMAScript 6 规范重新定义了浏览器模块，集之前两个系统之长于一身，并通过更简单的声明性语法暴露出来。浏览器对原生模块的支持越来越好，但也提供了稳健的工具以实现从不支持到支持 ES6 模块的过渡。

