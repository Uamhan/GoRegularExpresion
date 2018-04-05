package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//state struct which contains a symbol of type rune and two pointers for edges
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

//nfa stuct which has a pointer to the start and end of the nfa
type nfa struct {
	initial *state
	accept  *state
}

//main method
func main() {
	//creates scanner to take user input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter you regular expresion (infix notation)  or: |  one or more: * concatonate : .")
	scanner.Scan()
	//takes in user input to infix string
	infix := scanner.Text()
	fmt.Println("enter string you wish to check against your expresion")
	scanner.Scan()
	//takes in user input to checkString string
	checkString := scanner.Text()
	//converts the infix notation to postfix notation
	postfix := infixToPostfix(infix)
	//creates a nfa from postfix and trys to match checkstring with it
	fmt.Println(poMatch(strings.ToLower(postfix), strings.ToLower(checkString)))
}

//addes all reachable states from a given state to a state aray
//recursive function to navigate through all choices
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)
	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}

//function that trys to match regular expresion with a string
func poMatch(postfix string, s string) bool {
	//sets isMatch to false by default
	isMatch := false
	//using the postfixTonNfa function we get the nfa of the input
	poNfa := postfixToNfa(postfix)
	//initialise current and next nfa pointer arrays
	current := []*state{}
	next := []*state{}
	//sets current to out helper function add state
	current = addState(current[:], poNfa.initial, poNfa.accept)
	//loops through each rune in the input string
	for _, r := range s {
		//loops through each nfa pointer in current
		for _, c := range current {
			//if the symbol of the current nfa point equals the rune r
			//we set next to equal our add state function with the inputs below
			if c.symbol == r {
				next = addState(next[:], c.edge1, poNfa.accept)
			}
		}
		//we set the current to = next and the next to equal a blank pointer array
		current, next = next, []*state{}
	}
	//loops through each nfa pointer in current
	for _, c := range current {
		// if the current state is equal to the accept state of the nfa is match is set to true
		if c == poNfa.accept {
			isMatch = true
			break
		}
	}

	return isMatch
}

//postfix notation is used to create a NFA
func postfixToNfa(postfix string) *nfa {
	//stack of nfa pointers
	nfaStack := []*nfa{}
	//for loop that loops through each rune of the function input
	for _, r := range postfix {
		//switch statement for each rune that handles . | * seperately than everything else
		switch r {

		case '.':
			//pops two fragments off of the stack concatentates the two and addes it back onto the stack
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial
			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '|':
			//pops two fragments off the stack creates a new initial state from frag1 and 2's initial states
			//creates an empty accept state  then adds a point to these new values to the stack
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		case '*':
			//pops 1 fragment off the stack creates and empty accept state
			//creates a new initial state with edge1 being the frag initial
			//state and edge 2 being a referance to the new accept state
			//then appends a pointer to a nfa with these new values to the stack
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		default:
			//creates an empty accept state
			//creates a new initial state wiht symbol equaling the current rune
			//and edge 1 equaling the new accept state then appends a pointer to
			//a new nfa using these new values
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}

	//check only one
	return nfaStack[0]
}

//converts infix expresion to postfix expresion using the shunting yard algorithm
func infixToPostfix(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8} //map of runes giving order of precedance to these special characters
	postfix := []rune{}                               // array of runes that will end up as the postfix notation
	stack := []rune{}                                 // array of runes used as a stack for the shunting yard algorithm
	//for loop that will loop through every character in the input(infix)
	for _, r := range infix {
		//switch statement that handles brackets and specials characters
		switch {
		//in the case of an open bracket we append it to the stack
		case r == '(':
			stack = append(stack, r)
			//in the case of a closeing bracket we pop items off the stack and append
			//them onto the postfix untill we meet an open bracket we then remove the
			//open bracket fromm the stack
		case r == ')':
			for stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			//as logn as the length of the stack is greater than 0 and the precedance
			//of the current special is greater than the precendance of the rune
			//on the top of the stack we append the top of the stack to postfix
		case specials[r] > 0:
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, r)
			//default case we simply append the current rune to the postfix
		default:
			postfix = append(postfix, r)
		}
	}
	//after we have completed iterating over the input we pop any remaining items
	//on the stack onto the postfic
	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	//returns the post fix
	returnString := string(postfix)
	return returnString
}
