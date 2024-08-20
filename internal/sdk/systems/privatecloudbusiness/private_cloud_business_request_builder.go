package privatecloudbusiness

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrivateCloudBusinessRequestBuilder builds and executes requests for operations under \private-cloud-business
type PrivateCloudBusinessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewPrivateCloudBusinessRequestBuilderInternal instantiates a new PrivateCloudBusinessRequestBuilder and sets the default values.
func NewPrivateCloudBusinessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivateCloudBusinessRequestBuilder) {
    m := &PrivateCloudBusinessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business", pathParameters),
    }
    return m
}
// NewPrivateCloudBusinessRequestBuilder instantiates a new PrivateCloudBusinessRequestBuilder and sets the default values.
func NewPrivateCloudBusinessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivateCloudBusinessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivateCloudBusinessRequestBuilderInternal(urlParams, requestAdapter)
}
// V1beta1 the v1beta1 property
// returns a *V1beta1RequestBuilder when successful
func (m *PrivateCloudBusinessRequestBuilder) V1beta1()(*V1beta1RequestBuilder) {
    return NewV1beta1RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
