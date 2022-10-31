package index3

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"otherpay-test/common"
	"time"
)

type ArticleComment struct {
}

var _ = Suite(&ArticleComment{})

var (
	ArticleCommentUrl string = "http://localhost:8765/v1/user/replay"
)

type ArticleCommentReq struct {
	ArticleID      string `json:"article_id" example:"cjOdwnwH1IA0P6Dup56KrsGEpXA4CNKz9kWqNTeHAWo"`
	SourceType     string `json:"source_type"`
	Content        string `json:"content"`
	ParientComment string `json:"parient_comment"`
	CheckInfo
}



func (s *ArticleComment) TestArticleCommentCase00(goCheck *C) {
	//先注册一个用户，然后调用ArticleComment接口，添加评论
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
	goCheck.Assert(err, IsNil)
	var resp Response
	err = json.Unmarshal(respStr, &resp)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(resp.Code, Equals, uint32(0))

	reqArticleComment := ArticleCommentReq{
		ArticleID:      "UU5yS0VH_4DJpGqVzOtXCaEm2dYhtVpifO_TBqx5P-M",
		SourceType:     "mirror",
		Content:        "test comment 123",
		ParientComment: "",
		CheckInfo:      CheckInfo{
			Address: addr,
			Msg:     msg,
			Sign:    sign,
		},
	}

	respStr, err = common.DoPost(ArticleCommentUrl, common.ConvToJSON(reqArticleComment))
	var respArticleComment Response
	err = json.Unmarshal(respStr, &respArticleComment)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(respArticleComment.Code, Equals, uint32(0))
}
