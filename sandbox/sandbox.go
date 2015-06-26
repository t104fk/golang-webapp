package sandbox
import "log"

func Sandbox() {
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
