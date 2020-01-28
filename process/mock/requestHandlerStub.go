package mock

type RequestHandlerStub struct {
	RequestShardHeaderCalled        func(shardID uint32, hash []byte)
	RequestMetaHeaderCalled         func(hash []byte)
	RequestMetaHeaderByNonceCalled  func(nonce uint64)
	RequestShardHeaderByNonceCalled func(shardID uint32, nonce uint64)
	RequestTransactionHandlerCalled func(destShardID uint32, txHashes [][]byte)
	RequestScrHandlerCalled         func(destShardID uint32, txHashes [][]byte)
	RequestRewardTxHandlerCalled    func(destShardID uint32, txHashes [][]byte)
	RequestMiniBlockHandlerCalled   func(destShardID uint32, miniblockHash []byte)
	RequestMiniBlocksHandlerCalled  func(destShardID uint32, miniblocksHashes [][]byte)
	RequestTrieNodesCalled          func(destShardID uint32, hash []byte, topic string)
}

func (rhs *RequestHandlerStub) SetEpoch(epoch uint32) {
}

func (rhs *RequestHandlerStub) RequestShardHeader(shardID uint32, hash []byte) {
	if rhs.RequestShardHeaderCalled == nil {
		return
	}
	rhs.RequestShardHeaderCalled(shardID, hash)
}

func (rhs *RequestHandlerStub) RequestMetaHeader(hash []byte) {
	if rhs.RequestMetaHeaderCalled == nil {
		return
	}
	rhs.RequestMetaHeaderCalled(hash)
}

func (rhs *RequestHandlerStub) RequestMetaHeaderByNonce(nonce uint64) {
	if rhs.RequestMetaHeaderByNonceCalled == nil {
		return
	}
	rhs.RequestMetaHeaderByNonceCalled(nonce)
}

func (rhs *RequestHandlerStub) RequestShardHeaderByNonce(shardID uint32, nonce uint64) {
	if rhs.RequestShardHeaderByNonceCalled == nil {
		return
	}
	rhs.RequestShardHeaderByNonceCalled(shardID, nonce)
}

func (rhs *RequestHandlerStub) RequestTransaction(destShardID uint32, txHashes [][]byte) {
	if rhs.RequestTransactionHandlerCalled == nil {
		return
	}
	rhs.RequestTransactionHandlerCalled(destShardID, txHashes)
}

func (rhs *RequestHandlerStub) RequestUnsignedTransactions(destShardID uint32, txHashes [][]byte) {
	if rhs.RequestScrHandlerCalled == nil {
		return
	}
	rhs.RequestScrHandlerCalled(destShardID, txHashes)
}

func (rhs *RequestHandlerStub) RequestRewardTransactions(destShardID uint32, txHashes [][]byte) {
	if rhs.RequestRewardTxHandlerCalled == nil {
		return
	}
	rhs.RequestRewardTxHandlerCalled(destShardID, txHashes)
}

func (rhs *RequestHandlerStub) RequestMiniBlock(destShardID uint32, miniblockHash []byte) {
	if rhs.RequestMiniBlockHandlerCalled == nil {
		return
	}
	rhs.RequestMiniBlockHandlerCalled(destShardID, miniblockHash)
}

func (rhs *RequestHandlerStub) RequestMiniBlocks(destShardID uint32, miniblocksHashes [][]byte) {
	if rhs.RequestMiniBlocksHandlerCalled == nil {
		return
	}
	rhs.RequestMiniBlocksHandlerCalled(destShardID, miniblocksHashes)
}

func (rhs *RequestHandlerStub) RequestTrieNodes(destShardID uint32, miniblockHash []byte, topic string) {
	if rhs.RequestTrieNodesCalled == nil {
		return
	}
	rhs.RequestTrieNodesCalled(destShardID, miniblockHash, topic)
}

// IsInterfaceNil returns true if there is no value under the interface
func (rhs *RequestHandlerStub) IsInterfaceNil() bool {
	return rhs == nil
}
