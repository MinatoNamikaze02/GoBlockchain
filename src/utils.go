package blockchain


func FindPOFString(difficulty int) string{
	str := ""
	for i := 0; i < difficulty; i++ {
		str += "0"
	}
	return str
}

func IsValidProof(b *Blockchain,block Block, blockHash string) bool {
	return (blockHash[:b.Difficulty] == FindPOFString(b.Difficulty) /*&& blockHash == block.CalculateHash()*/)
}
