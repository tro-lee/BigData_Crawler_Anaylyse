package main

import (
	"douban/manager"
	"douban/utils"
	"encoding/json"
	"fmt"
)

func main() {
	manager := manager.New()
	go manager.Run()

	result := <-manager.Result

	fmt.Println("本次取到的数据长度为：", len(result))
	jsonData, _ := json.Marshal(result)
	utils.JsonToFile(jsonData, "result.json")
}
