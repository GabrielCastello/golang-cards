package main

import (
	"fmt"
	// "log"
	// "net/http"
)

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// card := newCard()

	// fmt.Println(card)
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

	// fmt.Println("Go Docker Tutorial!!!!")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello World !!!")
	// })

	// log.Fatal(http.ListenAndServe(":8081", nil))

}

func newCard() string {
	return "Fice of Diamonds"
}
