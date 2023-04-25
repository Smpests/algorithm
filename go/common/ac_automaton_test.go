package common

import (
	"testing"
	"fmt"

	"github.com/anknown/ahocorasick"
	"github.com/stretchr/testify/assert"
)

func TestAC(t *testing.T) {
	dict := [][]rune{{'天', '下', '太', '平'}, {'宁', '有', '种'}}
	machine := new(goahocorasick.Machine)
	if err := machine.Build(dict); err != nil {
        fmt.Println(err)
        return
    }

	content := []rune("天下太平久则久矣")
	terms := machine.MultiPatternSearch(content, false)
	fmt.Printf("%d, %s\n", terms[0].Pos, string(terms[0].Word))
	assert.Equal(t, 1, len(terms))
}