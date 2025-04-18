# Go_string_utils

Some functions for string and number convertions and file manipulation

## `Int32ToString()`

### Usage

`Int32ToString(x *int32) string`

### Example

Code: 

`"fmt"` and `"reflect"` imported

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

`"fmt"` imported

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

### Usage

```
file1 := "file1.txt"
file2 := "file2.txt"
sep := " | "
fast_diff(&file1, &file2, &sep)
```

Output

```
> oui |
non | non
>  |
> P |
ll | ll
>  |
K | K
 | > oui
 | >
```


