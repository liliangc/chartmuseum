/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"github.com/urfave/cli"
)

type (
	configVar struct {
		Type    configVarType
		Default interface{}
		CLIFlag cli.Flag
	}

	configVarType string
)

// Will be populated in init() below
var CLIFlags []cli.Flag

var (
	stringType configVarType = "string"
	intType    configVarType = "int"
	boolType   configVarType = "bool"
)

var configVars = map[string]configVar{
	"genindex": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "gen-index",
			Usage:  "generate index.yaml, print to stdout and exit",
			EnvVar: "GEN_INDEX",
		},
	},
	"debug": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "debug",
			Usage:  "show debug messages",
			EnvVar: "DEBUG",
		},
	},
	"logjson": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "log-json",
			Usage:  "output structured logs as json",
			EnvVar: "LOG_JSON",
		},
	},
	"loghealth": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "log-health",
			Usage:  "log inbound /health requests",
			EnvVar: "LOG_HEALTH",
		},
	},
	"disablemetrics": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-metrics",
			Usage:  "disable Prometheus metrics",
			EnvVar: "DISABLE_METRICS",
		},
	},
	"disableapi": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-api",
			Usage:  "disable all routes prefixed with /api",
			EnvVar: "DISABLE_API",
		},
	},
	"disablestatefiles": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-statefiles",
			Usage:  "disable use of index-cache.yaml",
			EnvVar: "DISABLE_STATEFILES",
		},
	},
	"allowoverwrite": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "allow-overwrite",
			Usage:  "allow chart versions to be re-uploaded without ?force querystring",
			EnvVar: "ALLOW_OVERWRITE",
		},
	},
	"disableforceoverwrite": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "disable-force-overwrite",
			Usage:  "do not allow chart versions to be re-uploaded, even with ?force querystring",
			EnvVar: "DISABLE_FORCE_OVERWRITE",
		},
	},
	"port": {
		Type:    intType,
		Default: 8080,
		CLIFlag: cli.IntFlag{
			Name:   "port",
			Usage:  "port to listen on",
			EnvVar: "PORT",
		},
	},
	"charturl": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "chart-url",
			Usage:  "absolute url for .tgzs in index.yaml",
			EnvVar: "CHART_URL",
		},
	},
	"basicauth.user": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "basic-auth-user",
			Usage:  "username for basic http authentication",
			EnvVar: "BASIC_AUTH_USER",
		},
	},
	"basicauth.pass": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "basic-auth-pass",
			Usage:  "password for basic http authentication",
			EnvVar: "BASIC_AUTH_PASS",
		},
	},
	"authanonymousget": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "auth-anonymous-get",
			Usage:  "allow anonymous GET operations when auth is used",
			EnvVar: "AUTH_ANONYMOUS_GET",
		},
	},
	"tls.cert": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "tls-cert",
			Usage:  "path to tls certificate chain file",
			EnvVar: "TLS_CERT",
		},
	},
	"tls.key": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "tls-key",
			Usage:  "path to tls key file",
			EnvVar: "TLS_KEY",
		},
	},
	"cache.store": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "cache",
			Usage:  "cache store, can be one of: redis",
			EnvVar: "CACHE",
		},
	},
	"cache.redis.addr": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "cache-redis-addr",
			Usage:  "address of Redis service (host:port)",
			EnvVar: "CACHE_REDIS_ADDR",
		},
	},
	"cache.redis.password": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "cache-redis-password",
			Usage:  "Redis requirepass server configuration",
			EnvVar: "CACHE_REDIS_PASSWORD",
		},
	},
	"cache.redis.db": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "cache-redis-db",
			Usage:  "Redis database to be selected after connect",
			EnvVar: "CACHE_REDIS_DB",
			Value:  0,
		},
	},
	"storage.backend": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage",
			Usage:  "storage backend, can be one of: local, amazon, google, oracle",
			EnvVar: "STORAGE",
		},
	},
	"storage.local.rootdir": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-local-rootdir",
			Usage:  "directory to store charts for local storage backend",
			EnvVar: "STORAGE_LOCAL_ROOTDIR",
		},
	},
	"storage.amazon.bucket": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-bucket",
			Usage:  "s3 bucket to store charts for amazon storage backend",
			EnvVar: "STORAGE_AMAZON_BUCKET",
		},
	},
	"storage.amazon.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-prefix",
			Usage:  "prefix to store charts for --storage-amazon-bucket",
			EnvVar: "STORAGE_AMAZON_PREFIX",
		},
	},
	"storage.amazon.region": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-region",
			Usage:  "region of --storage-amazon-bucket",
			EnvVar: "STORAGE_AMAZON_REGION",
		},
	},
	"storage.amazon.endpoint": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-endpoint",
			Usage:  "alternative s3 endpoint",
			EnvVar: "STORAGE_AMAZON_ENDPOINT",
		},
	},
	"storage.amazon.sse": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-amazon-sse",
			Usage:  "server side encryption algorithm",
			EnvVar: "STORAGE_AMAZON_SSE",
		},
	},
	"storage.google.bucket": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-google-bucket",
			Usage:  "gcs bucket to store charts for google storage backend",
			EnvVar: "STORAGE_GOOGLE_BUCKET",
		},
	},
	"storage.google.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-google-prefix",
			Usage:  "prefix to store charts for --storage-google-bucket",
			EnvVar: "STORAGE_GOOGLE_PREFIX",
		},
	},
	"storage.oracle.bucket": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-oracle-bucket",
			Usage:  "ocs bucket to store charts for oracle cloud storage",
			EnvVar: "STORAGE_ORACLE_BUCKET",
		},
	},
	"storage.oracle.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-oracle-prefix",
			Usage:  "prefix to store charts for --storage-oracle-bucket",
			EnvVar: "STORAGE_ORACLE_PREFIX",
		},
	},
	"storage.oracle.region": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-oracle-region",
			Usage:  "region to store charts for --storage-oracle-bucket",
			EnvVar: "STORAGE_ORACLE_REGION",
		},
	},
	"storage.oracle.compartmentid": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-oracle-compartmentid",
			Usage:  "compartment ocid of --storage-oracle-bucket",
			EnvVar: "STORAGE_ORACLE_COMPARTMENTID",
		},
	},
	"storage.microsoft.container": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-microsoft-container",
			Usage:  "container to store charts for microsoft storage backend",
			EnvVar: "STORAGE_MICROSOFT_CONTAINER",
		},
	},
	"storage.microsoft.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-microsoft-prefix",
			Usage:  "prefix to store charts for --storage-microsoft-prefix",
			EnvVar: "STORAGE_MICROSOFT_PREFIX",
		},
	},
	"storage.alibaba.bucket": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-alibaba-bucket",
			Usage:  "OSS bucket to store charts for Alibaba Cloud storage backend",
			EnvVar: "STORAGE_ALIBABA_BUCKET",
		},
	},
	"storage.alibaba.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-alibaba-prefix",
			Usage:  "prefix to store charts for --storage-alibaba-cloud-bucket",
			EnvVar: "STORAGE_ALIBABA_PREFIX",
		},
	},
	"storage.alibaba.endpoint": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-alibaba-endpoint",
			Usage:  "OSS endpoint",
			EnvVar: "STORAGE_ALIBABA_ENDPOINT",
		},
	},
	"storage.alibaba.sse": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-alibaba-sse",
			Usage:  "server side encryption algorithm for Alibaba Cloud storage backend, AES256 or KMS",
			EnvVar: "STORAGE_ALIBABA_SSE",
		},
	},
	"storage.openstack.container": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-openstack-container",
			Usage:  "container to store charts for openstack storage backend",
			EnvVar: "STORAGE_OPENSTACK_CONTAINER",
		},
	},
	"storage.openstack.prefix": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-openstack-prefix",
			Usage:  "prefix to store charts for --storage-openstack-container",
			EnvVar: "STORAGE_OPENSTACK_PREFIX",
		},
	},
	"storage.openstack.region": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-openstack-region",
			Usage:  "region of --storage-openstack-container",
			EnvVar: "STORAGE_OPENSTACK_REGION",
		},
	},
	"storage.openstack.cacert": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "storage-openstack-cacert",
			Usage:  "path to a CA cert bundle for your openstack endpoint",
			EnvVar: "STORAGE_OPENSTACK_CACERT",
		},
	},
	"chartpostformfieldname": {
		Type:    stringType,
		Default: "chart",
		CLIFlag: cli.StringFlag{
			Name:   "chart-post-form-field-name",
			Usage:  "form field which will be queried for the chart file content",
			EnvVar: "CHART_POST_FORM_FIELD_NAME",
		},
	},
	"provpostformfieldname": {
		Type:    stringType,
		Default: "prov",
		CLIFlag: cli.StringFlag{
			Name:   "prov-post-form-field-name",
			Usage:  "form field which will be queried for the provenance file content",
			EnvVar: "PROV_POST_FORM_FIELD_NAME",
		},
	},
	"maxstorageobjects": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "max-storage-objects",
			Usage:  "maximum number of objects allowed in storage (per tenant)",
			EnvVar: "MAX_STORAGE_OBJECTS",
		},
	},
	"maxuploadsize": {
		Type:    intType,
		Default: 1024 * 1024 * 20, // 20MB, per Helm's limit
		CLIFlag: cli.IntFlag{
			Name:   "max-upload-size",
			Usage:  "max size of post body (in bytes)",
			EnvVar: "MAX_UPLOAD_SIZE",
			Value:  1024 * 1024 * 20,
		},
	},
	"indexlimit": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "index-limit",
			Usage:  "parallel scan limit for the repo indexer",
			EnvVar: "INDEX_LIMIT",
		},
	},
	"contextpath": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "context-path",
			Usage:  "base context path",
			EnvVar: "CONTEXT_PATH",
		},
	},
	"depth": {
		Type:    intType,
		Default: 0,
		CLIFlag: cli.IntFlag{
			Name:   "depth",
			Usage:  "levels of nested repos for multitenancy",
			EnvVar: "DEPTH",
		},
	},
	"bearerauth": {
		Type:    boolType,
		Default: false,
		CLIFlag: cli.BoolFlag{
			Name:   "bearer-auth",
			Usage:  "enable bearer auth",
			EnvVar: "BEARER_AUTH",
		},
	},
	"authtype": {
		Type:    stringType,
		Default: "token",
		CLIFlag: cli.StringFlag{
			Name:   "auth-type",
			Usage:  "type of auth (currently only supports token)",
			EnvVar: "AUTH_TYPE",
		},
	},
	"authrealm": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "auth-realm",
			Usage:  "authorization server url",
			EnvVar: "AUTH_REALM",
		},
	},
	"authservice": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "auth-service",
			Usage:  "authorization server service name",
			EnvVar: "AUTH_SERVICE",
		},
	},
	"authissuer": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "auth-issuer",
			Usage:  "authorization server name",
			EnvVar: "AUTH_ISSUER",
		},
	},
	"authcertpath": {
		Type:    stringType,
		Default: "",
		CLIFlag: cli.StringFlag{
			Name:   "auth-cert-path",
			Usage:  "path to authorization server public pem file",
			EnvVar: "AUTH_CERT_PATH",
		},
	},
}

func populateCLIFlags() {
	CLIFlags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "chartmuseum configuration file",
			EnvVar: "CONFIG",
		},
	}
	for _, configVar := range configVars {
		if flag := configVar.CLIFlag; flag != nil {
			CLIFlags = append(CLIFlags, flag)
		}
	}
}

func init() {
	populateCLIFlags()
}
