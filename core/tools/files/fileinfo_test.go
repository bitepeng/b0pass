package files

import "testing"

func TestFileInfo(t *testing.T) {
	// GetBinPath 获取当前可执行文件的路径
	t1 := GetBinPath()
	t.Log(t1) //C:\...\go-build640703584\b001

	// GetCodePath 获取当前代码文件路径
	t2 := GetCodePath()
	t.Log(t2) //d:\...\b0passv2.0\core\tools

	// IfImage 根据文件名判断是否是图片
	t3 := IfImage("1.png")
	t.Log(t3) //true

	// PathExists 判断文件夹是否存在
	t4, _ := PathExists("C:\\")
	t.Log(t4) //true

	// GetSize 显示友好的文件大小
	t5 := GetSize(12348)
	t.Log(t5) //12K

	t6 := ListDirData("C:\\WCH.CN\\CH341SER", "CH341SER")
	t.Log(t6) //[map[date:11-14 ext:dir indexs:1 name:CH341SER path:CH341SER/CH341SER size:4096 sizes:4K type:dir]]
}
