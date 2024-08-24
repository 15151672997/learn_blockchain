package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	//1区块高度
	Height int64
	//2.上一个区块哈希
	PrevBlockHash []byte
	//3.交易数据
	Data []byte
	//4.时间戳
	Timestamp int64
	//5.hash
	Hash []byte
}

// 生成hash
func (b *Block) SetHash() {
	//1.height []byte
	//2. 时间戳 []byte
	//3.拼接所有属性
	//4.生成hash

	heightByter := IntToHex(b.Height)
	timeStr := strconv.FormatInt(b.Timestamp, 2) //2表示将时间戳转换为2进制
	timeStamp := []byte(timeStr)

	//拼接数组
	blockBytes := bytes.Join([][]byte{heightByter, b.PrevBlockHash, b.Data, timeStamp, b.Hash}, []byte{})
	//生成hash
	hash := sha256.Sum256(blockBytes)

	b.Hash = hash[:]

} // 1创建新的区块方法
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	blokc := Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
	}

	//hash是由区块中的数据产生的数据
	blokc.SetHash()
	return &blokc

}

// 2生成创世区块
func CreateGenesisBlock(data string) *Block {

	//高度//prevhas可知

	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0})

}
