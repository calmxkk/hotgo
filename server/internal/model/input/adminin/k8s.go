package adminin

// type MemberUpdateCashInp struct {
// 	Name      string `json:"name" v:"required#支付宝姓名不能为空"  dc:"支付宝姓名"`
// 	PayeeCode string `json:"payeeCode" v:"required#支付宝收款码不能为空"  dc:"支付宝收款码"`
// 	Account   string `json:"account" v:"required#支付宝账号不能为空"  dc:"支付宝账号"`
// 	Password  string `json:"password" v:"required#密码不能为空"  dc:"密码"`
// }

type Certificate struct {
	KeyData  []byte `json:"keyData"`
	CertData []byte `json:"certData"`
}

type Proxy struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Forward struct {
	ApiServer string `json:"apiServer"`
	Proxy     Proxy  `json:"proxy"   storm:"inline"`
}

type Connect struct {
	Direction string  `json:"direction"`
	Forward   Forward `json:"forward" storm:"inline"`
}

type Authentication struct {
	Mode              string      `json:"mode"`
	BearerToken       string      `json:"bearerToken"`
	Certificate       Certificate `json:"certificate" storm:"inline"`
	ConfigFileContent []byte      `json:"configFileContent"`
}

type Spec struct {
	Connect        Connect        `json:"connect" storm:"inline"`
	Authentication Authentication `json:"authentication" storm:"inline"`
	Local          bool           `json:"local"`
}

type Status struct {
	Version string `json:"version"`
	Phase   string `json:"phase"`
	Message string `json:"message"`
}

type Cluster struct {
	ApiVersion    string      `json:"apiVersion"`
	Kind          string      `json:"kind"`
	CaCertificate Certificate `json:"caCertificate" storm:"inline"`
	Spec          Spec        `json:"spec" storm:"inline"`
	PrivateKey    []byte      `json:"privateKey"`
	Status        Status      `json:"status" storm:"inline"`
	Labels        []string    `json:"labels"`
}
