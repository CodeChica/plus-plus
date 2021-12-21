# CSS

CSS is short for Cascading Style Sheets. It's a language that is used to tell
the browser how to make your page look.

In CSS you specify which types of element should get certain settings.

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

## Grid Layout

{% include youtube.html youtube_id="rPOX2hcaFvw" %}

<pre>
 Desktop                           Mobile

 -------------------------------   ----------
 | header                      |   | header |
 -------------------------------   |--------|
 |     ||                      |   | nav    |
 |     ||                      |   |--------|
 | nav || main                 |   |        |
 |     ||                      |   | main   |
 |     ||                      |   |        |
 |     ||----------------------|   |--------|
 |     || footer               |   | footer |
 -------------------------------   ----------
</pre>

```html
<!DOCTYPE html>
<html>
  <head><title>Grid Layout</title></head>
  <body>
    <header></header>
    <nav></nav>
    <main></main>
    <footer></footer>
  </body>
</html>
```

```css
body {
  min-height: 100vh;
  display: grid;
  grid-gap: 1em;
  grid: 'header' auto 'nav' auto 'main' 1fr 'footer' auto / 1fr;
}
header { grid-area: header; }
nav { grid-area: nav; }
main { grid-area: main; }
footer { grid-area: footer; }

@media (min-width: 40em) {
  body {
    grid: 'header header' auto 'nav main' 1fr 'nav footer' auto / 12em 1fr;
  }
}
```

[Grid layout](https://codepen.io/miriamsuzanne/pen/JjPeQYP?editors=1100)

## Flex Layout

[Flexbox](https://developer.mozilla.org/en-US/docs/Learn/CSS/CSS_layout/Flexbox)

<!--
TODO:: Describe

* inline css
* document `<style>` section
* external css file.
* code comments
* selectors
  * element type
  * id - there can only be one
  * class
  * '.img' vs '#img' vs 'img'
* grid system
  * flex vs percentage/pixels
    * 960/12
* box model
* responsive design (desktop, mobile, tablet)
* confetti css
* testing a selector in the browser. `querySelector(selector)`

-->

## Animations

{% include youtube.html youtube_id="etQEBtvI4zQ" %}
