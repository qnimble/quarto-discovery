# Arduino pluggable discovery for serial ports

The `quarto-discovery` tool is a command line program that interacts via stdio. It accepts commands as plain ASCII strings terminated with LF `\n` and sends response as JSON. It is forked and based on ['serial-discovery'](https://github.com/arduino/serial-discovery).

## How to build

Install a recent go environment (>=13.0) and run `go build`. The executable `quarto-discovery` will be produced in your working directory.

## Usage

After startup, the tool waits for commands. The available commands are: `HELLO`, `START`, `STOP`, `QUIT`, `LIST` and `START_SYNC`.

#### HELLO command

The `HELLO` command is used to establish the pluggable discovery protocol between client and discovery.
The format of the command is:

`HELLO <PROTOCOL_VERSION> "<USER_AGENT>"`

for example:

`HELLO 1 "Arduino IDE"`

or:

`HELLO 1 "arduino-cli"`

in this case the protocol version requested by the client is `1` (at the moment of writing there were no other revisions of the protocol).
The response to the command is:

```json
{
  "eventType": "hello",
  "protocolVersion": 1,
  "message": "OK"
}
```

`protocolVersion` is the protocol version that the discovery is going to use in the remainder of the communication.

#### START command

The `START` starts the internal subroutines of the discovery that looks for ports. This command must be called before `LIST` or `START_SYNC`. The response to the start command is:

```json
{
  "eventType": "start",
  "message": "OK"
}
```

#### STOP command

The `STOP` command stops the discovery internal subroutines and free some resources. This command should be called if the client wants to pause the discovery for a while. The response to the stop command is:

```json
{
  "eventType": "stop",
  "message": "OK"
}
```

#### QUIT command

The `QUIT` command terminates the discovery. The response to quit is:

```json
{
  "eventType": "quit",
  "message": "OK"
}
```

after this output the tool quits.

#### LIST command

The `LIST` command returns a list of the currently available serial ports. The format of the response is the following:

```json
{
  "eventType": "list",
  "ports": [
    {
      "address": "COM16",
      "label": "COM16",
      "properties": {
        "pid": "0x0941",
        "vid": "0x1781",
        "serialNumber": "12345678"
      },
      "protocol": "serial",
      "protocolLabel": "Quarto"
    }
  ]
}
```

The `ports` field contains a list of the available Quarto devices.

The list command is a one-shot command, if you need continuous monitoring of ports you should use `START_SYNC` command.

#### START_SYNC command

The `START_SYNC` command puts the tool in "events" mode: the discovery will send `add` and `remove` events each time a new port is detected or removed respectively.
The immediate response to the command is:

```json
{
  "eventType": "start_sync",
  "message": "OK"
}
```

after that the discovery enters in "events" mode.

The `add` events looks like the following:

```json
{
  "eventType": "add",
  "port": {
    "address": "COM16",
    "label": "COM16",
    "properties": {
      "pid": "0x0941",
      "vid": "0x1781",
      "serialNumber": "12345678"
    },
    "protocol": "serial",
    "protocolLabel": "Quarto"
  }
}
```

it basically gather the same information as the `list` event but for a single port. After calling `START_SYNC` a bunch of `add` events may be generated in sequence to report all the ports available at the moment of the start.

The `remove` event looks like this:

```json
{
  "eventType": "remove",
  "port": {
    "address": "COM16",
    "protocol": "serial"
  }
}
```

in this case only the `address` and `protocol` fields are reported.

### Example of usage

A possible transcript of the discovery usage:

```
$ ./quarto-discovery
START
{
  "eventType": "start",
  "message": "OK"
}
LIST
{
  "eventType": "list",
  "port":
  {
    "address": "COM16",
    "label": "COM16",
    "properties": {
       "pid": "0x0941",
       "vid": "0x1781",
       "serialNumber": "12345678"
    },
    "protocol": "serial",
    "protocolLabel": "Quarto"
  }
}
START_SYNC
{
  "eventType": "start_sync",
  "message": "OK"
}
{                                  <--- this event has been immediately sent
  "eventType": "add",
  "port":
  {
    "address": "COM16",
    "label": "COM16",
    "properties": {
       "pid": "0x0941",
       "vid": "0x1781",
       "serialNumber": "12345678"
    },
    "protocol": "serial",
    "protocolLabel": "Quarto"
  }
}
{                                  <--- the board has been disconnected here
  "eventType": "remove",
  "port": {
    "address": "COM16",
    "protocol": "serial"
  }
}
{                                  <--- the board has been connected again
  "eventType": "add",
  "port":
  {
    "address": "COM16",
    "label": "COM16",
    "properties": {
       "pid": "0x0941",
       "vid": "0x1781",
       "serialNumber": "12345678"
    },
    "protocol": "serial",
    "protocolLabel": "Quarto"
  }
}
QUIT
{
  "eventType": "quit",
  "message": "OK"
}
$
```

## License

Copyright (c) 2018 ARDUINO SA (www.arduino.cc)
Copyright (c) 2021 qNimble Inc (qnimble.com)

The software is released under the GNU General Public License, which covers the main body
of the quarto-discovery code. The terms of this license can be found at:
https://www.gnu.org/licenses/gpl-3.0.en.html

See [LICENSE.txt](https://github.com/qnimble/quarto-discovery/blob/main/LICENSE.txt) for details.
