package index3

import (
	"encoding/json"
	"fmt"
	. "gopkg.in/check.v1"
	"otherpay-test/client"
	"otherpay-test/common"
	"time"
)

type ArticleComment struct {
}

var _ = Suite(&ArticleComment{})

var (
	ArticleCommentUrl string = "http://localhost:8765/v1/article/comment"
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
	common.PrintInfo("urlLoginOrRegister_resp: %v", string(respStr))
	goCheck.Assert(err, IsNil)

	var resp Response
	err = json.Unmarshal(respStr, &resp)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(resp.Code, Equals, uint32(0))

	//评论文章
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
	common.PrintInfo("ArticleCommentUrl_resp: %v", string(respStr))

	var respArticleComment Response
	err = json.Unmarshal(respStr, &respArticleComment)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(respArticleComment.Code, Equals, uint32(0))
	commentID := respArticleComment.Data.(map[string]interface{})["id"].(string)
	goCheck.Assert(commentID, Not(Equals), "")
	//parientComment != "", 评论回复
	reqArticleComment1 := ArticleCommentReq {
		ArticleID:      "UU5yS0VH_4DJpGqVzOtXCaEm2dYhtVpifO_TBqx5P-M",
		SourceType:     "mirror",
		Content:        "test comment 456 after 123",
		ParientComment: commentID,
		CheckInfo:      CheckInfo{
			Address: addr,
			Msg:     msg,
			Sign:    sign,
		},
	}
	respStr1, err1 := common.DoPost(ArticleCommentUrl, common.ConvToJSON(reqArticleComment1))
	goCheck.Assert(err1, IsNil)
	common.PrintInfo("ArticleCommentUrl_resp1: %v", string(respStr1))
	err = json.Unmarshal(respStr1, &respArticleComment)
	goCheck.Assert(err, IsNil)
	commentID1 := respArticleComment.Data.(map[string]interface{})["id"].(string)
	sql := fmt.Sprintf("select article_id, content, parent, from_user_id, to_user_id from comment where id = \"%s\"", commentID1)
	rows, err := client.MysqlClientIndex3().Query(sql)
	goCheck.Assert(err, IsNil)

	var articleID string
	var content string
	var parent string
	var fromUserID string
	var toUserID string
	for rows.Next() {
		rows.Scan(&articleID, &content, &parent, &fromUserID, &toUserID)
	}
	goCheck.Assert(articleID, Equals, "UU5yS0VH_4DJpGqVzOtXCaEm2dYhtVpifO_TBqx5P-M")
	goCheck.Assert(content, Equals, "test comment 456 after 123")
	goCheck.Assert(parent, Equals, commentID)
	goCheck.Assert(fromUserID, Equals, addr)
	goCheck.Assert(toUserID, Equals, addr)


}
