# JQuery

## 基本使用

jQuery 封装了 JavaScript 常用的功能代码，优化了 DOM 操作、事件处理、动画设计和 Ajax 交互。

jQuery的官网地址： https://jquery.com/，各版本下载：https://www.bootcdn.cn/jquery/

mdn 界面上的事件：https://www.w3.org/TR/DOM-Level-3-Events

### jQuery的入口函数

jQuery中常见的两种入口函数：

```javascript
// 第一种: 简单易用。
$(function () {   
    ...  // 此处是页面 DOM 加载完成的入口
}) ; 

// 第二种: 繁琐，但是也可以实现
$(document).ready(function(){
   ...  //  此处是页面DOM加载完成的入口
});
```

1. 等着 DOM 结构渲染完毕即可执行内部代码，不必等到所有外部资源加载完成，jQuery 帮我们完成了封装。
2. 相当于原生 js 中的 DOMContentLoaded。
3. 不同于原生 js 中的 load 事件是等页面文档、外部的 js 文件、css文件、图片加载完毕才执行内部代码。
4. 更推荐使用第一种方式。

### 顶级对象 $

`$` 是 jQuery 的别称，在代码中可以使用 jQuery 代替，但一般为了方便，通常都直接使用 `$`  。

`$` 是jQuery的顶级对象，相当于原生JavaScript中的 window。把元素利用$包装成jQuery对象，就可以调用jQuery 的方法。

### jQuery 与原生JS的区别

1. 用原生 JS 获取来的对象就是 DOM 对象
2. jQuery 方法获取的元素就是 jQuery 对象。
3. jQuery 对象本质是： 利用$对DOM 对象包装后产生的对象（伪数组形式存储）。
4. 只有 jQuery 对象才能使用 jQuery 方法，DOM 对象则使用原生的 JavaScirpt 方法。

### jQuery 对象和 DOM 对象转换

DOM 对象与 jQuery 对象之间是可以相互转换的。因为原生 js 比 jQuery 更大，原生的一些属性和方法 jQuery 没有给我们封装. 要想使用这些属性和方法需要把jQuery对象转换为DOM对象才能使用。

```javascript
// 1.DOM对象转换成jQuery对象，方法只有一种
var box = document.getElementById('box');  // 获取DOM对象
var jQueryObject = $(box);  // 把DOM对象转换为 jQuery 对象

// 2.jQuery 对象转换为 DOM 对象有两种方法：
//   2.1 jQuery对象[索引值]
var domObject1 = $('div')[0]

//   2.2 jQuery对象.get(索引值)
var domObject2 = $('div').get(0)
```

总结：实际开发比较常用的是把DOM对象转换为jQuery对象，这样能够调用功能更加强大的jQuery中的方法。

## jQuery 选择器

原生 JS 获取元素方式很多，很杂，而且兼容性情况不一致，因此 jQuery 给我们做了封装，使获取元素统一标准。

### 基础选择器

```js
$("选择器")   //  里面选择器直接写 CSS 选择器即可，但是要加引号
```

| 名称       | 用法            | 描述                     |
| ---------- | --------------- | ------------------------ |
| ID选择器   | $("#id")        | 获取指定ID的元素         |
| 全选选择器 | $("*")          | 匹配所有元素             |
| 类选择器   | $(".class")     | 获取同一类class的元素    |
| 标签选择器 | $("div")        | 获取同一类标签的所有元素 |
| 并集选择器 | $("div,p,li")   | 选取多个元素             |
| 交集选择器 | $("li.current") | 交集元素                 |

### 层级选择器

层级选择器最常用的两个分别为：后代选择器和子代选择器。

| 名称       | 用法       | 描述                                               |
| ---------- | ---------- | -------------------------------------------------- |
| 子代选择器 | $("ul>li") | 使用>号，获取亲儿子层级的元素,不会获取孙子层级元素 |
| 后代选择器 | $("ul li") | 使用空格，代表后代选择器，包括孙子等               |

**基础选择器和层级选择器案例代码**

```html
<body>
    <div>我是div</div>
    <div class="nav">我是nav div</div>
    <p>我是p</p>
    <ul>
        <li>我是ul 的</li>
        <li>我是ul 的</li>        
        <li>我是ul 的</li>
    </ul>
    <script>
        $(function() {
            console.log($(".nav"));
            console.log($("ul li"));
        })
    </script>
</body>
```

### 伪类选择器

筛选选择器，顾名思义就是在所有的选项中选择满足条件的进行筛选选择。

| 语法       | 用法          | 描述                               |
| ---------- | ------------- | ---------------------------------- |
| :first     | $('li:first') | 获取第一个li元素                   |
| :last      | $('li:last')  | 获取最后一个li元素                 |
| :eq(index) | $('li:eq(2)') | 获取索引号为2的li元素，索引从0开始 |
| :odd       | $('li:odd')   | 获取索引号为奇数的li元素           |
| :even      | $('li:even')  | 获取索引号为偶数的li元素           |

**案例代码**

```html
<body>
    <ul>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
    </ul>
    <ol>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
        <li>多个里面筛选几个</li>
    </ol>
    <script>
        $(function() {
            $("ul li:first").css("color", "red");
            $("ul li:eq(2)").css("color", "blue");
            $("ol li:odd").css("color", "skyblue");
            $("ol li:even").css("color", "pink");
        })
    </script>
</body>
```

另:  jQuery中还有一些筛选方法，类似DOM中的通过一个节点找另外一个节点，父、子、兄以外有所加强。

### 筛选方法

| 语法               | 用法.                          | 说明                                             |
| ------------------ | ------------------------------ | ------------------------------------------------ |
| parent()           | $("1i").parent()               | 查找父级                                         |
| children(selector) | $("ul").children("1i")         | 相当于$("ul>li")，(亲儿子)                       |
| find(selector)     | $("u1").find("1i")             | 相当于$ ("ul li"),后代选择器                     |
| siblings(selector) | $(".first").siblings("li")     | 查找兄弟节点，不包括自己                         |
| nextAll([expr])    | $(".first").nextAll()          | 当前元索后所有的同辈元素                         |
| prevtAll([expr])   | $(".last").prevAll()           | 当前元素前所有的同辈元素                         |
| hasClass(class)    | $('div').hasClass("protected") | 检查当前的元素是否含有某个特定的类，有则返回true |
| eq(index)          | $("li").eq(2)                  | 相当于$("li:eq(2)")                              |

