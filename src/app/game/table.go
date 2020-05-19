package game

import "fmt"

//Table struct for represeting where tokens are going to be
type Table struct {
	matrix []Move
}

//AddMove to table
func (t *Table) AddMove(m Move) {
	mov := m
	for i := range mov.Tokens {
		mov.Tokens[i].Selected = 0
	}
	t.matrix = append(t.matrix, mov)
}

//CheckTable for every every Move
func (t *Table) CheckTable() bool {
	for _, v := range t.matrix {
		v.FindScore()
		if v.Score < 0 {
			return false
		}
	}
	return true
}

//Print table
func (t Table) Print() {
	var long = 0
	for _, row := range t.matrix {
		if long > 18 {
			long = 0
			fmt.Printf("\n")
		}
		row.Print()
		fmt.Printf(" ")
		long += len(row.Tokens)
	}
	if len(t.matrix) > 0 {
		fmt.Printf("\n\n")
	}

}
