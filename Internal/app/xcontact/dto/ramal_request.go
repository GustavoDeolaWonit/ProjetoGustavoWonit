package dto

type RamalRequest struct {
	Numero           string `json:"numero" binding:"required" example:"1200"`
	Nome             string `json:"nome" binding:"required" example:"Gustavo"`
	Senha            string `json:"senha" binding:"required" example:"minhaSenhaForte"`
	Grupo            string `json:"grupo" binding:"required" example:"GR-RAMAIS"`
	Allow            string `json:"allow" example:"opus,g729,ulaw"`
	Insecure         string `json:"insecure" example:"port,invite"`
	SubscribeContext string `json:"subscribecontext" example:"BLF_X5"`
	PickupGroup      string `json:"pickupgroup" example:"1"`
	CallGroup        int    `json:"callgroup" example:"1"`
	Transport        string `json:"transport" example:"udp"`
	CallLimit        int    `json:"call_limit" example:"3"`
	Nat              string `json:"nat" example:"no"`
	Mac              string `json:"mac" example:"00:11:22:33:44:55"`
	AccountCode      string `json:"accountcode" example:"ACC123"`
	DtmfMode         string `json:"dtmfmode" example:"rfc2833"`
	Language         string `json:"language" example:"pt-BR"`
	MusicOnHold      string `json:"musiconhold" example:"default"`
	CallerID         string `json:"callerid" example:"Gustavo <1200>"`
	HasSIP           string `json:"hassip" example:"yes"`
	Encryption       string `json:"encryption" example:"no"`
	Avpf             string `json:"avpf" example:"no"`
	IceSupport       string `json:"icesupport" example:"no"`
	DtlsEnable       string `json:"dtlsenable" example:"no"`
	DtlsVerify       string `json:"dtlsverify" example:"verify"`
	DtlsCertFile     string `json:"dtlscertfile" example:"/etc/ssl/certs/mycert.pem"`
	DtlsSetup        string `json:"dtlssetup" example:"actpass"`
	RtcpMux          string `json:"rtcp_mux" example:"no"`
	XcTipo           string `json:"xc_tipo" example:"ramal"`
	TemMulti         bool   `json:"tem_multi" example:"false"`
	MultiRegra       string `json:"multi_regra" example:""`
}

//RamalMulti       []string `json:"ramalMulti,omitempty" example:"[]"`
