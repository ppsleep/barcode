// Package code128 can create Code128 barcodes
package code128

import (
	"fmt"
	"github.com/ppsleep/barcode"
)

// A returns a Code 128-A barcode for the given content
func A(code string) (*barcode.CodesStruct, error) {
	if len(code) > 228 {
		return nil, fmt.Errorf("Content length should be between 1 and 228 runes")
	}
	special := map[int]int{
		241: 102, 242: 97, 243: 96, 244: 101,
	}
	codes := []int{103}
	checker := 103
	code_bytes := []byte(code)
	for k, v := range code_bytes {
		index := 0
		if v >= 32 && v <= 95 {
			index = int(v) - 32
		} else if v >= 0 && v < 32 {
			index = int(v) + 64
		} else if v > 241 && v < 244 {
			codes = append(codes, special[int(v)])
			index = special[int(v)]
		} else {
			return nil, fmt.Errorf("\"%s\" could not be encoded", string(v))
		}
		codes = append(codes, index)
		checker = checker + index*(k+1)
	}
	return encode(codes, checker)
}

// B returns a Code 128-B barcode for the given content
func B(code string) (*barcode.CodesStruct, error) {
	if len(code) > 228 {
		return nil, fmt.Errorf("Content length should be between 1 and 228 runes")
	}
	special := map[int]int{
		241: 102, 242: 97, 243: 96, 244: 100,
	}
	codes := []int{104}
	checker := 104
	code_bytes := []byte(code)
	for k, v := range code_bytes {
		index := 0
		if v >= 32 && v <= 127 {
			index = int(v) - 32
		} else if v > 241 && v < 244 {
			codes = append(codes, special[int(v)])
			index = special[int(v)]
		} else {
			return nil, fmt.Errorf("\"%s\" could not be encoded", string(v))
		}
		codes = append(codes, index)
		checker = checker + index*(k+1)
	}
	return encode(codes, checker)
}

// C returns a Code 128-C barcode for the given content
func C(code string) (*barcode.CodesStruct, error) {
	len := len(code)
	if len > 228 {
		return nil, fmt.Errorf("Content length should be between 1 and 228 runes")
	}
	if len%2 != 0 {
		return nil, fmt.Errorf("Length must be even")
	}
	code_bytes := []byte(code)
	codes := []int{105}
	checker := 105
	for i := 0; i < len; i = i + 2 {
		if code_bytes[i] < 48 || code_bytes[i] > 57 || code_bytes[i+1] < 48 || code_bytes[i+1] > 57 {
			return nil, fmt.Errorf("Only digits allowed")
		}
		item := int((code_bytes[i]-48)*10 + code_bytes[i+1] - 48)
		checker = checker + item*(i/2+1)
		codes = append(codes, item)
	}
	return encode(codes, checker)
}

func encode(codes []int, sum int) (*barcode.CodesStruct, error) {
	codes = append(codes, sum%103)
	codes = append(codes, 106)
	codes = append(codes, 107)
	pattern := pattern()
	var result []barcode.CodeItemStruct
	width := 0
	for _, v := range codes {
		for i := 0; i < 6; i++ {
			line := true
			if i%2 != 0 {
				line = false
			}
			p := pattern[v][i]
			width = width + p
			item := barcode.CodeItemStruct{IsLine: line, Width: p}
			result = append(result, item)
		}
	}
	return &barcode.CodesStruct{
		Codes: result,
		Width: width,
	}, nil
}

