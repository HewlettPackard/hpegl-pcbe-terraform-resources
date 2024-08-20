package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines
type V1beta1VirtualMachinesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesRequestBuilderGetQueryParameters list all the virtual machines across registered hypervisor managers.
type V1beta1VirtualMachinesRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of resources returned in the response.The returned set of resources must match the criteria in the filter query parameter.A comparison compares a property name to a literal. The following comparisons are supported:* “eq” : Is a property equal to value. Valid for number, boolean and string properties.* “ne” : Is a property not equal to value. Valid for number, boolean and string properties.* “gt” : Is a property greater than a value. Valid for number or string timestamp properties.* “lt” : Is a property less than a value. Valid for number or string timestamp properties* “in” : Is a value in a property (that is an array of strings)Filters are supported on the following attributes:* status* state* appType* hypervisorManagerInfo/name* hypervisorManagerInfo/displayName* hypervisorManagerInfo/id* hostInfo/id* hostInfo/name* hostInfo/displayName* clusterInfo/id* clusterInfo/name* clusterInfo/displayName* protectionJobInfo/protectionPolicyInfo/id* protectionJobInfo/protectionPolicyInfo/name* vmProtectionGroupsInfo/id* vmProtectionGroupsInfo/name* appInfo/vmware/type* appInfo/vmware/moref* appInfo/vmware/datastoresInfo/id* appInfo/vmware/datastoresInfo/name* appInfo/vmware/datastoresInfo/displayName* appInfo/vmware/datacenterInfo/displayName* appInfo/vmware/datacenterInfo/name* appInfo/vmware/datacenterInfo/id* volumesInfo/id* volumesInfo/storageSystemInfo/id* volumesInfo/storageSystemInfo/serialNumber* volumesInfo/storageSystemInfo/name* volumesInfo/storageSystemInfo/vendorName* volumesInfo/storageFolderInfo/id* volumesInfo/storageFolderInfo/name* volumesInfo/storagePoolInfo/id* volumesInfo/storagePoolInfo/name* powerState* createdAt* name* services* allowedOperations* hciClusterUuid* id* type* capacityInBytes* computeInfo/numCpuCores* computeInfo/memorySizeInMib* vmPerfMetricInfo/cpuAllocatedInMhz* vmPerfMetricInfo/cpuUsedInMhz* vmPerfMetricInfo/storageAllocatedInKb* vmPerfMetricInfo/storageUsedInBytes* vmPerfMetricInfo/memoryAllocatedInMb* vmPerfMetricInfo/memoryUsedInMb* vmPerfMetricInfo/totalReadIops* vmPerfMetricInfo/totalWriteIops* vmPerfMetricInfo/averageReadLatency* vmPerfMetricInfo/averageWriteLatency* displayName* protectionStatus* recoveryPointsExist* vclsVm
    Filter *string `uriparametername:"filter"`
    // The numbers of items to return
    Limit *int32 `uriparametername:"limit"`
    // The number of items to skip before starting to collect the result set
    Offset *int32 `uriparametername:"offset"`
    // The select query parameter is used to limit the properties returned with a resource or collection-level GET.Multiple properties can be listed to be returned. The server must only return the set of properties requested by the client.The property “select” is the name of the select query parameter; its value is the list of properties to return separated by commas.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a directionindicator ("asc" or "desc"). If no direction indicator is specified, thedefault order is ascending.
    Sort *string `uriparametername:"sort"`
}
// V1beta1VirtualMachinesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1VirtualMachinesRequestBuilderGetQueryParameters
}
// V1beta1VirtualMachinesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.virtualMachines.item collection
// returns a *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder when successful
func (m *V1beta1VirtualMachinesRequestBuilder) ById(id string)(*V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1VirtualMachinesVirtualMachinesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1VirtualMachinesRequestBuilderInternal instantiates a new V1beta1VirtualMachinesRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesRequestBuilder) {
    m := &V1beta1VirtualMachinesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesRequestBuilder instantiates a new V1beta1VirtualMachinesRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all the virtual machines across registered hypervisor managers.
// Deprecated: This method is obsolete. Use GetAsVirtualMachinesGetResponse instead.
// returns a V1beta1VirtualMachinesResponseable when successful
// returns a V1beta1VirtualMachines400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachines401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachines403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachines500Error error when the service returns a 500 status code
func (m *V1beta1VirtualMachinesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesRequestBuilderGetRequestConfiguration)(V1beta1VirtualMachinesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachines400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachines401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachines403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachines500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesResponseable), nil
}
// GetAsVirtualMachinesGetResponse list all the virtual machines across registered hypervisor managers.
// returns a V1beta1VirtualMachinesGetResponseable when successful
// returns a V1beta1VirtualMachines400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachines401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachines403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachines500Error error when the service returns a 500 status code
func (m *V1beta1VirtualMachinesRequestBuilder) GetAsVirtualMachinesGetResponse(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesRequestBuilderGetRequestConfiguration)(V1beta1VirtualMachinesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachines400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachines401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachines403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachines500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesGetResponseable), nil
}
// Migrate the migrate property
// returns a *V1beta1VirtualMachinesMigrateRequestBuilder when successful
func (m *V1beta1VirtualMachinesRequestBuilder) Migrate()(*V1beta1VirtualMachinesMigrateRequestBuilder) {
    return NewV1beta1VirtualMachinesMigrateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post deploys one or more virtual machines in HCI environment with specified template and storage provisioning policy.
// Deprecated: This method is obsolete. Use PostAsVirtualMachinesPostResponse instead.
// returns a V1beta1VirtualMachinesResponseable when successful
// returns a V1beta1VirtualMachines400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachines401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachines403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachines500Error error when the service returns a 500 status code
func (m *V1beta1VirtualMachinesRequestBuilder) Post(ctx context.Context, body V1beta1VirtualMachinesPostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachines400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachines401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachines403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachines500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesResponseable), nil
}
// PostAsVirtualMachinesPostResponse deploys one or more virtual machines in HCI environment with specified template and storage provisioning policy.
// returns a V1beta1VirtualMachinesPostResponseable when successful
// returns a V1beta1VirtualMachines400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachines401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachines403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachines500Error error when the service returns a 500 status code
func (m *V1beta1VirtualMachinesRequestBuilder) PostAsVirtualMachinesPostResponse(ctx context.Context, body V1beta1VirtualMachinesPostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachines400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachines401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachines403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachines500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesPostResponseable), nil
}
// ToGetRequestInformation list all the virtual machines across registered hypervisor managers.
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation deploys one or more virtual machines in HCI environment with specified template and storage provisioning policy.
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1VirtualMachinesPostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1VirtualMachinesRequestBuilder when successful
func (m *V1beta1VirtualMachinesRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesRequestBuilder) {
    return NewV1beta1VirtualMachinesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
