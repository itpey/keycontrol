<div align="center">
  <img src="https://raw.githubusercontent.com/egonelbre/gophers/master/vector/computer/gamer.svg" width="180" height="180">
</div>

<h1 align="center">KeyControl</h1>

<p align="center">
  KeyControl is a Go-based HTTP server that allows users to simulate keyboard key presses remotely via HTTP requests. This server is useful for automation, testing, and remote control scenarios.
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/itpey/keycontrol">
    <img src="https://pkg.go.dev/badge/github.com/itpey/keycontrol.svg" alt="Go Reference">
  </a>
  <a href="https://github.com/itpey/keycontrol/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/itpey/keycontrol" alt="license">
  </a>
</p>

## Table of Contents

- [Overview](#Overview)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
- [API Endpoints](api-endpoints)
- [Configuration](#configuration)
- [Key Code Mapping](#key-code-mapping)
- [Feedback and Contributions](#feedback-and-contributions)
- [License](#license)

## Overview

This project provides a simple HTTP server that allows you to simulate key presses on the server machine by making HTTP requests with specific parameters. It uses the [micmonay/keybd_event](https://github.com/micmonay/keybd_event) package to handle the keyboard events.

## Features

- Simulate keyboard key presses via HTTP endpoints
- Support for various key combinations (e.g., Ctrl+Alt+Delete)
- Easy-to-use API for integrating with other applications
- Customizable port and authentication token for security

## Getting Started

### Installation

#### Using Binary Releases

- Go to the [Releases](https://github.com/itpey/keycontrol/releases) page of KeyControl repository.
- Download the latest release for your operating system.

#### Using go install

If you have Go installed on your system, you can install the latest version of KeyControl directly using the go install command:

```bash
go install github.com/itpey/keycontrol@latest
```

Ensure that your `GOBIN` directory is in your `PATH` for the installed binary to be accessible globally.

### Usage

Once installed, you can use KeyControl from the command line:

```bash
$ keycontrol -h
```

## API Endpoints

### `/health` (GET)

Check the health status of the server.

#### Example:

```bash
curl -X GET "http://localhost:8080/health"
```

### `/press` (POST)

Simulate key press.

#### Parameters:

- `keycodes (required)`: Comma-separated list of virtual key codes to simulate.
- `shift (optional)`: Set to true to press SHIFT key.
- `ctrl (optional)`: Set to true to press CTRL key.
- `alt (optional)`: Set to true to press ALT key.
- `super (optional)`: Set to true to press Windows/Super key.

#### Example:

```bash
curl -X POST "http://localhost:8080/press?keycodes=A,B,C&ctrl=true" \
  -H "Authorization: keycontrol"
```

## Configuration

### Command-line Flags

You can configure the KeyControl server using the following command-line flags:

- `-- port, -p`: Specifies the server port (default: 8080).
- `--auth, -a`: Specifies the authentication token (default: keycontrol).

#### Example:

```bash
$ keycontrol -p 9000 -a mytoken
```

### Environment Variables

Customize the server port and authentication token using the following environment variables:

- `KEY_CONTROL_PORT`: Sets the server port (default: 8080).
- `KEY_CONTROL_AUTH_TOKEN`: Sets the authentication token (default: keycontrol).

#### Example:

```bash
$ export KEY_CONTROL_PORT=9000
$ export KEY_CONTROL_AUTH_TOKEN=mytoken
$ keycontrol
```

## Key Code Mapping

Refer to the source code for a complete list of key codes.

## Feedback and Contributions

If you encounter any issues or have suggestions for improvement, please [open an issue](https://github.com/itpey/keycontrol/issues) on GitHub.

We welcome contributions! Fork the repository, make your changes, and submit a pull request.

## License

KeyControl is open-source software released under the Apache License, Version 2.0. You can find a copy of the license in the [LICENSE](https://github.com/itpey/keycontrol/blob/main/LICENSE) file.
