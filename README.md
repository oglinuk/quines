# Quines: Self-Reproducing Programs

This is a repository containing different implementations of
self-reproducing programs. It is not meant to be a showcase of various
languages, but rather a showcase of different implementations/styles in
the languages I primarily use.

## What is a quine?

Ken Thomposon describes a quine in his paper [*Reflections on Trusting
Trust*](https://github.com/oglinuk/ken-thompsons-rott-quine/blob/master/rott.pdf)
as, "More precisely stated, the problem is to write a source program
that, when compiled and executed, will produce as output an exact copy of
its source."

## Why?

All it does it output it's source? What could be any use of that? The
description given by Ken does not limit the quine to what it can do in
the time between starting and outputting its own source. An example of
what I mean is the [self-compiling quines](self-compiling), which compile
themselves before outputting their source code. It is my intention to
explore the depths of quines and what they are truly capable of.

## TODO

* [ ] Implement a [basic quine](basic)
	* [X] [C](basic/c)
	* [X] [Go](basic/go)
	* [X] [Python 3](basic/python)
	* [ ] C++
	* [ ] BASH
	* [ ] JavaScript

* [ ] Implement a [hello world quine](hello-world)
	* [X] [C](hello-world/c)
	* [X] [Go](hello-world/go)
	* [ ] C++
	* [ ] BASH
	* [ ] JavaScript

* [ ] Implement a [self-compiling quine](self-compiling)
	* [X] [C](self-compiling/c)
	* [X] [Go](self-compiling/go)
	* [ ] C++

* [ ] Implement a [self-modifying quine](self-modifying)
	* [ ] C
	* [X] [Go](self-modifying/go)
	* [ ] C++

* [ ] Implement a [frankenquine](frankenquine)
	* [ ] C
	* [X] [Go](frankenquine/go)
	* [ ] C++
