package main

import (
    "fmt"
    "github.com/tailored-style/pattern-generator/geometry"
    "github.com/tailored-style/pattern-generator/patterns"
)

func main() {
    neck := 40.0
    chest := 100.0
    scyeDepth := 24.4
    naturalWaistLength := 44.6
    halfBack := 20.0
    sleeveLength := 85.0
    shirtLength := 81.0
    cuffSize := 24.0

    fmt.Println("MEASUREMENTS")
    fmt.Println(fmt.Sprintf("Neck: %.1f cm", neck))
    fmt.Println(fmt.Sprintf("Chest: %.1f cm", chest))
    fmt.Println(fmt.Sprintf("Scye Depth: %.1f cm", scyeDepth))
    fmt.Println(fmt.Sprintf("Natural Waist Length: %.1f cm", naturalWaistLength))
    fmt.Println(fmt.Sprintf("Half Back: %.1f cm", halfBack))
    fmt.Println(fmt.Sprintf("Sleeve Length: %.1f cm", sleeveLength))
    fmt.Println(fmt.Sprintf("Shirt Length: %.1f cm", shirtLength))
    fmt.Println(fmt.Sprintf("Cuff Size: %.1f cm", cuffSize))
    fmt.Println()
    
    fmt.Println("BODY SECTION")

    block := patterns.NewTailoredShirtBlock(
        neck,
        chest,
        scyeDepth,
        naturalWaistLength,
        halfBack,
        shirtLength,
    )

    //w := new(bytes.Buffer)
	//block.PrintInstructions(w)
    //fmt.Println(w.String())

    //fmt.Println()

    fmt.Println("PLOTTING")
    printPlots(block.GetPoints())

    lines := block.GetLines()
    for _, line := range lines {
        fmt.Printf(line.ToEnglish())
    }

    fmt.Println()

    fmt.Println("AUTOCAD")
    fmt.Println("--- BEGIN ---")

    //fmt.Println("(setq oldosmode (getvar 'osmode))")
    //fmt.Println("(setvar 'osmode (boole 4 1 oldosmode))")

    for _, line := range lines {
        fmt.Printf(line.ToAutoCAD())
    }

    //fmt.Println("(setvar 'osmode oldosmode)")
    fmt.Println("--- END ---")
}

func printPlots(plots map[string]geometry.Point) {
    for key, point := range plots {
        fmt.Printf("Point %s: %v\n", key, point)
    }
}
