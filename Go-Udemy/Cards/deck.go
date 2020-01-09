package main

import ("fmt"
		"strings"
		"io/ioutil"
		"os"
		"math/rand"
		"time"
		)

type deck []string

func newDeck() deck {
	cards := deck {}
	cardSuites:= []string{"Cubes","Spades","Diamonds", "Hearts"}
	cardNumbers:= []string{"Ace","Two","Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}
	for _,suite:=range cardSuites{
		for _,number:=range cardNumbers{
			cards = append(cards, fmt.Sprintf("%s of %s",number,suite))
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck,deck) {
	return d[:handsize],d[handsize:]
}

//Create will create a deck of cards.
func (d deck) print(){
	for i,card := range d{
		fmt.Println(i,card)
	}
}

//toString will get you the string representation of the deck
func (d deck) toString()string{
	cardSlice := []string(d)
	return strings.Join(cardSlice,",")
}

func (d deck) saveToFile (fileName string) error  {
	return ioutil.WriteFile(fileName,[]byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs,err:=ioutil.ReadFile(fileName)
	if err!=nil{
		//For handling errors you either choose to continue or determine that the error is critical.
		fmt.Println("Error:",err)
		os.Exit(-1)
	}else{
		s := strings.Split(string(bs),",")
		return deck(s)
	}
	return nil
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)
	
	for i := range d{
		np:=seed.Intn(len(d)-1)
		d[i],d[np] = d[np],d[i]
	}
	
}