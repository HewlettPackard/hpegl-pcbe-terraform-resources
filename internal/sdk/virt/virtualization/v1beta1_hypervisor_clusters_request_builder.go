package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorClustersRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-clusters
type V1beta1HypervisorClustersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorClustersRequestBuilderGetQueryParameters list all the hypervisors clusters across registered hypervisor managers.
type V1beta1HypervisorClustersRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of resources returned in the response.The returned set of resources must match the criteria in the filter query parameter.A comparison compares a property name to a literal. The following comparisons are supported:* “eq” : Is a property equal to value. Valid for number, boolean and string properties.* “ne” : Is a property not equal to value. Valid for number, boolean and string properties.* “gt” : Is a property greater than a value. Valid for number or string timestamp properties.* “lt” : Is a property less than a value. Valid for number or string timestamp properties* “in” : Is a value in a property (that is an array of strings)Filters are supported on the following attributes:* clusterType* appType* id* state* hypervisorManagerInfo/name* hypervisorManagerInfo/displayName* hypervisorManagerInfo/id* status* createdAt* hciClusterUuid* name* services* displayName* appInfo/vmware/moref
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
// V1beta1HypervisorClustersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorClustersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1HypervisorClustersRequestBuilderGetQueryParameters
}
// ByClusterId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.hypervisorClusters.item collection
// returns a *V1beta1HypervisorClustersClusterItemRequestBuilder when successful
func (m *V1beta1HypervisorClustersRequestBuilder) ByClusterId(clusterId string)(*V1beta1HypervisorClustersClusterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if clusterId != "" {
        urlTplParams["cluster%2Did"] = clusterId
    }
    return NewV1beta1HypervisorClustersClusterItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1HypervisorClustersRequestBuilderInternal instantiates a new V1beta1HypervisorClustersRequestBuilder and sets the default values.
func NewV1beta1HypervisorClustersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorClustersRequestBuilder) {
    m := &V1beta1HypervisorClustersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-clusters{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorClustersRequestBuilder instantiates a new V1beta1HypervisorClustersRequestBuilder and sets the default values.
func NewV1beta1HypervisorClustersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorClustersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorClustersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all the hypervisors clusters across registered hypervisor managers.
// Deprecated: This method is obsolete. Use GetAsHypervisorClustersGetResponse instead.
// returns a V1beta1HypervisorClustersResponseable when successful
// returns a V1beta1HypervisorClusters400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorClusters401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorClusters403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorClusters500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorClustersRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorClustersRequestBuilderGetRequestConfiguration)(V1beta1HypervisorClustersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorClusters400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorClusters401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorClusters403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorClusters500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorClustersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorClustersResponseable), nil
}
// GetAsHypervisorClustersGetResponse list all the hypervisors clusters across registered hypervisor managers.
// returns a V1beta1HypervisorClustersGetResponseable when successful
// returns a V1beta1HypervisorClusters400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorClusters401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorClusters403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorClusters500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorClustersRequestBuilder) GetAsHypervisorClustersGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorClustersRequestBuilderGetRequestConfiguration)(V1beta1HypervisorClustersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorClusters400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorClusters401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorClusters403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorClusters500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorClustersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorClustersGetResponseable), nil
}
// ToGetRequestInformation list all the hypervisors clusters across registered hypervisor managers.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorClustersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorClustersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1HypervisorClustersRequestBuilder when successful
func (m *V1beta1HypervisorClustersRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorClustersRequestBuilder) {
    return NewV1beta1HypervisorClustersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
