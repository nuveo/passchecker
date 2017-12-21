package gopsc

import (
	"testing"
)

func TestCheckStrength(t *testing.T) {
	p := New()
	val, _ := p.GetStrength("test long pass")
	if val != 10.932313333 {
		t.Errorf("Expected %v but returned %v", 10.932313333, val)
	}
}
