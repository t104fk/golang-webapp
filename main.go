package main
import (
	"golang-webapp/routes"
	"gopkg.in/alecthomas/kingpin.v2"
	"golang-webapp/config"
	"golang-webapp/db"
	"os"
	"golang-webapp/sandbox"
)

func main() {
	app := kingpin.New("ecs-bgdep-test-api", "Test API for ECS Blue-Green Deployment.")
	app.Version("0.0.1")

	config.Load(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
	db.InitDb()
	router := routes.Init()

	sandbox.Sandbox()

	router.Run(":3000")
}
