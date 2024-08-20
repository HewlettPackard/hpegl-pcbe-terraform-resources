package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1DatastoresRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\datastores
type V1beta1DatastoresRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DatastoresRequestBuilderGetQueryParameters get all datastores for registered hypervisor managers.
type V1beta1DatastoresRequestBuilderGetQueryParameters struct {
    // An expression to filter the results.A comparison compares a property name to a literal. The following comparisons are supported:* “eq” : Is a property equal to value. Valid for number, boolean and string properties.* “ne” : Is a property not equal to value. Valid for number, boolean and string properties.* “gt” : Is a property greater than a value. Valid for number or string timestamp properties.* “lt” : Is a property less than a value. Valid for number or string timestamp properties* “in” : Is a value in a property (that is an array of strings)Filters are supported on the following attributes:* status* state* appType* hypervisorManagerInfo/name* hypervisorManagerInfo/displayName* hypervisorManagerInfo/id* hostsInfo/id* hostsInfo/name* hostsInfo/displayName* clusterInfo/id* clusterInfo/name* clusterInfo/displayName* protectionJobInfo/protectionPolicyInfo/id* protectionJobInfo/protectionPolicyInfo/name* vmProtectionGroupsInfo/id* vmProtectionGroupsInfo/name* volumesInfo/id* volumesInfo/storageSystemInfo/id* volumesInfo/storageSystemInfo/serialNumber* volumesInfo/storageSystemInfo/name* volumesInfo/storageSystemInfo/vendorName* volumesInfo/storageSystemInfo/type* volumesInfo/storageFolderInfo/id* volumesInfo/storageFolderInfo/name* volumesInfo/storagePoolInfo/id* volumesInfo/storagePoolInfo/name* datastoreType* createdAt* name* services* allowedOperations* capacityInBytes* capacityFree* displayName* replicationInfo/name* replicationInfo/id* hciClusterUuid
    Filter *string `uriparametername:"filter"`
    // The maximum number of items to include in the response. Use offset in conjunction with limit for pagination.
    Limit *int32 `uriparametername:"limit"`
    // The number of items to skip before starting to collect the result set. Use offset in conjunction with limit for pagination.
    Offset *int32 `uriparametername:"offset"`
    // The select query parameter is used to limit the properties returned with a resource or collection-level GET.Multiple properties can be listed to be returned. The server must only return the set of properties requested by the client.The property “select” is the name of the select query parameter; its value is the list of properties to return separated by commas.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a directionindicator ("asc" or "desc"). 
    Sort *string `uriparametername:"sort"`
}
// V1beta1DatastoresRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DatastoresRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1DatastoresRequestBuilderGetQueryParameters
}
// V1beta1DatastoresRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DatastoresRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.datastores.item collection
// returns a *V1beta1DatastoresDatastoresItemRequestBuilder when successful
func (m *V1beta1DatastoresRequestBuilder) ById(id string)(*V1beta1DatastoresDatastoresItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1DatastoresDatastoresItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1DatastoresRequestBuilderInternal instantiates a new V1beta1DatastoresRequestBuilder and sets the default values.
func NewV1beta1DatastoresRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DatastoresRequestBuilder) {
    m := &V1beta1DatastoresRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/datastores{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1DatastoresRequestBuilder instantiates a new V1beta1DatastoresRequestBuilder and sets the default values.
func NewV1beta1DatastoresRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DatastoresRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DatastoresRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all datastores for registered hypervisor managers.
// Deprecated: This method is obsolete. Use GetAsDatastoresGetResponse instead.
// returns a V1beta1DatastoresResponseable when successful
// returns a V1beta1Datastores400Error error when the service returns a 400 status code
// returns a V1beta1Datastores401Error error when the service returns a 401 status code
// returns a V1beta1Datastores403Error error when the service returns a 403 status code
// returns a V1beta1Datastores500Error error when the service returns a 500 status code
func (m *V1beta1DatastoresRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1DatastoresRequestBuilderGetRequestConfiguration)(V1beta1DatastoresResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Datastores400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Datastores401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Datastores403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Datastores500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresResponseable), nil
}
// GetAsDatastoresGetResponse get all datastores for registered hypervisor managers.
// returns a V1beta1DatastoresGetResponseable when successful
// returns a V1beta1Datastores400Error error when the service returns a 400 status code
// returns a V1beta1Datastores401Error error when the service returns a 401 status code
// returns a V1beta1Datastores403Error error when the service returns a 403 status code
// returns a V1beta1Datastores500Error error when the service returns a 500 status code
func (m *V1beta1DatastoresRequestBuilder) GetAsDatastoresGetResponse(ctx context.Context, requestConfiguration *V1beta1DatastoresRequestBuilderGetRequestConfiguration)(V1beta1DatastoresGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Datastores400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Datastores401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Datastores403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Datastores500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresGetResponseable), nil
}
// Post create datastore
// Deprecated: This method is obsolete. Use PostAsDatastoresPostResponse instead.
// returns a V1beta1DatastoresResponseable when successful
// returns a V1beta1Datastores400Error error when the service returns a 400 status code
// returns a V1beta1Datastores401Error error when the service returns a 401 status code
// returns a V1beta1Datastores403Error error when the service returns a 403 status code
// returns a V1beta1Datastores500Error error when the service returns a 500 status code
func (m *V1beta1DatastoresRequestBuilder) Post(ctx context.Context, body V1beta1DatastoresPostRequestBodyable, requestConfiguration *V1beta1DatastoresRequestBuilderPostRequestConfiguration)(V1beta1DatastoresResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Datastores400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Datastores401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Datastores403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Datastores500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresResponseable), nil
}
// PostAsDatastoresPostResponse create datastore
// returns a V1beta1DatastoresPostResponseable when successful
// returns a V1beta1Datastores400Error error when the service returns a 400 status code
// returns a V1beta1Datastores401Error error when the service returns a 401 status code
// returns a V1beta1Datastores403Error error when the service returns a 403 status code
// returns a V1beta1Datastores500Error error when the service returns a 500 status code
func (m *V1beta1DatastoresRequestBuilder) PostAsDatastoresPostResponse(ctx context.Context, body V1beta1DatastoresPostRequestBodyable, requestConfiguration *V1beta1DatastoresRequestBuilderPostRequestConfiguration)(V1beta1DatastoresPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Datastores400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Datastores401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Datastores403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Datastores500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresPostResponseable), nil
}
// ToGetRequestInformation get all datastores for registered hypervisor managers.
// returns a *RequestInformation when successful
func (m *V1beta1DatastoresRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1DatastoresRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create datastore
// returns a *RequestInformation when successful
func (m *V1beta1DatastoresRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1DatastoresPostRequestBodyable, requestConfiguration *V1beta1DatastoresRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1DatastoresRequestBuilder when successful
func (m *V1beta1DatastoresRequestBuilder) WithUrl(rawUrl string)(*V1beta1DatastoresRequestBuilder) {
    return NewV1beta1DatastoresRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
