# go-cmd-sample
go-cmd-sample



This repo shows how to put your "main" package code into a "cmd" folder to make things easier to manage, when you have multiple main.go files.  
This seems a common pattern with go projects, especially those with multiple entrypoints.  

This is a bit unusual for me as a long time Java dev, as creating multiple Java classes, each with a 'main' method, does not cause any problems.
The difference is, however, when running Java you usually specify the entrypoint, whereas when running a go binary you do not - the entrypoint is fixed at compile time.

## Step 1 - when you have a single "cmd/main.go" and no other .go files

(see branch: step_1)

1.1. Put your main.go into cmd
1.2. Run this to build it:
    go build -o go-cmd-sample-app ./...
1.3. Run your app with:
    ./go-cmd-sample-app

## Step 2 - you then want more packages to add more code, and call those packages from your "main.go"

TODO