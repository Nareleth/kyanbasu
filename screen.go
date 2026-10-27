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
    for y := range s.Height {
        for x := range s.Width {
            this_cell := s.Back[y][x]
            if this_cell != s.Front[y][x] {
                s.Term.Move(x, y)
                s.Term.SetColor(this_cell.FG)
                // How to set background color?
                s.Term.WriteRune(this_cell.Char)
                s.Term.SetColor(Color_Reset)
                s.Front[y][x] = this_cell
            }
        }
    }

    // Clear back buffer
    for y := range s.Back {
        for x:= range s.Back[y] {
            s.Back[y][x] = Cell{Char: ' ', FG: Color_Reset, BG: Color_Reset}
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

