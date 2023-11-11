// Package main provides main    
package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string{
	
	if name == "" {
		return Hello("world", language)
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string){
	switch language{
		case "French":
			prefix= frenchHelloPrefix 
		case "Spanish":
			prefix= spanishHelloPrefix 
		default:
			prefix= englishHelloPrefix
		}

		return
	}


func main(){
	fmt.Println(Hello("Chris", "English"))
}
