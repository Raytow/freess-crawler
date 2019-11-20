package main

type SS struct {
	ip string
	port string
	password string
	method string
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

type GuiConfig struct {
	Index int `json:"index"`
	Random bool `json:"random"`
	Global bool `json:"global"`
	Enabled bool `json:"enabled"`
	ShareOverLan bool `json:"shareOverLan"`
	IsDefault bool `json:"isDefault"`
	LocalPort int `json:"localPort"`
	PacUrl *string `json:"pacUrl"`
	UseOnlinePac bool `json:"useOnlinePac"`
	ReconnectTimes int `json:"reconnectTimes"`
	RandomAlgorithm int `json:"randomAlgorithm"`
	TTL int `json:"TTL"`
	ProxyEnable bool `json:"proxyEnable"`
	ProxyType int `json:"proxyType"`
	ProxyHost *string `json:"proxyHost"`
	ProxyPort int `json:"proxyPort"`
	ProxyAuthUser *string `json:"proxyAuthUser"`
	ProxyAuthPass *string `json:"proxyAuthPass"`
	AuthUser *string `json:"authUser"`
	AuthPass *string `json:"authPass"`
	Autoban bool `json:"autoban"`
	Configs []*ConfigEle `json:"configs"`
}

type ConfigEle struct {
	Remarks string `json:"remarks"`
	Server string `json:"server"`
	ServerPort int `json:"server_port"`
	Method string `json:"method"`
	Obfs string `json:"obfs"`
	Obfsparam string `json:"obfsparam"`
	RemarksBase64 string `json:"remarks_base64"`
	Password string `json:"password"`
	TcpOverUdp bool `json:"tcp_over_udp"`
	UdpOverTcp bool `json:"udp_over_tcp"`
	Protocol string `json:"protocol"`
	ObfsUdp bool `json:"obfs_udp"`
	Enable bool `json:"enable"`
}