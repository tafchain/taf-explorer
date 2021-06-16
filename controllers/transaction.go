package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"tafexplorer/api"
	"tafexplorer/models"
)

type TransactionController struct {
	beego.Controller
}

// @Title transactionList
// @Description 交易列表
// @Param	page_index		json	int	true		"第几页"
// @Param	page_size		json	int	true		"页大小"
// @Success 200 {object} api.TransListResp
// @Failure 403 body is empty
// @router /list [post]
func (c *TransactionController) TransactionList() {
	r := resp{c: &c.Controller}
	req := &api.TransListReq{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		r.ErrParamUnmarshal(err)
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp, err := models.TransactionList(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "block list failed")
		return
	}
	r.Ok(resp)
}

// @Title transactionInfo
// @Description 交易详情
// @Param	trans_id		json	string	true		"交易id"
// @Success 200 {object} api.TransInfoResp
// @Failure 403 body is empty
// @router /info [post]
func (c *TransactionController) TransactionInfo() {
	r := resp{c: &c.Controller}
	req := &api.TransInfoReq{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		r.ErrParamUnmarshal(err)
	}
	resp, err := models.TransactionInfo(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "block list failed")
		return
	}
	r.Ok(resp)
}
