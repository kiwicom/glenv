# GlEnv

Jump into your favourite repository folder and GitLab env. variables are 
automatically loaded into your shell. 

This tool in combination with [direnv](https://direnv.net/) will export you 
project's env. variables from GitLab.

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

The `glenv` make sense in comboe with `direnv`. Please install also `direnv`:

```
$ brew install direnv
```

The direnv need some post-installation setup. For ZShell users, add the 
following line at the end of the `~/.zshrc` file and restart the shell:

```
$ eval "$(direnv hook zsh)"
```

If you're using different shell, follow the instructions in [direnv documentation](https://direnv.net/docs/hook.html).

The `glenv` is available in Homebrew as well.:

```
$ brew install kiwi/tools/glen
```

Make sure you have `GITLAB_TOKEN` env. variable present in your environment with
your [personal access token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.htm).

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

