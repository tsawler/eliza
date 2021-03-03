# Eliza

This is a simple port of the Eliza program to Go.

Eliza was originally written in the 1960s at the MIT Artificial Intelligence Laboratory by Joseph Weizenbaum.

This code is  based on a Python program by Joe Strout, Jeff Epler and Jez Higgins.

The original code is available [here](https://github.com/jezhiggins/eliza.py).

You can learn more about Eliza on [Wikipedia](https://en.wikipedia.org/wiki/ELIZA).


## Running the code

First, you have to install Go.

To run this program on a Mac, change to the directory where the code is, and type:

~~~
go run *.go
~~~

On Windows, type:

~~~
go run .
~~~

## Building an executable

Build in the usual way:

~~~
go build -o eliza main.go doctor.go
~~~