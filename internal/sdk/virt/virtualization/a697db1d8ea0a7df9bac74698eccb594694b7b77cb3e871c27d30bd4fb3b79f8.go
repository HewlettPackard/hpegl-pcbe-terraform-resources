package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An identifier for the resource, usually a UUID.
    id *string
    // The self reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo instantiates a new V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo and sets the default values.
func NewV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo()(*V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) {
    m := &V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) GetId()(*string) {
    return m.id
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) SetId(value *string)() {
    m.id = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponse_subscriptionInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetId(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
