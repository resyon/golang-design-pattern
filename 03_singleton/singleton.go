package singleton

import "sync"

//Singleton 是单例模式类
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}


//HungrySingleton 为饿汉式, 个人认为的一种实现
type HungrySingleton struct{}

var  hungrySingleton *HungrySingleton

func init(){
	hungrySingleton = &HungrySingleton{}
}

//GetHungryInstance 返回饿汉式单例对象
func GetHungryInstance() *HungrySingleton{
	return hungrySingleton
}
