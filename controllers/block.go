package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"tafexplorer/api"
	"tafexplorer/models"
)

type BlockController struct {
	beego.Controller
}

// @Title blockList
// @Description 区块列表
// @Param	page_index		json	int	true		"第几页"
// @Param	page_size		json	int	true		"页大小"
// @Success 200 {object} api.BlockListResp
// @Failure 403 body is empty
// @router /list [post]
func (c *BlockController) BlockList() {
	r := resp{c: &c.Controller}
	req := &api.BlockListReq{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		r.ErrParamUnmarshal(err)
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp, err := models.BlockList(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "block list failed")
		return
	}
	r.Ok(resp)
}

// @Title blockInfo
// @Description 区块详情
// @Param	block_num	json	int64	true		"区块高度"
// @Success 200 {object} api.BlockInfoResp
// @Failure 403 body is empty
// @router /info [post]
func (c *BlockController) BlockInfo() {
	r := resp{c: &c.Controller}
	req := &api.BlockInfoReq{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		r.ErrParamUnmarshal(err)
	}
	resp, err := models.BlockInfo(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "block info failed")
		return
	}
	r.Ok(resp)
}
