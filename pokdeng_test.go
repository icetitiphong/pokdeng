package pokdeng

import (
	"testing"
)

func TestPlayer1ShouldBeWinnerWhenScoreMoreThanPlayer2(t *testing.T) {
	fakeP1 := 9
	fakeP2 := 5
	want := "Player 1 is winner"
	get := Pokdeng(fakeP1, fakeP2)

	if want != get {
		t.Errorf("want %q but got %q\n", want, get)
	}
}

func TestPlayer2ShouldBeWinner_WhenScoreMoreThanPlayer1(t *testing.T) {
	fakeP1 := 2
	fakeP2 := 5
	want := "Player 2 is winner"
	get := Pokdeng(fakeP1, fakeP2)

	if want != get {
		t.Errorf("want %q but got %q\n", want, get)
	}
}

func TestPlayer1AndPlayer2ShouldBeDraw_WhenScoreIsSame(t *testing.T) {
	fakeP1 := 5
	fakeP2 := 5
	want := "draw"
	get := Pokdeng(fakeP1, fakeP2)

	if want != get {
		t.Errorf("want %q but got %q\n", want, get)
	}
}

func TestMaxumumScoreShouldNotOver9(t *testing.T) {
	fake := []int{10, 5}
	want := 5
	get := CalScore(fake)

	if want != get {
		t.Errorf("want %q but got %q\n", want, get)
	}
}

func TestMaxumumScoreShouldBeError_WhenScoreisOver9(t *testing.T) {
	fake := []int{10, 10}
	want := 0
	get := CalScore(fake)

	if want != get {
		t.Errorf("want %q but got %q\n", want, get)
	}
}

func TestCreateCardMustBeEqual40Cards(t *testing.T) {

	want := 40
	get := CreateCard()

	if want != len(get) {
		t.Errorf("want %q but got %q\n", want, get)
	}
}
