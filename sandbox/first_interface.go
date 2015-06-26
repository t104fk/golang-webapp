package sandbox
import "log"

type Sand interface {
	GetString() string
	SetString(string)
}

type First struct {
	text string
}

func (b *First) GetString() string {
	return b.text
}

func (b *First) SetString(text string) {
	b.text = text
}

func play() {
	box := &First{}
	box.SetString("pan!")
	log.Print(box.GetString())
}
