## CodeChica++

On this page you will find some guides for different subject that we will cover
in this course.

<!-- Start Terminal Guides -->
## Terminal

We use the terminal to interact with different commands and programs like
[git][git].

Each operating system comes with a different type of terminal and some commands
may work a little bit differently on different operating systems.

If you use macOS then you can open [Terminal.app][terminal.app].
On Microsoft Windows you can use [PowerShell][powershell].
For this course we will use the [terminal][integrated_terminal] that is built
into VSCode and a [Codespace][codespace].

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

List the all (including hidden) files using a long list format.

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
you run the `mkdir` command. To get the `mkdir` command to create the directory
if it isn't already there then use the `-p` option.

```bash
/workspaces/SparkleHub
$ mkdir -p skittles-and-rainbows
```

<!-- Start Git Guide -->
## Git

[git][git] is a tool that can be used to save changes that you make to your code.

[git][git] was created by Linus Torvalds so that he could make it easier for
others to contribute code to one of his projects, Linux. [git][git] lets us see
the very [first commit][linux_commit] for the Linux operating system.

When you `clone` a [git][git] repository this downloads a copy of all the
source code and history of changes to your computer. This make it possible for
you to read every single change (commit) that ever happened in a project.

Learning git will help you to contribute to Open Source Software,
become a better programmer, and learn how to look up historical artifacts about
code.

[git][git] allows us to make mistakes safely, and it makes it easier for us to
work together.

### Git Concepts

To get familiar with different concepts we will use
[Learn Git Branching][learngit] to help us get familiar with git.

### Git in the Terminal

In this section, you will learn how to use `git` from the terminal.

The best way to learn `git` is to use it. For this lesson you have the option of
using [your Codespace][codespace] or you can install [git][git] on your own
computer. You will also need a [terminal][terminal].

To get started we need to choose a directory on the computer where we can put
our code. Some coders like to use acronyms like `src`, which is short for
"source" or "source code". Let's follow this convention for now.

From your terminal create a directory for the [CodeChica][codechica] projects.

```bash
$ mkdir -p ~/src/github.com/CodeChica/
$ cd ~/src/github.com/CodeChica/
```

In this example we're going to contribute changes to an existing project.
To get started we need to create our own fork of the project. To create a fork
open https://github.com/CodeChica/novela in a web browser. Now click on the fork
button. If a dialog pops up, choose your username as the location to fork the
project too.

By `forking` the project, this instructs GitHub to create a copy of the entire
project just for you. In a few moments, you should get redirected to
https://github.com/<your-username>/novela. This is your very own copy of the
project that you are free to edit and do whatever you want with.

Let's `clone` this copy of the novela project to your computer so that you can
make edits to it. Replace '<your-username>' with your username. e.g. `git clone
git@github.com:xlgmokha/novela.git`.

```bash
$ git clone git@github.com:<your-username>/novela.git
$ cd novela
```

This cloned a repo from GitHub to a directory on our computer.
This project includes a file called `README.md` that has a story. This story
isn't finished yet, so we will need to add a new line to the story for others to
enjoy.

When you first cloned the project it starts on a branch named `main`. `main` is
the name of the branch where we can find the latest version of the story. To add
a new line to the story we need to create a new branch that will be all yours.

```bash
$ git checkout -b <your-username>-story
```

Now that you have your own `branch` let's make some edits.

Open `README.md` using your favourite editor or add a new line using the `echo`
command.

```bash
$ echo 'In a galaxy far far away...' >> README.md
```

Let's do a quick status check to see what's going on.

```bash
$ git status
On branch xlgmokha-story
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   README.md

no changes added to commit (use "git add" and/or "git commit -a")
```

[git][git] is telling us that it has detected that the file `README.md` has
changed. We can ask git to show us the changes using `git diff`

```bash
$ git diff
diff --git a/README.md b/README.md
index e45fe37..cc77fd0 100644
--- a/README.md
+++ b/README.md
@@ -1 +1,2 @@
 # la novela
+In a galaxy far far away
```

The `diff` output might look a little bit different for you because your sentence
is a little different from my mine, right?

LGTM, let's :shipit:!

Next, we can commit the change to tell [git][git] that we like the change and we
want it to save it for us. We're also going to need to give [git][git] a nice
description of the change so that in 100 years someone can read the description.

```bash
$ git add README.md
$ git commit -m 'Add a sentence to the story.'
```

