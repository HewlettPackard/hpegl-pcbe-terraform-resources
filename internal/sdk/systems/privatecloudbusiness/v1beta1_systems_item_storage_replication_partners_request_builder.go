package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemStorageReplicationPartnersRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\storage-replication-partners
type V1beta1SystemsItemStorageReplicationPartnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemStorageReplicationPartnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemStorageReplicationPartnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemStorageReplicationPartnersRequestBuilderInternal instantiates a new V1beta1SystemsItemStorageReplicationPartnersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemStorageReplicationPartnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemStorageReplicationPartnersRequestBuilder) {
    m := &V1beta1SystemsItemStorageReplicationPartnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/storage-replication-partners", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemStorageReplicationPartnersRequestBuilder instantiates a new V1beta1SystemsItemStorageReplicationPartnersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemStorageReplicationPartnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemStorageReplicationPartnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemStorageReplicationPartnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists replication partners of system's storage system in system details page view.Does not support Selection, Filtering and Sorting.
// Deprecated: This method is obsolete. Use GetAsStorageReplicationPartnersGetResponse instead.
// returns a V1beta1SystemsItemStorageReplicationPartnersResponseable when successful
// returns a V1beta1SystemsItemStorageReplicationPartners400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemStorageReplicationPartners401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemStorageReplicationPartners404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemStorageReplicationPartners500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemStorageReplicationPartnersRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemsItemStorageReplicationPartnersRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemStorageReplicationPartnersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemStorageReplicationPartners400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemStorageReplicationPartners401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemStorageReplicationPartners404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemStorageReplicationPartners500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemStorageReplicationPartnersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemStorageReplicationPartnersResponseable), nil
}
// GetAsStorageReplicationPartnersGetResponse lists replication partners of system's storage system in system details page view.Does not support Selection, Filtering and Sorting.
// returns a V1beta1SystemsItemStorageReplicationPartnersGetResponseable when successful
// returns a V1beta1SystemsItemStorageReplicationPartners400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemStorageReplicationPartners401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemStorageReplicationPartners404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemStorageReplicationPartners500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemStorageReplicationPartnersRequestBuilder) GetAsStorageReplicationPartnersGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemStorageReplicationPartnersRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemStorageReplicationPartnersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemStorageReplicationPartners400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemStorageReplicationPartners401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemStorageReplicationPartners404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemStorageReplicationPartners500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemStorageReplicationPartnersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemStorageReplicationPartnersGetResponseable), nil
}
// ToGetRequestInformation lists replication partners of system's storage system in system details page view.Does not support Selection, Filtering and Sorting.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemStorageReplicationPartnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SystemsItemStorageReplicationPartnersRequestBuilder when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemStorageReplicationPartnersRequestBuilder) {
    return NewV1beta1SystemsItemStorageReplicationPartnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
