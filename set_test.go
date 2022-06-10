package set

import (
	"testing"
)

func TestAddWithRepeated(t *testing.T) {

	hashSet := NewHashSet[playerKey, *player]()

	hashSet.Add(newPlayer("1"))
	hashSet.Add(newPlayer("2"))
	hashSet.Add(newPlayer("3"))

	//Repeated id must not be inserted
	hashSet.Add(newPlayer("3"))

	expected := 3
	got := hashSet.Size()

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestAddWithSameIdDifferentAttributes(t *testing.T) {

	player1 := newPlayer("1")
	player2 := newPlayer("1")

	//Setting different attributes but they have the same id
	player2.Name = "some random name"
	player2.Age = 50

	hashSet := NewHashSet[playerKey, *player]()

	hashSet.Add(player1)
	//Repeated id must not be inserted
	hashSet.Add(player2)

	expected := 1
	got := hashSet.Size()

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestIsEmptyWithEmptySet(t *testing.T) {
	hashSet := NewHashSet[playerKey, *player]()

	got := hashSet.IsEmpty()
	expected := true

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

}

func TestIsEmptyWithNoneEmptySet(t *testing.T) {
	hashSet := NewHashSet[playerKey, *player]()

	hashSet.Add(newPlayer("1"))

	got := hashSet.IsEmpty()
	expected := false

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

}

func TestContainsEmptySet(t *testing.T) {
	hashSet := NewHashSet[playerKey, *player]()

	got := hashSet.Contains(newPlayer("1"))
	expected := false

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

}

func TestContainsNoneExistentEntry(t *testing.T) {
	hashSet := NewHashSet[playerKey, *player]()

	hashSet.Add(newPlayer("1"))

	got := hashSet.Contains(newPlayer("2"))
	expected := false

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

}

func TestContainsExistentEntry(t *testing.T) {
	hashSet := NewHashSet[playerKey, *player]()

	hashSet.Add(newPlayer("1"))

	got := hashSet.Contains(newPlayer("1"))
	expected := true

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

}

func TestDeleteExistentEntry(t *testing.T) {
	hashSet := NewHashSet[playerKey, *player]()

	hashSet.Add(newPlayer("1"))

	if !hashSet.Contains(newPlayer("1")) {
		t.Errorf("Cannot add player 1")
	}

	//Removing player
	hashSet.Remove(newPlayer("1"))

	got := hashSet.Contains(newPlayer("1"))
	expected := false

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

}

func TestDeleteNoneExistentEntry(t *testing.T) {
	hashSet := NewHashSet[playerKey, *player]()

	hashSet.Add(newPlayer("1"))

	if !hashSet.Contains(newPlayer("1")) {
		t.Errorf("Cannot add player 1")
	}

	//Removing none existent player
	hashSet.Remove(newPlayer("2"))

	got := hashSet.Contains(newPlayer("1"))
	expected := true

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

}

func TestClearNoneEmptySet(t *testing.T) {
	hashSet := NewHashSet[playerKey, *player]()

	hashSet.Add(newPlayer("1"))
	hashSet.Add(newPlayer("2"))
	hashSet.Add(newPlayer("3"))

	hashSet.Clear()

	got := hashSet.Size()
	expected := 0

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

}

func TestClearEmptySet(t *testing.T) {
	hashSet := NewHashSet[playerKey, *player]()

	hashSet.Clear()

	got := hashSet.Size()
	expected := 0

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

}

type player struct {
	Id   string
	Age  int
	Name string
}

type playerKey struct {
	id string
}

func newPlayer(id string) *player {
	return &player{
		Id:   id,
		Name: "player " + id,
		Age:  18,
	}
}

func (p *player) GetKey() playerKey {
	return playerKey{
		id: p.Id,
	}
}
