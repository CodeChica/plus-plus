# 0x03 - Collaborating with others

<!--
* 30 min: CSS Animations demo. [Sparkle Party](https://github.com/codechica/sparkle-party)
* 30 min: Finish working on [Lesson 1 issue][issues].
* 30 min: Presentation on Source Control
* 15 min: [Play a game together](https://create.kahoot.it/share/git-terms/6bf90eb7-3c80-43d5-a0ff-b710cc767040)
* 60 min: Work on assigned [Lesson 2 issue][issues]
* 15 min: Heart, Fart and Shopping Cart


Part 1: Get an existing suite of tests passing.
Part 2: Start contributing tests that might be missing.
Part 3: Reproduce a bug as an automated test.
Part 4: Fix it
Part 5: Refactoring. Make it DRY, YAGNI and KISS(parkles). Gorilla/Banana.

Introduction to tests and pair programming. Break into pairs, work on a bug
fixing (make the tests pass), submit the changes as a pull request.
Provide example (SparkleHub) with problems to fix, work on the bug fix & share
how they fixed it. Get practice reviewing each others code and giving feedback.

Learn the error message. "What's wrong with this code" challenge.
  * Syntax, logic

How often will they be debugging?
How will this differentiate them from other devs?

## Outcomes

* Understand how to problem solve.
* Continue to work collaboratively to provide feedback.
* Learn the value of clean coding.
* Why are tests important?

[slides](./slides.html)

-->

Agenda:

* Group Activity (1 hour)
* Guest Speaker [@asciimike](https://github.com/asciimike) (30-45 minutes)
* Test Driven Development (1 hour)
* [Heart, Fart, Shopping Cart](https://docs.google.com/document/d/1mV7g-5E4GdLOG0j4l2xhJGNCw32uSDB9LMj6vbIFU9g/edit?usp=sharing)

## Group Activity

1. Create teams of 3 people.
1. Choose who will be the:
  * Product Manager
  * Designer
  * Engineer
1. Present your GitHub profile page to each other.
1. Pick one of the profile pages to improve.
  * Create a [User Story](./../../roles/product-manager.html#user-stories) to improve the profile page.
  * Create a [Wireframe](./../../roles/designer.html#what-is-a-wireframe) that satisfies the User Story.
  * Update the profile page to match the design.
1. Presentation
  * The Product Manager will present the teams [User Story](./../../roles/product-manager.html#user-stories)
  * The Designer will present the teams [Wireframe](./../../roles/designer.html#what-is-a-wireframe)
  * The Engineer will present the updated profile page.

{% include youtube.html youtube_id="AGfWZlAikPQ" %}

## Introduction

In this lesson we will introduce testing and pair programming.
We will continue to work collaboratively by submitting pull requests and reviewing each other's code.
We will also learn the value of writing clean, readable code.

By the end of this lesson:

* you will get an existing suite of test passing.
* you will write your first test.
* you will reproduce and fix your first bug.
* you will review code and give feedback.
* you will refactor your code.

{% include youtube.html youtube_id="1_uRNFkFX5Q" %}

## Why Test Driven Development (TDD)?

Test driven development is the process of writing tests before actually coding anything.
This helps a developer think about the requirements of a feature they're working on before writing the code for it.
You can be more sure that you haven't missed anything important about a feature's needs when you use this process.

Test driven development also makes refactoring much easier because we don't have to worry about breaking any existing code.
If we make changes, and something does change, a test will fail (essentially warning us), and we can make sure everything works again before pushing changes to the main branch of our project.
TDD also helps with better code coverage, which leads to fewer bugs in your application.
Finally, you can look back on your code and know exactly what you were trying to do by reading your tests.

To learn more, read our [Ruby Guide](../../guides/ruby.html)

### Write your first test

When we write tests, we follow the steps of "Red, Green, Refactor".
What this means is, we first write a test for a feature we want to build.
That feature has not been written yet, so our test will fail at first (red).
We write the simplest code to get our test to pass (green), and finally refactor to make our code better.

Use this pattern to get started:

> When \<context\> is established, it \<behavior\>

Example:

> When a sparkle is created, it saves the recipient

> When a sparkle is created, it saves the reason

```rb
require "minitest/autorun"

describe "Sparkle" do
  describe "when a sparkle is created" do
    it "saves the recipient" do
    end
  end
end
```

In order to write our test, we're going to use the pattern "Arrange, Act, Assert".

If you think of putting on a play, first we have the set the stage.

That's our "arrange" in this example:

```rb
require "minitest/autorun"

describe "Sparkle" do
  describe "when a sparkle is created" do
    before do
      # arrange
      @recipient = "@monalisa"
      @reason = "for helping me with my homework!"
    end

    it "saves the recipient" do
    end
  end
end
```

Next, is the "act" portion. We're "acting" upon our variables that we set up earlier.

```rb
require "minitest/autorun"

describe "Sparkle" do
  describe "when a sparkle is created" do
    before do
      # arrange
      @recipient = "@monalisa"
      @reason = "for helping me with my homework!"

      # act
      @sparkle = Sparkle.create(@recipient, @reason)
    end

    it "saves the recipient" do
    end
  end
end
```

Finally, we have "assert" which is checking to see if the result matches what we expect it to be.

```ruby
require "minitest/autorun"

describe Sparkle do
  describe "when a sparkle is created" do
    before do
      # arrange
      @recipient = "@monalisa"
      @reason = "for helping me with my homework!"

      # act
      @sparkle = Sparkle.create(@recipient, @reason)
    end

    it "saves the recipient" do
      # assert
      assert_equal @recipient, @sparkle.sparklee
    end
  end
end
```

Remember, our test won't pass right away because we haven't written the code for the feature yet, but we want to make sure that our test is failing for the right reasons.
From here, you should get a failing test message that looks like:

```bash
Expected: "@monalisa"
  Actual: nil
```

This tells us that our code currently doesn't satisfy the conditions described in the test.

When giving feedback, use the "critique sandwich" method.

1. Say something positive about your partner's code.
1. Give your critiques. What could be better? What do you have questions about?
1. End with another positive!

## Refactor

When you receive notes from your teammates, now is when you can look at how to improve your code.

Keep these things in mind when refactoring:

1. Does your code do what it should do well?
2. KISS (Keep it simple, silly).
3. Functions should be responsible for one thing in general. This is the Single Responsibility Principle (SRP). Do you have any functions that are long and can be broken up in this way?
4. Keep it DRY. DRY stands for "Don't Repeat Yourself". Do you see any code that repeats and could be put into it's own function to simplify your code? If so, go ahead and do that!
5. Keep methods small. Create chunks of code that you can reuse in different places.

None of these are hard and fast rules, so use your best judgement. Ultimately, we should care about "readability". Will you be able to make sense of your code when you come back to it later? Will others be able to understand your code and what it's doing?

### Rinse and repeat

Repeat the peer review and refactor steps until you are happy with the state of your code. Then, start from Step 1 again, and write more tests!

## Conclusion

In this lesson, you learned about test driven development and the reasons for its use. You also learned some intro Ruby concepts in order to get your first suite of tests to pass and to write your own tests. Finally, you practiced peer reviewing someone else's code, and refactoring to make your code better.

<!--
## Lesson 0x03

Software Engineering Lifecycle

* Thought -> Product
* Planning Stage, Requirements, Feasibility, Design, Prototyping,
  Implementation, Testing.
* Mini pop quiz: Which phase are we in right now?
* Encourage them to research outside of class.

Introduction to roles: What are the different types of jobs in tech?

* Product Manager
  * Product manager's guide to prioritizing things
  * Activity: Girls brainstorm a different feature for SparkleHub
* Designer
  * Designer's guide to designing things.
  * Wireframes - showing them the tools that developers use in the real world (how to create lines, shapes).
  * Usability, UI, UX, Design
  * Girls draw out what their feature idea could look like:
    Activity: Wireframe

Outcomes:

* Girls will learn about the software engineering life cycle, and explore an
  overview of the product manager & designer roles in tech.
* Girls will come up with an idea for a new feature, and a wireframe for it.


As an instructor, I want to be able to teach lesson 0x03 so that students can
understand the software engineering life cycle.

* [ ] Students learn an overview of what a product manager does.
* [ ] Students learn an overview of what a designer does.
* [ ] Students create a user story as the product manager.
* [ ] Students create a wireframe as a designer.


* [Presentation](./slides.html)


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

-->
