package patterns

import (
	"fmt"
	"github.com/tailored-style/pattern-generator/geometry"
	"io"
	"math"
)

type tailoredShirtBlock struct {
	points map[string]*geometry.Point
}

func NewTailoredShirtBlock(neck float64, chest float64, scyeDepth float64, naturalWaistLength float64, halfBack float64, shirtLength float64) Pattern {
	p := make(map[string]*geometry.Point)

	p["0"] = &geometry.Point{
		X: 0,
		Y: 0,
	}
	topLine := p["0"].Y

	zeroToOne := scyeDepth + 6.0
	p["1"] = p["0"].DrawDown(zeroToOne)
	scyeDepthLine := p["1"].Y

	zeroToTwo := naturalWaistLength + 2.5
	p["2"] = p["0"].DrawDown(zeroToTwo)
	waistline := p["2"].Y

	zeroToThree := shirtLength + 4.0
	p["3"] = p["0"].DrawDown(zeroToThree)
	hemline := p["3"].Y

	oneToFour := (chest / 2) + 10.0
	p["4"] = p["1"].DrawRight(oneToFour)
	p["5"] = p["4"].SquareToHorizontalLine(topLine)
	p["6"] = p["4"].SquareToHorizontalLine(hemline)

	zeroToSeven := (neck / 5) - 0.5
	p["7"] = p["0"].DrawRight(zeroToSeven)

	sevenToEight := 4.5
	p["8"] = p["7"].DrawUp(sevenToEight)

	zeroToNine := (zeroToOne / 5) + 4.0
	p["9"] = p["0"].DrawDown(zeroToNine)

	nineToTen := halfBack + 4.0
	p["10"] = p["9"].DrawRight(nineToTen)
	p["11"] = p["10"].SquareToHorizontalLine(scyeDepthLine)
	p["12"] = p["10"].SquareToHorizontalLine(topLine)

	twelveToThirteen := 1.5
	p["13"] = p["12"].DrawRight(twelveToThirteen)
	p["14"] = p["13"].DrawUp(2.0)

	tenToFifteen := 10.0
	p["15"] = p["10"].DrawLeft(tenToFifteen)

	tenToSixteen := 0.75
	p["16"] = p["10"].DrawDown(tenToSixteen)

	oneToSeventeen := (oneToFour / 2) + 0.5
	p["17"] = p["1"].DrawRight(oneToSeventeen)
	p["18"] = p["17"].SquareToHorizontalLine(waistline - 2.5)
	p["19"] = p["17"].SquareToHorizontalLine(hemline)

	fiveToTwenty := 4.5
	p["20"] = p["5"].DrawDown(fiveToTwenty)

	twentyToTwentyOne := (neck / 5) - 1.0
	p["21"] = p["20"].DrawLeft(twentyToTwentyOne)

	twentyToTwentyTwo := (neck / 5) - 2.5
	p["22"] = p["20"].DrawDown(twentyToTwentyTwo)

	tenToTwentyThree := 1.5
	p["23"] = p["10"].DrawDown(tenToTwentyThree)

	eightToFourteen := p["8"].DistanceTo(p["14"])
	twentyOneToTwentyFour := eightToFourteen + 0.5

	// Need to use pythagorean theorem to compute X value of 24
	twentyFourXFromTwentyOne := math.Sqrt(math.Pow(twentyOneToTwentyFour, 2.0) - math.Pow(p["23"].Y-p["20"].Y, 2.0))
	twentyFourX := p["21"].X - twentyFourXFromTwentyOne
	p["24"] = p["23"].SquareToVerticalLine(twentyFourX)

	oneToTwentyFive := (chest / 3) + 4.0
	p["25"] = p["1"].DrawRight(oneToTwentyFive)

	twentyFiveToTwentySix := 4.0
	p["26"] = p["25"].DrawUp(twentyFiveToTwentySix)

	p["27"] = p["24"].MidpointTo(p["26"])

	twentyTwoToTwentyEight := 1.5
	p["28"] = p["22"].DrawRight(twentyTwoToTwentyEight)

	twentyEightToTwentyNine := 3.5
	p["29"] = p["28"].DrawRight(twentyEightToTwentyNine)

	eighteenToThirty := 2.5
	p["30"] = p["18"].DrawRight(eighteenToThirty)

	eighteenToThirtyOne := 2.5
	p["31"] = p["18"].DrawLeft(eighteenToThirtyOne)

	nineteenToThirtyTwo := 8.0
	p["32"] = p["19"].DrawUp(nineteenToThirtyTwo)

	thirtyTwoToThirtyThree := 1.5
	p["33"] = p["32"].DrawRight(thirtyTwoToThirtyThree)

	thirtyTwoToThirtyFour := 1.5
	p["34"] = p["32"].DrawLeft(thirtyTwoToThirtyFour)

	p["35"] = p["19"].MidpointTo(p["6"])

	thirtyFiveToThirtySix := 3.0
	p["36"] = p["35"].DrawUp(thirtyFiveToThirtySix)

	p["37"] = p["3"].MidpointTo(p["19"])

	oneToEleven := p["1"].DistanceTo(p["11"])
	oneToThirtyEight := (oneToEleven / 2) + 2.0
	p["38"] = p["1"].DrawRight(oneToThirtyEight)

	thirtyEightToThirtyNine := 4.0
	p["39"] = p["38"].DrawDown(thirtyEightToThirtyNine)
	p["40"] = p["39"].SquareToHorizontalLine(waistline - 2.5)

	fortyToFortyOne := 16.0
	p["41"] = p["40"].DrawDown(fortyToFortyOne)

	return &tailoredShirtBlock{
		points: p,
	}
}

