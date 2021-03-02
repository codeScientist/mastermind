package mastermindlib

type Response struct {
	black_pins, white_pins, index int
}

func Create_Response(black_pins int, white_pins int, index int) Response {
	r := Response{black_pins, white_pins, index}
	return r
}

func (r *Response) Get_index() int {
	return (*r).index
}

func Cal_total_responses(total_pins int, total_selectable_pins int, is_single_use bool) int {
	/*
		total pins = 6, total selectable pins = 4
		black pins, white pins
		0 - 0,1,2,3,4
		1 - 0,1,2,3
		2 - 0,1,2
		3 - 0,1
		4 - 0

		from this 3,1 is not possible

		also if only single use is allowed then:
		0 - 0,1
		1 - 0
		is not possible because minimum white + black pins have to be 2
	*/
	total := (((total_selectable_pins + 2) * (total_selectable_pins + 1)) / 2) - 1
	if is_single_use {
		min_white := (2 * total_selectable_pins) - total_pins // minimum white + black pins present in any sequence
		if min_white > 0 {
			total = total - (((min_white + 1) * min_white) / 2)
		}
	}
	return total
}

func (r *Response) Set_response_and_index(gs *GameState, to_set_index bool) {
	/*
		0 - 0,1,2,3,4
		1 - 0,1,2,3
		2 - 0,1,2
		3 - 0,1
		4 - 0
		(0,0) will have index = 0
		(0,1) will have index = 1
		(1,0) will have index = 5
		and so on
	*/
	index := 0
	for black_pins := 0; black_pins <= (*gs).Total_selectable_pins; black_pins++ {
		// set max white pins
		max_white_pins := (*gs).Total_selectable_pins - black_pins
		if max_white_pins == 1 {
			max_white_pins = 0
		}

		// set min white pins
		min_white_pins := 0
		if (*gs).Is_single_use {
			min_white_pins = (2 * ((*gs).Total_selectable_pins)) - (*gs).Total_pins - black_pins
			if min_white_pins < 0 {
				min_white_pins = 0
			}
		}

		for white_pins := min_white_pins; white_pins <= max_white_pins; white_pins++ {
			// set white pins and black pins OR the index according to to_set_index
			if to_set_index && white_pins == (*r).white_pins && black_pins == (*r).black_pins {
				(*r).index = index
				return
			} else if !to_set_index && index == (*r).index {
				(*r).white_pins = white_pins
				(*r).black_pins = black_pins
				return
			}
			index++
		}
	}
}
