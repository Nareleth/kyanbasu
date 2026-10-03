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
    Reset   Color = "\x1b[0m"
    Black   Color = "\x1b[30m"
    Red     Color = "\x1b[31m"
    Green   Color = "\x1b[32m"
    Yellow  Color = "\x1b[33m"
    Blue    Color = "\x1b[34m"
    Magenta Color = "\x1b[35m"
    Cyan    Color = "\x1b[36m"
    White   Color = "\x1b[37m"
)
