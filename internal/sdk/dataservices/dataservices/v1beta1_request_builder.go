package dataservices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1RequestBuilder builds and executes requests for operations under \data-services\v1beta1
type V1beta1RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Announcements the announcements property
// returns a *V1beta1AnnouncementsRequestBuilder when successful
func (m *V1beta1RequestBuilder) Announcements()(*V1beta1AnnouncementsRequestBuilder) {
    return NewV1beta1AnnouncementsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AsyncOperations the asyncOperations property
// returns a *V1beta1AsyncOperationsRequestBuilder when successful
func (m *V1beta1RequestBuilder) AsyncOperations()(*V1beta1AsyncOperationsRequestBuilder) {
    return NewV1beta1AsyncOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1RequestBuilderInternal instantiates a new V1beta1RequestBuilder and sets the default values.
func NewV1beta1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1RequestBuilder) {
    m := &V1beta1RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1", pathParameters),
    }
    return m
}
// NewV1beta1RequestBuilder instantiates a new V1beta1RequestBuilder and sets the default values.
func NewV1beta1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1RequestBuilderInternal(urlParams, requestAdapter)
}
// DataServicesConnectors the dataServicesConnectors property
// returns a *V1beta1DataServicesConnectorsRequestBuilder when successful
func (m *V1beta1RequestBuilder) DataServicesConnectors()(*V1beta1DataServicesConnectorsRequestBuilder) {
    return NewV1beta1DataServicesConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DualAuthOperations the dualAuthOperations property
// returns a *V1beta1DualAuthOperationsRequestBuilder when successful
func (m *V1beta1RequestBuilder) DualAuthOperations()(*V1beta1DualAuthOperationsRequestBuilder) {
    return NewV1beta1DualAuthOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EmailSubscriptions the emailSubscriptions property
// returns a *V1beta1EmailSubscriptionsRequestBuilder when successful
func (m *V1beta1RequestBuilder) EmailSubscriptions()(*V1beta1EmailSubscriptionsRequestBuilder) {
    return NewV1beta1EmailSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Groups the groups property
// returns a *V1beta1GroupsRequestBuilder when successful
func (m *V1beta1RequestBuilder) Groups()(*V1beta1GroupsRequestBuilder) {
    return NewV1beta1GroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
// returns a *V1beta1IssuesRequestBuilder when successful
func (m *V1beta1RequestBuilder) Issues()(*V1beta1IssuesRequestBuilder) {
    return NewV1beta1IssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IssuesMetadata the issuesMetadata property
// returns a *V1beta1IssuesMetadataRequestBuilder when successful
func (m *V1beta1RequestBuilder) IssuesMetadata()(*V1beta1IssuesMetadataRequestBuilder) {
    return NewV1beta1IssuesMetadataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OneTimeTokens the oneTimeTokens property
// returns a *V1beta1OneTimeTokensRequestBuilder when successful
func (m *V1beta1RequestBuilder) OneTimeTokens()(*V1beta1OneTimeTokensRequestBuilder) {
    return NewV1beta1OneTimeTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecentSearches the recentSearches property
// returns a *V1beta1RecentSearchesRequestBuilder when successful
func (m *V1beta1RequestBuilder) RecentSearches()(*V1beta1RecentSearchesRequestBuilder) {
    return NewV1beta1RecentSearchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
// returns a *V1beta1SearchRequestBuilder when successful
func (m *V1beta1RequestBuilder) Search()(*V1beta1SearchRequestBuilder) {
    return NewV1beta1SearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchMetadata the searchMetadata property
// returns a *V1beta1SearchMetadataRequestBuilder when successful
func (m *V1beta1RequestBuilder) SearchMetadata()(*V1beta1SearchMetadataRequestBuilder) {
    return NewV1beta1SearchMetadataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchSummaries the searchSummaries property
// returns a *V1beta1SearchSummariesRequestBuilder when successful
func (m *V1beta1RequestBuilder) SearchSummaries()(*V1beta1SearchSummariesRequestBuilder) {
    return NewV1beta1SearchSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecretAssignments the secretAssignments property
// returns a *V1beta1SecretAssignmentsRequestBuilder when successful
func (m *V1beta1RequestBuilder) SecretAssignments()(*V1beta1SecretAssignmentsRequestBuilder) {
    return NewV1beta1SecretAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Secrets the secrets property
// returns a *V1beta1SecretsRequestBuilder when successful
func (m *V1beta1RequestBuilder) Secrets()(*V1beta1SecretsRequestBuilder) {
    return NewV1beta1SecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings the settings property
// returns a *V1beta1SettingsRequestBuilder when successful
func (m *V1beta1RequestBuilder) Settings()(*V1beta1SettingsRequestBuilder) {
    return NewV1beta1SettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareComponents the softwareComponents property
// returns a *V1beta1SoftwareComponentsRequestBuilder when successful
func (m *V1beta1RequestBuilder) SoftwareComponents()(*V1beta1SoftwareComponentsRequestBuilder) {
    return NewV1beta1SoftwareComponentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareReleases the softwareReleases property
// returns a *V1beta1SoftwareReleasesRequestBuilder when successful
func (m *V1beta1RequestBuilder) SoftwareReleases()(*V1beta1SoftwareReleasesRequestBuilder) {
    return NewV1beta1SoftwareReleasesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareUpgrades the softwareUpgrades property
// returns a *V1beta1SoftwareUpgradesRequestBuilder when successful
func (m *V1beta1RequestBuilder) SoftwareUpgrades()(*V1beta1SoftwareUpgradesRequestBuilder) {
    return NewV1beta1SoftwareUpgradesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StorageLocations the storageLocations property
// returns a *V1beta1StorageLocationsRequestBuilder when successful
func (m *V1beta1RequestBuilder) StorageLocations()(*V1beta1StorageLocationsRequestBuilder) {
    return NewV1beta1StorageLocationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags the tags property
// returns a *V1beta1TagsRequestBuilder when successful
func (m *V1beta1RequestBuilder) Tags()(*V1beta1TagsRequestBuilder) {
    return NewV1beta1TagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Webhooks the webhooks property
// returns a *V1beta1WebhooksRequestBuilder when successful
func (m *V1beta1RequestBuilder) Webhooks()(*V1beta1WebhooksRequestBuilder) {
    return NewV1beta1WebhooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WebhooksMetadata the webhooksMetadata property
// returns a *V1beta1WebhooksMetadataRequestBuilder when successful
func (m *V1beta1RequestBuilder) WebhooksMetadata()(*V1beta1WebhooksMetadataRequestBuilder) {
    return NewV1beta1WebhooksMetadataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
