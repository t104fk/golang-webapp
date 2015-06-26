package sandbox
import "golang-webapp/sandbox"

func Sandbox() {
	play()
//	sandbox.count(&sandbox.Second{}) // compile error
	count(&sandbox.Second{})
}
