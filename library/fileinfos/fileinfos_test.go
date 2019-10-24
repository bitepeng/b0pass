package fileinfos

import "testing"

func TestIfImage(t *testing.T) {
	tmp := IfImage("212/wqewqe/sadsad.png")
	t.Log(tmp)

	tmp = IfImage("212/wqewqe/sadsad.pdf")
	t.Log(tmp)
}
