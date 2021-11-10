package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0
	PAPER        = 1
	SCISSORS     = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	messageRnd := rand.Intn(3)

	computerChoice := ""
	roundResult := ""

	messageMap := make(map[int][]string)

	messageMap[PLAYERWINS] = []string{"GOAT!", "Well done.", "Greatest of all time."}
	messageMap[COMPUTERWINS] = []string{"Try again!", "Don`t give up!", "Keep trying! You can win!"}
	messageMap[DRAW] = []string{"Nobody wins :(", "Oh oh. You will have to play again I guess.", "Try playing again."}

	var message string

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		message = messageMap[DRAW][messageRnd]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		message = messageMap[PLAYERWINS][messageRnd]
	} else {
		roundResult = "Computer wins!"
		message = messageMap[COMPUTERWINS][messageRnd]
	}

	result := Round{
		Message:        message,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}
	return result
}
