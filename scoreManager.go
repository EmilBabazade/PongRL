package main

type ScoreManager struct {
	p1 int
	p2 int
}

func (s *ScoreManager) reset() {
	s.p1 = 0
	s.p2 = 0
}

func (s *ScoreManager) p1Score() {
	s.p1++
}

func (s *ScoreManager) p2Score() {
	s.p2++
}
