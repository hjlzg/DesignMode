**设计模式**

## 单例模式
* 简单new struct
* 加锁后new struct
* 判空后再加锁，再new struct
* sync.Once原子操作，once.Do()方法里new struct