package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"tafexplorer/api"
	"tafexplorer/models"
)

type VoteController struct {
	beego.Controller
}

// @Title vote count
// @Description 投票数量 人数的统计
// @Success 200 {object} api.VoteCountResp
// @Failure 403 body is empty
// @router /count [post]
func (c *VoteController) VoteCount() {
	r := resp{c: &c.Controller}
	req := &api.VoteCountReq{}
	resp, err := models.VoteCount(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "vote count failed")
		return
	}
	r.Ok(resp)
}

// @Title voteList
// @Description 投票数据
// @Success 200 {object} api.VoteListResp
// @Failure 403 body is empty
// @router /list [post]
func (c *VoteController) VoteList() {
	r := resp{c: &c.Controller}
	req := &api.VoteListReq{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		r.ErrParamUnmarshal(err)
	}
	resp, err := models.VoteList(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "block list failed")
		return
	}
	r.Ok(resp)
}
