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
