package dataservices

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent the Software Component a Software Release belongs to.
type V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the resource.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The name property
    name *string
    // The type property
    typeEscaped *string
}
// NewV1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent instantiates a new V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent and sets the default values.
func NewV1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent()(*V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) {
    m := &V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteUUIDValue("id", m.GetId())
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
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The ID of the resource.
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponent) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse_softwareComponentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetName()(*string)
    GetTypeEscaped()(*string)
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetName(value *string)()
    SetTypeEscaped(value *string)()
}
