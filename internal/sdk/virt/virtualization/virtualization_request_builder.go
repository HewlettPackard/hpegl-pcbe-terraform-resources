package virtualization

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// VirtualizationRequestBuilder builds and executes requests for operations under \virtualization
type VirtualizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewVirtualizationRequestBuilderInternal instantiates a new VirtualizationRequestBuilder and sets the default values.
func NewVirtualizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualizationRequestBuilder) {
    m := &VirtualizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization", pathParameters),
    }
    return m
}
// NewVirtualizationRequestBuilder instantiates a new VirtualizationRequestBuilder and sets the default values.
func NewVirtualizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualizationRequestBuilderInternal(urlParams, requestAdapter)
}
// V1beta1 the v1beta1 property
// returns a *V1beta1RequestBuilder when successful
func (m *VirtualizationRequestBuilder) V1beta1()(*V1beta1RequestBuilder) {
    return NewV1beta1RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
