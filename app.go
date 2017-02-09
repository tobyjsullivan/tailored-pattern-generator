package main

import (
	"fmt"
	"github.com/tailored-style/pattern-generator/patternfile"
	"github.com/tailored-style/pattern-generator/styles"
	"github.com/tailored-style/pattern-generator/pieces"
)

func main() {
	sample42In := &pieces.Measurements{
		ChestCircumference: 106.7, // 42"
		WaistCircumference: 91.4, // 36"
		HipCircumference: 109.2, // 43"
		NeckCircumference: 41.9, // 16 1/2"
		Height: 182.9, // 72"
	}

	//personal := &pieces.Measurements{
	//	ChestCircumference: 110.0,
	//	WaistCircumference: 96.5,
	//	HipCircumference: 110.5,
	//	NeckCircumference: 43.0,
	//	Height: 182.0,
	//}

	style := &styles.SN11001Shirt{
		Measurements: sample42In,
	}

	fmt.Println("Generating DXF...")
	pf := patternfile.NewPatternFile()
	err := pf.DrawPattern(style)
	if err != nil {
		panic(err.Error())
	}

	err = pf.SaveAs("/Users/toby/sandbox/v3-out.dxf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done.")
}
