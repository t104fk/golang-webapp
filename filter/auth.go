package filter
import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"golang-webapp/model"
	"golang-webapp/db"
)

func Auth(c *gin.Context) {
	cookie, e := c.Request.Cookie("test_uid")
	if e != nil {
		log.Printf("error: %s", e.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
		c.Abort()
		return
	}
	id := cookie.Value
	dbMap := db.Cfg.DbMap
	u, err := dbMap.Get(model.User{}, id)
	if err != nil {
		log.Printf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "db error"})
		c.Abort()
		return
	}
	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		c.Abort()
		return
	}
	user, ok := u.(*model.User)
	if !ok {
		log.Print("bad interface.")
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		c.Abort()
		return
	}
	log.Printf("authorized: %s", user.Name)
}
