package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersGetResponse_items_appInfo_vmware vMware specific app info.
type V1beta1HypervisorManagersGetResponse_items_appInfo_vmware struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Captures all the data center details of the hypervisor manager. VMware: A virtual data center is a container for all the inventory objects required to complete a fully functional environment for operating virtual machines. Microsoft VMM: Datacenter components include virtualization servers, networking components, and storage resources.
    datacenters []V1beta1HypervisorManagersGetResponse_items_appInfo_vmware_datacentersable
}
// NewV1beta1HypervisorManagersGetResponse_items_appInfo_vmware instantiates a new V1beta1HypervisorManagersGetResponse_items_appInfo_vmware and sets the default values.
func NewV1beta1HypervisorManagersGetResponse_items_appInfo_vmware()(*V1beta1HypervisorManagersGetResponse_items_appInfo_vmware) {
    m := &V1beta1HypervisorManagersGetResponse_items_appInfo_vmware{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorManagersGetResponse_items_appInfo_vmwareFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersGetResponse_items_appInfo_vmwareFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersGetResponse_items_appInfo_vmware(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorManagersGetResponse_items_appInfo_vmware) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatacenters gets the datacenters property value. Captures all the data center details of the hypervisor manager. VMware: A virtual data center is a container for all the inventory objects required to complete a fully functional environment for operating virtual machines. Microsoft VMM: Datacenter components include virtualization servers, networking components, and storage resources.
// returns a []V1beta1HypervisorManagersGetResponse_items_appInfo_vmware_datacentersable when successful
func (m *V1beta1HypervisorManagersGetResponse_items_appInfo_vmware) GetDatacenters()([]V1beta1HypervisorManagersGetResponse_items_appInfo_vmware_datacentersable) {
    return m.datacenters
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorManagersGetResponse_items_appInfo_vmware) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["datacenters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorManagersGetResponse_items_appInfo_vmware_datacentersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorManagersGetResponse_items_appInfo_vmware_datacentersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorManagersGetResponse_items_appInfo_vmware_datacentersable)
                }
            }
            m.SetDatacenters(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorManagersGetResponse_items_appInfo_vmware) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDatacenters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDatacenters()))
        for i, v := range m.GetDatacenters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("datacenters", cast)
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
func (m *V1beta1HypervisorManagersGetResponse_items_appInfo_vmware) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatacenters sets the datacenters property value. Captures all the data center details of the hypervisor manager. VMware: A virtual data center is a container for all the inventory objects required to complete a fully functional environment for operating virtual machines. Microsoft VMM: Datacenter components include virtualization servers, networking components, and storage resources.
func (m *V1beta1HypervisorManagersGetResponse_items_appInfo_vmware) SetDatacenters(value []V1beta1HypervisorManagersGetResponse_items_appInfo_vmware_datacentersable)() {
    m.datacenters = value
}
type V1beta1HypervisorManagersGetResponse_items_appInfo_vmwareable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatacenters()([]V1beta1HypervisorManagersGetResponse_items_appInfo_vmware_datacentersable)
    SetDatacenters(value []V1beta1HypervisorManagersGetResponse_items_appInfo_vmware_datacentersable)()
}
