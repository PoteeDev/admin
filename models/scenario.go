package models

type JsonScenario struct {
	Scenario Scenario `json:"scenario"`
}

type Scenario struct {
	Time     string             `yaml:"time" json:"time"`
	Period   string             `yaml:"period" json:"period"`
	Actions  []string           `yaml:"actions" json:"actions"`
	Services map[string]Service `yaml:"services" json:"services"`
}

type Service struct {
	Reputation int                 `yaml:"reputation" json:"reputation"`
	Domain     string              `yaml:"domain" json:"domain"`
	Script     string              `yaml:"script" json:"script"`
	Exploits   map[string]Exploits `yaml:"exploits" json:"exploits"`
}

type Exploits struct {
	Round  int    `yaml:"round" json:"round"`
	Cost   int    `yaml:"cost" json:"cost"`
	Period string `yaml:"period" json:"period"`
	Rounds []int  `yaml:"rounds" json:"rounds"`
}
