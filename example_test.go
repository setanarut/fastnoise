package fastnoise_test

import (
	"fmt"

	"github.com/setanarut/fastnoise"
)

func ExampleValue2D() {
	chars := " .:;+=xX$"
	noiseState := fastnoise.NewNoiseState[float32]()
	noiseState.Frequency = 0.05
	for y := range 10 {
		for x := range 80 {
			v := fastnoise.Value2D(x, y, noiseState)
			i := int((v + 1) / 2 * float32(len(chars)))
			fmt.Print(string(chars[i]))
		}
		fmt.Println()
	}
	// Output:
	// +=xxxxx==+++++=xXX$$XXx+;::..::;;++;:..  .:;+==xXXX$$$$$Xx=+:.   .:;=xXXXx=+;;:.
	// =xXXXXXx==+++==xXXXXXx=+;:.....::;::..   .:;+=xXXX$$$$$XXx=;:..  .:+=xXXxx=++;:.
	// xXX$$$XXx==++==xxXXxx=+;:.    ...::..    .:;=xXX$$$$$$XXxx=;:....:;+=xXXxx==+;:.
	// XX$$$$$Xxx======xxxx=+;::.     ......   ..;+=xX$$$$$XXXXx=+;::.::;+=xXXXXx==+;:.
	// X$$$$$$XXx===++=====++;::..   .......  ..:;+xX$$$$$XXXXxx=++;:::;+=xXXXXXx==+;:.
	// X$$$$$$XXxx=++++++++++;;::.......:......:;+=xX$$$$$XXXxxx==+;;;;+=xXX$XXXx=++;:.
	// XX$$$$$$Xxx=++;;;;;;+++;;;;::::::::::..::;+=xX$$$$XXXXXxxx==+++==xX$$$$Xxx=+;::.
	// xxXX$$$$XXx=+;;::::;;+++++++;;;;;;;:::::;++=xX$$$XXXXXXXXxx=====xX$$$$XXx=++;::.
	// +=xXX$$$$Xx=+;::..:;;+========++++;;;::;;++=xXXXXXXxxXXXXXXxxxxxXX$$$XXx=++;;:..
	// ;+=xX$$$$Xx=+;:...::;+=xxxxxxxx==++;;:::;++=xxXXxxxxxXXX$$XXXXXXXXXXXXx=+;;:::..
}

func ExampleValue3D() {
	chars := " .:;+=xX$"
	noiseState := fastnoise.NewNoiseState[float32]()
	noiseState.Frequency = 0.05
	for y := range 10 {
		for x := range 80 {
			v := fastnoise.Value3D(34, x, y, noiseState)
			i := int((v + 1) / 2 * float32(len(chars)))
			fmt.Print(string(chars[i]))
		}
		fmt.Println()
	}
	// Output:
	// XXXXXXXXx==++;;;;;;;;;;;;;++;;;:::::::;;;;++;;::......::;++===x====++;;;;;;++=xX
	// XXXX$$XXXxx=+++;;;;;;;;;;++++;;;;:::;;;;+++++;;;::::::::;++==xxx===++;;;;;;+==xx
	// XX$$$$$$XXxx==++++++++++++++++;;;;;;;;;+++++++;;;;;;;;;;;+==xxxxx==++;;;;;;+==xx
	// XX$$$$$$$XXx===+++++++++++++++;;;;;;;+++++++++++++++++++++==xxxxx==++;;;;;;++=xx
	// XX$$$$$$$XXxx===========++++++;;;;;+++++++++++++++++++++++==xxxxx==++;;;;;;++==x
	// xXX$$$$$$XXxx===========++++;;;;;;+++++=++++++=============xxxxxxx=++;;;;;;++===
	// xXX$$$$$XXxxxx=====xx===+++;;;;;;+++======++====xxxxxxxxx===xxxxxx==++;;;;;+++==
	// xxXXXXXXXXxxxxxxxxxxxx==++;;;:;;;;++===========xxxXXXXXxxx==xxxxxx==++;;;;;+++++
	// =xxxXXXXxxxxxxxxxxxxx===+;;:::::;;+====x======xxxXXXXXXXxxxxxxxxxx==++++;;;+++++
	// +==xxxxxxxxxxxxxxxxxx==++;::::::;;+==xxxxx===xxxXX$$$$XXXxxxxxxxxx==++++++++;;;;
}

