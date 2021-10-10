# 0x05 - The Software Engineer's role

## Introduction

In this lesson, you will learn an overview of the engineer's role in software development.
You will be placed in your teams and learn about the [final project requirements][final].
You will get to try out either the [Product Manager][product-manager],
[Designer][designer], or [Engineer][engineer] role!

By the end of this lesson you will:

* learn about the role of the Software Engineer.
* form teams and learn the [final project requirements][final].
* get to try out the role of either [Product Manager][product-manager], [Designer][designer] or [Engineer][engineer].

{% include youtube.html youtube_id="VSraiYX6Pl4" %}

[View Presentation](./slides.html) or learn more about [Elena the Engineer][engineer].

### Agenda

1. [The Software Game](./slides.html)
1. [Mini Coding Challenge](#warmup---mini-coding-challenge)
1. [Sam Klein, UX Designer](#guest-speaker--sam-klein)
1. [Example](#example--adding-confetti-to-sparklehub)
1. [Meet your teams](#activity---meet-your-teams)
1. [Heart, Fart and Shopping Cart](https://docs.google.com/document/d/1STo59fviyZraDr28txpKQUmQROLaVsEhyZT-QDtyjwc/edit?usp=sharing)

<hr />

## Warmup - Mini Coding Challenge

Get into groups and complete the following challenge.

1. Build the world's smallest website. [Hint](./../../guides/html.html)
1. Post a link to our website in the [#general][general] channel in [Slack][slack].

<hr />
## Guest Speaker - Sam Klein

{% include youtube.html youtube_id="rYeAHclkyCg" %}

<hr />

## Example - Adding Confetti to SparkleHub

The target audience is the Sparkler (person who is sending a Sparkle).
The feature is to add more fun!
The value is to notify the sender that the Sparkle has been sent and to create a
more interactive experience.

```plaintext
User Story: (Product Manager)

As a Sparkler,
I want it to rain confetti when I send a Sparkle,
so that I know that my Sparkle has been sent.

Acceptance Criteria: (Product Manager)

* Rainbow coloured confetti will rain down when a new Sparkle is created.
* The confetti will stop raining when it reaches the bottom of the screen.
* The text "You sent a Sparkle!" will appear behind the confetti.

Wire frames: (Designer)

  -------------------------------
  | o  @  o     o @ @ o @ o  0  |
  |   @ o    o @  @ o   o  o  o |
  |  o   @ o   @@ o   o  @ o o  |
  |      o                o     |
  |   o               o         |
  |                             |
  |     🎉 SPARKLE SENT  🎉     |
  |                             |
  |                             |
  |    /\/\      \🥳            |
  | /\/  \ \/\    ||\           |
  |/      \   \   /\            |
  -------------------------------
  |         SparkleHub          |
  -------------------------------

Tasks: (Engineer)

* [ ] Research how to create a confetti animation. (large)
* [ ] Build a prototype of the confetti animation. (large)
* [ ] Share a demo of the prototype and collect feedback. (medium)
* [ ] Trigger the confetti animation after the Sparkle is saved. (small)
* [ ] Stop the animation after `n` seconds using a timer. (small)
* [ ] Measure the memory usage in the browser. (medium)
* [ ] Add automated tests for each of the browsers that we support. (medium)
* [ ] Work with Design to choose colours that are accessible. (small)
* [ ] Add acceptance tests for the new animation. (medium)
```

<hr />

## Activity - Meet your Teams!

* For your final project, you will be collaborating in groups and wearing different hats, to give you the feeling of working in a real development environment!
* First, you will get to choose your groups. If students can't choose, the Instructor will assign teams.
* Then, once you are in your groups, choose one of the tech roles (Product Manager, Designer, Engineer). Each class, team members will switch and rotate roles!
* Share your feature idea & wireframe from last class with your team.
* As a group, choose 1-3 features to move forward with. Each team has to build at least one feature, if you’d like to include more, you have to prioritize which is most valuable.

<hr />

## Discussion

1. What does an engineer do?
1. What tech role sounds the most interesting to you?

<hr />

## Conclusion

In this lesson, you learned an overview of the engineer role in development,
and how they build features. You met your teams that you will be working with
on the final project, and chose a role. You began planning out your teams'
feature(s).

[ada]: ../../../heroes/ada-lovelace.html
[designer]: ./../../roles/designer.html
[engineer]: ./../../roles/software-engineer.html
[general]: https://codechica-plus-plus.slack.com/archives/C02CDMWDK7D
[product-manager]: ./../../roles/product-manager.html
[slack]: ./../../guides/slack.html
[user-story]: ./../../roles/product-manager.html#user-stories
[final]: ./../0x0A/
