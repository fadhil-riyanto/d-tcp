# DTCP (debug TCP)

this is advanced TCP debugger, inspired by 'nc' command in
*nix system

## motivation
while developing TCP socket based application, nc command will also send \n as a buffer to the tcp server, this is weird

sometimes, we need to generate dummy buffer, such 4096 char random buffer. this is reason I invented this tool

## usage
`./dtcp -p 8000`

spesify host using: `./dtcp -h 127.0.0.1 -p 8000`

by default, DTCP uses 127.0.0.1 loopback as a host

## keyboard key
asap in documentation

## contributing
any feature such udp support, or other are welcomed

## License
MIT, maintained by <a href="https://github.com/fadhil-riyanto/">Fadhil Riyanto</a>