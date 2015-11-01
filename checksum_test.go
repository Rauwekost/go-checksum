package checksum

import (
	"crypto"
	"testing"
)

const testString = "The quick brown fox jumps over the lazy dog"
const testFile = "examples/testfile.txt"

func TestHashes(t *testing.T) {
	hash, err := String(testString, crypto.MD5)
	if err != nil || hash != "9e107d9d372bb6826bd81d3542a419d6" {
		t.Error("MD5 string test failed")
	}
	hash, err = String(testString, crypto.SHA1)
	if err != nil || hash != "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12" {
		t.Error("SHA1 string test failed")
	}
	hash, err = String(testString, crypto.SHA224)
	if err != nil || hash != "730e109bd7a8a32b1cb9d9a09aa2325d2430587ddbc0c38bad911525" {
		t.Error("SHA224 string test failed")
	}
	hash, err = String(testString, crypto.SHA256)
	if err != nil || hash != "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592" {
		t.Error("SHA256 string test failed")
	}
	hash, err = String(testString, crypto.SHA384)
	if err != nil || hash != "ca737f1014a48f4c0b6dd43cb177b0afd9e5169367544c494011e3317dbf9a509cb1e5dc1e85a941bbee3d7f2afbc9b1" {
		t.Error("SHA284 string test failed")
	}
	hash, err = String(testString, crypto.SHA512)
	if err != nil || hash != "07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6" {
		t.Error("SHA512 string test failed")
	}
}

func TestFile(t *testing.T) {
	hash, err := File(testFile, crypto.MD5)
	if err != nil || hash != "4b068eabe53d384de505733b7346ac15" {
		t.Error("MD5 file test failed")
	}
	hash, err = File(testFile, crypto.SHA1)
	if err != nil || hash != "13bd73dbb1f0f31118b18ca39f182bb78862e7f9" {
		t.Error("SHA1 file test failed")
	}
	hash, err = File(testFile, crypto.SHA224)
	if err != nil || hash != "c3200804f8ed976eb79c0201236370e780f684f39abbad5da235ab41" {
		t.Error("SHA224 file test failed")
	}
	hash, err = File(testFile, crypto.SHA256)
	if err != nil || hash != "5312b2e55205aadc4e315aaf212a1a9781dcd5ef51a28ccb4641715fd0a58cff" {
		t.Error("SHA256 file test failed")
	}
	hash, err = File(testFile, crypto.SHA384)
	if err != nil || hash != "8901944bd4a7eba4da1aa0e0762e8dc49cff894885d3b222d993c8d41f7c5ae43fcbe490d1779c30664f9b85351660b5" {
		t.Error("SHA384 file test failed")
	}
	hash, err = File(testFile, crypto.SHA512)
	if err != nil || hash != "960ef1c3cd6e11f84f1a68e2bfb1083c7eba53b88c2fc2e4a6961b71830ef9fe61b3d5bd93636748d7e9084b3dc6ce54307f60e5f1bf07fe94e86ae7eddf0820" {
		t.Error("SHA512 file test failed")
	}
}
