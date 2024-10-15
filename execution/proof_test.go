package execution

import (
	// "bytes"
	"testing"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/holiman/uint256"
)

// Test the VerifyProof function
func TestVerifyProof(t *testing.T) {
	// Define mock proof, root, path, and value
	proof := [][]byte{
		[]byte{0x01, 0x02, 0x03}, // example node data
		[]byte{0x04, 0x05, 0x06}, // another node
	}
	root := keccak256([]byte("root"))  // mocked root hash
	path := []byte{0x01, 0x02}         // mocked path
	value := []byte{0x07, 0x08, 0x09}  // mocked value
	
	// Test valid proof (example scenario)
	valid, err := VerifyProof(proof, root, path, value)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !valid {
		t.Fatalf("Expected proof to be valid, but it was invalid")
	}

	// Test invalid proof (manipulate the proof to cause failure)
	invalidProof := [][]byte{
		[]byte{0x0A, 0x0B, 0x0C}, // wrong node data
	}
	valid, err = VerifyProof(invalidProof, root, path, value)
	if err == nil && valid {
		t.Fatalf("Expected proof to be invalid, but it was valid")
	}
}

// Test the pathsMatch function
func TestPathsMatch(t *testing.T) {
	// Define mock paths
	p1 := []byte{0x01, 0x02, 0x03}
	p2 := []byte{0x01, 0x02, 0x03}
	
	// Test matching paths
	if !pathsMatch(p1, 0, p2, 0) {
		t.Fatalf("Expected paths to match, but they didn't")
	}
	
	// Test non-matching paths
	p3 := []byte{0x04, 0x05, 0x06}
	if pathsMatch(p1, 0, p3, 0) {
		t.Fatalf("Expected paths to not match, but they did")
	}
}

// Test the isEmptyValue function
func TestIsEmptyValue(t *testing.T) {
	// Create a default account RLP encoding
	emptyAccount := Account{
		Nonce:       0,
		Balance:     uint256.NewInt(0).ToBig(),
		StorageHash: [32]byte{0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21},
		CodeHash:    [32]byte{0xc5, 0xd2, 0x46, 0x01, 0x86, 0xf7, 0x23, 0x3c, 0x92, 0x7e, 0x7d, 0xb2, 0xdc, 0xc7, 0x03, 0xc0, 0xe5, 0x00, 0xb6, 0x53, 0xca, 0x82, 0x27, 0x3b, 0x7b, 0xfa, 0xd8, 0x04, 0x5d, 0x85, 0xa4, 0x70},
	}
	encodedEmptyAccount, _ := rlp.EncodeToBytes(emptyAccount)
	
	// Test empty account case
	if !isEmptyValue(encodedEmptyAccount) {
		t.Fatalf("Expected empty account value to return true")
	}
	
	// Test non-empty account case
	nonEmptyValue := []byte{0x01, 0x02, 0x03}
	if isEmptyValue(nonEmptyValue) {
		t.Fatalf("Expected non-empty value to return false")
	}
}

// Test getNibble function
func TestGetNibble(t *testing.T) {
	path := []byte{0x12, 0x34}
	
	// Test for the first nibble
	if n := getNibble(path, 0); n != 0x1 {
		t.Fatalf("Expected nibble 0x1, got %x", n)
	}
	
	// Test for the second nibble
	if n := getNibble(path, 1); n != 0x2 {
		t.Fatalf("Expected nibble 0x2, got %x", n)
	}
}

// Test skipLength function
func TestSkipLength(t *testing.T) {
	// Test different cases for skipLength
	node := []byte{0x00, 0x12}
	if skipLength(node) != 2 {
		t.Fatalf("Expected skip length 2, got %d", skipLength(node))
	}

	node = []byte{0x01, 0x23}
	if skipLength(node) != 1 {
		t.Fatalf("Expected skip length 1, got %d", skipLength(node))
	}
}

// Test the sharedPrefixLength function
func TestSharedPrefixLength(t *testing.T) {
	path1 := []byte{0x12, 0x34}
	path2 := []byte{0x12, 0x35}
	
	// Test case where paths share a prefix
	if sharedPrefixLength(path1, 0, path2) != 3 {
		t.Fatalf("Expected shared prefix length of 3")
	}
	
	// Test case where paths don't share a prefix
	if sharedPrefixLength(path1, 0, []byte{0x56, 0x78}) != 0 {
		t.Fatalf("Expected shared prefix length of 0")
	}
}