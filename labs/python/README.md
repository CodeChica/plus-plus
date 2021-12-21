# Lab Python

## Getting Started

In the commands below, the bit you need to type is after the dollar sign (`$`). That dollar sign and the bit before it is what your terminal will show you when it's waiting for you to type a command.

1. Open a [terminal](https://codechica.ca/guides/terminal.html)
1. Navigate to this lab by changing the working directory.

    ```bash
    /workspaces/plus-plus-main/ (main ✗) $ cd labs/python/
    ```

1. Run the tests:

    ```bash
    /workspaces/plus-plus-main/labs/python (main ✗) $ python test_code.py
    ```

This test is calling the `sum` function in `maths.py` and checking that it does what it should. The output from the test will include something like this:

```shell
AssertionError: 1 + 1 should be 2
```

It looks like that `sum` function isn't doing what it should. Let's fix it.

## Understanding the code

1. Open `maths.py` in your editor

You'll see some code that looks like this:

```python
def sum(x, y):
    return 5
```

If you've read the [JavaScript guide](https://codechica.ca/guides/javascript.html), you'll have heard of functions and variables. Python is a different language, but the idea of functions and variables is the same. We'll be using those ideas here.

This code defines a function. What's its name?

<details>
<summary>Open this for the answer</summary>

The function is called `sum`. The keyword `def` at the beginning means that a function is being `defined`.
</details>

This function has two _parameters_. A parameter is a special variable which is used to hold the data that is sent to the function.

What are the names of the two parameters here?

<details>
<summary>Open this for the answer</summary>

The parameters are `x` and `y`. Normally it's better to use more descriptive names for variables, but for simple code like this, `x` and `y` are fine.

Since our function is called `sum` and has two parameters, you could call it and provide it with its data by writing some code like this:

```python
sum(10, 20)
```

There's code like this in the tests.
</details>

Finally, you'll see the `return` keyword. Can you guess what that does?

<details>
<summary>Open this for the answer</summary>

`return` is used here to send a value back to the code which called the function. In this case, the function is called by the test and the test examines the value it gets back.
</details>

## Fixing the code

1. Create a new branch - replace `<your-username>` with your GitHub username.

    ```bash
    /workspaces/plus-plus-main/labs/pythong (main ✗) $ git checkout -b <your-username>-branch
    ```
1. Make the necessary changes in `maths.py` so that the tests will pass. Be sure to keep the spaces at the beginning of the line, since indentation (spaces at the start of lines) matters in Python - it shows how code is grouped together.

**Tip:** You use `+` (plus) to add numbers together in Python (and many other languages).

1. Run the tests to check they're passing:

    ```bash
    /workspaces/plus-plus-main/labs/python (main ✗) $ python test_code.py
    ```

If your changes have worked, you should see `OK`.

1. Commit your change.

    ```bash
    /workspaces/plus-plus-main/labs/python (main ✗) $ git commit -am "fix: make the sum function work"
    ```

### Optional: Extension task

If you'd like to try something a little more challening, give this extension task a go.

1. Open `test_code.py` in your editor and look for these lines:

```python
    @unittest.skip("extension task")
    def test_multiply(self):
```

1. Remove the whole of that `@unittest.skip("extension task")` line. Now that second test will no longer be skipped over when the tests run.

1. Run the tests again:

    ```bash
    /workspaces/plus-plus-main/labs/python (main ✗) $ python test_code.py
    ```

You'll see an error like this:

```shell
NameError: global name 'multiply' is not defined
```

This is a bit more complicated than last time - your tests are expecting a `multiply` method, but that doesn't even exist.

1. Copying what's been done for `sum` (don't forget the spaces at the start of the lines), can you add a new `multiply` method in `maths.py` which multiplies together its two parameters?

1. As you go, run the tests to see if they're passing:

    ```bash
    /workspaces/plus-plus-main/labs/python (main ✗) $ python test_code.py
    ```

If your changes have worked, you should see `OK`. You'll probably have to make several changes and re-run the tests a few times until you get it right. Don't worry - that's what normal programming is like!

**Tip:** You use `*` (asterisk) for multiplication in Python (and many other languages)

If you see an error like `IndentationError: expected an indented block`, it means that you forgot the spaces at the start of the code inside the function. You need four spaces to tell Python that this code belongs "inside" the function.

1. Commit your change.

    ```bash
    /workspaces/plus-plus-main/labs/python (main ✗) $ git commit -am "feat: add a multiply function"
    ```

### Submitting your work

1. Submit a Pull Request from your feature branch to the `main` branch.
1. :shipit:
