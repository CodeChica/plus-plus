# JavaScript

JavaScript is a scripting language that was originally designed to run in a web
browser. It lets programmers write code that can execute in the web browser of
the visitors who visit a programmers website.

{% include youtube.html youtube_id="nItSSTwBvSU" %}

## How do I add JavaScript to my website?

To add JavaScript to our websites we can use the [`<script>`][script] tag to add inline or external JavaScript.

Here's an example of inline JavaScript in a file named `index.html`.

```html
<!-- index.html -->
<!DOCTYPE html>
<html>
  <head>
    <title>👀</title>
    <script>
      function clickButton() {
        alert("👋 Oh, hi!");
      }
    </script>
  </head>
  <body>
    <button type="button" onclick="clickButton();">Click Me</button>
  </body>
</html>
```
<hr />

Here's an example of how we can load an external JavaScript file named
`index.js` from an HTML file named `index.html`.

```html
<!-- index.html -->
<!DOCTYPE html>
<html>
  <head>
    <title>👀</title>
    <script type="text/javascript" src="index.js" ></script>
  </head>
  <body>
    <button type="button" onclick="clickButton();">Click Me</button>
  </body>
</html>
```

```javascript
// index.js
function clickButton() {
  alert("👋 Oh, hi!");
}
```

{% include youtube.html youtube_id="LFa9fnQGb3g" %}

<hr />

Sometimes our JavaScript will need to access HTML elements that are on the page.
To make sure that the HTML [DOM](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model)
is finished loading before our JavaScript code starts to execute we might need to wait until the [DOMContentLoaded](https://developer.mozilla.org/en-US/docs/Web/API/Window/DOMContentLoaded_event) event has
happened.


The code below will attach a function to the `click` event of the HTML [body element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body) so that when you click on the page it will show an alert with a
message.

```html
<!DOCTYPE html>
<html>
  <head>
    <title>👀</title>
    <script>
      document.addEventListener('DOMContentLoaded', function(event) {
        var bodyElement = document.querySelector('p');

        bodyElement.addEventListener('click', function(event) {
          alert("👋 Oh, hi!");
        })
      });
    </script>
  </head>
  <body>
    <p style="background-color: pink;">Click me!</p>
  </body>
</html>
```

Read [A re-introduction to JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/A_re-introduction_to_JavaScript) for a thorough tutorial on JavaScript.

{% include youtube.html youtube_id="_y9oxzTGERs" %}



## Data Types

* [Array](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array)
* [Boolean](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Boolean)
* [Date](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date)
* [Number](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number)
* [Object](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object)
* [String](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String)

## Variables

```javascript
let age = 14;
let birthday = new Date(2007, 09, 19);
let name = "Monalisa";
let mona = {
  Name: name,
  BirthDay: birthday,
  Age: age
};
```

* [Declarations](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements#declarations)

## Functions

```javascript
function greet(name) {
  return "Hello " + name;
}

let greeting = greet("Monalisa");
console.log(greeting);
```

{% include youtube.html youtube_id="AY6X5jZZ_JE" %}

Read more about [Functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/function).

## Loops

```javascript
for (let i = 0; i < 3; i++) {
  console.log("Happy Birthday to you");
}
```

* [Loops](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements#iterations)


## Additional Resources

{% include youtube.html youtube_id="QLatPwsbDrQ" %}

{% include youtube.html youtube_id="7qobB9DM0Ck" %}

[script]: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
