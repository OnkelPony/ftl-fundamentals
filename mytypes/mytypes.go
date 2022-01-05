package mytypes

import "strings"

type StringUppercaser struct {
	Contents strings.Builder
}

type IntBuilder struct {
	Contents []int
}

type MyInt int

func (i MyInt) Triple() MyInt {
	return 3 * i
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

func (ib IntBuilder) Sum() int {
	var result int
	for _, v := range ib.Contents {
		result += v
	}
	return result
}
