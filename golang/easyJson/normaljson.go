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
	fmt.Println("elapsedTime for marshal:", time.Now().Sub(t1))

	val := []byte(`{"name":"sharathbj","age":23}`)

	t2 := time.Now()
	err := json.Unmarshal(val,&d)
	fmt.Println("elapsedTime for Unmarshal:",err, time.Now().Sub(t2))

}
