# Git

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

## Source Control

Programmers use an **editor** to write code, and a **terminal** to
issues commands to a compiler/transpiler/interpreter to convert
code into something that a computer can execute.

Programmers also need a way to save their code, back it up and share it with
others. This is where a source control tool can help us.

> Source control is a way to manage changes to code over time.

Sometimes people also call it **version control** or **software configuration management**.

These terms are meant to describe a way to make sure that we don't lose our changes that
we make to code and that we have a way to keep track of the different changes that
different people are are making to the code.

[git][git] is a tool used to keep a record of the changes that we make to our code.
[GitHub][github] is a website that lets us save our git repos so that we don't
have to worry if our computer crashes.

## Git Concepts

To get familiar with different concepts we will use
[Learn Git Branching][learngit] to help us get familiar with git.

{% include youtube.html youtube_id="UNCiMhrTyMo" %}

<!--
::TODO insert video going over the different git concepts in the terminal

* init
* add
* commit
-->

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
To get started we need to create our own fork (copy) of the project. To create a fork
open https://github.com/CodeChica/novela in a web browser. Now click on the fork
button. If a dialog pops up, choose your username as the location to fork the
project to.

{% include youtube.html youtube_id="vLHFQYPTwfw" %}

By `forking` the project, this instructs GitHub to create a copy of the entire
project just for you. In a few moments, you should get redirected to
https://github.com/<your-username>/novela. This is your very own copy of the
project that you are free to edit and do whatever you want with.

Let's `clone` this copy of the novela project to your computer so that you can
make edits to it. Replace '<your-username>' with your GitHub username.

e.g. `git clone git@github.com:xlgmokha/novela.git`.

```bash
$ git clone git@github.com:<your-username>/novela.git
$ cd novela
```

This cloned a repo from GitHub to a directory on our computer.
This project includes a file called `README.md` that has a story. This story
isn't finished yet, so we will need to add a new line to the story for others to
enjoy.

When you first clone the project it starts on a branch named `main`. `main` is
the name of the branch where we can find the latest version of the story. To add
a new line to the story we need to create a new branch that will be all yours.

```bash
$ git checkout -b <your-username>-story
```

Now that you have your own `branch`, let's make some edits.

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

### Conventional Commits

Writing effective commit messages for your future self and for others is
important to help explain why changes are introduced in a project.

Below is a brief guide on how to structure your git commit message using a
convention known as `semantic commit messages`.

```bash
Format: <type>(<scope>): <subject>

<scope> is optional

feat: add hat wobble
^--^  ^------------^
|     |
|     +-> Summary in present tense.
|
+-------> Type: chore, docs, feat, fix, refactor, style, or test.
```

* chore: updating build tasks etc; no production code change
* docs: changes to the documentation
* feat: new feature for the user, not a new feature for build script
* fix: bug fix for the user, not a fix to a build script
* refactor: refactoring production code, eg. renaming a variable
* style: formatting, missing semi colons, etc; no production code change
* test: adding missing tests, refactoring tests; no production code change

[Read more about Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

## Git in VSCode (SCM)

It's also possible to interact with git using the
[Source Control Manager (SCM)][scm] that is built into VSCode.
To learn more about it [read the documentation][scm] or watch the [VSCode
Tour](./vscode.html).

## Glossary

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
changes into the owner's repository.

Repository (repo)
: A repo is a directory of code managed by git.

[alacritty]: https://github.com/alacritty/alacritty
[codechica]: https://github.com/CodeChica/
[codespace]: ./github.html#codespaces
[git]: https://git-scm.com/
[git_handbook]: https://guides.github.com/introduction/git-handbook/
[github]: https://github.com/
[learngit]: https://learngitbranching.js.org/
[linux_commit]: https://github.com/torvalds/linux/commit/1da177e4c3f41524e886b7f1b8a0c1fc7321cac2
[scm]: https://code.visualstudio.com/Docs/editor/versioncontrol
[terminal]: ./terminal.html
[vscode]: https://code.visualstudio.com/
