package main

import (
	"fmt"
	"github.com/codeScientist/mastermind/mastermindlib"
)

const default_total_pins = 6
const default_total_selectable_pins = 4
const default_is_single_use = true

func main() {
	game_state, _ := mastermindlib.Create_GameState(default_total_pins, default_total_selectable_pins, default_is_single_use)

	fmt.Println("Generating Guess Tree...")
	gtn := mastermindlib.Get_GuessTree(&game_state)
	gods_num := gtn.Cal_height()
	fmt.Printf("Current configuration:\ntotal pins: %d\ntotal select pins: %d\nsingle use: %t\ngods num: %d\n\n", default_total_pins, default_total_selectable_pins, default_is_single_use, gods_num)

	for {
		if gtn == nil {
			fmt.Println("No guess possible: INCORRECT INPUT")
			break
		}
		fmt.Printf("Computer's guess:\n")
		(*gtn).Print_Guess()
		fmt.Printf("\n\n")
		var bp, wp int
		fmt.Printf("Please enter number of black pins: ")
		fmt.Scanf("%d", &bp)
		fmt.Printf("Please enter number of white pins: ")
		fmt.Scanf("%d", &wp)
		r := mastermindlib.Create_Response(bp, wp, 0)
		r.Set_response_and_index(&game_state, true)

		if bp == game_state.Total_selectable_pins {
			fmt.Println("Computer WON")
			break
		}

		gtn = (*gtn).Get_next_GuessTreeNode((&r).Get_index())
	}

	// mode := true
	// for tp := 0; tp <= 6; tp++ {
	// 	for tsp := 1; tsp <= tp; tsp++ {
	// 		game_state, _ := mastermindlib.Create_GameState(tp, tsp, mode)
	// 		gtn := mastermindlib.Get_GuessTree(&game_state)
	// 		gods_num := gtn.Cal_height()
	// 		fmt.Printf("Current configuration:\ntotal pins: %d\ntotal select pins: %d\nsingle use: %t\ngods num: %d\n\n", tp, tsp, mode, gods_num)
	// 	}
	// }
}
