package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"tafexplorer/api"
	"tafexplorer/models"
)

type HomepageController struct {
	beego.Controller
}

// @Title homepage count
// @Description 首页的那些统计值
// @Success 200 {object} api.HomepageCountResp
// @Failure 403 body is empty
// @router /count [post]
func (c *HomepageController) HomepageCount() {
	r := resp{c: &c.Controller}
	req := &api.HomepageCountReq{}
	resp, err := models.HomepageCount(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "block info failed")
		return
	}
	r.Ok(resp)
}

// @Title homepage search
// @Description 搜索功能
// @Param	search_content	json	string	true		"搜索内容"
// @Success 200 {object} api.HomepageSearchResp
// @Failure 403 body is empty
// @router /search [post]
func (c *HomepageController) HomepageSearch() {
	r := resp{c: &c.Controller}
	req := &api.HomepageSearchReq{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		r.ErrParamUnmarshal(err)
	}
	resp, err := models.HomepageSearch(req)
	if err != nil {
		r.ErrWith(err, BlockDbErr, "block info failed")
		return
	}
	r.Ok(resp)
}
