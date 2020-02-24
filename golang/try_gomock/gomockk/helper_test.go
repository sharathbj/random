package gomockk

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"

	mocks "./mocks"
	mockk "github.com/golang/mock/gomock"
)

func TestRequest_GetInfo(t *testing.T) {
	ctl := mockk.NewController(t)
	defer ctl.Finish()
	mocker := mocks.NewMockRequester(ctl)
	mockReq := &RequestInp{Request:mocker}
	mocker.EXPECT().Crawl(fmt.Sprintf("https://voyager.goibibo.com/api/v1/flights_search/find_node_by_name_v2/?search_query=%s&limit=1", "BLR")).Return("Bangalore").AnyTimes()
	mockReq.URL = "https://voyager.goibibo.com/api/v1/flights_search/find_node_by_name_v2/?search_query=BLR&limit=1"
	city := mockReq.GetInfo(mocker)
	assert.Equal(t,"Bangalore",city,"error!!!!")
}