Next, let's push our change from our computer up to GitHub. This will make sure
that we have a backup in the cloud just in case our computer crashes on us.

```bash
$ git push origin <your-username>-story
```

Next open https://github.com/your-username/novela in a
browser and follow the instructions to create a Pull Request.

Woohoo, we did things! But wait... what did we do?

### Git in VSCode (SCM)

It's also possible to interact with git using the
[Source Control Manager (SCM)][scm] that is built into VSCode.
To learn more about it [read the documentation][scm].

### GitHub

GitHub is a website that lets us save our git repositories for free. It is the
largest code hosting platform on the planet. It's a great place to create a
backup of your code and collaborate with other coders.

### Glossary

Branch
: A branch is a part of a tree which grows out from the trunk. It's also used to
describe an independent branch of code.

Commit
: A commit captures a group of changes in 1 or more files.

Diff
: A diff is the "difference" between the old version and the new version of a file.

Fork
: A fork is not only a useful utensil but it is also what we call a copy of a
project or repository.

Pull Request
: A pull request is a way to share code from one repository to another. The
author of the Pull Request will ask the owner of a repository to pull in their
changes into the owners repository.

Repository (repo)
: A repo is a directory of code managed by git.

<!-- Start Codespace Guide -->

## Codespaces

We use [Codespaces](https://github.com/codespaces) to make sure that everyone
has access to the tools and software that they need for this course
so that everyone can learn to code.

A [Codespace](https://github.com/codespaces) is a development environment
that is created just for you. You can think of it like it is your own
computer. When you create a Codespace it also comes with a bunch of tools
and software already installed. So you can focus on writing code instead of
figuring out how to install git, Linux, Ruby, Node.js etc.

### [Create my Codespace](https://docs.github.com/en/codespaces/developing-in-codespaces/creating-a-codespace#creating-a-codespace)

1. Open the [SparkleHub](https://github.com/codechica/SparkleHub) repo in a browser.
1. Click on "Code"
1. Click on "Codespaces"
1. Click on "New codespace"
1. Select "2 core"
1. Click on "Create codespace"
1. Click on "Open in VSCode"
1. Accept the project extensions
1. Be patient
1. Open a Terminal ("Terminal" -> "New Terminal")
1. Type `./script/server`
1. Open [http://localhost:3000/](http://localhost:3000/) in your web browser.

This is your copy of the final project that you are free to explore.
Take a look at the files in the directory [app/views](https://github.com/CodeChica/SparkleHub/blob/main/app/views/).

If you want to turn off the server press "Ctrl" + "c" in the terminal.

### [Delete my Codespace](https://docs.github.com/en/codespaces/developing-in-codespaces/deleting-a-codespace)

When you are finished with your Codespace it is a good idea to clean it up.

1. Open [Your codespaces](https://github.com/codespaces) in a web browser.
1. Click on the `...` menu on the right side.
1. Choose the "Delete" option.

## Markdown

Markdown is a lightweight and easy-to-use syntax for styling your writing. It includes conventions for

```markdown
Syntax highlighted code block

# Header 1
## Header 2
### Header 3

- Bulleted
- List

1. Numbered
2. List

**Bold** and _Italic_ and `Code` text

[Link](url) and ![Image](src)
```

For more details see [GitHub Flavored Markdown](https://guides.github.com/features/mastering-markdown/).

### Help

[Try the troubleshooting guide](https://docs.github.com/en/codespaces/codespaces-reference/troubleshooting-your-codespace)
or [open a new question](https://github.com/CodeChica/plus-plus/discussions/categories/q-a) if you get stuck.

[Edit this page](https://github.com/CodeChica/plus-plus/edit/gh-pages/guides/index.md)

[codechica]: https://github.com/CodeChica/
[codespace]: https://github.com/CodeChica/plus-plus/blob/main/doc/guides/codespaces.md#creating-your-codespace
[dotfiles]: https://dotfiles.github.io/
[git]: https://git-scm.com/
[integrated_terminal]: https://code.visualstudio.com/docs/editor/integrated-terminal
[learngit]: https://learngitbranching.js.org/
[linux_commit]: https://github.com/torvalds/linux/commit/1da177e4c3f41524e886b7f1b8a0c1fc7321cac2
[powershell]: https://docs.microsoft.com/en-us/powershell/
[scm]:  https://code.visualstudio.com/docs/editor/versioncontrol
[terminal.app]: https://en.wikipedia.org/wiki/Terminal_(macOS)
