package kyanbasu

/*
    Colors are represented to the terminal in ANSI Escape Codes.
    The "reset" color returns color back to the terminal default.
*/


/* Types */

// Color
type Color string

/* Consts */

// Colors
const (
    Reset   Color = "\033[0m"
    Black   Color = "\033[30m"
    Red     Color = "\033[31m"
    Green   Color = "\033[32m"
    Yellow  Color = "\033[33m"
    Blue    Color = "\033[34m"
    Magenta Color = "\033[35m"
    Cyan    Color = "\033[36m"
    White   Color = "\033[37m"
)
