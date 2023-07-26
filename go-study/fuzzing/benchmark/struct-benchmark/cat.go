package main

type Cat struct {
	name string
}

// pointer receiver
func (this *Cat) setName(name string) {
	this.name = name
}

// value receiver
func (this Cat) changeName(name string) {
	this.name = name
}
