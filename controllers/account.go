package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"tafexplorer/api"
	"tafexplorer/models"
)

type AccountController struct {
	beego.Controller
}

// @Title 账户详情
// @Description account info
// @Param	account_name	json	string	true		"查询账户名"
// @Success 200 {object} api.AccountInfoResp
// @Failure 403 body is empty
// @router /info [post]
func (c *AccountController) AccountInfo() {
	r := resp{c: &c.Controller}
	req := &api.AccountInfoReq{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		r.ErrParamUnmarshal(err)
	}
	resp, err := models.AccountInfo(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "account info failed")
		return
	}
	r.Ok(resp)
}
