package sandbox

func Sandbox() {
	first()
//	sandbox.count(&sandbox.Second{}) // compile error
	second(&Second{})
	third(&Third{
		First{},
		1,
	})
}
