package expr

// lexer provides a lexer generator
type lexer struct {
	bytes    []byte
	position int
	ch       byte
}

// newLexer returns a new lexer from expression input
func newLexer(s string) *lexer {
	return &lexer{
		bytes: []byte(s),
	}
}

// next returns next byte from expression
func (l *lexer) next() byte {
	if l.position >= len(l.bytes) {
		l.ch = 0
	} else {
		l.ch = l.bytes[l.position]
		l.position++
	}

	return l.ch
}

// reverse reverses to previous position in expression
func (l *lexer) reverse() {
	if l.position > 0 {
		l.position--
	}
}
