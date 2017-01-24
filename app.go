package main

import (
    "fmt"
    "math"
    "github.com/tailored-style/pattern-generator/geometry"
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

	p := make([]geometry.Point, 42)

    fmt.Println("Square both ways from 0")
    p[0] = geometry.Point{0, 0}
    topLine := p[0].Y

    zeroToOne := scyeDepth + 6.0
    instruct(0, 1, zeroToOne, "square across")
    p[1] = p[0].DrawDown(zeroToOne)
    scyeDepthLine := p[1].Y

    zeroToTwo := naturalWaistLength + 2.5
    instruct(0, 2, zeroToTwo, "square across")
    p[2] = p[0].DrawDown(zeroToTwo)
    waistline := p[2].Y

    zeroToThree := shirtLength + 4.0
    instruct(0, 3, zeroToThree, "square across")
    p[3] = p[0].DrawDown(zeroToThree)
    hemline := p[3].Y

    oneToFour := (chest / 2) + 10.0
    instruct(1, 4, oneToFour, "square up to 5 and down to 6 on hem line")
    p[4] = p[1].DrawRight(oneToFour)
    p[5] = p[4].SquareToHorizontalLine(topLine)
    p[6] = p[4].SquareToHorizontalLine(hemline)

    zeroToSeven := (neck / 5) - 0.5
    instruct(0, 7, zeroToSeven, "square up")
    p[7] = p[0].DrawRight(zeroToSeven)

    sevenToEight := 4.5
    instruct(7, 8, sevenToEight, "draw in neck curve")
    p[8] = p[7].DrawUp(sevenToEight)

    zeroToNine := (zeroToOne / 5) + 4.0
    instruct(0, 9, zeroToNine, "square out")
    p[9] = p[0].DrawDown(zeroToNine)

    nineToTen := halfBack + 4.0
    instruct(9, 10, nineToTen, "square down to 11 on scye depth line and up to 12")
    p[10] = p[9].DrawRight(nineToTen)
    p[11] = p[10].SquareToHorizontalLine(scyeDepthLine)
    p[12] = p[10].SquareToHorizontalLine(topLine)

    twelveToThirteen := 1.5
    instruct(12, 13, twelveToThirteen, "square up 2 cm to 14. Join 8-14")
    p[13] = p[12].DrawRight(twelveToThirteen)
    p[14] = p[13].DrawUp(2.0)

    tenToFifteen := 10.0
    instruct(10, 15, tenToFifteen, "")
    p[15] = p[10].DrawLeft(tenToFifteen)

    tenToSixteen := 0.75
    instruct(10, 16, tenToSixteen, "join 15-16 with slight curve")
    p[16] = p[10].DrawDown(tenToSixteen)

    oneToSeventeen := (oneToFour / 2) + 0.5
    instruct(1, 17, oneToSeventeen, "square down to 18, 2.5 cm below waistline and 19 on hemline")
    p[17] = p[1].DrawRight(oneToSeventeen)
    p[18] = p[17].SquareToHorizontalLine(waistline - 2.5)
    p[19] = p[17].SquareToHorizontalLine(hemline)

    fiveToTwenty := 4.5
    instruct(5, 20, fiveToTwenty, "square out")
    p[20] = p[5].DrawDown(fiveToTwenty)

    twentyToTwentyOne := (neck / 5) - 1.0
    instruct(20, 21, twentyToTwentyOne, "")
    p[21] = p[20].DrawLeft(twentyToTwentyOne)

    twentyToTwentyTwo := (neck / 5) - 2.5
    instruct(20, 22, twentyToTwentyTwo, "draw in neck curve")
    p[22] = p[20].DrawDown(twentyToTwentyTwo)

    tenToTwentyThree := 1.5
    instruct(10, 23, tenToTwentyThree, "square out")
    p[23] = p[10].DrawDown(tenToTwentyThree)

    eightToFourteen := p[8].DistanceTo(p[14])
    twentyOneToTwentyFour := eightToFourteen + 0.5
    instruct(21, 24, twentyOneToTwentyFour, "")

    // Need to use pythagorean theorem to compute X value of 24
    twentyFourXFromTwentyOne := math.Sqrt(math.Pow(twentyOneToTwentyFour, 2.0) - math.Pow(p[23].Y - p[20].Y, 2.0))
    twentyFourX := p[21].X - twentyFourXFromTwentyOne 
    p[24] = p[23].SquareToVerticalLine(twentyFourX)

    oneToTwentyFive := (chest / 3) + 4.0
    instruct(1, 25, oneToTwentyFive, "")
    p[25] = p[1].DrawRight(oneToTwentyFive)

    twentyFiveToTwentySix := 4.0
    instruct(25, 26, twentyFiveToTwentySix, "join 24-26")
    p[26] = p[25].DrawUp(twentyFiveToTwentySix)

    twentyFourToTwentySeven := p[24].DistanceTo(p[26]) / 2
    instruct(24, 27, twentyFourToTwentySeven, "")
    p[27] = p[24].MidpointTo(p[26])

    fmt.Println("Draw an armhole shape through points 14-10, and 16, 17, 26, 24; curve arm scye inwards 1 cm at 27.")

    twentyTwoToTwentyEight := 1.5
    instruct(22, 28, twentyTwoToTwentyEight, "button stand; square down")
    p[28] = p[22].DrawRight(twentyTwoToTwentyEight)

    twentyEightToTwentyNine := 3.5
    instruct(28, 29, twentyEightToTwentyNine, "facing; square down; shape to edge at neckline")
    p[29] = p[28].DrawRight(twentyEightToTwentyNine)

    eighteenToThirty := 2.5
    instruct(18, 30, eighteenToThirty, "")
    p[30] = p[18].DrawRight(eighteenToThirty)

    eighteenToThirtyOne := 2.5
    instruct(18, 31, eighteenToThirtyOne, "")
    p[31] = p[18].DrawLeft(eighteenToThirtyOne)

    nineteenToThirtyTwo := 8.0
    instruct(19, 32, nineteenToThirtyTwo, "square across")
    p[32] = p[19].DrawUp(nineteenToThirtyTwo)

    thirtyTwoToThirtyThree := 1.5
    instruct(32, 33, thirtyTwoToThirtyThree, "")
    p[33] = p[32].DrawRight(thirtyTwoToThirtyThree)

    thirtyTwoToThirtyFour := 1.5
    instruct(32, 34, thirtyTwoToThirtyFour, "draw in curved seams")
    p[34] = p[32].DrawLeft(thirtyTwoToThirtyFour)

    nineteenToThirtyFive := p[19].DistanceTo(p[6]) / 2
    instruct(19, 35, nineteenToThirtyFive, "square up")
    p[35] = p[19].MidpointTo(p[6])

    thirtyFiveToThirtySix := 3.0
    instruct(35, 36, thirtyFiveToThirtySix, "square across to front edge")
    p[36] = p[35].DrawUp(thirtyFiveToThirtySix)

    nineteenToThirtySeven := p[3].DistanceTo(p[19]) / 2
    instruct(19, 37, nineteenToThirtySeven, "")
    p[37] = p[3].MidpointTo(p[19])

    fmt.Println("Draw shaped curves as shown from 33-36 and 34-37.")

    oneToEleven := p[1].DistanceTo(p[11])
    oneToThirtyEight := (oneToEleven / 2) + 2.0
    instruct(1, 38, oneToThirtyEight, "")
    p[38] = p[1].DrawRight(oneToThirtyEight)

    fmt.Println("If added waist shaping is required, construct a dart in the back section:")

    thirtyEightToThirtyNine := 4.0
    instruct(38, 39, thirtyEightToThirtyNine, "square down to 40, 2.5 cm below waistline")
    p[39] = p[38].DrawDown(thirtyEightToThirtyNine)
    p[40] = p[39].SquareToHorizontalLine(waistline - 2.5)

    fourtyToFourtyOne := 16.0
    instruct(40, 41, fourtyToFourtyOne, "draw a 1.5 cm dart on the line 39-41")
    p[41] = p[40].DrawDown(fourtyToFourtyOne)

    fmt.Println()

    fmt.Println("PLOTTING")
    printPlots(p)

    fmt.Println()
}

func instruct(start int, end int, dist float64, additional string) {
    fmt.Printf("%d-%d\t%.1f cm", start, end, dist)
    if (additional != "") {
        fmt.Printf("; %v", additional)
    }
    fmt.Println(".")
}

func printPlots(plots []geometry.Point) {
    for key, point := range plots {
        fmt.Printf("Point %d: %v\n", key, point)
    }
}

