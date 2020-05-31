package game

import "fmt"

//Table struct for represeting where tokens are going to be
type Table struct {
	Matrix []Move
}

//AddMove to table
func (t *Table) AddMove(m Move) {
	mov := m
	for i := range mov.Tokens {
		mov.Tokens[i].Selected = 0
	}
	t.Matrix = append(t.Matrix, mov)
}

//CheckTable for every every Move
func (t *Table) CheckTable() bool {
	for _, v := range t.Matrix {
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
	for _, row := range t.Matrix {
		if long > 18 {
			long = 0
			fmt.Printf("\n")
		}
		row.Print()
		fmt.Printf(" ")
		long += len(row.Tokens)
	}
	if len(t.Matrix) > 0 {
		fmt.Printf("\n\n")
	}

}

//ChangeSelected Movement
func (t *Table) ChangeSelected(right bool) {
	pos := t.findselected()
	val1 := 0
	val2 := 1
	if pos != -1 {
		if right {
			if t.Matrix[pos].FindSelected() == len(t.Matrix[pos].Tokens)-1 {
				if pos < len(t.Matrix)-1 {
					if t.Matrix[pos].Tokens[t.Matrix[pos].FindSelected()].Selected == 3 {
						val1 = 2
					}
					if t.Matrix[pos+1].Tokens[0].Selected == 2 {
						val2 = 3
					}
					t.Matrix[pos].Tokens[t.Matrix[pos].FindSelected()].Selected, t.Matrix[pos+1].Tokens[0].Selected = val1, val2
				}
			} else {
				t.Matrix[pos].ChangeSelected(true)
			}
		} else {
			if t.Matrix[pos].FindSelected() == 0 {
				if pos > 0 {
					if t.Matrix[pos].Tokens[0].Selected == 3 {
						val1 = 2
					}
					if t.Matrix[pos-1].Tokens[len(t.Matrix[pos-1].Tokens)-1].Selected == 2 {
						val2 = 3
					}
					t.Matrix[pos].Tokens[0].Selected, t.Matrix[pos-1].Tokens[len(t.Matrix[pos-1].Tokens)-1].Selected = val1, val2
				}
			} else {
				t.Matrix[pos].ChangeSelected(false)
			}
		}
	}
}

//findselected move
func (t Table) findselected() int {
	for i, v := range t.Matrix {
		if v.FindSelected() != -1 {
			return i
		}
	}
	return -1
}

//ResetSelected moves
func (t *Table) ResetSelected() {
	for i := range t.Matrix {
		t.Matrix[i].ResetSelected()
	}
}
