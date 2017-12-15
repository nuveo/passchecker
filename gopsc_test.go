package gopsc

import "testing"
import "fmt"

func TestCheckStrength(t *testing.T) {
	val := CheckStrength("test")
	fmt.Printf("%2.9g\n", val)
}
