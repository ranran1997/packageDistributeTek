package routers

import (
	"github.com/astaxie/beego"
	"packageDistributeTek/controllers"
)

func init() {
	//自动升级需要支持的接口
	beego.Router("/TestAlive", &controllers.MainController{}, "get:TestAlive")
	beego.Router("/Update", &controllers.MainController{}, "get:Update")
	beego.Router("/NewUpdate", &controllers.MainController{}, "get:NewUpdate") //
	beego.Router("/UpdateNow", &controllers.MainController{}, "get:UpdateNow") //

	// 主页
	beego.Router("/", &controllers.MainController{}, "get:Index")

	// 拣选
	beego.Router("/PickUpIndex", &controllers.MainController{}, "get:PickUpIndex")
	beego.Router("/SubmitPickupID", &controllers.MainController{}, "get:SubmitPickupID")
	beego.Router("/GetUncompltedOrdersCount", &controllers.MainController{}, "get:GetUncompltedOrdersCount")
	// beego.Router("/PickingupOrdersDetailIndex", &controllers.MainController{}, "get:PickingupOrdersDetailIndex")

	// 订单
	beego.Router("/OrderListIndex", &controllers.MainController{}, "get:OrderListIndex")
	beego.Router("/OrderInfoList", &controllers.MainController{}, "get:OrderInfoList")
	beego.Router("/AddOrderIndex", &controllers.MainController{}, "get:AddOrderIndex")
	beego.Router("/RemoveOrder", &controllers.MainController{}, "get:RemoveOrder")
	beego.Router("/UploadInfoFromFile", &controllers.MainController{}, "post:UploadInfoFromFile")
	beego.Router("/OrderDetailIndex", &controllers.MainController{}, "get:OrderDetailIndex")
	beego.Router("/OrderDetail", &controllers.MainController{}, "get:OrderDetail")
	beego.Router("/ClearCompletedOrders", &controllers.MainController{}, "get:ClearCompletedOrders")

	// 订单分配信息
	beego.Router("/OrderToExpressmanManagementIndex", &controllers.MainController{}, "get:OrderToExpressmanManagementIndex")
	beego.Router("/OrderToExpressmanList", &controllers.MainController{}, "get:OrderToExpressmanList")
	beego.Router("/AddDistributeInfoIndex", &controllers.MainController{}, "get:AddDistributeInfoIndex")

	// 商品条码信息
	beego.Router("/ProductManagementIndex", &controllers.MainController{}, "get:ProductManagementIndex")
	beego.Router("/ProductList", &controllers.MainController{}, "get:ProductList")
	beego.Router("/AddProductIndex", &controllers.MainController{}, "get:AddProductIndex")

}
