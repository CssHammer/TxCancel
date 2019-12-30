package middleware

import (
	"txCancel/config/api"
)

type Middlewares struct {
	sourceTypeList map[string]struct{}
}

func NewMiddlewares(conf *api.Config) *Middlewares {
	// map of valid source types, faster than slice
	sourceTypeList := make(map[string]struct{})
	for _, sourceType := range conf.SourceTypes {
		sourceTypeList[sourceType] = struct{}{}
	}

	return &Middlewares{sourceTypeList: sourceTypeList}
}
