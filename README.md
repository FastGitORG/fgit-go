# fgit-go

[![CircleCI](https://circleci.com/gh/fastgitorg/fgit-go/tree/master.svg?style=svg)](https://circleci.com/gh/fastgitorg/fgit-go/tree/master)

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
