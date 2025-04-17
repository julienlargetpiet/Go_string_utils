package Go_string_utils

func Int32ToString(x *int32) string {
  const base int32 = 10
  var remainder int32
  rtn_str := ""
  for *x > 0 {
    remainder = *x % base
    rtn_str = string(remainder + 48) + rtn_str
    *x -= remainder
    *x /= 10
  }
  return rtn_str
}


