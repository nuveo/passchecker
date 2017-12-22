package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/nuveo/passchecker"
)

func main() {
	js.Global.Set("passchecker", map[string]interface{}{
		"New": New,
	})
}

func New() *js.Object {
	return js.MakeWrapper(&passchecker.Password{
		WeightOfCharacters: passchecker.EnUS,
		CharacterSequence:  passchecker.EnUSSequence,
		StrengthRangeStart: passchecker.EnUSStrengthRangeStart,
		StrengthRangeEnd:   passchecker.EnUSStrengthRangeEnd,
	})
}
