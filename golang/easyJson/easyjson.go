package main

import "fmt"
import "time"
import ej "random/golang/easyJson/model"

func main() {
	t1 := time.Now()
	var d ej.Data
	d.Name = "sharathbj"
	d.Age = 23
	data, _ := d.MarshalJSON()
	fmt.Println(string(data))
	fmt.Println("elapsedTime:", time.Now().Sub(t1))
}
