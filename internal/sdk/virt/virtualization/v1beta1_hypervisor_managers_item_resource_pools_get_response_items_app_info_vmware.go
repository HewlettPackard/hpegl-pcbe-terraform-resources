package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // References to the datacenter that house this virtual machine.
    datacenterInfo V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware_datacenterInfoable
    // VMware provided moref for this resource pool.
    moref *string
}
// NewV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware instantiates a new V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware and sets the default values.
func NewV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware()(*V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware) {
    m := &V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmwareFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmwareFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatacenterInfo gets the datacenterInfo property value. References to the datacenter that house this virtual machine.
// returns a V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware_datacenterInfoable when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware) GetDatacenterInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware_datacenterInfoable) {
    return m.datacenterInfo
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["datacenterInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware_datacenterInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatacenterInfo(val.(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware_datacenterInfoable))
        }
        return nil
    }
    res["moref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMoref(val)
        }
        return nil
    }
    return res
}
// GetMoref gets the moref property value. VMware provided moref for this resource pool.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware) GetMoref()(*string) {
    return m.moref
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("datacenterInfo", m.GetDatacenterInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("moref", m.GetMoref())
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
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatacenterInfo sets the datacenterInfo property value. References to the datacenter that house this virtual machine.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware) SetDatacenterInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware_datacenterInfoable)() {
    m.datacenterInfo = value
}
// SetMoref sets the moref property value. VMware provided moref for this resource pool.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware) SetMoref(value *string)() {
    m.moref = value
}
type V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmwareable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatacenterInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware_datacenterInfoable)
    GetMoref()(*string)
    SetDatacenterInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfo_vmware_datacenterInfoable)()
    SetMoref(value *string)()
}
