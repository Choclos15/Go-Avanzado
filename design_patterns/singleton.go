package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

var db *Database
var lock sync.Mutex

func (Database) CreateSingleConection() {
	fmt.Println("Creating Singleton For Database")
	time.Sleep(2000)
	fmt.Println("Creation Done")
}

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &Database{}
		db.CreateSingleConection()
	} else {
		fmt.Println("DB already created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}
