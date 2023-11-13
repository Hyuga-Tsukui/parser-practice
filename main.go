package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Source struct {
	Str string
	Pos int
}

var MatchDigitRegexp = regexp.MustCompile(`\d`)

//func (s *Source) number() int {
//	for i, r := range s.Str {
//		if !MatchDigitRegexp.MatchString(string(r)) {
//			s.Pos = i
//			num, _ := strconv.Atoi(s.Str[:i])
//			return num
//		}
//	}
//	num, _ := strconv.Atoi(s.Str)
//	return num
//}

func (s *Source) number() int {
	var buf strings.Builder
	for r := s.peek(); r > 0 && !MatchDigitRegexp.MatchString(string(r)); {
		buf.WriteRune(r)
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
	s.Pos++
}

func (s *Source) expr() int {
	log.Printf("%d\n", s.Pos)
	x := s.number()
	log.Printf("%d\n", s.Pos)
	if s.peek() == '+' {
		s.next()
		log.Printf("%d\n", s.Pos)
		y := s.number()
		log.Printf("%d\n", y)

		x += y
	}
	return x
}

//func number(s *Source) (int, error) {
//	for i, r := range s.Str {
//		if !MatchDigitRegexp.MatchString(string(r)) {
//			s.Pos = i
//			num, err := strconv.Atoi(s.Str[:i])
//			if err != nil {
//				return 0, err
//			}
//			return num, nil
//		}
//	}
//	return strconv.Atoi(s.Str)
//}

func main() {
	exp := "12+34+56"
	source := &Source{Str: exp}
	fmt.Println(source.expr())
	fmt.Println(source)
}
