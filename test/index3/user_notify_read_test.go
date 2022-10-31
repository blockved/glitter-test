package index3

import (
	"encoding/json"
	"fmt"
	. "gopkg.in/check.v1"
	"otherpay-test/client"
	"otherpay-test/common"
	"time"
)

type UserNotifyRead struct {
}

var _ = Suite(&UserNotifyRead{})

var (
	UserNotifyReadUrl string = "http://localhost:8765/v1/user/notify/read"
)

func (s *UserNotifyRead) TestUserNotifyReadCase00(goCheck *C) {
	//用户登录，然后调用UserNotifyRead接口，校验user_notify表中addr的count被置0
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
	reqUserNotifyRead := RequestToken{
		Address: addr,
		Token:   token,
	}
	respStr, err = common.DoPost(UserNotifyReadUrl, common.ConvToJSON(reqUserNotifyRead))
	var respUserNotifyRead Response
	err = json.Unmarshal(respStr, &respUserNotifyRead)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(respUserNotifyRead.Code, Equals, uint32(0))
	sql := fmt.Sprintf("select count from user_notify where address = \"%s\"", addr)
	rows, err := client.MysqlClientIndex3().Query(sql)
	goCheck.Assert(err, IsNil)
	num := 0
	for rows.Next() {
		rows.Scan(&num)
	}
	goCheck.Assert(num, Equals, 0)
}
