package main

import "fmt"

type IDockerConfigParser interface {
	Parse()
}

type jsonDockerConfigParser struct {
}

func (jsonDockerConfigParser) Parse() {
	fmt.Println("json parse successful")
}

type yamlDockerConfigParser struct {
}

func (yamlDockerConfigParser) Parse() {
	fmt.Println("yaml parse successful")
}

type IDockerConfigParserFactory interface {
	CreateParser() IDockerConfigParser
}

type yamlDockerConfigParserFactory struct {
}

func (yamlDockerConfigParserFactory) CreateParser() IDockerConfigParser {
	return yamlDockerConfigParser{}
}

type jsonDockerConfigParserFactory struct {
}

func (jsonDockerConfigParserFactory) CreateParser() IDockerConfigParser {
	return jsonDockerConfigParser{}
}

func NewIDockerConfigParserFactory(name string) IDockerConfigParserFactory {
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
	factory = NewIDockerConfigParserFactory(name)
	factory.CreateParser().Parse()

	name = "json"
	factory = NewIDockerConfigParserFactory(name)
	factory.CreateParser().Parse()
}
