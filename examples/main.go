package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/crgimenes/gopsc"
)

func main() {
	f, err := os.Open("badpasswordlist.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := gopsc.New()
	r := 0.0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		passwd := scanner.Text()
		v, _ := p.GetStrength(passwd)
		if v > r {
			r = v
		}
		if v > 6.0 && v < 8.0 {
			fmt.Printf("%02.5f\t%v\n", v, passwd)
		}
	}

	//fmt.Println("Max:", r)
	err = scanner.Err()
	if err != nil {
		panic(err)
	}

}
