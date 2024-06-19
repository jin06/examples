Go struct 
========
1. go 结构体定义
- 用户自定义类型，包含一系列类型数据的集合。
    - go内置的数据类型：数字（int, uint, float等），byte, string, 数组，切片等
    - 结构体
    - 方法
    - 接口

2. 定义结构体
- 关键字 type， struct
```golang
 type test struct {

 }
``` 
3. 结构体字段的作用域
- 结构体作用域（是否可以在包外部直接使用，读取和更改）
    - 由结构体名字的首字母决定
- 字段是否可以在包外直接使用（读取和更改）
    - 由字段首字母决定，大写可以在包外部直接使用。其他只能在包内部直接使用。
- 如何在包外部使用私有结构体

4. 结构体的方法
- 定义方法
    - 方法接收者是引用：方法内使用的是结构体的引用，改变结构体属性值，结构体本身的属性也会改变
    - 方法接收者是结构体：方法内使用的是结构体的复制，改变结构体属性值，结构体本身的属性不会改变
    - 匿名方法接收者方式：方法内部不使用结构体的属性和其他方法

5. 面向对象问题
- 封装问题：结构体本身，接口。
- 继承问题： 属性、方法的继承
    - 属性拼接：
        - 结构体本身可以包含其他结构体
        - 一个结构体的属性包含了其他结构体，这个结构体就可以使用其他结构体的方法
- 多态问题：接口+结构体