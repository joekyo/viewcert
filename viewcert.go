// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: viewcert example.com")
	}
	addr := os.Args[1]
	if _, _, err := net.SplitHostPort(addr); err != nil {
		addr = addr + ":443"
	}
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	state := conn.ConnectionState()

	printTLSVersion(state.Version)
	printTLSCertificates(state.PeerCertificates)
}

func printTLSVersion(version uint16) {
	fmt.Print("tls version: ")
	switch version {
	case tls.VersionTLS10:
		fmt.Println("1.0")
	case tls.VersionTLS11:
		fmt.Println("1.1")
	case tls.VersionTLS12:
		fmt.Println("1.2")
	case tls.VersionTLS13:
		fmt.Println("1.3")
	default:
		fmt.Println("unknown")
	}
}

func printTLSCertificates(certs []*x509.Certificate) {
	for _, cert := range certs {
		fmt.Println()
		fmt.Println("SN:", cert.SerialNumber)
		fmt.Println("issuer:", cert.Issuer)
		fmt.Println("subject:", cert.Subject)
		fmt.Println("not before:", cert.NotBefore)
		fmt.Println("not after:", cert.NotAfter)
		fmt.Println("key usage:", getKeyUsage(cert.KeyUsage))
		fmt.Println("is CA:", cert.IsCA)
	}
}

func getKeyUsage(keyUsage x509.KeyUsage) string {
	usage := make([]string, 0)
	if keyUsage&x509.KeyUsageDigitalSignature != 0 {
		usage = append(usage, "DigitalSignature")
	}
	if keyUsage&x509.KeyUsageContentCommitment != 0 {
		usage = append(usage, "ContentCommitment")
	}
	if keyUsage&x509.KeyUsageKeyEncipherment != 0 {
		usage = append(usage, "KeyEncipherment")
	}
	if keyUsage&x509.KeyUsageDataEncipherment != 0 {
		usage = append(usage, "DataEncipherment")
	}
	if keyUsage&x509.KeyUsageKeyAgreement != 0 {
		usage = append(usage, "KeyAgreement")
	}
	if keyUsage&x509.KeyUsageCertSign != 0 {
		usage = append(usage, "CertSign")
	}
	if keyUsage&x509.KeyUsageCRLSign != 0 {
		usage = append(usage, "CRLSign")
	}
	if keyUsage&x509.KeyUsageEncipherOnly != 0 {
		usage = append(usage, "EncipherOnly")
	}
	if keyUsage&x509.KeyUsageDecipherOnly != 0 {
		usage = append(usage, "DecipherOnly")
	}
	if len(usage) == 0 {
		return "unknown"
	}
	return strings.Join(usage, ", ")
}
