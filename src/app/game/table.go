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

//ChangeSelected Movement
func (t *Table) ChangeSelected(right bool) {
	pos := t.findselected()
	val1 := 0
	val2 := 1
	if pos != -1 {
		if right {
			if t.matrix[pos].FindSelected() == len(t.matrix[pos].Tokens)-1 {
				if pos < len(t.matrix)-1 {
					if t.matrix[pos].Tokens[t.matrix[pos].FindSelected()].Selected == 3 {
						val1 = 2
					}
					if t.matrix[pos+1].Tokens[0].Selected == 2 {
						val2 = 3
					}
					t.matrix[pos].Tokens[t.matrix[pos].FindSelected()].Selected, t.matrix[pos+1].Tokens[0].Selected = val1, val2
				}
			} else {
				t.matrix[pos].ChangeSelected(true)
			}
		} else {
			if t.matrix[pos].FindSelected() == 0 {
				if pos > 0 {
					if t.matrix[pos].Tokens[0].Selected == 3 {
						val1 = 2
					}
					if t.matrix[pos-1].Tokens[len(t.matrix[pos-1].Tokens)-1].Selected == 2 {
						val2 = 3
					}
					t.matrix[pos].Tokens[0].Selected, t.matrix[pos-1].Tokens[len(t.matrix[pos-1].Tokens)-1].Selected = val1, val2
				}
			} else {
				t.matrix[pos].ChangeSelected(false)
			}
		}
	}
}

//findselected move
func (t Table) findselected() int {
	for i, v := range t.matrix {
		if v.FindSelected() != -1 {
			return i
		}
	}
	return -1
}

//ResetSelected moves
func (t *Table) ResetSelected() {
	for i := range t.matrix {
		t.matrix[i].ResetSelected()
	}
}
