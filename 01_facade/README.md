# 外观模式

API 为facade 模块的外观接口，大部分代码使用此接口简化对facade类的访问。

facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。

> 用于隐藏细节简化访问，和`装饰器模式`(decorator)有区别， 装饰器用于动态调整对象行为， 不可单纯按字面意思理解