//creating the instance only when first requested
// Pros: Fast startup time. Thread/goroutine safe. Highly concurrent (subsequent calls bypass the lock entirely after the first creation).
// Cons: Slightly more complex than eager initialization.
package singleton

import "sync"

type SingletonLazy struct {
	a string
}

var (
	instance *SingletonLazy
	once     sync.Once
)

func GetInstanceLazy() *SingletonLazy {
	once.Do(func() {
		instance = &SingletonLazy{
			a: "Hey how are you b",
		}
	})
	return instance
}
