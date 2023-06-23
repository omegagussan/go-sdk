/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"net/http"

	"github.com/openfga/go-sdk/credentials"
	"github.com/openfga/go-sdk/internal/utils"
)

const (
	SdkVersion = "0.2.2"

	defaultUserAgent = "openfga-sdk go/0.2.2"
)

// RetryParams configures configuration for retry in case of HTTP too many request
type RetryParams struct {
	MaxRetry    int `json:"maxRetry,omitempty"`
	MinWaitInMs int `json:"minWaitInMs,omitempty"`
}

// Configuration stores the configuration of the API client
type Configuration struct {
	ApiScheme      string                   `json:"api_scheme,omitempty"`
	ApiHost        string                   `json:"api_host,omitempty"`
	StoreId        string                   `json:"store_id,omitempty"`
	Credentials    *credentials.Credentials `json:"credentials,omitempty"`
	DefaultHeaders map[string]string        `json:"default_headers,omitempty"`
	UserAgent      string                   `json:"user_agent,omitempty"`
	Debug          bool                     `json:"debug,omitempty"`
	HTTPClient     *http.Client
	RetryParams    *RetryParams
}

// DefaultRetryParams returns the default retry parameters
func DefaultRetryParams() *RetryParams {
	return &RetryParams{
		MaxRetry:    15,
		MinWaitInMs: 100,
	}
}

func GetSdkUserAgent() string {
	return defaultUserAgent
}

// NewConfiguration returns a new Configuration object
func NewConfiguration(config Configuration) (*Configuration, error) {
	cfg := &Configuration{
		ApiScheme:      config.ApiScheme,
		ApiHost:        config.ApiHost,
		StoreId:        config.StoreId,
		Credentials:    config.Credentials,
		DefaultHeaders: config.DefaultHeaders,
		UserAgent:      config.UserAgent,
		Debug:          false,
		RetryParams:    config.RetryParams,
	}

	if cfg.ApiScheme == "" {
		cfg.ApiScheme = "https"
	}

	if cfg.UserAgent == "" {
		cfg.UserAgent = GetSdkUserAgent()
	}

	if cfg.DefaultHeaders == nil {
		cfg.DefaultHeaders = make(map[string]string)
	}

	err := cfg.ValidateConfig()

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeaders[key] = value
}

// ValidateConfig ensures that the given configuration is valid
func (c *Configuration) ValidateConfig() error {
	if c.ApiHost == "" {
		return reportError("Configuration.ApiHost is required")
	}

	if c.ApiScheme == "" {
		return reportError("Configuration.ApiScheme is required")
	}

	if !IsWellFormedUri(c.ApiScheme + "://" + c.ApiHost) {
		return reportError("Configuration.ApiScheme and Configuration.ApiHost (%s) do not generate a valid uri", c.ApiScheme+"://"+c.ApiHost)
	}

	if c.Credentials != nil {
		if err := c.Credentials.ValidateCredentialsConfig(); err != nil {
			return reportError("Credentials are invalid: %v", err)
		}
	}

	if c.RetryParams != nil && c.RetryParams.MaxRetry > 15 {
		return reportError("Configuration.RetryParams.MaxRetry exceeds maximum allowed limit of 15")
	}

	if c.StoreId != "" && !internalutils.IsWellFormedUlidString(c.StoreId) {
		return reportError("Configuration.StoreId is not a valid ulid")
	}

	return nil
}
