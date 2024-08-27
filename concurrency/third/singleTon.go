package main

import (
	"fmt"
	"sync"
)



type Config struct {
	Name string
}

var (	
	mx = sync.Mutex{}
	once sync.Once
	
	config *Config)

func NewConfig() *Config{
	mx.Lock()
	defer mx.Unlock()
	if config == nil{
		fmt.Println("create single")
		config = &Config{}
	}
	return config
}


func NewConfig2()*Config{
	once.Do(func() {
		fmt.Println("create")
		config = &Config{}
	})
	return config

}


func main(){
	for i := 0; i < 10; i++ {
		ii:= NewConfig()
		fmt.Printf("%p\n", ii)
	}
}