# Design-Pattern-with-Golang


## 01_Singleton
[掘金文章：golang 设计模式-单例模式](https://juejin.cn/post/7124720007447052302#heading-6)
## 02_Factory
[go语言中文网：Golang 工厂模式](https://studygolang.com/articles/27954)
[Go设计模式02-工厂模式&DI容器](https://lailin.xyz/post/factory.html#%E5%B7%A5%E5%8E%82%E6%96%B9%E6%B3%95)

### 简单工厂
- 一个工厂，调用不同商品的构造函数进行构造
- 工厂 + 抽象商品 + 商品
### 工厂方法
- 不同工厂生产不同商品
- 抽象工厂 + 工厂 + 抽象商品 + 商品
- 生成工厂/商品的选择函数仅创建对象，使用者调用具体函数
### 抽象工厂
- 一个工厂生产一类商品