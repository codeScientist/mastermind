package mastermindlib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GuessTreeNode struct {
	sequence     Sequence
	next_guesses []*GuessTreeNode
}

func Get_GuessTree(gs *GameState) *GuessTreeNode {
	// create the full path to file
	fileloc := gen_fileloc(gs)

	// check if the file exists
	_, err := os.Stat(fileloc)
	if err != nil && os.IsNotExist(err) {
		// create the file
		c := make(chan *GuessTreeNode)
		go Gen_GuessTree(gs, nil, c)
		gtn := <-c
		write_file(gs, gtn, nil)
		return gtn
	}

	// read the file
	return read_file(gs, nil)
}

func gen_fileloc(gs *GameState) string {
	mode := 'm'
	if (*gs).Is_single_use {
		mode = 's'
	}
	filename := fmt.Sprintf("%d_%d_%c.txt", (*gs).Total_pins, (*gs).Total_selectable_pins, mode)
	filepath := "./GuessTrees/"
	fileloc := filepath + filename
	return fileloc
}

func write_file(gs *GameState, root *GuessTreeNode, bufferedWriter *bufio.Writer) {
	if bufferedWriter == nil {
		fileloc := gen_fileloc(gs)
		newFile, _ := os.Create(fileloc)
		newFile.Close()
		file, _ := os.OpenFile(fileloc, os.O_WRONLY, 0666)
		bufferedWriter = bufio.NewWriter(file)
		defer file.Close()
	}

	s := ""
	for _, e := range (*root).sequence.sarray {
		if len(s) == 0 {
			s = strconv.Itoa(e)
		} else {
			s = s + " " + strconv.Itoa(e)
		}
	}
	s = s + "\n"

	bufferedWriter.WriteString(s)
	bufferedWriter.Flush()
	for _, e := range (*root).next_guesses {
		if e == nil {
			bufferedWriter.WriteString("0\n")
		} else {
			write_file(gs, e, bufferedWriter)
		}
	}
	bufferedWriter.Flush()
	return
}

func read_file(gs *GameState, scanner *bufio.Scanner) *GuessTreeNode {
	if scanner == nil {
		fileloc := gen_fileloc(gs)
		file, _ := os.Open(fileloc)
		scanner = bufio.NewScanner(file)
		defer file.Close()
	}

	success := scanner.Scan()
	if success == false {
		return nil
	}

	l := scanner.Text()
	if l == "0" {
		return nil
	}

	s := Create_Sequence(gs)
	sa := strings.Fields(l)
	for i, e := range sa {
		s.sarray[i], _ = strconv.Atoi(e)
	}

	gtn := GuessTreeNode{sequence: s, next_guesses: make([]*GuessTreeNode, (*gs).Total_responses)}

	for i := range gtn.next_guesses {
		gtn.next_guesses[i] = read_file(gs, scanner)
	}
	return &gtn
}

func Gen_GuessTree(gs *GameState, root *GuessListNode, c chan *GuessTreeNode) {
	guess_seq, e := Gen_minimax_next_guess(gs, root, nil, nil)
	if e != nil {
		c <- nil
		return
	}

	gtn := GuessTreeNode{sequence: guess_seq, next_guesses: make([]*GuessTreeNode, (*gs).Total_responses)}

	for i := 0; i < (*gs).Total_responses-1; i++ {
		res := Response{0, 0, i}
		res.Set_response_and_index(gs, false)
		gln := Create_GuessListNode(&guess_seq, &res)
		root_copy := root.Copy_GuessList()
		e := root_copy.Add_GuessListNode(&gln)
		if e != nil {
			root_copy = &gln
		}
		ch := make(chan *GuessTreeNode)
		go Gen_GuessTree(gs, root_copy, ch)
		gtn_ng := <-ch
		gtn.next_guesses[i] = gtn_ng
	}
	gtn.next_guesses[(*gs).Total_responses-1] = nil
	c <- &gtn
	return
}

func (root *GuessTreeNode) Cal_height() int {
	if root == nil {
		return 0
	}
	max := 0
	for i := range (*root).next_guesses {
		height := (*root).next_guesses[i].Cal_height()
		if height > max {
			max = height
		}
	}
	return max + 1
}

func (root *GuessTreeNode) Print_Guess() {
	(*root).sequence.Print_Sequence()
}

func (root *GuessTreeNode) Get_next_GuessTreeNode(index int) *GuessTreeNode {
	return (*root).next_guesses[index]
}
