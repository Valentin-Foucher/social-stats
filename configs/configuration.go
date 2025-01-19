package configs

type Server struct {
	BaseUrl string `yaml:"baseUrl"`
	Port    int    `yaml:"port"`
}

type Strava struct {
	PublicKey string `yaml:"public"`
	SecretKey string `yaml:"secret"`
	BaseUrl   string `yaml:"baseUrl"`
}

type Http struct {
	MaxRetries int `yaml:"maxRetries"`
}

type Configuration struct {
	Server *Server `yaml:"server"`
	Strava *Strava `yaml:"strava"`
	Http   *Http   `yaml:"http"`
}
