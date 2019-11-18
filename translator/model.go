package translator

type Config struct {
	KeyFrom        string `mapstructure:"key_from" json:"key_from"`
	Key            string `mapstructure:"key" json:"key"`
	Type           string `mapstructure:"type" json:"type"`
	DocType        string `mapstructure:"doc_type" json:"doc_type"`
	TargetLanguage string `mapstructure:"target_language" json:"target_language"`
	Brief          bool   `mapstructure:"brief"`
}
