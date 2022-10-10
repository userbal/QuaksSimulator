package main

import (
	"fmt"
	"math/rand"
)

var blowUpThreshold int = 7

var victoryPoints int = 0

type potLocation struct {
	money         int
	victoryPoints int
	ruby          bool
}

type boardLocation struct {
	value   int
	ratTail bool
}

type chip struct {
	color    string
	value    int
	function func(value int) int
}

type playerBoard struct {
	pot         []potLocation
	ingredients []chip
	position    int
	whites      int
}

var scoreBoard []boardLocation

var bag []chip

var pot playerBoard

func draw() bool {
	r := rand.Intn(len(bag))

	fmt.Println("You pulled a", bag[r].color, bag[r].value)

	transferChip(r)

	return pot.whites > blowUpThreshold
}

func transferChip(r int) {
	pot.ingredients = append(pot.ingredients, bag[r])

	pot.position += bag[r].function(bag[r].value)

	//remove selected chip from bag
	bag = append(bag[:r], bag[r+1:]...)
}

func endRound(isBlownUp bool) {
	// dice

	// green
	// black
	// purple

	// ruby

	// victory points
	if !isBlownUp {
		victoryPoints += pot.pot[pot.position].victoryPoints
	}

	// buy chips

	fmt.Printf("victoryPoints: %v\n", victoryPoints)
	// reset pot
	pot = setupPot()
}

func start() {
	isBlownUp := false

	var response string
	for i := 0; i < 9; i++ {
		fmt.Println("Round", i+1)
	round:
		for {
			// showBag()
			fmt.Println("draw chip? y/n")
			fmt.Scan(&response)
			if response == "y" {
				isBlownUp = draw()
				if isBlownUp {
					fmt.Println("You blew up!!")
					break round
				}

			} else if response == "n" {
				break round
			}
		}
		endRound(isBlownUp)
	}
}

func main() {
	//setup
	scoreBoard = setupBoard()

	pot = setupPot()

	bag = setupBag()

	rand.Seed(2)

	start()
}

// TODO buying chips
// TODO green, black and purple
// TODO dice
// TODO rubies
// TODO rat tails
// TODO multiple players
// TODO make an api that can connect to machine learning through python?
// TODO convert static data like board and pot locations to json files
