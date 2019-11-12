package main

type SS struct {
	ip string
	port string
	password string
	security string
}

type SSR struct {
	ip string
	port string
	password string
	security string
	protocol string
	obfuscation string
}

type V2ray struct {
	address string
	port string
	uid string
	network string
	path string
	tls string
}