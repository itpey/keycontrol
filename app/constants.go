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

import "github.com/urfave/cli/v2"

const (
	appVersion     = "0.1.0"
	appDescription = `KeyControl is a command-line application that starts an HTTP server to simulate keyboard events remotely. It allows you to send key press commands via HTTP requests, making it useful for automation, remote control, or integration with other systems.`
	appCopyright   = "Apache-2.0 license\nFor more information, visit the GitHub repository: https://github.com/itpey/keycontrol"
	appRepoURL     = "https://github.com/itpey/keycontrol"
)

const (
	appNameArt = `░█░█░█▀▀░█░█░█▀▀░█▀█░█▀█░▀█▀░█▀▄░█▀█░█░░
░█▀▄░█▀▀░░█░░█░░░█░█░█░█░░█░░█▀▄░█░█░█░░
░▀░▀░▀▀▀░░▀░░▀▀▀░▀▀▀░▀░▀░░▀░░▀░▀░▀▀▀░▀▀▀`
)

var (
	appAuthors = []*cli.Author{{Name: "itpey", Email: "itpey@github.com"}}
)
