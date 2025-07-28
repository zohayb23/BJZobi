// Create a minimal CLI terminal blackjack game
// The game should be played in the terminal
// The game should be played with a deck of cards
// The game should be played with a player and a dealer
// The game should be played with a deck of cards

// The player starts with with two cards at random, both shown to the user
// The dealer starts with two cards, one shown, and one hidden,
// the hidden card is revealed after the player stands with his cards or if the player busts and loses the game.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// clearScreen clears the terminal screen
func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape code to clear screen
}

// Define a card struct with Suit and Value fields
type Card struct {
	Suit  string // Hearts, Diamonds, Clubs, Spades
	Value string // 2-10, Jack, Queen, King, Ace
}

// Create a function to build a standard deck of 52 cards (Slice of Cards)
// Need to combine suits and values
func buildDeck() []Card {
	deck := []Card{} // Initialize an empty slice ,"deck" that holds all cards
	// make a slice of suits with type string
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	//make a slice of values with type string
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	//use nested loops to combine each suit with each value
	for _, suit := range suits { //Outer loop goes through each suit in suits slice
		for _, value := range values { //Inner loop goes through each value in values slice
			// store each values in the deck slice
			card := Card{Suit: suit, Value: value} //For every combintaion of suit and value, it creates a Card Struct
			deck = append(deck, card)              //adds the cards into the empty deck slice using append()

		}

	}
	return deck //returns 52 unique cards

}

// Write a shuffle function using rand.Shuffle
// Called after buildDeck() is ran
func shuffleDeck(deck []Card) { //call the deck
	rand.Shuffle(len(deck), func(i, j int) { //Shuffles the length of the deck, function swaps two cards
		deck[i], deck[j] = deck[j], deck[i] //swaps the cards from random indexes in the deck
	})
}

// Write a function to draw a card from the deck
// deck must be shuffled prior, shuffle deck to be called in main before drawCard()
// takes one card out of the deck
// return the card so we know what we drew
// return the remaining deck without that card so we don't draw it again
// draws from the end of the slice and returns both, last card drawn and the remaining cards in the deck
func drawCard(deck []Card) (Card, []Card) { // Card = the single card you drew, []Card = the remaining deck slice
	lastCard := deck[len(deck)-1] // the last card
	newDeck := deck[:len(deck)-1] // everything but the last card
	// := tells Go “create a new variable and assign it.
	return lastCard, newDeck
}

// getCardValue returns the numeric Blackjack value of a single card.
// Face cards (Jack, Queen, King) are worth 10.
// Ace is worth 11 (can be adjusted later in handValue).
// Number cards are converted from string to int.
// since values are currently type String, we need to convert them to numbers so the game can function
func getCardValue(card Card) int {
	switch card.Value { //change for loop and if statement to switch statement
	case "Jack", "Queen", "King": // case is like if statement, if card.Value is equal to J,Q,K. ret 10
		return 10
	case "Ace": // else if card is Ace return 11, need to make Ace logic in handValue() because depends on the total hand value
		return 11
	default:
		num, _ := strconv.Atoi(card.Value) // Convert string like "7" to int 7
		//researched online how to convert string to int, found strconv import that has a method to make it easier
		return num
	}
}

// Implment hand value, uses card values to add up whats in the players/dealers hand
// Remember to handle Aces properly, 1 or 11
// Input: A slice of Card represents a hand
// Output: The total Blackjack value of that hand
func handValue(hand []Card) int { // slice of card, "hand", return total hand value as int
	total := 0    // hold the sum of the hand
	aceCount := 0 // count the amount of Aces in hand/deck

	//loop through each card
	for _, card := range hand { //for each card in hand, _ = blank identifier, could replace with i = index, but we dont care about index, only the value
		value := getCardValue(card) //get the value
		total += value              // add it to the total
		if card.Value == "Ace" {    // if card value is Ace
			aceCount++ // increment aceAcount
		}
	}
	for total > 21 && aceCount > 0 { //if total hand value is greater than 21 and ace count is higher than 0
		total -= 10 // minus the total by 10 to use the Ace as a 1 rather than an 11
		aceCount--  //reduce aceCount by 1 because we handled one ace in the hand
	}
	return total //returns total int value of value in hand
}

