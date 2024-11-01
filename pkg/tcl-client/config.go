package tcl_client

type Config struct {
	Host      string `envconfig:"optional"`
	Namespace string `envconfig:"default=default"`
	Prefix    string
}
