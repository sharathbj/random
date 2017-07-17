package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"random/golang/protobuf/models"
)

func main() {
	fmt.Println("Hey Dude!!")

	data := &models.ProtoData{}
	data.Name = proto.String("sharath bj")
	data.Age = proto.Int32(23)
	data.From = proto.String("Ramanagara")
	data.State = proto.String("Karnataka")
	data.Country = proto.String("Indian")

	fmt.Println(data)
}
