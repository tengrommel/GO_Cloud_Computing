package singleton

import "sync"

// Singleton是单例模式类
type Singleton struct {}

var singletonInst *Singleton
var singletonInstLock sync.Mutex

// GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	if singletonInst == nil{
		singletonInstLock.Lock()
		if singletonInst == nil{
			singletonInst = &Singleton{}
		}
	}
	return singletonInst
}