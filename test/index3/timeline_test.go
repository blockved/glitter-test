package index3

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"otherpay-test/common"
	"time"
)

type Timeline struct {
}

var _ = Suite(&Timeline{})

var (
	TimelineUrl string = "http://localhost:8765/v1/articles/timeline"
)

type RequestArticleTimeline struct {
	RequestToken
	PageInfo
}

type PageInfo struct {
	Page  int `form:"page" example:"1"`
	Limit int `form:"limit" example:"20"`
}

func (s *Timeline) TestTimelineCase00(goCheck *C) {
	//有addr
	privateHex := "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce"
	msg := "I am registing for index3"
	addr, sign := common.GetSignNew(privateHex, msg)
	req := RequestLoginOrRegister{
		CheckInfo: CheckInfo{
			Address: addr,
			Msg:     msg,
			Sign:    sign,
		},
		LoginTime: time.Now().Unix(),
	}
	respStr, err := common.DoPost(urlLoginOrRegister, common.ConvToJSON(req))
	common.PrintInfo("urlLoginOrRegister_resp: %v", string(respStr))
	goCheck.Assert(err, IsNil)
	var resp Response
	_ = json.Unmarshal(respStr, &resp)
	token := resp.Data.(map[string]interface{})["token"].(string)

	TimelineReq := RequestArticleTimeline{
		RequestToken: RequestToken{
			Address: addr,
			Token:   token,
		},
		PageInfo: PageInfo{
			Page:  1,
			Limit: 10,
		},
	}
	respStr, err = common.DoPost(TimelineUrl, common.ConvToJSON(TimelineReq))
	common.PrintInfo("TimelineUrl_resp: %v", string(respStr))
}
