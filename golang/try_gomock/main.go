package main

import(
	"fmt"
	"./gomockk"
)


func main(){
	fmt.Printf("hello ")
	req := gomockk.RequestInp{}
	requester := getRequester()
	req.Request = requester
	req.URL = fmt.Sprintf("https://voyager.goibibo.com/api/v1/flights_search/find_node_by_name_v2/?search_query=%s&limit=1", "BLR")
	fmt.Println(req.GetInfo(requester))
}

func getRequester() gomockk.Requester{
	return &gomockk.RequestInp{}
}