package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Enable/Disable the InfoSight Remote Support for support engineers to login to Data Services Connector via InfoSight Interactive Device Acceess feature
    remoteAccessEnabled *bool
}
// NewV1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody instantiates a new V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody and sets the default values.
func NewV1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody()(*V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody) {
    m := &V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["remoteAccessEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAccessEnabled(val)
        }
        return nil
    }
    return res
}
// GetRemoteAccessEnabled gets the remoteAccessEnabled property value. Enable/Disable the InfoSight Remote Support for support engineers to login to Data Services Connector via InfoSight Interactive Device Acceess feature
// returns a *bool when successful
func (m *V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody) GetRemoteAccessEnabled()(*bool) {
    return m.remoteAccessEnabled
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("remoteAccessEnabled", m.GetRemoteAccessEnabled())
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
func (m *V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRemoteAccessEnabled sets the remoteAccessEnabled property value. Enable/Disable the InfoSight Remote Support for support engineers to login to Data Services Connector via InfoSight Interactive Device Acceess feature
func (m *V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBody) SetRemoteAccessEnabled(value *bool)() {
    m.remoteAccessEnabled = value
}
type V1beta1DataServicesConnectorsItemSetRemoteSupportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRemoteAccessEnabled()(*bool)
    SetRemoteAccessEnabled(value *bool)()
}
