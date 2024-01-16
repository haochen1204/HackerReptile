package templateType

// Template
// @Description: yaml文件格式
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
	Name   string `yaml:"name"`
}

type Args struct {
	URL       string `yaml:"url,omitempty"`
	By        string `yaml:"by,omitempty"`
	Xpath     string `yaml:"xpath,omitempty"`
	Id        string `yaml:"id,omitempty"`
	Jspath    string `yaml:"jspath,omitempty"`
	Nodeid    string `yaml:"nodeid,omitempty"`
	Query     string `yaml:"query,omitempty"`
	Queryall  string `yaml:"queryall,omitempty"`
	Value     string `yaml:"value,omitempty"`
	Target    string `yaml:"target,omitempty"`
	Attribute string `yaml:"attribute,omitempty"`
	Keys      string `yaml:"keys,omitempty"`
	To        string `yaml:"to"`
}
