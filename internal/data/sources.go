package data

import "github.com/rbrick/elections/internal/data/sources/us/texas"

type SourceMap map[string]Source

func (s SourceMap) Register(aliases []string, source Source) {
	for _, alias := range aliases {
		s[alias] = source
	}
}

// country code -> region map

var (
	Countries []map[string]SourceMap = []map[string]SourceMap{
		{
			"US": USSources,
		},
	}
)

var USSources = SourceMap{}

func init() {
	USSources.Register([]string{"TX", "TEXAS"}, &texas.Source{})
}
