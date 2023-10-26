//go:build !darwin && !windows
// +build !darwin,!windows

package keys

// CmdKey 主电脑键盘
func SendKey(k string) {
	//Linux暂不支持
}
