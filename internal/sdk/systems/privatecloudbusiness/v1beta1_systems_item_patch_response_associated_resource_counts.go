package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemPatchResponse_associatedResourceCounts associated Resource Information of system.
type V1beta1SystemsItemPatchResponse_associatedResourceCounts struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Total Datastore Count of the system
    datastoreCount *float64
    // Total Hypervisor Cluster Count of the system
    hypervisorClusterCount *float64
    // Total Servers Count of the system
    serversCount *float64
    // Total Storage Array Count of the system
    storageArrayCount *float64
    // Total Switches Count of the system
    switchesCount *float64
    // Total Virtual Machine Count of the system
    vmCount *float64
}
// NewV1beta1SystemsItemPatchResponse_associatedResourceCounts instantiates a new V1beta1SystemsItemPatchResponse_associatedResourceCounts and sets the default values.
func NewV1beta1SystemsItemPatchResponse_associatedResourceCounts()(*V1beta1SystemsItemPatchResponse_associatedResourceCounts) {
    m := &V1beta1SystemsItemPatchResponse_associatedResourceCounts{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemPatchResponse_associatedResourceCountsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemPatchResponse_associatedResourceCountsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemPatchResponse_associatedResourceCounts(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatastoreCount gets the datastoreCount property value. Total Datastore Count of the system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) GetDatastoreCount()(*float64) {
    return m.datastoreCount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["datastoreCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatastoreCount(val)
        }
        return nil
    }
    res["hypervisorClusterCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorClusterCount(val)
        }
        return nil
    }
    res["serversCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServersCount(val)
        }
        return nil
    }
    res["storageArrayCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageArrayCount(val)
        }
        return nil
    }
    res["switchesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSwitchesCount(val)
        }
        return nil
    }
    res["vmCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmCount(val)
        }
        return nil
    }
    return res
}
// GetHypervisorClusterCount gets the hypervisorClusterCount property value. Total Hypervisor Cluster Count of the system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) GetHypervisorClusterCount()(*float64) {
    return m.hypervisorClusterCount
}
// GetServersCount gets the serversCount property value. Total Servers Count of the system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) GetServersCount()(*float64) {
    return m.serversCount
}
// GetStorageArrayCount gets the storageArrayCount property value. Total Storage Array Count of the system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) GetStorageArrayCount()(*float64) {
    return m.storageArrayCount
}
// GetSwitchesCount gets the switchesCount property value. Total Switches Count of the system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) GetSwitchesCount()(*float64) {
    return m.switchesCount
}
// GetVmCount gets the vmCount property value. Total Virtual Machine Count of the system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) GetVmCount()(*float64) {
    return m.vmCount
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("datastoreCount", m.GetDatastoreCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("hypervisorClusterCount", m.GetHypervisorClusterCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("serversCount", m.GetServersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("storageArrayCount", m.GetStorageArrayCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("switchesCount", m.GetSwitchesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("vmCount", m.GetVmCount())
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
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatastoreCount sets the datastoreCount property value. Total Datastore Count of the system
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) SetDatastoreCount(value *float64)() {
    m.datastoreCount = value
}
// SetHypervisorClusterCount sets the hypervisorClusterCount property value. Total Hypervisor Cluster Count of the system
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) SetHypervisorClusterCount(value *float64)() {
    m.hypervisorClusterCount = value
}
// SetServersCount sets the serversCount property value. Total Servers Count of the system
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) SetServersCount(value *float64)() {
    m.serversCount = value
}
// SetStorageArrayCount sets the storageArrayCount property value. Total Storage Array Count of the system
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) SetStorageArrayCount(value *float64)() {
    m.storageArrayCount = value
}
// SetSwitchesCount sets the switchesCount property value. Total Switches Count of the system
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) SetSwitchesCount(value *float64)() {
    m.switchesCount = value
}
// SetVmCount sets the vmCount property value. Total Virtual Machine Count of the system
func (m *V1beta1SystemsItemPatchResponse_associatedResourceCounts) SetVmCount(value *float64)() {
    m.vmCount = value
}
type V1beta1SystemsItemPatchResponse_associatedResourceCountsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatastoreCount()(*float64)
    GetHypervisorClusterCount()(*float64)
    GetServersCount()(*float64)
    GetStorageArrayCount()(*float64)
    GetSwitchesCount()(*float64)
    GetVmCount()(*float64)
    SetDatastoreCount(value *float64)()
    SetHypervisorClusterCount(value *float64)()
    SetServersCount(value *float64)()
    SetStorageArrayCount(value *float64)()
    SetSwitchesCount(value *float64)()
    SetVmCount(value *float64)()
}
