package dataservices

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SoftwareUpgradesGetResponse_items an upgrade for a previously installed Software Release.
type V1beta1SoftwareUpgradesGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether access to the release is currently allowed. If it is not allowed, the`notes` property contains actions to complete to make it allowed.
    allowed *bool
    // The ID of the resource.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Actions to completed before the upgrade is allowed. Empty if `allowed` is `true`.
    notes *string
    // A release of software.
    softwareRelease V1beta1SoftwareUpgradesGetResponse_items_softwareReleaseable
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1SoftwareUpgradesGetResponse_items instantiates a new V1beta1SoftwareUpgradesGetResponse_items and sets the default values.
func NewV1beta1SoftwareUpgradesGetResponse_items()(*V1beta1SoftwareUpgradesGetResponse_items) {
    m := &V1beta1SoftwareUpgradesGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SoftwareUpgradesGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SoftwareUpgradesGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SoftwareUpgradesGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SoftwareUpgradesGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowed gets the allowed property value. Whether access to the release is currently allowed. If it is not allowed, the`notes` property contains actions to complete to make it allowed.
// returns a *bool when successful
func (m *V1beta1SoftwareUpgradesGetResponse_items) GetAllowed()(*bool) {
    return m.allowed
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SoftwareUpgradesGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowed(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["softwareRelease"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SoftwareUpgradesGetResponse_items_softwareReleaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareRelease(val.(V1beta1SoftwareUpgradesGetResponse_items_softwareReleaseable))
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
// GetId gets the id property value. The ID of the resource.
// returns a *UUID when successful
func (m *V1beta1SoftwareUpgradesGetResponse_items) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetNotes gets the notes property value. Actions to completed before the upgrade is allowed. Empty if `allowed` is `true`.
// returns a *string when successful
func (m *V1beta1SoftwareUpgradesGetResponse_items) GetNotes()(*string) {
    return m.notes
}
// GetSoftwareRelease gets the softwareRelease property value. A release of software.
// returns a V1beta1SoftwareUpgradesGetResponse_items_softwareReleaseable when successful
func (m *V1beta1SoftwareUpgradesGetResponse_items) GetSoftwareRelease()(V1beta1SoftwareUpgradesGetResponse_items_softwareReleaseable) {
    return m.softwareRelease
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1SoftwareUpgradesGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1SoftwareUpgradesGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowed", m.GetAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("softwareRelease", m.GetSoftwareRelease())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1SoftwareUpgradesGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowed sets the allowed property value. Whether access to the release is currently allowed. If it is not allowed, the`notes` property contains actions to complete to make it allowed.
func (m *V1beta1SoftwareUpgradesGetResponse_items) SetAllowed(value *bool)() {
    m.allowed = value
}
// SetId sets the id property value. The ID of the resource.
func (m *V1beta1SoftwareUpgradesGetResponse_items) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetNotes sets the notes property value. Actions to completed before the upgrade is allowed. Empty if `allowed` is `true`.
func (m *V1beta1SoftwareUpgradesGetResponse_items) SetNotes(value *string)() {
    m.notes = value
}
// SetSoftwareRelease sets the softwareRelease property value. A release of software.
func (m *V1beta1SoftwareUpgradesGetResponse_items) SetSoftwareRelease(value V1beta1SoftwareUpgradesGetResponse_items_softwareReleaseable)() {
    m.softwareRelease = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1SoftwareUpgradesGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1SoftwareUpgradesGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowed()(*bool)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetNotes()(*string)
    GetSoftwareRelease()(V1beta1SoftwareUpgradesGetResponse_items_softwareReleaseable)
    GetTypeEscaped()(*string)
    SetAllowed(value *bool)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetNotes(value *string)()
    SetSoftwareRelease(value V1beta1SoftwareUpgradesGetResponse_items_softwareReleaseable)()
    SetTypeEscaped(value *string)()
}
