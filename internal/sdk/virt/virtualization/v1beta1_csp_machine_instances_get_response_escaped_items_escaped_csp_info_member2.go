package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2 machine instance properties whose values are defined by the Azure cloud service provider.
type V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Cloud provider availability zones
    availabilityZones []string
    // Number of CPU cores configured.
    cpuCoreCount *int32
    // The time at which the machine instance was launched.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Cloud provider region
    cspRegion *string
    // List of tags assigned to the resource
    cspTags []V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_cspTagsable
    // A cloud provider designated machine instance type.
    instanceType *string
    // Name of the key pair associated with the machine instance, if any.
    keyPairName *string
    // Network properties for the machine instance.
    networkInfo V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_networkInfoable
    // A cloud provider designated machine platform.
    platform *string
    // The device name for the root volume, for example /dev/sda1.
    rootDevice *string
    // Azure cloud service provider designated unique identifier for the machine instance.
    uniqueId *string
}
// NewV1beta1CspMachineInstancesGetResponse_items_cspInfoMember2 instantiates a new V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2 and sets the default values.
func NewV1beta1CspMachineInstancesGetResponse_items_cspInfoMember2()(*V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) {
    m := &V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesGetResponse_items_cspInfoMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesGetResponse_items_cspInfoMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesGetResponse_items_cspInfoMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailabilityZones gets the availabilityZones property value. Cloud provider availability zones
// returns a []string when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetAvailabilityZones()([]string) {
    return m.availabilityZones
}
// GetCpuCoreCount gets the cpuCoreCount property value. Number of CPU cores configured.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetCpuCoreCount()(*int32) {
    return m.cpuCoreCount
}
// GetCreatedAt gets the createdAt property value. The time at which the machine instance was launched.
// returns a *Time when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCspRegion gets the cspRegion property value. Cloud provider region
// returns a *string when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetCspRegion()(*string) {
    return m.cspRegion
}
// GetCspTags gets the cspTags property value. List of tags assigned to the resource
// returns a []V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_cspTagsable when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetCspTags()([]V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_cspTagsable) {
    return m.cspTags
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availabilityZones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAvailabilityZones(res)
        }
        return nil
    }
    res["cpuCoreCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuCoreCount(val)
        }
        return nil
    }
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["cspRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspRegion(val)
        }
        return nil
    }
    res["cspTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_cspTagsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_cspTagsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_cspTagsable)
                }
            }
            m.SetCspTags(res)
        }
        return nil
    }
    res["instanceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceType(val)
        }
        return nil
    }
    res["keyPairName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyPairName(val)
        }
        return nil
    }
    res["networkInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_networkInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkInfo(val.(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_networkInfoable))
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val)
        }
        return nil
    }
    res["rootDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootDevice(val)
        }
        return nil
    }
    res["uniqueId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueId(val)
        }
        return nil
    }
    return res
}
// GetInstanceType gets the instanceType property value. A cloud provider designated machine instance type.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetInstanceType()(*string) {
    return m.instanceType
}
// GetKeyPairName gets the keyPairName property value. Name of the key pair associated with the machine instance, if any.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetKeyPairName()(*string) {
    return m.keyPairName
}
// GetNetworkInfo gets the networkInfo property value. Network properties for the machine instance.
// returns a V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_networkInfoable when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetNetworkInfo()(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_networkInfoable) {
    return m.networkInfo
}
// GetPlatform gets the platform property value. A cloud provider designated machine platform.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetPlatform()(*string) {
    return m.platform
}
// GetRootDevice gets the rootDevice property value. The device name for the root volume, for example /dev/sda1.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetRootDevice()(*string) {
    return m.rootDevice
}
// GetUniqueId gets the uniqueId property value. Azure cloud service provider designated unique identifier for the machine instance.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) GetUniqueId()(*string) {
    return m.uniqueId
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCspTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCspTags()))
        for i, v := range m.GetCspTags() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("cspTags", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("networkInfo", m.GetNetworkInfo())
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
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailabilityZones sets the availabilityZones property value. Cloud provider availability zones
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetAvailabilityZones(value []string)() {
    m.availabilityZones = value
}
// SetCpuCoreCount sets the cpuCoreCount property value. Number of CPU cores configured.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetCpuCoreCount(value *int32)() {
    m.cpuCoreCount = value
}
// SetCreatedAt sets the createdAt property value. The time at which the machine instance was launched.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCspRegion sets the cspRegion property value. Cloud provider region
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetCspRegion(value *string)() {
    m.cspRegion = value
}
// SetCspTags sets the cspTags property value. List of tags assigned to the resource
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetCspTags(value []V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_cspTagsable)() {
    m.cspTags = value
}
// SetInstanceType sets the instanceType property value. A cloud provider designated machine instance type.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetInstanceType(value *string)() {
    m.instanceType = value
}
// SetKeyPairName sets the keyPairName property value. Name of the key pair associated with the machine instance, if any.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetKeyPairName(value *string)() {
    m.keyPairName = value
}
// SetNetworkInfo sets the networkInfo property value. Network properties for the machine instance.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetNetworkInfo(value V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_networkInfoable)() {
    m.networkInfo = value
}
// SetPlatform sets the platform property value. A cloud provider designated machine platform.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetPlatform(value *string)() {
    m.platform = value
}
// SetRootDevice sets the rootDevice property value. The device name for the root volume, for example /dev/sda1.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetRootDevice(value *string)() {
    m.rootDevice = value
}
// SetUniqueId sets the uniqueId property value. Azure cloud service provider designated unique identifier for the machine instance.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2) SetUniqueId(value *string)() {
    m.uniqueId = value
}
type V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityZones()([]string)
    GetCpuCoreCount()(*int32)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCspRegion()(*string)
    GetCspTags()([]V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_cspTagsable)
    GetInstanceType()(*string)
    GetKeyPairName()(*string)
    GetNetworkInfo()(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_networkInfoable)
    GetPlatform()(*string)
    GetRootDevice()(*string)
    GetUniqueId()(*string)
    SetAvailabilityZones(value []string)()
    SetCpuCoreCount(value *int32)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCspRegion(value *string)()
    SetCspTags(value []V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_cspTagsable)()
    SetInstanceType(value *string)()
    SetKeyPairName(value *string)()
    SetNetworkInfo(value V1beta1CspMachineInstancesGetResponse_items_cspInfoMember2_networkInfoable)()
    SetPlatform(value *string)()
    SetRootDevice(value *string)()
    SetUniqueId(value *string)()
}
