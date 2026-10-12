package kyanbasu

/*
    A Cell type is a literal rune in the terminal.
    We can customize it with colors and render characters to the terminal.
*/

/* Types */

// Cell
type Cell struct {
    Char    rune
    FG      Color
    BG      Color
}

