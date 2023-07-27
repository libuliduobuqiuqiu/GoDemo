package ReflectDemo

type ObjOperation interface {
	GetOptions() map[string]string
}

type Man struct {
	sex string
	Person
}

type Person struct {
	Name    string
	Age     int
	Options map[string]string
}

func (p *Person) GetOptions() map[string]string {
	return p.Options
}

func (p *Person) SetOptions(o map[string]string) {
	p.Options = o
}
