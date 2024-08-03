# CSC 3310 Concepts of Programming Languages -- Go Programming Assignment: 4Point Grammar Lexical and Syntax Analysis


## Description
This was a project done in my Concept of Programming Languages course at SPU in October of 2021. The overall objective of this assignment was to demonstrate understanding of the compilation process by implementing the lexical and syntactical analysis phases. The instructions for this assignment were as follows:
 
Write a program in Go that takes a program written in 4Point, and outputs:
1. If the code has lexical or syntax errors, the error that was found. Use panic version of error handling (once an error is found report the error and stop the process).
1. If the code is OK, depending on a command line flag the program will produce:
   1.	If the flag is `-s` the program will output function calls in Scheme that is going to be called by a program in Scheme that will calculate properties of those three points.
   1. If the flag is `-p` the program will output a series of queries about those three points.

The program runs like this:
```
prompt>go run . input.txt -s
; Processing Input File input.txt
; Lexical and Syntax analysis passed
; Generating Scheme Code
(process-triangle (make-point 2 3) (make-point 1 4) (make-point 3 4))
prompt>
```

## Grammar of 4Point

```
START      --> STMT_LIST
STMT_LIST  --> STMT. |
               STMT; STMT_LIST
STMT       --> POINT_DEF |
               TEST
POINT_DEF  --> ID = point(NUM, NUM)
TEST       --> test(OPTION, POINT_LIST)
ID         --> LETTER+
NUM        --> DIGIT+
OPTION     --> triangle |
               square
POINT_LIST --> ID |
               ID, POINT_LIST
LETTER     --> a | b | c | d | e | f | g | ... | z
DIGIT      --> 0 | 1 | 2 | 3 | 4 | 5 | 6 | ... | 9

```

The tokens of this grammar are:

Token | Lexeme
------ | ------
`POINT` | `point`
`ID` | `identifier`
`NUM` | `234`
`SEMICOLON` | `;`
`COMMA` | `,`
`PERIOD` | `.`
`LPAREN` | `(`
`RPAREN` | `)`
`ASSIGN` | `=`
`TRIANGLE` | `triangle`
`SQUARE` | `square`
`TEST` | `test`

Given the following program written in this language:
```
a = point(2, 3);
b = point(1, 1);
c = point(1, 3);
d = point(0, 0);
test(square, a, b, c, d);
test(triangle, a, b, c).
```
The tokens that it would generate are:
```
ID  a
ASSIGN
POINT
LPAREN
NUM 2
COMMA
NUM 3
RPAREN
SEMICOLON
ID  b
ASSIGN
POINT
LPAREN
NUM 1
COMMA
NUM 1
RPAREN
SEMICOLON
ID  c
ASSIGN
POINT
LPAREN
NUM 1
COMMA
NUM 3
RPAREN
SEMICOLON
ID  d
ASSIGN
POINT
LPAREN
NUM 0
COMMA
NUM 0
RPAREN
SEMICOLON
TEST
LPAREN
SQUARE
COMMA
ID a
COMMA
ID b
COMMA
ID c
COMMA
ID d
RPAREN
SEMICOLON
TEST
LPAREN
TRIANGLE
COMMA
ID a
COMMA
ID b
COMMA
ID c
RPAREN
PERIOD
```

Notice that the ID and NUM tokens have their lexeme associated. Also notice that in the language the elements do not need to be separated by space.


## How to run the program

The following examples assume that `input.txt` contains the following code:
```
a = point(2, 3);
b = point(1, 1);
c = point(1, 3);
d = point(0, 0);
test(square, a, b, c, d);
test(triangle, a, b, c).
```

### Scheme Output
To generate scheme output you will add the `-s` flag at the end of the command:
```
prompt> go run . input.txt -s
; processing input file input.txt
; Lexical and Syntax analysis passed
; Generating Scheme Code
(process-square (make-point 2 3) (make-point 1 1) (make-point 1 3) (make-point 0 0))
(process-triangle (make-point 2 3) (make-point 1 1) (make-point 1 3))
```

### Prolog Output
To generate scheme output you will add the `-p` flag at the end of the command:
```
prompt> go run .  input.txt -p
/* processing input file input.txt
   Lexical and Syntax analysis passed
   Generating Prolog Code */

 /* Processing test(square, a, b, c, d) */
 query(square(point2d(2, 3), point2d(1, 1), point2d(1, 3), point2d(0, 0)))

 /* Professing test(triangle, a, b, c) */
 query(line(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(triangle(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(vertical(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(horizontal(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(equilateral(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(isosceles(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(right(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(scalene(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(acute(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(obtuse(point2d(2,3), point2d(1,1), point2d(1, 3))).
 
 /* Query Processing */
 writeln(T) :- write(T), nl.
 main:- forall(query(Q), Q-> (writeln(‘yes’)) ; (writeln(‘no’))),
       halt.

```
