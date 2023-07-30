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