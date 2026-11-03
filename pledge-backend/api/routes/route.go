package routes

import (
	"github.com/gin-gonic/gin"
	"pledge-backend/api/controllers"
	"pledge-backend/api/middlewares"
	"pledge-backend/config"
)

func InitRoute(e *gin.Engine) *gin.Engine {

	e.GET("/healthz", controllers.Health)

	// version group
	v2Group := e.Group("/api/v" + config.Config.Env.Version)

	v2Group.GET("/deployment", controllers.Deployment)
	v2Group.GET("/sync/status", controllers.SyncStatus)

	// pledge-defi backend
	poolController := controllers.PoolController{}
	v2Group.GET("/poolBaseInfo", middlewares.RequireSnapshot(), poolController.PoolBaseInfo)                     //pool base information
	v2Group.GET("/poolDataInfo", middlewares.RequireSnapshot(), poolController.PoolDataInfo)                     //pool data information
	v2Group.GET("/token", poolController.TokenList)                                                              //pool token information
	v2Group.POST("/pool/debtTokenList", middlewares.CheckToken(), poolController.DebtTokenList)                  //pool debtTokenList
	v2Group.POST("/pool/search", middlewares.CheckToken(), middlewares.RequireSnapshot(), poolController.Search) //pool search

	// plgr-usdt price

	v2Group.GET("/price", controllers.DisabledProjectPrice) //new price on ku-coin-exchange

	// pledge-defi admin backend
	multiSignPoolController := controllers.MultiSignPoolController{}
	v2Group.POST("/pool/setMultiSign", middlewares.CheckToken(), multiSignPoolController.SetMultiSign) //multi-sign set
	v2Group.POST("/pool/getMultiSign", middlewares.CheckToken(), multiSignPoolController.GetMultiSign) //multi-sign get

	userController := controllers.UserController{}
	v2Group.POST("/user/login", userController.Login)                             // login
	v2Group.POST("/user/logout", middlewares.CheckToken(), userController.Logout) // logout

	return e
}
