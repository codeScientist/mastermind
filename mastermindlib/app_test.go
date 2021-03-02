package mastermindlib

import (
	"testing"
)

func Test_Cal_total_responses(t *testing.T) {
	default_total_pins := 6
	default_total_selectable_pins := 4
	default_is_single_use := true

	correct_total_responses := 11
	total_responses := Cal_total_responses(default_total_pins, default_total_selectable_pins, default_is_single_use)
	if total_responses != correct_total_responses {
		t.Errorf("Cal_total_responses(%d, %d, %t) = %d; want %d", default_total_pins, default_total_selectable_pins, default_is_single_use, total_responses, correct_total_responses)
	}

	default_is_single_use = false
	correct_total_responses = 14
	total_responses = Cal_total_responses(default_total_pins, default_total_selectable_pins, default_is_single_use)
	if total_responses != correct_total_responses {
		t.Errorf("Cal_total_responses(%d, %d, %t) = %d; want %d", default_total_pins, default_total_selectable_pins, default_is_single_use, total_responses, correct_total_responses)
	}
}

func Test_Set_response_and_index(t *testing.T) {
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
	default_total_pins := 6
	default_total_selectable_pins := 4
	default_is_single_use := false

	game_state := Create_GameState(default_total_pins, default_total_selectable_pins, default_is_single_use)

	res := Response{black_pins: 2, white_pins: 2, index: -1}
	res.Set_response_and_index(&game_state, true)
	correct_res := Response{black_pins: 2, white_pins: 2, index: 11}
	if res.index != correct_res.index {
		init_res := Response{black_pins: 2, white_pins: 2, index: -1}
		t.Errorf("game state = %+v\nres = %+v\nres.Set_response_and_index(&game_state, true) = %+v; want %+v", game_state, init_res, res, correct_res)
	}
}

func Test_Cal_total_sequences(t *testing.T) {
	default_total_pins := 6
	default_total_selectable_pins := 4
	default_is_single_use := true

	correct_total_sequences := 360
	total_sequences := Cal_total_sequences(default_total_pins, default_total_selectable_pins, default_is_single_use)
	if total_sequences != correct_total_sequences {
		t.Errorf("Cal_total_sequences(%d, %d, %t) = %d; want %d", default_total_pins, default_total_selectable_pins, default_is_single_use, total_sequences, correct_total_sequences)
	}

	default_is_single_use = false
	correct_total_sequences = 1296
	total_sequences = Cal_total_sequences(default_total_pins, default_total_selectable_pins, default_is_single_use)
	if total_sequences != correct_total_sequences {
		t.Errorf("Cal_total_sequences(%d, %d, %t) = %d; want %d", default_total_pins, default_total_selectable_pins, default_is_single_use, total_sequences, correct_total_sequences)
	}
}

func Test_Set_lowest_seq(t *testing.T) {
	default_total_pins := 6
	default_total_selectable_pins := 4
	default_is_single_use := true
	game_state := Create_GameState(default_total_pins, default_total_selectable_pins, default_is_single_use)

	seq := Create_Sequence(&game_state)
	(&seq).Set_lowest_seq(&game_state)
	correct_seq := Create_Sequence(&game_state)
	correct_seq.sarray = []int{1, 2, 3, 4}

	seq_match := true
	for i, e := range seq.sarray {
		if correct_seq.sarray[i] != e {
			seq_match = false
			break
		}
	}

	if !seq_match {
		t.Errorf("game state = %+v\n(&seq).Set_lowest_seq(&game_state) = %v; want %v", game_state, seq, correct_seq)
	}
}

func Test_Evaluate_seq(t *testing.T) {
	default_total_pins := 6
	default_total_selectable_pins := 4
	default_is_single_use := false

	game_state := Create_GameState(default_total_pins, default_total_selectable_pins, default_is_single_use)
	var s1, s2 Sequence
	var res, correct_res Response
	s1 = Create_Sequence(&game_state)
	s2 = Create_Sequence(&game_state)

	s1.sarray = []int{1, 2, 1, 2}
	s2.sarray = []int{2, 1, 1, 2}
	res = (&s1).Evaluate_seq(&game_state, &s2)
	correct_res = Response{black_pins: 2, white_pins: 2, index: 11}
	if res.black_pins != correct_res.black_pins || res.white_pins != correct_res.white_pins || res.index != correct_res.index {
		t.Errorf("game state = %+v\ns1 = %v\ns2 = %v\n(&s1).Evaluate_seq(&game_state, &s2) = %+v; want %+v", game_state, s1.sarray, s2.sarray, res, correct_res)
	}
}

func Test_Gen_next_seq(t *testing.T) {
	default_total_pins := 6
	default_total_selectable_pins := 4
	default_is_single_use := false

	game_state := Create_GameState(default_total_pins, default_total_selectable_pins, default_is_single_use)

	seq := Create_Sequence(&game_state)
	seq.sarray = []int{1, 2, 1, 2}

	next_seq, e := (&seq).Gen_next_seq(&game_state)

	var correct_next_seq = Create_Sequence(&game_state)
	correct_next_seq.sarray = []int{1, 2, 1, 3}

	seq_match := true
	for i, e := range next_seq.sarray {
		if correct_next_seq.sarray[i] != e {
			seq_match = false
			break
		}
	}

	if !seq_match || e != nil {
		t.Errorf("game_state = %+v\nseq = %v\nseq.Gen_next_seq(&game_state) = %v; want %v", game_state, seq.sarray, next_seq.sarray, correct_next_seq.sarray)
	}
}

func Test_Is_single_use_seq_valid(t *testing.T) {
	default_total_pins := 6
	default_total_selectable_pins := 4
	default_is_single_use := true

	game_state := Create_GameState(default_total_pins, default_total_selectable_pins, default_is_single_use)

	seq := Create_Sequence(&game_state)
	seq.sarray = []int{1, 2, 1, 2}

	res := (&seq).Is_single_use_seq_valid()
	correct_res := false
	if res != correct_res {
		t.Errorf("game_state = %+v\nseq = %v\nseq.Is_single_use_seq_valid() = %t; want %t", game_state, seq.sarray, res, correct_res)
	}
}

// func Test_Gen_minimax_next_guess(t *testing.T) {
// 	default_total_pins := 6
// 	default_total_selectable_pins := 4
// 	default_is_single_use := true

// 	game_state := Create_GameState(default_total_pins, default_total_selectable_pins, default_is_single_use)

// 	Gen_minimax_next_guess(&game_state, nil, nil)
// }
