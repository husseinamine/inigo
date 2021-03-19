package ini

import (
	"fmt"
	"strings"
)

type SectionKey string
type Section map[string]string

func Loads(config string) map[SectionKey]Section {
	data := make(map[SectionKey]Section)
	section := SectionKey("")

	for _, line := range strings.Split(config, "\n") {
		if strings.HasPrefix(line, ";") {
			continue
		} else if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			line = strings.Replace(line, "[", "", 1)
			line = strings.Replace(line, "]", "", 1)

			section = SectionKey(strings.TrimSpace(line))

			if _, ok := data[section]; !ok {
				data[section] = make(Section)
			}
		} else {
			if strings.Contains(line, "=") {
				key, value := strings.Split(line, "=")[0], strings.Join(strings.Split(line, "=")[1:], "=")

				key = strings.TrimSpace(key)
				key = strings.Trim(key, "\"")
				key = strings.Trim(key, "'")

				value = strings.TrimSpace(value)
				value = strings.Trim(value, "\"")
				value = strings.Trim(value, "'")

				data[section][key] = value
			}
		}
	}

	return data
}

func Dumps(json map[SectionKey]Section) string {
	text := ""

	for section := range json {
		text += fmt.Sprintf("[%s]\n", section)

		for key, value := range json[section] {
			text += fmt.Sprintf("%s = %s\n", key, value)
		}

		text += "\n"
	}

	return strings.TrimSpace(text)
}
