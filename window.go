// The window represents a specific drawing area within the screen.
// These drawing windows logically separate from each other and have relative cell writing.
// A window is made up of a reference to the parent drawing screen, absolute coordinates, absolute dimensions, as well as parameters for controlling the view (z layer and hidden logic)

package kyanbasu


/* Types */
type Window struct {
    Screen          *Screen
    X,Y             int
    Width,Height    int
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
    }
}


/* Methods */
// Write a cell to a relative position in the window.
func (w *Window) SetCell(x, y int, cell Cell) {
    if cell.FG == "" {
        cell.FG = Color_Reset
    }

    if cell.BG == "" {
        cell.BG = Color_Reset
    }

    absX := w.X + x
    absY := w.Y + y

    w.Screen.Back[absY][absX] = cell
}


// Write formated text to a relative position in the window.
func (w *Window) WriteString(x,y int, s string, args ...any) {
    absX := w.X + x
    absY := w.Y + y

    w.Screen.Term.Move(absX, absY)
    w.Screen.Term.WriteString(s, args...)
}
