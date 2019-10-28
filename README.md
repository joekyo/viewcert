### Build

```shell
go build viewcert.go
```


### Run

```shell
./viewcert example.com
tls version: 1.3

SN: 21020869104500376438182461249190639870
issuer: CN=DigiCert SHA2 Secure Server CA,O=DigiCert Inc,C=US
subject: CN=www.example.org,OU=Technology,O=Internet Corporation for Assigned Names and Numbers,L=Los Angeles,ST=California,C=US
not before: 2018-11-28 00:00:00 +0000 UTC
not after: 2020-12-02 12:00:00 +0000 UTC
key usage: DigitalSignature, KeyEncipherment
is CA: false

SN: 2646203786665923649276728595390119057
issuer: CN=DigiCert Global Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US
subject: CN=DigiCert SHA2 Secure Server CA,O=DigiCert Inc,C=US
not before: 2013-03-08 12:00:00 +0000 UTC
not after: 2023-03-08 12:00:00 +0000 UTC
key usage: DigitalSignature, CertSign, CRLSign
is CA: true

SN: 10944719598952040374951832963794454346
issuer: CN=DigiCert Global Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US
subject: CN=DigiCert Global Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US
not before: 2006-11-10 00:00:00 +0000 UTC
not after: 2031-11-10 00:00:00 +0000 UTC
key usage: DigitalSignature, CertSign, CRLSign
is CA: true
```


### Usage

When you run `viewcert example.com`, `viewcert` will connect to `example.com:443`.
If the server port is not `443` but other port like `8443`, you can run
`viewcert example.com:8443`.
