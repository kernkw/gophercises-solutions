package main

import (
	"strings"
	"testing"
)

func TestParseQuiz(t *testing.T) {
	lines := [][]string{{"1+1", "2"}, {"3+4", " 7 "}, {"0+1", "1"}}
	actualProblems := parseQuiz(lines)

	for i, line := range lines {
		expectedQuestion := line[0]
		expectedAnswer := strings.TrimSpace(line[1])

		if actualProblems[i].question != expectedQuestion {
			t.Errorf("actual question %s does not equal expected %s",
				actualProblems[i].question, expectedQuestion)
		}
		if actualProblems[i].answer != expectedAnswer {
			t.Errorf("actual answer %s does not equal expected %s",
				actualProblems[i].answer, expectedAnswer)
		}
	}
}

func TestCheckAnswerCorrect(t *testing.T) {
	var correct int
	expectedAns := "50"
	givenAns := "50"
	p := problem{answer: expectedAns}
	checkAnswer(givenAns, p, &correct)

	if correct != 1 {
		t.Errorf("Expected correct to be incremented by 1.")
	}
}

func TestCheckAnswerIncorrect(t *testing.T) {
	var correct int
	expectedAns := "50"
	givenAns := "51"
	p := problem{answer: expectedAns}
	checkAnswer(givenAns, p, &correct)

	if correct != 0 {
		t.Errorf("Expected correct to not be incremented by 1.")
	}
}
