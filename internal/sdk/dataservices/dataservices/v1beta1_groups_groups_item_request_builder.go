package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1GroupsGroupsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\groups\{id}
type V1beta1GroupsGroupsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1GroupsGroupsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1GroupsGroupsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1GroupsGroupsItemRequestBuilderGetQueryParameters returns group for a customer
type V1beta1GroupsGroupsItemRequestBuilderGetQueryParameters struct {
    // A list of properties to include in the response. Service currently only supports specification of all fields.
    Select *string `uriparametername:"select"`
}
// V1beta1GroupsGroupsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1GroupsGroupsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1GroupsGroupsItemRequestBuilderGetQueryParameters
}
// V1beta1GroupsGroupsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1GroupsGroupsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociatedResources the associatedResources property
// returns a *V1beta1GroupsItemAssociatedResourcesRequestBuilder when successful
func (m *V1beta1GroupsGroupsItemRequestBuilder) AssociatedResources()(*V1beta1GroupsItemAssociatedResourcesRequestBuilder) {
    return NewV1beta1GroupsItemAssociatedResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1GroupsGroupsItemRequestBuilderInternal instantiates a new V1beta1GroupsGroupsItemRequestBuilder and sets the default values.
func NewV1beta1GroupsGroupsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1GroupsGroupsItemRequestBuilder) {
    m := &V1beta1GroupsGroupsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/groups/{id}{?select*}", pathParameters),
    }
    return m
}
// NewV1beta1GroupsGroupsItemRequestBuilder instantiates a new V1beta1GroupsGroupsItemRequestBuilder and sets the default values.
func NewV1beta1GroupsGroupsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1GroupsGroupsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1GroupsGroupsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a Specific Group.
// returns a V1beta1GroupsItemGroups403Error error when the service returns a 403 status code
// returns a V1beta1GroupsItemGroups404Error error when the service returns a 404 status code
// returns a V1beta1GroupsItemGroups409Error error when the service returns a 409 status code
// returns a V1beta1GroupsItemGroups503Error error when the service returns a 503 status code
func (m *V1beta1GroupsGroupsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1GroupsGroupsItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": CreateV1beta1GroupsItemGroups403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1GroupsItemGroups404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1GroupsItemGroups409ErrorFromDiscriminatorValue,
        "503": CreateV1beta1GroupsItemGroups503ErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get returns group for a customer
// Deprecated: This method is obsolete. Use GetAsGroupsGetResponse instead.
// returns a V1beta1GroupsItemGroupsResponseable when successful
// returns a V1beta1GroupsItemGroups403Error error when the service returns a 403 status code
// returns a V1beta1GroupsItemGroups404Error error when the service returns a 404 status code
// returns a V1beta1GroupsItemGroups503Error error when the service returns a 503 status code
func (m *V1beta1GroupsGroupsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1GroupsGroupsItemRequestBuilderGetRequestConfiguration)(V1beta1GroupsItemGroupsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": CreateV1beta1GroupsItemGroups403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1GroupsItemGroups404ErrorFromDiscriminatorValue,
        "503": CreateV1beta1GroupsItemGroups503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsItemGroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsItemGroupsResponseable), nil
}
// GetAsGroupsGetResponse returns group for a customer
// returns a V1beta1GroupsItemGroupsGetResponseable when successful
// returns a V1beta1GroupsItemGroups403Error error when the service returns a 403 status code
// returns a V1beta1GroupsItemGroups404Error error when the service returns a 404 status code
// returns a V1beta1GroupsItemGroups503Error error when the service returns a 503 status code
func (m *V1beta1GroupsGroupsItemRequestBuilder) GetAsGroupsGetResponse(ctx context.Context, requestConfiguration *V1beta1GroupsGroupsItemRequestBuilderGetRequestConfiguration)(V1beta1GroupsItemGroupsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": CreateV1beta1GroupsItemGroups403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1GroupsItemGroups404ErrorFromDiscriminatorValue,
        "503": CreateV1beta1GroupsItemGroups503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsItemGroupsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsItemGroupsGetResponseable), nil
}
// Patch update a Group.
// Deprecated: This method is obsolete. Use PatchAsGroupsPatchResponse instead.
// returns a V1beta1GroupsItemGroupsResponseable when successful
// returns a V1beta1GroupsItemGroups400Error error when the service returns a 400 status code
// returns a V1beta1GroupsItemGroups403Error error when the service returns a 403 status code
// returns a V1beta1GroupsItemGroups404Error error when the service returns a 404 status code
// returns a V1beta1GroupsItemGroups409Error error when the service returns a 409 status code
// returns a V1beta1GroupsItemGroups503Error error when the service returns a 503 status code
func (m *V1beta1GroupsGroupsItemRequestBuilder) Patch(ctx context.Context, body V1beta1GroupsItemGroupsPatchRequestBodyable, requestConfiguration *V1beta1GroupsGroupsItemRequestBuilderPatchRequestConfiguration)(V1beta1GroupsItemGroupsResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1GroupsItemGroups400ErrorFromDiscriminatorValue,
        "403": CreateV1beta1GroupsItemGroups403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1GroupsItemGroups404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1GroupsItemGroups409ErrorFromDiscriminatorValue,
        "503": CreateV1beta1GroupsItemGroups503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsItemGroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsItemGroupsResponseable), nil
}
// PatchAsGroupsPatchResponse update a Group.
// returns a V1beta1GroupsItemGroupsPatchResponseable when successful
// returns a V1beta1GroupsItemGroups400Error error when the service returns a 400 status code
// returns a V1beta1GroupsItemGroups403Error error when the service returns a 403 status code
// returns a V1beta1GroupsItemGroups404Error error when the service returns a 404 status code
// returns a V1beta1GroupsItemGroups409Error error when the service returns a 409 status code
// returns a V1beta1GroupsItemGroups503Error error when the service returns a 503 status code
func (m *V1beta1GroupsGroupsItemRequestBuilder) PatchAsGroupsPatchResponse(ctx context.Context, body V1beta1GroupsItemGroupsPatchRequestBodyable, requestConfiguration *V1beta1GroupsGroupsItemRequestBuilderPatchRequestConfiguration)(V1beta1GroupsItemGroupsPatchResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1GroupsItemGroups400ErrorFromDiscriminatorValue,
        "403": CreateV1beta1GroupsItemGroups403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1GroupsItemGroups404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1GroupsItemGroups409ErrorFromDiscriminatorValue,
        "503": CreateV1beta1GroupsItemGroups503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsItemGroupsPatchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsItemGroupsPatchResponseable), nil
}
// ToDeleteRequestInformation delete a Specific Group.
// returns a *RequestInformation when successful
func (m *V1beta1GroupsGroupsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1GroupsGroupsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns group for a customer
// returns a *RequestInformation when successful
func (m *V1beta1GroupsGroupsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1GroupsGroupsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update a Group.
// returns a *RequestInformation when successful
func (m *V1beta1GroupsGroupsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body V1beta1GroupsItemGroupsPatchRequestBodyable, requestConfiguration *V1beta1GroupsGroupsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/merge-patch+json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1GroupsGroupsItemRequestBuilder when successful
func (m *V1beta1GroupsGroupsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1GroupsGroupsItemRequestBuilder) {
    return NewV1beta1GroupsGroupsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
