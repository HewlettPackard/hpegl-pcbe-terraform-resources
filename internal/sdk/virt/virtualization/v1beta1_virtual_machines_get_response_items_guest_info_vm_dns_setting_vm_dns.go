package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // IP Addresses associated with dns server.
    ipAddresses []string
    // The searchDomains property
    searchDomains []string
}
// NewV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns instantiates a new V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns()(*V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns) {
    m := &V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ipAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetIpAddresses(res)
        }
        return nil
    }
    res["searchDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSearchDomains(res)
        }
        return nil
    }
    return res
}
// GetIpAddresses gets the ipAddresses property value. IP Addresses associated with dns server.
// returns a []string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns) GetIpAddresses()([]string) {
    return m.ipAddresses
}
// GetSearchDomains gets the searchDomains property value. The searchDomains property
// returns a []string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns) GetSearchDomains()([]string) {
    return m.searchDomains
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIpAddresses() != nil {
        err := writer.WriteCollectionOfStringValues("ipAddresses", m.GetIpAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetSearchDomains() != nil {
        err := writer.WriteCollectionOfStringValues("searchDomains", m.GetSearchDomains())
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
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIpAddresses sets the ipAddresses property value. IP Addresses associated with dns server.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns) SetIpAddresses(value []string)() {
    m.ipAddresses = value
}
// SetSearchDomains sets the searchDomains property value. The searchDomains property
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDns) SetSearchDomains(value []string)() {
    m.searchDomains = value
}
type V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIpAddresses()([]string)
    GetSearchDomains()([]string)
    SetIpAddresses(value []string)()
    SetSearchDomains(value []string)()
}
