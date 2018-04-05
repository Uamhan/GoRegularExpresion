func inFixToPostFix(infix string)string
{
specials := map[rune]int{'*':10, '.':9,'|':8}
postfix := []rune{} 
stack := []rune{}

for _, r:=range infix
{
	switch{
	case r =='(':
		stack = append(stack,r)
	case r ==')':
		for stack[len(stack)-1]!='('{
			postfix = append(postfix,stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack=stack[:len(stack)-1]
	case specials[r]>0:
		for len(stack)>0 && specials[r] <= specials[s[len(stack)-1]]{
			postfix = append(postfix,stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		s= append(s,r)
	default:
		postfix = append(postfix,r)
	}
}

for len(stack) >0 {
	postfix = append(postfix,stack[len(stack)-1])
	stack = stack[:len(stack)-1]
}

return postfix
}

func main() {

}