//pro : simple and no locks are required
//con : what if the initialization taks too much amount of time haa ??  and memoryy

package singleton

type Singleton struct {
	a string
}

var singleton *Singleton

func init() {
	singleton = &Singleton{
		a: "Hello yaar!!",
	}
}

// GetInstanmce
func GetInstance() *Singleton {
	return singleton
}
