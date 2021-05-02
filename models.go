package main

type (
	Status struct {
		Operation string `json:"operation"`
		File      string `json:"file"`
	}

	Command struct {
		Name string `yaml:"name"`
		Arg  string `json:"arg"`
	}

	Configuration struct {
		Name      string  `yaml:"name"`
		Directory string  `yaml:"directory"`
		Regex     string  `yaml:"regex"`
		Command   Command `yaml:"command"`
	}

	Configurations struct {
		Configuration []Configuration `yaml:"configurations"`
	}
)
