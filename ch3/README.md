Go's type categories:

1 - Basic types: numbers, strings, and booleans

Integers:
 - int8, int16, int32, int64
 - uint8, uint16, uint32, uint64

 - int (32 or 64 bits, compiler choice)
 - uint (32 or 64 bits, compiler choice)

 - rune (int32 synonym - unicode code point)
 - byte (uint8 synonym - piece of raw data)
 - uintptr (width not specified, hold all the bits of a pointer)

Go's binary operators for arithmetic, logic, and comparison (decreasing precedence):

```
*   /   %   <<   >>   &   &^
+   -   |   ^
==  !=  <   <=   >    >=
&&
||
```

Binary comparison operators:

```
==  equal to
!=  not equal to
<   less than
<=  less than or equal to
>   greater than
>=  greater than or equal to
```

Note: basic types (booleans, numbers, and strings) are *comparable* (== and !=).
Note: integers, floating-point numbers, and strings are *ordered*.

Addition and subtraction operators:

```
+   unary positive (no effect)
-   unary negation
```

Bitwise operators

```
&   bitwise AND
|   bitwise OR
^   bitwise XOR
&^  bit clear (AND NOT)
<<  left shift
>>  right shift
```

2 - Aggregate types: arrays and structs
3 - Reference types: pointers, slices, maps, functions, and channels
4 - Interface types: ???