# Eliza

This is a simple port of the Eliza program to Go.

Eliza was originally written in the 1960s at the MIT Artificial Intelligence Laboratory by Joseph Weizenbaum. This code
is based on a Python program by Joe Strout, Jeff Epler and Jez Higgins. The original code is
available [here](https://github.com/jezhiggins/eliza.py).

You can learn more about Eliza on [Wikipedia](https://en.wikipedia.org/wiki/ELIZA).

## Running the code

To run this program go to the directory where this code lives, and type:

~~~
go run .
~~~

## Building an executable

Build in the usual way:

~~~
go build -o eliza main.go doctor.go
~~~
