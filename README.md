# my-cli
### https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177
> get module

`go get -u github.com/spf13/cobra`

> get CLI tool

`go get -u github.com/spf13/cobra/cobra`

> initiate go mod

`go mod init my-calc`

> install module

`go install my-calc`

> cobra init

`cobra init --pkg-name my-calc`

> add sub command

`cobra add add`

> add sub command to add

`cobra add even`

> add sub command to add (odd)

`cobra add odd`

* https://github.com/spf13/cobra/blob/master/shell_completions.md

> bash completion

`cobra add completion`

`my-cli completion bash > /etc/bash_completion.d/my-cli`
