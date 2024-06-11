# JavaScript

## javascript认识

### javascript作用及组成

- 作用
  - 表单动态校验(密码强度检测)
  - 网页特效
- 组成

![image-20200709030342715](assets\image-20200709030342715.png)

### ECMAScript 

ECMAScript是ECMA国际进行标准化的一门语言，它往往被称为JavaScript或JScript，实际上后两者是ECMAScript语言的实现和扩展

![image-20200709030400475](assets\image-20200709030400475.png)

ECMAScript：规定了JS的变成语法和基础核心知识，是所有浏览器厂商共同遵守的一套JS语法工业标准。

### DOM  文档对象模型

称DOM，是W3C组织推荐处理可扩展标记语言的标准编程接口。通过 DOM 提供的接口可以对页面上的各种元素进行操作（大小、位置、颜色等）

### BOM  浏览器对象模型

简称BOM，是指浏览器对象模型，它提供了独立于内容的、可以与浏览器窗口进行互动的对象结构。通过BOM可以操作浏览器窗口，比如弹出框、控制浏览器跳转、获取分辨率等。

## JavaScript 写法

JavaScript 有3种书写方式，行内式，内嵌式，外部

### 行内式

```html
<input type="button" value="点我试试" onclick="alert('Hello World')" />
```

- 可以将单行或少量JS代码写在时间属性中（**以on开头的属性**）

- 注意单双引号，HTML推荐双引号，JS推荐单引号
- 可读性差，html中编写大量JS代码，不方便阅读
- 引号易错，多层嵌套引号容易匹配混淆
- 特殊情况下使用，一般不使用

### 内嵌式

```html
<script>
    alert('Hello  World~!');
</script>
```

- 可以将多行JS代码写道script标签中
- 内嵌JS是学习时常用的方式

### 外部JS文件

```html
<script src="my.js"></script>
```

- 利于HTML页面代码结构化，把大段JS代码独立于HTML页面外，美观，利于复用
- 引用外部JS文件的script标签中可以不写代码
- 适合JS代码量比较大的情况

## 基本数据类型

### 注释

为了提高代码的可读性，JS与CSS一样，也提供了注释功能。JS中的注释主要有两种，分别是 单行注释 和 多行注释。

| 符号         | 说明     |
| ------------ | -------- |
| //注释内容   | 单行注释 |
| /*注释内容*/ | 多行注释 |

### 输入输出语句

#### 方法

为了方便信息的输入输出，JS中提供了一些输入输出语句，其常用的语句如下：

| 方法             | 说明                              | 归属   |
| ---------------- | --------------------------------- | ------ |
| alert(msg)       | 浏览器弹出警示框                  | 浏览器 |
| console.log(msg) | 浏览器控制台打印输出信息          | 浏览器 |
| prompt(info)     | 浏览器弹出输入框，用户可以输入    | 浏览器 |
| confirm          | 浏览器弹窗`是否`选择框,用户可选择 | 浏览器 |

#### 弹窗  alert() 

主要用来显示消息给客户,弹出的这个带有信息的小窗口被称为 **模态窗**。“modal” 意味着用户**不能与页面的其他部分**（例如点击其他按钮等）进行交互，直到他们处理完窗口。

```javascript
alert("Hello");
```

#### 控制台输出  console.log() 

控制台输出的内容，用来给程序员自己看运行时的消息

```javascript
console.log("hello word !")
```

#### 输入框  prompt  

prompt函数接收两个参数，`prompt(title, [default]);`

- title：显示给用户的文本
- 可选的第二个参数，指定input输入框的初始值

访问者可以在提示输入栏中输入一些内容，然后按“确定”键。然后我们在 result 中获取该文本。或者他们可以按取消键或按 Esc 键取消输入，然后我们得到 null 作为 result 。

```javascript
let age = prompt('How old are you?', 100); 
	alert('You are ${age} years old!); 
```

#### 是否弹窗  confirm  

confirm 函数显示一个带有 question 以及确定和取消两个按钮的模态窗口。点击确定返回 true ，点击取消返回 false 。

```javascript
let isBoss = confirm("Are you the boss?"); 
    alert( isBoss ); 
 // 如果“确定”按钮被按下，则显示 true
```

### 变量

变量是数据的“命名存储”。我们可以使用变量来保存商品、访客和其他信息。

| 方法  | 简介                                          |
| ----- | --------------------------------------------- |
| let   | 声明局部变量                                  |
| var   | 声明全局变量                                  |
| const | 使用 const 声明的变量称为“常量”，不能被修改。 |

#### var  声明全局变量

```javascript
var age; //  声明一个 名称为age 的变量     
```

#### let 声明局部变量

变量声明复制可以以单行形式也可以以多行形式，var方法与适用

```javascript
let user = 'John'; 
let age = 25; 
let message = 'Hello';

//等价于
let user = 'John', 
    age = 25, 
    message = 'Hello';

//等价于
let user = 'John', age = 25, message = 'Hello';
```

#### 变量语法拓展

| 情况                           | 说明                    | 结果      |
| ------------------------------ | ----------------------- | --------- |
| var  age ; console.log (age);  | 只声明 不赋值           | undefined |
| console.log(age)               | 不声明 不赋值  直接使用 | 报错      |
| age   = 10; console.log (age); | 不声明   只赋值         | 10        |

#### 变量命名规范

- 由字母(A-Za-z)、数字(0-9)、下划线(_)、美元符号( $ )组成，如：usrAge, num01, _name
- 严格区分大小写。var app; 和 var App; 是两个变量
- 不能 以数字开头。  18age   是错误的
- 不能 是关键字、保留字。例如：var、for、while
- 变量名必须有意义。 MMD   BBD        nl   →     age  
- 遵守驼峰命名法。首字母小写，后面单词的首字母需要大写。myFirstName

### 数据类型

#### 数据类型总览

JavaScript 中有八种基本的数据类型（译注：前七种为基本数据类型，也称为原始类型，而 object 为复杂数据类型）。

| 数据类型  | 说明                                                     |
| --------- | -------------------------------------------------------- |
| number    | 任何类型的数字：整数或浮点数，在 ±(253-1) 范围内的整数。 |
| bigint    | 任意长度的整数。                                         |
| string    | 字符串：一个字符串可以包含 0 个或多个字符                |
| boolean   | 布尔值类型，true或false                                  |
| Null      | 用于未知的值，只有一个 null 值的独立类型。               |
| undefined | 用于未定义的值， 只有一个 undefined 值的独立类型         |
| symbol    | 用于唯一的标识符。                                       |
| object    | 用于更复杂的数据结构。                                   |

#### 数字型 number

number既可保存整数，也可以保存小数(浮点数)

```javascript
let age = 21;       // 整数
let Age = 21.3747;  // 小数     
```

除了常规的数字，还包括所谓的“特殊数值（“special numeric values”）”也属于这种类型：Infinity 、 -Infinity 和 NaN。

Infinity表示无穷大，比任何一个数字都大的特殊值

```JavaScript
alert( 1 / 0 ); 
// 等价于
alert( Infinity );
```

NaN 代表一个计算错误。它是一个不正确的或者一个未定义的数学操作所得到的结果。

```JavaScript
alert( "not a number" / 2 ); 
// NaN，这样的除法是错误的
```

NaN 是粘性的。任何对 NaN 的进一步数学运算都会返回 NaN：

```JavaScript
alert( NaN + 1 ); // NaN 
alert( 3 * NaN ); // NaN 
alert( "not a number" / 2 - 1 ); // NaN
alert( NaN ** 0); // 1
```

所以，如果在数学表达式中有一个 NaN ，会被传播到最终结果（只有一个例外：`NaN ** 0`结果为 1) 

#### 字符串型  string

在 JavaScript 中，有三种包含字符串的方式：三引号，双引号，反引号，反引号是 **功能扩展** 引号。它们允许我们通过将变量和表达式包装在 ${…} 中，来将它们嵌入到字符串中：

```JavaScript
let name = "John"; 
// 嵌入一个变量 
alert( `Hello, ${name}!` ); 
// 嵌入一个表达式 
alert( `the result is ${1 + 2}` ); 
```

类似HTML里面的特殊字符，字符串中也有特殊字符，称之为转义符，转义符都是以\开头

| 转义符 | 解释说明                          |
| ------ | --------------------------------- |
| \n     | 换行符，n   是   newline   的意思 |
| \ \    | 斜杠   \                          |
| \'     | '   单引号                        |
| \"     | ”双引号                           |
| \t     | tab  缩进                         |
| \b     | 空格 ，b   是   blank  的意思     |

字符串是由若干字符组成的，这些字符的数量就是字符串的长度。通过字符串的 length 属性可以获取整个字符串的长度。

```JavaScript
var strMsg = "我是帅气多金的程序猿！";
alert(strMsg.length); // 显示 11
```

- 多个字符串之间可以使用 + 进行拼接，其拼接方式为 字符串 + 任何类型 = 拼接之后的新字符串
- 拼接前会把与字符串相加的任何类型转成字符串，再拼接成一个新的字符串
- 数值相加 ，字符相连

```JavaScript
//1.1 字符串 "相加"
alert('hello' + ' ' + 'world'); // hello world
//1.2 数值字符串 "相加"
alert('100' + '100'); // 100100
//1.3 数值字符串 + 数值
alert('11' + 12);     // 1112 
```

- 字符串拼接加强

  ```JavaScript
  // 1. 检测获取字符串的长度 length
  var str = 'my name is andy';
  console.log(str.length); // 15
  // 2. 字符串的拼接 +  只要有字符串和其他类型相拼接 最终的结果是字符串类型
  console.log('沙漠' + '骆驼'); // 字符串的 沙漠骆驼
  console.log('张三' + 18); // '张三老师18'
  console.log('张三' + true); // 张三true
  console.log(12 + 12); // 24
  console.log('12' + 12); // '1212'
  ```

  - 经常会将字符串和变量来拼接，变量可以很方便地修改里面的值
  - 变量是不能添加引号的，因为加引号的变量会变成字符串
  - 如果变量两侧都有字符串拼接，口诀“引引加加 ”，删掉数字，变量写加中间

#### 布尔型 Boolean

布尔类型有两个值：true 和 false ，其中 true 表示真（对），而 false 表示假（错）。布尔型和数字型相加的时候， true 的值为 1 ，false 的值为 0。

```javascript
console.log(true + 1);  // 2
console.log(false + 1); // 1
```

布尔值也可以作为比较的结果

```JavaScript
let isGreater = 4 > 1; 
alert( isGreater ); // true（比较的结果是 "yes"）
```

#### Undefined和Null

一个声明后没有被赋值的变量会有一个默认值undefined ( 如果进行相连或者相加时，注意结果）

 ```JavaScript
var variable;
console.log(variable); // undefined
console.log('你好' + variable);  // 你好 undefined
console.log(11 + variable);     // NaN
console.log(true + variable);   //  NaN
 ```

一个声明变量给 null 值，里面存的值为空

```JavaScript
var vari = null;
console.log('你好' + vari);  // 你好null
console.log(11 + vari);     // 11
console.log(true + vari);   //  1
```

#### object 类型和 symbol 类型

object 类型是一个特殊的类型。其他所有的数据类型都被称为“原始类型”，因为它们的值只包含一个单独的内容（字符串、数字或者其他）。相反， object 则用于储存数据集合和更复杂的实体。

symbol 类型用于创建对象的唯一标识符。

### 数据类型转换

#### 获取数据类型  typeof

typeof 可用来获取检测变量的数据类型

```JavaScript
typeof undefined // "undefined" 
typeof 0 // "number" 
typeof 10n // "bigint" 
typeof true // "boolean" 
typeof "foo" // "string" 
typeof Symbol("id") // "symbol" 
typeof Math // "object" (1) 
typeof null // "object" (2) 
typeof alert // "function" (3)
```

- Math 是一个提供数学运算的内建 object 。
- **typeof null** 的结果为 "object" 。这是官方承认的 typeof 的错误，这个问题来自于JavaScript 语言的早期阶段，并为了兼容性而保留了下来。 null 绝对不是一个 object 。 null有自己的类型，它是一个特殊值。 typeof 的行为在这里是错误的。
- **typeof alert **的结果是 "function" ，因为 alert 在 JavaScript 语言中是一个函数。在 JavaScript 语言中没有一个特别的 “function” 类型。函数隶属于 object 类型。但是 typeof 会对函数区分对待，并返回 "function" 。这也是来自于 JavaScript 语言早期的问题。从技术上讲，这种行为是不正确的，但在实际编程中却非常方便。

#### 转换为字符串

| 方法              | 简介                         |
| ----------------- | ---------------------------- |
| toString()        | 转成字符串                   |
| String() 强制转换 | 转成字符串                   |
| 加号拼接字符串    | 和字符串拼接的结果都是字符串 |

三种转换方式，更多第三种加号拼接字符串转换方式， 这一种方式也称之为隐式转换。


 #### 转换为数字型

| 方式                   | 说明                         | 案例                |
| ---------------------- | ---------------------------- | ------------------- |
| parseInt(string)函数   | 将string类型转成整数数值型   | parselnt('78')      |
| parseFloat(string)函数 | 将string类型转成浮点数数值型 | parseFloat('78.21') |
| Number() 强制转换函数  | 将string类型转换为数值型     | Number('12')        |
| js隐式转换(- * /)      | 利用算术运算隐式转换为数值型 | '12'-0              |

  - 注意 parseInt 和 parseFloat 单词的大小写
  - 隐式转换是在进行算数运算的时候，JS 自动转换了数据类型

#### 转换为布尔型

| 方式          | 说明               | 案例             |
| ------------- | ------------------ | ---------------- |
| Boolean()函数 | 其他类型转成布尔值 | Boolean('true'); |

  - 代表空、否定的值会被转换为 false  ，如 ''、0、NaN、null、undefined  
  - 其余值都会被转换为 true

```javascript
console.log(Boolean('')); // false
console.log(Boolean(0)); // false
console.log(Boolean(NaN)); // false
console.log(Boolean(null)); // false
console.log(Boolean(undefined)); // false
console.log(Boolean('小白')); // true
console.log(Boolean(12)); // true
```

## 运算符

### 算数运算符

| 运算符 | 描述         | 实例                     |
| ------ | ------------ | ------------------------ |
| +      | 加           | 10+ 20= 30               |
| -      | 减           | 10- 20=-10               |
| *      | 乘           | 10 * 20= 200             |
| /      | 除           | 10/ 20=0.5               |
| %      | 取余数(取模) | 返回除法的余数 9 % 62= 1 |

### 浮点数精度问题

浮点数值的最高精度是 17 位小数，但在进行算术计算时其精确度远远不如整数。

````JavaScript
var result = 0.1 + 0.2;    // 结果不是 0.3，而是：0.30000000000000004
console.log(0.07 * 100);   // 结果不是 7，  而是7.000000000000001
````

### 递增和递减运算符

在 JavaScript 中，递增（++）和递减（ -- ）既可以放在变量前面，也可以放在变量后面。放在变量前面时，我们可以称为前置递增（递减）运算符，放在变量后面时，我们可以称为后置递增（递减）运算符。

++num 前置递增，就是自加1，类似于 num =  num + 1，但是 ++num 写起来更简单。

````JavaScript
var  num = 10;
alert(++num + 10);   // 21
````

### 比较运算符

比较运算符（关系运算符）是两个数据进行比较时所使用的运算符，比较运算后，会返回一个布尔值（true / false）作为比较运算的结果。

| 运算符名称 | 说明                       | 案例.       | 结果. |
| ---------- | -------------------------- | ----------- | ----- |
| <          | 小于号                     | 1 <2        | true  |
| >          | 大于号                     | 1>2         | false |
| >=         | 大于等于号(大于或者等于)   | 2>=2        | true  |
| <=         | 小于等于号(小于或者等于)   | 3<= 2       | false |
| ==         | 判等号(会转型)             | 37 == 37    | true  |
| !=         | 不等号:                    | 37 != 37    | false |
| `=== !==`  | 全等要求值和数据类型都一致 | 37 === '37' | false |

**等号比较:**

| 符号 | 作用 | 用法                                   |
| ---- | ---- | -------------------------------------- |
| =    | 赋值 | 把右边给左边                           |
| ==   | 判断 | 判断两边值是否相等(注意此时有隐式转换) |
| ===  | 全等 | 判断两边的值和数据类型是否完全相同     |

```JavaScript
console.log(18 == '18');
console.log(18 === '18'); 
```

### 逻辑运算符

逻辑运算符是用来进行布尔值运算的运算符，其返回值也是布尔值。后面开发中经常用于多个条件的判断

| 逻辑运算符 | 说明                   | 案例           |
| ---------- | ---------------------- | -------------- |
| &&         | "逻辑与"，简称"与" and | true && false  |
| \|\|       | "逻辑或"，简称"或" or  | true\|\| false |
| !          | 逻辑非，简称"非" not   | ! true         |

```JavaScript
var isOk = !true;
console.log(isOk);  // false
```

**短路运算（逻辑中断）**

短路运算的原理：当有多个表达式（值）时,左边的表达式值可以确定结果时,就不再继续运算右边的表达式的值;

- 逻辑与

```JavaScript
// 如果第一个表达式的值为真，则返回表达式2
// 如果第一个表达式的值为假，则返回表达式1
console.log( 123 && 456 );        // 456
console.log( 0 && 456 );          // 0
console.log( 123 && 456&& 789 );  // 789
```

  - 逻辑或



```JavaScript
// 如果第一个表达式的值为真，则返回表达式1
// 如果第一个表达式的值为假，则返回表达式2
console.log( 123 || 456 );         //  123
 console.log( 0 ||  456 );          //  456
 console.log( 123 || 456 || 789 );  //  123
```

### 赋值运算符

用来把数据赋值给变量的运算符。


| 赋值运算符 | 说明                 | 案例                       |
| ---------- | -------------------- | -------------------------- |
| =          | 直接赋值             | var usrName = '我是值;     |
| +=、-=     | 加、减一个数后在赋值 | var age= 10; age+=5; // 15 |
| *=、/=、%= | 乘除、取模后在赋值   | varage= 2; age*=5;// 10    |

```JavaScript
var age = 10;
age += 5;  // 相当于 age = age + 5;
age -= 5;  // 相当于 age = age - 5;
age *= 10; // 相当于 age = age * 10;
```

### 运算符优先级

| 优先级 | 运算符     | 顺序          |
| ------ | ---------- | ------------- |
| 1      | 小括号     | ()            |
| 2      | 一元运算符 | ++ -- !       |
| 3      | 算数运算符 | 先*1%后+ .    |
| 4      | 关系运算符 | >>=<<=        |
| 5      | 相等运算符 | == != === !== |
| 6      | 逻辑运算符 | 先 && 后 II   |
| 7      | 赋值运算符 | =             |
| 8      | 逗号运算符 | ,             |

- 一元运算符里面的逻辑非优先级很高
- **逻辑与**比**逻辑或**优先级高

## 基本流程控制

### if 语句

#### 语法结构

```JavaScript
// 条件成立执行代码，否则什么也不做
if (条件表达式) {
    // 条件成立执行的代码语句
}
```

语句可以理解为一个行为，循环语句和分支语句就是典型的语句。一个程序由很多个语句组成，一般情况下，会分割成一个一个的语句。

#### if else语句（双分支语句）

- 语法结构

  ```js
  // 条件成立  执行 if 里面代码，否则执行else 里面的代码
  if (条件表达式) {
      // [如果] 条件成立执行的代码
  } else {
      // [否则] 执行的代码
  }
  ```

#### if else if 语句(多分支语句)

- 语法结构

  ```js
  // 适合于检查多重条件。
  if (条件表达式1) {
      语句1；
  } else if (条件表达式2)  {
      语句2；
  } else if (条件表达式3)  {
     语句3；
   ....
  } else {
      // 上述条件都不成立执行此处代码
  }
  ```

####  三元表达式

- 语法结构

  ```js
  表达式1 ? 表达式2 : 表达式3;
  ```

- 执行思路

  - 如果表达式1为 true ，则返回表达式2的值，如果表达式1为 false，则返回表达式3的值
  - 简单理解： 就类似于  if  else （双分支） 的简写

####  switch 分支流程控制

- 语法结构

  switch 语句也是多分支语句，它用于基于不同的条件来执行不同的代码。当要针对变量设置一系列的特定值的选项时，就可以使用 switch。

  ```js
  switch( 表达式 ){ 
      case value1:
          // 表达式 等于 value1 时要执行的代码
          break;
      case value2:
          // 表达式 等于 value2 时要执行的代码
          break;
      default:
          // 表达式 不等于任何一个 value 时要执行的代码
  }
  
  ```

  - switch ：开关 转换  ， case ：小例子   选项

  - 关键字 switch 后面括号内可以是表达式或值， 通常是一个变量

  - 关键字 case , 后跟一个选项的表达式或值，后面跟一个冒号

  - switch 表达式的值会与结构中的 case 的值做比较 

  - 如果存在匹配全等(===) ，则与该 case 关联的代码块会被执行，并在遇到 break 时停止，整个 switch 语句代码执行结束

  - 如果所有的 case 的值都和表达式的值不匹配，则执行 default 里的代码

    **注意： 执行case 里面的语句时，如果没有break，则继续执行下一个case里面的语句。**

- switch 语句和 if else if 语句的区别

  - 一般情况下，它们两个语句可以相互替换
  - switch...case 语句通常处理 case为比较确定值的情况， 而 if…else…语句更加灵活，常用于范围判断(大于、等于某个范围)
  - switch 语句进行条件判断后直接执行到程序的条件语句，效率更高。而if…else 语句有几种条件，就得判断多少次。
  - 当分支比较少时，if… else语句的执行效率比 switch语句高。
  - 当分支比较多时，switch语句的执行效率比较高，而且结构更清晰。 

### for 循环

语法结构

```js
for(初始化变量; 条件表达式; 操作表达式 ){
    //循环体
}
```

| 名称       | 作用                                                         |
| ---------- | ------------------------------------------------------------ |
| 初始化变量 | 通常被用于初始化一个计数器，该表达式可以使用 var 关键字声明新的变量，这个变量帮我们来记录次数。 |
| 条件表达式 | 用于确定每一次循环是否能被执行。如果结果是 true 就继续循环，否则退出循环。 |
| 操作表达式 | 用于确定每一次循环是否能被执行。如果结果是 true 就继续循环，否则退出循环。 |

执行过程：

1. 初始化变量，初始化操作在整个 for 循环只会执行一次。

   执行条件表达式，如果为true，则执行循环体语句，否则退出循环，循环结束。

   执行操作表达式，此时第一轮结束。

2. 第二轮开始，直接去执行条件表达式（不再初始化变量），如果为 true ，则去执行循环体语句，否则退出循环。

3. 继续执行操作表达式，第二轮结束。

4. 后续跟第二轮一致，直至条件表达式为假，结束整个 for 循环。

断点调试：

```
断点调试是指自己在程序的某一行设置一个断点，调试时，程序运行到这一行就会停住，然后你可以一步一步往下调试，调试过程中可以看各个变量当前的值，出错的话，调试到出错的代码行即显示错误，停下。断点调试可以帮助观察程序的运行过程
```

```html
断点调试的流程：
1、浏览器中按 F12--> sources -->找到需要调试的文件-->在程序的某一行设置断点
2、Watch: 监视，通过watch可以监视变量的值的变化，非常的常用。
3、摁下F11，程序单步执行，让程序一行一行的执行，这个时候，观察watch中变量的值的变化。
```

#### 常用方法

```javascript
# 1 方式一：最基本的
for (let i=0; i<3; i++) {
  console.log(i)
}

# 2 in 循环  es5的语法
for(let 成员 in 对象){
    循环的代码块
  
  
# 3 for of   es6的循环
  for(item of arr){
    console.log('item =>', item)
  }
  
 
# 4 数组foreach循环 (数组)
   var a=[33,22,888]
   a.forEach(function (value,index){
        console.log(value)
        console.log(index)
    })

# 5 jq  $each 循环
$.each(可迭代对象,function (key,value) {
  });
```





### while循环

while语句的语法结构如下：

```js
while (条件表达式) {
    // 循环体代码 
}
```

执行思路：

- 1 先执行条件表达式，如果结果为 true，则执行循环体代码；如果为 false，则退出循环，执行后面代码
- 2 执行循环体代码
- 3 循环体代码执行完毕后，程序会继续判断执行条件表达式，如条件仍为true，则会继续执行循环体，直到循环条件为 false 时，整个循环过程才会结束

注意：

- 使用 while 循环时一定要注意，它必须要有退出条件，否则会成为死循环

####  continue、break

continue 关键字用于立即跳出本次循环，继续下一次循环（本次循环体中 continue 之后的代码就会少执行一次）。

例如，吃5个包子，第3个有虫子，就扔掉第3个，继续吃第4个第5个包子，其代码实现如下：

```js
 for (var i = 1; i <= 5; i++) {
     if (i == 3) {
         console.log('这个包子有虫子，扔掉');
         continue; // 跳出本次循环，跳出的是第3次循环 
      }
      console.log('我正在吃第' + i + '个包子呢');
 }
```

## 函数与对象

###  数组

####  创建数组

JS 中创建数组有两种方式：

- 利用  new 创建数组  

  ```js
  var 数组名 = new Array() ；
  var arr = new Array();   // 创建一个新的空数组
  ```

  注意 Array () ，A 要大写    

- 利用数组字面量创建数组

  ```js
  //1. 使用数组字面量方式创建空的数组
  var  数组名 = []；
  //2. 使用数组字面量方式创建带初始值的数组
  var  数组名 = ['小白','小黑','大黄','瑞奇'];
  ```

  - 数组的字面量是方括号 [ ] 

- 声明数组并赋值称为数组的初始化

  - 这种字面量方式也是我们以后最多使用的方式

- 数组元素的类型

  数组中可以存放任意类型的数据，例如字符串，数字，布尔值等。

  ```js
  var arrStus = ['小白',12,true,28.9];
  ```

####  获取数组中的元素

索引 (下标) ：用来访问数组元素的序号（数组下标从 0 开始）。

数组可以通过索引来访问、设置、修改对应的数组元素，可以通过“数组名[索引]”的形式来获取数组中的元素。

```js
// 定义数组
var arrStus = [1,2,3];
// 获取数组中的第2个元素
alert(arrStus[1]);    
```

注意：如果访问时数组没有和索引值对应的元素，则得到的值是undefined

####  遍历数组

- 数组遍历

  把数组中的每个元素从头到尾都访问一次（类似学生的点名），可以通过 for 循环索引遍历数组中的每一项


```js
var arr = ['red','green', 'blue'];
for(var i = 0; i < arr.length; i++){
    console.log(arrStus[i]);
}
```

- 数组的长度

  数组的长度：默认情况下表示数组中元素的个数

  使用“数组名.length”可以访问数组元素的数量（数组长度）。

  ```js
  var arrStus = [1,2,3];
  alert(arrStus.length);  // 3
  
  
  ```

    **注意**：

  - 此处数组的长度是数组元素的个数 ，不要和数组的索引号混淆。

- 当我们数组里面的元素个数发生了变化，这个 length 属性跟着一起变化

  - 数组的length属性可以被修改：

- 如果设置的length属性值大于数组的元素个数，则会在数组末尾出现空白元素；

  - 如果设置的length属性值小于数组的元素个数，则会把超过该值的数组元素删除

####  数组中新增元素

数组中可以通过以下方式在数组的末尾插入新元素：

```js
数组[ 数组.length ] = 新数据;
```

###  函数

####  函数的使用

**声明函数**

```js
// 声明函数
function 函数名() {
    //函数体代码
}

```

- function 是声明函数的关键字,必须小写

- 由于函数一般是为了实现某个功能才定义的， 所以通常我们将函数名命名为动词，比如 getSum

**调用函数**

```js
// 调用函数
函数名();  // 通过调用函数名来执行函数体代码

```

- 调用的时候千万不要忘记添加小括号

- 口诀：函数不调用，自己不执行

  注意：声明函数本身并不会执行代码，只有调用函数时才会执行函数体代码。

**函数的封装**

函数的封装是把一个或者多个功能通过函数的方式封装起来，对外只提供一个简单的函数接口

例子：封装计算1-100累加和

```js
/* 
   计算1-100之间值的函数
*/
// 声明函数
function getSum(){
  var sumNum = 0;// 准备一个变量，保存数字和
  for (var i = 1; i <= 100; i++) {
    sumNum += i;// 把每个数值 都累加 到变量中
  }
  alert(sumNum);
}
// 调用函数
getSum();

```

####  函数的参数

| 参数 | 说明                                                       |
| ---- | ---------------------------------------------------------- |
| 形参 | 形式上的参数  函数定义的时候传递的参数  当前并不知道是什么 |
| 实参 | 实际上的参数  函数调用的时候传递的参数  实参是传递给形参的 |

参数的作用 : 在函数内部某些值不能固定，我们可以通过参数在调用函数时传递不同的值进去。

函数参数的运用：

```js
// 带参数的函数声明
function 函数名(形参1, 形参2 , 形参3...) { // 可以定义任意多的参数，用逗号分隔
  // 函数体
}
// 带参数的函数调用
函数名(实参1, 实参2, 实参3...); 

```

1. 调用的时候实参值是传递给形参的
2. 形参简单理解为：不用声明的变量
3. 实参和形参的多个参数之间用逗号（,）分隔

**形参与实参数据不匹配时**

| 参数个数             | 说明                               |
| -------------------- | ---------------------------------- |
| 实参个数等于形参个数 | 输出正确结果                       |
| 实参个数多余形参个数 | 只取形参的个数                     |
| 实参个数小于形参个数 | 多的形参定义为undefined，结果为NaN |

在JavaScript中，形参的默认值是undefined。

####  函数的返回值

返回值：函数调用整体代表的数据；函数执行完成后可以通过return语句将指定数据返回 。

```js
// 声明函数
function 函数名（）{
    ...
    return  需要返回的值；
}
// 调用函数
函数名();    // 此时调用函数就可以得到函数体内return 后面的值

```

如果函数没有 return ，返回的值是 undefined

break ,continue ,return 的区别

| 方法     | 简介                                                         |
| -------- | ------------------------------------------------------------ |
| break    | 结束当前的循环体（如 for、while                              |
| continue | 跳出本次循环，继续执行下次循环（如 for、while）              |
| return   | 不仅可以退出循环，还能够返回 return 语句中的值，同时还可以结束当前的函数体内的代码 |

####  函数的两种声明方式

- 自定义函数方式(命名函数)

  利用函数关键字 function 自定义函数方式

  ```js
  // 声明定义方式
  function fn() {...}
  // 调用  
  fn();  
  
  ```

  - 因为有名字，所以也被称为命名函数
  - 调用函数的代码既可以放到声明函数的前面，也可以放在声明函数的后面

- 函数表达式方式(匿名函数）

  利用函数表达式方式的写法如下： 

  ```js
  // 这是函数表达式写法，匿名函数后面跟分号结束
  var fn = function(){...}；
  // 调用的方式，函数调用必须写到函数体下面
  fn();
  
  ```

  - 因为函数没有名字，所以也被称为匿名函数
  - 这个fn 里面存储的是一个函数  
  - 函数表达式方式原理跟声明变量方式是一致的
  - 函数调用的代码必须写到函数体后面

###  对象

就是花括号 { } 里面包含了表达这个具体事物（对象）的属性和方法；{ } 里面采取键值对的形式表示 

- 键：相当于属性名

- 值：相当于属性值，可以是任意类型的值（数字类型、字符串类型、布尔类型，函数类型等）

  ```js
  var star = {
      name : '张三',
      age : 18,
      sex : '男',
      sayHi : function(){
          alert('大家好啊~');
      }
  };
  
  ```

上述代码中 star即是创建的对象。

#### 对象的使用

- 对象的属性

  - 对象中存储**具体数据**的 "键值对"中的 "键"称为对象的属性，即对象中存储具体数据的项

- 对象的方法

  - 对象中存储**函数**的 "键值对"中的 "键"称为对象的方法，即对象中存储函数的项

- 访问对象的属性

  - 对象里面的属性调用 : 对象.属性名 ，这个小点 . 就理解为“ 的 ”  

  - 对象里面属性的另一种调用方式 : 对象[‘属性名’]，注意方括号里面的属性必须加引号      

    示例代码如下：

    ```js
    console.log(star.name)     // 调用名字属性
    console.log(star['name'])  // 调用名字属性
    
    ```

- 调用对象的方法

  - 对象里面的方法调用：对象.方法名() ，注意这个方法名字后面一定加括号 

    示例代码如下：

    ```js
    star.sayHi();              // 调用 sayHi 方法,注意，一定不要忘记带后面的括号
    
    ```

- 变量、属性、函数、方法总结

  属性是对象的一部分，而变量不是对象的一部分，变量是单独存储数据的容器

  - 变量：单独声明赋值，单独存在
  - 属性：对象里面的变量称为属性，不需要声明，用来描述该对象的特征

+ 方法是对象的一部分，函数不是对象的一部分，函数是单独封装操作的容器
  + 函数：单独存在的，通过“函数名()”的方式就可以调用
  + 方法：对象里面的函数称为方法，方法不需要声明，使用“对象.方法名()”的方式就可以调用，方法用来描述该对象的行为和功能。 

#### 遍历对象

  for... in 语句用于对数组或者对象的属性进行循环操作。

  其语法如下：

  ```js
for (变量 in 对象名字) {
	// 在此执行代码
}

  ```

  语法中的变量是自定义的，它需要符合命名规范，通常我们会将这个变量写为 k 或者 key。

  ```js
for (var k in obj) {
    console.log(k);      // 这里的 k 是属性名
    console.log(obj[k]); // 这里的 obj[k] 是属性值
}
  ```

###  作用域

####  作用域概述

通常来说，一段程序代码中所用到的名字并不总是有效和可用的，而限定这个名字的可用性的代码范围就是这个名字的作用域。作用域的使用提高了程序逻辑的局部性，增强了程序的可靠性，减少了名字冲突。

JavaScript（es6前）中的作用域有两种：

- 全局作用域
- 局部作用域（函数作用域）	

####  全局作用域

作用于所有代码执行的环境(整个 script 标签内部)或者一个独立的 js 文件。

####  局部作用域

作用于函数内的代码环境，就是局部作用域。 因为跟函数有关系，所以也称为函数作用域。

####  JS没有块级作用域

- 块作用域由 { } 包括。

- 在其他编程语言中（如 java、c#等），在 if 语句、循环语句中创建的变量，仅仅只能在本 if 语句、本循环语句中使用，如下面的Java代码：	

  java有块级作用域：

  ```java
  if(true){
    int num = 123;
    system.out.print(num);  // 123
  }
  system.out.print(num);    // 报错
  
  
  ```

  以上java代码会报错，是因为代码中 { } 即一块作用域，其中声明的变量 num，在 “{ }” 之外不能使用；

  而与之类似的JavaScript代码，则不会报错：



  Js中没有块级作用域（在ES6之前）

  ```js
if(true){
  var num = 123;
  console.log(num); //123
}
console.log(num);   //123


  ```

###  变量的作用域

在JavaScript中，根据作用域的不同，变量可以分为两种：

- 全局变量
- 局部变量

**全局变量**

在全局作用域下声明的变量叫做全局变量（在函数外部定义的变量）。

- 全局变量在代码的任何位置都可以使用
- 在全局作用域下 var 声明的变量 是全局变量
- 特殊情况下，在函数内不使用 var 声明的变量也是全局变量（不建议使用）

**局部变量**

在局部作用域下声明的变量叫做局部变量（在函数内部定义的变量）

- 局部变量只能在该函数内部使用
- 在函数内部 var 声明的变量是局部变量
- 函数的形参实际上就是局部变量

**全局变量和局部变量的区别**

- 全局变量：在任何一个地方都可以使用，只有在浏览器关闭时才会被销毁，因此比较占内存
- 局部变量：只在函数内部使用，当其所在的代码块被执行时，会被初始化；当代码块运行结束后，就会被销毁，因此更节省内存空间。

# BOM对象

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

## history对象

history 对象表示当前窗口首次使用以来用户的导航历史记录。因为 history 是 window 的属性，所以每个 window 都有自己的 history 对象。出于安全考虑，这个对象不会暴露用户访问过的 URL，但可以通过它在不知道实际 URL 的情况下前进和后退。

```javascript
// 后退一页
history.go(-1); 
history.back(); 

// 前进一页
history.go(1);
history.forward();

// 前进两页
history.go(2);

// 导航到最近的 wrox.com 页面
history.go("wrox.com"); 

// 导航到最近的 nczonline.net 页面
history.go("nczonline.net");

// 判断窗口是否位第一个窗口，常被用于创建前进后退，以及确定页面是不是历史记录中的第一条
if (history.length == 1){ 
 // 这是用户窗口中的第一个页面
}
```

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

# Arguments对象

`arguments`是一个对应于传递给函数的参数的类数组对象。`arguments`对象是所有（非箭头）函数中都可用的**局部变量**。可以使用`arguments`对象在函数中引用函数的参数。此对象包含传递给函数的每个参数，第一个参数在索引 0 处。例如，如果一个函数传递了三个参数，可以以如下方式引用他们：

```javascript
// 引用参数
arguments[0]
arguments[1]
arguments[2]

// 设置参数
arguments[1] = 'new value';
```

`arguments`对象不是一个 Array 。它类似于`Array`，但**除了 length 属性和索引元素**外没有任何`Array`属性。例如，它没有 pop 方法。但是它可以被转换为一个真正的`Array`。

如果调用的参数多于正式声明接受的参数，则可以使用`arguments`对象。这种技术对于可以传递可变数量的参数的函数很有用。使用 [`arguments.length`](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Functions/arguments/length)来确定传递给函数参数的个数，然后使用`arguments`对象来处理每个参数。要确定函数[签名](https://developer.mozilla.org/zh-CN/docs/Glossary/Signature/Function)中（输入）参数的数量，请使用[`Function.length`](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Function/length)属性。

# 不定参数和默认参数

## 不定参数

我们通常使用可变参函数来构造API，可变参函数可接受任意数量的参数。

```javascript
function info(v1,...data){
       console.log(v1,data);
   }
info(11);
   info(11,22);
   info(11,22,333,444,55);

function info(v1,v2,v3,v4){
        console.log(v1,v2,v3,v4);
    }

    info(11,22,333,444);
    nums = [22,33,44,55,66,77,88];
    info(11,...nums)
```

## 默认参数

# 解构赋值

```html
let info = {name:"武沛齐",email:"wupeiqi@live.com",addr:"北京"};
let {name,addr} = info;
    
console.log(name);
console.log(addr);

<script>
function getData(n1,{name,addr}){
    //let {name,addr} = info; 
	console.log(name);
	console.log(addr);
}
    
let info = {name:"hsiangya",email:"hsiangyatang@gmail.com",addr:"北京"};
getData(111,info);	
</script>
```

# 模块导入导出

1. 定义模块

```javascript
// filename:index.js
const name = 'hsiangya'
let age = 19

function ShowInfo(){
    console.log(name,age);
}

// 1. 导出变量或函数
export {age,ShowInfo};
// 2. 导出对象
obj = {
    n1:age;
    n2:ShowInfo
}
export default obj;
```

2. 引用

```javascript
<!--声明模块-->
<script type="model">
    // 1. 导入变量或函数
	import * as md from "./xxx/index.js";
	console.log(md.age);
	md.ShowInfo();
	// 2. 导入对象
	import md from"./xxx/index.js";
	console.log(md.n1)
	md.n2()
</script>
```

# Fetch

## Fetch API 

Fetch API 能够执行 XMLHttpRequest 对象的所有任务，但更容易使用，接口也更现代化，能够在 Web 工作线程等现代 Web 工具中使用。XMLHttpRequest 可以选择异步，而 Fetch API 则必须是异步。

Fetch API 是 WHATWG 的一个“活标准”（living standard），用规范原文说，就是“Fetch 标准定义请求、响应，以及绑定二者的流程：获取（fetch）”。

Fetch API 本身是使用 JavaScript 请求资源的优秀工具，同时这个 API 也能够应用在服务线程（service worker）中，提供拦截、重定向和修改通过 fetch()生成的请求接口。

### 基本用法

fetch() 方法是暴露在全局作用域中的，包括主页面执行线程、模块和工作线程。调用这个方法，浏览器就会向给定 URL 发送请求。

1、分派请求

fetch()只有一个必需的参数 input。多数情况下，这个参数是要获取资源的 URL。这个方法返回

一个期约：

```javascript
let r = fetch('/bar'); 

console.log(r); // Promise <pending> 
```

URL 的格式（相对路径、绝对路径等）的解释与 XHR 对象一样。

请求完成、资源可用时，期约会解决为一个 Response 对象。这个对象是 API 的封装，可以通过它

取得相应资源。获取资源要使用这个对象的属性和方法，掌握响应的情况并将负载转换为有用的形式，

如下所示：

```javascript
  fetch('bar.txt').then((response) => {
    console.log(response);
  });

// Response { type: "basic", url: ... }
```



2、读取响应

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



3、处理状态码和请求失败

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

违反 CORS、无网络连接、HTTPS 错配及其他浏览器/网络策略问题都会导致期约被拒绝。

可以通过 url 属性检查通过 fetch()发送请求时使用的完整 URL：

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



4、自定义选项

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

### 常见 Fetch 请求模式

与 XMLHttpRequest 一样，fetch()既可以发送数据也可以接收数据。使用 init 对象参数，可以配置 fetch()在请求体中发送各种序列化的数据。

1、发送 JSON 数据

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

2、在请求体中发送参数

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

3、发送文件

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



4、加载 **Blob** 文件

Fetch API也能提供 Blob 类型的响应，而 Blob 又可以兼容多种浏览器 API。一种常见的做法是明确将图片文件加载到内存，然后将其添加到 HTML图片元素。为此，可以使用响应对象上暴露的 blob()方法。这个方法返回一个期约，解决为一个 Blob 的实例。然后，可以将这个实例传给 URL.createObjectUrl() 以生成可以添加给图片元素 src 属性的值：

```javascript
const imageElement = document.querySelector('img');

fetch('my-image.png') 
 .then((response) => response.blob()) 
 .then((blob) => { 
 imageElement.src = URL.createObjectURL(blob); 
 });
```

5、发送跨源请求

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

6、中断请求

Fetch API 支持通过 AbortController/AbortSignal 对中断请求。调用 AbortController. abort()会中断所有网络传输，特别适合希望停止传输大型负载的情况。中断进行中的 fetch()请求会导致包含错误的拒绝。

```javascript
let abortController = new AbortController(); 

fetch('wikipedia.zip', { signal: abortController.signal }) 
 .catch(() => console.log('aborted!'); 
 
// 10 毫秒后中断请求
setTimeout(() => abortController.abort(), 10); 

// 已经中断
```

### **Headers** 对象

Headers 对象是所有外发请求和入站响应头部的容器。每个外发的 Request 实例都包含一个空的Headers 实例，可以通过 Request.prototype.headers 访问，每个入站 Response 实例也可以通过Response.prototype.headers 访问包含着响应头部的 Headers 对象。这两个属性都是可修改属性。另外，使用 new Headers()也可以创建一个新实例。

1、**Headers** 与 **Map** 的相似之处

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

2、 **Headers** 独有的特性

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

3、头部护卫

某些情况下，并非所有 HTTP 头部都可以被客户端修改，而 Headers 对象使用护卫来防止不被允许的修改。不同的护卫设置会改变 set()、append()和 delete()的行为。违反护卫限制会抛出TypeError。

Headers 实例会因来源不同而展现不同的行为，它们的行为由护卫来控制。JavaScript 可以决定Headers 实例的护卫设置。下表列出了不同的护卫设置和每种设置对应的行为。

| 护 卫           | 适用情形                                                     | 限 制                                                        |
| --------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| none            | 在通过构造函数创建 Headers 实例时激活                        | 无                                                           |
| request         | 在通过构造函数初始化 Request对象，且 mode 值为非 no-cors 时激活 | 不允许修改禁止修改的头部（参见 MDN 文档中的 forbidden header name 词条） |
| request-no-cors | 在通过构造函数初始化 Request对象，且 mode值为 no-cors 时激活 | 不允许修改非简单头部（参见 MDN 文档中的simple header 词条）  |
| response        | 在通过构造函数初始化 Response 对象时激活                     | 不允许修改禁止修改的响应头部（参见 MDN 文档中的 forbidden response header name 词条） |
| immutable       | 在通过 error()或 redirect()静态方法初始化 Response 对象时激活 | 不允许修改任何头部                                           |

### **Request** 对象

顾名思义，Request 对象是获取资源请求的接口。这个接口暴露了请求的相关信息，也暴露了使用请求体的不同方式。

1、创建 **Request** 对象

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

2、克隆 **Request** 对象

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

3、在 **fetch()** 中使用 **Request** 对象

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

### **Response** 对象

顾名思义，Response 对象是获取资源响应的接口。这个接口暴露了响应的相关信息，也暴露了使用响应体的不同方式。

1、 创建 **Response** 对象

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

2、读取响应状态信息

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

3、克隆 **Response** 对象

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

### **Request**、**Response** 及 **Body** 混入

Request 和 Response 都使用了 Fetch API 的 Body 混入，以实现两者承担有效载荷的能力。这个混入为两个类型提供了只读的 body 属性（实现为 ReadableStream）、只读的 bodyUsed 布尔值（表示 body 流是否已读）和一组方法，用于从流中读取内容并将结果转换为某种 JavaScript 对象类型。

通常，将 Request 和 Response 主体作为流来使用主要有两个原因。一个原因是有效载荷的大小可能会导致网络延迟，另一个原因是流 API 本身在处理有效载荷方面是有优势的。除此之外，最好是一次性获取资源主体。

Body 混入提供了 5 个方法，用于将 ReadableStream 转存到缓冲区的内存里，将缓冲区转换为某种 JavaScript 对象类型，以及通过期约来产生结果。在解决之前，期约会等待主体流报告完成及缓冲被解析。这意味着客户端必须等待响应的资源完全加载才能访问其内容。



1、**Body.text()**

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

2、**Body.json()**

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

3、 **Body.formData()**

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

4、**Body.arrayBuffer()**

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

5、**Body.blob()**

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

