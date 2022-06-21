package parenthesis

func IsValid(s string) bool {
	return isValid(s)
}

func isValid(s string) bool {
	st := &stack{[]rune{}, -1}
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			st.push(ch)
		case ')', ']', '}':
			if top, ok := st.pop(); ok && ch == pairs[top] {
				// do nothing
			} else {
				return false
			}
		default:
			// do nothing
		}
	}
	_, ok := st.pop()
	return !ok
}

var pairs = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
}

type stack struct {
	data []rune
	pos  int
}

func (st *stack) push(r rune) {
	st.pos++
	if st.pos >= len(st.data) {
		st.data = append(st.data, r)
	} else {
		st.data[st.pos] = r
	}
}

func (st *stack) peek() (rune, bool) {
	if st.pos == -1 {
		return 0, false
	}
	return st.data[st.pos], true
}

func (st *stack) pop() (rune, bool) {
	if r, ok := st.peek(); ok {
		st.pos--
		return r, ok
	} else {
		return -1, ok
	}
}
