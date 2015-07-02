package private
import "log"

type private struct {
	Name string
}

func (p *private) shout() {
	log.Print(p.Name)
}

func Grobal(s string) {
	log.Print(s)
}
