package mastermindlib

import "errors"

type GameState struct {
	Total_pins, Total_selectable_pins, Total_responses, Total_sequences int
	Is_single_use                                                       bool
}

func Create_GameState(total_pins int, total_selectable_pins int, is_single_use bool) (GameState, error) {
	if is_single_use && total_pins < total_selectable_pins {
		return GameState{total_pins, total_selectable_pins, 0, 0, is_single_use}, errors.New("this configuration is not possible, total_pins cannot be less than total_selectable_pins in single_use settings")
	}
	total_responses := Cal_total_responses(total_pins, total_selectable_pins, is_single_use)
	total_sequences := Cal_total_sequences(total_pins, total_selectable_pins, is_single_use)
	return GameState{total_pins, total_selectable_pins, total_responses, total_sequences, is_single_use}, nil
}
