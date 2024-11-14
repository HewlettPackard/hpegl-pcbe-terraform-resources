package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo information about the OS on the server.
type V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The license property
    license *string
    // The version property
    version *string
}
// NewV1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo instantiates a new V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo and sets the default values.
func NewV1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo()(*V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo) {
    m := &V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["license"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicense(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLicense gets the license property value. The license property
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo) GetLicense()(*string) {
    return m.license
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("license", m.GetLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLicense sets the license property value. The license property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo) SetLicense(value *string)() {
    m.license = value
}
// SetVersion sets the version property value. The version property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfo) SetVersion(value *string)() {
    m.version = value
}
type V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLicense()(*string)
    GetVersion()(*string)
    SetLicense(value *string)()
    SetVersion(value *string)()
}
