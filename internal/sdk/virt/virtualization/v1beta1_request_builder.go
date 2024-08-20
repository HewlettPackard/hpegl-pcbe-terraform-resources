package virtualization

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1RequestBuilder builds and executes requests for operations under \virtualization\v1beta1
type V1beta1RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1beta1RequestBuilderInternal instantiates a new V1beta1RequestBuilder and sets the default values.
func NewV1beta1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1RequestBuilder) {
    m := &V1beta1RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1", pathParameters),
    }
    return m
}
// NewV1beta1RequestBuilder instantiates a new V1beta1RequestBuilder and sets the default values.
func NewV1beta1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1RequestBuilderInternal(urlParams, requestAdapter)
}
// CspAccounts the cspAccounts property
// returns a *V1beta1CspAccountsRequestBuilder when successful
func (m *V1beta1RequestBuilder) CspAccounts()(*V1beta1CspAccountsRequestBuilder) {
    return NewV1beta1CspAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspMachineImages the cspMachineImages property
// returns a *V1beta1CspMachineImagesRequestBuilder when successful
func (m *V1beta1RequestBuilder) CspMachineImages()(*V1beta1CspMachineImagesRequestBuilder) {
    return NewV1beta1CspMachineImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspMachineInstances the cspMachineInstances property
// returns a *V1beta1CspMachineInstancesRequestBuilder when successful
func (m *V1beta1RequestBuilder) CspMachineInstances()(*V1beta1CspMachineInstancesRequestBuilder) {
    return NewV1beta1CspMachineInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspMachineInstanceTypes the cspMachineInstanceTypes property
// returns a *V1beta1CspMachineInstanceTypesRequestBuilder when successful
func (m *V1beta1RequestBuilder) CspMachineInstanceTypes()(*V1beta1CspMachineInstanceTypesRequestBuilder) {
    return NewV1beta1CspMachineInstanceTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspRdsInstances the cspRdsInstances property
// returns a *V1beta1CspRdsInstancesRequestBuilder when successful
func (m *V1beta1RequestBuilder) CspRdsInstances()(*V1beta1CspRdsInstancesRequestBuilder) {
    return NewV1beta1CspRdsInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspVolumes the cspVolumes property
// returns a *V1beta1CspVolumesRequestBuilder when successful
func (m *V1beta1RequestBuilder) CspVolumes()(*V1beta1CspVolumesRequestBuilder) {
    return NewV1beta1CspVolumesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Datastores the datastores property
// returns a *V1beta1DatastoresRequestBuilder when successful
func (m *V1beta1RequestBuilder) Datastores()(*V1beta1DatastoresRequestBuilder) {
    return NewV1beta1DatastoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HypervisorClusters the hypervisorClusters property
// returns a *V1beta1HypervisorClustersRequestBuilder when successful
func (m *V1beta1RequestBuilder) HypervisorClusters()(*V1beta1HypervisorClustersRequestBuilder) {
    return NewV1beta1HypervisorClustersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HypervisorHosts the hypervisorHosts property
// returns a *V1beta1HypervisorHostsRequestBuilder when successful
func (m *V1beta1RequestBuilder) HypervisorHosts()(*V1beta1HypervisorHostsRequestBuilder) {
    return NewV1beta1HypervisorHostsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HypervisorManagers the hypervisorManagers property
// returns a *V1beta1HypervisorManagersRequestBuilder when successful
func (m *V1beta1RequestBuilder) HypervisorManagers()(*V1beta1HypervisorManagersRequestBuilder) {
    return NewV1beta1HypervisorManagersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VirtualMachines the virtualMachines property
// returns a *V1beta1VirtualMachinesRequestBuilder when successful
func (m *V1beta1RequestBuilder) VirtualMachines()(*V1beta1VirtualMachinesRequestBuilder) {
    return NewV1beta1VirtualMachinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
