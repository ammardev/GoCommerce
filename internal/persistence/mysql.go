package persistence

import "strings"

type BaseMySqlRepository struct {
	UpdatesBuilder updatesBuilder
}

type updatesBuilder struct {
	fields []string
	Values []any
}

func (builder *updatesBuilder) Add(field string, value any) {
	if value == nil {
		return
	}

	builder.fields = append(builder.fields, field+" = ?")
	builder.Values = append(builder.Values, value)
}

func (builder *updatesBuilder) GetQuery() string {
	return strings.Join(builder.fields, ", ")
}
