package kyanbasu

import (
    "fmt"
)


/*
    The panel is a relative grid-based drawing area within the canvas.
    You can have 1 panel that covers the entire canvas, or multiple ones that represent split drawing areas.
    Cells that are drawn within a panel are positioned relative to the panel itself, not the whole terminal.
*/


/* Types */

// A Panel takes positioning (x, y) and size (width, height) that can be within or the entire coordinates of the canvas.
type Panel struct {
    X           int
    Y           int
    Width       int
    Height      int
    Style       PanelStyle
    CanvasGrid  [][]Cell
}


// PanelStyle is used for defining the premade border style of the panel
type PanelStyle int


// PanelRunes are used internally to make building the border easier
type PanelRunes struct {
    PanelRow, PanelCol, PanelNW, PanelNE, PanelSW, PanelSE rune
}


/* Consts */
const (
    None PanelStyle = iota
    Light
    Heavy
)


/* Functions */

// Initialize a new Panel
func NewPanel(x, y, width, height int, canvas *Canvas) *Panel {
    return &Panel {
        X:          x,
        Y:          y,
        Width:      width,
        Height:     height,
        CanvasGrid: canvas.Back,
    }
}


// Draw the Panel Borders using the selected border style. This func uses the canvas to write the rune on the border of the panel which is not a normal valid drawing coordinate.
func (p *Panel) DrawBorders(canvas *Canvas) {
    // Init the dimensions of the border
    length  := p.X - 1 + p.Width
    size    := p.Y -1 + p.Height

    // Get the selected character set
    set := p.StyleBorders()

    // If no style is selected, end the function
    if set.PanelRow == 0 {
        return
    }

    // Draw the X axis of the border
    for row := p.X; row <= length; row++ {
        // Move the cursor to the current index
        canvas.Move(row, p.Y)

        // Draw a rune char based on the posiiton of the index (for corners or non corners)
        switch row {
            case p.X:       canvas.WriteRune(set.PanelNW)
            case length:    canvas.WriteRune(set.PanelNE)
            default:        canvas.WriteRune(set.PanelRow)
        }

        // Repeat the same steps for the bottom row
        canvas.Move(row, size)
        switch row {
            case p.X:       canvas.WriteRune(set.PanelSW)
            case length:    canvas.WriteRune(set.PanelSE)
            default:        canvas.WriteRune(set.PanelRow)
        }
    }

    // Draw the Y axis of the border
    for col := p.Y; col <= size; col++ {
        canvas.Move(p.X, col)
        switch col {
            case p.Y:   canvas.WriteRune(set.PanelNW)
            case size:  canvas.WriteRune(set.PanelSW)
            default:    canvas.WriteRune(set.PanelCol)
        }

        // Repeat the same steps for the right column
        canvas.Move(length, col)
        switch col {
            case p.Y:   canvas.WriteRune(set.PanelNE)
            case size:  canvas.WriteRune(set.PanelSE)
            default:    canvas.WriteRune(set.PanelCol)
        }
        
    }


}


// Set the value of a cell relative to the panel coordinates. This is the main method for drawing directly into the panel.
func (p *Panel) SetCell(x, y int, char rune) error {
    // Translate relative panel coords with absolute terminal position.
    AbsX := p.X + x
    AbsY := p.Y + y

    // Check if the provided drawing coordinates are within the panel boundary
    if AbsX >= p.X && AbsX < p.X + p.Width && AbsY >= p.Y && AbsY < p.Y + p.Height {
        p.CanvasGrid[AbsY][AbsX] = Cell{Char: char}
        return nil
    } else {
        return fmt.Errorf("Cell out of bounds.")
    }
}


// Applies border Style
func (p *Panel) StyleBorders() PanelRunes {
    switch p.Style {
        case Light:
            return PanelRunes {
                PanelRow:   RuneBoxLightRow,
                PanelCol:   RuneBoxLightCol,
                PanelNW:    RuneBoxLightNW,
                PanelNE:    RuneBoxLightNE,
                PanelSW:    RuneBoxLightSW,
                PanelSE:    RuneBoxLightSE,
            }

        case Heavy:
            return PanelRunes {
                PanelRow:   RuneBoxHeavyRow,
                PanelCol:   RuneBoxHeavyCol,
                PanelNW:    RuneBoxHeavyNW,
                PanelNE:    RuneBoxHeavyNE,
                PanelSW:    RuneBoxHeavySW,
                PanelSE:    RuneBoxHeavySE,
            }

        default: return PanelRunes{}
    }
}


// Write text within a panel relative to the panel coordinates. Follows printf format.
func (p *Panel) WriteText(x, y int, text string, args ...any) {
    // Create formatted string with args.
    formatted := fmt.Sprintf(text, args...)

    // Itterate through the range of chars and print them to the terminal.
    for i, char := range formatted {
        p.SetCell(x + i, y, char)
    }
}
