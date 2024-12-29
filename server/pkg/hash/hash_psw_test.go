package hash

import "testing"

func TestHashPsw(t *testing.T) {
	t.Log(HashPsw("123456"))
	t.Log(CheckPsw(HashPsw("123456"), "123456"))
	t.Log(CheckPsw(HashPsw("123456"), "123457"))
}
