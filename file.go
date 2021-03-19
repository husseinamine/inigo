package ini

import (
	"io/ioutil"
	"os"
)

func Load(path string) (map[SectionKey]Section, error) {
	f, err := os.Open(path)
	defer f.Close()

	content, _ := ioutil.ReadAll(f)

	return Loads(string(content)), err
}

func Dump(json map[SectionKey]Section, path string) error {
	data := Dumps(json)

	return ioutil.WriteFile(path, []byte(data), 777)
}
