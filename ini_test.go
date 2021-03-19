package ini

import (
	"fmt"
	"testing"
)

func TestLoads(t *testing.T) {
	j := Loads("[main]\nkey = 'value==val;lues'\n; key3 = 'value3\nkey432' =     \"423890fsda=432fds=f=\"")

	fmt.Println(j["main"]["key"])

	j2 := make(map[SectionKey]Section)

	j2["main"] = make(Section)
	j2["main"]["key"] = "value"

	j2["debug"] = make(Section)
	j2["debug"]["keyfsafdsa"] = "value::==1=23=:??!#?/'\"a\"'"

	fmt.Println(Dumps(j2))

	Dump("test.ini", j2)

	j3 := Load("test.ini")

	fmt.Println(j3["main"]["key"])
}
