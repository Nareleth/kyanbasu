// The window represents a specific drawing area within the screen.
// These drawing windows logically separate from each other and have relative cell writing.
// A window is made up of a reference to the parent drawing screen, absolute coordinates, absolute dimensions, as well as parameters for controlling the view (z layer and hidden logic)

package kyanbasu


/* Types */
type Window struct {
    Screen          *Screen
    X,Y             int
    Width,Height    int
    Visible         bool
}


/* Functions */
// Initializes a new window
func NewWindow(screen *Screen, x,y,w,h int) *Window {
    return &Window {
        Screen:     screen,
        X:          x,
        Y:          y,
        Width:      w,
        Height:     h,
        Visible:    true,
    }
}


/* Methods */
// Draw the border around the window
func (w *Window) DrawBorder() {
    if !w.Visible {
        return
    }

    left    := -1
    right   := w.Width -2
    top     := -1
    bottom  := w.Height -2

    var (
        border_row = Cell{Char: RuneBoxLightRow}
        border_col = Cell{Char: RuneBoxLightCol}
        border_nw  = Cell{Char: RuneBoxLightNW}
        border_ne  = Cell{Char: RuneBoxLightNE}
        border_sw  = Cell{Char: RuneBoxLightSW}
        border_se  = Cell{Char: RuneBoxLightSE}
    ) 

    for x := 0; x < right; x++ {
        w.SetCell(x, top,       border_row)
        w.SetCell(x, bottom,    border_row)
    }

    for y := 0; y < bottom; y++ {
        w.SetCell(left,    y,  border_col)
        w.SetCell(right,   y,  border_col)
    }


    w.SetCell(left,     top,        border_nw)
    w.SetCell(right,    top,        border_ne)
    w.SetCell(left,     bottom,     border_sw)
    w.SetCell(right,     bottom,    border_se)
}


// Hide the window.
func (w *Window) Hide() {
   w.Visible = false 
}

// Write a cell to a relative position in the window.
func (w *Window) SetCell(x, y int, cell Cell) {
    if !w.Visible {
        return
    }

    if cell.FG == "" {
        cell.FG = Color_Reset
    }

    if cell.BG == "" {
        cell.BG = Color_Reset
    }

    absX := w.X + x +1
    absY := w.Y + y +1

    // Clamp coordinates within the screen
    if absX < w.X || absY < w.Y || absX >= w.X + w.Width || absY >= w.Y + w.Height {
        return
    }

    w.Screen.Back[absY][absX] = cell
}


// Show the window if it was hidden.
func (w *Window) Show() {
    w.Visible = true
}


// Write formated text to a relative position in the window.
func (w *Window) WriteString(x,y int, s string, args ...any) {
    if !w.Visible {
        return
    }

    absX := w.X + x +1
    absY := w.Y + y +1

    // Clamp coordinates within the screen
    if absX < w.X || absY < w.Y || absX >= w.X + w.Width || absY >= w.Y + w.Height {
        return
    }

    w.Screen.Term.Move(absX, absY)
    w.Screen.Term.WriteString(s, args...)
}
