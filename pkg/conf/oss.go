package conf

type Oss struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	EndPoint  string `yaml:"end_point"`
	Bucket    string `yaml:"bucket"`
	Domain    string `yaml:"domain"`
}
