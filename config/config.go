// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package config

import (
	"errors"
	"io/ioutil"
	"net/http"
	sdkUtils "github.com/yunify/qingstor-sdk-go/utils"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/yunify/qingstor-sdk-go/logger"
)

// A Config stores a configuration of this sdk.
type Config struct {
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`

	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Protocol          string `yaml:"protocol"`
	ConnectionRetries int    `yaml:"connection_retries"`

	AdditionalUserAgent string `yaml:"additional_user_agent"`

	LogLevel string `yaml:"log_level"`

	HttpSettings HttpClientSettings

	Connection *http.Client
}

type HttpClientSettings struct {
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
	ReadTimeout time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" `
	TLSHandshakeTimeout time.Duration `yaml:"tls_timeout"`
	IdleConnTimeout time.Duration `yaml:"idle_timeout"`
	TcpKeepAlive time.Duration `yaml:"tcp_keepalive_time"`
	DualStack bool `yaml:"dual_stack"`
	MaxIdleConns int `yaml:"max_idle_conns"`
	MaxIdleConnsPerHost int `yaml:"max_idle_conns_per_host"`
}

var DefaultHttpClientSettings = HttpClientSettings{
	ConnectTimeout: time.Second * 30,
	ReadTimeout: time.Second * 30,
	WriteTimeout: time.Second * 30,
	TLSHandshakeTimeout: time.Second * 10,
	IdleConnTimeout: time.Second * 20,
	TcpKeepAlive: 0,
	DualStack: false,
	MaxIdleConns: 100,
	MaxIdleConnsPerHost: 10,
}

// New create a Config with given AccessKeyID and SecretAccessKey.
func New(accessKeyID, secretAccessKey string) (c *Config, err error) {
	c, err = NewDefault()
	if err != nil {
		c = nil
		return
	}
	c.AccessKeyID = accessKeyID
	c.SecretAccessKey = secretAccessKey

	c.HttpSettings = DefaultHttpClientSettings

	c.InitHttpClient()
	return
}

// NewDefault create a Config with default configuration.
func NewDefault() (c *Config, err error) {
	c = &Config{}
	err = c.LoadDefaultConfig()
	if err != nil {
		c = nil
		return
	}
	c.HttpSettings = DefaultHttpClientSettings

	c.InitHttpClient()
	return
}

// Check checks the configuration.
func (c *Config) Check() (err error) {
	if c.AccessKeyID == "" {
		err = errors.New("access key ID not specified")
		return
	}
	if c.SecretAccessKey == "" {
		err = errors.New("secret access key not specified")
		return
	}

	if c.Host == "" {
		err = errors.New("server host not specified")
		return
	}
	if c.Port <= 0 {
		err = errors.New("server port not specified")
		return
	}
	if c.Protocol == "" {
		err = errors.New("server protocol not specified")
		return
	}

	if c.AdditionalUserAgent != "" {
		for _, x := range c.AdditionalUserAgent {
			// Allow space(32) to ~(126) in ASCII Table, exclude "(34).
			if int(x) < 32 || int(x) > 126 || int(x) == 32 || int(x) == 34 {
				err = errors.New("additional User-Agent contains characters that not allowed")
				return
			}
		}
	}

	err = logger.CheckLevel(c.LogLevel)
	if err != nil {
		return
	}

	return
}

// LoadDefaultConfig loads the default configuration for Config.
// It returns error if yaml decode failed.
func (c *Config) LoadDefaultConfig() (err error) {
	err = yaml.Unmarshal([]byte(DefaultConfigFileContent), c)
	if err != nil {
		logger.Errorf(nil, "Config parse error, %v.", err)
		return
	}

	logger.SetLevel(c.LogLevel)
	return
}

// LoadUserConfig loads user configuration in ~/.qingstor/config.yaml for Config.
// It returns error if file not found.
func (c *Config) LoadUserConfig() (err error) {
	_, err = os.Stat(GetUserConfigFilePath())
	if err != nil {
		logger.Warnf(nil, "Installing default config file to %s.", GetUserConfigFilePath())
		InstallDefaultUserConfig()
	}

	c.HttpSettings = DefaultHttpClientSettings
	return c.LoadConfigFromFilePath(GetUserConfigFilePath())
}

// LoadConfigFromFilePath loads configuration from a specified local path.
// It returns error if file not found or yaml decode failed.
func (c *Config) LoadConfigFromFilePath(filePath string) (err error) {
	if strings.Index(filePath, "~/") == 0 {
		filePath = strings.Replace(filePath, "~/", getHome()+"/", 1)
	}

	yamlString, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Errorf(nil, "File not found: %s.", filePath)
		return err
	}

	return c.LoadConfigFromContent(yamlString)
}

// LoadConfigFromContent loads configuration from a given byte slice.
// It returns error if yaml decode failed.
func (c *Config) LoadConfigFromContent(content []byte) (err error) {
	c.LoadDefaultConfig()

	err = yaml.Unmarshal(content, c)
	if err != nil {
		logger.Errorf(nil, "Config parse error, %v.", err)
		return
	}

	err = c.Check()
	if err != nil {
		return
	}

	logger.SetLevel(c.LogLevel)
	return
}

func (c *Config) InitHttpClient() {
	dialer := sdkUtils.NewDialer(
		c.HttpSettings.ConnectTimeout,
		c.HttpSettings.ReadTimeout,
		c.HttpSettings.WriteTimeout,
	)
	dialer.KeepAlive = c.HttpSettings.TcpKeepAlive
	// XXX: DualStack enables RFC 6555-compliant "Happy Eyeballs" dialing
	// when the network is "tcp" and the destination is a host name
	// with both IPv4 and IPv6 addresses. This allows a client to
	// tolerate networks where one address family is silently broken
	dialer.DualStack = c.HttpSettings.DualStack
	c.Connection = &http.Client{
		// We do not use the timeout in http client,
		// because this timeout is for the whole http body read/write,
		// it's unsuitable for various length of files and network condition.
		// We provide a wraper in utils/conn.go of net.Dialer to make io timeout to the http connection
		// for individual buffer I/O operation,
		Timeout: 0,
		Transport: &http.Transport{
			DialContext: dialer.DialContext,
			MaxIdleConns:          c.HttpSettings.MaxIdleConns,
			MaxIdleConnsPerHost:   c.HttpSettings.MaxIdleConnsPerHost,
			IdleConnTimeout:       c.HttpSettings.IdleConnTimeout,
			TLSHandshakeTimeout:   c.HttpSettings.TLSHandshakeTimeout, //Default
			ExpectContinueTimeout: 2 * time.Second,
		},
	}
}
