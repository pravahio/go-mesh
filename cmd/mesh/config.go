package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"

	json "encoding/json"

	cli "gopkg.in/urfave/cli.v1"
)

type Config struct {
	ConfigMap map[string]string
}

func NewConfig(ctx *cli.Context) *Config {
	c := &Config{
		ConfigMap: make(map[string]string),
	}

	configFile := ctx.String(CONFIG_FILE)

	c.applyConfigFromFile(configFile, ctx)
	c.applyConfigFromConsole(ctx)

	return c
}

func (c *Config) applyConfigFromFile(file string, ctx *cli.Context) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	parsedVal := make(map[string]interface{})

	err = json.Unmarshal(data, &parsedVal)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// TODO: Is passed key in our config or user passing some bogus key config.
	for k, v := range parsedVal {
		if k != BOOTSTRAP_SERVER {
			c.ConfigMap[k] = v.(string)
		} else {
			s := ""
			va := reflect.ValueOf(v)

			for i := 0; i < va.Len(); i++ {
				s += fmt.Sprint(va.Index(i))
				if i < va.Len()-1 {
					s += ","
				}
			}
			c.ConfigMap[k] = s
		}
	}

	return nil
}

func (c *Config) applyConfigFromConsole(ctx *cli.Context) {
	for _, n := range ctx.GlobalFlagNames() {
		g := ctx.Generic(n)

		v := g.(flag.Value).String()
		if v == "false" || v == "" {
			continue
		}
		c.ConfigMap[n] = v
	}
}

func (c *Config) Bool(n string) bool {
	if val, ok := c.ConfigMap[n]; ok {
		parsed, err := strconv.ParseBool(val)
		if err != nil {
			return false
		}
		return parsed
	}
	return false
}

func (c *Config) String(n string) string {

	if val, ok := c.ConfigMap[n]; ok {
		return val
	}
	return ""
}
