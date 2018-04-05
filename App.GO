package main

import "fmt"

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
	nfa := postfixToNfa("ab.c*|")
	fmt.Println(nfa)
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
