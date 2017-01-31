package main

import (
	"fmt"
	"github.com/tailored-style/pattern-generator/patterns"
	"github.com/tobyjsullivan/dxf"
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

	fmt.Println("Generating DXF...")
	d := dxf.NewDrawing()
	d.Header().LtScale = 100.0
	err := patterns.DrawDXF(block, d)
	if err != nil {
		panic(err.Error())
	}

	err = d.SaveAs("/Users/toby/sandbox/test-tailored.dxf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Generating TORSO SLOPER DXF...")
	torsoBlock := &patterns.TorsoSloper{
		ChestCircumference: 106.7,
		BackInterscyeLength: 43.2,
		ShoulderToShoulder: 46.4,
		ArmLength: 63.8,
		BicepCircumference: 36.2,
		Height: 182.9,
	}
	d2 := dxf.NewDrawing()
	d2.Header().LtScale = 100.0
	err = patterns.DrawDXF(torsoBlock, d2)
	if err != nil {
		panic(err.Error())
	}

	err = d2.SaveAs("/Users/toby/sandbox/test-torso.dxf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done.")
}
