# go-cmd-sample

This repo shows how to put your "main" package code into a "cmd" folder to make things easier to manage, when you have multiple main.go files.  
This seems a common pattern with go projects, especially those with multiple entrypoints.  

This is a bit unusual for me as a long time Java dev, as creating multiple Java classes, each with a 'main' method, does not cause any problems.  
The difference is, when running Java you usually specify the entrypoint class, whereas when running a binary built with Go you do not - the entrypoint is fixed at compile time.

## Step 1 - when you have a single "cmd/main.go" and no other .go files
(see "step_1" branch for complete code)

1.1. We place our "main.go" into the "cmd" directory.  
See the "step_1" branch for the code

1.2. Use this command to build the binary:  

    go build -o go-cmd-sample-app ./...

Notice the triple dot, which is a bit unusual. In Go, "x/..." means "every path under x/"

1.3. Run your app with:  

    ./go-cmd-sample-app

References:
- go build: [here](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)

## Step 2 - you then want more packages to add more code, and call those packages from your "main.go"
(see "step_2" branch for complete code)

2.1. We now introduce a new file "useful.go" in the root repo dir and use it from our "main.go"  

2.2. We modified main.go to reference package "useful" and call a function from it, and print the result.  
The new lines in main.go look like this:
```
usefulResult := useful.MyUsefulFunction()
fmt.Printf("usefulResult: %s\n", usefulResult)
```

Note: The above code uses the word "useful" to refer to the useful *package*, not the name of the useful.go file.

Note that main.go also needs an extra import line near the top:
```
import (
	"fmt"
	useful "go-cmd-sample"
)
```

2.3. Try to build using the same command as in Step 1  

    go build -o go-cmd-sample-app ./...

You'll see an error like this:
```
go: cannot write multiple packages to non-directory go-cmd-sample-app
```

This seems to be a quirk of Go when using a wildcard, so instead we'll 
use a specific name on the build command, as follows:  

    go build -o go-cmd-sample-app cmd/main.go

2.5. Run your app with:  

    ./go-cmd-sample-app

You'll see the output, including the string returned from the call out to the function in the 'useful' package.

