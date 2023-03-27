//go:build linux
// +build linux

package main

import "github.com/TKMAX777/RemoteRelativeInput/linux/client"

func main() {
	client.StartClient()
}
