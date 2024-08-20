package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Domain name of the dns server.
    domainName *string
    // Host Name of the dns server.
    hostName *string
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValuesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValuesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDomainName gets the domainName property value. Domain name of the dns server.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues) GetDomainName()(*string) {
    return m.domainName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["hostName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
        }
        return nil
    }
    return res
}
// GetHostName gets the hostName property value. Host Name of the dns server.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues) GetHostName()(*string) {
    return m.hostName
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostName", m.GetHostName())
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDomainName sets the domainName property value. Domain name of the dns server.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues) SetDomainName(value *string)() {
    m.domainName = value
}
// SetHostName sets the hostName property value. Host Name of the dns server.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValues) SetHostName(value *string)() {
    m.hostName = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_vmDnsSetting_vmDnsValuesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDomainName()(*string)
    GetHostName()(*string)
    SetDomainName(value *string)()
    SetHostName(value *string)()
}
