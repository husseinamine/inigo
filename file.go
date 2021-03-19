package ini

import (
	"io/ioutil"
)

func Load(path string) (map[SectionKey]Section, error) {
	f, err := ioutil.ReadFile(path)

	return Loads(string(f)), err
}

func Dump(json map[SectionKey]Section, path string) error {
	data := Dumps(json)

	return ioutil.WriteFile(path, []byte(data), 777)
}
