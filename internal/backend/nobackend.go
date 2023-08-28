// Generated code. DO NOT EDIT.

package backend

import (
	"crypto"
	"crypto/cipher"
	"hash"
	"io"
)

const Enabled = false

func NewSHA1() hash.Hash     { panic("cryptobackend: not available") }
func NewSHA224() hash.Hash   { panic("cryptobackend: not available") }
func NewSHA256() hash.Hash   { panic("cryptobackend: not available") }
func NewSHA384() hash.Hash   { panic("cryptobackend: not available") }
func NewSHA512() hash.Hash   { panic("cryptobackend: not available") }
func NewSHA3_224() hash.Hash { panic("cryptobackend: not available") }
func NewSHA3_256() hash.Hash { panic("cryptobackend: not available") }
func NewSHA3_384() hash.Hash { panic("cryptobackend: not available") }
func NewSHA3_512() hash.Hash { panic("cryptobackend: not available") }

func SHA1(p []byte) (sum [20]byte)     { panic("cryptobackend: not available") }
func SHA224(p []byte) (sum [28]byte)   { panic("cryptobackend: not available") }
func SHA256(p []byte) (sum [32]byte)   { panic("cryptobackend: not available") }
func SHA384(p []byte) (sum [48]byte)   { panic("cryptobackend: not available") }
func SHA512(p []byte) (sum [64]byte)   { panic("cryptobackend: not available") }
func SHA3_224(p []byte) (sum [28]byte) { panic("cryptobackend: not available") }
func SHA3_256(p []byte) (sum [32]byte) { panic("cryptobackend: not available") }
func SHA3_384(p []byte) (sum [48]byte) { panic("cryptobackend: not available") }
func SHA3_512(p []byte) (sum [64]byte) { panic("cryptobackend: not available") }

func SupportsHash(h crypto.Hash) bool { panic("cryptobackend: not available") }

func NewHMAC(h func() hash.Hash, key []byte) hash.Hash { panic("cryptobackend: not available") }

func NewAESCipher(key []byte) (cipher.Block, error) { panic("cryptobackend: not available") }
func NewGCMTLS(c cipher.Block) (cipher.AEAD, error) { panic("cryptobackend: not available") }

func ExpandHKDF(h func() hash.Hash, pseudorandomKey, info []byte) (io.Reader, error) {
	panic("cryptobackend: not available")
}

func ExtractHKDF(h func() hash.Hash, secret, salt []byte) ([]byte, error) {
	panic("cryptobackend: not available")
}
