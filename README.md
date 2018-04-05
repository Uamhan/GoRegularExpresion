# GoRegularExpresion
program in the Go programming language that can build a non-deterministic finite automaton from a regular expression, and can use the NFA to check if the regular expression matches any given string of text.
## Installation
The code in this repository is writen in GO. It will require you to download and instal the GO compiler,
Which can be found at https://golang.org/
## Usage
### 1 : You need to clone the repository to your desired location using Git.
```bash
> git clone https://github.com/data-representation/go-examples
```
### 2 : Compile the go code using the following command
```bash
> go build filename.go
```
### 3 : You now have an executable file which can be run with the this command.
```bash
> ./filename.exe
```
## Code explanation
this regular expresion engine recognises the special characters | meaning OR, * one or many, . meaning concatonate this engine can be broken down into three main section

### 1 : converting the users regular expresion from infix notation to postfix notation
in this program we achieve this action by useing the infixToPostfix function which takes the infix string in and returns it in post fix format
![infix to post fix code](https://github.com/Uamhan/GoRegularExpresion/blob/master/images/infixpostfix.PNG)

### 2 : creating an Non deterministic finite automaton from the postfix notation
in this program we achieve this action by using the postFixToNfa function which takes in the postfix string and returns a pointer to a non deterministic finite automaton
![NFA code](https://github.com/Uamhan/GoRegularExpresion/blob/master/images/postfixtonfa.PNG)

### 3 : matching the users string against this automaton
in this program we achive this action by using the poMatch function which takes in the postfix notation string and the striing you wish to match it against and returns true or false on weather they match or not.
![postfix match code](https://github.com/Uamhan/GoRegularExpresion/blob/master/images/pomatch.PNG)

### In Summary
#### the user is asked to enter a regular expresion in the infix notation
#### the user is asked to enter a string to match it with
#### the program takes the regular expresion string and converts it from infix notation to postfix notation
#### the program then creates a non deterministic finite automaton from this post fix regular expresion
#### the user input string is matched against this new automaton
#### the progrogram returns either true or false on weather the string matched the expresion


## Testing

## References