func pattern() [][]int {
	return [][]int{
		{2, 1, 2, 2, 2, 2}, // 00
		{2, 2, 2, 1, 2, 2}, // 01
		{2, 2, 2, 2, 2, 1}, // 02
		{1, 2, 1, 2, 2, 3}, // 03
		{1, 2, 1, 3, 2, 2}, // 04
		{1, 3, 1, 2, 2, 2}, // 05
		{1, 2, 2, 2, 1, 3}, // 06
		{1, 2, 2, 3, 1, 2}, // 07
		{1, 3, 2, 2, 1, 2}, // 08
		{2, 2, 1, 2, 1, 3}, // 09
		{2, 2, 1, 3, 1, 2}, // 10
		{2, 3, 1, 2, 1, 2}, // 11
		{1, 1, 2, 2, 3, 2}, // 12
		{1, 2, 2, 1, 3, 2}, // 13
		{1, 2, 2, 2, 3, 1}, // 14
		{1, 1, 3, 2, 2, 2}, // 15
		{1, 2, 3, 1, 2, 2}, // 16
		{1, 2, 3, 2, 2, 1}, // 17
		{2, 2, 3, 2, 1, 1}, // 18
		{2, 2, 1, 1, 3, 2}, // 19
		{2, 2, 1, 2, 3, 1}, // 20
		{2, 1, 3, 2, 1, 2}, // 21
		{2, 2, 3, 1, 1, 2}, // 22
		{3, 1, 2, 1, 3, 1}, // 23
		{3, 1, 1, 2, 2, 2}, // 24
		{3, 2, 1, 1, 2, 2}, // 25
		{3, 2, 1, 2, 2, 1}, // 26
		{3, 1, 2, 2, 1, 2}, // 27
		{3, 2, 2, 1, 1, 2}, // 28
		{3, 2, 2, 2, 1, 1}, // 29
		{2, 1, 2, 1, 2, 3}, // 30
		{2, 1, 2, 3, 2, 1}, // 31
		{2, 3, 2, 1, 2, 1}, // 32
		{1, 1, 1, 3, 2, 3}, // 33
		{1, 3, 1, 1, 2, 3}, // 34
		{1, 3, 1, 3, 2, 1}, // 35
		{1, 1, 2, 3, 1, 3}, // 36
		{1, 3, 2, 1, 1, 3}, // 37
		{1, 3, 2, 3, 1, 1}, // 38
		{2, 1, 1, 3, 1, 3}, // 39
		{2, 3, 1, 1, 1, 3}, // 40
		{2, 3, 1, 3, 1, 1}, // 41
		{1, 1, 2, 1, 3, 3}, // 42
		{1, 1, 2, 3, 3, 1}, // 43
		{1, 3, 2, 1, 3, 1}, // 44
		{1, 1, 3, 1, 2, 3}, // 45
		{1, 1, 3, 3, 2, 1}, // 46
		{1, 3, 3, 1, 2, 1}, // 47
		{3, 1, 3, 1, 2, 1}, // 48
		{2, 1, 1, 3, 3, 1}, // 49
		{2, 3, 1, 1, 3, 1}, // 50
		{2, 1, 3, 1, 1, 3}, // 51
		{2, 1, 3, 3, 1, 1}, // 52
		{2, 1, 3, 1, 3, 1}, // 53
		{3, 1, 1, 1, 2, 3}, // 54
		{3, 1, 1, 3, 2, 1}, // 55
		{3, 3, 1, 1, 2, 1}, // 56
		{3, 1, 2, 1, 1, 3}, // 57
		{3, 1, 2, 3, 1, 1}, // 58
		{3, 3, 2, 1, 1, 1}, // 59
		{3, 1, 4, 1, 1, 1}, // 60
		{2, 2, 1, 4, 1, 1}, // 61
		{4, 3, 1, 1, 1, 1}, // 62
		{1, 1, 1, 2, 2, 4}, // 63
		{1, 1, 1, 4, 2, 2}, // 64
		{1, 2, 1, 1, 2, 4}, // 65
		{1, 2, 1, 4, 2, 1}, // 66
		{1, 4, 1, 1, 2, 2}, // 67
		{1, 4, 1, 2, 2, 1}, // 68
		{1, 1, 2, 2, 1, 4}, // 69
		{1, 1, 2, 4, 1, 2}, // 70
		{1, 2, 2, 1, 1, 4}, // 71
		{1, 2, 2, 4, 1, 1}, // 72
		{1, 4, 2, 1, 1, 2}, // 73
		{1, 4, 2, 2, 1, 1}, // 74
		{2, 4, 1, 2, 1, 1}, // 75
		{2, 2, 1, 1, 1, 4}, // 76
		{4, 1, 3, 1, 1, 1}, // 77
		{2, 4, 1, 1, 1, 2}, // 78
		{1, 3, 4, 1, 1, 1}, // 79
		{1, 1, 1, 2, 4, 2}, // 80
		{1, 2, 1, 1, 4, 2}, // 81
		{1, 2, 1, 2, 4, 1}, // 82
		{1, 1, 4, 2, 1, 2}, // 83
		{1, 2, 4, 1, 1, 2}, // 84
		{1, 2, 4, 2, 1, 1}, // 85
		{4, 1, 1, 2, 1, 2}, // 86
		{4, 2, 1, 1, 1, 2}, // 87
		{4, 2, 1, 2, 1, 1}, // 88
		{2, 1, 2, 1, 4, 1}, // 89
		{2, 1, 4, 1, 2, 1}, // 90
		{4, 1, 2, 1, 2, 1}, // 91
		{1, 1, 1, 1, 4, 3}, // 92
		{1, 1, 1, 3, 4, 1}, // 93
		{1, 3, 1, 1, 4, 1}, // 94
		{1, 1, 4, 1, 1, 3}, // 95
		{1, 1, 4, 3, 1, 1}, // 96
		{4, 1, 1, 1, 1, 3}, // 97
		{4, 1, 1, 3, 1, 1}, // 98
		{1, 1, 3, 1, 4, 1}, // 99
		{1, 1, 4, 1, 3, 1}, // 100
		{3, 1, 1, 1, 4, 1}, // 101
		{4, 1, 1, 1, 3, 1}, // 102
		{2, 1, 1, 4, 1, 2}, // 103 START A
		{2, 1, 1, 2, 1, 4}, // 104 START B
		{2, 1, 1, 2, 3, 2}, // 105 START C
		{2, 3, 3, 1, 1, 1}, // STOP
		{2, 0, 0, 0, 0, 0}, // END
	}
}
