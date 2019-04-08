<h3 align="center">
	<img
		width="40"
		alt="NV3"
		src="https://s3.eu-central-1.amazonaws.com/clogger/media/nv3-bw.png">
</h3>

<h1 align="center">
	GO STRUCTER
</h1>

<h4 align="center">Generate structs from JSON definitions</h4>

<p align="center">
	<strong>
		<a href="https://github.com/strackovski/structer/issues">Issues</a>
		•
		<a href="https://github.com/strackovski/structer/wiki">Wiki</a>
		•
		<a href="examples">Examples</a>
	</strong>
</p>


<p align="center">
	<img src="https://s3.eu-central-1.amazonaws.com/terrarium-pro/terrarium-logo-tmp.png" width="230">
</p>

# Overview
> [![Build](https://img.shields.io/badge/build-passing-brightgreen.svg)](#, "Build passing.")
[![Coverage](https://img.shields.io/badge/coverage-78%25-yellowgreen.svg)](#, "Coverage 78%.")
[![License](https://img.shields.io/badge/license-Apache%202-blue.svg)](#, "View this license type.")
[![Water](https://img.shields.io/badge/water-added-brightgreen.svg)](#, "Just add water.")
>
> ### A JSON-to-struct generator for Go
>
>Because sometimes you just have to generate a struct from json.

# Installation

>### Requirements
>
>- Go 1.11

## Quick Start

Install structer using the installation script by executing the following command (execute in directory with write permissions):

```bash
curl https://s3.eu-central-1.amazonaws.com/structer/dist/structer_installer.sh | bash
```

The installer script is available for download [here](https://s3.eu-central-1.amazonaws.com/structer/dist/structer_installer.sh). 

When installed, run the `structer` command from terminal:

```bash
structer
```

You should see the following output:

```bash
Structer is a tool for creating struct types from definitions stored as JSON.

Usage:
  structer [command]

Available Commands:
  help        Help about any command
  make        Makes a struct from JSON data.

Flags:
      --config string   config file (default is $HOME/.structer.yaml)
  -h, --help            help for structer
  -t, --toggle          Help message for toggle

Use "structer [command] --help" for more information about a command.
```

To get familiar with the main command, `make`, use the help option 

```bash
structer help make
```

```bash
Makes a struct from JSON data.

Usage:
  structer make [flags]

Flags:
      --format string   the format of the input (default "json")
  -h, --help            help for make
      --in string       path to input file or directory
      --out string      path to output directory (default "generated")

Global Flags:
      --config string   config file (default is $HOME/.structer.yaml)
```


Use `structer make` providing the path to json file in the `--in` parameter:
 
```bash
structer make --in <PATH_TO_JSON_FILE>
```

> :bulb: **TIP**
> 
> For more examples visit the [Examples](examples) directory.

## Download Binaries

> :bulb: **TIP**
> 
> Download the latest release of structer for your platform.

- The fastest way to install is by downloading and running the installer script as shown in [Quick Start](#quick-start)
- You can download the `src` archive if you'd like to build this tool form source. See [Building From Source]
(#building-from-source) for details.


| File name                    | Kind    | OS      | Arch   | Size | Checksum                                                         |
|------------------------------|---------|---------|--------|------|------------------------------------------------------------------|
| [structer-installer.sh](https://s3.eu-central-1.amazonaws.com/structer/dist/structer_installer.sh)    | Installer  | -       | -      | 3.4KB | [View](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_darwin_386.checksum) |
| [structer_0.1.0_darwin_386](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_darwin_386) | Archive | OSX     | x86 | 10.9MB | [View](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_darwin_386.checksum) |
| [structer_0.1.0_darwin_amd64](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_darwin_amd64) | Archive | OSX     | x86-64 | 12.3MB| [View](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_darwin_amd64.checksum) |
| [structer_0.1.0_linux_386](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_linux_386)    | Archive | Linux   | x86    | 10.9MB | [View](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_linux_386.checksum) 
| [structer_0.1.0_linux_amd64](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_linux_amd64)  | Archive | Linux   | x86-64 | 12.9MB | [View](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_linux_amd64.checksum) |
| [structer_0.1.0_freebsd_386](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_freebsd_386)   | Archive | freebsd | x86    | 10.9MB | [View](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_freebsd_386.checksum) |
| [structer_0.1.0_freebsd_amd64](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_freebsd_amd64)  | Archive | freebsd | x86-64    | 12.3MB | [View](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_freebsd_amd64.checksum) |
| [structer_0.1.0_windows_386](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_windows_386)   | Archive | Windows | x86    | 10.9MB | [View](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_windows_386.checksum) |
| [structer_0.1.0_windows_amd64](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_windows_amd64)  | Archive | Windows | x86-64 | 12.4MB | [View](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/structer_0.1.0_windows_amd64.checksum) |

Builds for other platforms available [here](https://s3.eu-central-1.amazonaws.com/structer/0.1.0/).

> :truck: **Release 0.1.0**
> 
> See [release log](CHANGELOG.md) for more information about the current and previous structer releases.

## Building From Source

Either download the source package or clone this repository, cd into the directory where you have the source files, and execute:

```bash
make && make install
```

This will build and install the structer executable to your path.

That's it. Now you can run:

```bash
structer
```

Makefile offers the following options:

```bash
# Start the build process, cross-building all supported platforms:
make

# Start the build process for the current platform:
make current

# Start the build process for a specific platform:
make platform darwin_amd64

# The publish option is reserved for project maintainers. 
# It will publish the distributable files to the build server 
# after a successful build.
make publish

# Make install assumes there's a build file waiting in the build 
# directory, and tries to move it to system path:
make install

# Downloads and installs the prebuilt binary for the current platform:
make web_install

```


# Usage

Make sure you try the [Quick Start](#quick-start) guide first. See the [examples](examples/README.md) directory for more examples. 
Below is a very simple one:

```bash
structer
structer help
structer make --in <PATH_TO_JSON_FILE>
```

# Running Tests

> :bulb: **TIP**
> 
> TBD

```bash
structer test run
```

# Contributing

Contributions are welcome. Check [Issues & Roadmap](#issues-and-roadmap) below to see how you can help with priority features, bugs, 
tests or 
docs. 

Contribution guidelines are described in the [Contributing document](#).

# Issues and Roadmap

Issue tracker top:

> :zap: **[STR-2](https://github.com/strackovski/structer/issues/1)** Enable (extendable) multiple input formats

> :bulb: **[STRU-3](https://github.com/strackovski/structer/issues/3)** TBD


# Contact

Feel free to get in touch if you have a question or request in connection to this program or source repository.

**Open, public communication is the prefered choice:**

- Open an [issue](https://github.com/strackovski/structer/issues), or ask a [question on wiki](https://github.com/strackovski/structer/wiki) 
- There's a mailing list and a Google Group at [structer.groups.nv3.eu](https://structer.groups.nv3.eu)

Private inquiries over email. Find emails in:
 
- contributors in [CONTRIBUTORS.md](CONTRIBUTORS.md) for emails

# License

Licensed under the Apache License, Version 2.0. See the LICENSE file distributed with this source code for the full text license agreement.

<h1 align="right">
	<img
		width="20"
		alt="Logo"
		src="https://s3.eu-central-1.amazonaws.com/clogger/media/nv3-bw.png">
</h1>