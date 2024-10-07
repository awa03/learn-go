package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func scoreText(text []byte) float64 {
  // Score based on frequency in alphabet
	frequency := map[byte]float64{
		' ': 13.0, 'E': 12.7, 'T': 9.1, 'A': 8.2, 'O': 7.5, 'I': 7.0, 'N': 6.7, 'S': 6.3, 'H': 6.1, 'R': 6.0,
		'D': 4.3, 'L': 4.0, 'U': 2.8, 'C': 2.8, 'M': 2.4, 'F': 2.2, 'Y': 2.0, 'W': 2.0, 'G': 2.0, 'P': 1.9,
		'B': 1.5, 'V': 1.0, 'K': 0.8, 'X': 0.1, 'Q': 0.1, 'J': 0.1, 'Z': 0.1,
	}
	score := 0.0
	for _, b := range text {
		// Treat both lower and upper cases as the same for simplicity
		if b >= 'a' && b <= 'z' {
			b -= 32 // Convert to uppercase
		}
		if val, ok := frequency[b]; ok {
			score += val
		}
	}
	return score
}

func xorWithKey(buf [] byte, key byte) []byte {
  result := make([]byte, len(buf))
  // xor each byte
  for i:=range buf {
    result[i] = buf[i] ^ key
  }
  return result
}

func findBestKey(hexStr string) (byte, string, float64) {
  bytes, err := hex.DecodeString(hexStr)

  if err != nil {
    log.Fatal(err)
  } 

  var bestKey byte
	var bestScore float64
	var bestResult []byte

  for key := 0 ; key < 256 ; key++ {
    result := xorWithKey(bytes, byte(key))
    score  := scoreText(result)

    if score > bestScore {
      bestScore = score
      bestKey = byte(key)
      bestResult = result
    }
  }
  return bestKey, string(bestResult), bestScore
}


func main(){
  cypher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

  fmt.Println("Hex Encoded String:\n", cypher, "\n")

  bestKey, bestResult, bestScore := findBestKey(cypher)

  fmt.Println("Best Key:", bestKey)
  fmt.Println("Best Result:", bestResult)
  fmt.Println("Best Score:", bestScore)

}
