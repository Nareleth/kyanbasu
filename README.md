# Kyanbasu~ (きゃんばす)
## Overview
A go terminal ui (tui) framework built for games and animations live in your terminal!
It utilizes a 4 layer tiered system to achieve fantastic results.
The terminal layer is a driver for all underlying terminal manipulation.


# TODO:
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