func ExampleDomainWarp2D() {
	chars := " .:;+=xX$"
	noiseState := fastnoise.NewNoiseState[float32]()
	noiseState.Frequency = 0.05
	noiseState.DomainWarpType = fastnoise.DomainWarpOpenSimplex2Reduced
	noiseState.DomainWarpAmp = 30
	for y := range 10 {
		for x := range 80 {
			wx, wy := fastnoise.DomainWarp2D(x, y, noiseState)
			v := fastnoise.Value2D(wx, wy, noiseState)
			i := int((v + 1) / 2 * float32(len(chars)))
			fmt.Print(string(chars[i]))
		}
		fmt.Println()
	}
	// Output:
	// ++=xXX=+;;;+++==+;. :x$x;..    .:;;++;::....:;;. .;++==xXX$XX=+xX+. ;X$Xx+;;;;;:
	// =xxX$x+;;+++;;;;:..;xXx;;;+++;:....::..      .:. .;+====xxXxx==XX; .=X$Xx++;;;;:
	// XX$$X=+;+++++;;:::;xXx++xX$$XXx+;:...         .. .;==========xX$x:.;X$Xx=++;;;;;
	// X$Xx=+++====++;;;+=x=++x$XXxxxxx=+;:.    ...  .. .+xXXxx===xxX$X+.;x$Xx+;;;;;;;:
	// +=++;++=xx===++;+=x=;+xx=+;;;;+++=+;:............:=X$$$XXXXX$$X=;;x$X=;:..::::::
	// .::;+=xxxxx==++++==;;==+::;++;;:;+=+;::::::::::..+X$$$$$$$$$$X=++xXx;:.    ..:::
	// :;+=xxxxxxxx==+++=+:;++::;===++::;+=++;;;;;;;;:::x$$XXXxxxxxx=+=xX=;::::........
	// =xXXXXXXx======++=;:;+:.:=xxx=+;::;+=++++++++;::+X$XXxx====++=xXX=++=xxx=+;::...
	// X$$$$$Xx=++++==++=;.:;..;xXXxx=+:.:;++++++++;;:+xXXXXx=======xXXx++xX$$XXx=+;::.
	// $$$$$$Xx=;;;;=====;.::..+xXXxx=+;..::;;;;;;;::;=XxX$XxxxXXXXXXXx==xX$$$XXx===+;;
}
func ExampleDomainWarp3D() {
	chars := " .:;+=xX$"
	noiseState := fastnoise.NewNoiseState[float32]()
	noiseState.Frequency = 0.05
	noiseState.DomainWarpType = fastnoise.DomainWarpOpenSimplex2Reduced
	noiseState.DomainWarpAmp = 30
	for y := range 10 {
		for x := range 80 {
			wx, wy, wz := fastnoise.DomainWarp3D(y, x, 300, noiseState)
			v := fastnoise.Value3D(wx, wy, wz, noiseState)
			i := int((v + 1) / 2 * float32(len(chars)))
			fmt.Print(string(chars[i]))
		}
		fmt.Println()
	}
	// Output:
	// ......:;;:.  .::;;++=xxx=++;;:::::::::::::::;++;:::+:.   .::;;::.:;=xXx+:.   ..:
	// :::.. .:;:.   ..:;+=xXXx=+:::::::::::::::::::;+;;::;;.    .::::..:+=xxx+:.   .::
	// :::.. .:+;:....:;+=xXXXx=;::::::::::::::::::::;+;::;;:    ..::.::;+=xxx+:.....:;
	// :::....:;;;:::::;=xXXXx=;;::::::::::::::::::::;+;;;;;:. ...::::;;+=xxx=;:....::;
	// :::....::;;;:::;+=xXXx=+;;::...:::::::::::::.::+;;;+;:....::++++==xxx=+;::.:::;+
	// ::::..:::;;:::::;+====+++;::....:::::::::::...:+;++++;:..:;+++==xxxx==;:::::::;+
	// ;:::::::::::....::;+++++++;:....:::::::::::..::+++==+;::::;++==xxXxx=;;:::::::;+
	// ;;;;;;;;:::..    .::;;+==+;::...::::::::::::::;++====+;:::;++==xxXx=+;::::::::;+
	// +++;;;;;;:::..    .::;====+;::..:::;;;:::::::;+==xx==;:::;;;++=xxx=+;;;;;:::::;+
	// +++++++;;;;;::..  ..:;==xx=+;:::::;;;;+++;;++==xxx==;:::::::;;++==+;;;;++;;;:::+
}
