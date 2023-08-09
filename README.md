# Golang
Golang

1.  How do we run the code in our project?

    To run the code in our project, simply write `go run file.go` file.go is name of the file we want to execute.

    -> go run VS go build ?
    `go run` compiles and executes one or two files.
    `go build` compiles a bunch of go source code files.

2. What does `package main` mean?

    `package` is a collection of common source code files.
    
    -> Types of packages
    i) `Executable`, Generates a file that we can run.
    ii) `Reusable`, Code used as "helpers".  Good place to put reusable logic. 

3. Why do we use `import` statements?

    To give our package access to code written in another package

4. What does `import "fmt" ` mean?

    ->The "fmt" package stands for "format" and provides functions for formatted I/O (input/output), such as printing to the console or reading input from the user.
    
5. What's that `func` thing?

    `func` is a keyword to declare a function
    
6. How is the `main.go` file organized?

# Variables
    var card string = "Ace of Strings"
    
    `var` is about create a new variable.
    `card` is the name of the variable will be `card`.
    `string` only a `string` will ever be assigned to this variable.
    `Ace of Strings` Assign the value to the variable.

## Basic Go Types
 1. > `bool` true and false
 2. > `string` -> "Hi!", "How are you?"
 3. > `int` -> 0 , -1000, 9999
 4. > `float64` -> 10.000001, 0.00009, -100.003

## Functions and Return Types

    `func newCard() string {

    }`

    1. `newCard` -> Define a function called **newCard**
    2. `string` -> When executed, this function will return a value of type **string**

## Slices and For Loops
    - Go has two basic data structures for handling lists of records
    
    1. Array > Fixed length list of things
    2. Slice > An array that can grow or shrink and every element in a slice must be of same type

    -> slice[0:2] in which `0` means startIndexIncluding and 2 means upToNotIncluding