package game

//Game struct for zipping all other classes
type Game struct {
	Players []Player
	Table   Table
	playing int
}

//Print game to screen
func (g Game) Print() {
	g.Table.Print()
	g.Players[0].Print()
}

//CreateGame function
func CreateGame(online bool) Game {
	g := Game{}
	if !online {
		g.Players = append(g.Players, CreatePlayer())
		g.playing = 1
	}
	return g
}

//ChangeSection for changing from player to table and viceversa
func (g *Game) ChangeSection(up bool, p int) {
	if up {
		if len(g.Table.Matrix) > 0 {
			pos := g.Players[p].FindSelected()
			if pos != -1 {
				if g.Players[p].Tokens[pos].Selected == 1 {
					g.Players[p].Tokens[pos].Selected = 0
				} else {
					g.Players[p].Tokens[pos].Selected = 2
				}
				if g.Table.Matrix[0].Tokens[0].Selected == 0 {
					g.Table.Matrix[0].Tokens[0].Selected = 1
				} else {
					g.Table.Matrix[0].Tokens[0].Selected = 3
				}
			}
		}
	} else {
		if len(g.Players[p].Tokens) > 0 {
			pos := -1
			for i, v := range g.Table.Matrix {
				if v.FindSelected() != -1 {
					pos = i
				}
			}
			if pos != -1 {
				if g.Table.Matrix[pos].Tokens[g.Table.Matrix[pos].FindSelected()].Selected == 1 {
					g.Table.Matrix[pos].Tokens[g.Table.Matrix[pos].FindSelected()].Selected = 0
				} else {
					g.Table.Matrix[pos].Tokens[g.Table.Matrix[pos].FindSelected()].Selected = 2
				}
				if g.Players[p].Tokens[0].Selected == 0 {
					g.Players[p].Tokens[0].Selected = 1
				} else {
					g.Players[p].Tokens[0].Selected = 3
				}
			}
		}
	}
}

//ChangeSelected between game
func (g *Game) ChangeSelected(right bool, p int) {
	if right {
		if g.Players[p].FindSelected() == -1 {
			g.Table.ChangeSelected(true)
		} else {
			g.Players[p].ChangeSelected(true)
		}
	} else {
		if g.Players[p].FindSelected() == -1 {
			g.Table.ChangeSelected(false)
		} else {
			g.Players[p].ChangeSelected(false)
		}
	}
}

//ResetSelected function
func (g *Game) ResetSelected() {
	g.Table.ResetSelected()
	for i := range g.Players {
		g.Players[i].ResetSelected()
	}
}

//CreateMove function
func (g *Game) CreateMove(p int) Move {
	mov := Move{}
	for _, v := range g.Players[p].Tokens {
		if v.Selected == 2 || v.Selected == 3 {
			mov.Tokens = append(mov.Tokens, v)
		}
	}
	//Borramos de player
	g.Players[p].RemoveMove(mov)
	for _, v := range g.Table.Matrix {
		for _, h := range v.Tokens {
			if h.Selected == 2 || h.Selected == 3 {
				mov.Tokens = append(mov.Tokens, h)
			}
		}
	}

	//Borramos de table
	for i := 0; i < len(g.Table.Matrix); i++ {
		for j := 0; j < len(g.Table.Matrix[i].Tokens); j++ {
			if g.Table.Matrix[i].Tokens[j].Selected == 2 || g.Table.Matrix[i].Tokens[j].Selected == 3 {
				g.Table.Matrix[i].Tokens = append(g.Table.Matrix[i].Tokens[:j], g.Table.Matrix[i].Tokens[j+1:]...)
				j = j - 1
			}
		}
	}

	//Borramos los movimientos vacios
	for i := 0; i < len(g.Table.Matrix); i++ {
		if len(g.Table.Matrix[i].Tokens) == 0 {
			g.Table.Matrix = append(g.Table.Matrix[:i], g.Table.Matrix[i+1:]...)
		}
	}

	return mov
}

//SelectToken function
func (g *Game) SelectToken(p int) {
	if g.Players[p].FindSelected() != -1 {
		g.Players[p].SelectToken()
	} else {
		pos := g.Table.findselected()
		if g.Table.Matrix[pos].Tokens[g.Table.Matrix[pos].FindSelected()].Selected == 1 {
			g.Table.Matrix[pos].Tokens[g.Table.Matrix[pos].FindSelected()].Selected = 3
		} else if g.Table.Matrix[pos].Tokens[g.Table.Matrix[pos].FindSelected()].Selected == 3 {
			g.Table.Matrix[pos].Tokens[g.Table.Matrix[pos].FindSelected()].Selected = 1
		}
	}
}
