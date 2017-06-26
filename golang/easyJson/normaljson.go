package main

import (
	"fmt"
	"time"
	"encoding/json"
	model "random/golang/easyJson/model"
)
func main() {
	t1 := time.Now()
	var d model.Data
	d.Name = "sharathbj"
	d.Age = 23
	data, _ := json.Marshal(d)
	fmt.Println(string(data))
	fmt.Println("elapsedTime:", time.Now().Sub(t1))
}
