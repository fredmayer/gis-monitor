package config

type Config struct {
	LogLevel string `mapstructure:"log-level"`

	Interval int64 `mapstructure:"interval"`

	Db Db `mapstructure:"db"`

	Gis Gis `mapstructure:"gis"`

	Tg Tg `mapstructure:"tg"`
}

type Tg struct {
	Url   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
}

type Gis struct {
	Host     string            `mapstructure:"host"`
	Endpoint string            `mapstructure:"endpoint"`
	Link     string            `mapstructure:"link"`
	Params   map[string]string `mapstructure:"params"`
}

type Db struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}
