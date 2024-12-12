package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DataServicesConnectorsItemGenerateTotpPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // TOTP expiry time in seconds.
    expiryInSeconds *int32
    // Time-based one-time password (TOTP) used to login to DSC VM UI and execute on-prem APIs
    totp *string
}
// NewV1beta1DataServicesConnectorsItemGenerateTotpPostResponse instantiates a new V1beta1DataServicesConnectorsItemGenerateTotpPostResponse and sets the default values.
func NewV1beta1DataServicesConnectorsItemGenerateTotpPostResponse()(*V1beta1DataServicesConnectorsItemGenerateTotpPostResponse) {
    m := &V1beta1DataServicesConnectorsItemGenerateTotpPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsItemGenerateTotpPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemGenerateTotpPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemGenerateTotpPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsItemGenerateTotpPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiryInSeconds gets the expiryInSeconds property value. TOTP expiry time in seconds.
// returns a *int32 when successful
func (m *V1beta1DataServicesConnectorsItemGenerateTotpPostResponse) GetExpiryInSeconds()(*int32) {
    return m.expiryInSeconds
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsItemGenerateTotpPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expiryInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryInSeconds(val)
        }
        return nil
    }
    res["totp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotp(val)
        }
        return nil
    }
    return res
}
// GetTotp gets the totp property value. Time-based one-time password (TOTP) used to login to DSC VM UI and execute on-prem APIs
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemGenerateTotpPostResponse) GetTotp()(*string) {
    return m.totp
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsItemGenerateTotpPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("expiryInSeconds", m.GetExpiryInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("totp", m.GetTotp())
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
func (m *V1beta1DataServicesConnectorsItemGenerateTotpPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiryInSeconds sets the expiryInSeconds property value. TOTP expiry time in seconds.
func (m *V1beta1DataServicesConnectorsItemGenerateTotpPostResponse) SetExpiryInSeconds(value *int32)() {
    m.expiryInSeconds = value
}
// SetTotp sets the totp property value. Time-based one-time password (TOTP) used to login to DSC VM UI and execute on-prem APIs
func (m *V1beta1DataServicesConnectorsItemGenerateTotpPostResponse) SetTotp(value *string)() {
    m.totp = value
}
type V1beta1DataServicesConnectorsItemGenerateTotpPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiryInSeconds()(*int32)
    GetTotp()(*string)
    SetExpiryInSeconds(value *int32)()
    SetTotp(value *string)()
}
