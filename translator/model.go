package translator

type Config struct {
	KeyFrom string `mapstructure:"key_from"`
	Key     string `mapstructure:"key"`
	Type    string `mapstructure:"type"`
	DocType string `mapstructure:"doc_type"`
}
