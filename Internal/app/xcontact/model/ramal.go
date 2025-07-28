package model

type Ramal struct {
	Id               int     `json:"id,omitempty"`
	Numero           string  `json:"numero"`
	Nome             string  `json:"nome"`
	Senha            string  `json:"senha"`
	Grupo            string  `json:"grupo"`
	Allow            string  `json:"allow"`
	Insecure         string  `json:"insecure"`
	SubscribeContext string  `json:"subscribecontext"`
	PickupGroup      string  `json:"pickupgroup"`
	CallGroup        int     `json:"callgroup"`
	Transport        string  `json:"transport"`
	CallLimit        int     `json:"call_limit"`
	Nat              string  `json:"nat"`
	Mac              string  `json:"mac"`
	AccountCode      string  `json:"accountcode"`
	DtmfMode         string  `json:"dtmfmode"`
	Language         string  `json:"language"`
	MusicOnHold      string  `json:"musiconhold"`
	CallerID         string  `json:"callerid"`
	HasSIP           string  `json:"hassip"`
	Encryption       string  `json:"encryption"`
	Avpf             string  `json:"avpf"`
	IceSupport       string  `json:"icesupport"`
	DtlsEnable       string  `json:"dtlsenable"`
	DtlsVerify       string  `json:"dtlsverify"`
	DtlsCertFile     string  `json:"dtlscertfile"`
	DtlsSetup        string  `json:"dtlssetup"`
	RtcpMux          string  `json:"rtcp_mux"`
	XcTipo           string  `json:"xc_tipo"`
	TemMulti         bool    `json:"tem_multi"`
	MultiRegra       string  `json:"multi_regra"`
	DefaultUser      *string `json:"defaultuser,omitempty"`
	UserAgent        *string `json:"useragent,omitempty"`
	IpAddr           *string `json:"ipaddr,omitempty"`
	LastMS           *int    `json:"lastms,omitempty"`
}

//RamalMulti       []string `json:"ramalMulti,omitempty"`
