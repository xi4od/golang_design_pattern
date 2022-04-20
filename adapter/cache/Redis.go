package cache

import "fmt"

type Redis struct {

}

func(t Redis) Save2Cache(){
	fmt.Println("Save data to Redis")
}
