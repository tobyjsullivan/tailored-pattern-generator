package main

import (
	"fmt"
	"github.com/tailored-style/pattern-generator/patternfile"
	"github.com/tailored-style/pattern-generator/styles"
)

func main() {
	style := &styles.SN11001Shirt{}

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
