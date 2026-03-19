# lazypodman

A simple terminal UI for Podman and Podman Compose, written in Go.

![Gif](/docs/resources/demo3.gif)

## Why lazypodman

`lazypodman` is a fork from [lazydocker](https://github.com/jesseduffield/lazydocker) but for podman.

## Requirements

- Podman installed
- Podman socket enabled
- a compose provider available through `podman compose`
- Go `1.19+` for installation from source via `go install`

## Installation

Install with Go:

```sh
go install github.com/m7medVision/lazypodman@latest
```

## Quick Start

Example rootless setup:

```sh
systemctl --user enable --now podman.socket
lazypodman
```

## Usage

Run:

```sh
lazypodman
```

Optional alias:

```sh
echo "alias lzp='lazypodman'" >> ~/.zshrc
```

## Configuration

- open the config from the project panel with `o`
- edit the config directly with `e`
- see `docs/Config.md` for available options and command templates

## Contributing

Contributions are welcome. Open an issue or pull request in:

- `https://github.com/m7medVision/lazypodman`
