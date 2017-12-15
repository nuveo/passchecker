package gopsc

import (
	"fmt"
	"testing"
)

func TestCheckStrength(t *testing.T) {
	p := New()
	val := p.GetStrength("aaaaaa")
	fmt.Printf("%.9f\n", val)
	val = p.GetStrength("aaaaab")
	fmt.Printf("%.9f\n", val)
}
