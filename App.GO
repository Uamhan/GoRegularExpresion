//converts infix expresion to postfix expresion using the shunting yard algorithm
func inFixToPostFix(infix string)string
{
specials := map[rune]int{'*':10, '.':9,'|':8} //map of runes giving order of precedance to these special characters
postfix := []rune{} // array of runes that will end up as the postfix notation 
stack := []rune{} // array of runes used as a stack for the shunting yard algorithm
//for loop that will loop through every character in the input(infix)
for _, r:=range infix
{
	//switch statement that handles brackets and specials characters
	switch{
		//in the case of an open bracket we append it to the stack
	case r =='(':
		stack = append(stack,r)
		//in the case of a closeing bracket we pop items off the stack and append
		//them onto the postfix untill we meet an open bracket we then remove the
		//open bracket fromm the stack
	case r ==')':
		for stack[len(stack)-1]!='('{
			postfix = append(postfix,stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack=stack[:len(stack)-1]
		//as logn as the length of the stack is greater than 0 and the precedance
		//of the current special is greater than the precendance of the rune
		//on the top of the stack we append the top of the stack to postfix
	case specials[r]>0:
		for len(stack)>0 && specials[r] <= specials[s[len(stack)-1]]{
			postfix = append(postfix,stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		s= append(s,r)
		//default case we simply append the current rune to the postfix
	default:
		postfix = append(postfix,r)
	}
}
//after we have completed iterating over the input we pop any remaining items
//on the stack onto the postfic
for len(stack) >0 {
	postfix = append(postfix,stack[len(stack)-1])
	stack = stack[:len(stack)-1]
}
//returns the post fix
return postfix
}

func main() {

}