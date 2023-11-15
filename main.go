package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Source struct {
	Str string
	Pos int
}

var MatchDigitRegexp = regexp.MustCompile(`\d`)

func (s *Source) number() int {
	var buf strings.Builder
	for s.peek() > 0 && MatchDigitRegexp.MatchString(string(s.peek())) {
		buf.WriteRune(s.peek())
		s.next()
	}
	num, _ := strconv.Atoi(buf.String())
	return num
}

func (s *Source) peek() rune {
	if s.Pos < len(s.Str) {
		return rune(s.Str[s.Pos])
	}
	return -1
}

func (s *Source) next() {
	s.Pos += 1
}

func (s *Source) Expr() int {
	x := s.term()
	for {
		switch s.peek() {
		case '+':
			s.next()
			x += s.term()
			continue
		case '-':
			s.next()
			x -= s.term()
			continue
		}
		break
	}
	return x
}

func (s *Source) term() int {
	x := s.factor()
	for {
		switch s.peek() {
		case '*':
			s.next()
			x *= s.factor()
			continue
		case '/':
			s.next()
			x /= s.factor()
			continue
		}
		break
	}
	return x
}

// factor = [spaces], ("(", <Exper>, ")") | number, [spaces].
func (s *Source) factor() int {
	var ret int
	s.spaces()
	if s.peek() == '(' {
		s.next()
		ret = s.Expr()
		if s.peek() == ')' {
			s.next()
		}
	} else {
		ret = s.number()
	}
	s.spaces()
	return ret
}

func (s *Source) spaces() {
	for s.peek() == ' ' {
		s.next()
	}
}

func main() {
	exp := "-1-1"
	source := &Source{Str: exp}
	fmt.Printf("%s=%d\n", exp, source.Expr())
}
