**Data Types**
bool => false
string => ""
int int8 int16 int32 int64 => 0
uint uint8 uint16 uint32 uint64 => 0
float32 float64 => 0
complex64 complex128 => 0 + 0i

byte -> alias for uint8
rune -> alias for int32 -> represents a Unicode code point

**Format Specifiers**
%v -> value
%T -> type
%%v -> literal % sign

%t -> boolean

%d -> decimal
%b -> binary
%o -> octal
%x -> hexadecimal

%e -> scientific notation
%f/%F -> floating point
%.xf -> x represents the precision omit for 0 precision
%g -> large notation

%s -> string
%q -> double quoted string

%n@ -> n is width use - for width from left and @ is the format specifier

slices, maps -> passed by ref
arrays, structs -> passed by val

assignment -> creates a shallow copy
copy() -> creates a deep copy

foo bar baz qux quux corge grault garply waldo fred plugh xyzzy thud
