package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.GlobalControllerRouter["tafexplorer/controllers:BlockController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:BlockController"],
		beego.ControllerComments{
			Method:           "BlockList",
			Router:           "/list",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:BlockController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:BlockController"],
		beego.ControllerComments{
			Method:           "BlockInfo",
			Router:           "/info",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:AccountController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:AccountController"],
		beego.ControllerComments{
			Method:           "AccountInfo",
			Router:           "/info",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:ContractController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:ContractController"],
		beego.ControllerComments{
			Method:           "ContractList",
			Router:           "/list",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:ContractController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:ContractController"],
		beego.ControllerComments{
			Method:           "ContractInfo",
			Router:           "/info",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:HomepageController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:HomepageController"],
		beego.ControllerComments{
			Method:           "HomepageCount",
			Router:           "/count",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:HomepageController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:HomepageController"],
		beego.ControllerComments{
			Method:           "HomepageSearch",
			Router:           "/search",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:TransactionController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:TransactionController"],
		beego.ControllerComments{
			Method:           "TransactionList",
			Router:           "/list",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:TransactionController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:TransactionController"],
		beego.ControllerComments{
			Method:           "TransactionInfo",
			Router:           "/info",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:VoteController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:VoteController"],
		beego.ControllerComments{
			Method:           "VoteCount",
			Router:           "/count",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["tafexplorer/controllers:VoteController"] = append(beego.GlobalControllerRouter["tafexplorer/controllers:VoteController"],
		beego.ControllerComments{
			Method:           "VoteList",
			Router:           "/list",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}
