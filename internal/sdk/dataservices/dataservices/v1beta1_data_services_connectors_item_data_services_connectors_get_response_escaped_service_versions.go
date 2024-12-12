package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the DSC service.
    serviceName *string
    // Version of the DSC service.
    serviceVersion *string
}
// NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions instantiates a new V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions and sets the default values.
func NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions()(*V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions) {
    m := &V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["serviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceName(val)
        }
        return nil
    }
    res["serviceVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceVersion(val)
        }
        return nil
    }
    return res
}
// GetServiceName gets the serviceName property value. Name of the DSC service.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions) GetServiceName()(*string) {
    return m.serviceName
}
// GetServiceVersion gets the serviceVersion property value. Version of the DSC service.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions) GetServiceVersion()(*string) {
    return m.serviceVersion
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("serviceName", m.GetServiceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serviceVersion", m.GetServiceVersion())
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
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetServiceName sets the serviceName property value. Name of the DSC service.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions) SetServiceName(value *string)() {
    m.serviceName = value
}
// SetServiceVersion sets the serviceVersion property value. Version of the DSC service.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersions) SetServiceVersion(value *string)() {
    m.serviceVersion = value
}
type V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetServiceName()(*string)
    GetServiceVersion()(*string)
    SetServiceName(value *string)()
    SetServiceVersion(value *string)()
}
