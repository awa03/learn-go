package main
import (
  "fmt"

  // For hashing
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "io"
  "log"
  "net/http"
  "os"
  "time"
  "github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Block struct {
  Index     int     // position of the data recorded
  Timestamp string  // automatically determined
  BPM       int     // beats per minute, heart rate
  Hash      string  // identifier for record
  PrevHash  string  // previous record in chain
}

// Through using hashes we are saving data because we only 
// need to store that information. By storing previous hashes
// we can ensure the integrity of the blockchain. We can ensure
// that the data is the same through the hashes. 

func calculateHash(block Block) string {
  record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
  h := sha256.New()
  h.Write([]byte(record))
  hashed := h.Sum(nil)
  return hex.EncodeToString(hashed)
}

// Implementing a linked list ot connect the blocks
// Within the generate block function we are initializing 
// a new block, and adding current block as the previous block
// element of the new block. 

func generateBlock(oldBlock Block, BPM int) (Block, error) {
  var newBlock Block
  t := time.Now()

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String()
  newBlock.BPM = BPM
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Hash = calculateHash(newBlock)
  
  return newBlock, nil
}

// this function will check the validity of a block. We need to ensure that
// a block is not a duplicate of the previous hash, and that the index is 
// valid

func isBlockValid(newBlock, oldBlock Block) bool {
  // Ensure indexing is consistent and valid
  if oldBlock.Index + 1 != newBlock.Index {
    return false
  }
  // ensure that the old has isnt the same as the new blocks previous
  if oldBlock.Hash != newBlock.PrevHash {
    return false
  }
  // ensure the hash is not the same
  if calculateHash(newBlock) != newBlock.Hash {
    return false
  }
  return true
}

func replaceChain(newBlock []Block) {
  if len(newBlock) > len(BlockChain) {
    BlockChain = newBlock
  }
}

// our blockchain
var BlockChain []Block

func main(){
   
}
