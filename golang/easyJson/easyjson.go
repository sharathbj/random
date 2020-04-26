package main

import (
	"fmt"
)
import "time"
import ej "random/golang/easyJson/model"

func main() {
	t1 := time.Now()
	var d ej.Data
	d.Name = "sharathbj"
	d.Age = 23
	data, _ := d.MarshalJSON()
	fmt.Println(string(data))
	fmt.Println("elapsedTime for marshal:", time.Now().Sub(t1))

	val := []byte(`{"name":"sharathbj","age":23}`)

	t2 := time.Now()
	err :=  d.UnmarshalJSON(val)
	fmt.Println("elapsedTime for Unmarshal:",err, time.Now().Sub(t2))
}
