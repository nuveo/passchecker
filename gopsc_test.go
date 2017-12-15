package gopsc

import "testing"
import "fmt"

func TestCheckStrength(t *testing.T) {
	val := CheckStrength("aaaaaa")
	fmt.Printf("%.9f\n", val)
	val = CheckStrength("aaaaab")
	fmt.Printf("%.9f\n", val)
}
