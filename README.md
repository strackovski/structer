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
	<img src="https://s3.eu-central-1.amazonaws.com/terrarium-pro/terrarium-logo-tmp.png" width="220">
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

An installer script is available at our .... Once ready, run

```bash
./structer-install.sh
```
 
When installed, run structer's `make` command from terminal, providing the path to json file in the `--file` parameter:
 
```bash
structer make --in <PATH_TO_JSON_FILE>
```

## Download Binaries

> :bulb: **TIP**
> 
> Download the latest release of structer for your platform.

- The fastest way to install is by downloading and running the installer script as shown in [Quick Start](#quick-start)
- You can download the `src` archive if you'd like to build this tool form source. See [Building From Source]
(#building-from-source) for details.


| File name                    | Kind    | OS      | Arch   | Size | Checksum                                                         |
|------------------------------|---------|---------|--------|------|------------------------------------------------------------------|
| [structer-installer.sh]()    | Installer  | -       | -      | 1MB | [View]() |
| [structer_0.1.0_src.tar.gz]()    | Source  | -       | -      | 6MB | [View]() |
| [structer_0.1.0_darwin_amd64.tar.gz]() | Archive | OSX     | x86-64 | 4.7MB| [View]() |
| [structer_0.1.0_linux_386.tar.gz]()    | Archive | Linux   | x86    | 4.7MB | [View]() 
| [structer_0.1.0_linux_amd64.tar.gz]()  | Archive | Linux   | x86-64 | 4.7MB | [View]() |
| [structer_0.1.0_windows_386.zip]()     | Archive | Windows | x86    | 4.7MB | [View]() |
| [structer_0.1.0_windows_amd64.zip]()   | Archive | Windows | x86-64 | 4.7MB | [View]() |


> :truck: **Release 0.1.0**
> 
> See [release log](CHANGELOG.md) for more information about the current and previous structer releases.

## Building From Source

Either download the source package or clone this repository, cd into the directory where you have the source files, and execute:

```bash
make
```

That's it. Now you can run:

```bash
structer
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

Issue tracker top 5 by priority:

> :exclamation: **[STRU-329](#)**  Some bug request makes sense

> :zap: **[STRU-330](#)** For more information about structer release

> :fire: **[STRU-293](#)** Some bug report confirmed

> :question: **[STRU-230](#)** Some feature request makes sense
 
> :grey_question: **[STRU-101](#)** Some feature request makes sense

> :bulb: **[STRU-109](#)** Some feature request makes sense


# Contact

Feel free to get in touch if you have a question or request in connection to this program or source repository.

**Open, public communication is the prefered choice:**

- Open an [issue](issues), or ask a [question on wiki](wiki) 
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