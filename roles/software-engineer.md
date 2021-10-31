![Elena the Engineer](./../assets/images/elena-the-engineer-slim.png)

_What image comes to mind when you think "engineer"? Many people think software engineers just sit at their computer writing lines of code all day. However, this role has many different aspects to it and uses lots of skills!_

By definition, a software engineer uses Computer Science and Programming principles to develop software solutions, build digital products, database programs, and computer systems for businesses.

> "Most developers create applications with the user experience in mind. Behind the scenes, they have to juggle multiple software projects and consult with their coworkers on how things should be done instead of just coding the entire time." - Xavier, Software Engineer at Codeacademy

Engineers get to collaborate with other people, such as Designers, Data Scientists, and Project Managers.

{% include youtube.html youtube_id="G-gAohI9qDA" %}

## What does an Engineer do?

* Collaborate with other people, such as Designers, Data Scientists, and Project Managers.
* Review code, do research, and meet with other teams to find the best solution for whatever problem they are trying to solve or feature they are building.
* Develop a functional piece of software.

Other related roles that have overlap with an Engineer are a Programmer and Web Developer.
You may see these terms used a lot, later on in your career journey.
Often, they are different names to describe the same role.
For example, some companies call programmers web developers or front-end engineers,
while others may call them software or UI (User Interface) engineers.
Web developers are focused on building user-facing applications.
Software engineers are more likely to work on computer systems as a whole.

> Did you know? Only 2% of employed Engineers are Latinas.

Software development is a male-dominated field, which can be discouraging.
We need more Latina and women Engineers! Knowing about the lack of diversity in
the tech world is all the more reason to keep pursuing your passion in coding,
break these barriers, and believe in yourself!
This field was founded by women like [Ada Lovelace][ada].

## Engineer's Guide to Building

As an Engineer, it is our task to work with Design to come up with a set of
wireframes to figure out how we can build this feature. Then it is up to the
Engineer to identify any technical risks with building the feature, breaking
down the work into smaller tasks and coming up with estimates for how long it
will take to build this feature.

When we have a flushed out User Story with Acceptance Criteria, one or more
wireframes, a task break down with estimates for the work then we can begin.

### Problem Solving Algorithm

1. Understand the problem.
1. Research solutions to this problem.
1. Choose a solution (or design a new one, if a solution does not exist)
1. Implement the solution.
1. Review the solution. (`goto` step 3 if the solution isn't working well)

When we start building this feature, we might identify new constraints that we
didn't think of when we were analyzing the work. Sometimes this new information
might change our estimates, or the direction for how we do this work. Because of
this it is important to continue to communicate progress with your team mates in
[Product management][product-manager] and [Design][designer] so that they can help by reducing scope
from the feature or putting together new wireframes that might make the problem
easier to solve.

To accomplish these goals, Engineers add [tasks](#task-breakdown) to a [User Story][user-story]
to break down a large feature into smaller pieces of work. Engineers, will try
to reduce the risk of delivering large features by shipping small changes more
often. Engineers will write [automated unit tests](#automated-tests),
apply [design principles](#design-principles), use [design patterns](#design-patterns)
and [refactor code](#refactoring) to keep code quality high and to make sure that
the code is easy to maintain. Access to new features can be controlled via
[feature flags](#feature-flags). This makes it possible to release a new feature in stages and
let a small population of users try out the feature before releasing it to
everyone.

## Automated Tests

{% include youtube.html youtube_id="OaMzkeUgZa8" %}

## Design Patterns

{% include youtube.html youtube_id="rtmFCcjEgEw" %}

## Design Principles

## Feature Flags

## Refactoring

## Task Breakdown

Before we start working on a new feature we need to make sure that we understand
the feature. To make sure that we understand the feature we will write a list of
tasks that we think we need to complete in order to complete the work. The list
of tasks can be as big or small as you want. Learning to break down large pieces
of work into small tasks takes practice.

Example task break down:

```plaintext
Tasks:

* [ ] Research how to create a confetti animation. (2 days)
* [ ] Build a prototype of the confetti animation. (2 days)
* [ ] Share a demo of the prototype and collect feedback. (1 day)
* [ ] Trigger the confetti animation after the Sparkle is saved. (1/2 day)
* [ ] Stop the animation after `n` seconds using a timer. (1/2 day)
* [ ] Measure the memory usage in the browser. (1 day)
* [ ] Add automated tests for each of the browsers that we support. (2 days)
* [ ] Work with Design to choose colours that are accessible. (1/2 day)
* [ ] Add acceptance tests for the new animation. (2 days)
```

[ada]: ../../../heroes/ada-lovelace.html
[designer]: ./../../roles/designer.html
[product-manager]: ./../../roles/product-manager.html
[user-story]: ./../../roles/product-manager.html#user-stories
