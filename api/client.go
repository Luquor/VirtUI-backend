package api

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var URL_API_LXD = "127.0.0.1:8443"

type Client struct {
	Client *http.Client
}

var Cli = Client{
	Client: generateClient(),
}

func (c Client) Get(endpoint string) string {
	resp, err := c.Client.Get(fmt.Sprintf("https://%s%s", URL_API_LXD, endpoint))
	if err != nil {
		log.Fatal(err)
	}
	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(msg)
}

// TO DO : Définir plus précisément le type de "data"
func (c Client) Post(endpoint string, data any) string {
	jsonData, _ := json.Marshal(data)
	reader := bytes.NewReader(jsonData)
	resp, err := c.Client.Post(fmt.Sprintf("https://%s%s", URL_API_LXD, endpoint), "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}
	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(msg)
}

func (c Client) Delete(endpoint string) string {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://%s%s", URL_API_LXD, endpoint), nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, _ := c.Client.Do(req)
	msg, _ := io.ReadAll(resp.Body)
	return string(msg)
}

func (c Client) Put(endpoint string, data any) string {
	jsonData, _ := json.Marshal(data)
	reader := bytes.NewReader(jsonData)
	resp, _ := c.Client.Post(fmt.Sprintf("https://%s%s", URL_API_LXD, endpoint), "application/json", reader)
	msg, _ := io.ReadAll(resp.Body)
	return string(msg)
}

var (
	CACertFilePath = "tls/api_server.crt"
	CertFilePath   = "tls/client.crt"
	KeyFilePath    = "tls/client.key"
)

func generateClient() *http.Client {
	// load tls certificates
	clientTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
		return nil
	}
	// Configure the client to trust TLS server certs issued by a CA.
	certPool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}
	if caCertPEM, err := os.ReadFile(CACertFilePath); err != nil {
		panic(err)
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		panic("invalid cert in CA PEM")
	}
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientTLSCert},
	}
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{Transport: tr}
	return client
}
