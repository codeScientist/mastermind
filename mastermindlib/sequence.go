package mastermindlib

import (
	"errors"
	"fmt"
	"math/rand"
)

type Sequence struct {
	sarray []int
}

func Create_Sequence(gs *GameState) Sequence {
	s := Sequence{sarray: make([]int, (*gs).Total_selectable_pins)}
	return s
}

func (s *Sequence) Print_Sequence() {
	fmt.Printf("%v", (*s).sarray)
}

func Cal_total_sequences(total_pins int, total_selectable_pins int, is_single_use bool) int {
	total := 1
	if is_single_use {
		for i := 0; i < total_selectable_pins; i++ {
			total = total * (total_pins - i)
		}
	} else {
		for i := 0; i < total_selectable_pins; i++ {
			total = total * total_pins
		}
	}
	return total
}

func (s *Sequence) Set_lowest_seq(gs *GameState) {
	if (*gs).Is_single_use {
		for i := range (*s).sarray {
			(*s).sarray[i] = i + 1
		}
	} else {
		for i := range (*s).sarray {
			(*s).sarray[i] = 1
		}
	}
}

func (s *Sequence) Set_seq(s2 []int) {
	for i := range (*s).sarray {
		(*s).sarray[i] = s2[i]
	}
}

func (s *Sequence) Copy_seq() Sequence {
	new_seq := Sequence{sarray: make([]int, len((*s).sarray))}
	for i, e := range (*s).sarray {
		new_seq.sarray[i] = e
	}
	return new_seq
}

func (s *Sequence) increment_seq(gs *GameState, index int) error {
	if index < 0 {
		return errors.New("cannot increment the sequence further")
	}

	(*s).sarray[index]++

	if (*s).sarray[index] > (*gs).Total_pins {
		(*s).sarray[index] = 1
		return s.increment_seq(gs, index-1)
	}

	return nil
}

func (s *Sequence) Gen_next_seq(gs *GameState) (Sequence, error) {
	next_seq := s.Copy_seq()
	e := (&next_seq).increment_seq(gs, len((*s).sarray)-1)
	for (*gs).Is_single_use && !(&next_seq).Is_single_use_seq_valid() && e == nil {
		e = (&next_seq).increment_seq(gs, len((*s).sarray)-1)
	}
	return next_seq, e
}

func (s *Sequence) Set_random_seq(gs *GameState) {
	if (*gs).Is_single_use {
		bag := make([]int, (*gs).Total_selectable_pins)
		for i := range bag {
			bag[i] = i + 1
		}
		for i := range (*s).sarray {
			bag_index := rand.Intn(len(bag))
			(*s).sarray[i] = bag[bag_index]
			bag[bag_index] = bag[len(bag)-1]
			bag = bag[:len(bag)-1]
		}
	} else {
		for i := range (*s).sarray {
			(*s).sarray[i] = rand.Intn((*gs).Total_pins) + 1
		}
	}
}

func (s *Sequence) Is_single_use_seq_valid() bool {
	for i := 0; i < (len((*s).sarray) - 1); i++ {
		for j := i + 1; j < len((*s).sarray); j++ {
			if (*s).sarray[i] == (*s).sarray[j] {
				return false
			}
		}
	}
	return true
}

func (s *Sequence) Evaluate_seq(gs *GameState, s2 *Sequence) Response {
	res := Response{0, 0, -1}
	for _, ie := range (*s).sarray {
		for j, je := range (*s2).sarray {
			if ie == je {
				(*s2).sarray[j] = -1 * je
				res.white_pins++
				break
			}
		}
	}
	// correct the negative values in s2
	for i, e := range (*s2).sarray {
		if e < 0 {
			(*s2).sarray[i] = -1 * e
		}
	}

	for i, ie := range (*s).sarray {
		if ie == (*s2).sarray[i] {
			res.white_pins--
			res.black_pins++
		}
	}

	(&res).Set_response_and_index(gs, true)
	return res
}
