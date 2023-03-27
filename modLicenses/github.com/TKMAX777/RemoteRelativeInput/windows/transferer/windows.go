package transferer

import (
	"bufio"
	"fmt"
	"os"
	"os/user"

	"github.com/natefinch/npipe"
)

func StartTransferer() {
	userinfo, err := user.Current()
	if err != nil {
		panic(err)
	}

	conn, err := npipe.Dial(`\\.\pipe\RemoteRelativeInput\` + userinfo.Uid)
	if err != nil {
		panic(err)
	}

	go func() {
		var worker = bufio.NewScanner(conn)
		for worker.Scan() {
			fmt.Println(worker.Text())
		}
	}()

	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}
}
