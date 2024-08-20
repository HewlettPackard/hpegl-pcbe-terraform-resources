package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\system-software-catalogs\{id}
type V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderGetQueryParameters returns the system software catalog for the specified id.Includes versions of the catalog and all constituent software components along with the End User License Agreementand a list of systems with update path to the corresponding catalog version.
type V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderGetQueryParameters struct {
    // A list of properties in the items collection to include in the response.
    Select *string `uriparametername:"select"`
}
// V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderGetQueryParameters
}
// NewV1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderInternal instantiates a new V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder and sets the default values.
func NewV1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder) {
    m := &V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/system-software-catalogs/{id}{?select*}", pathParameters),
    }
    return m
}
// NewV1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder instantiates a new V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder and sets the default values.
func NewV1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the system software catalog for the specified id.Includes versions of the catalog and all constituent software components along with the End User License Agreementand a list of systems with update path to the corresponding catalog version.
// Deprecated: This method is obsolete. Use GetAsSystemSoftwareCatalogsGetResponse instead.
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponseable when successful
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs400Error error when the service returns a 400 status code
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs401Error error when the service returns a 401 status code
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs404Error error when the service returns a 404 status code
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs500Error error when the service returns a 500 status code
func (m *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderGetRequestConfiguration)(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponseable), nil
}
// GetAsSystemSoftwareCatalogsGetResponse returns the system software catalog for the specified id.Includes versions of the catalog and all constituent software components along with the End User License Agreementand a list of systems with update path to the corresponding catalog version.
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseable when successful
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs400Error error when the service returns a 400 status code
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs401Error error when the service returns a 401 status code
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs404Error error when the service returns a 404 status code
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs500Error error when the service returns a 500 status code
func (m *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder) GetAsSystemSoftwareCatalogsGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderGetRequestConfiguration)(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogs500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseable), nil
}
// ToGetRequestInformation returns the system software catalog for the specified id.Includes versions of the catalog and all constituent software components along with the End User License Agreementand a list of systems with update path to the corresponding catalog version.
// returns a *RequestInformation when successful
func (m *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder when successful
func (m *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder) {
    return NewV1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
