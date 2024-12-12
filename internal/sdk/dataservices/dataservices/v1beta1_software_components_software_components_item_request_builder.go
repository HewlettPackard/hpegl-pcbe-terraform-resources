package dataservices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\software-components\{id}
type V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilderInternal instantiates a new V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder and sets the default values.
func NewV1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder) {
    m := &V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/software-components/{id}", pathParameters),
    }
    return m
}
// NewV1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder instantiates a new V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder and sets the default values.
func NewV1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// InstallRelease the installRelease property
// returns a *V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder when successful
func (m *V1beta1SoftwareComponentsSoftwareComponentsItemRequestBuilder) InstallRelease()(*V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder) {
    return NewV1beta1SoftwareComponentsItemInstallReleaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
