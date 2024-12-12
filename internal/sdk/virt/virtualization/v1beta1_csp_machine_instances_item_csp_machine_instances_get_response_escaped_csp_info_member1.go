package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1 machine instance properties whose values are defined by the AWS cloud service provider.
type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1 struct {
    // ID of the cloud provider access profile associated with the machine instance(for example the ARN of an AWS IAM instance profile), if any.
    accessProfileId *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Cloud provider availability zone
    availabilityZone *string
    // Number of CPU cores configured.
    cpuCoreCount *int32
    // The time at which the machine instance was launched.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Cloud provider region
    cspRegion *string
    // List of tags assigned to the resource
    cspTags []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_cspTagsable
    // A cloud provider designated machine instance type.
    instanceType *string
    // Name of the key pair associated with the machine instance, if any.
    keyPairName *string
    // Network properties for the machine instance.
    networkInfo V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_networkInfoable
    // A cloud provider designated machine platform.
    platform *string
    // The device name for the root volume, for example /dev/sda1.
    rootDevice *string
    // A cloud provider designated virtualization type.
    virtualizationType *string
}
// NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1 instantiates a new V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1 and sets the default values.
func NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1()(*V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) {
    m := &V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1(), nil
}
// GetAccessProfileId gets the accessProfileId property value. ID of the cloud provider access profile associated with the machine instance(for example the ARN of an AWS IAM instance profile), if any.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetAccessProfileId()(*string) {
    return m.accessProfileId
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailabilityZone gets the availabilityZone property value. Cloud provider availability zone
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetAvailabilityZone()(*string) {
    return m.availabilityZone
}
// GetCpuCoreCount gets the cpuCoreCount property value. Number of CPU cores configured.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetCpuCoreCount()(*int32) {
    return m.cpuCoreCount
}
// GetCreatedAt gets the createdAt property value. The time at which the machine instance was launched.
// returns a *Time when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCspRegion gets the cspRegion property value. Cloud provider region
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetCspRegion()(*string) {
    return m.cspRegion
}
// GetCspTags gets the cspTags property value. List of tags assigned to the resource
// returns a []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_cspTagsable when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetCspTags()([]V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_cspTagsable) {
    return m.cspTags
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessProfileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessProfileId(val)
        }
        return nil
    }
    res["availabilityZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityZone(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_cspTagsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_cspTagsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_cspTagsable)
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
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_networkInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkInfo(val.(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_networkInfoable))
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
    res["virtualizationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualizationType(val)
        }
        return nil
    }
    return res
}
// GetInstanceType gets the instanceType property value. A cloud provider designated machine instance type.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetInstanceType()(*string) {
    return m.instanceType
}
// GetKeyPairName gets the keyPairName property value. Name of the key pair associated with the machine instance, if any.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetKeyPairName()(*string) {
    return m.keyPairName
}
// GetNetworkInfo gets the networkInfo property value. Network properties for the machine instance.
// returns a V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_networkInfoable when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetNetworkInfo()(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_networkInfoable) {
    return m.networkInfo
}
// GetPlatform gets the platform property value. A cloud provider designated machine platform.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetPlatform()(*string) {
    return m.platform
}
// GetRootDevice gets the rootDevice property value. The device name for the root volume, for example /dev/sda1.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetRootDevice()(*string) {
    return m.rootDevice
}
// GetVirtualizationType gets the virtualizationType property value. A cloud provider designated virtualization type.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) GetVirtualizationType()(*string) {
    return m.virtualizationType
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAccessProfileId sets the accessProfileId property value. ID of the cloud provider access profile associated with the machine instance(for example the ARN of an AWS IAM instance profile), if any.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetAccessProfileId(value *string)() {
    m.accessProfileId = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailabilityZone sets the availabilityZone property value. Cloud provider availability zone
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetAvailabilityZone(value *string)() {
    m.availabilityZone = value
}
// SetCpuCoreCount sets the cpuCoreCount property value. Number of CPU cores configured.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetCpuCoreCount(value *int32)() {
    m.cpuCoreCount = value
}
// SetCreatedAt sets the createdAt property value. The time at which the machine instance was launched.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCspRegion sets the cspRegion property value. Cloud provider region
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetCspRegion(value *string)() {
    m.cspRegion = value
}
// SetCspTags sets the cspTags property value. List of tags assigned to the resource
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetCspTags(value []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_cspTagsable)() {
    m.cspTags = value
}
// SetInstanceType sets the instanceType property value. A cloud provider designated machine instance type.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetInstanceType(value *string)() {
    m.instanceType = value
}
// SetKeyPairName sets the keyPairName property value. Name of the key pair associated with the machine instance, if any.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetKeyPairName(value *string)() {
    m.keyPairName = value
}
// SetNetworkInfo sets the networkInfo property value. Network properties for the machine instance.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetNetworkInfo(value V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_networkInfoable)() {
    m.networkInfo = value
}
// SetPlatform sets the platform property value. A cloud provider designated machine platform.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetPlatform(value *string)() {
    m.platform = value
}
// SetRootDevice sets the rootDevice property value. The device name for the root volume, for example /dev/sda1.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetRootDevice(value *string)() {
    m.rootDevice = value
}
// SetVirtualizationType sets the virtualizationType property value. A cloud provider designated virtualization type.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1) SetVirtualizationType(value *string)() {
    m.virtualizationType = value
}
type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessProfileId()(*string)
    GetAvailabilityZone()(*string)
    GetCpuCoreCount()(*int32)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCspRegion()(*string)
    GetCspTags()([]V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_cspTagsable)
    GetInstanceType()(*string)
    GetKeyPairName()(*string)
    GetNetworkInfo()(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_networkInfoable)
    GetPlatform()(*string)
    GetRootDevice()(*string)
    GetVirtualizationType()(*string)
    SetAccessProfileId(value *string)()
    SetAvailabilityZone(value *string)()
    SetCpuCoreCount(value *int32)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCspRegion(value *string)()
    SetCspTags(value []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_cspTagsable)()
    SetInstanceType(value *string)()
    SetKeyPairName(value *string)()
    SetNetworkInfo(value V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember1_networkInfoable)()
    SetPlatform(value *string)()
    SetRootDevice(value *string)()
    SetVirtualizationType(value *string)()
}
