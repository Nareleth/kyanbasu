package kyanbasu

import (
    "os"
    "golang.org/x/term"
)


/*
    To be replaced in the future, I just don't know how to do this part myself.
    x/term is utilized to expand the capabilities of the terminal.
*/


/* Functions */

// Get the size of the terminal
func GetSize() (int, int) {
    width, height, _ := term.GetSize(int(os.Stdout.Fd()))
    return width, height
}

