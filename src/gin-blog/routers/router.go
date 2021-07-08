package routers

import (
	"github.com/dearllu/ddddd/src/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"
)

func InitRouter() *gin.Engine {
	r := gin.New() //创建GIN引擎，功能和gin.DEFAULT相同

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags",v1.GetTags)
	}
}
