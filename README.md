[![version](https://shields.io/github/v/release/kiwicom/glenv)

# glenv

Jump into your favourite repository folder and GitLab env. variables are 
automatically loaded into your shell. This tool in combination with 
[direnv](https://direnv.net/) will export you project's env. variables 
from GitLab.

## How it works

First, I have to say the glenv requires [direnv](https://direnv.net/) installed.

Let's imagine we have clonned the `hello/project` repository into some folder.

```
$ clone git git@gitlab.com:glenv-demo/hello/demo.git
```

Let's jump into cloned `demo` folder.

```
$ cd ~/demo
direnv: loading ~/demo/.envrc
direnv: export +CI_PROJECT_ID +DEMO_ENV ...
```

The `direnv` automatically inject all env. variables from GitLab into my shell.

When I jump-out from this folder:

```
$ cd ..
direnv: unloading
```

The `direnv` automatically unload all GitLab env. variables.


## How to install

You can use Homebrew

```
brew install kiwicom/tap/glenv
```

Or you can download archive with binary here in [GitHub releases](https://github.com/kiwicom/glenv/releases/).
Installation is simple anyway. Download the desired archive for your system, 
unpack it and place it in some folder, which is in your PATH.

No need to configure. Only `GITLAB_TOKEN` variable need to be present in your 
environment.

## How to enable it in repository

If you don't have `.envrc` in your repository directory, call the command:

```
$ glenv init
```

The command create `.envrc` file for you. Now you can allow the directory

```
$ direnv allow .
```

enjoy :-)

