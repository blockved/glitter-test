package index3

import (
	"encoding/json"
	"fmt"
	. "gopkg.in/check.v1"
	"otherpay-test/client"
	"otherpay-test/common"
	"time"
)

type UserInfoModify struct {
}

var _ = Suite(&UserInfoModify{})

var (
	userInfoModifyUrl string = "http://localhost:8765/v1/user/info/modify"
)

type RequestModifyUserInfo struct {
	CheckInfo
	NickName    string `json:"nick_name"`
	HeadPicture string `json:"head_picture"`
}

func (s *UserInfoModify) TestUserInfoModifyCase00(goCheck *C) {
	//先注册一个用户，然后调用modify接口，修改用户信息
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

	nickName := "new_nick_name"
	headPicture := "http://www.headpictuer.png"
	modifyReq := RequestModifyUserInfo{
		CheckInfo: CheckInfo{
			Address: addr,
			Msg:     msg,
			Sign:    sign,
		},
		NickName:    nickName,
		HeadPicture: headPicture,
	}
	respModifyStr, err := common.DoPost(userInfoModifyUrl, common.ConvToJSON(modifyReq))
	goCheck.Assert(err, IsNil)
	var respModify Response
	err = json.Unmarshal(respModifyStr, &respModify)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(respModify.Code, Equals, uint32(0))
	sql := fmt.Sprintf("select nick_name, head_picture from mirror_user where address = \"%s\"", addr)
	rows, err := client.MysqlClientIndex3().Query(sql)
	var nickNameDb string
	var headPictureDb string
	for rows.Next() {
		rows.Scan(&nickNameDb, &headPictureDb)
		break
	}
	goCheck.Assert(nickNameDb, Equals, nickName)
	goCheck.Assert(headPictureDb, Equals, headPicture)

	//nickName不修改, headPicture修改
	headPicture2 := "head_picture_2"
	modifyReq1 := RequestModifyUserInfo{
		CheckInfo: CheckInfo{
			Address: addr,
			Msg:     msg,
			Sign:    sign,
		},
		NickName:    "",
		HeadPicture: headPicture2,
	}
	respModifyStr1, err := common.DoPost(userInfoModifyUrl, common.ConvToJSON(modifyReq1))
	goCheck.Assert(err, IsNil)
	err = json.Unmarshal(respModifyStr1, &respModify)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(respModify.Code, Equals, uint32(0))
	//校验db
	rows1, err := client.MysqlClientIndex3().Query(sql)
	for rows1.Next() {
		rows1.Scan(&nickNameDb, &headPictureDb)
		break
	}
	goCheck.Assert(nickNameDb, Equals, nickName)
	goCheck.Assert(headPictureDb, Equals, headPicture2)

	//都不修改
	modifyReq2 := RequestModifyUserInfo{
		CheckInfo: CheckInfo{
			Address: addr,
			Msg:     msg,
			Sign:    sign,
		},
		NickName:    "",
		HeadPicture: "",
	}
	respModifyStr2, err := common.DoPost(userInfoModifyUrl, common.ConvToJSON(modifyReq2))
	goCheck.Assert(err, IsNil)
	err = json.Unmarshal(respModifyStr2, &respModify)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(respModify.Code, Equals, uint32(0))

	//参数不合法
	modifyReq3 := RequestModifyUserInfo{
		CheckInfo: CheckInfo{
			Address: addr,
			Msg:     msg + " test",
			Sign:    sign,
		},
		NickName:    "",
		HeadPicture: "",
	}
	respModifyStr3, err := common.DoPost(userInfoModifyUrl, common.ConvToJSON(modifyReq3))
	goCheck.Assert(err, IsNil)
	err = json.Unmarshal(respModifyStr3, &respModify)
	goCheck.Assert(err, IsNil)
	goCheck.Assert(respModify.Code, Not(Equals), uint32(0))

}
