package templateType

type Template struct {
	ID       string `yaml:"id"`
	Info     Info   `yaml:"info"`
	Headless []struct {
		Steps []Step `yaml:"steps"`
	} `yaml:"headless"`
}

type Info struct {
	Name        string `yaml:"name"`
	Author      string `yaml:"author"`
	Description string `yaml:"description"`
}

type Step struct {
	Args   Args   `yaml:"args"`
	Action string `yaml:"action"`
}

type Args struct {
	URL       string `yaml:"url,omitempty"`
	By        string `yaml:"by,omitempty"`
	Xpath     string `yaml:"xpath,omitempty"`
	Value     string `yaml:"value,omitempty"`
	Target    string `yaml:"target,omitempty"`
	Attribute string `yaml:"attribute,omitempty"`
}
