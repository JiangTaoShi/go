package env

import "os"

var (
	active Environment
	dev    Environment = &environment{value: "dev"}
	fat    Environment = &environment{value: "fat"}
	uat    Environment = &environment{value: "uat"}
	pro    Environment = &environment{value: "pro"}
)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsFat() bool
	IsUat() bool
	IsPro() bool
}

type environment struct {
	Environment
	value string
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == "dev"
}

func (e *environment) IsFat() bool {
	return e.value == "fat"
}

func (e *environment) IsUat() bool {
	return e.value == "uat"
}

func (e *environment) IsPro() bool {
	return e.value == "pro"
}

func init() {
	envStr := os.Getenv("GOENV")
	switch envStr {
	case "dev":
		active = dev
	case "fat":
		active = fat
	case "uat":
		active = uat
	case "pro":
		active = pro
	default:
		active = dev
	}
}

// Active 当前配置的env
func Active() Environment {
	return active
}
