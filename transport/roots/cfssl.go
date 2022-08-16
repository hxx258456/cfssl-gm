package roots

import (
	"encoding/json"
	"errors"

	"github.com/hxx258456/ccgo/x509"

	"github.com/hxx258456/cfssl-gm/api/client"
	"github.com/hxx258456/cfssl-gm/helpers"
	"github.com/hxx258456/cfssl-gm/info"
)

// This package contains CFSSL integration.

// NewCFSSL produces a new CFSSL root.
func NewCFSSL(metadata map[string]string) ([]*x509.Certificate, error) {
	host, ok := metadata["host"]
	if !ok {
		return nil, errors.New("transport: CFSSL root provider requires a host")
	}

	label := metadata["label"]
	profile := metadata["profile"]
	cert, err := helpers.LoadClientCertificate(metadata["mutual-tls-cert"], metadata["mutual-tls-key"])
	if err != nil {
		return nil, err
	}
	remoteCAs, err := helpers.LoadPEMCertPool(metadata["tls-remote-ca"])
	if err != nil {
		return nil, err
	}
	srv := client.NewServerTLS(host, helpers.CreateTLSConfig(remoteCAs, cert))
	data, err := json.Marshal(info.Req{Label: label, Profile: profile})
	if err != nil {
		return nil, err
	}

	resp, err := srv.Info(data)
	if err != nil {
		return nil, err
	}

	return helpers.ParseCertificatesPEM([]byte(resp.Certificate))
}
