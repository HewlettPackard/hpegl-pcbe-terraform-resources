package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers\{hypervisor-id}\hypervisor-library-images
type V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderGetQueryParameters list all virtual machine images from the hypervisor library.
type V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of resources returned in the response.The returned set of resources must match the criteria in the filter query parameter.A comparison compares a property name to a literal. The following comparisons are supported:* “eq” : Is a property equal to value. Valid for number, boolean and string properties.* “gt” : Is a property greater than a value. Valid for number or string timestamp properties.* “lt” : Is a property less than a value. Valid for number or string timestamp properties* “in” : Is a value in a property (that is an array of strings)Filters are supported on the following attributes:* fileType* name* services* sizeInBytes* subscribed
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
// V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderGetQueryParameters
}
// ByHypervisorLibraryImageId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.hypervisorManagers.item.hypervisorLibraryImages.item collection
// returns a *V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder) ByHypervisorLibraryImageId(hypervisorLibraryImageId string)(*V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hypervisorLibraryImageId != "" {
        urlTplParams["hypervisor%2Dlibrary%2Dimage%2Did"] = hypervisorLibraryImageId
    }
    return NewV1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderInternal instantiates a new V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder) {
    m := &V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers/{hypervisor%2Did}/hypervisor-library-images{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder instantiates a new V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all virtual machine images from the hypervisor library.
// Deprecated: This method is obsolete. Use GetAsHypervisorLibraryImagesGetResponse instead.
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesResponseable when successful
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImages400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImages401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImages403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImages500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemHypervisorLibraryImagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagersItemHypervisorLibraryImages400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagersItemHypervisorLibraryImages401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemHypervisorLibraryImages403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemHypervisorLibraryImages500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemHypervisorLibraryImagesResponseable), nil
}
// GetAsHypervisorLibraryImagesGetResponse list all virtual machine images from the hypervisor library.
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponseable when successful
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImages400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImages401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImages403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImages500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder) GetAsHypervisorLibraryImagesGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagersItemHypervisorLibraryImages400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagersItemHypervisorLibraryImages401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemHypervisorLibraryImages403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemHypervisorLibraryImages500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponseable), nil
}
// ToGetRequestInformation list all virtual machine images from the hypervisor library.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder) {
    return NewV1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
