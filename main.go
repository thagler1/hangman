package main


import (
    "bufio"
    "fmt"
	"os"
	"time"
	"math/rand"
	"io/ioutil"
	"strings"
	"os/exec"
)

const filepath  string = "C:\\Users\\todd\\Documents\\Go\\words.txt"

func get_input()(string){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Guess: ")
	guess, _ := reader.ReadString('\n')
	fmt.Println("You have guessed: ",guess)
	return guess
}

func rand_number(max int)(int){
	rand.Seed(time.Now().Unix())
    return rand.Intn(max)
}

func load_words()(string){
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
    fmt.Println("load words failed ..") // descriptive error, kinda 
}
	lines := strings.Split(string(content), "\n")
	r := rand_number(len(lines))

	return lines[r]
}

func print_board(word string, guesses []byte){
	board := make([]string, len(word))
	
	fmt.Println("The word to guess is : ",word)
	for i:=0; i<len(word); i++{ //for each letter
		found := false
		for c:=0; c<len(guesses); c++{ // fpr each letter go through each guess
			
			if word[i] == guesses[c]{
				found = true
				board = append(board, string(word[i]))
			} 
			

		}
		if found == false{
			board = append(board, "_")
		}
	}

	fmt.Println(board)
	
}





func start_game(guesses []byte)(string){
	words := load_words()
	print_board(words, guesses)

	return words
	
}

func check_win(word string, guesses []byte)(bool){ //False if you won. This will exit mains for loop :/
	
	for i:=0; i<len(word); i++{ //for each letter in word
		found := false
		for c:=0; c<len(guesses); c++{ //for each guess
			if word[i] == guesses[c]{
				found = true
			}

		}
		if found == false{
			return true
		}
	}

	return false

}



func main() {
	guesses := make([]byte, 100)
	word := start_game(guesses)
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
	
	x:=true
	for x==true{
		print_board(word, guesses)
		guess := get_input()
		newguess := []byte(guess)
		
		guesses = append(guesses, newguess...)
		
		
		x = check_win(word, guesses)



		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		c.Run()
				if x==false{
			fmt.Println("You won!!")
		}
		

	}

}

	


	