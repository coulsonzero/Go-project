package 封装

type Student struct {
	Person
	score  float64
	gender bool
}

func (s *Student) GetScore() float64 {
	return s.score
}

func (s *Student) getGender() bool {
	return s.gender
}
