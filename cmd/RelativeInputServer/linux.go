//go:build linux
// +build linux

package main

import "github.com/TKMAX777/RemoteRelativeInput/linux/host"

func main() {
	host.StartServer()
}
