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
	//torsoBlock32S := &patterns.TorsoSloper{
	//	ChestCircumference: 81.3, // 32
	//	BackInterscyeLength: 36.8, // 14 1/2
	//	ShoulderToShoulder: 40.0, // 15 3/4
	//	ArmLength: 58.4, // 23
	//	BicepCircumference: 26.7, // 10 1/2
	//	Height: 152.4, // 60
	//}
	//torsoBlock38R := &patterns.TorsoSloper{
	//	ChestCircumference: 96.5, // 38
	//	BackInterscyeLength: 40.6, // 16
	//	ShoulderToShoulder: 43.8, // 17 1/4
	//	ArmLength: 63, // 24 7/8
	//	BicepCircumference: 32.4, // 12 3/4
	//	Height: 172.7, // 5' 8" = 68
	//}
	torsoBlock42R := &patterns.TorsoSloper{
		ChestCircumference: 106.7,
		BackInterscyeLength: 43.2,
		ShoulderToShoulder: 46.4,
		ArmLength: 63.8,
		BicepCircumference: 36.2,
		Height: 182.9,
	}

	//torsoBlock50T := &patterns.TorsoSloper{
	//	ChestCircumference: 127, // 50
	//	BackInterscyeLength: 48.3, // 19
	//	ShoulderToShoulder: 51.4, // 20 1/4
	//	ArmLength: 68.9, // 27 1/8
	//	BicepCircumference: 43.8, // 17 1/4
	//	Height: 208.3, // 82
	//}

	layers := []patterns.Pattern{
		//torsoBlock32S,
		//torsoBlock38R,
		torsoBlock42R,
		//torsoBlock50T,
	}

	d2 := dxf.NewDrawing()
	d2.Header().LtScale = 100.0

	for _, p := range layers {
		err = patterns.DrawDXF(p, d2)
		if err != nil {
			panic(err.Error())
		}
	}

	err = d2.SaveAs("/Users/toby/sandbox/test-torso.dxf")
	if err != nil {
		panic(err.Error())
	}

	v3Torso := patterns.NewV3Torso(&patterns.V3Measurements{
		ChestCircumference: 106.7,
		Height: 182.9,
		WaistCircumference: 106.7,
		NeckCircumference: 40.0,
		HipCircumference: 106.7,
		BiceptCircumference: 36.2,
	})

	d3 := dxf.NewDrawing()
	d3.Header().LtScale = 100.0
	v3Torso.DrawDXF(d3)
	d3.SaveAs("/Users/toby/sandbox/v3-out.dxf")

	fmt.Println("Done.")
}
