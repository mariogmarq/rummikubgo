package game

import "fmt"

//Move struct for representing multiple tiles
type Move struct {
	Score  int
	Tokens []Token
}

//IsValid move
func (m Move) IsValid() bool {
	if len(m.Tokens) < 3 {
		return false
	}
	var iswildcard = 0

	//is a series
	var colors []int
	var numbers []int
	isSeries := true
	for i := range m.Tokens {
		colors = append(colors, m.Tokens[i].Color)
		numbers = append(numbers, m.Tokens[i].Value)
	}

	for i := 0; numbers[i] == 0; i++ {
		iswildcard = i + 1
	}

	for i, v := range numbers {
		if numbers[iswildcard] != v && i != 0 {
			isSeries = false
		}
	}

	for i := iswildcard; i < len(colors)-1; i++ {
		for j := i + 1; j < len(colors); j++ {
			if colors[i] == colors[j] {
				isSeries = false
			}
		}
	}

	if isSeries {
		return true
	}

	isStair := true

	for i := iswildcard; i < len(colors); i++ {
		if colors[iswildcard] != colors[i] && i != iswildcard {
			isStair = false
		}
	}

	skipped := 0
	for i := iswildcard; i < len(numbers)-1; i++ {
		if numbers[i]+1+skipped != numbers[i+1] {
			if skipped == iswildcard {
				isStair = false
			} else {
				skipped++
				i--
			}
		}
	}

	return isStair
}

//FindScore of move
func (m *Move) FindScore() {
	if !m.IsValid() {
		m.Score = -1
	} else {
		sum := 0
		for _, v := range m.Tokens {
			sum += v.Value
		}
		m.Score = sum
	}
}

//Print move
func (m Move) Print() {
	for _, v := range m.Tokens {
		v.Print()
		fmt.Printf("/")
	}
}

//FindSelected inside move
func (m Move) FindSelected() int {
	for i, v := range m.Tokens {
		if v.Selected == 1 || v.Selected == 3 {
			return i
		}
	}
	return -1
}

//ChangeSelected in move
func (m *Move) ChangeSelected(right bool) {
	pos := m.FindSelected()
	val1 := 0
	val2 := 1
	if pos != -1 {
		if right {
			if pos < len(m.Tokens)-1 {
				if m.Tokens[pos].Selected == 3 {
					val1 = 2
				}
				if m.Tokens[pos+1].Selected == 2 {
					val2 = 3
				}
				m.Tokens[pos].Selected, m.Tokens[pos+1].Selected = val1, val2
			}
		} else {
			if pos > 0 {
				if m.Tokens[pos].Selected == 3 {
					val1 = 2
				}
				if m.Tokens[pos-1].Selected == 2 {
					val2 = 3
				}
				m.Tokens[pos].Selected, m.Tokens[pos-1].Selected = val1, val2
			}
		}
	}
}

//ResetSelected tokens
func (m *Move) ResetSelected() {
	for i := range m.Tokens {
		m.Tokens[i].Selected = 0
	}
}
