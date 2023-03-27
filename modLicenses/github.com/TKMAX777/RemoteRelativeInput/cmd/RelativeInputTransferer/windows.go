//go:build windows
// +build windows

package main

import "github.com/TKMAX777/RemoteRelativeInput/windows/transferer"

func main() {
	transferer.StartTransferer()
}
