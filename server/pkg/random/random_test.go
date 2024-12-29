package random

import "testing"

func TestRandomStr(t *testing.T) {
	t.Log(CreateRandomStr(1, 8))
	t.Log(CreateRandomStr(2, 6))
	t.Log(CreateRandomStr(3, 6))
}
