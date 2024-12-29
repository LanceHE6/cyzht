package random

import "testing"

func TestRandomStr(t *testing.T) {
	t.Log(CreateRandomStr(6, Number))
	t.Log(CreateRandomStr(4, Letter))
	t.Log(CreateRandomStr(8, NumberAndLetter))
}
