package main

import "fmt"

type IDockerConfigParser interface {
	parseDockerConfig()
}

type jsonDockerConfigParser struct {
}

func (jsonDockerConfigParser) parseDockerConfig() {
	fmt.Println("docker json parse successful")
}

type yamlDockerConfigParser struct {
}

func (yamlDockerConfigParser) parseDockerConfig() {
	fmt.Println("docker yaml parse successful")
}

type ISystemConfigParser interface {
	parseSystemConfig()
}

type jsonSystemConfigParser struct {
}

func (jsonSystemConfigParser) parseSystemConfig() {
	fmt.Println("system json parse successful")
}

type yamlSystemConfigParser struct {
}

func (yamlSystemConfigParser) parseSystemConfig() {
	fmt.Println("system yaml parse successful")
}

type IDockerConfigParserFactory interface {
	createDockerConfigParser() IDockerConfigParser
	createSystemConfigParser() ISystemConfigParser
}

type jsonDockerConfigParserFactory struct {
}

func (jsonDockerConfigParserFactory) createDockerConfigParser() IDockerConfigParser {
	return jsonDockerConfigParser{}
}

func (jsonDockerConfigParserFactory) createSystemConfigParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}

type yamlDockerConfigParserFactory struct {
}

func (yamlDockerConfigParserFactory) createDockerConfigParser() IDockerConfigParser {
	return yamlDockerConfigParser{}
}

func (yamlDockerConfigParserFactory) createSystemConfigParser() ISystemConfigParser {
	return yamlSystemConfigParser{}
}

func newIDockerConfigParserFactory(name string) IDockerConfigParserFactory {
	switch name {
	case "json":
		return jsonDockerConfigParserFactory{}
	case "yaml":
		return yamlDockerConfigParserFactory{}
	}
	return nil
}

func main() {
	var factory IDockerConfigParserFactory
	var name string
	name = "yaml"
	factory = newIDockerConfigParserFactory(name)
	factory.createDockerConfigParser().parseDockerConfig()
	factory.createSystemConfigParser().parseSystemConfig()

	name = "json"
	factory = newIDockerConfigParserFactory(name)
	factory.createDockerConfigParser().parseDockerConfig()
	factory.createSystemConfigParser().parseSystemConfig()
}


