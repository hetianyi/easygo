package base

import (
	"gopkg.in/yaml.v2"
)

// ParseYamlFromString 从字符串解析yaml为对象
func ParseYamlFromString(input string, target interface{}) error {
	return yaml.Unmarshal([]byte(input), target)
}

// Convert2Yaml 将对象转为yaml文本
func Convert2Yaml(c interface{}) (string, error) {
	s, err := yaml.Marshal(c)
	return string(s), err
}
