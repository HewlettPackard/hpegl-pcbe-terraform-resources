package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesMigrateRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines\migrate
type V1beta1VirtualMachinesMigrateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesMigrateRequestBuilderPostQueryParameters migrate virtual machines to another cluster or datastore
type V1beta1VirtualMachinesMigrateRequestBuilderPostQueryParameters struct {
    // Type of the query parameter (vms or sourceDataStoreIDs)
    ParameterType *string `uriparametername:"parameter%2Dtype"`
    // A list of VM IDs to migrate or a list of SourceDatastoreIDs. If you specify the list of SourceDatastoreIDs, all VMs within the specified datastores are migrated. If you specify a list of VM IDs, only those VMs are migrated.
    Vms []string `uriparametername:"vms"`
}
// V1beta1VirtualMachinesMigrateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesMigrateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1VirtualMachinesMigrateRequestBuilderPostQueryParameters
}
// NewV1beta1VirtualMachinesMigrateRequestBuilderInternal instantiates a new V1beta1VirtualMachinesMigrateRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesMigrateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesMigrateRequestBuilder) {
    m := &V1beta1VirtualMachinesMigrateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines/migrate{?parameter%2Dtype*,vms*}", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesMigrateRequestBuilder instantiates a new V1beta1VirtualMachinesMigrateRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesMigrateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesMigrateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesMigrateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post migrate virtual machines to another cluster or datastore
// Deprecated: This method is obsolete. Use PostAsMigratePostResponse instead.
// returns a V1beta1VirtualMachinesMigrateResponseable when successful
// returns a V1beta1VirtualMachinesMigrate400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesMigrate401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesMigrate403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesMigrate500Error error when the service returns a 500 status code
func (m *V1beta1VirtualMachinesMigrateRequestBuilder) Post(ctx context.Context, body V1beta1VirtualMachinesMigratePostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesMigrateRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesMigrateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesMigrate400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesMigrate401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesMigrate403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesMigrate500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesMigrateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesMigrateResponseable), nil
}
// PostAsMigratePostResponse migrate virtual machines to another cluster or datastore
// returns a V1beta1VirtualMachinesMigratePostResponseable when successful
// returns a V1beta1VirtualMachinesMigrate400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesMigrate401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesMigrate403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesMigrate500Error error when the service returns a 500 status code
func (m *V1beta1VirtualMachinesMigrateRequestBuilder) PostAsMigratePostResponse(ctx context.Context, body V1beta1VirtualMachinesMigratePostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesMigrateRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesMigratePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesMigrate400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesMigrate401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesMigrate403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesMigrate500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesMigratePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesMigratePostResponseable), nil
}
// ToPostRequestInformation migrate virtual machines to another cluster or datastore
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesMigrateRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1VirtualMachinesMigratePostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesMigrateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
// returns a *V1beta1VirtualMachinesMigrateRequestBuilder when successful
func (m *V1beta1VirtualMachinesMigrateRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesMigrateRequestBuilder) {
    return NewV1beta1VirtualMachinesMigrateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
