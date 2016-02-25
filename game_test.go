package liarsdice

import (
	"testing"
	"fmt"
)

func TestNewPlayer(t *testing.T) {
	player := NewPlayer()
	
	if len(player.dice) != dicePerPlayer {
		t.Error("unable to initialize dice for player")
	}
}

func TestNewGame(t *testing.T) {
	game := NewGame(4)
	if len(game.players) != 4 {
		t.Errorf("did not have specified number of players")
	}
}

func TestTotalDice(t *testing.T) {
	game := NewGame(4)
	if game.TotalDice() != 20 {
		t.Errorf("TotalDice is not calculating correctly")
	}
}

func TestDiceInPlay(t *testing.T) {
	game := NewGame(4)
	_ = game.Move(1,3,5)
	if game.DiceInPlay() != 3 {
		t.Errorf("DiceInPlay is not calculating correctly")
	}
}

func TestDiceLeft(t *testing.T) {
	game := NewGame(4)
	game.Move(3,5,3)
	if game.DiceLeft() != 15 {
		t.Errorf("DiceLeft is not calculating correctly")
	}
}

func TestMove(t *testing.T) {
	game := NewGame(4)
	_ = game.Move(1,3,5)
	
	if len(game.moves) != 1 {
		t.Errorf("Move is not successfully being processed")
	}

	err := game.Move(1,35,4)
	if err == nil {
		t.Errorf("An invalid move was allowed")
	}
}

func TestClaim(t *testing.T) {
	game := NewGame(4)
	_ = game.Move(1,2,3)
	_ = game.Move(2,1,3)
	prob := game.Claim(19.0, 3.0)

	if prob < 0.00000000000505 || prob > 0.0000000000051 {
		t.Errorf("The claim calculations are not coming out right")
	}
}

func TestProbability(t *testing.T) {

	game := NewGame(4)

	prob := game.Probability(1.0, 1.0)
	if fmt.Sprintf("%.2f", prob) != "0.17" {
		t.Errorf("Probability is not working like it should")
	}

}

func TestTotalDiceWithValue(t *testing.T) {
	game := NewGame(4)
	_ = game.Move(1,2,3)
	
	if game.TotalDiceWithValue(3) == 0 {
		t.Errorf("not calculating TotalDiceWithValue correctly")
	}
}

func TestChallenge(t *testing.T) {
	game := NewGame(4)
	
	// It is highly unlikely that the randomized dice will all be 1...
	if game.Challenge(20, 1) != false {
		t.Errorf("Challenge is not computing false correctly")
	}

	// It is impossible for this to not be true
	if game.Challenge(1, game.players[1].dice[1].value) == true {
		t.Errorf("Challenge is not computing truth correctly")
	}
}

func TestExampleFromExercise(t *testing.T) {
	game := NewGame(4)

	game.Move(3,5,3)
	game.Move(4,5,3)
	game.Move(1,5,3)
	game.Move(2,4,3)

	claim := game.Claim(20, 3)
	
	if fmt.Sprintf("%.2f", claim) != "0.17" {
		t.Errorf("end to end did not work for the claim")
	}
}
