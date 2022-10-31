package index3

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"otherpay-test/common"
	"time"
)

type UserReplay struct {
}

var _ = Suite(&UserReplay{})

var (
	UserReplayUrl string = "http://localhost:8765/v1/user/replay"
)

func (s *UserReplay) TestUserReplayCase00(goCheck *C) {
	//先注册一个用户，然后调用UserReplay接口，查看用户的回复内容
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
	reqUserReplay := RequestToken{
		Address: addr,
		Token:   token,
	}

	respStr, err = common.DoPost(UserReplayUrl, common.ConvToJSON(reqUserReplay))
	var respUserReplay Response
	err = json.Unmarshal(respStr, &respUserReplay)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(respUserReplay.Code, Equals, uint32(0))
	replays := respUserReplay.Data.(Replays)
	goCheck.Assert(len(replays), Not(Equals), 0)
}