// test cases func
func runTest() {
	fmt.Println("Welcome to Blackjack")
	//Tests Build deck
	test := buildDeck() //initializes the build of the deck
	fmt.Println("Result of building the deck", test)
	fmt.Println("Deck size:", len(test))

	//Test Shuffle deck
	fmt.Println("Before shuffle:", test[:5]) // print first 5 to see order
	shuffleDeck(test)
	fmt.Println("After shuffle:", test[:5]) // print first 5 to see shuffled order

	//Tests drawCard()
	card, test := drawCard(test)
	fmt.Println("Drew:", card)
	fmt.Println("Deck size after draw:", len(test))
	card2, test := drawCard(test)
	fmt.Println("Drew:", card2)
	fmt.Println("Deck size after draw:", len(test))
	card3, test := drawCard(test)
	fmt.Println("Drew:", card3)
	fmt.Println("Deck size after draw:", len(test))
	card4, test := drawCard(test)
	fmt.Println("Drew:", card4)
	fmt.Println("Deck size after draw:", len(test))

	//test getCardValue()
	// Test cards to check their values
	testCards := []Card{
		{Suit: "Hearts", Value: "7"},
		{Suit: "Clubs", Value: "King"},
		{Suit: "Diamonds", Value: "Ace"},
		{Suit: "Spades", Value: "10"},
		{Suit: "Hearts", Value: "Queen"},
		{Suit: "Clubs", Value: "Jack"},
	}
	for _, card := range testCards {
		val := getCardValue(card)
		fmt.Printf("Card: %s of %s → Value: %d\n", card.Value, card.Suit, val)
	}

	//test handValue()
	hand1 := []Card{
		{Suit: "Hearts", Value: "Ace"},
		{Suit: "Spades", Value: "9"},
	}
	fmt.Println("Hand1 (Ace + 9):", handValue(hand1)) // expect 20

	// Hand: Ace + Ace + 9 (should be 21)
	hand2 := []Card{
		{Suit: "Hearts", Value: "Ace"},
		{Suit: "Spades", Value: "Ace"},
		{Suit: "Clubs", Value: "9"},
	}
	fmt.Println("Hand2 (Ace + Ace + 9):", handValue(hand2)) // expect 21

	// invalidCard := Card{Suit: "Hearts", Value: "Joker"}
	// val := getCardValue(invalidCard)
	// fmt.Printf("Invalid card value test: %s of %s → Value: %d\n", invalidCard.Value, invalidCard.Suit, val)
	// deck := buildDeck()
	// shuffleDeck(deck)
	// for len(deck) > 0 {
	// 	var card Card
	// 	card, deck = drawCard(deck)
	// 	fmt.Printf("Drew card: %s of %s | Cards left: %d\n", card.Value, card.Suit, len(deck))
	// }
	// // Try drawing one more card from empty deck (optional, you might want to handle this case)
	// if len(deck) == 0 {
	// 	fmt.Println("Deck is empty, no more cards to draw.")
	// }
}

