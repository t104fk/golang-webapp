package sandbox

import (
	"golang-webapp/sandbox/codegangsta/cli"
	"golang-webapp/sandbox/private"
	"log"
	"runtime"
)

// Sandbox is
func Sandbox() {
	interfaceTest()
	cliTest()
	pathTest()
	scopeTest()
	namedExec()
}

// interface_test is
func interfaceTest() {
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

func cliTest() {
	cli.Start()
}

func pathTest() {
	log.Print(runtime.GOROOT())
}

func scopeTest() {
	private.Grobal("scope")
	p := &private.Private{}
	log.Print(p)
}
