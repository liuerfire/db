package sqlgen

import (
	"fmt"
	"strings"
)

type Values []Value

type Value struct {
	v interface{}
}

func (self Value) String() string {
	if raw, ok := self.v.(Raw); ok {
		return raw.Raw
	}
	return mustParse(sqlValueQuote, Raw{fmt.Sprintf(`%v`, self.v)})
}

func (self Values) String() string {
	l := len(self)

	if l > 0 {
		chunks := make([]string, 0, l)

		for i := 0; i < l; i++ {
			chunks = append(chunks, self[i].String())
		}

		return strings.Join(chunks, sqlValueSeparator)
	}

	return ""
}
