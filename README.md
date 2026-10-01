# Kyanbasu~ (きゃんばす)
## Overview
A go framework built for games and animations live in your terminal!
Kyanbasu utilizes a cell-based grid built on panels within a canvas to have full control of every column and row.


# TODO:
## Terminal Layer: all underlying terminal manipulation
- add description
- term struct:
    - old state
    X- out (writer)
    X- in (os.stdin)
    X- fd (int - its the stdin stdout stderr file descriptors)
X- clear screen
- raw mode (open)
- defer raw exit (panic too) (close)
- size func (uses x/term) returns the size - doesnt store it
- hide cursor
- show cursor
- defer show cursor (panic too)
- move cursor
- write
- color
- read key input

## Key
need to capture key input

## Cell:
    - rune
    - fg
    - bg

## Screen Layer: handles the terminal rendering
- screen struct:
    - term
    - width height (uses terms size func) 
    - back buffer
    - front buffer
- resize
- diff logic
- flush

## Window Layer: what game code interacts with to set cells
- struct:
    - x y
    - w h
    - Z (z layer)
    - hidden bool (is it shown?)
    - cells array
- newwindow
- logical pen? (to write within the window without moving the term cursor)
- draw border (apparently it can be done way cleaner than what ive been doing)


## Subwindow Layer: represent sections within the window - still in the window file
- subwindow func (init)
- relative adding cells
- z layering
