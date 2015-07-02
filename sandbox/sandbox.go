package sandbox
import (
	"log"
	"golang-webapp/sandbox/codegangsta/cli"
	"runtime"
	"golang-webapp/sandbox/private"
)

func Sandbox() {
//	interface_test()
	cli_test()
	path_test()
	scope_test()
}

func interface_test() {
	first()
//	sandbox.count(&sandbox.Second{}) // compile error
	second(&Second{})
	third(&Third{
		First{},
		1,
	})
	log.Print(fourth_interface(&Implemented{}))
	log.Print("log.Print(fourth_type(&First{}))")
	log.Print(fourth_type(&First{}))
	log.Print("log.Print(fourth_type(&Implemented{}))")
	log.Print(fourth_type(&Implemented{}))
}

func cli_test() {
	cli.Start()
}

func path_test() {
	log.Print(runtime.GOROOT())
}

func scope_test() {
	private.Grobal("scope")
	p := &private.Private{}
	log.Print(p)
}
