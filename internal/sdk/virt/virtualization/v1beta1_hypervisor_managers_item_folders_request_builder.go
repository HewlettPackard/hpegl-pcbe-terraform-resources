package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorManagersItemFoldersRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers\{hypervisor-id}\folders
type V1beta1HypervisorManagersItemFoldersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorManagersItemFoldersRequestBuilderGetQueryParameters list all the hypervisors folders in a registered hypervisor manager.
type V1beta1HypervisorManagersItemFoldersRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of resources returned in the response.The returned set of resources must match the criteria in the filter query parameter.A comparison compares a property name to a literal. The following comparisons are supported:* “eq” : Is a property equal to value. Valid for number, boolean and string properties.* “ne” : Is a property not equal to value. Valid for number, boolean and string properties.* “gt” : Is a property greater than a value. Valid for number or string timestamp properties.* “lt” : Is a property less than a value. Valid for number or string timestamp properties* “in” : Is a value in a property (that is an array of strings)Filters are supported on the following attributes:* folderType* createdAt* name* services* displayName
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
// V1beta1HypervisorManagersItemFoldersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersItemFoldersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1HypervisorManagersItemFoldersRequestBuilderGetQueryParameters
}
// ByFolderId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.hypervisorManagers.item.folders.item collection
// returns a *V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemFoldersRequestBuilder) ByFolderId(folderId string)(*V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if folderId != "" {
        urlTplParams["folder%2Did"] = folderId
    }
    return NewV1beta1HypervisorManagersItemFoldersFolderItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1HypervisorManagersItemFoldersRequestBuilderInternal instantiates a new V1beta1HypervisorManagersItemFoldersRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemFoldersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemFoldersRequestBuilder) {
    m := &V1beta1HypervisorManagersItemFoldersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers/{hypervisor%2Did}/folders{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersItemFoldersRequestBuilder instantiates a new V1beta1HypervisorManagersItemFoldersRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemFoldersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemFoldersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersItemFoldersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all the hypervisors folders in a registered hypervisor manager.
// Deprecated: This method is obsolete. Use GetAsFoldersGetResponse instead.
// returns a V1beta1HypervisorManagersItemFoldersResponseable when successful
// returns a V1beta1HypervisorManagersItemFolders400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagersItemFolders401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemFolders403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemFolders500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemFoldersRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemFoldersRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemFoldersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagersItemFolders400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagersItemFolders401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemFolders403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemFolders500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemFoldersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemFoldersResponseable), nil
}
// GetAsFoldersGetResponse list all the hypervisors folders in a registered hypervisor manager.
// returns a V1beta1HypervisorManagersItemFoldersGetResponseable when successful
// returns a V1beta1HypervisorManagersItemFolders400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagersItemFolders401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemFolders403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemFolders500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemFoldersRequestBuilder) GetAsFoldersGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemFoldersRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemFoldersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagersItemFolders400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagersItemFolders401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemFolders403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemFolders500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemFoldersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemFoldersGetResponseable), nil
}
// ToGetRequestInformation list all the hypervisors folders in a registered hypervisor manager.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersItemFoldersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemFoldersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1HypervisorManagersItemFoldersRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemFoldersRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersItemFoldersRequestBuilder) {
    return NewV1beta1HypervisorManagersItemFoldersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
