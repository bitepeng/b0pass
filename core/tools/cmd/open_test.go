package cmd

import "testing"

func TestOpen(t *testing.T) {
	if err := Open("https://baidu.com"); err != nil {
		t.Error(err)
	}
}

func TestOpen2(t *testing.T) {
	if err := Open("/Users/"); err != nil {
		t.Error(err)
	}
}
