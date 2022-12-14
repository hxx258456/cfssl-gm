// Package derhelpers implements common functionality
// on DER encoded data
package derhelpers

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"

	"github.com/hxx258456/ccgo/sm2"
	"github.com/hxx258456/ccgo/x509"

	cferr "github.com/hxx258456/cfssl-gm/errors"
)

// ParsePrivateKeyDER parses a PKCS #1, PKCS #8, ECDSA, or Ed25519 DER-encoded
// private key. The key must not be in PEM format.
func ParsePrivateKeyDER(keyDER []byte) (key crypto.Signer, err error) {
	generalKey, err := x509.ParsePKCS8PrivateKey(keyDER)
	if err != nil {
		generalKey, err = x509.ParsePKCS1PrivateKey(keyDER)
		if err != nil {
			generalKey, err = x509.ParseECPrivateKey(keyDER)
			if err != nil {
				generalKey, err = ParseEd25519PrivateKey(keyDER)
				if err != nil {
					// We don't include the actual error into
					// the final error. The reason might be
					// we don't want to leak any info about
					// the private key.
					return nil, cferr.New(cferr.PrivateKeyError,
						cferr.ParseFailed)
				}
			}
		}
	}

	switch generalKey := generalKey.(type) {
	case *sm2.PrivateKey:
		return generalKey, nil
	case *rsa.PrivateKey:
		return generalKey, nil
	case *ecdsa.PrivateKey:
		return generalKey, nil
	case ed25519.PrivateKey:
		return generalKey, nil
	}

	// should never reach here
	return nil, cferr.New(cferr.PrivateKeyError, cferr.ParseFailed)
}
