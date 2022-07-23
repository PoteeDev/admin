package models

type JsonScenario struct {
	Scenario Scenario `json:"scenario"`
}

type Scenario struct {
	Rounds   int                `yaml:"rounds" json:"rounds"`
	Period   string             `yaml:"period" json:"period"`
	Actions  []string           `yaml:"actions" json:"actions"`
	Services map[string]Service `yaml:"services" json:"services"`
}

type Service struct {
	Hp       int                 `yaml:"hp" json:"hp"`
	Domain   string              `yaml:"domain" json:"domain"`
	Script   string              `yaml:"script" json:"script"`
	Exploits map[string]Exploits `yaml:"exploits" json:"exploits"`
}

type Exploits struct {
	Round int `yaml:"round" json:"round"`
	Cost  int `yaml:"cost" json:"cost"`
}
