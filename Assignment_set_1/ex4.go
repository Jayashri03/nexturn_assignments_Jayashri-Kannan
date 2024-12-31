<<<<<<< HEAD
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Question string
	Options  []string
	Answer   int // The index of the correct answer
}

var questionBank = []Question{
	{
		Question: "What is the capital of France?",
		Options:  []string{"1. Berlin", "2. Madrid", "3. Paris", "4. Rome"},
		Answer:   3,
	},
	{
		Question: "What is 2 + 2?",
		Options:  []string{"1. 3", "2. 4", "3. 5", "4. 6"},
		Answer:   2,
	},
	{
		Question: "Who wrote 'To Kill a Mockingbird'?",
		Options:  []string{"1. Harper Lee", "2. J.K. Rowling", "3. Ernest Hemingway", "4. Jane Austen"},
		Answer:   1,
	},
}

func takeQuiz(questions []Question, timePerQuestion time.Duration) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	score := 0

	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		for _, option := range q.Options {
			fmt.Println(option)
		}
		
		answerCh := make(chan string)
		go func() {
			fmt.Print("Your answer (or type 'exit' to quit): ")
			if scanner.Scan() {
				answerCh <- scanner.Text()
			}
		}()

		timer := time.NewTimer(timePerQuestion)
		select {
		case <-timer.C:
			fmt.Println("\nTime's up for this question!")
			continue
		case answer := <-answerCh:
			timer.Stop()
			if strings.ToLower(answer) == "exit" {
				fmt.Println("\nExiting the quiz.")
				return score, nil
			}

			selectedOption, err := strconv.Atoi(answer)
			if err != nil || selectedOption < 1 || selectedOption > len(q.Options) {
				fmt.Println("Invalid input. Skipping question.")
				continue
			}

			if selectedOption == q.Answer {
				score++
			}
		}
	}

	return score, nil
}

func classifyPerformance(score, totalQuestions int) string {
	percentage := (float64(score) / float64(totalQuestions)) * 100
	if percentage >= 80 {
		return "Excellent"
	} else if percentage >= 50 {
		return "Good"
	} else {
		return "Needs Improvement"
	}
}

func main() {
	const timePerQuestion = 10 * time.Second

	fmt.Println("Welcome to the Online Examination System!")
	fmt.Printf("You have %d seconds to answer each question.\n", timePerQuestion/time.Second)
	fmt.Println("Type 'exit' anytime to quit the quiz.")

	score, err := takeQuiz(questionBank, timePerQuestion)
	if err != nil {
		fmt.Println("Error during quiz:", err)
		return
	}

	fmt.Printf("\nQuiz completed! Your score: %d/%d\n", score, len(questionBank))
	performance := classifyPerformance(score, len(questionBank))
	fmt.Printf("Your performance: %s\n", performance)

}
=======
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Question string
	Options  []string
	Answer   int // The index of the correct answer
}

var questionBank = []Question{
	{
		Question: "What is the capital of France?",
		Options:  []string{"1. Berlin", "2. Madrid", "3. Paris", "4. Rome"},
		Answer:   3,
	},
	{
		Question: "What is 2 + 2?",
		Options:  []string{"1. 3", "2. 4", "3. 5", "4. 6"},
		Answer:   2,
	},
	{
		Question: "Who wrote 'To Kill a Mockingbird'?",
		Options:  []string{"1. Harper Lee", "2. J.K. Rowling", "3. Ernest Hemingway", "4. Jane Austen"},
		Answer:   1,
	},
}

func takeQuiz(questions []Question, timePerQuestion time.Duration) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	score := 0

	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		for _, option := range q.Options {
			fmt.Println(option)
		}
		
		answerCh := make(chan string)
		go func() {
			fmt.Print("Your answer (or type 'exit' to quit): ")
			if scanner.Scan() {
				answerCh <- scanner.Text()
			}
		}()

		timer := time.NewTimer(timePerQuestion)
		select {
		case <-timer.C:
			fmt.Println("\nTime's up for this question!")
			continue
		case answer := <-answerCh:
			timer.Stop()
			if strings.ToLower(answer) == "exit" {
				fmt.Println("\nExiting the quiz.")
				return score, nil
			}

			selectedOption, err := strconv.Atoi(answer)
			if err != nil || selectedOption < 1 || selectedOption > len(q.Options) {
				fmt.Println("Invalid input. Skipping question.")
				continue
			}

			if selectedOption == q.Answer {
				score++
			}
		}
	}

	return score, nil
}

func classifyPerformance(score, totalQuestions int) string {
	percentage := (float64(score) / float64(totalQuestions)) * 100
	if percentage >= 80 {
		return "Excellent"
	} else if percentage >= 50 {
		return "Good"
	} else {
		return "Needs Improvement"
	}
}

func main() {
	const timePerQuestion = 10 * time.Second

	fmt.Println("Welcome to the Online Examination System!")
	fmt.Printf("You have %d seconds to answer each question.\n", timePerQuestion/time.Second)
	fmt.Println("Type 'exit' anytime to quit the quiz.")

	score, err := takeQuiz(questionBank, timePerQuestion)
	if err != nil {
		fmt.Println("Error during quiz:", err)
		return
	}

	fmt.Printf("\nQuiz completed! Your score: %d/%d\n", score, len(questionBank))
	performance := classifyPerformance(score, len(questionBank))
	fmt.Printf("Your performance: %s\n", performance)

}
>>>>>>> e0759172fad731f5d36f1ff310d9cd155d1db974
