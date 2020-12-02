# Daily Coding Problem: Problem #718 [Medium]  

This problem was asked by Apple.

Gray code is a binary code where each successive value differ in only one bit,
as well as when wrapping around.
Gray code is common in hardware so that we don't see temporary spurious
values during transitions.

Given a number of bits n, generate a possible gray code for it.

For example, for n = 2, one gray code would be [00, 01, 11, 10].

## Analysis

I went with the classic Gray Code algorithm:

1. Reflect the current bits, giving twice as many values, albeit 2 of each.
2. Add a new bit to each value, 0-value bits to the top half,
1-value bits to the bottom, the part reflected in step 1.

The first few steps go like this:

Current bits:
```
0
1
```

Reflect these bits:
```
0
1
1
0
```

Add a bit prefix to each value, 0-bits on first half,
1-bits on second half:
```
00
01
11
10
```

The trick in this algorithm is the ability to find "half way"
through the array of values:
Computer integer arithmetic rounds down, so 1/2 evaluates to 0.
I used two indexes, one starting at zero and incrementing,
one starting at 2^m and decrementing.
Looping stops when the decrementing index gets to less than or
equal to the incrementing index.

I tried this with 2 data structures:

1. [one Go byte](gray1.go) per conceptual bit
2. [one Go integer bit](gray2.go) per conceptual bit

## Interview Analysis

This problem seems difficult for a "medium" rating.
Generating Gray Codes via algorithms isn't something that a lot
of people are going to know off the top of their heads.
I used what I remember as "Gray Code Classic",
but I didn't google for an algorithm,
and I know there's several.
Getting past that is the first hurdle.

A candidate can get hung up on the integer arithmetic rounding problem.
An algorithm could look correct to a human,
but give incomprehensible output based on rounding 1/2 or 4/8 or 32/64 to zero.
I think this problem exists regardless of the data structure chosen.
This is also one of those mistakes that even experienced developers make.

Choice of data structure could affect how easy the code
is to write: I tried a slice-of-slices,
with a Go `byte` holding each conceptual bit value,
and a slice of unsigned integers,
with each bit in an unsigned integer for a conceptual bit.
There's not a lot of difference to the variants.

Aside from generation algorithm,
and a facet of computer integer arithmetic that everyone
forgets and then has trouble with,
this isn't a bad problem if the interviewer wants to see
the candidate write some code.
There's an algorithm, a choice of data structure,
looping, maybe with multiple indexes,
and possibly bitwise operations.
There's enough variables involved that the candidate's
skill in naming gets exercized.

If the interviewer is willing to give away a Gray Code
generation algorithm,
this could be a really good coding problem for a mid-
or senior-level developer.
