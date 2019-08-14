package main

import (
	"./app/pkg/setting"
	"./app/routers"
)

// // User s
// type User struct {
// 	Name string `bson:"name"`
// 	Age  int    `bson:"age"`
// }

func main() {
	r := routers.InitRouter()
	r.Run("127.0.0.1:" + setting.HTTPPort)
	// test1
}
