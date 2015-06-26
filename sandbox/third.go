package sandbox
import (
	"fmt"
	"log"
)

type Third struct {
	First
	tel int
}

func (t *Third) GetString() string {
	return fmt.Sprintf("%d", t.tel)
}

func (t *Third) SetString(text string) {
	t.First.SetString(text)
}

func third(s Sand) {
	s.SetString("third")
	log.Print(s.GetString())
}
