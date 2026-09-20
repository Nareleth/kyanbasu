package main

import (
    "github.com/Nareleth/kyanbasu"
)


func main() {
    // Define a key and value for colors
    type ColorValue struct {
        Name string
        Code kyanbasu.Color
    }

    // Initialize color slice
    colors := []ColorValue {
        {"Black",   kyanbasu.Black},
        {"Red",     kyanbasu.Red},
        {"Green",   kyanbasu.Green},
        {"Yellow",  kyanbasu.Yellow},
        {"Blue",    kyanbasu.Blue},
        {"Magenta", kyanbasu.Magenta},
        {"Cyan",    kyanbasu.Cyan},
        {"White",   kyanbasu.White},
    }

    // Adjustable row length
    rowSize := 10

    // Unicode Block - FG will be filled later
    tile := kyanbasu.Cell{Char: kyanbasu.RuneBlockFull, FG: kyanbasu.Black}

    // Get terminal dimensions
    w, h := kyanbasu.GetSize()

    // Initialize a new Canvas with the size of the terminal    
    canvas := kyanbasu.NewCanvas(0, 0, w, h)

    // Initialize a single Panel for the full canvas size
    panel := kyanbasu.NewPanel(0,0, rowSize, len(colors), canvas)

    // Initialize a second panel for color names
    panel2 := kyanbasu.NewPanel(rowSize +1, 0, w - rowSize -1, len(colors), canvas)

    // Clear the screen
    canvas.Clear()

    // Create a grid of blocks
    grid := make([][]kyanbasu.Cell, len(colors))

    // Iterate the range of colors
    for i, color := range colors {
        // Init a tile
        grid[i] = make([]kyanbasu.Cell, rowSize)
        for x:=0; x < rowSize; x++ {
            // Set tile color
            tile.BG = color.Code

            // Set the cell
            panel.SetCell(x, i, tile)

            // Write color name in second panel
            panel2.WriteText(0, i, color.Name)
        }
    }

    // Move cursor out of canvas
    canvas.Move(0,h)

    // Flush the buffer content to terminal
    canvas.Flush()

}
