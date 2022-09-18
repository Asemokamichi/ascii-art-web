package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strings"
)

var CorrectBanner = map[string]string{
	"standard.txt":   "ac85e83127e49ec42487f272d9b9db8b",
	"shadow.txt":     "a49d5fcb0d5c59b2e77674aa3ab8bbb1",
	"thinkertoy.txt": "bf1d925662e40f5278b26a0531bfdb63",
}

// SetAsciiArt function is
func SetAsciiArt(s, banner string) (string, int) {
	if s == "" || banner != "standard.txt" && banner != "shadow.txt" && banner != "thinkertoy.txt" {
		return http.StatusText(http.StatusBadRequest), http.StatusBadRequest
	}
	content, err := ioutil.ReadFile("internal/utils/banner//" + banner)
	if err != nil || !CheckingTheHashCode(content, CorrectBanner[banner]) {
		return http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError
	}
	arr := SplitLines(strings.Split(string(content), "\n"))
	// func() {}
	m := make(map[rune][]string)
	for i, w := range arr {
		m[rune(i+32)] = w
	}

	for _, r := range s {
		if (r < 32 || r > 126) && r != 10 && r != 13 {
			return http.StatusText(http.StatusBadRequest), http.StatusBadRequest
		}
	}

	result := [][8]string{}
	lines := [8]string{}

	str := strings.Split(s, "\r\n")
	for _, w := range str {
		for _, q := range w {
			for j := range lines {
				lines[j] += m[q][j]
			}
		}
		result = append(result, lines)
		lines = [8]string{}
	}

	return ToString(result), 0
}

// ToString function is
func ToString(arr [][8]string) string {
	s := ""
	for _, w := range arr {
		for _, q := range w {
			if len(q) == 0 {
				s += "\n"
				break
			}
			s += q + "\n"
		}
	}
	return s
}

// SplitLine funtion is
func SplitLines(lines []string) [][]string {
	symbol := []string{}
	symbols := [][]string{}
	for i, line := range lines {
		if line != "" {
			symbol = append(symbol, line)
		}
		if (line == "" || i == len(lines)-1) && len(symbol) > 0 {
			symbols = append(symbols, symbol)
			symbol = []string{}
		}
	}
	return symbols
}

func CheckingTheHashCode(content []byte, hash string) bool {
	h := md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil)) == hash
}
