package main

type calc struct {
	form string
	x    uint
	y    uint
	z    uint
}

func (cal *calc) SetCalc(form string, x, y uint) {
	cal.form = form
	cal.x = x
	cal.y = y
	cal.calc_z()
}

func (cal *calc) GetCalc() (string, uint, uint, uint) {
	return cal.form, cal.z, cal.x, cal.y
}

func (cal *calc) calc_z() {
	cal.z = cal.x + cal.y
}

func (cal *calc) calc_div() (string, uint, uint, uint) {
	cal.calc_z()
	div_result1 := cal.x / cal.y
	div_result2 := cal.x / cal.z
	div_result3 := cal.y / cal.z

	if cal.x > cal.y {
		return cal.form, cal.x, cal.y, div_result1
	} else if cal.x < cal.y {
		return cal.form, cal.x, cal.z, div_result2
	} else {
		return cal.form, cal.y, cal.z, div_result3
	}
}
