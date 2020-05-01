package singleTon

import "sync"

type singleTon struct {
}

var instance *singleTon

func GetInstance() *singleTon{
	if instance==nil{
		instance=&singleTon{}   //<-非线程安全的
	}
	return instance
}

var mu sync.Mutex
func GetInstance2() *singleTon{
	mu.Lock()
	defer mu.Unlock()

	if instance==nil{
		instance=&singleTon{}   // <--- 如果实例已经被创建就没有必要锁写
	}
	return instance
}

func GetInstance3() *singleTon{
	if instance==nil{     // <-- 不够完善. 他并不是完全的原子性
		mu.Lock()
		defer mu.Unlock()

		if instance==nil{
			instance=&singleTon{}
		}
	}
	return instance
}

var once sync.Once
func GetInstance4() *singleTon{
	once.Do(func(){
		instance=&singleTon{}
	})
	return instance
}