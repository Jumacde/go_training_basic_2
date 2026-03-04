package main

type mathe struct {
	x      int
	y      int
	result int
}

func (m *mathe) SetMathe(x, y, result int) {
	m.x = x
	m.y = y
	m.result = result
	m.addtion()
}

func (m *mathe) GetMathe() (int, int, int) {
	return m.x, m.y, m.result
}

func (m *mathe) addtion() (int, int, int) {
	m.result = m.x + m.y
	return m.x, m.y, m.result
}
