// Screen Demo
// Test the capabilities of the functionality of the screen layer.
// This demo showcases cell control with a tilemap.
// The tilemap values are stored in an external JSON file.

// AAYOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO
// Use this tutorial for json champ:
// https://tutorialedge.net/golang/parsing-json-with-golang/

package main

import (
    //"fmt"
    //"io"
    "math/rand/v2"
    //"os"
    "time"
    "github.com/Nareleth/kyanbasu"
)

const(
    FPS             = 30                                // Set FPS
    Tilemap_File    = "tilemap.json"                    // Input tilemap file
    // TEMP
    World_Size = 100
)

// Global vars
var Seed = rand.New(rand.NewPCG(12345, 0))   // Define the seed values

// TEMP
var tileGrass = kyanbasu.Cell{Char: '.', FG: kyanbasu.Color_Green,}
var tileWater = kyanbasu.Cell{Char: '~', FG: kyanbasu.Color_Blue,}


// Generate a very simple tilemap
func GenerateMap() [][]kyanbasu.Cell {
    // Initialize empty world
    world := make([][]kyanbasu.Cell, World_Size)

    // Create ocean layer
    for y := range world {
        world[y] = make([]kyanbasu.Cell, World_Size)
        for x := range world[y] {
            world[y][x] = tileWater
        }
    }

    // Add land
    for y := range world {
        for x := range world[y] {
            n := Seed.IntN(100)
            switch {
                case n >= 20:
                    world[y][x] = tileGrass
            }

        }

    }
    
    // Return the map
    return world
}


// Draw the Tilemap on the screen buffer
func SetMap(screen *kyanbasu.Screen, tilemap [][]kyanbasu.Cell) {
    // Set visible dimensions
    cvx, cvy := 0, 0

    for y := 0; y < screen.Height; y++ {
        for x := 0; x < screen.Width; x++ {
            // Assign camera coords to tiles
            mapX := cvx + x
            mapY := cvy + y

            // Skip drawing tiles that dont exist
            if mapX >= World_Size || mapY >= World_Size {
                continue
            }

            screen.SetCell(x, y, tilemap[mapY][mapX])
        }
    }

    /*
    for y := range tilemap {
        for x := range tilemap [y] {
            screen.SetCell(x, y, tilemap[y][x])
        }
    }
    */
}


/*
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
*/


func main() {
    // Initializes a new Term object to control terminal manipulation.
    // We defer the exit behavior now so it can be captured on close.
    term := kyanbasu.NewTerm()
    defer term.CloseTerm()

    // Initializes a new Screen object for controlling rendering and flushing.
    screen := kyanbasu.NewScreen(term)

    // Read the tilemap file
    //ReadFile(Tilemap_File)

    // Generate the Tilemap
    tilemap := GenerateMap()

    // Draw the tilemap
    SetMap(screen, tilemap)
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
