***This code is not maintained, and may break in future versions of Go. You will not get support.***

# gocalc
This calculator is a command-line tool that supports basic arithmetic operations, variables, and functions. It reads input from the user, parses it into tokens, and performs the appropriate calculations. The calculator will continue to run until the user closes it, allowing the user to enter multiple expressions.

## Syntax
Simple expressions are arithmetic expressions that consist of numbers and operators. The calculator supports the following operators:

- **+**: addition
- **-**: subtraction
- **\***: multiplication
- **/**: division

For example, the following are valid simple expressions:

```
2 + 3
4 - 5
6 * 7
8 / 2
```

## Variable definitions
Variable definitions allow the user to assign a value to a variable and reuse it in future expressions. Variable definitions have the following syntax:

```go
def VARNAME VALUE
```

where `VARNAME` is the name of the variable and `VALUE` is the value to assign to the variable. `VALUE` can be a number or another variable.

For example, the following are valid variable definitions:

```
def x 2
def y 3
def z x + y
```

## Function definitions
Function definitions allow the user to define custom functions and reuse them in future expressions. Function definitions have the following syntax:

```go
fn FUNCNAME ARG1 ARG2 ...
```

where `FUNCNAME` is the name of the function and `ARG1`, `ARG2`, etc. are the arguments to the function. Arguments can be numbers or variables.

The calculator supports the following built-in functions:

- sin(x): calculates the sine of x (in radians)
- cos(x): calculates the cosine of x (in radians)
- mean(x1, x2, ...): calculates the mean of all the arguments

For example, the following are valid function definitions:

```
fn addTwo x y
fn square x x * x
```

## Examples
Here are some examples of expressions that the calculator can evaluate:

```
> 2 + 3
Result: 5.00
> 4 - 5
Result: -1.00
> 6 * 7
Result: 42.00
> 8 / 2
Result: 4.00
> def x 2
> def y 3
> def z x + y
> z * 4
Result: 20.00
> fn addTwo x y
> addTwo 3 4
Result: 7.00
> fn square x x * x
> square 5
Result: 25.00
> sin(0.5)
Result: 0.47
> cos(0.5)
Result: 0.87
> mean(1, 2, 3, 4, 5)
Result: 3.00
```
