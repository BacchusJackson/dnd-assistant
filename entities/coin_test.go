package entities

import "testing"

func checkTestTable(t *testing.T, expected []uint, results ...uint) {
	for i, value := range expected {
		checkError(t, value, results[i])
	}
}

func TestCopper(t *testing.T) {
	testCopper := Copper(105)
	c, s, g, p := testCopper.Value()
	checkTestTable(t, []uint{105, 10, 1, 0}, c, s, g, p)
	t.Log(testCopper)
}

func TestSilver(t *testing.T) {
	testSilver := Silver(465)
	c, s, g, p := testSilver.Value()
	checkTestTable(t, []uint{4650, 465, 46, 0}, c, s, g, p)
	t.Log(testSilver)
}

func TestGold(t *testing.T) {
	testGold := Gold(5401)
	c, s, g, p := testGold.Value()
	checkTestTable(t, []uint{540100, 54010, 5401, 54}, c, s, g, p)
	t.Log(testGold)
}

func TestPlatinum(t *testing.T) {
	testPlatinum := Platinum(27)
	c, s, g, p := testPlatinum.Value()
	checkTestTable(t, []uint{270000, 27000, 2700, 27}, c, s, g, p)
	t.Log(testPlatinum)
}
