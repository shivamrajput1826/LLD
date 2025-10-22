package creational

import (
	"fmt"
	"sync"
)

type Config struct {
	AppName string
}

var (
	instance *Config
	once     sync.Once
)

func GetInstance() *Config {
	once.Do(func() {
		fmt.Println("Initializing only once")
		instance = &Config{AppName: "My-Design-App"}
	})
	return instance

}

func Singleton() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cfg := GetInstance()
			fmt.Printf("Goroutine %d -> %p, %s\n", id, cfg, cfg.AppName)
		}(i)
	}
	wg.Wait()
}