func (p *tailoredShirtBlock) GetPoints() map[string]geometry.Point {
	newMap := make(map[string]geometry.Point)

	for key, val := range p.points {
		newMap[key] = *val
	}

	return newMap
}

func (p *tailoredShirtBlock) PrintInstructions(w io.Writer) {
	fmt.Fprintln(w, "Square both ways from 0")

	p.inst(w, "0", "1", "square across")
	p.inst(w, "0", "2", "square across")
	p.inst(w, "0", "3", "square across")
	p.inst(w, "1", "4", "square up to 5 and down to 6 on hem line")
	p.inst(w, "0", "7", "square up")
	p.inst(w, "7", "8", "draw in neck curve")
	p.inst(w, "0", "9", "square out")
	p.inst(w, "9", "10", "square down to 11 on scye depth line and up to 12")
	p.inst(w, "12", "13", "square up 2 cm to 14. Join 8-14")
	p.inst(w, "10", "15", "")
	p.inst(w, "10", "16", "join 15-16 with slight curve")
	p.inst(w, "1", "17", "square down to 18, 2.5 cm below waistline and 19 on hemline")
	p.inst(w, "5", "20", "square out")
	p.inst(w, "20", "21", "")
	p.inst(w, "20", "22", "draw in neck curve")
	p.inst(w, "10", "23", "square out")
	p.inst(w, "21", "24", "")
	p.inst(w, "1", "25", "")
	p.inst(w, "25", "26", "join 24-26")
	p.inst(w, "24", "27", "")
	fmt.Fprintln(w, "Draw an armhole shape through points 14-10, and 16, 17, 26, 24; curve arm scye inwards 1 cm at 27.")
	p.inst(w, "22", "28", "button stand; square down")
	p.inst(w, "28", "29", "facing; square down; shape to edge at neckline")
	p.inst(w, "18", "30", "")
	p.inst(w, "18", "31", "")
	p.inst(w, "19", "32", "square across")
	p.inst(w, "32", "33", "")
	p.inst(w, "32", "34", "draw in curved seams")
	p.inst(w, "19", "35", "square up")
	p.inst(w, "35", "36", "square across to front edge")
	p.inst(w, "19", "37", "")

	fmt.Fprintln(w, "Draw shaped curves as shown from 33-36 and 34-37.")

	p.inst(w, "1", "38", "")

	fmt.Fprintln(w, "If added waist shaping is required, construct a dart in the back section:")

	p.inst(w, "38", "39", "square down to 40, 2.5 cm below waistline")
	p.inst(w, "40", "41", "draw a 1.5 cm dart on the line 39-41")
}

func (p *tailoredShirtBlock) inst(w io.Writer, start string, end string, additional string) {
	dist := p.points[start].DistanceTo(p.points[end])

	fmt.Fprintf(w, "%s-%s\t%.1f cm", start, end, dist)
	if additional != "" {
		fmt.Fprintf(w, "; %v", additional)
	}
	fmt.Fprintln(w, ".")
}

func (p *tailoredShirtBlock) GetLines() []geometry.Line {
	lines := []geometry.Line{}

	lines = append(lines, &geometry.StraightLine{
		Start: p.points["0"],
		End:   p.points["9"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["0"],
		End:   p.points["8"],
	})

	lines = append(lines, &geometry.StraightLine{
		Start: p.points["8"],
		End:   p.points["14"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["14"],
		End:   p.points["10"],
	})

	lines = append(lines, &geometry.StraightLine{
		Start: p.points["10"],
		End:   p.points["9"],
	})

	lines = append(lines, &geometry.StraightLine{
		Start: p.points["9"],
		End:   p.points["3"],
	})

	lines = append(lines, &geometry.StraightLine{
		Start: p.points["3"],
		End:   p.points["37"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["34"],
		End:   p.points["37"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["34"],
		End:   p.points["31"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["17"],
		End:   p.points["31"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["17"],
		End:   p.points["16"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["15"],
		End:   p.points["16"],
	})

	lines = append(lines, &geometry.StraightLine{
		Start: p.points["9"],
		End:   p.points["15"],
	})

	bottomRight := &geometry.Point{X: p.points["29"].X, Y: p.points["36"].Y}
	lines = append(lines, &geometry.StraightLine{
		Start: p.points["29"],
		End:   bottomRight,
	})

	lines = append(lines, &geometry.StraightLine{
		Start: bottomRight,
		End:   p.points["36"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["33"],
		End:   p.points["36"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["33"],
		End:   p.points["30"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["17"],
		End:   p.points["30"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["17"],
		End:   p.points["24"],
	})

	lines = append(lines, &geometry.StraightLine{
		Start: p.points["24"],
		End:   p.points["21"],
	})

	lines = append(lines, &geometry.CurvedLine{
		Start: p.points["22"],
		End:   p.points["21"],
	})

	lines = append(lines, &geometry.StraightLine{
		Start: p.points["22"],
		End:   p.points["29"],
	})

	return lines
}
