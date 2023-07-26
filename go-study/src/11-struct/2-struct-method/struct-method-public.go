package main

type student struct {
	name string
	age  int
}

// private method
func (s *student) setName(name string) {
	s.name = name
}

// GetName
// public method
func (s *student) GetName() string {
	return s.name
}