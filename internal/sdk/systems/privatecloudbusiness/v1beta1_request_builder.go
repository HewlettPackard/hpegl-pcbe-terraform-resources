package privatecloudbusiness

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1RequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1
type V1beta1RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1beta1RequestBuilderInternal instantiates a new V1beta1RequestBuilder and sets the default values.
func NewV1beta1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1RequestBuilder) {
    m := &V1beta1RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1", pathParameters),
    }
    return m
}
// NewV1beta1RequestBuilder instantiates a new V1beta1RequestBuilder and sets the default values.
func NewV1beta1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1RequestBuilderInternal(urlParams, requestAdapter)
}
// Systems the systems property
// returns a *V1beta1SystemsRequestBuilder when successful
func (m *V1beta1RequestBuilder) Systems()(*V1beta1SystemsRequestBuilder) {
    return NewV1beta1SystemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SystemSoftwareCatalogs the systemSoftwareCatalogs property
// returns a *V1beta1SystemSoftwareCatalogsRequestBuilder when successful
func (m *V1beta1RequestBuilder) SystemSoftwareCatalogs()(*V1beta1SystemSoftwareCatalogsRequestBuilder) {
    return NewV1beta1SystemSoftwareCatalogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
