package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorManagersItemTagsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers\{hypervisor-id}\tags
type V1beta1HypervisorManagersItemTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorManagersItemTagsRequestBuilderGetQueryParameters list all the tag resources in a registered hypervisor manager.
type V1beta1HypervisorManagersItemTagsRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of resources returned in the response.The returned set of resources must match the criteria in the filter query parameter.A comparision compares a property name to a literal. The comparisons supported are the following:* “eq” : Is a property equal to value. Valid for number, boolean and string properties.* “ne” : Is a property not equal to value. Valid for number, boolean and string properties.* “gt” : Is a property greater than a value. Valid for number or string timestamp properties.* “lt” : Is a property less than a value. Valid for number or string timestamp properties* “in” : Is a value in a property (that is an array of strings)Filters are supported on the following attributes:* state* createdAt* name* services* displayName* associatedResources/id* categoryName
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
// V1beta1HypervisorManagersItemTagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersItemTagsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1HypervisorManagersItemTagsRequestBuilderGetQueryParameters
}
// ByTagId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.hypervisorManagers.item.tags.item collection
// returns a *V1beta1HypervisorManagersItemTagsTagItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemTagsRequestBuilder) ByTagId(tagId string)(*V1beta1HypervisorManagersItemTagsTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tagId != "" {
        urlTplParams["tag%2Did"] = tagId
    }
    return NewV1beta1HypervisorManagersItemTagsTagItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1HypervisorManagersItemTagsRequestBuilderInternal instantiates a new V1beta1HypervisorManagersItemTagsRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemTagsRequestBuilder) {
    m := &V1beta1HypervisorManagersItemTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers/{hypervisor%2Did}/tags{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersItemTagsRequestBuilder instantiates a new V1beta1HypervisorManagersItemTagsRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersItemTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all the tag resources in a registered hypervisor manager.
// Deprecated: This method is obsolete. Use GetAsTagsGetResponse instead.
// returns a V1beta1HypervisorManagersItemTagsResponseable when successful
// returns a V1beta1HypervisorManagersItemTags400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagersItemTags401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemTags403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemTags500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemTagsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemTagsRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemTagsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagersItemTags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagersItemTags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemTags403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemTags500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemTagsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemTagsResponseable), nil
}
// GetAsTagsGetResponse list all the tag resources in a registered hypervisor manager.
// returns a V1beta1HypervisorManagersItemTagsGetResponseable when successful
// returns a V1beta1HypervisorManagersItemTags400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagersItemTags401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemTags403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemTags500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemTagsRequestBuilder) GetAsTagsGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemTagsRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemTagsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagersItemTags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagersItemTags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemTags403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemTags500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemTagsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemTagsGetResponseable), nil
}
// ToGetRequestInformation list all the tag resources in a registered hypervisor manager.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersItemTagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemTagsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1HypervisorManagersItemTagsRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemTagsRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersItemTagsRequestBuilder) {
    return NewV1beta1HypervisorManagersItemTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
