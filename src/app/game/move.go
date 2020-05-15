package game

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

	//is a series
	var colors []int
	var numbers []int
	isSeries := true
	for i := range m.Tokens {
		colors = append(colors, m.Tokens[i].Color)
		numbers = append(numbers, m.Tokens[i].Value)
	}

	for i, v := range numbers {
		if numbers[0] != v && v != WILDCARD && i != 0 {
			isSeries = false
		}
	}

	for i := 0; i < len(colors)-1; i++ {
		for j := i + 1; j < len(colors); j++ {
			if colors[i] == colors[j] && colors[i] != WILDCARD {
				isSeries = false
			}
		}
	}

	if isSeries {
		return true
	}

	isStair := true

	for i, v := range colors {
		if colors[0] != v && v != WILDCARD && i != 0 {
			isStair = false
		}
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] != numbers[i+1] {
			isStair = false
		}
	}

	return isStair
}

//FindScore of move
func (m *Move) FindScore() {
	if !m.IsValid() {
		m.Score = -1
	}
	sum := 0
	for _, v := range m.Tokens {
		sum += v.Value
	}
	m.Score = sum
}

//Print move
func (m Move) Print() {
	for _, v := range m.Tokens {
		v.Print()
	}
}
