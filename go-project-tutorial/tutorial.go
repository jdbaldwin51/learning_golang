package main


import "fmt"

func main() {
    fmt.Println("welcome to my quiz game")

    fmt.Printf("Enter you name: ")
    var name string
    fmt.Scan(&name)

    fmt.Printf("Hello, %v, welcome to the game!\n", name)

    fmt.Printf("Enter you age: ")
    var age uint
    fmt.Scan(&age)

    if age >= 10 {
        fmt.Println("Great, you can play!")
    } else{
        fmt.Println("Sorry you're not old enough to play")
        return
    }

    fmt.Printf("What is better, the RTX 3000 or RTX 3090? \n")
    var answer string
    fmt.Scan(&answer)

    fmt.Println(answer)
}