## 常用技巧

### jQuery 设置样式

```javascript
$('div').css('属性', '值')    
```

### 隐式迭代

```javascript
// 遍历内部 DOM 元素（伪数组形式存储）的过程就叫做隐式迭代。
// 简单理解：给匹配到的所有元素进行循环遍历，执行相应的方法，而不用我们再进行循环，简化我们的操作，方便我们调用。
$('div').hide();  // 页面中所有的div全部隐藏，不用循环操作
```

```html
<body>
<div>惊喜不，意外不</div>
<div>惊喜不，意外不</div>
<div>惊喜不，意外不</div>
<div>惊喜不，意外不</div>
<ul>
    <li>相同的操作</li>
    <li>相同的操作</li>
    <li>相同的操作</li>
</ul>
<script>
    // 1. 获取四个div元素
    console.log($("div"));
    // 2. 给四个div设置背景颜色为粉色 jquery对象不能使用style
    $("div").css("background", "pink");
    // 3. 隐式迭代就是把匹配的所有元素内部进行遍历循环，给每一个元素添加css这个方法
    $("ul li").css("color", "red");
</script>
</body>

```

### 排他思想

```html
<body>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <script>    
        $(function () {        // 1. 隐式迭代        
            $("button").click(function () { 
                // 2. 当前的元素变化背景颜色            
                $(this).css("background", "pink");            
                // 3. 其余的兄弟去掉背景颜色 隐式迭代            
                $(this).siblings("button").css("background", "");        
            });
        })
    </script>
</body>
```

### 链式编程

```html
<body>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <button>快速</button>
    <script>    
        $(function () {        
            // 隐式迭代 给所有的按钮都绑定了点击事件        
            $("button").click(function () {            
                // 链式编程            
                $(this).css("color", "red").siblings().css("color", "");
            });
        })
    </script>
</body>
```

## jQuery 样式操作

jQuery中常用的样式操作有两种：css() 和 设置类样式方法

### 操作 css 方法

jQuery 可以使用 css 方法来修改简单元素样式； 也可以操作类，修改多个样式。

常用以下三种形式 : 

```javascript
// 1.参数只写属性名，则是返回属性值
var strColor = $(this).css('color');
// 2.  参数是属性名，属性值，逗号分隔，是设置一组样式，属性必须加引号，值如果是数字可以不用跟单位和引号。
$(this).css('color', 'red');
// 3.  参数可以是对象形式，方便设置多组样式。属性名和属性值用冒号隔开， 属性可以不用加引号
$(this).css({
    "color":"white",
    "font-size":"20px"
});
```

###  设置类样式方法

作用等同于以前的 classList，可以操作类样式， 注意操作类里面的参数不要加点。

常用的三种设置类样式方法：

```javascript
// 1.添加类
$("div").addClass("current");
// 2.删除类
$("div").removeClass("current");
// 3.切换类
$("div").toggleClass("current");
```

1. 设置类样式方法比较适合样式多时操作，可以弥补css()的不足。
2. 原生 JS 中 className 会覆盖元素原先里面的类名，jQuery 里面类操作只是对指定类进行操作，不影响原先的类名。

##  jQuery 效果

### 方法

| 方法          | 简介                   |
| ------------- | ---------------------- |
| show()        | 显示参数               |
| hide()        | 隐藏参数               |
| toggle()      | 切换参数               |
| slideDown()   | 上滑动                 |
| slideUp()     | 下滑动                 |
| slideToggle() | 滑动切换               |
| fadeIn()      | 淡入                   |
| fadeOut()     | 淡出                   |
| fadeToggle()  | 淡入淡出切换           |
| fadeTo()      | 淡入淡出并调整不透明度 |
| animate()     | 自定义动画             |
| stop()        | 停止动画               |
| hover()       | 事件切换               |
|               |                        |

动画或者效果一旦触发就会执行，如果多次触发，就造成多个动画或者效果排队执行。jQuery为我们提供另一个方法，可以停止动画排队：stop() ;

### 显示隐藏

显示隐藏动画，常见有三个方法：show() / hide() / toggle() ;

#### 显示  show()

```javascript
show([speed, [easing], [fn]])
```

- 参数都可以省略，无动画直接显示。

- speed: 

  三种预定速度之一的字符串（'slow', 'normal', 'fast'） 或表示动画时长的毫秒数值

- easing:  用来指定切换效果，默认是 'waing', 可选参数 'liner'。

- fn: 回调函数，在动画完成时执行的函数，每个元素执行一次。

#### 影藏  hide()

```
hide([speed, [easing], [fn]])
```

- 参数都可以省略，无动画直接显示。

- speed: 

  三种预定速度之一的字符串（'slow', 'normal', 'fast'） 或表示动画市场的毫秒数值

- easing: (Optional) 用来指定切换效果，默认是 'waing', 可选参数 'liner'。

- fn: 回调函数，在动画完成时执行的函数，每个元素执行一次。

#### 切换  toggle()

```
toggle([speed, [easing], [fn]])
```

- 参数都可以省略，无动画直接显示。

- speed: 

  三种预定速度之一的字符串（'slow', 'normal', 'fast'） 或表示动画市场的毫秒数值

- easing: (Optional) 用来指定切换效果，默认是 'waing', 可选参数 'liner'。

- fn: 回调函数，在动画完成时执行的函数，每个元素执行一次。

平时一般不带参数，直接显示影藏即可

```html
<body>
    <button>显示</button>
    <button>隐藏</button>
    <button>切换</button>
    <div></div>
    <script>
        $(function() { 
            $("button").eq(0).click(function() { 
                $("div").show(1000, function() { 
                    alert(1);
                });
            })
            $("button").eq(1).click(function() { 
                $("div").hide(1000, function() { 
                    alert(1);
                });
            })
            $("button").eq(2).click(function() {
                $("div").toggle(1000);
            })
      // 一般情况下，我们都不加参数直接显示隐藏就可以了
        });
    </script>
</body>
```

### 滑入滑出

滑入滑出动画，常见有三个方法：slideDown() / slideUp() / slideToggle() ; 

#### 下滑  slideDown

```javascript
slideDown([speed, [easing] , [fn]])
```

- 参数都可以省略。

