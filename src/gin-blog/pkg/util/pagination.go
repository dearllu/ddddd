package util
<<<<<<< Updated upstream
import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"gin-blog/pkg/setting"
)
=======

import (
	"github.com/Unknwon/com"
	"github.com/dearllu/ddddd/src/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)
//获取分页页码
>>>>>>> Stashed changes
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
<<<<<<< Updated upstream
}
=======
}
>>>>>>> Stashed changes
