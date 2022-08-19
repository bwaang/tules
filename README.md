# tules
Simple util scripts to help with everyday repetitive commands

## Installation

To install clone down this repository:

```bash
git clone git@github.com:bwaang/tules.git
```

Set your `$TULES_HOME` to the directory you cloned this repository to.  You can use `pwd` to figure that out from the previous step.  To make this stay, you can update this in your `~/.bash_profile` (Mac) or `~/.bashrc` (Linux)

```bash
export TULES_HOME=`path-to-tules-directory`
```

Next update your `$PATH` in your `~/.bash_profile` to include the `bin` directory of this repository, this allows you to execute scripts from anywhere.

```bash
$PATH="$PATH:$TULES_HOME/bin"
```
## GO Support

To compile and build with go 1.18 and module support, add the following to your .zshrc directory

```bash
export GOBIN="$TULES_HOME/bin"
export GOPATH="$TULES_HOME"
```
