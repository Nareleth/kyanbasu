// The screen holds the terminal rendering logic.
// It utilizes a 2d cell buffer and holds tools for drawing and writing cells.

package kyanbasu


/* Types */
// A screen drawing area
type Screen struct {
    Term            *Term
    Width,Height    int
    Back            [][]Cell
    Front           [][]Cell
}


/* Functions */
// Initializes a new screen
func NewScreen(term *Term) *Screen {
    w, h := term.GetSize()
    back    := make([][]Cell, h)
    front   := make([][]Cell, h)
    
    for i := range front {
        back[i]     = make([]Cell, w)
        front[i]    = make([]Cell, w)
    }

    return &Screen {
        Term:   term,
        Width:  w,
        Height: h,
        Back:   back,
        Front:  front,
    }
}


// Flush the screen buffer
func (s *Screen) Flush() {
    for y := range s.Front {
        for x := range s.Front[y] {
            if s.Front[y][x] != s.Back[y][x] {
                s.Front[y][x] = s.Back[y][x]
                s.Term.Move(x, y)
                s.Term.SetColor(s.Front[y][x].FG)
                // How to set background color?
                s.Term.WriteRune(s.Front[y][x].Char)
                s.Term.SetColor(Color_Reset)
            }
        }
    }

    s.Term.Flush()
}


// Set a specific cell
func (s *Screen) SetCell(x, y int, cell Cell) {
    if cell.FG == "" {
        cell.FG = Color_Reset
    }

    if cell.BG == "" {
        cell.BG = Color_Reset
    }

    s.Back[y][x] = cell
}

