package main

import (
	"fmt"
	"github.com/tobyjsullivan/dxf"
	"github.com/tailored-style/pattern-generator/patterns/v3"
	"github.com/tailored-style/pattern-generator/patterns"
)

func main() {

	pattern := &v3.Pattern{}

	fmt.Println("Generating DXF...")
	d := dxf.NewDrawing()
	d.Header().LtScale = 1.0
	err := patterns.DrawDXF(pattern, d)
	if err != nil {
		panic(err.Error())
	}

	err = d.SaveAs("/Users/toby/sandbox/v3-out.dxf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done.")
}
