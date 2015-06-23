package routes
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang-webapp/model"
	"log"
	"golang-webapp/form"
	"golang-webapp/db"
)

func GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	dbMap := db.Cfg.DbMap
	art, err := dbMap.Get(model.Article{}, id)
	if err != nil {
		log.Printf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "db error"})
		return
	}
	if art == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "article": art})
}

func PostArticle(c *gin.Context) {
	var json form.ArticleJSON
	c.Bind(&json)
	article := model.NewArticle(json.Title, json.Content)
	if !model.ValidArticle(&article) {
		log.Printf("validate error.")
		c.JSON(http.StatusBadRequest, gin.H{"status": "validate error"})
		return
	}
	dbMap := db.Cfg.DbMap
	err := dbMap.Insert(&article)
	if err != nil {
		log.Printf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "db error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
