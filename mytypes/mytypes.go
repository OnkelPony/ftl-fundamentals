package mytypes

import "strings"

type StringUppercaser struct {
	Contents strings.Builder
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}
