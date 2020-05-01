**设计模式**

## 单例模式

意图：使程序运行期间只存在一个实例对象
* 简单new struct
* 加锁后new struct
* 判空后再加锁，再new struct
* sync.Once原子操作，once.Do()方法里new struct