package wallet

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"log"
)

// refactor this to
//
//	https://pkg.go.dev/github.com/algorand/go-algorand-sdk/mnemonic
const (
	checkSumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

func (w *Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)
	versionedHash := append([]byte{version}, pubHash...)
	checkSum := Checksum(versionedHash)
	fullHash := append(versionedHash, checkSum...)
	address := Base58Encode(fullHash)
	//	fmt.Printf("Pubkey: %x\n", w.PublicKey)
	//	fmt.Printf("pub hash %x\n", pubHash)
	//	fmt.Printf("address %x\n", address)
	return address

}
func ValidateAddress(address string) bool {
	publicKeyHash := Base56Decode([]byte(address))
	actualCheckSum := publicKeyHash[len(publicKeyHash)-checkSumLength:]
	version := publicKeyHash[0]

	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-checkSumLength]
	targetChecksum := Checksum(append([]byte{version}, publicKeyHash...))

	return bytes.Equal(actualCheckSum, targetChecksum)

}

func NewKeyPair() (ed25519.PrivateKey, ed25519.PublicKey) {

	public, private, err := ed25519.GenerateKey(rand.Reader)

	if err != nil {
		log.Panic(err)
	}

	return private, public
}
func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}
func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)
	hasher := sha256.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}
	publicRipeMD := hasher.Sum(nil)
	return publicRipeMD

}
func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:checkSumLength]
}
