package xcert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestRSATLSConnection(t *testing.T) {

	// generate a test certificate to use
	priv, _ := rsa.GenerateKey(rand.Reader, 4096)

	duration30Days, _ := time.ParseDuration("-30h")
	notBefore := time.Now().Add(duration30Days) // valid 30 hours ago
	duration1Year, _ := time.ParseDuration("90h")
	notAfter := notBefore.Add(duration1Year) // for 90 hours

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Zhangdale"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	template.IPAddresses = append(template.IPAddresses, net.ParseIP("127.0.0.1"))
	template.IPAddresses = append(template.IPAddresses, net.ParseIP("::"))
	template.DNSNames = append(template.DNSNames, "localhost")

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)

	c := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	b := x509.MarshalPKCS1PrivateKey(priv)
	k := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: b})

	// we mock up a server connection to test the client
	listenerKeyPair, _ := tls.X509KeyPair(c, k)

	var listenerTLSConfig tls.Config
	listenerTLSConfig.Certificates = make([]tls.Certificate, 0)
	listenerTLSConfig.Certificates = append(listenerTLSConfig.Certificates, listenerKeyPair)
	_, _ = tls.Listen("tcp", ":0", &listenerTLSConfig)
	// todo

	// mock up the client side too
	clientTLSCertPool := x509.NewCertPool()
	clientTLSCertPool.AppendCertsFromPEM(c)

	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:    clientTLSCertPool,
			ServerName: "localhost",
		},
	}
	// http.DefaultClient.Do()
	// todo
}
