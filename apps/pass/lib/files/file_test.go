package files

import "testing"

func Test_CheckFile(t *testing.T) {
	header, can := CheckFile("D:\\WorkRoot\\code-github\\B0Pass\\b0passv2.0\\main\\files\\new\\pbwl-service_8096.jar")
	t.Log("header=", header, "canread=", can)
}
