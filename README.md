<!-- markdownlint-disable MD033 -->
<h1 align="center">pair</h1>

<p align="center">
  <em>Pair commits with ticket ids and co-authors made simple.</em>
</p>

![pair-demo](https://github.com/user-attachments/assets/27dec47a-5c02-40d8-88f6-0d82c603ae53)

## Contents

- [Why](#why)
- [Install](#install)
  - [Download from GitHub](#download-from-github)
  - [Build it locally](#build-it-locally)
- [Commands](#commands)
- [Configuration](#configuration)

## Why?

I'm constantly forgetting what ticket I'm working on so need to find it at commit time,
then manually add the co-authors each commit.

Now I don't need to worry about that, I can set the values with `pair` and have
them automatically added to my commits each time.

## Install

### Download from GitHub

Find the latest version for your system on the
[GitHub releases page](https://github.com/tx3stn/pair/releases).

### Build it locally

If you have go installed, you can clone this repo and run:

```bash
make install
```

This will build the binary and then copy it to `/usr/local/bin/pair` so it will be
available on your path. Nothing more to it.

## Commands

Run `pair --help` for a full, up to date list of available commands.

### `on`

Set the ticket id you are working on.

> [!tip]
> Set `ticketPrefix` in your config file if your tickets always have the same prefix,
> and you can just type the part after the prefix.

You can pass the ticket directly to the command, or just run `pair on` for an interactive
prompt to enter the ticket id.

![pair-on-demo](https://github.com/user-attachments/assets/413e073c-bd6c-4a5a-8362-84a3635f5827)

### `with`

Select co-authors from the list defined in your config file.

![pair-with-demo](https://github.com/user-attachments/assets/8ad2b3ad-d0b7-4c23-a43b-5b9b16b2649b)

### `commit`

Commit your staged changes using the values you have already set.
If you have not set a ticket id or selected co-authors you will be prompted to
set them before running the commit.

> [!tip]
> Want to sign your commits?
> Set `-s` in the `commitArgs` field in your config file.

![pair-commit-demo](https://github.com/user-attachments/assets/5bb45be6-a3bd-4fbd-bd2f-db7f23ebb2bd))

### `done`

When you're all done with the ticket or work you were doing, you can run `pair done`
to clear the session data temp files.

![pair-done-demo](https://github.com/user-attachments/assets/1664d5e6-1493-4393-85c5-22e073f07fe2)

## Configuration

By default `pair` checks for a config file called `pair.json` in your `$XDG_CONFIG_DIR`, 
or `$HOME/.config` in a directory e.g: `$HOME/.config/pair.json`.

### Schema

You can see an example of the config file format in the
[pair.json](./.schema/pair.json) in the schema directory.

> [!TIP]
> Make sure you add the `$schema` keyword to the top of your config file to
> for in editor validation and descriptions of what fields are used for.
