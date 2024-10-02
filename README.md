# go-cmd-sample
go-cmd-sample



Shows how to put your "main" package code into a "cmd" folder to make things easier to manage, when you have multiple main.go files

## Step 1 - when you have a single "cmd/main.go" and no other .go files

1.1. Put your main.go into cmd
1.2. Run this to build it:
    go build ./...

## Step 2 - you then want more packages to add more code, and call those packages from your "main.go"

