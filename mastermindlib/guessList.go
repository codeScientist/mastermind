package mastermindlib

import (
	"errors"
	"fmt"
)

type GuessListNode struct {
	sequence Sequence
	response Response
	next     *GuessListNode
}

func Create_GuessListNode(s *Sequence, r *Response) GuessListNode {
	gln := GuessListNode{sequence: *s, response: *r, next: nil}
	return gln
}

func (root *GuessListNode) Copy_GuessList() *GuessListNode {
	if root == nil {
		return nil
	}
	cp_seq := (*root).sequence.Copy_seq()
	cp_res := Response{black_pins: (*root).response.black_pins, white_pins: (*root).response.white_pins, index: (*root).response.index}
	gln := Create_GuessListNode(&cp_seq, &cp_res)
	gln.next = (*root).next.Copy_GuessList()
	return &gln
}

func (root *GuessListNode) Add_GuessListNode(gln *GuessListNode) error {
	if root == nil {
		return errors.New("no guess list to add guess list node to")
	}
	if (*root).next == nil {
		(*root).next = gln
		return nil
	}
	return (*root).next.Add_GuessListNode(gln)
}

func (root *GuessListNode) Get_last_guessListNode() (*GuessListNode, error) {
	if root == nil {
		return root, errors.New("no guesses in the guess list, no last guess to return")
	}
	if (*root).next == nil {
		return root, nil
	}
	return (*root).next.Get_last_guessListNode()
}

func (root *GuessListNode) Cal_count_guessListNode() int {
	if root == nil {
		return 0
	}

	return 1 + (*root).next.Cal_count_guessListNode()
}

func (root *GuessListNode) Check_seq(gs *GameState, seq *Sequence) bool {
	if root == nil {
		return true
	}
	res := (&(*root).sequence).Evaluate_seq(gs, seq)
	if res.index != (*root).response.index {
		return false
	}
	return (*root).next.Check_seq(gs, seq)
}

func (root *GuessListNode) Print_guess_list() {
	if root == nil {
		fmt.Printf("No guesses\n")
		return
	}
	// header of the guess list
	fmt.Printf("+----+--------------------+-------+-------+\n")
	fmt.Printf("|  # |              guess | black | white |\n")
	fmt.Printf("|----|--------------------|-------|-------|\n")

	currentNode := root
	index := 0
	for {
		index = index + 1
		fmt.Printf("|%3d |", index)

		for i := 0; i < 20-2*len((*currentNode).sequence.sarray); i++ {
			fmt.Printf(" ")
		}
		for _, e := range (*currentNode).sequence.sarray {
			fmt.Printf("%d ", e)
		}
		fmt.Printf("|")

		fmt.Printf("%6d |", (*currentNode).response.black_pins)
		fmt.Printf("%6d |\n", (*currentNode).response.white_pins)

		if (*currentNode).next == nil {
			break
		} else {
			currentNode = (*currentNode).next
		}
	}

	fmt.Printf("+----+--------------------+-------+-------+\n")
}
