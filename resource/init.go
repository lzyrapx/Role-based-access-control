package resource

var parsers = &resParser{}

func init() {
	RegisterParseFunc(ParseGlobalRes, ParseNameRes, ParseURLRes)
}