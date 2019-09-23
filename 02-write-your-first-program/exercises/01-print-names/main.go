// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// ---------------------------------------------------------
// EXERCISE: Print names
//
//  Print your name and your best friend's name using
//  Println twice
//
// EXPECTED OUTPUT
//  YourName
//  YourBestFriendName
//
// BONUS
//  Use `go run` first.
//  And after that use `go build` and run your program.
// ---------------------------------------------------------


import "fmt"
func main() {
	fmt.Println("Yu Weng")
	fmt.Println("Jesus")
}

//******* Testing section **********//
// Method 1: (after setting working directory to current directory) go run main.go 
// Method 2: (after setting working directory to current directory) go build main.go --which produces the main.exe file in the same directory
// followed by: .\main.exe