package ini

import (
	"fmt"
	"testing"
)

func TestLoads(t *testing.T) {
	j := Loads("[main]\nkey = 'value==val;lues'\n; key3 = 'value3\nkey432' =     \"423890fsda=432fds=f=\"")

	fmt.Println(j["main"]["key"])

	j2 := make(map[SectionKey]SectionValue)

	j2["main"] = make(SectionValue)
	j2["main"]["key"] = "value"

	j2["debug"] = make(SectionValue)
	j2["debug"]["keyfsafdsa"] = "value::==1=23=:??!#?/'\"a\"'"

	fmt.Println(Dumps(j2))
}
