# Examples

Package allows to automatize testing of `Marshal*` and `Unmarshal*` methods.

Suppose we have the following type:

`binary/byte.go`
```go
package binary

import (
	"fmt"
	"strconv"
)

type Byte byte

func (b Byte) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%08b", b)), nil
}

func (b *Byte) UnmarshalText(data []byte) error {
	u, err := strconv.ParseUint(string(data), 2, 8)
	if err != nil {
		return err
	}
	*b = Byte(u)
	return nil
}
```

Package `go.lstv.dev/util/test` allowing to create tests by following way:

`binary/byte_test.go`
```go
package binary

import (
	"testing"

	"go.lstv.dev/util/test"
)

var ByteTextCases = []test.CaseText[Byte]{
	{ // 0
		Data:  "00000000",
		Value: 0,
	},
	{ // 1
		Data:  "00000001",
		Value: 1,
	},
	{ // 2
		Data:  "11111111",
		Value: 255,
	},
	{ // 3
		Constraint: test.OnlyUnmarshal,
		Error:      test.Error("strconv.ParseUint: parsing \"2\": invalid syntax"),
		Data:       "2",
	},
}

func TestByteMarshalText(t *testing.T) {
	test.MarshalText(t, ByteTextCases)
}

func TestByteUnmarshalText(t *testing.T) {
	test.UnmarshalText(t, ByteTextCases, nil)
}
```
