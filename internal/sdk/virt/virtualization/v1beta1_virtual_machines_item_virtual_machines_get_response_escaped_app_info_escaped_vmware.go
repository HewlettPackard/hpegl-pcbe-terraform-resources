package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // References to the datacenter that house this virtual machine.
    datacenterInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datacenterInfoable
    // References to all datastores that house virtual disks of this virtual machine.
    datastoresInfo []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoable
    // VMware provided moref for this virtual machine.
    moref *string
    // Information about the VMware's resource pool to which the VM belongs to.
    resourcePoolInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_resourcePoolInfoable
    // Information about the VMware tools installed in this virtual machine.
    toolsInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoable
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmwareFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmwareFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatacenterInfo gets the datacenterInfo property value. References to the datacenter that house this virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datacenterInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) GetDatacenterInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datacenterInfoable) {
    return m.datacenterInfo
}
// GetDatastoresInfo gets the datastoresInfo property value. References to all datastores that house virtual disks of this virtual machine.
// returns a []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) GetDatastoresInfo()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoable) {
    return m.datastoresInfo
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["datacenterInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datacenterInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatacenterInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datacenterInfoable))
        }
        return nil
    }
    res["datastoresInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoable)
                }
            }
            m.SetDatastoresInfo(res)
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
    res["resourcePoolInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_resourcePoolInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourcePoolInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_resourcePoolInfoable))
        }
        return nil
    }
    res["toolsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToolsInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoable))
        }
        return nil
    }
    return res
}
// GetMoref gets the moref property value. VMware provided moref for this virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) GetMoref()(*string) {
    return m.moref
}
// GetResourcePoolInfo gets the resourcePoolInfo property value. Information about the VMware's resource pool to which the VM belongs to.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_resourcePoolInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) GetResourcePoolInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_resourcePoolInfoable) {
    return m.resourcePoolInfo
}
// GetToolsInfo gets the toolsInfo property value. Information about the VMware tools installed in this virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) GetToolsInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoable) {
    return m.toolsInfo
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("datacenterInfo", m.GetDatacenterInfo())
        if err != nil {
            return err
        }
    }
    if m.GetDatastoresInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDatastoresInfo()))
        for i, v := range m.GetDatastoresInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("datastoresInfo", cast)
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
        err := writer.WriteObjectValue("resourcePoolInfo", m.GetResourcePoolInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("toolsInfo", m.GetToolsInfo())
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatacenterInfo sets the datacenterInfo property value. References to the datacenter that house this virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) SetDatacenterInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datacenterInfoable)() {
    m.datacenterInfo = value
}
// SetDatastoresInfo sets the datastoresInfo property value. References to all datastores that house virtual disks of this virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) SetDatastoresInfo(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoable)() {
    m.datastoresInfo = value
}
// SetMoref sets the moref property value. VMware provided moref for this virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) SetMoref(value *string)() {
    m.moref = value
}
// SetResourcePoolInfo sets the resourcePoolInfo property value. Information about the VMware's resource pool to which the VM belongs to.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) SetResourcePoolInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_resourcePoolInfoable)() {
    m.resourcePoolInfo = value
}
// SetToolsInfo sets the toolsInfo property value. Information about the VMware tools installed in this virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware) SetToolsInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoable)() {
    m.toolsInfo = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmwareable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatacenterInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datacenterInfoable)
    GetDatastoresInfo()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoable)
    GetMoref()(*string)
    GetResourcePoolInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_resourcePoolInfoable)
    GetToolsInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoable)
    SetDatacenterInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datacenterInfoable)()
    SetDatastoresInfo(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoable)()
    SetMoref(value *string)()
    SetResourcePoolInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_resourcePoolInfoable)()
    SetToolsInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoable)()
}
