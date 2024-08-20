package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse a subscription whose values are defined by, and synchronized with, a cloud service provider.
type V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse struct {
    // Information about the cloud provider account where the resource is located.
    accountInfo V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse_accountInfoable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Resource identifier assigned by the cloud provider.
    cspId *string
    // Resource name assigned by the cloud provider, if any.
    cspName *string
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // An identifier for the resource, usually a UUID.
    id *string
    // A system specified name for the resource.
    name *string
    // The self reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse instantiates a new V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse and sets the default values.
func NewV1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse()(*V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) {
    m := &V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse(), nil
}
// GetAccountInfo gets the accountInfo property value. Information about the cloud provider account where the resource is located.
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse_accountInfoable when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetAccountInfo()(V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse_accountInfoable) {
    return m.accountInfo
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCspId gets the cspId property value. Resource identifier assigned by the cloud provider.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetCspId()(*string) {
    return m.cspId
}
// GetCspName gets the cspName property value. Resource name assigned by the cloud provider, if any.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetCspName()(*string) {
    return m.cspName
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse_accountInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountInfo(val.(V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse_accountInfoable))
        }
        return nil
    }
    res["cspId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspId(val)
        }
        return nil
    }
    res["cspName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspName(val)
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["generation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["resourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceUri(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountInfo sets the accountInfo property value. Information about the cloud provider account where the resource is located.
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetAccountInfo(value V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse_accountInfoable)() {
    m.accountInfo = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCspId sets the cspId property value. Resource identifier assigned by the cloud provider.
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetCspId(value *string)() {
    m.cspId = value
}
// SetCspName sets the cspName property value. Resource name assigned by the cloud provider, if any.
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetCspName(value *string)() {
    m.cspName = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountInfo()(V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse_accountInfoable)
    GetCspId()(*string)
    GetCspName()(*string)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetAccountInfo(value V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponse_accountInfoable)()
    SetCspId(value *string)()
    SetCspName(value *string)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
