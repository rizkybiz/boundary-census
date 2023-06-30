package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/shipyard-run/hclconfig"
	"github.com/shipyard-run/hclconfig/types"
)

var (
	nomadEnvPrefix    = "NOMAD_"
	boundaryEnvPrefix = "BOUNDARY_"
)

// Config defines a struct that holds the config for the controller
type Config struct {
	types.ResourceMetadata `hcl:",remain"`

	Nomad    *Nomad    `hcl:"nomad,block"`
	Boundary *Boundary `hcl:"boundary,block"`
}

// Nomad is configuration specific to the Nomad scheduler
type Nomad struct {
	Address   string `hcl:"address,optional"`
	Token     string `hcl:"token,optional"`
	Region    string `hcl:"region,optional"`
	Namespace string `hcl:"namespace,optional"`
}

// / Boundary is configuration specific to Boundary
type Boundary struct {
	Enterprise     bool     `hcl:"enterprise,optional"`
	OrgID          string   `hcl:"org_id"`
	DefaultProject string   `hcl:"default_project,optional"`
	DefaultGroups  []string `hcl:"default_groups,optional"`

	AuthMethodID string `hcl:"auth_method_id"`
	Username     string `hcl:"username"`
	Password     string `hcl:"password"`
	Address      string `hcl:"address"`

	DefaultIngressFilter string `hcl:"default_ingress_filter,optional"`
	DefaultEgressFilter  string `hcl:"default_egress_filter,optional"`
}

// Process is called by hclconfig when it finds a Config resource
// you can do validation and cleanup here
func (c *Config) Process() error {
	c.Boundary.DefaultIngressFilter = strings.TrimSpace(c.Boundary.DefaultIngressFilter)
	c.Boundary.DefaultEgressFilter = strings.TrimSpace(c.Boundary.DefaultEgressFilter)

	// check for ENV VARS
	c.checkEnvVars()

	if !c.Boundary.Enterprise {
		if c.Boundary.DefaultIngressFilter != "" || c.Boundary.DefaultEgressFilter != "" {
			return fmt.Errorf("ingress filters are not supported by oss boundary")
		}
	}

	return nil
}

// Parse the given HCL config file and return the Config
func Parse(config string) (*Config, error) {
	p := hclconfig.NewParser(hclconfig.DefaultOptions())
	p.RegisterType("config", &Config{})
	p.RegisterFunction("trim", func(s string) (string, error) {
		return strings.TrimSpace(s), nil
	})

	c := hclconfig.NewConfig()
	err := p.ParseFile(config, c)
	if err != nil {
		return nil, fmt.Errorf("unable to process file: %s, error: %s", config, err)
	}

	r, err := c.FindResourcesByType("config")
	if err != nil {
		return nil, fmt.Errorf("unable to process file: %s, error: %s", config, err)
	}

	if len(r) != 1 {
		return nil, fmt.Errorf("unable to process file: %s, file does not contain a single config resource", config)
	}

	return r[0].(*Config), nil
}

// checkEnvVars will check for existing env vars and override
// the config values for those env vars
func (c *Config) checkEnvVars() {

	n := nomadEnvPrefix
	b := boundaryEnvPrefix

	if x := os.Getenv(n + "ADDRESS"); x != "" {
		c.Nomad.Address = x
	}
	if x := os.Getenv(n + "TOKEN"); x != "" {
		c.Nomad.Token = x
	}
	if x := os.Getenv(n + "REGION"); x != "" {
		c.Nomad.Region = x
	}
	if x := os.Getenv(n + "NAMESPACE"); x != "" {
		c.Nomad.Namespace = x
	}
	switch x := os.Getenv(b + "ENTERPRISE"); x {
	case "":
		break
	case "true":
		c.Boundary.Enterprise = true
	case "false":
		c.Boundary.Enterprise = false
	}
	if x := os.Getenv(b + "ORG_ID"); x != "" {
		c.Boundary.OrgID = x
	}
	if x := os.Getenv(b + "DEFAULT_PROJECT"); x != "" {
		c.Boundary.DefaultProject = x
	}
	if x := os.Getenv(b + "DEFAULT_GROUPS"); x != "" {
		dp := strings.Split(x, ",")
		c.Boundary.DefaultGroups = dp
	}
	if x := os.Getenv(b + "AUTH_METHOD_ID"); x != "" {
		c.Boundary.AuthMethodID = x
	}
	if x := os.Getenv(b + "USERNAME"); x != "" {
		c.Boundary.Username = x
	}
	if x := os.Getenv(b + "PASSWORD"); x != "" {
		c.Boundary.Password = x
	}
	if x := os.Getenv(b + "ADDRESS"); x != "" {
		c.Boundary.Address = x
	}
	if x := os.Getenv(b + "DEFAULT_INGRESS_FILTER"); x != "" {
		c.Boundary.DefaultIngressFilter = x
	}
	if x := os.Getenv(b + "DEFAULT_EGRESS_FILTER"); x != "" {
		c.Boundary.DefaultEgressFilter = x
	}
}
