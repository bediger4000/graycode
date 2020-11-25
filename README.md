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
1-value bits to the bottom half, which was reflected in step 1.

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
Getting past that is the first hurdle.

Choice of data structure could affect how easy the code
is to write: I tried a slice-of-slices,
with a Go `byte` holding each conceptual bit vale,
and a slice of unsigned integers,
with each bit in an unsigned integer for a conceptual bit.
There's not a lot of difference to the variants.

A candidate can get hung up on the integer arithmetic rounding problem.
An algorithm could look correct to a human,
but give incomprehensible output based on rounding 1/2 or 4/8 or 32/64 to zero.
I think this problem exists regardless of the data structure chosen.
