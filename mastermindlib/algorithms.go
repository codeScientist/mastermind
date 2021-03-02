package mastermindlib

import (
	// "fmt"
	"errors"
)

func Gen_minimax_next_guess(gs *GameState, root *GuessListNode, possible_correct_seqs []Sequence, equivalent_seqs []Sequence) (Sequence, error) {
	// first we need to get the possible_correct_seqs
	if possible_correct_seqs == nil {
		possible_correct_seqs = get_all_possible_correct_seqs(gs, root)
	}

	if len(possible_correct_seqs) == 0 {
		return Create_Sequence(gs), errors.New("given guesses and their responses is not possible")
	}

	if len(possible_correct_seqs) == 1 {
		return possible_correct_seqs[0], nil
	}

	// fmt.Printf("possible seqs: %v\n", possible_correct_seqs)

	/*
		now we need to find which guess sequence most evenly divides
		the possible_correct_seq across all the responses
		this is a heavy operation
		we may be able to optimize it by eliminating equivalent sequences

		for the first guess you only need to check for 5 seqs

		seq		count	num_colors	color_dist_index	total_dist
		1111	6		1			1					1
		1112	120		2			1					4
		1122	90		2			2					3
		1233	720		3			1					6
		1234	360		4			1					1
	*/

	if equivalent_seqs == nil {
		equivalent_seqs = get_equivalent_seqs(gs, root)
	}

	// fmt.Printf("equivalent seqs: %v\n", equivalent_seqs)

	var all_res []Response
	for i := 0; i < (*gs).Total_responses; i++ {
		res := Response{0, 0, i}
		(&res).Set_response_and_index(gs, false)
		all_res = append(all_res, res)
	}

	// fmt.Printf("possible res: %v\n", all_res)

	var best_guess_seq Sequence
	var best_divide int
	best_divide = -1

	// var distribution map[int][]Sequence
	// distribution = make(map[int][]Sequence)

	for _, guess_seq := range equivalent_seqs {
		divide := get_max_division(gs, &guess_seq, all_res, possible_correct_seqs)

		// _, ok := distribution[divide]
		// if !ok {
		// 	distribution[divide] = []Sequence{guess_seq}
		// } else {
		// 	distribution[divide] = append(distribution[divide], guess_seq)
		// }

		// fmt.Printf("divide: %d for seq %v\n", divide, guess_seq.sarray)
		if best_divide == -1 || divide < best_divide {
			best_divide = divide
			best_guess_seq = guess_seq
		}
	}
	// fmt.Printf("distribution: \n%v\n", distribution)

	// fmt.Printf("best seq: %v, best divide: %d\n", best_guess_seq.sarray, best_divide)
	return best_guess_seq, nil
}

func get_max_division(gs *GameState, guess_seq *Sequence, all_res []Response, possible_correct_seq []Sequence) int {
	max := 0
	for _, res := range all_res {
		count := 0
		gln := Create_GuessListNode(guess_seq, &res)
		for _, s := range possible_correct_seq {
			if (&gln).Check_seq(gs, &s) {
				// fmt.Printf("s: %v\n", s.sarray)
				count++
			}
		}

		// fmt.Printf("count: %d for seq: %v and res: %+v\n", count, (*guess_seq).sarray, res)

		if count > max {
			max = count
		}
	}

	return max
}

func get_equivalent_seqs(gs *GameState, root *GuessListNode) []Sequence {
	var equivalent_seqs []Sequence

	if root.Cal_count_guessListNode() == 0 {
		if (*gs).Is_single_use {
			seq := Create_Sequence(gs)
			(&seq).Set_lowest_seq(gs)
			equivalent_seqs = append(equivalent_seqs, seq)
		} else {
			max_colors := get_max_color_pegs(gs)
			for i := 1; i <= max_colors; i++ {
				combinations := get_color_combinations(gs, i)
				for _, e := range combinations {
					s := gen_seq_from_color_comb(gs, e)
					equivalent_seqs = append(equivalent_seqs, s)
				}
			}
		}
	} else {
		var e error
		e = nil
		seq := Create_Sequence(gs)
		(&seq).Set_lowest_seq(gs)
		for e == nil {
			equivalent_seqs = append(equivalent_seqs, seq)
			seq, e = (&seq).Gen_next_seq(gs)
		}
	}

	return equivalent_seqs
}

func gen_seq_from_color_comb(gs *GameState, comb []int) Sequence {
	seq := Create_Sequence(gs)
	num := 1
	index := 0
	for _, e := range comb {
		for j := 0; j < e; j++ {
			seq.sarray[index] = num
			index++
		}
		num++
	}
	return seq
}

func get_max_color_pegs(gs *GameState) int {
	if (*gs).Total_pins > (*gs).Total_selectable_pins {
		return (*gs).Total_selectable_pins
	} else {
		return (*gs).Total_pins
	}
}

func get_color_combinations(gs *GameState, num_colors int) [][]int {
	var queue [][]int
	var combinations [][]int

	// get lowest combination
	total_sum := (*gs).Total_selectable_pins
	comb := make([]int, num_colors)
	for i := range comb {
		comb[i] = 1
	}
	comb[0] = total_sum - num_colors + 1

	queue = append(queue, comb)

	for len(queue) > 0 {
		comb = queue[0]
		queue = queue[1:]
		combinations = append(combinations, comb)
		for i := 0; i < len(comb)-1; i++ {
			if comb[i]-comb[i+1] >= 2 {
				new_comb := make([]int, num_colors)
				for j := range new_comb {
					new_comb[j] = comb[j]
				}
				new_comb[i]--
				new_comb[i+1]++
				queue = append(queue, new_comb)
			}
		}
	}

	return combinations
}

func get_all_possible_correct_seqs(gs *GameState, root *GuessListNode) []Sequence {
	var all_possible_correct_seqs []Sequence
	var e error
	e = nil
	seq := Create_Sequence(gs)
	(&seq).Set_lowest_seq(gs)

	for e == nil {
		if root.Check_seq(gs, &seq) {
			all_possible_correct_seqs = append(all_possible_correct_seqs, seq)
		}
		seq, e = (&seq).Gen_next_seq(gs)
	}

	return all_possible_correct_seqs
}
