package nameparser

import "strings"

type NameParser struct {
}

func (n NameParser) Parse(fullName string) ParsedName {
	split := strings.Split(fullName, " ")
	return ParsedName{
		Name:    split[0],
		Surname: strings.Join(split[1:], " "),
	}
}

type ParsedName struct {
	Name    string
	Surname string
}

type Parser struct {
	nameParser NameParser
}

func (p Parser) Parse(fullName string) ParsedName {
	return p.nameParser.Parse(fullName)
}
