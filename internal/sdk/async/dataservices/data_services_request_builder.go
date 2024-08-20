package dataservices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DataServicesRequestBuilder builds and executes requests for operations under \data-services
type DataServicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewDataServicesRequestBuilderInternal instantiates a new DataServicesRequestBuilder and sets the default values.
func NewDataServicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataServicesRequestBuilder) {
    m := &DataServicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services", pathParameters),
    }
    return m
}
// NewDataServicesRequestBuilder instantiates a new DataServicesRequestBuilder and sets the default values.
func NewDataServicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataServicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataServicesRequestBuilderInternal(urlParams, requestAdapter)
}
// V1beta1 the v1beta1 property
// returns a *V1beta1RequestBuilder when successful
func (m *DataServicesRequestBuilder) V1beta1()(*V1beta1RequestBuilder) {
    return NewV1beta1RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
