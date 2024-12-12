package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // IP Addresses information of the Network Interface.
    ipAddresses []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesable
    // Mac Address of the Network Interface.
    macAddress *string
    // Nic of the Network Interface.
    nic *string
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ipAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesable)
                }
            }
            m.SetIpAddresses(res)
        }
        return nil
    }
    res["macAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddress(val)
        }
        return nil
    }
    res["nic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNic(val)
        }
        return nil
    }
    return res
}
// GetIpAddresses gets the ipAddresses property value. IP Addresses information of the Network Interface.
// returns a []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) GetIpAddresses()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesable) {
    return m.ipAddresses
}
// GetMacAddress gets the macAddress property value. Mac Address of the Network Interface.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) GetMacAddress()(*string) {
    return m.macAddress
}
// GetNic gets the nic property value. Nic of the Network Interface.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) GetNic()(*string) {
    return m.nic
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIpAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIpAddresses()))
        for i, v := range m.GetIpAddresses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("ipAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("nic", m.GetNic())
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIpAddresses sets the ipAddresses property value. IP Addresses information of the Network Interface.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) SetIpAddresses(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesable)() {
    m.ipAddresses = value
}
// SetMacAddress sets the macAddress property value. Mac Address of the Network Interface.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) SetMacAddress(value *string)() {
    m.macAddress = value
}
// SetNic sets the nic property value. Nic of the Network Interface.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo) SetNic(value *string)() {
    m.nic = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIpAddresses()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesable)
    GetMacAddress()(*string)
    GetNic()(*string)
    SetIpAddresses(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesable)()
    SetMacAddress(value *string)()
    SetNic(value *string)()
}
