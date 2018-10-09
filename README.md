## go 设计模式

[![Build Status](https://travis-ci.org/xiaomeng79/go-design-pattern.svg?branch=master)](https://travis-ci.org/xiaomeng79/go-design-pattern) 
[![GitHub license](https://img.shields.io/github/license/xiaomeng79/go-design-pattern.svg)](https://github.com/xiaomeng79/go-design-pattern/blob/master/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/xiaomeng79/go-design-pattern.svg)](https://github.com/xiaomeng79/go-design-pattern/issues) [![GitHub stars](https://img.shields.io/github/stars/xiaomeng79/go-design-pattern.svg)](https://github.com/xiaomeng79/go-design-pattern/stargazers) ![go1.11](https://img.shields.io/badge/go-1.11-brightgreen.svg)


####  设计模式的六大原则
1、开闭原则
 
对扩展开放，对修改关闭,简而言之：使用接口和抽象类
 
2、里氏代换原则

任何基类可以出现的地方，子类一定可以出现

3、依赖倒转原则

针对接口编程，依赖于抽象而不依赖于具体

4、接口隔离原则

使用多个隔离的接口，比使用单个接口要好,降低耦合，参考io包

5、迪米特法则，又称最少知道原则

一个实体应当尽量少地与其他实体之间发生相互作用，使得系统功能模块相对独立

6、合成复用原则

尽量使用合成/聚合的方式，而不是使用继承

*总结：多使用接口，接口组合，针对接口编程*

#### 创建型模式

- single 单例模式
- abstractFactory 抽象工厂
- builder 建造者模式
- factoryMethod 工厂方法
- prototype 原型模式
- simpleFactory 简单工厂模式

#### 结构型模式

- adapter 适配器模式
- bridge 桥接模式
- composite 组合模式
- decorator 装饰器模式
- facade 外观模式
- flyweight 享元模式
- proxy 代理模式
- options 选项模式

#### 行为型模式

- chain 责任链模式
- command 命令模式
- interperter 解释器模式
- iterator 迭代器模式
- mediator 中介者模式
- memento 备忘录模式
- observer 观察者模式
- state 状态模式
- strategy 策略模式
- template 模板模式
- visitor 访问者模式

##### 参考资料

[tmrts](https://github.com/tmrts/go-patterns)

[BPing](https://github.com/BPing/golang_design_pattern)

[qibin0506](https://github.com/qibin0506/go-designpattern)

[HCLAC](https://github.com/HCLAC/DesignPattern)



