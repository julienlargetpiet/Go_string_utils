# Go_string_utils

Some functions for string and number convertions and file manipulation

## `Int32ToString()`

### Usage

`Int32ToString(x *int32) string`

### Example

Code: 

```
var x int32 = 234
var rtn_str string
rtn_str = Int32ToString(&x)
fmt.Println(rtn_str)
fmt.Println(reflect.TypeOf(rtn_str))
```

Result:

```
234
string
```

## StringToInt32

### Usage

Code: 

```
var x string = "234"
rtn_val := StringToInt32(x)
fmt.Println(rtn_val)
```

Results:

```
234
```

## fast_diff

Special implementation of the `diff` command
The fisrt file must be the one that have been updated.

### Usage

Code:

```
file1 := "file1.txt"
file2 := "file2.txt"
sep := " | "
fast_diff(&file1, &file2, &sep)
```

Output

```
oui | -
non | non
 | -
P | -
ll | ll
 | -
K | K
 | + oui
 | + pp
 | + mm
 | + mm
```

Code:

```
file1 := "file1b.txt"
file2 := "file2b.txt"
sep := " | "
x := fast_diff2(&file1, &file2, &sep)
fmt.Println(x)
```

Output:

```
oui | -
non | non
 | -
ll | ll
ll | -
K | -
```

## fast_diff2

Same thing than `fast_diff` but returns a `string`

### Usage

Code:

```
file1 := "file1.txt"
file2 := "file2.txt"
sep := " | "
x := fast_diff2(&file1, &file2, &sep)
fmt.Println(x)
```

Output

```
oui | -
non | non
 | -
P | -
ll | ll
 | -
K | K
 | + oui
 | + pp
 | + mm
 | + mm
```

Code:

```
file1 := "file1b.txt"
file2 := "file2b.txt"
sep := " | "
x := fast_diff2(&file1, &file2, &sep)
fmt.Println(x)
```

Output:

```
oui | -
non | non
 | -
ll | ll
ll | -
K | -
```
