# 🚨Lux
Lux is a command-line interface for controlling and monitoring Govee lighting strips built in Go. Lux provides it's users with the ability to manage their lighting strips from their desktop computer from any network with an internet connection

## Installation
We provide a simple installer for each release for Windows but if that's not suitable for you, you can easily compile Lux directly from source.
You can get the latest copy of Lux from our [GitHub releases](https://github.com/BanDev/Lux/releases), or via [winget](https://docs.microsoft.com/en-us/windows/package-manager/winget/)

```shell
winget install BanDev.Lux
```

> If you choose to self compile, you will need to add the path to `lux.exe` to your `Path` environment variables in order for Lux to be detected as a command in PowerShell & CMD

If you're already using a terminal, restart it so that it can register the changes to the environment variables. Then, run the following command to begin setting up your instance of Lux

```shell
lux setup
```

## Commands
* `lux devices` Lists the devices registered on the user's account
* `lux query $id` Returns data about the state of a device
* `lux turn $id on` Turn a device on or off
* `lux turn $id $perc` Alter the brightness of a device
* `lux color $id $color` Alter the color of a device
* `lux setup` Connect Lux to your Govee account using your API Key
* `lux help` Lists a list of commands and information about the CLI

## Contributing
Looking to contribute to Lux? That's great! There are a couple of ways to help out. Translations, bug reports and pull requests are all greatly appreciated. Please refer to the [contributing guidelines](https://github.com/BanDev/Lux/blob/main/CONTRIBUTING.md) to get started

## License
```
Lux
Copyright © 2022 Jack Devey

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>
```
