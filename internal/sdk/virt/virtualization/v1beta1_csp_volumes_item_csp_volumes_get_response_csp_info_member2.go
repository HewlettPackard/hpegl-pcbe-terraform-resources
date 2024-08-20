package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2 volume properties whose values are defined by, and synchronized with, the Azure cloud service provider.
type V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Cloud provider availability zones
    availabilityZones []string
    // The time at which the volume was created
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Cloud provider region
    cspRegion *string
    // List of tags assigned to the resource
    cspTags []V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2_cspTagsable
    // The number of I/O operations per second configured for the volume.  The meaning canvary based on the volume type.
    iops *int32
    // Indicates whether the volume is encrypted.
    isEncrypted *bool
    // Volume size in GiB
    sizeInGiB *int32
    // Azure cloud service provider designated sku name
    skuName *string
    // Azure cloud service provider designated unique identifier for the volume.
    uniqueId *string
}
// NewV1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2 instantiates a new V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2 and sets the default values.
func NewV1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2()(*V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) {
    m := &V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailabilityZones gets the availabilityZones property value. Cloud provider availability zones
// returns a []string when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetAvailabilityZones()([]string) {
    return m.availabilityZones
}
// GetCreatedAt gets the createdAt property value. The time at which the volume was created
// returns a *Time when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCspRegion gets the cspRegion property value. Cloud provider region
// returns a *string when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetCspRegion()(*string) {
    return m.cspRegion
}
// GetCspTags gets the cspTags property value. List of tags assigned to the resource
// returns a []V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2_cspTagsable when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetCspTags()([]V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2_cspTagsable) {
    return m.cspTags
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2_cspTagsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2_cspTagsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2_cspTagsable)
                }
            }
            m.SetCspTags(res)
        }
        return nil
    }
    res["iops"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIops(val)
        }
        return nil
    }
    res["isEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["sizeInGiB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInGiB(val)
        }
        return nil
    }
    res["skuName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuName(val)
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
// GetIops gets the iops property value. The number of I/O operations per second configured for the volume.  The meaning canvary based on the volume type.
// returns a *int32 when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetIops()(*int32) {
    return m.iops
}
// GetIsEncrypted gets the isEncrypted property value. Indicates whether the volume is encrypted.
// returns a *bool when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetIsEncrypted()(*bool) {
    return m.isEncrypted
}
// GetSizeInGiB gets the sizeInGiB property value. Volume size in GiB
// returns a *int32 when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetSizeInGiB()(*int32) {
    return m.sizeInGiB
}
// GetSkuName gets the skuName property value. Azure cloud service provider designated sku name
// returns a *string when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetSkuName()(*string) {
    return m.skuName
}
// GetUniqueId gets the uniqueId property value. Azure cloud service provider designated unique identifier for the volume.
// returns a *string when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) GetUniqueId()(*string) {
    return m.uniqueId
}
// Serialize serializes information the current object
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailabilityZones sets the availabilityZones property value. Cloud provider availability zones
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetAvailabilityZones(value []string)() {
    m.availabilityZones = value
}
// SetCreatedAt sets the createdAt property value. The time at which the volume was created
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCspRegion sets the cspRegion property value. Cloud provider region
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetCspRegion(value *string)() {
    m.cspRegion = value
}
// SetCspTags sets the cspTags property value. List of tags assigned to the resource
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetCspTags(value []V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2_cspTagsable)() {
    m.cspTags = value
}
// SetIops sets the iops property value. The number of I/O operations per second configured for the volume.  The meaning canvary based on the volume type.
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetIops(value *int32)() {
    m.iops = value
}
// SetIsEncrypted sets the isEncrypted property value. Indicates whether the volume is encrypted.
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// SetSizeInGiB sets the sizeInGiB property value. Volume size in GiB
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetSizeInGiB(value *int32)() {
    m.sizeInGiB = value
}
// SetSkuName sets the skuName property value. Azure cloud service provider designated sku name
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetSkuName(value *string)() {
    m.skuName = value
}
// SetUniqueId sets the uniqueId property value. Azure cloud service provider designated unique identifier for the volume.
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2) SetUniqueId(value *string)() {
    m.uniqueId = value
}
type V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityZones()([]string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCspRegion()(*string)
    GetCspTags()([]V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2_cspTagsable)
    GetIops()(*int32)
    GetIsEncrypted()(*bool)
    GetSizeInGiB()(*int32)
    GetSkuName()(*string)
    GetUniqueId()(*string)
    SetAvailabilityZones(value []string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCspRegion(value *string)()
    SetCspTags(value []V1beta1CspVolumesItemCspVolumesGetResponse_cspInfoMember2_cspTagsable)()
    SetIops(value *int32)()
    SetIsEncrypted(value *bool)()
    SetSizeInGiB(value *int32)()
    SetSkuName(value *string)()
    SetUniqueId(value *string)()
}
