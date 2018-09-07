Error or Panic
-

[![CircleCI](https://circleci.com/gh/cn007b/eop.svg?style=svg)](https://circleci.com/gh/cn007b/eop)
[![Go Report Card](https://goreportcard.com/badge/github.com/cn007b/eop)](https://goreportcard.com/report/github.com/cn007b/eop)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b05afa73b94146d4bd65543d32e7627b)](https://www.codacy.com/app/cn007b/eop?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cn007b/eop&amp;utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/caf54b60fe8ae4d9a423/maintainability)](https://codeclimate.com/github/cn007b/eop/maintainability)

Here you can find 2 approaches how to deal with errors:
first one - return error, another one - panic.
<br>Which better - it's up to you!

## Prerequisites

Common stuff between all approaches - workflow:
here I have one controller which calls service function
which works like a facade and performs sign up into:

* facebook
* twitter
* pinterest

you can see it on sequence diagram:

![sequence diagram](/sequenceDiagram.png)

## The point

You can go through `src/app/examples` directory and find implementations for different approaches,
you can compare different approaches and decide which one is better for you.

## Usage

To check that all approaches work in same way you can run:

````bash
git clone https://github.com/cn007b/eop.git && cd eop
export GOPATH=$PWD

# 1: return error
go run -race src/app/main.go --strategy=error
go run -race src/app/main.go --strategy=error --username=bond
go run -race src/app/main.go --strategy=error --username=james

# 2: panic
go run -race src/app/main.go --strategy=panic
go run -race src/app/main.go --strategy=panic --username=bond
go run -race src/app/main.go --strategy=panic --username=james

# 3: return error from goroutine
go run -race src/app/main.go --strategy=error_goroutine
go run -race src/app/main.go --strategy=error_goroutine --username=bond
go run -race src/app/main.go --strategy=error_goroutine --username=james

# 4: panic in goroutine
go run -race src/app/main.go --strategy=panic_goroutine
go run -race src/app/main.go --strategy=panic_goroutine --username=bond
go run -race src/app/main.go --strategy=panic_goroutine --username=james

# 5: panic in goroutine like PRO 😎
go run -race src/app/main.go --strategy=panic_goroutine_pro
go run -race src/app/main.go --strategy=panic_goroutine_pro --username=bond
go run -race src/app/main.go --strategy=panic_goroutine_pro --username=james
````
