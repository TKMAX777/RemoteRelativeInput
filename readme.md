# RemoteRelativeInput
## About This Program
This program is designed to allow relative input in an RDP (VNC) session by wrapping an existing remote desktop client window with another window and sending the client's input information using an SSH session. Currently, only sessions from a Windows machine to a Windows or Linux machine are supported.

![sample](https://gyazo.com/5b6e57408136ba4fcebfd2525b7dc232.gif)

## install

### Server

#### Debian / Ubuntu

```sh
sudo apt install xdotool golang-go
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputServer@latest
```

#### Windows

- The [Go](https://go.dev/doc/install) and [OpenSSH Server](https://docs.microsoft.com/en-us/windows-server/administration/openssh/openssh_install_firstuse) must be installed before installation.
- Windows requires a separate worker program to send messages to the user session.

```
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputServer@latest
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputWorker@latest
```

### Client
- The [Go](https://go.dev/doc/install) and [OpenSSH Client](https://docs.microsoft.com/en-us/windows-server/administration/openssh/openssh_install_firstuse) must be installed before installation.
- Currently, it is reported that it does not work properly in PowerShell.<br>

```sh
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputClient@latest
```

## Usage

### Connect to Debian / Ubuntu

```sh
set CLIENT_NAME=<hostname> - Remote Desktop Connection
RelativeInputClient.exe | ssh <hostname> /home/<UserName>/go/bin/RelativeInputServer
```

- Pressing the F8 key toggles between relative and absolute input
- CLIENT_NAME should be the name of an existing remote desktop client window.

### Connect to Windows

1. Start the worker program on the host machine.

```
start /d "C:\Users\<HostFolderName>\go\bin" RelativeInputWorker.exe
```

2. Starts an SSH session from the client machine.

```
set CLIENT_NAME=<HostAddress> - Remote Desktop Connection"
RelativeInputClient.exe | ssh <HostUsername>@<HostAddress> "C:\Users\<HostFolderName>\go\bin\RelativeInputServer.exe"
```

- CLIENT_NAME should be the name of an existing remote desktop client window.

3. Enter host user password

4. Press Yes in the message box displayed on the host machine.

5. Press OK in the message box displayed on the client machine.

6. Enjoy!

- Pressing the F8 key toggles between relative and absolute input
