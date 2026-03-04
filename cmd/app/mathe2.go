package main

type Mathe2 struct {
	x, y, result uint
}

func (m2 *Mathe2) SetMathe2(x, y, result uint) {
	m2.x = x
	m2.y = y
	m2.result = result
	m2.calc_mathe2_mul()
}

func (m2 *Mathe2) GetMathe2() (uint, uint, uint) {
	return m2.x, m2.y, m2.result
}

func (m2 *Mathe2) calc_mathe2_mul() (uint, uint, uint) {
	m2.result = m2.x * m2.y
	return m2.x, m2.y, m2.result
}
