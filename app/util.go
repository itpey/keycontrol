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

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
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
	case "MINUS":
		return keybd_event.VK_MINUS
	case "EQUAL":
		return keybd_event.VK_EQUAL
	case "TAB":
		return keybd_event.VK_TAB
	case "ENTER":
		return keybd_event.VK_ENTER
	case "SEMICOLON":
		return keybd_event.VK_SEMICOLON
	case "GRAVE":
		return keybd_event.VK_GRAVE
	case "BACKSLASH":
		return keybd_event.VK_BACKSLASH
	case "COMMA":
		return keybd_event.VK_COMMA
	case "SLASH":
		return keybd_event.VK_SLASH
	case "SPACE":
		return keybd_event.VK_SPACE
	case "CAPSLOCK":
		return keybd_event.VK_CAPSLOCK
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
	case "DELETE":
		return keybd_event.VK_DELETE
	case "HELP":
		return keybd_event.VK_HELP
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
