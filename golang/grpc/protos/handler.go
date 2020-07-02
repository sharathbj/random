package protos

import "context"

type SamplePing struct {

}

func (p *SamplePing)PingCheck(c context.Context, ps *PingStruct) (*PingStruct,error) {
	return &PingStruct{Resp:"Acknowledgement from server :)"},nil
}