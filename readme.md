# DTCP (debug TCP)

this is advanced TCP debugger, inspired by 'nc' command in
*nix system

## motivation
while developing TCP socket based application, nc command will also send \n as a buffer to the tcp server, this is weird.
sometimes, we also need to generate dummy buffer, such 4096 char random buffer. or normal may http raw request in order get response from server as raw tcp.

this is reason I invented this tool

## usage and install
`make binary`
`make install`
`dtcp 127.0.0.1 8000`

which 127.0.0.1 is IP address, 8000 is portnumber

## command and keyboard
`esc` : exit
`/genbuf` : generate repeat character as many as N
synopsis: "string" character, "int" how many?
`/genrandbuf` : generate random string as many as N
synopsis: "int" how many?
`/n` : generate LF linebreak
`/r` : generate CR linebreak
`/nr` : generate CRLF linebreak
`/nr2` : generate CRLFCRLF linebreak (usually used in HTTP protocol)



## contributing
any feature such udp support, IPV6 support, or other are welcomed

## License
MIT, maintained by <a href="https://github.com/fadhil-riyanto/">Fadhil Riyanto</a>