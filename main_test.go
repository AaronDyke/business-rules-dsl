package main

import (
	"testing"
)

func testTokenSliceEq(a, b []token) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBasicLexer(t *testing.T) {
	basics := map[string]token{
		"+": {
			tokenType: Plus,
			startPos:  0,
			text:      "+",
		},
		"-": {
			tokenType: Minus,
			startPos:  0,
			text:      "-",
		},
		"*": {
			tokenType: Multiply,
			startPos:  0,
			text:      "*",
		},
		"/": {
			tokenType: Divide,
			startPos:  0,
			text:      "/",
		},
		"<": {
			tokenType: LessThan,
			startPos:  0,
			text:      "<",
		},
		">": {
			tokenType: GreaterThan,
			startPos:  0,
			text:      ">",
		},
		"<=": {
			tokenType: LessThanEqual,
			startPos:  0,
			text:      "<=",
		},
		">=": {
			tokenType: GreaterThanEqual,
			startPos:  0,
			text:      ">=",
		},
		"AND": {
			tokenType: And,
			startPos:  0,
			text:      "AND",
		},
		"OR": {
			tokenType: Or,
			startPos:  0,
			text:      "OR",
		},
		"TRUE": {
			tokenType: True,
			startPos:  0,
			text:      "TRUE",
		},
		"FALSE": {
			tokenType: False,
			startPos:  0,
			text:      "FALSE",
		},
		"\"String\"": {
			tokenType: String,
			startPos:  0,
			text:      "\"String\"",
		},
		"33": {
			tokenType: Number,
			startPos:  0,
			text:      "33",
		},
		"UNKNOWN_VARIABLE": {
			tokenType: Unknown,
			startPos:  0,
			text:      "UNKNOWN_VARIABLE",
		},
	}

	for basicStr, basicToken := range basics {
		var lexToken = lexer([]byte(basicStr))
		if testTokenSliceEq(lexToken, []token{basicToken}) == false {
			t.Fatalf(`Failed matching basic token %s`, basicStr)
		}
	}
}
