package chain

import (
	"gopherchain/block"
)

type Chain struct {
	blocks []block.Block
}

func (chain *Chain) Peek() block.Block {
	return chain.blocks[len(chain.blocks) - 1]
}

func (chain *Chain) Add(block *block.Block) {
	chain.blocks = append(chain.blocks, *block)
}

func (chain *Chain) Blocks() []block.Block {
	blocksClone := []block.Block{}
	copy(blocksClone, chain.blocks)
	return chain.blocks
}