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
	"runtime"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func Create() *cli.App {
	app := &cli.App{
		Name:        "keycontrol",
		Usage:       "Run a server for controlling keyboard events remotely",
		Description: appDescription,
		Copyright:   appCopyright,
		Authors:     appAuthors,
		Version:     appVersion,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number for the KeyControl server",
				EnvVars:  []string{"KEY_CONTROL_PORT"},
				Value:    "8080",
				Required: false,
			},

			&cli.StringFlag{
				Name:     "auth",
				Aliases:  []string{"a"},
				Usage:    "Authentication token required for API access",
				EnvVars:  []string{"KEY_CONTROL_AUTH_TOKEN"},
				Value:    "keycontrol",
				Required: false,
			},
		},

		Action: func(c *cli.Context) error {
			clearConsole()
			port = c.String("port")
			authToken = c.String("auth")

			fmt.Println(color.HiBlueString(appNameArt))
			fmt.Println("Welcome to KeyControl!")
			fmt.Print(color.YellowString("Version: "))
			fmt.Printf("%s | Go version: %s | OS/Arch: %s/%s\n", appVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH)
			fmt.Println(color.YellowString("Status: ") + color.GreenString("Online"))
			fmt.Println(color.YellowString("Documentation: ") + appRepoURL)

			serveServer()

			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v", "ver", "about"},
				Usage:   "Print the KeyControl version information",
				Action: func(c *cli.Context) error {
					return showVersion()
				},
			},
		},
	}

	return app
}
