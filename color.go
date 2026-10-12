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
    Color_Reset   Color = "\x1b[0m"
    Color_Black   Color = "\x1b[30m"
    Color_Red     Color = "\x1b[31m"
    Color_Green   Color = "\x1b[32m"
    Color_Yellow  Color = "\x1b[33m"
    Color_Blue    Color = "\x1b[34m"
    Color_Magenta Color = "\x1b[35m"
    Color_Cyan    Color = "\x1b[36m"
    Color_White   Color = "\x1b[37m"
)
