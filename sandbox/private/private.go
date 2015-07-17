package private

import "log"

// Private is
type Private struct {
	Name string
}

// shout is
func (p *Private) shout() {
	log.Print(p.Name)
}

// Grobal is
func Grobal(s string) {
	log.Print(s)
}
