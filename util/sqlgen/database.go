package sqlgen

import (
	"fmt"
)

type Database struct {
	v string
}

func (self Database) String() string {
	return mustParse(sqlIdentifierQuote, Raw{fmt.Sprintf(`%v`, self.v)})
}
