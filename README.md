# Go_string_utils

Some functions for string and number convertions

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

"`fmt`" imported

```
var x string = "234"
rtn_val := StringToInt32(x)
fmt.Println(rtn_val)
```

Results:

```
234
```


