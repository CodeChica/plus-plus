# Lab go

## Getting Started

1. Open a [terminal](https://codechica.ca/guides/terminal.html)
1. Navigate to this lab by changing the working directory.

    ```bash
    /workspaces/plus-plus-main/ (main ✗) $ cd labs/golang/
    ```

1. Run all the tests:

    ```bash
    /workspaces/plus-plus-main/labs/golang (main ✗) $ go test ./...
    ```

1. Create a new branch.

    ```bash
    /workspaces/plus-plus-main/labs/golang (main ✗) $ git checkout -b xlgmokha-branch
    ```

1. Open `sum.test.js` in your editor.
1. Remove each instance of `t.Skip()`
1. Re-run all the tests.

    ```bash
    /workspaces/plus-plus-main/labs/golang (xlgmokha-branch ✗) $ go test ./...
    ```

1. Make each test pass one by one.
1. Commit your change.

    ```bash
    /workspaces/plus-plus-main/labs/golang (main ✗) $ git commit -am "fix: a failing test"
    ```

1. Submit a Pull Request from your feature branch to the `main` branch.
1. :shipit:
