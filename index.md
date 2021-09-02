## Welcome - CodeChica

Welcome to CodeChica++.

To make sure that you are prepared for the first class we've put together this
guide to help you get comfortable with the tools that we'll be using in this
class. All the tools that we'll be using in this course is **free** to download,
install and use.

The [CodeChica channel][youtube] contains all the videos that we'll be using for
this course.

The following links are pages that you might want to bookmark.

* [My Issues](https://github.com/issues/assigned)
* The [Plus Plus repo](https://github.com/CodeChica/plus-plus)
* The [SparkleHub repo](https://github.com/CodeChica/SparkleHub)

### Slack

We use [Slack][slack] for group chat to share content with each other and to ask
for help. If you haven't received an invitation to join the CodeChica++ slack
workspace then send me an [email][email] and we'll make sure that you get added.

::TODO Video tour

[Learn more about Slack](https://slack.com/intl/en-ca/features)

### GitHub

We'll be using our [CodeChica organization on GitHub][organization] to host our
code, provide development environments via [codespaces], facilitate
[discussions][discussions] and [track our progress][curriculum].

::TODO Video tour

[Learn more about GitHub][github].

### VS Code

We'll be using [VSCode][vscode] as our editor, integrated terminal and Codespace
explorer to work on labs and the final project.

::TODO Video tour

* Sign in to GitHub
* Method 1: Codespaces
* Method 2: Docker
* Method 3: Computer - [git](https://git-scm.com/downloads)

[Learn more about VSCode][vscode]

<!-- Start Lesson 0x01 -->

## Lesson 0x01

* Review HTML, CSS, Javascript
* Introduction to source control.
* Let's peek at the final project, SparkleHub.

### My First Issue

1. [Open](https://github.com/CodeChica/plus-plus/issues/new/choose) in a new tab.
1. Click the "Get Started" button next to "Lesson 0x01"
1. Click the "assign yourself" link.
1. Click "Submit new issue"

![Creating my first issue](/images/create-first-issue.gif)

This is your tracking issue to help you take notes and keep track of your
progress.

When you have completed all the steps in this issue add a comment to this issue
that says:

```plaintext
@xlgmokha I'm done!
```

If you get stuck on any of the sections add a comment to this issue that says:

```plaintext
@xlgmokha I need some help!
```

### HTML

Hyper text markup language (HTML) is a markup language that is used to render
content in web browsers. (.e.g. Apple Safari, Google Chrome, Mozilla Firefox).

HTML is used to describe the *content* of a document. Remember the internet was
first created as a way to share documents so that people could learn from each
other. So the terms *web page* and *document* mean the same thing.

The worlds smallest HTML page:

```html
<html>
  <head>
    <title>This text will appear in the browser tab</title>
  </head>
  <body>
    <div id="container">
      <h1>Hello, world!</h1>
      <p>
        I like sandwiches.
        Do you like sandwiches?
      </p>
    </div>
  </body>
</html>
```

Every HTML webpage must have an `<html>`, `<head>` and `<body>` element, like
the one above.

* The `<body>` is the part that you can see in a browser.
* The `<head>` is the part that includes metadata about the web page.

HTML Elements (a.k.a. tags) have a name, and attributes.
Some elements are self closing and some elements can contain other elements.

* `<a />` is an *anchor* link.
* `<div></div>` is used to *divide* content.
* `<ul>` is an *unordered* list.
* `<ol>` is an *ordered* list.
* `<li>` is a list item and goes inside a list.
* `<p>` is a *paragraph*.
* `<h1>` is a level 1 heading
* ...
* `<h6>` is a level 6 heading

### CSS

CSS is short for Cascading Style Sheets. It's a language that is used to tell
the browser how to make your page look.

In CSS you specify which types of elements should get certain settings.

e.g.

```css
/* Find all 'div' elements and give them a pink background. */
div {
  background-color: pink;
}

/* Find all elements with a `hero` class and center their text. */
.hero {
  text-align: center;
}

/*
  Find the one element with an id of `container`
  and give it a width of 90% of the screen then make sure
  the left and right margins are even so that the content
  is centered.
*/
#container {
  width: 90%;
  margin-left: auto;
  margin-right: auto;
}
```

### JavaScript

JavaScript is a language that allows you to write code that can react to events
in the browser. When you *click* on a button, this raises an event that can be
handled by JavaScript code. JavaScript can then be used to manipulate the
Document Object Model (a.k.a. DOM) to perform actions like making a network
request to a server or validate data entered on a page.

e.g.

The code below will render an alert box with the text "Ouch!" when you click
on a page.

```javascript
document.addEventListener('DOMContentLoaded', function(event) {
  var bodyElement = document.querySelector('body');

  bodyElement.addEventListener('click', function(event) {
    // This is a great way to irritate visitors to your website.
    alert('Ouch!');
  })
});
```

### Activity

Break up into groups of 2 or 3 and work through the sections named `Introduction to HTML`
and `Customize your profile` together.

### Source Control

Programmers need a few tools in order to write code.

* Editor ([VSCode][vscode])
* Terminal ([Alacritty][alacritty])
* Source Control ([git][git])

We use our **editor** to edit code, and we use our Terminal to convert that
code into something that we can run. We also use a Terminal to interact with our
source control tool to save our changes and push them up to GitHub to make sure
that we don't lose any of our work.

Web Developers also need a **Web Browser** with [Dev Tools][devtools] like
[Google Chrome][chrome].

Programmers also need a way to save their code, back it up and share it with
others. This is where a source control tool can help us.

> Source control is a way to manage changes to code.

Sometimes people also call it by other names like **version control**, and
**software configuration management**.

It's meant to describe a way to make sure that you don't lose code and that
programmers have a way to keep track of different changes that are made to code.

In this class, we'll use a tool called [git][git] to keep a
record of the changes that we make to our code. Then we'll use a website called
[GitHub (free)][github] to save our code for us so that we don't lose it.

Let's play [The Git Game][git_game].

<!-- Start Lesson 0x02 -->
## Lesson 0x02

<!-- Start Lesson 0x03 -->
## Lesson 0x03

<!-- Start Lesson 0x04 -->
## Lesson 0x04

```plaintext
# Agenda

* Introduction
* Product Overview
* Customer Feedback


Feedback:

* view profile of peers (high priority)
* no way to see sent/received sparkles
  * see sparkles that we got self esteem boost
  * forgot to sparkle certain people.
    * can't see who received
  * filter by our sparkles that we sent
  * filter by our sparkles that we receive
  * separate page to see my sparkles
  * privacy settings
    * private sparkles.
      * interpersonal conflicts.
* no notifications when sending/receiving
* avatar icon bug
  * login via Facebook
  * login via GitHub
* colours: branding colors
  * colour palette
  * red (need to chat with marketing)
  * align colors with branding. (custom colors)

authenticating users
* what kind of timeframe are we looking at?
  * 3 months to the full organization.

Spike:
* GitHub Auth
* Facebook Auth

Temporary solution:
* add a textbox to enter username

Remove Avatar for now.
```

<!-- Start Lesson 0x05 -->
## Lesson 0x05

<!-- Start Lesson 0x06 -->
## Lesson 0x06

## [Guides](/guides/)
## [Issues](https://github.com/CodeChica/plus-plus/issues/choose)
## [Plus Plus](https://github.com/CodeChica/plus-plus)
## [SparkleHub](https://github.com/CodeChica/SparkleHub)

## Help!

Having trouble? Check out our [discussions][discussions] or ask for help in [Slack][slack] and we’ll help you sort it out.

[Edit this page](https://github.com/CodeChica/plus-plus/edit/gh-pages/index.md)

[alacritty]: https://github.com/alacritty/alacritty
[chrome]: https://www.google.com/chrome/
[curriculum]: https://github.com/CodeChica/plus-plus/issues/new/choose
[devtools]: https://developer.chrome.com/docs/devtools/
[discussions]: https://github.com/CodeChica/plus-plus/discussions
[docker]: https://docs.docker.com/get-docker/
[email]: mailto:mo@mokhan.ca&subject=CodeChica++
[git]: https://git-scm.com/
[git_game]: https://learngitbranching.js.org/
[github]: https://github.com/
[organization]: https://github.com/CodeChica
[q-and-a]: https://github.com/CodeChica/plus-plus/discussions/categories/q-a
[slack]: https://codechica-plus-plus.slack.com/
[vscode]: https://code.visualstudio.com/
[youtube]: https://www.youtube.com/playlist?list=PLaZatV79bZCRtD6yCw-goNH5Keh8ovMQp
[zoom]: https://zoom.us/
