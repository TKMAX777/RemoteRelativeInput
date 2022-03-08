# RemoteRelativeInput
## Abstruct
This program is designed to allow relative input in an RDP (VNC) session by sending the client's input information using an SSH session. Currently only supported when connecting from Windows to a Linux machine.

![sample](https://gyazo.com/5b6e57408136ba4fcebfd2525b7dc232.gif)

## install

### Server

```sh
sudo apt install xdotool 
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputServer@latest
```

### Client

```sh
go install github.com/TKMAX777/RemoteRelativeInput/cmd/RelativeInputClient@latest
CLIENT_NAME="192.168.***.*** - Remote Desktop" RelativeInputClient | ssh 192.168.***.*** /home/.../go/bin/RelativeInputServer
```

## Usage

```sh
CLIENT_NAME="192.168.***.*** - Remote Desktop" RelativeInputClient | ssh 192.168.***.*** /home/<UserName>/go/bin/RelativeInputServer
```

- Pressing the F8 key toggles between relative and absolute input
