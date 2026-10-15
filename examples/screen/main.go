// Screen Demo
// Test the capabilities of the functionality of the screen layer.
// This demo showcases cell control with a tilemap.
// The tilemap values are stored in an external JSON file.

// AAYOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO
// Use this tutorial for json champ:
// https://tutorialedge.net/golang/parsing-json-with-golang/

package main

import (
    "fmt"
    "io"
    "os"
    "time"
    "github.com/Nareleth/kyanbasu"
)

const(
    FPS             = 30                // Set FPS
    Tilemap_File    = "tilemap.json"      // Input tilemap file
    // TEMP
    World_Size = 10
)

// TEMP
var tileGrass = kyanbasu.Cell{Char: '.', FG: kyanbasu.Color_Green,}


// Draw the visible tile map
func DrawMap(screen *kyanbasu.Screen) {
    for y := 1; y < 10; y++ {
        for x := 1; x < 10; x++ {
            screen.SetCell(x, y, tileGrass)
        }
    }
}


// Read our tilemap file
func ReadFile(file string) {
    // Open file 
    jsonFile, err := os.Open(file)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Defer file close
    defer jsonFile.Close()

    // Read file as byte array
    file_value, err := io.ReadAll(jsonFile)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // need to somehow initialize the tiles as cells. maybe make a tile type? or reuse cell?
}


func main() {
    // Initializes a new Term object to control terminal manipulation.
    // We defer the exit behavior now so it can be captured on close.
    term := kyanbasu.NewTerm()
    defer term.CloseTerm()

    // Initializes a new Screen object for controlling rendering and flushing.
    screen := kyanbasu.NewScreen(term)

    // Read the tilemap file
    ReadFile(Tilemap_File)

    // Draw the tilemap
    DrawMap(screen)

/*
    // Game loop
    for {
        // Flush all terminal drawing instructions
        screen.Flush()

        // Wait to run again for fps limit
        time.Sleep(time.Second / time.Duration(FPS))
    }
*/
    // temp until i can use the loop
    screen.Flush()
    time.Sleep(5 * time.Second)
}
