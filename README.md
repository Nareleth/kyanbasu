# Kyanbasu~ (きゃんばす)
## Overview
A go terminal ui (tui) framework built for games and animations live in your terminal!
It utilizes a 4 layer tiered system to achieve fantastic results.
The terminal layer is a driver for all underlying terminal manipulation.


# TODO:
## Screen Layer: handles the terminal rendering
- Figured out how to get FG and BG but requires an entire color rework.

X- screen struct:
    X- term
    X- width height (uses terms size func) 
    X- back buffer
    X- front buffer
X- diff logic
X- flush
- how to draw FG and BG
- resize
- additional drawing logic (filling areas like rectangles)
- rename back and front lol i hate it

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

## Extra:
- fps
- input layer (easier time establishing a goroutine for key input)

## Bugs:
### Term:
- cant ctrl+c to close (implement a fix in input layer)

## Example:
### Screen:
- maybe add more map generation logic?
- test writing BG folor with FG braille
- Read JSON
- potentially may wanna add a form of unmarshall method for cells
