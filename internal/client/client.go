package client

import (
	"context"
	"time"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/auth"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/defaults"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/async"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/systems"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/microsoft/kiota-http-go"
)

type PCBeClient struct {
	Config Config
}

type Config struct {
	Host         string
	Token        string
	HTTPDump     bool
	MaxPolls     int32
	PollInterval float32
	HTTPTimeout  time.Duration
}

func (c *PCBeClient) NewAsyncClient(ctx context.Context) (*async.ApiClient, error) {
	var middlewares []nethttplibrary.Middleware

	tp := auth.NewPcbeAccessTokenProvider(c.Config.Token)
	authProvider := authentication.NewBaseBearerTokenAuthenticationProvider(tp)
	observeOpts := nethttplibrary.ObservabilityOptions{}
	headerOpts := nethttplibrary.NewHeadersInspectionOptions()
	headerOpts.InspectResponseHeaders = true
	headerOptsHandler := nethttplibrary.NewHeadersInspectionHandlerWithOptions(*headerOpts)
	middlewares = append(middlewares, headerOptsHandler)
	httpclient := nethttplibrary.GetDefaultClient(middlewares...)
	httpclient.Timeout = c.Config.HTTPTimeout

	adapter, err := nethttplibrary.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClientAndObservabilityOptions(authProvider, nil, nil, httpclient, observeOpts) // nolint: lll
	if err != nil {
		return nil, err
	}

	adapter.SetBaseUrl(c.Config.Host)
	asyncClient := async.NewApiClient(adapter)

	return asyncClient, nil
}

func (c *PCBeClient) NewVirtClient(
	ctx context.Context,
) (*virt.ApiClient, *nethttplibrary.HeadersInspectionOptions, error) {
	var middlewares []nethttplibrary.Middleware

	tp := auth.NewPcbeAccessTokenProvider(c.Config.Token)
	authProvider := authentication.NewBaseBearerTokenAuthenticationProvider(tp)
	observeOpts := nethttplibrary.ObservabilityOptions{}
	headerOpts := nethttplibrary.NewHeadersInspectionOptions()
	headerOpts.InspectResponseHeaders = true
	headerOptsHandler := nethttplibrary.NewHeadersInspectionHandlerWithOptions(*headerOpts)
	middlewares = append(middlewares, headerOptsHandler)
	httpclient := nethttplibrary.GetDefaultClient(middlewares...)
	httpclient.Timeout = c.Config.HTTPTimeout

	adapter, err := nethttplibrary.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClientAndObservabilityOptions(authProvider, nil, nil, httpclient, observeOpts) // nolint: lll
	if err != nil {
		return nil, nil, err
	}

	adapter.SetBaseUrl(c.Config.Host)
	virtClient := virt.NewApiClient(adapter)

	return virtClient, headerOpts, nil
}

func (c *PCBeClient) NewSysClient(
	ctx context.Context,
) (*systems.ApiClient, *nethttplibrary.HeadersInspectionOptions, error) {
	var middlewares []nethttplibrary.Middleware

	tp := auth.NewPcbeAccessTokenProvider(c.Config.Token)
	authProvider := authentication.NewBaseBearerTokenAuthenticationProvider(tp)
	observeOpts := nethttplibrary.ObservabilityOptions{}
	headerOpts := nethttplibrary.NewHeadersInspectionOptions()
	headerOpts.InspectResponseHeaders = true
	headerOptsHandler := nethttplibrary.NewHeadersInspectionHandlerWithOptions(*headerOpts)
	middlewares = append(middlewares, headerOptsHandler)
	httpclient := nethttplibrary.GetDefaultClient(middlewares...)
	httpclient.Timeout = c.Config.HTTPTimeout

	adapter, err := nethttplibrary.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClientAndObservabilityOptions(authProvider, nil, nil, httpclient, observeOpts) // nolint: lll
	if err != nil {
		return nil, nil, err
	}

	adapter.SetBaseUrl(c.Config.Host)
	sysClient := systems.NewApiClient(adapter)

	return sysClient, headerOpts, nil
}

func NewPCBeClient(ctx context.Context, cfg Config) (*PCBeClient, error) {
	if cfg.MaxPolls == 0 {
		cfg.MaxPolls = defaults.MaxPolls
	}
	if cfg.PollInterval == 0.0 {
		cfg.PollInterval = defaults.PollInterval
	}
	client := &PCBeClient{
		Config: cfg,
	}

	return client, nil
}
