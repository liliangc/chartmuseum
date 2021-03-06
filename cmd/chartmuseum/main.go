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

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/helm/chartmuseum/pkg/cache"
	"github.com/helm/chartmuseum/pkg/chartmuseum"
	"github.com/helm/chartmuseum/pkg/config"
	"github.com/helm/chartmuseum/pkg/storage"

	"github.com/urfave/cli"
)

var (
	crash = log.Fatal

	newServer = chartmuseum.NewServer

	// Version is the semantic version (added at compile time)
	Version string

	// Revision is the git commit id (added at compile time)
	Revision string
)

func main() {
	app := cli.NewApp()
	app.Name = "ChartMuseum"
	app.Version = fmt.Sprintf("%s (build %s)", Version, Revision)
	app.Usage = "Helm Chart Repository with support for Amazon S3, Google Cloud Storage, Oracle Cloud Infrastructure Object Storage and Openstack"
	app.Action = cliHandler
	app.Flags = config.CLIFlags
	app.Run(os.Args)
}

func cliHandler(c *cli.Context) {
	conf := config.NewConfig()
	err := conf.UpdateFromCLIContext(c)
	if err != nil {
		crash(err)
	}

	backend := backendFromConfig(conf)
	store := storeFromConfig(conf)

	options := chartmuseum.ServerOptions{
		StorageBackend:         backend,
		ExternalCacheStore:     store,
		ChartURL:               conf.GetString("charturl"),
		TlsCert:                conf.GetString("tls.cert"),
		TlsKey:                 conf.GetString("tls.key"),
		Username:               conf.GetString("basicauth.user"),
		Password:               conf.GetString("basicauth.pass"),
		ChartPostFormFieldName: conf.GetString("chartpostformfieldname"),
		ProvPostFormFieldName:  conf.GetString("provpostformfieldname"),
		ContextPath:            conf.GetString("contextpath"),
		LogJSON:                conf.GetBool("logjson"),
		LogHealth:              conf.GetBool("loghealth"),
		Debug:                  conf.GetBool("debug"),
		EnableAPI:              !conf.GetBool("disableapi"),
		UseStatefiles:          !conf.GetBool("disablestatefiles"),
		AllowOverwrite:         conf.GetBool("allowoverwrite"),
		AllowForceOverwrite:    !conf.GetBool("disableforceoverwrite"),
		EnableMetrics:          !conf.GetBool("disablemetrics"),
		AnonymousGet:           conf.GetBool("authanonymousget"),
		GenIndex:               conf.GetBool("genindex"),
		MaxStorageObjects:      conf.GetInt("maxstorageobjects"),
		IndexLimit:             conf.GetInt("indexlimit"),
		Depth:                  conf.GetInt("depth"),
		MaxUploadSize:          conf.GetInt("maxuploadsize"),
		BearerAuth:             conf.GetBool("bearerauth"),
		AuthType:               conf.GetString("authtype"),
		AuthRealm:              conf.GetString("authrealm"),
		AuthService:            conf.GetString("authservice"),
		AuthIssuer:             conf.GetString("authissuer"),
		AuthCertPath:           conf.GetString("authcertpath"),
	}

	server, err := newServer(options)
	if err != nil {
		crash(err)
	}

	server.Listen(conf.GetInt("port"))
}

func backendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.backend"})

	var backend storage.Backend

	storageFlag := strings.ToLower(conf.GetString("storage.backend"))
	switch storageFlag {
	case "local":
		backend = localBackendFromConfig(conf)
	case "amazon":
		backend = amazonBackendFromConfig(conf)
	case "google":
		backend = googleBackendFromConfig(conf)
	case "oracle":
		backend = oracleBackendFromConfig(conf)
	case "microsoft":
		backend = microsoftBackendFromConfig(conf)
	case "alibaba":
		backend = alibabaBackendFromConfig(conf)
	case "openstack":
		backend = openstackBackendFromConfig(conf)
	default:
		crash("Unsupported storage backend: ", storageFlag)
	}

	return backend
}

func localBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.local.rootdir"})
	return storage.Backend(storage.NewLocalFilesystemBackend(
		conf.GetString("storage.local.rootdir"),
	))
}

func amazonBackendFromConfig(conf *config.Config) storage.Backend {
	// If using alternative s3 endpoint (e.g. Minio) default region to us-east-1
	if conf.GetString("storage.amazon.endpoint") != "" && conf.GetString("storage.amazon.region") == "" {
		conf.Set("storage.amazon.region", "us-east-1")
	}
	crashIfConfigMissingVars(conf, []string{"storage.amazon.bucket", "storage.amazon.region"})
	return storage.Backend(storage.NewAmazonS3Backend(
		conf.GetString("storage.amazon.bucket"),
		conf.GetString("storage.amazon.prefix"),
		conf.GetString("storage.amazon.region"),
		conf.GetString("storage.amazon.endpoint"),
		conf.GetString("storage.amazon.sse"),
	))
}

func googleBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.google.bucket"})
	return storage.Backend(storage.NewGoogleCSBackend(
		conf.GetString("storage.google.bucket"),
		conf.GetString("storage.google.prefix"),
	))
}

func oracleBackendFromConfig(conf *config.Config) storage.Backend {
        crashIfConfigMissingVars(conf, []string{"storage.oracle.bucket", "storage.oracle.compartmentid"})
        return storage.Backend(storage.NewOracleCSBackend(
		conf.GetString("storage.oracle.bucket"),
		conf.GetString("storage.oracle.prefix"),
		conf.GetString("storage.oracle.region"),
		conf.GetString("storage.oracle.compartmentid"),
        ))
}

func microsoftBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.microsoft.container"})
	return storage.Backend(storage.NewMicrosoftBlobBackend(
		conf.GetString("storage.microsoft.container"),
		conf.GetString("storage.microsoft.prefix"),
	))
}

func alibabaBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.alibaba.bucket"})
	return storage.Backend(storage.NewAlibabaCloudOSSBackend(
		conf.GetString("storage.alibaba.bucket"),
		conf.GetString("storage.alibaba.prefix"),
		conf.GetString("storage.alibaba.endpoint"),
		conf.GetString("storage.alibaba.sse"),
	))
}

func openstackBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.openstack.container", "storage.openstack.region"})
	return storage.Backend(storage.NewOpenstackOSBackend(
		conf.GetString("storage.openstack.container"),
		conf.GetString("storage.openstack.prefix"),
		conf.GetString("storage.openstack.region"),
		conf.GetString("storage.openstack.cacert"),
	))
}

func storeFromConfig(conf *config.Config) cache.Store {
	if conf.GetString("cache.store") == "" {
		return nil
	}

	var store cache.Store

	cacheFlag := strings.ToLower(conf.GetString("cache.store"))
	switch cacheFlag {
	case "redis":
		store = redisCacheFromConfig(conf)
	default:
		crash("Unsupported cache store: ", cacheFlag)
	}

	return store
}

func redisCacheFromConfig(conf *config.Config) cache.Store {
	crashIfConfigMissingVars(conf, []string{"cache.redis.addr"})
	return cache.Store(cache.NewRedisStore(
		conf.GetString("cache.redis.addr"),
		conf.GetString("cache.redis.password"),
		conf.GetInt("cache.redis.db"),
	))
}

func crashIfConfigMissingVars(conf *config.Config, vars []string) {
	missing := []string{}
	for _, v := range vars {
		if conf.GetString(v) == "" {
			flag := config.GetCLIFlagFromVarName(v)
			missing = append(missing, fmt.Sprintf("--%s", flag))
		}
	}
	if len(missing) > 0 {
		crash("Missing required flags(s): ", strings.Join(missing, ", "))
	}
}
