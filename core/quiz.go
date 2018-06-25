package core

import (
	"exercises/quiz1/csvparser"
	"bufio"
	"os"
	"log"
	"fmt"
)

type Score struct {
	Score int
	Correct int
	Wrong int
}

func (s *Score) AnsCorrect(){
	s.Score += 1
	s.Correct +=1
}

func (s *Score) AnsWrong(){
	s.Wrong += 1
}

func Startquiz(recs []*csvparser.Record) Score {
	score := &Score{0,0,0}
	for cnt, rec := range recs {
		askquestion(rec, score, cnt+1)
	}
	return *score
}

func askquestion(r *csvparser.Record, s *Score, cnt int){
	fmt.Println("\n\nQuestion #",cnt,": ", r.Question)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	fmt.Println("You have entered: ", str)
	if err != nil {
		log.Fatal(err)
        }
	if  r.CheckAnswer(str)  {
		s.AnsCorrect()
		fmt.Println("Your Answer is correct")
	} else {
		s.AnsWrong()
		fmt.Println("Your Answer is Wrong")
		fmt.Println("Correct Answer is ",r.Answer)
	}
}
