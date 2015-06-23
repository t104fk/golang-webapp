package routes
import (
	"github.com/gin-gonic/gin"
	"golang-webapp/filter"
)

func Init() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("")
	{
		v1.Use(filter.Auth)
		v1.GET("/article/:id", GetArticle)
		v1.POST("/article", PostArticle)
	}

	return router
}
