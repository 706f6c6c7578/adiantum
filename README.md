# adiantum

[Adiantum Encryption](https://security.googleblog.com/2019/02/introducing-adiantum-encryption-for.html)

This is a modified version of [ilius's CLI version](https://github.com/ilius/adiantum-cli) so that Windows users can use the software as well.

Modifications where made for key and nonce usage, so that users can load a key file instead of typing in a passphrase and a nonce file when encrypting or decrypting messages.

Nonce usage requires that for each encrypted message a new nonce has to be used.

Nonces should be at least 12 bytes in size.

A key must be 32 hex bytes in size!

$ adiantum -h, adiantum --help

### Usage of adiantum:

### adiantum [-d] keyfile noncefile < infile > outfile
