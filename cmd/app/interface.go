package main

type Calc interface {
	GetCalc() (string, uint, uint, uint)
	calc_div() (string, uint, uint, uint)

	GetSchool_score() (string, int, int, int)
	determinate_School_score()

	GetMathe() (int, int, int)
	addtion(int, int, int)

	GetMathe2() (uint, uint, uint)
	calc_mathe2_mul() (uint, uint, uint)
}
