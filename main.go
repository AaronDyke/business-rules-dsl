package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Start Reading file
	var fileName = "uppercase_and_boolean.brr"
	data, err := os.ReadFile(fmt.Sprintf("./example_rules/%s", fileName))
	check(err)
	// End Reading file
	var tokens = lexer(data)
	fmt.Println(tokens)
}

type TokenType int64

const (
	Unknown TokenType = iota
	LessThan
	LessThanEqual
	GreaterThan
	GreaterThanEqual
	And
	Or
	True
	False
	Plus
	Minus
	Multiply
	Divide
	Number
	String
	Variable
)

type token struct {
	tokenType TokenType
	startPos  int
	text      string
}

func lexer(data []byte) []token {
	var curPos = 0
	var tokens []token
	for curPos < len(data) {
		var tokenStartPosition = curPos
		var s = string(data[curPos])
		switch s {
		case " ":
			curPos = curPos + 1
		case "+":
			curPos = curPos + 1
			tokens = append(tokens, token{tokenType: Plus, startPos: tokenStartPosition, text: s})
		case "-":
			curPos = curPos + 1
			tokens = append(tokens, token{tokenType: Minus, startPos: tokenStartPosition, text: s})
		case "*":
			curPos = curPos + 1
			tokens = append(tokens, token{tokenType: Multiply, startPos: tokenStartPosition, text: s})
		case "/":
			curPos = curPos + 1
			tokens = append(tokens, token{tokenType: Divide, startPos: tokenStartPosition, text: s})
		case ">":
			curPos = curPos + 1
			if curPos < len(data) && string(data[curPos]) == "=" {
				// token is >=
				tokens = append(tokens, token{tokenType: GreaterThanEqual, startPos: tokenStartPosition, text: ">="})
				curPos = curPos + 1
			} else {
				// token is >
				tokens = append(tokens, token{tokenType: GreaterThan, startPos: tokenStartPosition, text: ">"})
			}
		case "<":
			curPos = curPos + 1
			if curPos < len(data) && string(data[curPos]) == "=" {
				// token is <=
				tokens = append(tokens, token{tokenType: LessThanEqual, startPos: tokenStartPosition, text: "<="})
				curPos = curPos + 1
			} else {
				// token is <
				tokens = append(tokens, token{tokenType: LessThan, startPos: tokenStartPosition, text: "<"})
			}
		default:
			var b bytes.Buffer
			for curPos < len(data) && string(data[curPos]) != " " {
				b.WriteByte(data[curPos])
				curPos = curPos + 1
			}

			if isNumber(b.String()) {
				// token is string
				tokens = append(tokens, token{tokenType: Number, startPos: tokenStartPosition, text: b.String()})
			} else if strings.ToUpper(b.String()) == "OR" {
				// token is OR
				tokens = append(tokens, token{tokenType: Or, startPos: tokenStartPosition, text: b.String()})
			} else if strings.ToUpper(b.String()) == "AND" {
				// token is AND
				tokens = append(tokens, token{tokenType: And, startPos: tokenStartPosition, text: b.String()})
			} else if strings.ToUpper(b.String()) == "TRUE" {
				// token is TRUE TODO: determine if support for case insensitive True is good or not
				tokens = append(tokens, token{tokenType: True, startPos: tokenStartPosition, text: b.String()})
			} else if strings.ToUpper(b.String()) == "FALSE" {
				// token is FALSE TODO: determine if support for case insensitive False is good or not
				tokens = append(tokens, token{tokenType: False, startPos: tokenStartPosition, text: b.String()})
			} else if isString(b.String()) {
				// token is "STRING"
				tokens = append(tokens, token{tokenType: String, startPos: tokenStartPosition, text: b.String()})
			} else if isVariable(b.String()) {
				// token is a variable
				tokens = append(tokens, token{tokenType: Variable, startPos: tokenStartPosition, text: b.String()})
			} else {
				// unknown token
				tokens = append(tokens, token{tokenType: Unknown, startPos: tokenStartPosition, text: b.String()})
			}
		}
	}
	return tokens
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isString(s string) bool {
	return (strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")) || (strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'"))
}

func isVariable(s string) bool {
	var HARD_CODED_VARIABLES = [1]string{"TEST"}
	for _, varName := range HARD_CODED_VARIABLES {
		if s == varName {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
