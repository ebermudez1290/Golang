package main

import ("testing"
		"os")

func TestNewDeck(t *testing.T){
	d:= newDeck();
	if len(d)!=52{
		t.Errorf("Expected deck length of 52,but got %v", len(d))
	}

	if d[0] !="Ace of Cubes"{
		t.Errorf("Expected first card of Ace of Cubes, but got %v", d[0])
	}

	if d[len(d)-1] !="King of Hearts"{
		t.Errorf("Expected first card of King of Hearts, but got %v", d[len(d)-1])
	}
}

	//we should delete 
	func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
		os.Remove("_decktesting")
		d := newDeck()
		d.saveToFile("_decktesting")

		ld:= newDeckFromFile("_decktesting")

		if  len(ld)!=52 {
			t.Errorf("Expected deck length of 52,but got %v", len(ld))
		}
		os.Remove("_decktesting")
	}