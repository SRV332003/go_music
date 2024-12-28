package main

import (
	"fmt"

	"github.com/SRV332003/go_music/handler"
	"github.com/SRV332003/go_music/player"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

func main() {

	handler.ClearScr()
	fmt.Print(color.CyanString("dhvani > "))
	player.StopIfPlaying()

	for {

		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		if char == ':' {

			fmt.Printf(color.HiRedString("\rdhvani >> "))
			handler.HandleWordInputs()
			fmt.Printf(color.CyanString("dhvani > "))

		} else {

			if handler.HandleCommand(char, key) {
				fmt.Printf(color.CyanString("dhvani > "))
				continue
			} else {
				fmt.Printf(color.CyanString("\rdhvani > "))
			}
		}
	}

}
