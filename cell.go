package kyanbasu


/* 
    A cell is a literal rune character printing to a terminal column.
    A cell type contains:
    Char:   a rune character.
    FG:     an ANSI value for the foreground color of the rune.
    BG:     an ANSI value for the background color of the column.
*/


/* Structs */

// Cell type
type Cell struct {
    Char    rune
}

