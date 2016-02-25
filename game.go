package liarsdice

import (
	"math"
	"math/rand"
)

const (
	dicePerPlayer = 5
	dieMax = 6
)

type Move struct {
	playerNumber int
	numDice int
	valueOfDie int
}

type Die struct {
	value int
}

type Player struct {
	dice []Die
}

type Game struct {
	players []*Player
	moves []*Move
}

type InvalidMoveError struct {
	message string
	move *Move
}

func (e *InvalidMoveError) Error() string {
	return "Invalid Move"
}

func NewGame(numPlayers int) *Game {
	game := new(Game)
	game.players = make([]*Player, numPlayers)

	for i := 0; i < numPlayers; i++ {
		game.players[i] = NewPlayer()
	}

	return game
}

func NewPlayer() *Player {
	player := new(Player)

	player.dice = make([]Die, dicePerPlayer)

	for i := 0; i < dicePerPlayer; i++ {
		player.dice[i] = Die{rand.Intn(dieMax + 1)}
	}

	return player
}


func (self *Game) Move(playerIndex int, numDice int, valueOfDice int) error {
	if numDice < self.DiceLeft() {
		self.moves = append(self.moves, &Move{playerIndex,numDice,valueOfDice})
		return nil
	} else {
		return &InvalidMoveError{"Not a valid move", &Move{playerIndex,numDice,valueOfDice}}
	}
}

func (self *Game) Probability(k float64, n float64) (result float64) {
	if k <= n {
		//fmt.Printf("(Factorial(%.2f) / (Factorial(%.2f) * Factorial(%.2f - %.2f))) * math.Pow(1.0/6.0, %.2f) * math.Pow(5.0/6.0, (%.2f - %.2f))\n",n,k,n,k,k,n,k)
		return (Factorial(n) / (Factorial(k) * Factorial(n - k))) * math.Pow(1.0/6.0, k) * math.Pow(5.0/6.0, (n - k))
	} else {
		return 0.0
	}
}

func (self *Game) Claim(numDice int, valueOfDice int) (result float64) {
	diceLeft := float64(self.DiceLeft())

	if diceLeft > 1 {
		return (self.Probability(diceLeft, diceLeft) + self.Probability(diceLeft - 1.0, diceLeft))
	} else {
		return self.Probability(diceLeft, diceLeft)
	}
}

func (self *Game) Challenge(numDice int, valueOfDice int) bool {
	return self.TotalDiceWithValue(valueOfDice) == numDice
}

func (self *Game) TotalDiceWithValue(valueOfDice int) int {
	totalDice := 0
	for _, player := range self.players {
		for _, die := range player.dice {
			if die.value == valueOfDice {
				totalDice++
			}
		}
	}
	return totalDice
}

func (self *Game) TotalDice() int {
	return dicePerPlayer * len(self.players)
}

func (self *Game) DiceInPlay() int {
	diceInPlay := 0
	for _, move := range self.moves {
		diceInPlay += move.numDice
	}
	return diceInPlay
}

func (self * Game) DiceLeft() int {
	return self.TotalDice() - self.DiceInPlay()
}

func Factorial(n float64)(result float64) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1.0
}
