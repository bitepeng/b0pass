//go:build darwin || windows
// +build darwin windows

package keys

import (
	"github.com/go-vgo/robotgo"
)

// SendKey 主电脑键盘
func SendKey(k string) {
	switch k {
	case "esc":
		robotgo.KeyTap(`esc`)
	case "wintab":
		robotgo.KeyTap("tab", "command")
	case "tab":
		robotgo.KeyTap("tab")
	case "enter":
		robotgo.KeyTap(`enter`)
	case "f5":
		robotgo.KeyTap(`f5`)
	case "f11":
		robotgo.KeyTap(`f11`)
	case "up":
		robotgo.KeyTap(`up`)
	case "down":
		robotgo.KeyTap(`down`)
	case "left":
		robotgo.KeyTap(`left`)
	case "right":
		robotgo.KeyTap(`right`)
	case "audio_mute":
		robotgo.KeyTap(`audio_mute`)
	case "audio_vol_up":
		robotgo.KeyTap(`audio_vol_up`)
	case "audio_vol_down":
		robotgo.KeyTap(`audio_vol_down`)
	case "mouse_up":
		robotgo.MoveRelative(0, -10)
	case "mouse_down":
		robotgo.MoveRelative(0, 10)
	case "mouse_left":
		robotgo.MoveRelative(-10, 0)
	case "mouse_right":
		robotgo.MoveRelative(10, 0)
	case "mouse_click":
		robotgo.Click("left", true)
	default:
		robotgo.KeyTap(k)
	}
}
