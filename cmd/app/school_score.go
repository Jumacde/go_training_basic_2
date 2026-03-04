package main

import "fmt"

type school_score struct {
	score_list []int
	name       string
	eng_score  int
	math_score int
	sum_score  int
}

func (s *school_score) SetSchool_score(name string, eng_score, math_score, sum_score int) {
	s.name = name
	s.eng_score = eng_score
	s.math_score = math_score
	s.sum_score = sum_score
	s.calc_sum_score()
	s.determinate_School_score()
}

func (s *school_score) GetSchool_score() (string, int, int, int) {
	return s.name, s.eng_score, s.math_score, s.sum_score
}

func (s *school_score) calc_sum_score() {
	if s.eng_score >= 0 && s.math_score >= 0 {
		sum := s.eng_score + s.math_score
		s.sum_score = sum
	} else {
		fmt.Println("error value")
	}
}

func (s *school_score) determinate_School_score() {
	s.calc_sum_score()
	if s.name == "Alice" {
		alice_score := []int{s.eng_score, s.math_score, s.sum_score}
		s.score_list = alice_score
	} else if s.name == "Bob" {
		bob_score := []int{s.eng_score, s.math_score, s.sum_score}
		s.score_list = bob_score
	}

}
