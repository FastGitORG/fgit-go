# fgit-go

Command line to do git operation with FastGit.

## What works

- Clone  
*git clone GITHUB_URL* command had been tested in Windows 10(build 1809)

- Stdout/stderr  
Real-time display has been tested in Windows 10(build 1809)

## Preparation

Before use fgit-go, install git by yourself. Add git to env PATH is also required.

Download source code of fgit-go, build and run.

## Difference between fgit

[fgit](https://github.com/fastgitorg/fgit) by @xkeyc only provides clone operation support, but fgit-go provides push and etc.

And, fgit-go is cross-platform.

## How does fgit-go work

To clone, fgit-go just replaces url to FastGit url.

To push, fgit-go modifies .git config temporarily.
