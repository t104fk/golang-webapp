package sandbox
import "log"

type Fourth interface {
	GetString() string
}

type Implemented struct {}

func (i *Implemented) GetString() string {
	return "implemented"
}

func (i *Implemented) SetString() {
	log.Print("set string")
}

func fourth_interface(v interface{}) string {
	f, ok := v.(Fourth)
	if ok {
		return f.GetString()
	} else {
		return "not implemented"
	}
}

func fourth_type(v interface{}) string {
	switch checked := v.(type) {
		case First:
			log.Print("first")
			return checked.GetString()
		// type can implement multiple interface,
		// so line from strong interface to weak.
		// っていうかこの場合設計が悪いｗ
		case Sand:
			log.Print("sand")
			return checked.GetString()
		case Fourth:
			log.Print("fourth")
			return checked.GetString()
		case string:
			return "not implemented"
	}
	return "no!!!"
}
