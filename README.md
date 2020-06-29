<img width="150" height="150" align="left" style="float: left; margin: 0 10px 0 0;" alt="fastgit" src="https://avatars2.githubusercontent.com/u/62810231?s=200&v=4">

# fgit-go

[![CircleCI](https://circleci.com/gh/fastgitorg/fgit-go/tree/master.svg?style=svg)](https://circleci.com/gh/fastgitorg/fgit-go/tree/master)
[![LGBT-CN](https://img.shields.io/badge/Support-LGBTQIA-FF0000?style=flat-square)](https://git.io/JfJiO)

Command line to do git operation with FastGit.

## What works

- Stdout/stderr  
Real-time display has been tested in Windows 10(build 1809)

- Push  
*`fgit push`* command had been tested in Windows 10(build 1809)

- Clone  
*`fgit clone GITHUB_URL`* command had been tested in Windows 10(build 1809)

## Preparation

Before use fgit-go, install `git` by yourself. Add git to env PATH is also **required**.

Download source code of fgit-go, build and run.

## Extra Syntax

### 1. debug

**SYNTAX:**

```bash
fgit debug [URL<string>] [help]
```

**FUNCTION:**

This command line is for debug. Will provide remote addr, local addr, and connection info.

**EXAMPLE:**

```bash
>fastgit debug
FastGit Debug Tool
==================
Remote Address: https://hub.fastgit.org
IP Address: [x.x.x.x]
Local Address: [x.x.x.x]
Test connection...Success
```

### 2. get

**SYNTAX:**

```bash
fgit get [URL<string>] [Path<string>] [--help]
```

**ALIAS:**

```bash
fgit dl
fgit download
```

**FUNCTION:**

This command line is for downloading. Will auto convert github download link to fastgit.

**EXAMPLE:**

```bash
>fgit get https://github.com/fastgitorg/fgit-go/archive/master.zip
File with the same name exists. New file will cover the old file.
Do you want to continue? [Y/n]y
Redirect url -> https://download.fastgit.org/fastgitorg/fgit-go/archive/master.zip
Downloading...
Finished.
```

## Difference between fgit

[fgit](https://github.com/fastgitorg/fgit) by @xkeyc only provides clone operation support, but fgit-go provides push and etc.

And, fgit-go is cross-platform.

## How does fgit-go work

To clone, fgit-go just replaces url to FastGit url.

To push, fgit-go modifies .git config temporarily. Like

```config
[core]
        repositoryformatversion = 0
        filemode = false
        bare = false
        logallrefupdates = true
        symlinks = false
        ignorecase = true
[remote "origin"]
        url = https://github.com/A/B
        fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
        remote = origin
        merge = refs/heads/master
```

To

```config
[core]
        repositoryformatversion = 0
        filemode = false
        bare = false
        logallrefupdates = true
        symlinks = false
        ignorecase = true
[remote "origin"]
        url = https://hub.fastgit.org/A/B
        fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
        remote = origin
        merge = refs/heads/master
```

## TODO

- [ ] `--verbose` flag
- [ ] `--node` flag

## Open-Source License - [GNU General Public License v3.0](LICENSE)

```license
                    GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.
```

-----
Proudly build with ❤️ & golang
