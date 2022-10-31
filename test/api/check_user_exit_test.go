package api

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"otherpay-test/common"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type CheckUserExist struct {
}

var _ = Suite(&CheckUserExist{})

var (
	urlCheckUserExist string = "http://localhost/api/v1/checkUserExist"
)

type CheckUserExistReq struct {
	Addr string `json:"addr"`
}

type CheckUserExistResp struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    CheckUserExistRespData `json:"data"`
}

type CheckUserExistRespData struct {
	IsExist bool `json:"is_exist"`
}

func (s *CheckUserExist) TestCheckUserExistCase00(goCheck *C) {
	//校验addr不存在，CheckUserExist接口返回is_exist=false
	//privateHex := "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce"
	//addr, sign, ts := common.GetSign(privateHex)
	//fmt.Println("sign: ", sign)
	req := CheckUserExistReq{
		Addr: "test_not_in_db",
	}
	respStr, err := common.DoPost(urlCheckUserExist, common.ConvToJSON(req))
	var resp CheckUserExistResp
	goCheck.Assert(err, IsNil)
	err = json.Unmarshal(respStr, &resp)
	goCheck.Assert(resp.Code, Equals, 0)
	goCheck.Assert(resp.Data.IsExist, Equals, false)
}

func (s *CheckUserExist) TestCheckUserExistCase01(goCheck *C) {
	//校验addr不存在，CheckUserExist接口返回is_exist=false
	//privateHex := "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce"
	//addr, sign, ts := common.GetSign(privateHex)
	//fmt.Println("sign: ", sign)
	req := CheckUserExistReq{
		Addr: "test_not_in_db",
	}
	respStr, err := common.DoPost(urlCheckUserExist, common.ConvToJSON(req))
	var resp CheckUserExistResp
	goCheck.Assert(err, IsNil)
	err = json.Unmarshal(respStr, &resp)
	goCheck.Assert(resp.Code, Equals, 0)
	goCheck.Assert(resp.Data.IsExist, Equals, false)
}
