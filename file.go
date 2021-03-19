package ini

import (
	"io/ioutil"
	"os"
)

func Load(path string) map[SectionKey]Section {
	f, _ := os.Open(path)

	content, _ := ioutil.ReadAll(f)

	return Loads(string(content))
}

func Dump(path string, json map[SectionKey]Section) error {
	data := Dumps(json)

	return ioutil.WriteFile(path, []byte(data), 777)
}
