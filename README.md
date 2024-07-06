# adiantum

[Adiantum Encryption](https://security.googleblog.com/2019/02/introducing-adiantum-encryption-for.html)

For each encrypted message a new nonce has to be used.

Nonces should be at least 12 bytes in size.

A keyfile must be 32 hex bytes, without LF/CRLF.

$ adiantum -h, adiantum --help

### Usage of adiantum:

### adiantum [-d] keyfile noncefile < infile > outfile
