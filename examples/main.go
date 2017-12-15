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
		v := p.GetStrength(passwd)
		if v > r {
			r = v
		}
		if v > 10.0 {
			fmt.Println(p.GetStrength(passwd), passwd)
		}
	}

	fmt.Println("Max:", r)
	err = scanner.Err()
	if err != nil {
		panic(err)
	}

}
