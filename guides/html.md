## HTML

Hyper text markup language (HTML) is a markup language that is used to render
content in web browsers. (e.g. Apple Safari, Google Chrome, Mozilla Firefox).

{% include youtube.html youtube_id="XNAM6MjPfok" %}

HTML is used to describe the *content* of a document. Remember the internet was
first created as a way to share documents so that people could learn from each
other. So the terms *web page* and *document* mean the same thing.

The world's smallest HTML page:

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

Every HTML webpage must have an [`<html>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html),
[`<head>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head) and
[`<body>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body) element, like the one above.

* The [`<body>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body) is the part that you can see in a browser.
* The [`<head>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head) is the part that includes metadata about the web page.

HTML Elements (a.k.a. tags) have a name, and attributes.
Some elements are self closing and some elements can contain other elements.

* [`<a />`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a) is an *anchor* link.
* [`<div></div>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div) is used to *divide* content.
* [`<ul>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul) is an *unordered* list.
* [`<ol>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol) is an *ordered* list.
* [`<li>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li) is a list item and goes inside a list.
* [`<p>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p) is a *paragraph*.
* [`<h1>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/h1) is a level 1 heading
* ...
* [`<h6>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/h6) is a level 6 heading

[HTML elements reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element)
[Publishing via GitHub](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/Publishing_your_website#publishing_via_github)

{% include youtube.html youtube_id="y6stlwCb7VU" %}


Linking HTML and CSS files

{% include youtube.html youtube_id="ZO98acQRLgY" %}


Adding Images

{% include youtube.html youtube_id="hn2Upok0-x4" %}


Divs, images and flex wrap

{% include youtube.html youtube_id="i1ayJd3_qpE" %}



### Tools

Web developers rely on many tools to help them build amazing web sites. Below is
a list of tools that we will use in this course.

#### dev-tools

Web Developers use a **Web Browser** with [Dev Tools][devtools] like [Google Chrome][chrome].

{% include youtube.html youtube_id="VYyQv0CSZOE" %}

#### http-server

In this class we'll use an HTTP server. An HTTP server is a program that sends webpages to browsers.

There are two ways to install the `http-server`.
Download the latest version from [here](https://github.com/xlg-pkg/http-server/releases)
or install from the [terminal](./terminal.html) by running:

  ```bash
  $ go install github.com/xlg-pkg/http-server@latest
  ```

Then we can start the `http-server` from the [terminal](./terminal.html).

  ```bash
  $ http-server
  ```

[chrome]: https://www.google.com/chrome/
[devtools]: https://developer.chrome.com/docs/devtools/overview/
