package domain

type Frobnicator struct {
	Id  string
	Bar string
}

func (frob *Frobnicator) GetId() string {
	return frob.Id
}

func NewFrobnicator(id string, bar string) *Frobnicator {
	return &Frobnicator{Id: id, Bar: bar}
}
