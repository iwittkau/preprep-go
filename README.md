preprep-go
===

A very simple Go source file pre-processor to ...

**It works, but the docs are still work in progress.**


# Why?

Why not? (I'll explain later ...)


# How?

`preprep-go` will convert comments in go files to `cpp` compatible directive blocks.


# Installation

```
go get -u github.com/iwittkau/preprep-go
``` 

This will install `preprep-go` in `$HOME/go/bin/`.

# Example 

Clone this repository and run `go generate ./...` from the project root.

## Generate a `.pgo` file

The `go:generate` directive will convert the `main.go` to a `main.pgo` using `preprep-go`.

[embedmd]:# (example/main.go)
```go
//go:generate preprep-go -i main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello from main")
	//#ifdef FEATURE
	feature()
	//#endif
}
```

This will generate a `main.pgo` like this:

[embedmd]:# (example/generated/example_main.pgo)
```pgo
//go:generate preprep-go -i main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello from main")
#ifdef FEATURE
	feature()
#endif
}
```


## Generate a new `.go` file from a `.pgo` file, using `cpp`

You can now use `cpp` to convert this `main.pgo` to a new `main.go`. The [`Makefile`](Makefile) in this repository uses a different folder for this:

```bash
mkdir build
cpp -DFEATURE -P example/main.pgo build/main.go
```

Now, there is the preprocessed `main.go` with the `feature()` call, because we have set `FEATURE`

[embedmd]:# (example/generated/example_main.go)
```go
//go:generate preprep-go -i main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello from main")

	feature()

}
```

## Compiling the final binary

Our new, pre-processed `main.go` calls `feature()` which is also part of the `main` package, but located in the `feature.go` file. This file needs to be present when we want to build `example` with `feature`. Also, there needs to be a build tag set in `feature.go`.

``` 
cp example/feature.go build/
go build -tags feature -o xmpl-feat ./build
rm -r build
``` 

# Ceavats

`example/` will not compile unless all tags are provided during build.


# Inspiration

This article by Svetlin Ralchev inspired `preprep-go`: 

> Go does not have a preprocessor to control the inclusion of platform specific code. Even though C preprocessor is intended to be used only with C, C++, and Objective-C source code, we will use it as a general text processor of Go source code.

[Svetlin Ralchev - Conditional compilation in Golang](http://blog.ralch.com/tutorial/golang-conditional-compilation/)
