package main

// Needed imports
import (
  "fmt"
  "encoding/hex"
  "encoding/base64"
)

func convert_hex_to_base_64(test_str string) (string, error) {
  // decode the hex values into bytes (groups of 8)
  bytes, err := hex.DecodeString(test_str)
  if err != nil {
    return "", err
  }

  base64Str := base64.StdEncoding.EncodeToString(bytes)
  // Convert those bytes to base 64, using encoding.
  return base64Str, nil
}

func main(){
  // Our Hex
  var str string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
  base64Str, err := convert_hex_to_base_64(str)

  // If theres an error
  if err != nil {
    fmt.Println("Uh Oh!")
  }

  fmt.Println(base64Str)
}
