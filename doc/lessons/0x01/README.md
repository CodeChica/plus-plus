# Lesson 0x01

* Introduction and review of HTML, CSS, Javascript
* Git Game with visuals to visualize source control.
* Introduction to source control.
  * why does it matter?

Outcomes:

* Girls will refresh coding skills, get excited about what they'll be learning.
* They will get to see SparkleHub, run setup and see the server.

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

## Why do we need source control?
## What is source control?
