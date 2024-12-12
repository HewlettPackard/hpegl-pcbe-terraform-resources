package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An IP address or FQDN of the NTP server.
    networkAddress *string
}
// NewV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp instantiates a new V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp and sets the default values.
func NewV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp()(*V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp) {
    m := &V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["networkAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkAddress(val)
        }
        return nil
    }
    return res
}
// GetNetworkAddress gets the networkAddress property value. An IP address or FQDN of the NTP server.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp) GetNetworkAddress()(*string) {
    return m.networkAddress
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("networkAddress", m.GetNetworkAddress())
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
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNetworkAddress sets the networkAddress property value. An IP address or FQDN of the NTP server.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntp) SetNetworkAddress(value *string)() {
    m.networkAddress = value
}
type V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetworkAddress()(*string)
    SetNetworkAddress(value *string)()
}
