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
	x := s.number()
	for {
		switch s.peek() {
		case '+':
			s.next()
			x += s.number()
			continue
		case '-':
			s.next()
			x -= s.number()
			continue
		case '*':
			s.next()
			x *= s.number()
			continue
		}
		break
	}
	return x
}

func main() {
	exp := "-1-1"
	source := &Source{Str: exp}
	fmt.Printf("%s=%d\n", exp, source.Expr())
}
