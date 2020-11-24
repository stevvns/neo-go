package state

import (
	"testing"

	"github.com/nspcc-dev/neo-go/internal/testserdes"
)

func TestEncodeDecodeStorageItem(t *testing.T) {
	storageItem := &StorageItem{
		Value:   []byte{},
		IsConst: false,
	}

	testserdes.EncodeDecodeBinary(t, storageItem, new(StorageItem))
}
