package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // NTP servers against which the Data Services Connector should synchronize.
    ntpServers []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp_ntpServersable
    // Reason for the current state of the NTP configuration.
    stateReason *string
}
// NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp instantiates a new V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp and sets the default values.
func NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp()(*V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp) {
    m := &V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ntpServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp_ntpServersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp_ntpServersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp_ntpServersable)
                }
            }
            m.SetNtpServers(res)
        }
        return nil
    }
    res["stateReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateReason(val)
        }
        return nil
    }
    return res
}
// GetNtpServers gets the ntpServers property value. NTP servers against which the Data Services Connector should synchronize.
// returns a []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp_ntpServersable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp) GetNtpServers()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp_ntpServersable) {
    return m.ntpServers
}
// GetStateReason gets the stateReason property value. Reason for the current state of the NTP configuration.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp) GetStateReason()(*string) {
    return m.stateReason
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNtpServers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNtpServers()))
        for i, v := range m.GetNtpServers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("ntpServers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stateReason", m.GetStateReason())
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
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNtpServers sets the ntpServers property value. NTP servers against which the Data Services Connector should synchronize.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp) SetNtpServers(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp_ntpServersable)() {
    m.ntpServers = value
}
// SetStateReason sets the stateReason property value. Reason for the current state of the NTP configuration.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp) SetStateReason(value *string)() {
    m.stateReason = value
}
type V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNtpServers()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp_ntpServersable)
    GetStateReason()(*string)
    SetNtpServers(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntp_ntpServersable)()
    SetStateReason(value *string)()
}
