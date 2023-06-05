package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck{
	cards := deck{}
	cardSuits :=[]string {"Spades","Hearts","Clubs","Diamonds"}
	cardValues :=[]string {"Aces","Two","Three","Four"}
	for _,suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards,value +" " + "of" + " " +suit)
		}
	}
	return cards
}

func (d deck) print(){
	for i,card := range d {
		fmt.Println(i,card)
	}
}



func (d deck) toString() string {
	return strings.Join((d),",")
}

func (d deck) saveToFile(filename string) error {
	 return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}

func newDeckFromFile(filename string) deck {
	bs,err := ioutil.ReadFile(filename)
	if err!= nil {
		fmt.Println("Error",err)
		os.Exit(1)
	}
	s := strings.Split(string(bs),",")
	return deck(s)
}

// package main

// func main() {
// 	cards := newDeckFromFile("my_Cards")
// 	cards.print()
	
// }
