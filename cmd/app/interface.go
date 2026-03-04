package main

type Calc interface {
	GetCalc() (string, uint, uint, uint)
	calc_div() (string, uint, uint, uint)

	GetSchool_score() (string, int, int, int)
	determinate_School_score()
}
