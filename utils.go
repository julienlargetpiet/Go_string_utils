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

func StringToInt32(x string) int32 {
  var ref_nb = [10]uint8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
  var rtn_val int32 = 0
  var lngth int = len(x)
  var i2 int32
  var cur_rn uint8
  var i int
  for i = 0; i + 1 < lngth; i++ {
    cur_rn = x[i]
    i2 = 0
    for cur_rn != ref_nb[i2]{
      i2++
    }
    rtn_val += i2
    rtn_val *= 10
  }
  cur_rn = x[i]
  i2 = 0
  for cur_rn != ref_nb[i2]{
    i2++
  }
  rtn_val += i2
  return rtn_val
}


