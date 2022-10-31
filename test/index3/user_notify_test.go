package index3

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"otherpay-test/common"
	"time"
)

type Replays []Replay

type Replay struct {
	Article  Article    `json:"article"`
	FromUser MirrorUser `json:"from_user"`
	ToUser   MirrorUser `json:"to_user"`
	Comment  Comment    `json:"comment"`
}

type MirrorUser struct {
	ID               int    `gorm:"primary_key" json:"id"`
	SourceID         string `json:"source_id"`
	Address          string `json:"address"`
	ParentCapability string `json:"parent_capability"`
	LastLoginTime    int    `json:"last_login_time"`
	NickName         string `json:"nick_name"`
	HeadPicture      string `json:"head_picture"`
}

type Comment struct {
	ID string `gorm:"primary_key" json:"id"`
	//SourceID         string `json:"source_id"`
	ArticleID  string `json:"article_id"`
	Content    string `json:"content"`
	Parent     string `json:"parent"`
	Timestamp  int64  `json:"timestamp"`
	FromUserID string `json:"from_user_id"`
	ToUserID   string `json:"to_user_id"`
	SourceType string `json:"source_type"`
}

type Article struct {
	ArticleId     string `json:"article_id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	CoverUrl      string `json:"cover_url"`
	ArticleLink   string `json:"article_link"`
	Language      string `json:"language"`
	PublishedTime int    `json:"published_time"`
	Category      string `json:"category"`
	Author        string `json:"author"`
	DisplayName   string `json:"display_name"`
	Ens           string `json:"ens"`
	AvatarUrl     string `json:"avatar_url"`
}

type UserNotify struct {
}

var _ = Suite(&UserNotify{})

var (
	UserNotifyUrl string = "http://localhost:8765/v1/user/notify"
)

type RequestToken struct {
	Address string `json:"address"`
	Token   string `json:"token"`
}

func (s *UserNotify) TestUserNotifyCase00(goCheck *C) {
	//先注册一个用户，然后调用usernotify接口，查看用户的回复内容
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
	token := resp.Data.(map[string]interface{})["token"].(string)
	//token := resp.Data.(ResponseToken).Token
	reqUserNotify := RequestToken{
		Address: addr,
		Token:   token,
	}

	respStr, err = common.DoPost(UserNotifyUrl, common.ConvToJSON(reqUserNotify))
	var respUserNotify Response
	err = json.Unmarshal(respStr, &respUserNotify)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(respUserNotify.Code, Equals, uint32(0))
	replays := respUserNotify.Data.(Replays)
	goCheck.Assert(len(replays), Not(Equals), 0)
}
