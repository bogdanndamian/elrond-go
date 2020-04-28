package mock

import (
	"github.com/ElrondNetwork/elrond-go/data"
)

// TrieStub -
type TrieStub struct {
	GetCalled                func(key []byte) ([]byte, error)
	UpdateCalled             func(key, value []byte) error
	DeleteCalled             func(key []byte) error
	RootCalled               func() ([]byte, error)
	CommitCalled             func() error
	RecreateCalled           func(root []byte) (data.Trie, error)
	CancelPruneCalled        func(rootHash []byte, identifier data.TriePruningIdentifier)
	PruneCalled              func(rootHash []byte, identifier data.TriePruningIdentifier)
	ResetOldHashesCalled     func() [][]byte
	AppendToOldHashesCalled  func([][]byte)
	GetSerializedNodesCalled func([]byte, uint64) ([][]byte, uint64, error)
	DatabaseCalled           func() data.DBWriteCacher
}

// EnterSnapshotMode -
func (ts *TrieStub) EnterSnapshotMode() {
}

// ExitSnapshotMode -
func (ts *TrieStub) ExitSnapshotMode() {
}

// ClosePersister -
func (ts *TrieStub) ClosePersister() error {
	return nil
}

// TakeSnapshot -
func (ts *TrieStub) TakeSnapshot(_ []byte) {
}

// SetCheckpoint -
func (ts *TrieStub) SetCheckpoint(_ []byte) {
}

// GetAllLeaves -
func (ts *TrieStub) GetAllLeaves() (map[string][]byte, error) {
	return make(map[string][]byte), nil
}

// IsPruningEnabled -
func (ts *TrieStub) IsPruningEnabled() bool {
	return false
}

// Get -
func (ts *TrieStub) Get(key []byte) ([]byte, error) {
	if ts.GetCalled != nil {
		return ts.GetCalled(key)
	}

	return nil, errNotImplemented
}

// Update -
func (ts *TrieStub) Update(key, value []byte) error {
	if ts.UpdateCalled != nil {
		return ts.UpdateCalled(key, value)
	}

	return errNotImplemented
}

// Delete -
func (ts *TrieStub) Delete(key []byte) error {
	if ts.DeleteCalled != nil {
		return ts.DeleteCalled(key)
	}

	return errNotImplemented
}

// Root -
func (ts *TrieStub) Root() ([]byte, error) {
	if ts.RootCalled != nil {
		return ts.RootCalled()
	}

	return nil, errNotImplemented
}

// Commit -
func (ts *TrieStub) Commit() error {
	if ts != nil {
		return ts.CommitCalled()
	}

	return errNotImplemented
}

// Recreate -
func (ts *TrieStub) Recreate(root []byte) (data.Trie, error) {
	if ts.RecreateCalled != nil {
		return ts.RecreateCalled(root)
	}

	return nil, errNotImplemented
}

// String -
func (ts *TrieStub) String() string {
	return "stub trie"
}

// IsInterfaceNil returns true if there is no value under the interface
func (ts *TrieStub) IsInterfaceNil() bool {
	return ts == nil
}

// CancelPrune invalidates the hashes that correspond to the given root hash from the eviction waiting list
func (ts *TrieStub) CancelPrune(rootHash []byte, identifier data.TriePruningIdentifier) {
	if ts.CancelPruneCalled != nil {
		ts.CancelPruneCalled(rootHash, identifier)
	}
}

// Prune removes from the database all the old hashes that correspond to the given root hash
func (ts *TrieStub) Prune(rootHash []byte, identifier data.TriePruningIdentifier) {
	if ts.PruneCalled != nil {
		ts.PruneCalled(rootHash, identifier)
	}
}

// ResetOldHashes resets the oldHashes and oldRoot variables and returns the old hashes
func (ts *TrieStub) ResetOldHashes() [][]byte {
	if ts.ResetOldHashesCalled != nil {
		return ts.ResetOldHashesCalled()
	}

	return nil
}

// AppendToOldHashes appends the given hashes to the trie's oldHashes variable
func (ts *TrieStub) AppendToOldHashes(hashes [][]byte) {
	if ts.AppendToOldHashesCalled != nil {
		ts.AppendToOldHashesCalled(hashes)
	}
}

// GetSerializedNodes -
func (ts *TrieStub) GetSerializedNodes(hash []byte, maxBuffToSend uint64) ([][]byte, uint64, error) {
	if ts.GetSerializedNodesCalled != nil {
		return ts.GetSerializedNodesCalled(hash, maxBuffToSend)
	}
	return nil, 0, nil
}

// Database -
func (ts *TrieStub) Database() data.DBWriteCacher {
	if ts.DatabaseCalled != nil {
		return ts.DatabaseCalled()
	}
	return nil
}

// GetDirtyHashes -
func (ts *TrieStub) GetDirtyHashes() (data.ModifiedHashes, error) {
	return nil, nil
}

// AddNewHashes -
func (ts *TrieStub) AddNewHashes(_ data.ModifiedHashes) {
}
