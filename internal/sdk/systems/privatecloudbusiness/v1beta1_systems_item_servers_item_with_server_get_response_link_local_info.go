package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo information about link local IP addresses.
type V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The linkLocalIpV4Address property
    linkLocalIpV4Address *string
    // The linkLocalIpV6Address property
    linkLocalIpV6Address *string
}
// NewV1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo instantiates a new V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo and sets the default values.
func NewV1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo()(*V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo) {
    m := &V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["linkLocalIpV4Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkLocalIpV4Address(val)
        }
        return nil
    }
    res["linkLocalIpV6Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkLocalIpV6Address(val)
        }
        return nil
    }
    return res
}
// GetLinkLocalIpV4Address gets the linkLocalIpV4Address property value. The linkLocalIpV4Address property
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo) GetLinkLocalIpV4Address()(*string) {
    return m.linkLocalIpV4Address
}
// GetLinkLocalIpV6Address gets the linkLocalIpV6Address property value. The linkLocalIpV6Address property
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo) GetLinkLocalIpV6Address()(*string) {
    return m.linkLocalIpV6Address
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("linkLocalIpV4Address", m.GetLinkLocalIpV4Address())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("linkLocalIpV6Address", m.GetLinkLocalIpV6Address())
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
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLinkLocalIpV4Address sets the linkLocalIpV4Address property value. The linkLocalIpV4Address property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo) SetLinkLocalIpV4Address(value *string)() {
    m.linkLocalIpV4Address = value
}
// SetLinkLocalIpV6Address sets the linkLocalIpV6Address property value. The linkLocalIpV6Address property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfo) SetLinkLocalIpV6Address(value *string)() {
    m.linkLocalIpV6Address = value
}
type V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLinkLocalIpV4Address()(*string)
    GetLinkLocalIpV6Address()(*string)
    SetLinkLocalIpV4Address(value *string)()
    SetLinkLocalIpV6Address(value *string)()
}
