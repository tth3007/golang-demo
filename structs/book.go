package structs

type Book struct {
	ID    int
	Name  string
	Pages uint64
}

func (b *Book) GetType() string {
	if b.Pages >= 100 {
		return "NOVEL"
	} else {
		return "SHORT"
	}
}