- speed:

  三种预定速度之一的字符串( “slow”, "normal”, or“fast”)或表示动画时长的毫秒数值(如:1000)

- easing:(Optional)用来指定切换效果，默认是“swing”，可用参数“linear”。

- fn:回调函数，在动画完成时执行的函数，每个元素执行一次。

#### 上滑  slideUp

```
slideUp([speed,[easing],[fn]])
```

- 参数都可以省略。

- speed:

  三种预定速度之一的字符串(“slow”, "normal”,or“fast”)或表示动画时长的毫秒数值(如:1000)。

- easing: (Optional)用来指定切换效果，默认是“swing”，可用参数“linear”。

- fn:回调函数，在动画完成时执行的函数，每个元素执行一次。

#### 滑动切换  slideToggle

```
slideToggle([speed,[easing], [fn]])
```

- 参数都可以省略。

- speed:

  三种预定速度之一的字符串(“slow”, "normal”,or“fast”)或表示动画时长的毫秒数值(如:1000).

- easing:(Optional)用来指定切换效果，默认是“swing”，可用参数“linear”。

- fn:回调函数，在动画完成时执行的函数，每个元素执行一次。

```html
<body>    
    <button>下拉滑动</button>    
    <button>上拉滑动</button>    
    <button>切换滑动</button>    
    <div></div>    
    <script>        
        $(function() {            
            $("button").eq(0).click(function() {
                // 下滑动 slideDown()                
                $("div").slideDown();
            })            
            $("button").eq(1).click(function() {
                // 上滑动 slideUp()                
                $("div").slideUp(500);
            })            
            $("button").eq(2).click(function() {
                // 滑动切换 slideToggle()                
                $("div").slideToggle(500);
            });
        });    
    </script>
</body>
```

### 停止动画  stop()

动画或者效果一旦触发就会执行，如果多次触发，就造成多个动画或者效果排队执行。停止动画排队的方法为：stop() ; 

- stop() 方法用于停止动画或效果。
- stop() 写到动画或者效果的前面， 相当于停止结束上一次的动画。
- 总结: 每次使用动画之前，先调用 stop() ,在调用动画。

### 淡入淡出

淡入淡出动画，常见有四个方法：fadeIn() / fadeOut() / fadeToggle() / fadeTo() ; 

#### 淡入  fadeIn()

```
fadeIn([speed, [easing], [fn]])
```

- 参数都可以省略。

- speed:

  三种预定速度之一的字符串( "slow" ，"normal" , or "fast” )或表示动画时长的毫秒数值(如: 1000)。

- easing: (Optional) 用来指定切换效果，默认是“swing” ，可用参数"linear"。

- fn: 回调函数，在动画完成时执行的函数，每个元素执行- -次。

#### 淡诎  fadeOut()

```
fadeOut([speed, [easing], [fn]])
```

- 参数都可以省略。

- speed:

  三种预定速度之一的字符串( "slow" ，"normal" , or "fast" )或表示动画时长的毫秒数值(如: 1000)。

- easing: (Optional) 用来指定切换效果,默认是"swing” ，可用参数"linear”.

- fn: 回调函数,在动画完成时执行的函数，每个元素执行一次。

#### 切换效果  fadeToggle()

```
fadeToggle([apeed, [eaaing], [fn]])
```

- 参数都可以省略。

- speed:

  三种预定速度之一的字符串( "slow" ，"normal" , or "fast" )或表示动画时长的毫秒数值(如: 1000)。

- easing: (Optional) 用来指定切换效果，默认是“swing” ， 可用参数"linear”

- fn: 回调函数，在动画完成时执行的函数，每个元素执行一次。

- 

#### 调整不透明度  fadeTo()

```
fadeTo([[speed]，opacity, [easing], [fn]])
```

+ opacity 透明度必须写，取值0~1之间。

+ speed: 

  三种预定速度之-的字符串( "slow" ，"normal" , or "fast" )或表示动画时长的毫秒数值(如: 1000)。 必须写;

- easing: (Optional) 用来指定切换效果，默认是"swing” 。可用参数"linear” 
- fn: 回调函数,在动画完成时执行的函数，每个元素执行一次。

```html
<body>
    <button>淡入效果</button>
    <button>淡出效果</button>
    <button>淡入淡出切换</button>
    <button>修改透明度</button>
    <div></div>
    <script>
        $(function() {
            $("button").eq(0).click(function() {
                // 淡入 fadeIn()
                $("div").fadeIn(1000);
            })
            $("button").eq(1).click(function() {
                // 淡出 fadeOut()
                $("div").fadeOut(1000);
            })
            $("button").eq(2).click(function() { 
                // 淡入淡出切换 fadeToggle()
                $("div").fadeToggle(1000);
            });
            $("button").eq(3).click(function() { 
                //  修改透明度 fadeTo() 这个速度和透明度要必须写
                $("div").fadeTo(1000, 0.5);
            });
        });
    </script>
</body>
```

### 自定义动画

自定义动画非常强大，通过参数的传递可以模拟以上所有动画，方法为：animate() ;

```
animate (params, [speed], [easing], [fn])
```

- params:

  想要更改的样式属性，以对象形式传递，必须写。属性名可以不用带引号，如果是复合属性则需要采取驼峰命名法borderLeft。其余参数都可以省略。

- speed:

  三种预定速度之一的字符串(“slow”,"normal”, or“fast”)或表示动画时长的毫秒数值(如:1000)。

- easing: (Optional)用来指定切换效果，默认是“swing”，可用参数“linear”。

- fn:回调函数，在动画完成时执行的函数，每个元素执行一次。

```html
<body>
    <button>动起来</button>
    <div></div>
    <script>
        $(function() {
            $("button").click(function() { 
                $("div").animate({
                    left: 500,
                    top: 300,
                    opacity: .4,
                    width: 500
                }, 500);
            }) 
        })
    </script>
</body>
```

### 事件切换  hover()

jQuery中为我们添加了一个新事件 hover() ; 功能类似 css 中的伪类 :hover 。介绍如下

```javascript
hover([over,]out)     // 其中over和out为两个函数
```

- over:鼠标移到元素上要触发的函数（相当于mouseenter）
- out:鼠标移出元素要触发的函数（相当于mouseleave）
- 如果只写一个函数，则鼠标经过和离开都会触发它

hover事件和停止动画排列案例

```html
<body>
    <ul class="nav">
        <li>
            <a href="#">微博</a>
            <ul>
                <li><a href="">私信</a></li>
                <li><a href="">评论</a></li>
                <li><a href="">@我</a></li>
            </ul>
        </li>
        <li>
            <a href="#">微博</a>
            <ul>
                <li><a href="">私信</a></li>
                <li><a href="">评论</a></li>
                <li><a href="">@我</a></li>
            </ul>
        </li>
    </ul>
    <script>
        $(function() {
            // 鼠标经过
            $(".nav>li").mouseover(function() {
                $(this) // jQuery 当前元素  this不要加引号
                //show() 显示元素  hide() 隐藏元素    
                $(this).children("ul").slideDown(200);         
            });
            //鼠标离开
            $(".nav>li").mouseout(function() {   
                $(this).children("ul").slideUp(200);         
            });
            // 1. 事件切换 hover 就是鼠标经过和离开的复合写法
            $(".nav>li").hover(function() {   
                $(this).children("ul").slideDown(200); 
            }, function() {   
                $(this).children("ul").slideUp(200);
            });            
            // 2. 事件切换 hover  如果只写一个函数，那么鼠标经过和鼠标离开都会触发这个函数
            $(".nav>li").hover(function() {
                // stop 方法必须写到动画的前面
                $(this).children("ul").stop().slideToggle();            });
        })
    </script>
</body>
```

### 案例：王者荣耀手风琴效果

https://pvp.qq.com/web201605/herolist.shtml

> 思路分析: 
> 1.鼠标经过某个小li 有两步操作：
> 2.当前小li 宽度变为 224px， 同时里面的小图片淡出，大图片淡入
> 3.其余兄弟小li宽度变为69px， 小图片淡入， 大图片淡出

```css
<style type="text/css">    
    * {
        margin: 0;
        padding: 0;
    }    
    
    img {
        display: block;
    }    
    
    ul {
        list-style: none;
    }    
    
    .king {
        width: 852px;
        margin: 100px auto;
        overflow: hidden;
        padding: 10px;
    }   
    
    .king ul {
        overflow: hidden;
    }    
    
    .king li {
        position: relative;
        float: left;
        width: 69px;
        height: 69px;
        margin-right: 10px;
    }    
    
    .king li.current {
        width: 224px;
    }    
    
    .king li.current .big {
        display: block;
    }    
    
    .king li.current .small {
        display: none;
    }    
    
    .big {
        width: 224px;
        display: none;
    }
    
    .small {
        position: absolute;
        top: 0;
        left: 0;
        width: 69px;
        height: 69px;
        border-radius: 5px;
    }
</style>
```

```html
<div class="king">
    <ul>
        <li class="current">
            <a href="#">
                <img src="images/129.jpg" alt="" class="small">
                <img src="images/129-freehover.png" alt="" class="big">
            </a>
        </li>
        <li>
            <a href="#">
                <img src="images/152.jpg" alt="" class="small">
                <img src="images/152-freehover.png" alt="" class="big">
            </a>
        </li>
        <li>
            <a href="#">
                <img src="images/194.jpg" alt="" class="small">
                <img src="images/194-freehover.png" alt="" class="big">
            </a>
        </li>
        <li>
            <a href="#">
                <img src="images/195.jpg" alt="" class="small">
                <img src="images/195-freehover.png" alt="" class="big">
            </a>
        </li>
        <li>
            <a href="#">
                <img src="images/508.jpg" alt="" class="small">
                <img src="images/508-freehover.png" alt="" class="big">
            </a>
        </li>
    </ul>
</div>
```

```html
<script type="text/javascript">
    $(function () {
        $(".king li").mouseenter(function () {
            /*当前小li 宽度变为 224px，同时里面的小图片淡出，大图片淡入 */
            $(this).stop().animate({
                width: 224
            }).find(".small").stop().fadeOut().siblings(".big").stop().fadeIn();
            /*其余兄弟小li宽度变为69px，小图片淡入，大图片淡出*/ 
            $(this).siblings("li").stop().animate({
                width: 69  }).find(".small").stop().fadeIn().siblings(".big").stop().fadeOut();
        })});
```

## jQuery 属性操作

### 方法

| 方法   | 简介           |
| ------ | -------------- |
| prop() | 元素固有属性值 |
| attr() | 自定义属性值   |
| data() | 数据缓存       |

- 能够操作  jQuery 属性 
- 能够操作  jQuery 元素
- 能够操作  jQuery 元素尺寸、位置

### 固有属性值  prop()

所谓元素固有属性就是元素本身自带的属性，比如 <a> 元素里面的 href ，比如 <input> 元素里面的 type。 

```javascript
// 1.获取属性语法
prop("属性")

// 设置属性语法
prop("属性",”属性值")
```

prop() 除了普通属性操作，更适合操作表单属性：disabled / checked / selected 等。

### 自定义属性值 attr()

用户自己给元素添加的属性，我们称为自定义属性。 比如给 div 添加 index =“1”。 

```JavaScript
// 获取属性语法
attr("属性")  //类似原生getAttribute()

// 设置属性语法
attr("属性"，“属性值") // 类似原生setAttribute()
```

attr() 除了普通属性操作，更适合操作自定义属性。（该方法也可以获取 H5 自定义属性）

### 数据缓存 data()

data() 方法可以在指定的元素上存取数据，并不会修改 DOM 元素结构。一旦页面刷新，之前存放的数据都将被移除。 

```JavaScript
// 附加数据语法
data("name","value") // 向被选元素附加数据

// 获取数据语法
date("name") //向被选元素获取数据
```

同时，还可以读取 HTML5 自定义属性  data-index ，得到的是数字型。

```html
<body>
    <a href="http://www.baidu.com" title="都挺好">都挺好</a>
    <input type="checkbox" name="input" checked="checked">
    <div index="1" data-index="2">我是div</div>
    <span>123</span>
    <script>
        $(function () {
            //1. element.prop("属性名") 获取元素固有的属性值
            console.log($("a").prop("href"));
            $("a").prop("title", "我们都挺好");
            $("input").change(function () {
                console.log("input checked: " + $(this).prop("checked"));
            });
            // console.log($("div").prop("index"));
            // 2. 元素的自定义属性 我们通过 attr()
            console.log("index: " + $("div").attr("index"));
            $("div").attr("index", 4);
            console.log("data-index: " + $("div").attr("data-index"));
            // 3. 数据缓存 data() 这个里面的数据是存放在元素的内存里面
            $("span").data("uname", "andy");
            console.log("data-uname: " + $("span").data("uname"));
            // 这个方法获取data-index h5自定义属性 第一个 不用写data-  而且返回的是数字型
            console.log('data-index: ' + $("div").data("index"));
        })
    </script>
</body>
```

## jQuery 文本属性值

### 方法

jQuery的文本属性值常见操作有三种：html() / text() / val() ; 分别对应JS中的 innerHTML 、innerText 和 value 属性。主要针对元素的内容还有表单的值操作。

| 方法   | 简介                         |
| ------ | ---------------------------- |
| html() | 获取/设置普通元素内容        |
| text() | 获取/设置普通元素文本内容    |
| val()  | 获取/设置表单的值val("内容") |

html() 可识别标签，text() 不识别标签。

### 获取普通元素内容  html() 

相当于原生inner HTML

```javascript
html() //获取元素的内容html("内容") // 设置元素的内容
```

### 获取普通元素文本内容  text()

相当与原生 innerText

```javascript
text()   //获取元素的文本内容text("文本内容") // 设置元素的文本内容
```

### 获取表单的值  val()  

相当于原生value

```javascript
val()  //获取表单的值val("内容") // 设置表单的值
```

```html
<body>
    <div>
        <span>我是内容</span>
    </div>
    <input type="text" value="请输入内容">
    <script>
        // 1. 获取设置元素内容 html()
        console.log($("div").html());
        // $("div").html("123");
        // 2. 获取设置元素文本内容 text()
        console.log($("div").text());
        $("div").text("123");
        // 3. 获取设置表单值 val()
        console.log($("input").val());
        $("input").val("123");
    </script>
</body>
```

## jQuery 元素操作

jQuery 元素操作主要是操作标签的遍历、创建、添加、删除等操作。

### 方法

| 方法                      | 简介                         |
| ------------------------- | ---------------------------- |
| each()                    | 遍历元素                     |
| append("内容")            | 内部添加元素                 |
| element.prepend("内容' ") | 把内容放入匹配元素内部最后面 |
| element.after("内容")     | 把内容放入目标元素后面       |
| element.before("内容")    | 把内容放入目标元素前面       |

### 遍历元素   each()

jQuery 隐式迭代是对同一类元素做了同样的操作。 如果想要给同一类元素做不同操作，就需要用到遍历。


```
$("div").each(function(index, domE1e) { xxx; })
```

- each() 方法遍历匹配的每一个元素。 主要用DOM处理。each 每一个
- 里面的回调函数有2个参数: index 是每个元素的索引号; demEle 是每个DOM元素对象，不是jquery对象
- 所以要想使用jquery方法，需要给这个dom元素转换为jquery对象$(domEle)

注意：此方法用于遍历 jQuery 对象中的每一项，回调函数中元素为 DOM 对象，想要使用 jQuery 方法需要转换。

```
$.each(object,function(index,element){xx;})
```

- $.each0方法可用于遍历任何对象。主要用于数据处理，比如数组，对象
- 里面的函数有2个参数:index是每个元素的索引号;element遍历内容

注意：此方法用于遍历 jQuery 对象中的每一项，回调函数中元素为 DOM 对象，想要使用 jQuery 方法需要转换。

**演示代码**

```html
<body>
    <div>1</div>
    <div>2</div>
    <div>3</div>
    <script>
        $(function() {
            // 如果针对于同一类元素做不同操作，需要用到遍历元素（类似for，但是比for强大）
            var sum = 0;
            var arr = ["red", "green", "blue"]; 
            // 1. each() 方法遍历元素
            $("div").each(function(i, domEle) {
                // 回调函数第一个参数一定是索引号  可以自己指定索引号号名称
                // console.log(i);
                // 回调函数第二个参数一定是 dom 元素对象，也是自己命名
                // console.log(domEle);
                // 使用jQuery方法需要转换
                $(domEle)
                $(domEle).css("color", arr[i]);
                sum += parseInt($(domEle).text());
            })
            console.log(sum);
            // 2. $.each() 方法遍历元素 主要用于遍历数据，处理数据
            // $.each($("div"), function(i, ele) {
            //     console.log(i);
            //     console.log(ele);
            // });
            // $.each(arr, function(i, ele) {
            //     console.log(i);
            //     console.log(ele);
            // })
            $.each({
                name: "andy",
                age: 18
            }, function(i, ele) {
                console.log(i);
                // 输出的是 name age 属性名
                console.log(ele);
                // 输出的是 andy  18 属性值
            })
        })
    </script>
</body>
```

### 案例：购物车案例模块-计算总计和总额

1. 把所有文本框中的值相加就是总额数量，总计同理。
2. 文本框里面的值不同，如果想要相加需要用 each() 遍历，声明一个变量做计数器，累加即可。

```javascript
/*4. 用户修改文本框的值 计算 小计模块  */
$(".itxt").change(function () {
// 先得到文本框的里面的值 乘以 当前商品的单价
    var n = $(this).val();
    // 当前商品的单价
    var p = $(this).parents(".p-num").siblings(".p-price").html();
    // console.log(p);
    p = p.substr(1);
    $(this).parents(".p-num").siblings(".p-sum").html("￥" + (p * n).toFixed(2));
    getSum();
});
/*5. 计算总计和总额模块*/
getSum();
function getSum() { 
    var count = 0; // 计算总件数
    var money = 0; // 计算总价钱
    $(".itxt").each(function (i, ele) {
        count += parseInt($(ele).val());
    });
    $(".amount-sum em").text(count);
    $(".p-sum").each(function (i, ele) {
        money += parseFloat($(ele).text().substr(1));
    });
    $(".price-sum em").text("￥" + money.toFixed(2));
}
```

### 创建、添加、删除

jQuery方法操作元素的创建、添加、删除方法很多，则重点使用部分，如下：

**语法总和**

+ **1. 创建**

  ```
  $("<li></li>");
  ```

动态的创建了一个 `<li>` 

+ 2.1. 内部添加

  ```
  element.append("内容")
  ```

把内容放入匹配元素内部最后面，类似原生appendChild.

```
  element.prepend("内容' ")
```

  把内容放入匹配元素内部最前面。

+ 2.2 外部添加

  ```
  element.after("内容") //把内容放入目标元素后面element.before("内容") // 把内容放入目标元素前面
  ```

① 内部添加元素，生成之后，它们是父子关系。

② 外部添加元素，生成之后，他们是兄弟关系。

+ 3 删除元素

  ```javascript
  element.remove() // 删除匹配的元素(本身)
  element.empty() // 删除匹配的元素集合中所有的子节点
  element.html("") // 清空匹配的元素内容
  ```

①remove 删除元素本身。

②empt() 和html("") 作用等价，都可以删除元素里面的内容，只不过html还可以设置内容。



注意：以上只是元素的创建、添加、删除方法的常用方法，其他方法请参详API。

**案例代码**

```html
<body>    <ul>        <li>原先的li</li>    </ul>    <div class="test">我是原先的div</div>    <script>        $(function() {            // 1. 创建元素            var li = $("<li>我是后来创建的li</li>");                  // 2. 添加元素            // 	2.1 内部添加            // $("ul").append(li);  内部添加并且放到内容的最后面             $("ul").prepend(li); // 内部添加并且放到内容的最前面            //  2.2 外部添加            var div = $("<div>我是后妈生的</div>");            // $(".test").after(div);            $(".test").before(div);                  // 3. 删除元素            // $("ul").remove(); 可以删除匹配的元素 自杀            // $("ul").empty(); // 可以删除匹配的元素里面的子节点 孩子            $("ul").html(""); // 可以删除匹配的元素里面的子节点 孩子        })    </script></body>
```

### 案例：购物车案例模块-删除商品模块

1. 核心思路：把商品remove() 删除元素即可
2. 有三个地方需要删除： 1. 商品后面的删除按钮 2. 删除选中的商品 3. 清理购物车
3. 商品后面的删除按钮： 一定是删除当前的商品，所以从 $(this) 出发
4. 删除选中的商品： 先判断小的复选框按钮是否选中状态，如果是选中，则删除对应的商品
5. 清理购物车： 则是把所有的商品全部删掉



### 案例：购物车案例模块-选中商品添加背景

1. 核心思路：选中的商品添加背景，不选中移除背景即可
2. 全选按钮点击：如果全选是选中的，则所有的商品添加背景，否则移除背景
3. 小的复选框点击： 如果是选中状态，则当前商品添加背景，否则移除背景
4. 这个背景，可以通过类名修改，添加类和删除类



##  jQuery 尺寸、位置操作

jQuery中分别为我们提供了两套快速获取和设置元素尺寸和位置的API，方便易用，内容如下。

### jQuery 尺寸操作

jQuery 尺寸操作包括元素宽高的获取和设置，且不一样的API对应不一样的盒子模型。

**语法**

| 语法                                | 用法                                               |
| ----------------------------------- | -------------------------------------------------- |
| width()/ height()                   | 取得匹配元素宽度和高度值只算width / height         |
| innerWidth()/ innerHieght()         | 取得匹配元素宽度和高度值包含padding                |
| outerWidth()/ outerHeight()         | 取得匹配元素宽度和高度值包含padding. border        |
| outerWidth(true)/ outerHeight(true) | 取得匹配元素宽度和高度值包含padding. borde、margin |

+ 以上参数为空，则是获取相应值，返回的是数字型。
+ 如果参数为数字,则是修改相应值。
+ 参数可以不必写单位。

**代码演示**

```html
<body>
    <div></div>
    <script>
        $(function() {
            // 1. width() / height() 获取设置元素 width和height大小
            console.log($("div").width());
            // $("div").width(300);
            // 2. innerWidth() / innerHeight()  获取设置元素 width和height + padding 大小
            console.log($("div").innerWidth());
            // 3. outerWidth()  / outerHeight()  获取设置元素 width和height + padding + border 大小
            console.log($("div").outerWidth());
            // 4. outerWidth(true) / outerHeight(true) 获取设置 width和height + padding + border + margin
            console.log($("div").outerWidth(true));
        })
    </script>
</body>
```

注意：有了这套 API 我们将可以快速获取和子的宽高，至于其他属性想要获取和设置，还要使用 css() 等方法配合。

### jQuery 位置操作

jQuery的位置操作主要有三个： offset()、position()、scrollTop()/scrollLeft() , 具体介绍如下: 

**语法**

1. offset设置或获取元素偏移

   ⑴ offset方法设置或返回被选元素相对于文档的偏移坐标，跟父级没有关系。

   ⑵ 该方法有2个属性 left、top 。offset().top 用于获取距离文档顶部的距离，offset().left 用于获取距离文档左侧的距离。

   ⑶ 可以设置元素的偏移: offset( {top:10, left: 30} );

2. position 获取元素偏移

   ⑴ position()方法用于返回被选元素相对于带有定位的父级偏移坐标，如果父级都没有定位，则以文档为准。 

   ⑵ 该方法有2个属性left、top。position().top用于获取距离定位父级顶部的距离，position().left用于获取距离定 位父级左侧的距离。

   ⑶ 该方法只能获取。

3. scrollTop()/scrollLeft()设置或获取元素被卷去的头部和左侧

   ⑴ scrollTop()方法设置或返回被选元素被卷去的头部。

   ⑵ 不跟参数是获取，参数为不带单位的数字则是设置被卷去的头部。

**代码演示**

```html
<body>    <div class="father">        <div class="son"></div>    </div>            <div class="back">返回顶部</div>    <div class="container"></div>       <script>        $(function() {            // 1. 获取设置距离文档的位置（偏移） offset            console.log($(".son").offset());            console.log($(".son").offset().top);            // $(".son").offset({            //     top: 200,            //     left: 200            // });                  // 2. 获取距离带有定位父级位置（偏移） position   如果没有带有定位的父级，则以文档为准            // 这个方法只能获取不能设置偏移            console.log($(".son").position());            // $(".son").position({            //     top: 200,            //     left: 200            // });            		// 3. 被卷去的头部      		$(document).scrollTop(100);            // 被卷去的头部 scrollTop()  / 被卷去的左侧 scrollLeft()            // 页面滚动事件            var boxTop = $(".container").offset().top;            $(window).scroll(function() {                // console.log(11);                console.log($(document).scrollTop());                if ($(document).scrollTop() >= boxTop) {                    $(".back").fadeIn();                } else {                    $(".back").fadeOut();                }            });            // 返回顶部            $(".back").click(function() {                // $(document).scrollTop(0);                $("body, html").stop().animate({                    scrollTop: 0                });                // $(document).stop().animate({                //     scrollTop: 0                // }); 不能是文档而是 html和body元素做动画            })        })    </script></body>
```

### 案例：带有动画的返回顶部

1. 核心原理： 使用animate动画返回顶部。
2. animate动画函数里面有个scrollTop 属性，可以设置位置
3. 但是是元素做动画，因此 $(“body,html”).animate({scrollTop: 0})

```html
    <style>        li {LIST-STYLE-TYPE: none;}        body {height: 2000px;}        .nav {width: 900px;overflow: hidden;height: 400px;background: cyan;margin: 0 auto;}        .nav ul {background: #88C6E5;overflow: hidden;font-size: 16px;}        .nav li {float: left;width: 88px;text-align: center;margin-right: 1px;line-height: 42px;}        .nav li a {display: block;color: #ffffff;text-decoration: none;}        .nav li a:hover {background: #68ACFA;}        .back {position: fixed;width: 50px;height: 50px;background-color: pink;right: 30px;bottom: 100px;display: none;}        .container {width: 900px;background-color: skyblue;margin: 0 auto;font-size: 24px;}        p {line-height: 32px;}    </style>
```

```html
<div class="back">返回顶部</div><div class="nav">    <ul>        <li class="this"><a href="#" title="书趣阁_笔趣阁">首页</a></li>        <li><a href="#">玄幻</a></li>        <li><a href="#">武侠</a></li>        <li><a href="#">都市</a></li>        <li><a href="#">历史</a></li>        <li><a href="#">侦探</a></li>        <li><a href="#">排行榜</a></li>        <li><a href="#">全本</a></li>        <li><a href="#">书架</a></li>    </ul></div><div class="container">    <p>1</p>    <p>2</p>    <p>3</p>    <p>4</p>    <p>5</p>    <p>6</p>    <p>7</p>    <p>8</p>    <p>9</p>    <p>10</p>    <p>11</p>    <p>12</p>    <p>13</p>    <p>14</p>    <p>15</p>    <p>16</p>    <p>17</p>    <p>18</p>    <p>19</p>    <p>20</p>    <p>21</p>    <p>22</p>    <p>23</p>    <p>24</p>    <p>25</p>    <p>26</p>    <p>27</p>    <p>28</p>    <p>29</p></div>
```

```html
<script>    $(function () {        var boxTop = $(".container").offset().top;        $(window).scroll(function () {            // console.log(11);            console.log($(document).scrollTop());            if ($(document).scrollTop() >= boxTop) {                $(".back").fadeIn();            } else {                $(".back").fadeOut();            }        });        // 返回顶部        $(".back").click(function () {            // $(document).scrollTop(0);            $("body, html").stop().animate({                scrollTop: 0            });        })    })</script>
```



## 事件

### jQuery 事件注册

jQuery 为我们提供了方便的事件注册机制，是开发人员抑郁操作优缺点如下：

- 优点: 操作简单，且不用担心事件覆盖等问题。
- 缺点: 普通的事件注册不能做事件委托，且无法实现事件解绑，需要借助其他方法。

**语法** 

```javascript
element.事件(function() {})$("div").click(function(){事件处理程序 })
```

其他事件和原生基本一致

比如 mouseover、mouseout、 blur、 focus、change、 keydown、 keyup、resize、scroll 等

**演示代码** 

```html
<body>
    <div></div>
    <script>
        $(function() {
            // 1. 单个事件注册
            $("div").click(function() {
                $(this).css("background", "purple");
            });
            $("div").mouseenter(function() {
                $(this).css("background", "skyblue");
            });
        })
    </script>
</body>
```

### jQuery 事件处理

因为普通注册事件方法的不足，jQuery又开发了多个处理方法，重点讲解如下：

- on(): 用于事件绑定，目前最好用的事件绑定方法
- off(): 事件解绑
- trigger() / triggerHandler(): 事件触发

#### 事件处理 on() 绑定事件

因为普通注册事件方法的不足，jQuery又创建了多个新的事件绑定方法bind() / live() / delegate() / on()等，其中最好用的是: on()

**语法**

+ on()方法优势1:

  可以绑定多个事件，多个处理事件处理

  ```javascript
  $("div").on({
      mouseover: function(){},
      mouseout: function(){},
      click: functioni (){}});
  ```

如果事件处理程序相同

```
  $("div").on("mouseover mouseout", function () {      $(this).toggleClass("current");  });
```

+ on()方法优势2:

  可以事件委派操作。事件委派定义是，把原来加给子元素身上的事件绑定在父元素身上，就是把事件委派给父元素。

  ```
  $('ul').on("click", "li"， function() {    alert ( 'he1lo world!') ;});
  ```

  在此之前有bind(), live(),delegate()等方法来处理事件绑定或者事件委派，最新版本的请用on替代他们。

+ on()方法优势3:

  动态创建的元素，click() 没有办法绑定事件。on()可以给动态生成的元素绑定事件

  ```
  $("div" ).on("click"，"P", function(){ 	alert("俺可以给动态生成的元素绑定事件")}):
  ```

```
  $("div").append($("<p>我是动态创建的p</p>"));
```

**演示代码** 

```html
<body>    <div></div>    <ul>        <li>我们都是好孩子</li>        <li>我们都是好孩子</li>        <li>我们都是好孩子</li>    </ul>    <ol></ol>    <script>        $(function() {            // (1) on可以绑定1个或者多个事件处理程序            // $("div").on({            //     mouseenter: function() {            //         $(this).css("background", "skyblue");            //     },            //     click: function() {            //         $(this).css("background", "purple");            //     }            // });            $("div").on("mouseenter mouseleave", function() {                $(this).toggleClass("current");            });              // (2) on可以实现事件委托（委派）            // click 是绑定在ul 身上的，但是 触发的对象是 ul 里面的小li            // $("ul li").click();            $("ul").on("click", "li", function() {                alert(11);            });            // (3) on可以给未来动态创建的元素绑定事件            $("ol").on("click", "li", function() {                alert(11);            })            var li = $("<li>我是后来创建的</li>");            $("ol").append(li);        })    </script></body>
```

#### 案例：发布微博案例

1. 点击发布按钮， 动态创建一个小li，放入文本框的内容和删除按钮， 并且添加到ul 中。
2. 点击的删除按钮，可以删除当前的微博留言。

```css
* {margin: 0;padding: 0}ul {list-style: none}.box {width: 600px;margin: 100px auto;border: 1px solid #000;padding: 20px;}textarea {width: 450px;height: 160px;outline: none;resize: none;}ul {width: 450px;padding-left: 80px;}ul li {line-height: 25px;border-bottom: 1px dashed #cccccc;display: none;}input {float: right;}ul li a {float: right;}
```



```html
<body><div class="box" id="weibo"><span>微博发布</span>     <textarea name="" class="txt" cols="30" rows="10"></textarea>    <button class="btn">发布</button>    <ul></ul></div></body>
```



```html
<script>    $(function () {        // 1.点击发布按钮， 动态创建一个小li，放入文本框的内容和删除按钮， 并且添加到ul 中                $(".btn").on("click", function () {            var li = $("<li></li>");            li.html($(".txt").val() + "<a href='javascript:;'> 删除</a>");            $("ul").prepend(li);            li.slideDown();            $(".txt").val("");        });        // 2.点击的删除按钮，可以删除当前的微博留言li                // $("ul a").click(function() {          // 此时的click不能给动态创建的a添加事件                //     alert(11);                // })              // on可以给动态创建的元素绑定事件                    $("ul").on("click", "a", function () {            $(this).parent().slideUp(function () {                $(this).remove();            });        });    })</script>
```



#### 事件处理 off() 解绑事件

当某个事件上面的逻辑，在特定需求下不需要的时候，可以把该事件上的逻辑移除，这个过程我们称为事件解绑。jQuery 为我们提供 了多种事件解绑方法：die() / undelegate() / off() 等，甚至还有只触发一次的事件绑定方法 one()，在这里我们重点讲解一下 off() ;

**语法**

off(方法可以移除通过on0方法添加的事件处理程序。

```javascript
$("p").off() // 解绑p元素所有事件处理程序
$("p").off("click") // 解绑p元素上面的点击事件后面的foo是监听函数名$("ul").off("click", "li"); // 解绑事件委托
```

如果有的事件只想触发一次，可以使用one()来绑定事件。

**演示代码**

```html
<body>    <div></div>    <ul>        <li>我们都是好孩子</li>        <li>我们都是好孩子</li>        <li>我们都是好孩子</li>    </ul>    <p>我是一个P标签</p>	<script>        $(function() {  			// 事件绑定            $("div").on({                click: function() {                    console.log("我点击了");                },                mouseover: function() {                    console.log('我鼠标经过了');                }            });            $("ul").on("click", "li", function() {                alert(11);            });              // 1. 事件解绑 off             // $("div").off();  // 这个是解除了div身上的所有事件            $("div").off("click"); // 这个是解除了div身上的点击事件            $("ul").off("click", "li");              // 2. one() 但是它只能触发事件一次            $("p").one("click", function() {                alert(11);            })        })    </script></body>
```

#### 事件处理 trigger() 自动触发事件

有些时候，在某些特定的条件下，我们希望某些事件能够自动触发, 比如轮播图自动播放功能跟点击右侧按钮一致。可以利用定时器自动触发右侧按钮点击事件，不必鼠标点击触发。由此 jQuery 为我们提供了两个自动触发事件 trigger() 和 triggerHandler() ; 

**语法**

+ 第一种: trigger()

  ```
  element.click() // 第一种简写形式element.trigger("type") //第二种自动触发模式
  ```

+ 第二种: triggerHandler()

  ```
  element.triggerHandler(type) // 第三种自动触发模式
  ```

triggerHandler模式不会触发元素的默认行为，这是和前面两种的区别。

**演示代码**

```html
<body>    <div></div>    <input type="text">          <script>    $(function() {      // 绑定事件      $("div").on("click", function() {        alert(11);      });      // 自动触发事件      // 1. 元素.事件()      // $("div").click();会触发元素的默认行为            // 2. 元素.trigger("事件")      // $("div").trigger("click");会触发元素的默认行为      $("input").trigger("focus");            // 3. 元素.triggerHandler("事件") 就是不会触发元素的默认行为      $("input").on("focus", function() {        $(this).val("你好吗");      });      // 一个会获取焦点，一个不会      $("div").triggerHandler("click");      // $("input").triggerHandler("focus");    });    </script></body>
```

### jQuery 事件对象

jQuery 对DOM中的事件对象 event 进行了封装，兼容性更好，获取更方便，使用变化不大。事件被触发，就会有事件对象的产生。

**语法**

```
element.on(events, [selector], function(event) {})
```

阻止默认行为: event.preventDefault() 或者return false

阻止冒泡: event.stopPropagation()

**演示代码**

```html
<body>    <div></div>	<script>        $(function() {            $(document).on("click", function() {                console.log("点击了document");            })            $("div").on("click", function(event) {                // console.log(event);                console.log("点击了div");                event.stopPropagation();            })        })    </script></body>
```

注意：jQuery中的 event 对象使用，可以借鉴 API 和 DOM 中的 event 。

### jQuery 拷贝对象

jQuery中分别为我们提供了两套快速获取和设置元素尺寸和位置的API，方便易用，内容如下。

**语法**

```
$.extend([deep]，target, object1， [objectN] )
```

1. deep:如果设为true为深拷贝，默认为false 浅拷贝
2. target:要拷贝的目标对象
3. object1:待拷贝到第一个对象的对象。
4. objectN:待拷贝到第N个对象的对象。
5. 浅拷贝目标对象引用的被拷贝的对象地址，修改目标对象会影响被拷贝对象。
6. 深拷贝，前面加true，完全克隆， 修改目标对象不会影响被拷贝对象。

**演示代码** 

```html
<script>    $(function() {        // 1.合并数据        var targetObj = {};        var obj = {            id: 1,            name: "andy"        };        // $.extend(target, obj);        $.extend(targetObj, obj);        console.log(targetObj);        // 2. 会覆盖 targetObj 里面原来的数据        var targetObj = {            id: 0        };        var obj = {            id: 1,            name: "andy"        };        // $.extend(target, obj);        $.extend(targetObj, obj);        console.log(targetObj);     })</script>
```

