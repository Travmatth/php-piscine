package Targaryen

type Targaryen struct {
}

type Fireproof interface {
	ResistsFire() bool
}

func NewTargaryen() (t *Targaryen) {
	t = &Targaryen{}
	return
}

func (t *Targaryen) ResistsFire() bool {
	return false
}

func GetBurned(person Fireproof) (s string) {
	if person.ResistsFire() {
		s = "emerges naked but unharmed"
	} else {
		s = "burns alive"
	}
	return
}
