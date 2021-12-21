## Terminal

We use the terminal to interact with different commands and programs like
[git][git].

Each operating system comes with a different type of terminal and some commands
may work a little bit differently on different operating systems.

If you use macOS then you can open [Terminal.app][terminal.app].
On Microsoft Windows you can use [PowerShell][powershell].
For this course we will use the [terminal][integrated_terminal] that is built
into VSCode and a [Codespace][codespace].

{% include youtube.html youtube_id="x73WTEltyHU" %}

### Getting Started

For this lesson you will need to [open your Codespace][codespace] and then
launch VSCode and [open a Terminal][integrated_terminal].

The terminal is a portal that allows you interact with your computer.
It provides a command line interface that looks like this:

```bash
/workspaces/SparkleHub
$
```

Start by typing command, like `ls`, followed by pressing the Enter key.
The terminal will execute that command and print the output from the command.

When you start up a new Codespace the default directory that the
terminal starts in is the `/workspaces/SparkleHub` directory. This directory
contains all the code for the CodeChica++ final project.

### `man`

If you forget what a specific command does you can find the answer in the `man`
page. `man` is short for instruction *manual*. To look up a command in the
manual type `man` then the name of the command.

For example:

```bash
$ man man
MAN(1)                             Manual pager utils                             MAN(1)

NAME
       man - an interface to the system reference manuals

SYNOPSIS
       man [man options] [[section] page ...] ...
       man -k [apropos options] regexp ...
       man -K [man options] [section] term ...
       man -f [whatis options] page ...
       man -l [man options] file ...
       man -w|-W [man options] page ...

DESCRIPTION
       man  is  the  system's manual pager.  Each page argument given to man is normally
       the name of a program, utility or function.  The manual page associated with each
       of these arguments is then found and displayed.  A section, if provided, will di‐
       rect man to look only in that section of the manual.  The default  action  is  to
       search  in  all  of the available sections following a pre-defined order (see DE‐
       FAULTS), and to show only the first page found, even if page  exists  in  several
       sections.
```

### `ls`

The `ls` command will *list* all the files in a directory.

```bash
$ man ls
LS(1)                                 User Commands                                LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List  information  about  the FILEs (the current directory by default).  Sort en‐
       tries alphabetically if none of -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.
```
- https://man7.org/linux/man-pages/man1/ls.1.html

#### Examples

List files in the current directory:

```bash
/workspaces/SparkleHub
$ ls
Gemfile       app              config.ru     package.json       tmp
Gemfile.lock  babel.config.js  db            postcss.config.js  vendor
README.md     bin              doc           script             yarn.lock
Rakefile      config           node_modules  test
```

List the files in the `app` directory:

```bash
/workspaces/SparkleHub
$ ls app/
controllers/  javascript/  models/  views/
```

List the files using a long list format.

```bash
/workspaces/SparkleHub
$ ls -l script/
total 12K
-rwxr-xr-x 1 username 185 Jul 26 20:46 server*
-rwxr-xr-x 1 username 320 Jul 26 20:46 setup*
-rwxr-xr-x 1 username 183 Jul 26 20:46 test*
```

List all files (including hidden ones) using a long list format.

```bash
/workspaces/SparkleHub
$ ls -al
total 404K
drwxr-xr-x 14 username 4.0K Jul 26 20:46 ./
drwxr-xr-x  3 username 4.0K Jul 26 20:46 ../
drwxr-xr-x  6 username 4.0K Jul 26 20:46 app/
-rw-r--r--  1 username 1.7K Jul 26 20:46 babel.config.js
drwxr-xr-x  2 username 4.0K Jul 26 20:46 bin/
-rw-r--r--  1 username    9 Jul 26 20:46 .browserslistrc
drwxr-xr-x  4 username 4.0K Jul 26 20:46 config/
-rw-r--r--  1 username  160 Jul 26 20:46 config.ru
drwxr-xr-x  3 username 4.0K Jul 26 20:46 db/
drwxr-xr-x  2 username 4.0K Jul 26 20:46 .devcontainer/
drwxr-xr-x  2 username 4.0K Jul 26 20:46 doc/
-rw-r--r--  1 username  200 Jul 26 20:46 Gemfile
-rw-r--r--  1 username 4.1K Jul 26 20:46 Gemfile.lock
drwxr-xr-x  8 username 4.0K Jul 26 20:46 .git/
drwxr-xr-x  4 username 4.0K Jul 26 20:46 .github/
-rw-r--r--  1 username  776 Jul 26 20:46 .gitignore
-rw-r--r--  1 username  302 Jul 26 20:46 package.json
-rw-r--r--  1 username  224 Jul 26 20:46 postcss.config.js
-rw-r--r--  1 username   68 Jul 26 20:46 Rakefile
-rw-r--r--  1 username  413 Jul 26 20:46 README.md
-rw-r--r--  1 username   11 Jul 26 20:46 .ruby-version
drwxr-xr-x  2 username 4.0K Jul 26 20:46 script/
drwxr-xr-x  6 username 4.0K Jul 26 20:46 test/
drwxr-xr-x  3 username 4.0K Jul 26 20:46 tmp/
drwxr-xr-x  3 username 4.0K Jul 26 20:46 vendor/
-rw-r--r--  1 username 300K Jul 26 20:46 yarn.lock
```

### `pwd`

The `pwd` command can be used to *print* the *working directory*. This is very
helpful when you can't remember what directory you are in.

```bash
$ man pwd
PWD(1)                                User Commands                               PWD(1)

NAME
       pwd - print name of current/working directory

SYNOPSIS
       pwd [OPTION]...

DESCRIPTION
       Print the full filename of the current working directory.
```
- https://man7.org/linux/man-pages/man1/pwd.1.html

### Examples

Print the current working directory.

```bash
$ pwd
/workspaces/SparkleHub
```

### Change Directory `cd`

The `cd` command can be used to *change directories*.
This builtin is useful when you want to switch to a different directory to do a
little bit of work from there.

```bash
$ help cd
cd: cd [-L|[-P [-e]] [-@]] [dir]
    Change the shell working directory.

    Change the current directory to DIR.  The default DIR is the value of the
    HOME shell variable.
```

#### Examples

Change the current directory to your $HOME directory:

```bash
/workspaces/SparkleHub
$ cd ~/
/home/username/
$
```

Change back to the previous directory:

```bash
/home/username/
$ cd -
/workspaces/SparkleHub
$
```

### `mkdir`

The `mkdir` command can be used to *make* a new *directory*.

```bash
$ man mkdir
MKDIR(1)                              User Commands                             MKDIR(1)

NAME
       mkdir - make directories

SYNOPSIS
       mkdir [OPTION]... DIRECTORY...

DESCRIPTION
       Create the DIRECTORY(ies), if they do not already exist.
```

#### Examples

To create a new directory named `butterfly`.

```bash
/workspaces/SparkleHub
$ mkdir butterfly
```

If a directory named `butterfly` already exists then you might get an error when
you run the `mkdir` command. Adding the `-p` option means that `mkdir` will
create the directory if it doesn't already exist and will just do nothing
(instead of giving you an error) if it does exist:

```bash
/workspaces/SparkleHub
$ mkdir -p skittles-and-rainbows
```

[codespace]: ./github.html#codespaces
[git]: ./git.html
[integrated_terminal]: ./vscode.html#integrated-terminal
[powershell]: https://docs.microsoft.com/en-us/powershell/
[terminal.app]: https://support.apple.com/en-ca/guide/terminal/welcome/mac
