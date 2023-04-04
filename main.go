package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

// declare global variables
var ia_started int = 0
var resposta string

// array of general responses
var gereral_response = [6]string{
	"Sim, Não, Talvez...",
	"Talvez",
	"Não sei",
	"Não quero responder",
	"Cansei de responder agora",
	"Acabou meu banco de dados",
}

func main() {
	// random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// phrase to be printed
	phrase := "Querido computador, gostaria de saber"
	index := 0

	// open keyboard
	if err := keyboard.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao abrir o teclado:", err)
		os.Exit(1)
	}
	defer keyboard.Close()
	// print instructions
	fmt.Println("Faça uma pergunta ao computador e pressione ENTER para ver a resposta. ESC para sair.")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Erro ao ler o teclado:", err)
			os.Exit(1)
		}

		// check if key pressed is ;
		if char == 59 {
			// check if ia_started is 0
			if ia_started == 0 {
				// ; is pressed, show the phrase
				ia_started = 1
			} else if ia_started == 1 {
				// if ; is pressed again, store the phrase
				ia_started = 2
			}
		}

		// if esc is pressed, break
		if key == keyboard.KeyEsc {
			break
		}

		// if enter is pressed, print a response
		if key == keyboard.KeyEnter {
			fmt.Println()
			// if ia_started is 2, print a secret response
			if ia_started == 2 {
				fmt.Println("Resposta: " + resposta)
			} else {
				// if ia_started is not started, print a general response
				// generate randon number between 0 and 5
				random := rand.Intn(6)
				fmt.Println("Resposta: " + gereral_response[random])
			}
			break
		}

		// if ia_started is 1, print phrase
		if ia_started == 1 {
			if char != 0 {
				fmt.Printf("%c", phrase[index])
				// store the char in resposta
				index = (index + 1) % len(phrase)
				if char != 59 {
					resposta += string(char)
				}
			}
			if key == keyboard.KeySpace {
				resposta += " "
			}

		} else {
			// if ia_started is 2, print the question
			if char != 0 {
				if char != 59 {
					fmt.Printf("%c", char)
				}
			}
			// if space is pressed, print a space
			if key == keyboard.KeySpace {
				fmt.Printf(" ")
			}
		}
	}
}
