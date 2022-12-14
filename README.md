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
## 10_组合模式
- **树形结构，每一层实现同一接口，使用同一方法进行统计/检索等等**
## 11_享元模式
- **复用对象中不必要多次出现的变量，减小内存开销；map存储复用对象，指针指向共享内存**
# 行为型模式
## 12_观察者模式
- **发布订阅机制：观察者提供Update方法，被观察者提供Add、Remove、Notify（调用所有Update方法）方法**
- Golang tips：
  - 函数参数是interface时，则传递值和指针均可接收
  - slice删除i位置元素方法
  - 注意：删除不要用range，range过程中len不变，但是删了元素有panic问题
  ```golang
  for i := 0; i < len(sub.observers); i++ {
		if sub.observers[i] == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
			i--
		}
	}
  ```
## 13_模版模式
- **业务逻辑的总体框架（模版）不变，内部不同实例的各个方法实现细节不同**
## 14_策略模式
- **通过函数设置不同的策略，避免if-else多重判断**
## 15_责任链模式
- **实现同一接口，添加组成链式结构，顺序执行**
- eg：Gin的中间件实现
## 16_状态模式
- **实现同一接口，每个状态不用的函数为空，进行状态转移（前进或后退到某一步），避免if-else状态转移**
- 缺点：不适合状态经常变化的场景
- 策略模式：彼此不知道相互存在
- 状态模式：彼此状态知道相互存在，进行前进或后退
## 17_迭代器模式
- **为对象提供一个迭代器（利于遍历），实现迭代器接口**
## 18_访问者模式
- **将某个操作（计算面积）作用于多个对象（正方形、长方形、圆形），_解藕操作和对象本身_**
- **被访问者提供Accept(v visitor)，调用visit函数，不管visitor是谁，增强可扩展性**
- **访问者visitor是一个接口，visitor传入具体被访问者，使用类型断言选择不同函数进行处理**
## 19_备忘录模式
- **提供一个副本/快照功能，进行内容恢复**
## 20_命令模式
- **把函数/请求封装成对象，实现同一接口，使发出请求的责任和执行请求的责任分割开**
- tips1：继承中如果子类太多的话，修改父类可能导致大量子类需要修改，不符合开闭原则
- tips2：main函数中使用chan和goroutine很好的实现了一个case，值得参考
## 21_解释器模式
- **统一实现interpreter接口，传进来字符串，封装具体分析方法，不暴露给外部**
## 22_中介者模式
- **添加中介者，使多对多的关系变成多对一再一对多的关系**
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
