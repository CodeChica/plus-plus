# JavaScript

{% include youtube.html youtube_id="_y9oxzTGERs" %}

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

* [Function](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/function)

## Loops

```javascript
for (let i = 0; i < 3; i++) {
  console.log("Happy Birthday to you");
}
```

* [Loops](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements#iterations)

<!--

TODO::
* `fetch()` data from the server.
* JSON API
* JSON
* REST
* variables
* datatypes
* functions
* events
-->
