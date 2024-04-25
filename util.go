// Copyright 2024 itpey
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/fatih/color"
	"github.com/micmonay/keybd_event"
)

func showVersion() error {
	fmt.Printf("Key Control version %s\n", appVersion)
	return nil
}

// simulateKeyPress simulates a key press event based on the provided parameters.
func simulateKeyPress(vkCode int, hasShift bool, hasCtrl bool, hasAlt bool, hasSuper bool) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Fatalf(color.RedString("Error: Creating key binding: %v"), err)
	}

	kb.SetKeys(vkCode)
	kb.HasSHIFT(hasShift)
	kb.HasCTRLR(hasCtrl)
	kb.HasALT(hasAlt)
	kb.HasSuper(hasSuper)

	err = kb.Launching()
	if err != nil {
		log.Fatalf(color.RedString("Error: Pressing key: %v", err))
	}

	time.Sleep(100 * time.Millisecond)
}

// getKeyCodeByKeyName maps key names to virtual key codes.
func getKeyCodeByKeyName(keyName string) int {
	switch keyName {
	case "SP1":
		return keybd_event.VK_SP1
	case "SP2":
		return keybd_event.VK_SP2
	case "SP3":
		return keybd_event.VK_SP3
	case "SP4":
		return keybd_event.VK_SP4
	case "SP5":
		return keybd_event.VK_SP5
	case "SP6":
		return keybd_event.VK_SP6
	case "SP7":
		return keybd_event.VK_SP7
	case "SP8":
		return keybd_event.VK_SP8
	case "SP9":
		return keybd_event.VK_SP9
	case "SP10":
		return keybd_event.VK_SP10
	case "SP11":
		return keybd_event.VK_SP11
	case "SP12":
		return keybd_event.VK_SP12
	case "ESC":
		return keybd_event.VK_ESC
	case "1":
		return keybd_event.VK_1
	case "2":
		return keybd_event.VK_2
	case "3":
		return keybd_event.VK_3
	case "4":
		return keybd_event.VK_4
	case "5":
		return keybd_event.VK_5
	case "6":
		return keybd_event.VK_6
	case "7":
		return keybd_event.VK_7
	case "8":
		return keybd_event.VK_8
	case "9":
		return keybd_event.VK_9
	case "0":
		return keybd_event.VK_0
	case "Q":
		return keybd_event.VK_Q
	case "W":
		return keybd_event.VK_W
	case "E":
		return keybd_event.VK_E
	case "R":
		return keybd_event.VK_R
	case "T":
		return keybd_event.VK_T
	case "Y":
		return keybd_event.VK_Y
	case "U":
		return keybd_event.VK_U
	case "I":
		return keybd_event.VK_I
	case "O":
		return keybd_event.VK_O
	case "P":
		return keybd_event.VK_P
	case "A":
		return keybd_event.VK_A
	case "S":
		return keybd_event.VK_S
	case "D":
		return keybd_event.VK_D
	case "F":
		return keybd_event.VK_F
	case "G":
		return keybd_event.VK_G
	case "H":
		return keybd_event.VK_H
	case "J":
		return keybd_event.VK_J
	case "K":
		return keybd_event.VK_K
	case "L":
		return keybd_event.VK_L
	case "Z":
		return keybd_event.VK_Z
	case "X":
		return keybd_event.VK_X
	case "C":
		return keybd_event.VK_C
	case "V":
		return keybd_event.VK_V
	case "B":
		return keybd_event.VK_B
	case "N":
		return keybd_event.VK_N
	case "M":
		return keybd_event.VK_M
	case "F1":
		return keybd_event.VK_F1
	case "F2":
		return keybd_event.VK_F2
	case "F3":
		return keybd_event.VK_F3
	case "F4":
		return keybd_event.VK_F4
	case "F5":
		return keybd_event.VK_F5
	case "F6":
		return keybd_event.VK_F6
	case "F7":
		return keybd_event.VK_F7
	case "F8":
		return keybd_event.VK_F8
	case "F9":
		return keybd_event.VK_F9
	case "F10":
		return keybd_event.VK_F10
	case "F11":
		return keybd_event.VK_F11
	case "F12":
		return keybd_event.VK_F12
	case "F13":
		return keybd_event.VK_F13
	case "F14":
		return keybd_event.VK_F14
	case "F15":
		return keybd_event.VK_F15
	case "F16":
		return keybd_event.VK_F16
	case "F17":
		return keybd_event.VK_F17
	case "F18":
		return keybd_event.VK_F18
	case "F19":
		return keybd_event.VK_F19
	case "F20":
		return keybd_event.VK_F20
	case "F21":
		return keybd_event.VK_F21
	case "F22":
		return keybd_event.VK_F22
	case "F23":
		return keybd_event.VK_F23
	case "F24":
		return keybd_event.VK_F24
	case "NUMLOCK":
		return keybd_event.VK_NUMLOCK
	case "SCROLLLOCK":
		return keybd_event.VK_SCROLLLOCK
	case "RESERVED":
		return keybd_event.VK_RESERVED
	case "MINUS":
		return keybd_event.VK_MINUS
	case "EQUAL":
		return keybd_event.VK_EQUAL
	case "BACKSPACE":
		return keybd_event.VK_BACKSPACE
	case "TAB":
		return keybd_event.VK_TAB
	case "LEFTBRACE":
		return keybd_event.VK_LEFTBRACE
	case "RIGHTBRACE":
		return keybd_event.VK_RIGHTBRACE
	case "ENTER":
		return keybd_event.VK_ENTER
	case "SEMICOLON":
		return keybd_event.VK_SEMICOLON
	case "APOSTROPHE":
		return keybd_event.VK_APOSTROPHE
	case "GRAVE":
		return keybd_event.VK_GRAVE
	case "BACKSLASH":
		return keybd_event.VK_BACKSLASH
	case "COMMA":
		return keybd_event.VK_COMMA
	case "DOT":
		return keybd_event.VK_DOT
	case "SLASH":
		return keybd_event.VK_SLASH
	case "KPASTERISK":
		return keybd_event.VK_KPASTERISK
	case "SPACE":
		return keybd_event.VK_SPACE
	case "CAPSLOCK":
		return keybd_event.VK_CAPSLOCK
	case "KP0":
		return keybd_event.VK_KP0
	case "KP1":
		return keybd_event.VK_KP1
	case "KP2":
		return keybd_event.VK_KP2
	case "KP3":
		return keybd_event.VK_KP3
	case "KP4":
		return keybd_event.VK_KP4
	case "KP5":
		return keybd_event.VK_KP5
	case "KP6":
		return keybd_event.VK_KP6
	case "KP7":
		return keybd_event.VK_KP7
	case "KP8":
		return keybd_event.VK_KP8
	case "KP9":
		return keybd_event.VK_KP9
	case "KPMINUS":
		return keybd_event.VK_KPMINUS
	case "KPPLUS":
		return keybd_event.VK_KPPLUS
	case "KPDOT":
		return keybd_event.VK_KPDOT
	case "LBUTTON":
		return keybd_event.VK_LBUTTON
	case "RBUTTON":
		return keybd_event.VK_RBUTTON
	case "CANCEL":
		return keybd_event.VK_CANCEL
	case "MBUTTON":
		return keybd_event.VK_MBUTTON
	case "XBUTTON1":
		return keybd_event.VK_XBUTTON1
	case "XBUTTON2":
		return keybd_event.VK_XBUTTON2
	case "BACK":
		return keybd_event.VK_BACK
	case "CLEAR":
		return keybd_event.VK_CLEAR
	case "PAUSE":
		return keybd_event.VK_PAUSE
	case "CAPITAL":
		return keybd_event.VK_CAPITAL
	case "KANA":
		return keybd_event.VK_KANA
	case "HANGUEL":
		return keybd_event.VK_HANGUEL
	case "HANGUL":
		return keybd_event.VK_HANGUL
	case "JUNJA":
		return keybd_event.VK_JUNJA
	case "FINAL":
		return keybd_event.VK_FINAL
	case "HANJA":
		return keybd_event.VK_HANJA
	case "KANJI":
		return keybd_event.VK_KANJI
	case "CONVERT":
		return keybd_event.VK_CONVERT
	case "NONCONVERT":
		return keybd_event.VK_NONCONVERT
	case "ACCEPT":
		return keybd_event.VK_ACCEPT
	case "MODECHANGE":
		return keybd_event.VK_MODECHANGE
	case "PAGEUP":
		return keybd_event.VK_PAGEUP
	case "PAGEDOWN":
		return keybd_event.VK_PAGEDOWN
	case "END":
		return keybd_event.VK_END
	case "HOME":
		return keybd_event.VK_HOME
	case "LEFT":
		return keybd_event.VK_LEFT
	case "UP":
		return keybd_event.VK_UP
	case "RIGHT":
		return keybd_event.VK_RIGHT
	case "DOWN":
		return keybd_event.VK_DOWN
	case "SELECT":
		return keybd_event.VK_SELECT
	case "PRINT":
		return keybd_event.VK_PRINT
	case "EXECUTE":
		return keybd_event.VK_EXECUTE
	case "SNAPSHOT":
		return keybd_event.VK_SNAPSHOT
	case "INSERT":
		return keybd_event.VK_INSERT
	case "DELETE":
		return keybd_event.VK_DELETE
	case "HELP":
		return keybd_event.VK_HELP
	case "SCROLL":
		return keybd_event.VK_SCROLL
	case "LMENU":
		return keybd_event.VK_LMENU
	case "RMENU":
		return keybd_event.VK_RMENU
	case "BROWSER_BACK":
		return keybd_event.VK_BROWSER_BACK
	case "BROWSER_FORWARD":
		return keybd_event.VK_BROWSER_FORWARD
	case "BROWSER_REFRESH":
		return keybd_event.VK_BROWSER_REFRESH
	case "BROWSER_STOP":
		return keybd_event.VK_BROWSER_STOP
	case "BROWSER_SEARCH":
		return keybd_event.VK_BROWSER_SEARCH
	case "BROWSER_FAVORITES":
		return keybd_event.VK_BROWSER_FAVORITES
	case "BROWSER_HOME":
		return keybd_event.VK_BROWSER_HOME
	case "VOLUME_MUTE":
		return keybd_event.VK_VOLUME_MUTE
	case "VOLUME_DOWN":
		return keybd_event.VK_VOLUME_DOWN
	case "VOLUME_UP":
		return keybd_event.VK_VOLUME_UP
	case "MEDIA_NEXT_TRACK":
		return keybd_event.VK_MEDIA_NEXT_TRACK
	case "MEDIA_PREV_TRACK":
		return keybd_event.VK_MEDIA_PREV_TRACK
	case "MEDIA_STOP":
		return keybd_event.VK_MEDIA_STOP
	case "MEDIA_PLAY_PAUSE":
		return keybd_event.VK_MEDIA_PLAY_PAUSE
	case "LAUNCH_MAIL":
		return keybd_event.VK_LAUNCH_MAIL
	case "LAUNCH_MEDIA_SELECT":
		return keybd_event.VK_LAUNCH_MEDIA_SELECT
	case "LAUNCH_APP1":
		return keybd_event.VK_LAUNCH_APP1
	case "LAUNCH_APP2":
		return keybd_event.VK_LAUNCH_APP2
	case "OEM_1":
		return keybd_event.VK_OEM_1
	case "OEM_PLUS":
		return keybd_event.VK_OEM_PLUS
	case "OEM_COMMA":
		return keybd_event.VK_OEM_COMMA
	case "OEM_MINUS":
		return keybd_event.VK_OEM_MINUS
	case "OEM_PERIOD":
		return keybd_event.VK_OEM_PERIOD
	case "OEM_2":
		return keybd_event.VK_OEM_2
	case "OEM_3":
		return keybd_event.VK_OEM_3
	case "OEM_4":
		return keybd_event.VK_OEM_4
	case "OEM_5":
		return keybd_event.VK_OEM_5
	case "OEM_6":
		return keybd_event.VK_OEM_6
	case "OEM_7":
		return keybd_event.VK_OEM_7
	case "OEM_8":
		return keybd_event.VK_OEM_8
	case "OEM_102":
		return keybd_event.VK_OEM_102
	case "PROCESSKEY":
		return keybd_event.VK_PROCESSKEY
	case "PACKET":
		return keybd_event.VK_PACKET
	case "ATTN":
		return keybd_event.VK_ATTN
	case "CRSEL":
		return keybd_event.VK_CRSEL
	case "EXSEL":
		return keybd_event.VK_EXSEL
	case "EREOF":
		return keybd_event.VK_EREOF
	case "PLAY":
		return keybd_event.VK_PLAY
	case "ZOOM":
		return keybd_event.VK_ZOOM
	case "NONAME":
		return keybd_event.VK_NONAME
	case "PA1":
		return keybd_event.VK_PA1
	case "OEM_CLEAR":
		return keybd_event.VK_OEM_CLEAR
	default:
		return 0
	}
}

func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
