package graffiti

import (
	"encoding/hex"
	"io/ioutil"
	"strings"

	types "github.com/prysmaticlabs/eth2-types"
	"github.com/prysmaticlabs/prysm/shared/hashutil"
	"gopkg.in/yaml.v2"
)

const (
	hexGraffitiPrefix = "hex"
	hex0xPrefix       = "0x"
)

// Graffiti is a graffiti container.
type Graffiti struct {
	Hash     [32]byte
	Default  string                          `yaml:"default,omitempty"`
	Ordered  []string                        `yaml:"ordered,omitempty"`
	Random   []string                        `yaml:"random,omitempty"`
	Specific map[types.ValidatorIndex]string `yaml:"specific,omitempty"`
}

// ParseGraffitiFile parses the graffiti file and returns the graffiti struct.
func ParseGraffitiFile(f string) (*Graffiti, error) {
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	g := &Graffiti{}
	if err := yaml.Unmarshal(yamlFile, g); err != nil {
		return nil, err
	}

	for i, o := range g.Specific {
		g.Specific[types.ValidatorIndex(i)] = ParseHexGraffiti(o)
	}

	for i, v := range g.Ordered {
		g.Ordered[i] = ParseHexGraffiti(v)
	}

	for i, v := range g.Random {
		g.Random[i] = ParseHexGraffiti(v)
	}

	g.Default = ParseHexGraffiti(g.Default)
	g.Hash = hashutil.Hash(yamlFile)

	return g, nil
}

// ParseHexGraffiti checks if a graffiti input is being represented in hex and converts it to ASCII if so
func ParseHexGraffiti(rawGraffiti string) string {
	splitGraffiti := strings.SplitN(rawGraffiti, ":", 2)
	if strings.ToLower(splitGraffiti[0]) == hexGraffitiPrefix {
		target := splitGraffiti[1]
		if target == "" {
			log.WithField("graffiti", rawGraffiti).Debug("Blank hex tag to be interpreted as itself")
			return rawGraffiti
		}
		if len(target) > 3 && target[:2] == hex0xPrefix {
			target = target[2:]
		}
		if target == "" {
			log.WithField("graffiti", rawGraffiti).Debug("Nothing after 0x prefix, hex tag to be interpreted as itself")
			return rawGraffiti
		}
		graffiti, err := hex.DecodeString(target)
		if err != nil {
			log.WithError(err).Debug("Error while decoding hex string")
			return rawGraffiti
		}
		return string(graffiti)
	}
	return rawGraffiti
}
