package boot

type Conf struct{
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
	ServiceName string `env:"SERVICE_NAME"`
}

var Config = &Conf{}
