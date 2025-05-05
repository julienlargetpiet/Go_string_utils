package Go_string_utils

import (
  "fmt"
  "os"
  "bufio"
)

func Int32ToString(x *int32) string {
  if *x == 0 {
    return "0"
  }
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

func fast_diff(file1 *string, file2 *string, sep *string) {
  var dataa string
  var datab string
  var comp bool = true
  filea, err := os.Open(*file1) 
  if err != nil {
    return
  }
  defer filea.Close()
  scannera := bufio.NewScanner(filea)
  fileb, err := os.Open(*file2) 
  if err != nil {
    return
  }
  defer fileb.Close()
  scannerb := bufio.NewScanner(fileb)
  for comp && scannerb.Scan() {
    datab = scannerb.Text()
    scannera.Scan()
    dataa = scannera.Text()
    for datab != dataa && scannera.Scan() {
      fmt.Printf("%v%v -\n", dataa, *sep)
      dataa = scannera.Text()
    }
    comp = (datab == dataa)
    if comp {
      fmt.Printf("%v%v%v\n", dataa, *sep, datab)
    } else {
      fmt.Printf("%v+ %v\n", *sep, datab)
    }
  }
  for scannerb.Scan() {
    datab = scannerb.Text()
    fmt.Printf("%v+ %v\n", *sep, datab)
  }
  if comp {
    for scannera.Scan() {
      dataa = scannera.Text()
      rtn_str = rtn_str + dataa + *sep + "- " + "\n"
    }
  }
}

func fast_diff2(file1 *string, file2 *string, sep *string) string {
  var dataa string
  var datab string
  var comp bool = true
  var rtn_str string = ""
  filea, err := os.Open(*file1) 
  if err != nil {
    return ""
  }
  defer filea.Close()
  scannera := bufio.NewScanner(filea)
  fileb, err := os.Open(*file2) 
  if err != nil {
    return ""
  }
  defer fileb.Close()
  scannerb := bufio.NewScanner(fileb)
  for comp && scannerb.Scan() {
    datab = scannerb.Text()
    scannera.Scan()
    dataa = scannera.Text()
    for datab != dataa && scannera.Scan() {
      rtn_str = rtn_str + dataa + *sep + "- \n"
      dataa = scannera.Text()
    }
    comp = (datab == dataa)
    if comp {
      rtn_str = rtn_str + dataa + *sep + datab + "\n"
    } else {
      rtn_str = rtn_str + *sep + "+ " + datab + "\n"
    }
  }
  for scannerb.Scan() {
    datab = scannerb.Text()
    rtn_str = rtn_str + *sep + "+ " + datab + "\n"
  }
  if comp {
    for scannera.Scan() {
      dataa = scannera.Text()
      rtn_str = rtn_str + dataa + *sep + "- " + "\n"
    }
  }
  return rtn_str
}

