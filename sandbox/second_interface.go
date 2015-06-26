package sandbox
import "log"

type Second struct {
	text string
}

func (s *Second) GetString() string {
	return s.text
}

func (s *Second) SetString(text string) {
	s.text = text
}

func count(sand Sand) {
	log.Print(sand.GetString())
}
