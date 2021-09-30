# Building a toy-chain

As the name "Blockchain" suggest, its just a data structure of _chained_ blocks. Each block contains some data we want to pass around in a p2p network, as well as a hash that is associated with the block itself. We can chain the blocks together by  storing the blocks previous block hash, similar to a linked list. That is to say that each previous hash in our blockchain represents the last block created up to the current block.

```go=
type Block struct {
	hash     []byte
	data     []byte
	prevHash []byte
}
```

Now that we have a block, we need a way to create a new hash for the current block. We can compose with a method that hashes the blocks current data and its previous hash. Let's use `sha256` for now. 

```go=
func (b *Block) Hash() {
	info := bytes.Join([][]byte{b.data, b.prev}, []byte{})
	hash := sha256.Sum256(info)
	b.hash = hash[:]
}
```

Cool. Notice that the info we're hashing is a 2D slice of the blocks data and it's previous hash.

Now we can create a new structure that represents our chain of blocks. We'll keep it as simple as possible for now, so a simple 1D array of Blocks will suffice. 

```go=
type BlockChain struct {
	blocks []*Block
}
```

Aside from a few helper methods to create and add blocks to the chain, that is pretty much what a block chain is. Simple, yet effective!

Take a look at the repo if you'd like to take a look at a runnable blockchain. 

In the future - let's look into adding merkle trees!
