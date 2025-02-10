package kata

import "unicode"
import "strings"


func CamelCase(s string) string {
  fieldsSLICE := strings.Fields(s)
  for index,value := range fieldsSLICE{
    w := []rune(value)
    if len(w) != 0{
      w[0] = unicode.ToUpper(w[0])
    }
    for j:= 1; j < len(w);j ++{
      w[j] = unicode.ToLower(w[j])
      
    }
    fieldsSLICE[index] = string(w)
    
    
    
  }
  res := ""
  for _,k := range fieldsSLICE{
    res += string(k)
    
  }
  return res
