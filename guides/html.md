## HTML

Hyper text markup language (HTML) is a markup language that is used to render
content in web browsers. (.e.g. Apple Safari, Google Chrome, Mozilla Firefox).

{% include youtube.html youtube_id="XNAM6MjPfok" %}

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

{% include youtube.html youtube_id="y6stlwCb7VU" %}

### Tools

Web Developers use a **Web Browser** with [Dev Tools][devtools] like [Google Chrome][chrome].

::TODO tour of Google Chrome developer tools.
