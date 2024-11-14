package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // gateway. Applies only to Alletra dHCI.
    gateway *int32
    // IP Address. Applies only to Alletra dHCI.
    ipAddress *string
    // Subnet mask. Applies only to Alletra dHCI.
    subnetMask *string
}
// NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo instantiates a new V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo and sets the default values.
func NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo()(*V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) {
    m := &V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["gateway"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGateway(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["subnetMask"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetMask(val)
        }
        return nil
    }
    return res
}
// GetGateway gets the gateway property value. gateway. Applies only to Alletra dHCI.
// returns a *int32 when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) GetGateway()(*int32) {
    return m.gateway
}
// GetIpAddress gets the ipAddress property value. IP Address. Applies only to Alletra dHCI.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetSubnetMask gets the subnetMask property value. Subnet mask. Applies only to Alletra dHCI.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) GetSubnetMask()(*string) {
    return m.subnetMask
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("gateway", m.GetGateway())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnetMask", m.GetSubnetMask())
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
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGateway sets the gateway property value. gateway. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) SetGateway(value *int32)() {
    m.gateway = value
}
// SetIpAddress sets the ipAddress property value. IP Address. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetSubnetMask sets the subnetMask property value. Subnet mask. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfo) SetSubnetMask(value *string)() {
    m.subnetMask = value
}
type V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGateway()(*int32)
    GetIpAddress()(*string)
    GetSubnetMask()(*string)
    SetGateway(value *int32)()
    SetIpAddress(value *string)()
    SetSubnetMask(value *string)()
}
