package parseconf

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

var configPath = "config.yaml"

// Config is the main strcut
type Config struct {
	Env []Env `yaml:"env"`
}

// Env is a sub struct of Config
type Env struct {
	Name             string `yaml:"name"`
	KeypathInstances string `yaml:"keypath_instances"`
	KeypathBastion   string `yaml:"keypath_bastion"`
	UsernameBastion  string `yaml:"username_bastion"`
	UsernameInstance string `yaml:"username_instance"`
	TagFilter        string `yaml:"tagFilter"`
	TagValuesPrefix  string `yaml:"tagValuesPrefix"`
	TagValueSuffix   string `yaml:"tagValuesSuffix"`
	AwsRegion        string `yaml:"awsRegion"`
}

// GetVarsForThisEnv return datas for specified env given in parameter
func GetVarsForThisEnv(env string, confs *Config) (error, *Env) {
	for _, x := range confs.Env {
		if x.Name == env {
			return nil, &x
		}
	}
	return fmt.Errorf("Unable to find env %s", env), &Env{}
}

// ParseConfigFile parse the file given in parameter
func ParseConfigFile() *Config {
	data, err := fileToBytes(configPath)
	if nil != err {
		log.Error("Unable to parse config file.")
		log.Fatal(err)
	}
	configs := Config{}
	err = yaml.Unmarshal([]byte(data), &configs)
	if nil != err {
		log.Error("Unable read config file. Check, maybe you made a typo. Try to validate it with http://www.yamllint.com/")
		log.Fatal(err)
	}
	return &configs
}

// FileToBytes take path in parameter and return byte slice content
func fileToBytes(path string) ([]byte, error) {
	dat, err := ioutil.ReadFile(path)
	return dat, err
}

// ListEnvFromConfig retrun a slice string from Config.Env
func ListEnvFromConfig(conf *Config) []string {
	var envs []string
	for _, x := range conf.Env {
		envs = append(envs, x.Name)
	}
	return envs
}
