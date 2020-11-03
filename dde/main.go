package dde

import (
	"io/ioutil"
	"log"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

type Entry struct {
	Model      string                 `yaml:"Model"`
	Type       string                 `yaml:"Type"`
	Attributes map[string]interface{} `yaml:"Attributes"`
}

type Environment struct {
	Registry map[string]interface{}
	Entries  map[string]Entry
}

func (e *Environment) Create(environmentFile string) {
	// Handle error properly.
	buf, _ := ioutil.ReadFile("./features/dde/" + environmentFile + ".yaml")

	var config []Entry
	yaml.Unmarshal(buf, &config)

	for _, entry := range config {
		e.Entries[entry.Model] = entry
	}

	for name, entry := range e.Entries {
		model := reflect.New(reflect.TypeOf(e.Registry[name])).Interface()
		mapstructure.Decode(entry.Attributes, model)

		log.Printf("Create from model: <%T> %+v", model, model)
	}
}

func (e *Environment) Destroy() {
	// Destroy stuff
}

func (e *Environment) Register(name string, item interface{}) {
	e.Registry[name] = item
}

func NewEnvironment() Environment {
	return Environment{
		Registry: make(map[string]interface{}, 0),
		Entries:  make(map[string]Entry, 0),
	}
}
