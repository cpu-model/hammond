package registration

import (
 "encoding/csv"
 "fmt"
 "io"
 "strconv"
 "strings"
)

// ParseRoutingCSV reads an explicitly transcribed wiring table.
// Header: key,drawbar,wheel. Key is 1..61; drawbar is 1..9;
// wheel is 1..91. Exactly 549 contacts must be specified.
// Wheel 0 is permitted only for an explicitly disconnected contact.
// Requiring all contacts prevents silently inventing foldback.
func ParseRoutingCSV(reader io.Reader) (Routing,error) {
 var routing Routing
 csvReader:=csv.NewReader(reader)
 csvReader.FieldsPerRecord=3
 header,err:=csvReader.Read()
 if err!=nil { return routing,fmt.Errorf("read routing header: %w",err) }
 if strings.TrimSpace(header[0])!="key" || strings.TrimSpace(header[1])!="drawbar" || strings.TrimSpace(header[2])!="wheel" {
  return routing,fmt.Errorf("expected header key,drawbar,wheel")
 }
 var seen [Keys][Harmonics]bool
 count:=0
 for {
  record,err:=csvReader.Read()
  if err==io.EOF { break }
  if err!=nil { return Routing{},fmt.Errorf("routing row %d: %w",count+2,err) }
  var nums [3]int
  for i,v:=range record {
   nums[i],err=strconv.Atoi(strings.TrimSpace(v))
   if err!=nil { return Routing{},fmt.Errorf("routing row %d: invalid integer %q",count+2,v) }
  }
  k,d,w:=nums[0],nums[1],nums[2]
  if k<1 || k>Keys || d<1 || d>Harmonics || w<0 || w>91 {
   return Routing{},fmt.Errorf("routing row %d: contact out of range",count+2)
  }
  if seen[k-1][d-1] { return Routing{},fmt.Errorf("duplicate contact key %d drawbar %d",k,d) }
  seen[k-1][d-1]=true
  routing[k-1][d-1]=uint8(w)
  count++
 }
 if count!=Keys*Harmonics {
  return Routing{},fmt.Errorf("incomplete routing: %d of %d contacts",count,Keys*Harmonics)
 }
 return routing,nil
}
