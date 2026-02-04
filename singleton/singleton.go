package singleton

import "sync"

type Singleton struct{}

var singletonInstance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singletonInstance = &Singleton{}
	})
	return singletonInstance
}
