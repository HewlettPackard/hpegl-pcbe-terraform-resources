package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // IP Address. Applies only to Alletra dHCI.
    ipAddress *string
}
// NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos instantiates a new V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos and sets the default values.
func NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos()(*V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos) {
    m := &V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetIpAddress gets the ipAddress property value. IP Address. Applies only to Alletra dHCI.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos) GetIpAddress()(*string) {
    return m.ipAddress
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
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
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIpAddress sets the ipAddress property value. IP Address. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos) SetIpAddress(value *string)() {
    m.ipAddress = value
}
type V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIpAddress()(*string)
    SetIpAddress(value *string)()
}
