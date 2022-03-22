package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"unicode"
	"unicode/utf8"
)

func fourPoint1() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("y"))
	var pc [256]byte
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	ans := 0
	for i := 0; i < len(c1); i++ {
		val1 := c1[i]
		val2 := c2[i]
		ans += int(math.Abs(float64(pc[val1] - pc[val2])))
	}
	fmt.Println("不同的字节数：", ans)
}

func fourPoint2() {
	flag.Parse()
	method := flag.String("s", "SHA256", "输入hash算法")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	switch *method {
	case "SHA256":
		fmt.Println(sha256.Sum256([]byte(line)))
	case "SHA512":
		fmt.Println(sha512.Sum512([]byte(line)))
	case "SHA384":
		fmt.Println(sha512.Sum384([]byte(line)))
	default:
		fmt.Println("错误的编码类型")
	}
}

func fourPoint3(array *[5]int) {
	for i, j := 0, len(*array)-1; i < j; i, j = i+1, j-1 {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}
}

func fourPoint4(s []int, n int) []int {
	for i := 0; i < n; i++ {
		s = append(s, s[i])
	}
	return s[n:]
}

func fourPoint5(strs []string) []string {
	index := 0
	for i, v := range strs {
		if i == 0 || strs[i] != strs[i-1] {
			strs[index] = v
			index++
		}
	}
	return strs[:index]
}

func fourPoint6(bytes []byte) []byte {
	str := string(bytes)
	runes := []rune(str)
	index := 0
	for i, v := range runes {
		if !unicode.IsSpace(v) {
			runes[index] = v
			index++
		} else {
			if i == 0 || !unicode.IsSpace(runes[i-1]) {
				runes[index] = v
				index++
			}
		}
	}
	return []byte(string(runes[:index]))
}

func fourPoint8() {
	counts := make(map[string]int)  // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		} else if unicode.IsSpace(r) {
			counts["space"]++
		} else if unicode.IsNumber(r) {
			counts["number"]++
		} else {
			counts["other"]++
		}
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func wordfreq() {
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file fail,", err)
		return
	}
	defer file.Close()
	wordFeq := make(map[string]int)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordFeq[word]++
	}
	for k, v := range wordFeq {
		fmt.Println("key:", k, " val:", v)
	}
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
	//变量和匿名成员变量相同，还是需要	wheel.Circle.X来访问匿名成员变量
	X int
}

func main() {
	//fourPoint2()
	//array := [...]int{1, 2, 3, 4, 5}
	//fourPoint3(&array)
	//fmt.Println(array)
	//s := []int{1, 2, 3, 4, 5, 6}
	//s = fourPoint4(s, 3)
	//fmt.Println(s)
	//strs := []string{"abc", "abc", "cd", "c", "c", "dd"}
	//strs = fourPoint5(strs)
	//fmt.Println(strs)
	//str := "he  l  lo  中        国"
	//bytes := []byte(str)
	//bytes = fourPoint6(bytes)
	//fmt.Println(string(bytes))
	wordfreq()

}
