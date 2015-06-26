package sandbox
import "log"

type Sand interface {
	GetString() string
	SetString(string)
}

type First struct {
	text string
}

func (f *First) GetString() string {
	return f.text
}

func (f *First) SetString(text string) {
	f.text = text
}

func first() {
	first := &First{}
	first.SetString("pan!")
	log.Print(first.GetString())
}
