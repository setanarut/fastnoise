package fastnoise_test

import (
	"fmt"

	"github.com/setanarut/fastnoise"
)

func ExampleState_Noise2D() {
	chars := " .:;+=xX$"
	noise := fastnoise.New[float32]()
	noise.Frequency = 0.05
	for y := range 15 {
		for x := range 100 {
			v := noise.Noise2D(x, y)
			i := int((v + 1) / 2 * float32(len(chars)))
			fmt.Print(string(chars[i]))
		}
		fmt.Println()
	}
	// Output:
	// +=xxxxx==+++++=xXX$$XXx+;::..::;;++;:..  .:;+==xXXX$$$$$Xx=+:.   .:;=xXXXx=+;;:.. ...:;;+++===xxXXXx
	// =xXXXXXx==+++==xXXXXXx=+;:.....::;::..   .:;+=xXXX$$$$$XXx=;:..  .:+=xXXxx=++;:..  ..::;++==xxXXXXXX
	// xXX$$$XXx==++==xxXXxx=+;:.    ...::..    .:;=xXX$$$$$$XXxx=;:....:;+=xXXxx==+;:..   ..:;+==xxXX$$$$X
	// XX$$$$$Xxx======xxxx=+;::.     ......   ..;+=xX$$$$$XXXXx=+;::.::;+=xXXXXx==+;:.    ..:;+=xXXX$$$$$$
	// X$$$$$$XXx===++=====++;::..   .......  ..:;+xX$$$$$XXXXxx=++;:::;+=xXXXXXx==+;:.    .:;+=xXXX$$$$$$$
	// X$$$$$$XXxx=++++++++++;;::.......:......:;+=xX$$$$$XXXxxx==+;;;;+=xXX$XXXx=++;:.    .:;+xxX$$$$$$$$$
	// XX$$$$$$Xxx=++;;;;;;+++;;;;::::::::::..::;+=xX$$$$XXXXXxxx==+++==xX$$$$Xxx=+;::.   ..:+=xX$$$$$$$$$$
	// xxXX$$$$XXx=+;;::::;;+++++++;;;;;;;:::::;++=xX$$$XXXXXXXXxx=====xX$$$$XXx=++;::.. ..:;+=XX$$$$$XXXXX
	// +=xXX$$$$Xx=+;::..:;;+========++++;;;::;;++=xXXXXXXxxXXXXXXxxxxxXX$$$XXx=++;;:.....::;=xX$$$$XXXXXXX
	// ;+=xX$$$$Xx=+;:...::;+=xxxxxxxx==++;;:::;++=xxXXxxxxxXXX$$XXXXXXXXXXXXx=+;;:::....::;+=xX$$$XXxxxxxx
	// :;+=xX$$$Xx=+:.. ..:;=xxXXXXXXxx==+;;:::;;+==xxxx==xxxXX$$$$XXXXXXXXxx=+;::::..::::;++=xXXXXXx====xx
	// .:;+xX$$$Xx=;:.   .:;=xXX$$$$XXx=++;::.:::;++========xXX$$$$$XXXXxxx==+;::....:::;;++=xxXXXxx==++==x
	//  .:+=X$$$Xx=;:.   .:;+=xX$$$$$Xx=+;:.....:;;++++++++==xXX$$$$XXXxx==+;;:......:;;+++==xxxxx==+++++=x
	//  .:+=xX$XXx+;:.   ..:+=xxXX$$XXx=;:..   ..:;;;;;;;;;++=xXX$$XXXxx==+;;:.    .:;++===xxxxx==+;;;;;+=x
	// ..:;=xXXXx=+;:.    .:;+==xxXXXx=+;:.    ..::;;;;::::;;+=xxXXXXxx==+;;:..   .:;+==xxxxxxx==+;::::;+=x
}
