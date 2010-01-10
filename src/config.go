package netsnail

type Config struct {
    Port int
}

func NewConfig() *Config {
    return new(Config)
}

func (this *Config) ParseArgs() *Config {

    return nil
}
