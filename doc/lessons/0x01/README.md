# Lesson 0x01

* Introduction and review of HTML, CSS, Javascript
* Git Game with visuals to visualize source control.
* Introduction to source control.
  * why does it matter?

Outcomes:

* Girls will refresh coding skills, get excited about what they'll be learning.
* They will get to see SparkleHub, run setup and see the server.

## Create your first Issue

1. [Open](https://github.com/CodeChica/plus-plus/issues/new/choose) in a new tab.
1. Click the "Get Started" button next to "Lesson 0x01"
1. Click the "assign yourself" link.
1. Click "Submit new issue"

This is your tracking issue to help you take notes and keep track of your
progress.

## Review

* HTML
* CSS
* JavaScript

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

## Tools

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

### Source Control

> Source control is a way to manage changes to code.

Sometimes people also call it by other names like **version control**, and
**software configuration management**.

It's meant to describe a way to make sure that you don't lose code and that
programmers have a way to keep track of different changes that are made to code.

In this class, we'll use a tool called [git][git] to keep a
record of the changes that we make to our code. Then we'll use a website called
[GitHub (free)][github] to save our code for us so that we don't lose it.

## Activity Time

Break up into groups of 2 or 3 and work through the sections named `Getting Started`
and `Customize your profile` together.

[alacritty]: https://github.com/alacritty/alacritty
[chrome]: https://www.google.com/chrome/
[devtools]: https://developer.chrome.com/docs/devtools/
[git]: https://git-scm.com/
[github]: https://github.com/
[vscode]: https://code.visualstudio.com/
