package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"tafexplorer/api"
	"tafexplorer/models"
)

type ContractController struct {
	beego.Controller
}

// @Title contract
// @Description 合约列表
// @Param	page_index		json	int	true		"第几页"
// @Param	page_size		json	int	true		"页大小"
// @Success 200 {object} api.ContractListResp
// @Failure 403 body is empty
// @router /list [post]
func (c *ContractController) ContractList() {
	r := resp{c: &c.Controller}
	pageIndex, _ := c.GetInt("page_index", 0)
	req := &api.ContractListReq{
		PageIndex: pageIndex,
	}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		r.ErrParamUnmarshal(err)
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp, err := models.ContractList(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "account info failed")
		return
	}
	r.Ok(resp)
}



// @Title 合约详情
// @Description 合约 info
// @Param	contract_name	json	string	true		"合约名字"
// @Success 200 {object} api.AccountInfoResp
// @Failure 403 body is empty
// @router /info [post]
func (c *ContractController) ContractInfo() {
	r := resp{c: &c.Controller}
	req := &api.ContractInfoReq{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		r.ErrParamUnmarshal(err)
	}
	resp, err := models.ContractInfo(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "account info failed")
		return
	}
	r.Ok(resp)
}