func main() {
	clearScreen() //Clears the terminal screen
	//Assign for both player and dealer to use
	var card Card
	//First lets build the deck
	deck := buildDeck()
	//Shuffle deck
	shuffleDeck(deck)
	//Initalize hands by decalring an emply slice of Card, {} intializes empty slice
	playerHand := []Card{}
	dealerHand := []Card{}
	//both the player and the dealer need to draw cards
	// call the drawCard()
	//Players card
	card1, deck := drawCard(deck) //player first card. card 1 is the card drawn, deck, is the where the card is taken from
	playerHand = append(playerHand, card1)
	card2, deck := drawCard(deck)
	playerHand = append(playerHand, card2)

	//Dealers card
	card3, deck := drawCard(deck) // dealer first card
	dealerHand = append(dealerHand, card3)
	card4, deck := drawCard(deck) // dealer second card
	dealerHand = append(dealerHand, card4)

	fmt.Println("Welcome to Blackjack")
	// fmt.Printf("Player's Hand: %s of %s, ", playerHand[0].Value, playerHand[0].Suit) //Display player hands first card
	// fmt.Printf("%s of %s", playerHand[1].Value, playerHand[1].Suit)                  //Display player hands second card
	// fmt.Println(" ")
	// fmt.Printf("Dealer's Hand: %s of %s, [Hidden]\n", //Displays dealer hands first card but keeps the second one hidden for now
	// 	dealerHand[0].Value, dealerHand[0].Suit)

	//Players Turn
	//Ask for a hit or stand
	//Display the player's hand and its value
	//Display the dealers first card and Hidden
	//Ask for user input: "Hit or Stand?"
	//Read input from terminal using fmt.Scanln(&choice)
	// If player chooses hit, drawsCard(deck), and adds it to the playerHand
	// Also need to update the deck with the new one returned
	// Then Calculate handValue(playerHand)
	// if the value is > 21, then print "Player Busts", end the game
	//If player chooses Stand, break out of the loop and move to the dealers turn
	// If input is invalid, give an error and loop again
	// Keep looping if the player keeps choosing Hit and hasnt busted
	for {
		//print players initial hand and total value using handValue(playerHand)
		//fmt.Println("Players hand and total value:", handValue(playerHand))
		//print dealers first card and hide the second one
		//fmt.Printf("Dealer's Hand: %s of %s, [Hidden]\n", dealerHand[0].Value, dealerHand[0].Suit)
		//declare a string variable to store the input
		// var choice string
		// fmt.Print("Hit or Stand? ") //prompt the user Hit or Stand
		// fmt.Scanln(&choice)         //uses the Scanln() to read the line of text entered by the user until a newline char is pressed and stores it in the choice variable
		// fmt.Printf("Player entered: %s\n", choice)

		for _, card := range playerHand {
			fmt.Printf("- %s of %s\n", card.Value, card.Suit)

		}
		fmt.Println("Total value:", handValue(playerHand))
		fmt.Printf("Dealer's Hand: %s of %s, [Hidden]\n", dealerHand[0].Value, dealerHand[0].Suit)
		var choice string
		fmt.Print("Hit or Stand? ") //prompt the user Hit or Stand
		fmt.Scanln(&choice)         //uses the Scanln() to read the line of text entered by the user until a newline char is pressed and stores it in the choice variable
		fmt.Printf("Player entered: %s\n", choice)
		if choice == "Hit" {
			card, deck = drawCard(deck) //players card drawn, deck(sliced array) its from, := creates a new variable
			playerHand = append(playerHand, card)
			//Update the players hand

			if handValue(playerHand) > 21 {
				fmt.Println("Player busts!, Dealer Wins!")
				fmt.Println("Players Total value:", handValue(playerHand))
				fmt.Println("Dealer's full hand:")
				for _, card := range dealerHand {
					fmt.Printf("- %s of %s\n", card.Value, card.Suit)
				}
				fmt.Println("Dealer's Total value:", handValue(dealerHand))
				return // exit the funciton, ending the game right away without going to dealers turn
			}
		} else if choice == "Stand" {
			fmt.Println("Player choose Stand, Dealers Turn,")
			fmt.Println("Players Total value:", handValue(playerHand))
			break // move to dealer turn
		} else {
			fmt.Println("Invalid input, please type Hit or Stand")
			continue // retries loop
		}

	} // End of  user for loop
	// Need to implement dealer turn to reveal hidden card and to see if dealer wins, busts, ties
	//Dealers Turn
	//Loop through dealers full hand
	println("Dealers Turn")
	println("Dealers Hand")
	for _, card := range dealerHand {
		fmt.Printf("- %s of %s\n", card.Value, card.Suit)
		//fmt.Printf("Dealer's Hand: %s of %s, [Hidden]\n", dealerHand[0].Value, dealerHand[0].Suit)
	}
	fmt.Println("Dealer's Total value:", handValue(dealerHand))
	for handValue(dealerHand) < 17 { //"Dealer Stand on 17 Rule" Must hit until their total is 17 or more, ensures dealer cant stop with a very low hand like 12
		card, deck = drawCard(deck)           //draws card, dealerDraw from deck
		dealerHand = append(dealerHand, card) // appends card to dealer hand
		//update deck
		//print dealer drew a card
		fmt.Println("Dealer drew a card: ", card.Value, card.Suit)
		//display dealers hand
		fmt.Println("Dealers hand: ", handValue(dealerHand), card.Suit, card.Value)
	}
	//Compare results
	if handValue(dealerHand) > 21 {
		fmt.Println("Dealer Busts!, Player wins!")
	} else if handValue(dealerHand) > handValue(playerHand) {
		fmt.Println("Dealer wins!")
	} else if handValue(dealerHand) < handValue(playerHand) {
		fmt.Println("Player wins!")
	} else {
		fmt.Println("Its a Tie!")
	}

	//runTest()
} // End of main()
