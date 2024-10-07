package main

import (
  "fmt"
  "log"
  "encoding/hex"
)

func fixed_xor(str1, str2 string) (string, error){
  // Decode the hex values
  buf1, err := hex.DecodeString(str1)
  if err != nil {
    return "", fmt.Errorf("Decoder error on line 1")
  }

  buf2, err := hex.DecodeString(str2)
  if err != nil {
    return "", fmt.Errorf("Decoder error on hex 2")
  }

  // if the lengths arent even its invalid
  if len(buf1) != len(buf2) {
    return "", fmt.Errorf("Strings must be of equal length")
  }

  // make a byte array to store the result
  result := make([]byte, len(buf1))

  // perform xor operations
  for i := range buf1 {
    result[i] = buf1[i] ^ buf2[i]
  }
  return string(result), nil
}

func main(){
  buf1 := "1c0111001f010100061a024b53535009181c"
  buf2 := "686974207468652062756c6c277320657965" 

  fmt.Println(" Testing buffer #1:\n", buf1, "\n")
  fmt.Println(" Testing Buffer #2:\n", buf2, "\n")

  result, err := fixed_xor(buf1, buf2);

  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf(" XOR result (hex): \n %x\n", result)
}
