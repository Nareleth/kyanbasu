package main


func main() {
    // Get values for terminal dimensions (width, height)
    w, h := GetSize()

    // Initialize a new Canvas with the size of the terminal
    canvas := NewCanvas(0, 0, w, h)

    // Initialize a single Panel with the full size of the canvas
    panel := NewPanel(0, 0, w, h, canvas)

    // Style the panel with the light boxset
    panel.Style = Light

    
    // Issue a screen clear in the canvas buffer
    canvas.Clear()

    // Draw the border of the panel
    panel.DrawBorders(canvas)

    // Write the terminal dimensions to the first row of the panel
    panel.WriteText(1, 1, "The Terminal size is (%d, %d)\n", w, h)

    // Flush the buffer content to the terminal and print it
    canvas.Flush()
}
