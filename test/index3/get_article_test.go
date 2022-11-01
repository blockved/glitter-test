package index3

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"otherpay-test/common"
	"time"
)

type GetArticle struct {
}

var _ = Suite(&GetArticle{})

var (
	GetArticleUrl string = "http://localhost:8765/v1/articles/get_by_id"
)

type RequestGetArticle struct {
	Address    string `json:"address"`
	Token      string `json:"token"`
	ArticleID  string `json:"article_id"`
	SourceType string `json:"source_type"`
}

func (s *GetArticle) TestGetArticleCase00(goCheck *C) {
	//先注册或者登录一个用户，然后调用GetArticle接口，查看返回的文章及评论数据
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

	getArticleReq := RequestGetArticle{
		Address:    addr,
		Token:      token,
		ArticleID:  "UU5yS0VH_4DJpGqVzOtXCaEm2dYhtVpifO_TBqx5P-M",
		SourceType: "mirror",
	}
	respStr, err = common.DoPost(GetArticleUrl, common.ConvToJSON(getArticleReq))
	common.PrintInfo("GetArticleUrl_resp: %v", string(respStr))
}
