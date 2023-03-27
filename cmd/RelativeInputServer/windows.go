//go:build windows
// +build windows

package main

import "github.com/TKMAX777/RemoteRelativeInput/windows/host"

func main() {
	host.StartServer()
}
