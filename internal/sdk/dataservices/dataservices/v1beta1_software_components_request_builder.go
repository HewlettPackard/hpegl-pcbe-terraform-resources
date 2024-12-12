package dataservices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1SoftwareComponentsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\software-components
type V1beta1SoftwareComponentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.softwareComponents.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder when successful
func (m *V1beta1SoftwareComponentsRequestBuilder) ById(id string)(*V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.softwareComponents.item collection
// returns a *V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder when successful
func (m *V1beta1SoftwareComponentsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SoftwareComponentsRequestBuilderInternal instantiates a new V1beta1SoftwareComponentsRequestBuilder and sets the default values.
func NewV1beta1SoftwareComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareComponentsRequestBuilder) {
    m := &V1beta1SoftwareComponentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/software-components", pathParameters),
    }
    return m
}
// NewV1beta1SoftwareComponentsRequestBuilder instantiates a new V1beta1SoftwareComponentsRequestBuilder and sets the default values.
func NewV1beta1SoftwareComponentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareComponentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SoftwareComponentsRequestBuilderInternal(urlParams, requestAdapter)
}
