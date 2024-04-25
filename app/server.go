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
	"net"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

var authToken string
var port string

// healthHandler responds with an HTTP status OK for health checks.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// keyPressHandler simulates a key press based on the provided key code and modifiers.
func keyPressHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Check if the request contains the correct token
	token := r.Header.Get("Authorization")
	if token == "" || token != authToken {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	keyCode := r.URL.Query().Get("keyCode")
	shift := r.URL.Query().Get("shift") == "true"
	ctrl := r.URL.Query().Get("ctrl") == "true"
	alt := r.URL.Query().Get("alt") == "true"
	super := r.URL.Query().Get("super") == "true"

	if keyCode == "" {
		http.Error(w, "keyCode parameter is required", http.StatusBadRequest)
		return
	}

	vkCode := getKeyCodeByKeyName(keyCode)
	if vkCode == 0 {
		http.Error(w, "Invalid keyCode provided", http.StatusBadRequest)
		return
	}
	var keyCodes []int
	keyCodes = append(keyCodes, vkCode)
	simulateKeyPress(keyCodes, shift, ctrl, alt, super)

	logMessage := fmt.Sprintf("Key press simulated for keyCode: %s", keyCode)
	if ctrl {
		logMessage += " | Ctrl"
	}
	if alt {
		logMessage += " | Alt"
	}
	if shift {
		logMessage += " | Shift"
	}
	if super {
		logMessage += " | Super"
	}

	log.Print(logMessage)
	fmt.Fprint(w, logMessage)

}

// keyPressHandler simulates a key press based on the provided key code and modifiers.
func multiKeyHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Check if the request contains the correct token
	token := r.Header.Get("Authorization")
	if token == "" || token != authToken {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	keysParam := r.URL.Query().Get("keyCodes")
	shift := r.URL.Query().Get("shift") == "true"
	ctrl := r.URL.Query().Get("ctrl") == "true"
	alt := r.URL.Query().Get("alt") == "true"
	super := r.URL.Query().Get("super") == "true"

	if keysParam == "" {
		http.Error(w, "keyCodes parameter is required", http.StatusBadRequest)
		return
	}

	keyNames := strings.Split(keysParam, ",")
	var keyCodes []int

	// Map key names to virtual key codes
	for _, keyName := range keyNames {
		vkCode := getKeyCodeByKeyName(keyName)
		if vkCode == 0 {
			http.Error(w, "Invalid keyCodes provided", http.StatusBadRequest)
			return
		}
		keyCodes = append(keyCodes, vkCode)
	}

	simulateKeyPress(keyCodes, shift, ctrl, alt, super)

	logMessage := fmt.Sprintf("Key press simulated for keyCodes: %s", strings.Join(keyNames, " | "))
	if ctrl {
		logMessage += " | Ctrl"
	}
	if alt {
		logMessage += " | Alt"
	}
	if shift {
		logMessage += " | Shift"
	}
	if super {
		logMessage += " | Super"
	}

	log.Print(logMessage)
	fmt.Fprint(w, logMessage)

}

// getConnectedIPs returns a slice of connected and active IP addresses of the server
func getConnectedIPs() []string {
	var ips []string

	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return ips
	}

	// Iterate through interfaces to find connected and active IP addresses
	for _, iface := range interfaces {
		if (iface.Flags&net.FlagUp != 0) && (iface.Flags&net.FlagLoopback == 0) && (iface.Flags&net.FlagPointToPoint == 0) && !strings.Contains(iface.Name, "Host-Only Network") {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					if v.IP.To4() != nil {
						ips = append(ips, fmt.Sprintf("http://%s:%s", v.IP.String(), port))
					}
				}
			}
		}

	}
	ips = append(ips, fmt.Sprintf("http://%s:%s", "127.0.0.1", port))
	return ips
}

// serveServer starts the HTTP server on the specified port.
func serveServer() {

	// Register route handlers
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/press", keyPressHandler)
	http.HandleFunc("/press/multi", multiKeyHandler)

	fmt.Print(color.YellowString("API: "))
	fmt.Println(strings.Join(getConnectedIPs(), " | "))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf(color.RedString("Error: Starting server: %v"), err)
	}

}
