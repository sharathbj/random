package gomockk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Requester interface {
	Crawl(url string) string
}

type RequestInp struct {
	URL string
	Request Requester
}

func (this *RequestInp)GetInfo(requester Requester) string {
	return this.Request.Crawl(this.URL)
}

func (this *RequestInp)Crawl(url string) string {
	//fmt.Println("inside crawl",url)
	client := http.Client{}
	client.Timeout=15*time.Second

	rs := RespStruct{}
	if resp,err := client.Get(url);err != nil{
		fmt.Println("http err",err)
		return ""
	}else {
		data, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(data, &rs)
	}
	return rs.Data.R[0].Ct.N
}

type RespStruct struct {
	Success bool `json:"success"`
	Data struct {
		R []struct {
			ID string `json:"_id"`
			N string `json:"n"`
			T string `json:"t"`
			Mt string `json:"mt"`
			Xtr struct {
				Ar int `json:"ar"`
				Cn string `json:"cn"`
				CntID string `json:"cnt_id"`
				Eid string `json:"eid"`
				Cc string `json:"cc"`
				Dis int `json:"dis"`
				Cid string `json:"cid"`
				DisType string `json:"dis_type"`
				Fsc int `json:"fsc"`
				CntN string `json:"cnt_n"`
			} `json:"xtr"`
			Iata string `json:"iata"`
			Ct struct {
				N string `json:"n"`
				ID string `json:"_id"`
			} `json:"ct"`
		} `json:"r"`
	} `json:"data"`
}