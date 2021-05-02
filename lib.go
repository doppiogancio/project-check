package main

import "regexp"

func Commands(status []Status, configurations []Configuration) []Configuration {
	commands := map[string]Configuration{}
	commandsAsList := []Configuration{}

	for _, s := range status {
		for _, configuration := range configurations {
			re := regexp.MustCompile(configuration.Regex)
			res := re.FindStringSubmatch(s.File)
			if len(res) == 0 {
				continue
			}

			commands[configuration.Name] = configuration
		}
	}

	for _, c := range commands {
		commandsAsList = append(commandsAsList, c)
	}

	return commandsAsList
}
