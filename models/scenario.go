package models

type JsonScenario struct {
	Scenario Scenario `json:"scenario"`
}

type Scenario struct {
	Time           string             `yaml:"time" json:"time" bson:"time"`
	Period         string             `yaml:"period" json:"period" bson:"period"`
	GloabalActions []string           `yaml:"global_actions" json:"global_actions" bson:"global_actions"`
	Actions        []string           `yaml:"actions" json:"actions" bson:"actions"`
	Services       map[string]Service `yaml:"services" json:"services" bson:"services"`
	NewsChatId     int                `yaml:"news_chat_id" json:"news_chat_id" bson:"news_chat_id"`
	News           []News             `yaml:"news" json:"news" bson:"news"`
}

type News struct {
	Round int    `yaml:"round" json:"round" bson:"round"`
	Text  string `yaml:"text" json:"text" bson:"text"`
	Mode  string `yaml:"mode" json:"mode" bson:"mode"`
}

type Service struct {
	Reputation  int                 `yaml:"reputation" json:"reputation" bson:"reputation"`
	Name        string              `yaml:"name" json:"name" bson:"name"`
	Description string              `yaml:"description" json:"description" bson:"description"`
	Domain      string              `yaml:"domain" json:"domain" bson:"domain"`
	Script      string              `yaml:"script" json:"script" bson:"script"`
	Checkers    []string            `yaml:"checkers" json:"checkers" bson:"checkers"`
	Exploits    map[string]Exploits `yaml:"exploits" json:"exploits" bson:"exploits"`
}

type Exploits struct {
	Round  int    `yaml:"round" json:"round" bson:"round"`
	Cost   int    `yaml:"cost" json:"cost" bson:"cost"`
	Period string `yaml:"period" json:"period" bson:"period"`
	Rounds []int  `yaml:"rounds" json:"rounds" bson:"rounds"`
}
