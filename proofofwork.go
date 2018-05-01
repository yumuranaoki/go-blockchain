package main

import (
	"bytes"
	"math/big"
  "fmt"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

//ProofOfWorkの難易度を決める
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//targetの右に数を追加。targetBitsが大きいほど、targetは小さくなる
	target.Lsh(target, uint(256-targetBits))
}

//dataにnonceを追加して、sha256前のデータを準備
//blockのメソッドでもいいのでは？
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
  var hashInt Big.Int
  var hash [32]byte
  nonce := 0

  fmt,Printf("Mining the block containing \"%s\"\n", pow.block.Data)
  for nonce < maxNonce {
    data := pow.prepareData(nonce)
    hash = sha256.Sum256(data)
    fmt.Printf("\r%x", hash)
    hashInt.SetBytes(hash[:])

    if hashInt.Cmp(pow.target) == -1 {
        break
    } else {
        nonce++
    }

  }
  fmt.Print("\n\n")

  return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {

}
