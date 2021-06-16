// @APIVersion 1.0.1
// @Title TAF-Explorer API
// @Description the TAF-Explorer About
package routers

import (
	"github.com/astaxie/beego"
	"tafexplorer/controllers"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/block",
			beego.NSInclude(
				&controllers.BlockController{},
			),
		),
		beego.NSNamespace("/account",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),
		beego.NSNamespace("/contract",
			beego.NSInclude(
				&controllers.ContractController{},
			),
		),
		beego.NSNamespace("/homepage",
			beego.NSInclude(
				&controllers.HomepageController{},
			),
		),
		beego.NSNamespace("/transaction",
			beego.NSInclude(
				&controllers.TransactionController{},
			),
		),
		beego.NSNamespace("/vote",
			beego.NSInclude(
				&controllers.VoteController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
