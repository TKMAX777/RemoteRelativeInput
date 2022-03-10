# RemoteRelativeInput
## About This Program
This program is designed to allow relative input in an RDP (VNC) session by sending the client's input information using an SSH session. Currently, only sessions from a Windows machine to a Windows or Linux machine are supported.

![sample](https://gyazo.com/5b6e57408136ba4fcebfd2525b7dc232.gif)

## install

### Server

#### Debian / Ubuntu

```sh
sudo apt install xdotool 
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputServer@latest
```

#### Windows

- Windows requires a separate worker program to send messages to the user session.

```
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputServer@latest
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputWorker@latest
```

- Refer to the following for how to install OpenSSH server on Windows.

[Get started with OpenSSH | Microsoft Docs](https://docs.microsoft.com/ja-jp/windows-server/administration/openssh/openssh_install_firstuse)

### Client

```sh
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputClient@latest
```

## Usage

### Connect to Debian / Ubuntu

```sh
CLIENT_NAME="192.168.***.*** - Remote Desktop" RelativeInputClient | ssh 192.168.***.*** /home/<UserName>/go/bin/RelativeInputServer
```

- Pressing the F8 key toggles between relative and absolute input

### Connect to Windows

1. Start the worker program on the host machine.

```
start /d "C:\Users\<UserName>\go\bin" RelativeInputServer.exe
```

2. Starts an SSH session on the client machine.

```
CLIENT_NAME='192.168.***.*** - Remote Desktop' RelativeInputClient.exe |ssh 192.168.***.*** "C:\Users\<UserName>\go\bin\RelativeInputServer.exe"
```
