# Design-Pattern-with-Golang

## 本项目主要目的
  - 个人学习设计模式
  - v1.0.0：尽量用一行表达一个设计模式，前期避免过于冗杂的概念晦涩难懂
## 主要参考文章
  - [mohuishou的博客](https://lailin.xyz/post/singleton.html)
  - [《深入设计模式》](https://refactoringguru.cn/design-patterns/catalog)
# 建造者模式
## 01_单例模式
- **一个类只有一个实例/只有一份数据**
- [掘金文章：golang 设计模式-单例模式](https://juejin.cn/post/7124720007447052302#heading-6)
## 02_工厂模式
- **创建类型不同的相关对象**
### 简单工厂
- **一个工厂，调用不同商品的构造函数进行构造**
- 工厂 + 抽象商品 + 商品
### 工厂方法
- **不同工厂生产不同商品**
- 抽象工厂 + 工厂 + 抽象商品 + 商品
- 生成工厂/商品的选择函数仅创建对象，使用者调用具体函数
### 抽象工厂
- **一个工厂生产一类商品**

## 03_建造者模式
- **创建参数复杂的对象**


- 通用做法：为每一个struct提供一个builder，builder的每个变量创建set函数，调用builder的各种函数进行参数修改，最后把builder的所有值赋给struct，得到对象
- [Options模式](https://mp.weixin.qq.com/s/z2w_MArNTjJfm9kbCFOOnA)
- Golang做法：是一个很通用的写法，必填参数（变量）+ 可选参数（opt...），可选参数通过传入的opt函数进行调用，否则使用默认值

## 04_原型模式
- **通过已经创建好的示例作为原型，进行复制克隆（**深拷贝**）**


- 好处：高效；安全，无须知道对象创建细节
- 深拷贝：拷贝整个对象，不共享内存
- 浅拷贝：拷贝对象指针，共享内存


- Golang：
  - go中传递都是值传递
  - 值类型的数据，默认全部都是深复制，Array、Int、String、Struct、Float，Bool
  - 引用类型的数据，默认全部都是浅复制，Slice，Map
  - 传入值参数e时候
    - res := e（深拷贝）
    - res := &e（浅拷贝）
  - 传入指针参数e时候：
    - res := *e（拷贝的是e本身的对象，深拷贝）
    - res := e（拷贝的仅仅是指针）
- Golang语法点：深拷贝除了上述*e方式，还常结合**序列化和反序列化**完成


- json相关问题
  - json.Unmarshal()第二个参数必须为指针
  - 想要json序列化/反序列化的key必须首字母大写，否则json包访问不到私有变量

# 结构性模式
## 05_代理模式
- **Golang实现关键：代理类和被代理类实现同一接口，代理类中持有被代理类对象**
## 06_桥接模式
- **用组合关系（has）代替继承关系（is）来实现，增强可扩展性**
- eg：不同机器搭配不同零件

## 07_装饰器模式
- **给原始类增强功能，均扩展同一接口，避免列出各种组合类型代码量暴增；**
- eg：披萨可以只要披萨饼，也可以搭配吐司、番茄、洋葱等等配料，要基于披萨饼进行扩展
## 08_适配器模式
- **将不兼容的接口转换为兼容的接口**
- adapter变量包含一个原本实例，内部做适配转换
- 创建一个interface，实例成不同的类，调用同一个方法使用不同的类，屏蔽底层不同的创建方法（因为内部实现适配转换）
## 09_外观模式
- **复杂的内部系统只对外提供一个易用接口**
- 
## 失血、贫血、充血、胀血
- 失血模型：只有属性的get set方法的纯数据类，所有的业务逻辑完全由Service层来完成的，由于没有dao，Service直接操作数据库，进行数据持久化。
  - model：只包含get set方法
- 贫血模型：贫血模型中包含了一些业务逻辑，但不包含依赖持久层的业务逻辑。这部分依赖于持久层的业务逻辑将会放到服务层中。
  - service ：组合服务，也叫事务服务
  - model：除包含get set方法，还包含原子服务
  - dao：数据持久化
- 充血模型：胀血模型反而是另外一种的失血模型，因为服务层消失了，领域层干了服务层的事，到头来还是什么都没变
  - service ：组合服务 也叫事务服务
  - model：除包含get set方法，还包含原子服务和数据持久化的逻辑
- 胀血模型：把和业务逻辑不想关的其他应用逻辑（如授权、事务等）都放到领域模型，胀血模型反而是另外一种的失血模型，因为服务层消失了，领域层干了服务层的事，到头来还是什么都没变

## 依赖倒置、控制反转、依赖注入、
### 依赖倒置（Dependency Inversion Principle，DIP）
- 高层模块不应该依赖低层模块，两者都应该依赖其抽象；抽象不应该依赖细节，细节应该依赖抽象。
- 我的理解：使用抽象，**面向接口编程**，使用接口方法，而不是具体实现的函数，可扩展性好
### 控制反转（Inversion of Control，IOC）
- 依赖方不需要关心被依赖对象的创建，只需要使用**传进来**的被依赖对象即可
### 依赖注入（Dependency Injection，DI）
- 面向接口编程，而不是应该面向实现编程
- 核心思想：只需要声明自己依赖的接口，而不要真正去实例化它，真正的实例会在合适的时机注入进去
- **现在的IOC实现几乎都是用依赖注入的方式。**
