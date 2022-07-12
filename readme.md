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

### Connect to Windows

1. Open Remote Desktop Connection and connect to your server like usual and have it in Maximize Windowed (**NOT FULL SCREEN**)

2. Open cmd and start the worker program on the host machine.

```
start /d "C:\Users\<HostFolderName>\go\bin" RelativeInputWorker.exe
```

☆ replace <HostFolderName> with the name of the folder of your account located in C:\Users

3. Starts an SSH session from the client machine on cmd.

```
set CLIENT_NAME=<HostAddress> - Remote Desktop Connection"
RelativeInputClient.exe | ssh <HostUsername>@<HostAddress> "C:\Users\<HostFolderName>\go\bin\RelativeInputServer.exe"
```

replace <HostFolderName> with the name of the folder of your account located in C:\Users
replace <HostUsername> with the Host username
replace <HostAddress> with the IP of your host

4. Ignore the message box and click on cmd tab, and enter host user password

5. Press Yes in the message box displayed on the host machine.

6. Press OK in the message box displayed on the client machine.

7. Enjoy!

☆ If your mouse disappeared hit F8 and go back to your Remote Desktop Connection Then choose RDP Input Wrapper and hit F8

### Connect to Debian / Ubuntu

```sh
set CLIENT_NAME=<hostname> - Remote Desktop Connection
RelativeInputClient.exe | ssh <hostname> /home/<UserName>/go/bin/RelativeInputServer
```

☆ If your mouse disappeared hit F8 and go back to your Remote Desktop Connection Then choose RDP Input Wrapper and hit F8
☆ replace <HostFolderName> with the name of the folder of your account located in C:\Users
