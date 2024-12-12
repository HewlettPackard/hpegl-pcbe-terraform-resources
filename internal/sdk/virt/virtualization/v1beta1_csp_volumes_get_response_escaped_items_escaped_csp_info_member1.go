package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspVolumesGetResponse_items_cspInfoMember1 volume properties whose values are defined by, and synchronized with, the AWS cloud service provider.
type V1beta1CspVolumesGetResponse_items_cspInfoMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Cloud provider availability zone
    availabilityZone *string
    // The time at which the volume was created
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Cloud provider region
    cspRegion *string
    // List of tags assigned to the resource
    cspTags []V1beta1CspVolumesGetResponse_items_cspInfoMember1_cspTagsable
    // The number of I/O operations per second configured for the volume.  The meaning canvary based on the volume type.
    iops *int32
    // Indicates whether the volume is encrypted.
    isEncrypted *bool
    // Volume size in GiB
    sizeInGiB *int32
    // A cloud provider designated volume type
    volumeType *string
}
// NewV1beta1CspVolumesGetResponse_items_cspInfoMember1 instantiates a new V1beta1CspVolumesGetResponse_items_cspInfoMember1 and sets the default values.
func NewV1beta1CspVolumesGetResponse_items_cspInfoMember1()(*V1beta1CspVolumesGetResponse_items_cspInfoMember1) {
    m := &V1beta1CspVolumesGetResponse_items_cspInfoMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspVolumesGetResponse_items_cspInfoMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspVolumesGetResponse_items_cspInfoMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspVolumesGetResponse_items_cspInfoMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailabilityZone gets the availabilityZone property value. Cloud provider availability zone
// returns a *string when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetAvailabilityZone()(*string) {
    return m.availabilityZone
}
// GetCreatedAt gets the createdAt property value. The time at which the volume was created
// returns a *Time when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCspRegion gets the cspRegion property value. Cloud provider region
// returns a *string when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetCspRegion()(*string) {
    return m.cspRegion
}
// GetCspTags gets the cspTags property value. List of tags assigned to the resource
// returns a []V1beta1CspVolumesGetResponse_items_cspInfoMember1_cspTagsable when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetCspTags()([]V1beta1CspVolumesGetResponse_items_cspInfoMember1_cspTagsable) {
    return m.cspTags
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspVolumesGetResponse_items_cspInfoMember1_cspTagsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspVolumesGetResponse_items_cspInfoMember1_cspTagsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspVolumesGetResponse_items_cspInfoMember1_cspTagsable)
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
    res["volumeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeType(val)
        }
        return nil
    }
    return res
}
// GetIops gets the iops property value. The number of I/O operations per second configured for the volume.  The meaning canvary based on the volume type.
// returns a *int32 when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetIops()(*int32) {
    return m.iops
}
// GetIsEncrypted gets the isEncrypted property value. Indicates whether the volume is encrypted.
// returns a *bool when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetIsEncrypted()(*bool) {
    return m.isEncrypted
}
// GetSizeInGiB gets the sizeInGiB property value. Volume size in GiB
// returns a *int32 when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetSizeInGiB()(*int32) {
    return m.sizeInGiB
}
// GetVolumeType gets the volumeType property value. A cloud provider designated volume type
// returns a *string when successful
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) GetVolumeType()(*string) {
    return m.volumeType
}
// Serialize serializes information the current object
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailabilityZone sets the availabilityZone property value. Cloud provider availability zone
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) SetAvailabilityZone(value *string)() {
    m.availabilityZone = value
}
// SetCreatedAt sets the createdAt property value. The time at which the volume was created
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCspRegion sets the cspRegion property value. Cloud provider region
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) SetCspRegion(value *string)() {
    m.cspRegion = value
}
// SetCspTags sets the cspTags property value. List of tags assigned to the resource
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) SetCspTags(value []V1beta1CspVolumesGetResponse_items_cspInfoMember1_cspTagsable)() {
    m.cspTags = value
}
// SetIops sets the iops property value. The number of I/O operations per second configured for the volume.  The meaning canvary based on the volume type.
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) SetIops(value *int32)() {
    m.iops = value
}
// SetIsEncrypted sets the isEncrypted property value. Indicates whether the volume is encrypted.
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// SetSizeInGiB sets the sizeInGiB property value. Volume size in GiB
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) SetSizeInGiB(value *int32)() {
    m.sizeInGiB = value
}
// SetVolumeType sets the volumeType property value. A cloud provider designated volume type
func (m *V1beta1CspVolumesGetResponse_items_cspInfoMember1) SetVolumeType(value *string)() {
    m.volumeType = value
}
type V1beta1CspVolumesGetResponse_items_cspInfoMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityZone()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCspRegion()(*string)
    GetCspTags()([]V1beta1CspVolumesGetResponse_items_cspInfoMember1_cspTagsable)
    GetIops()(*int32)
    GetIsEncrypted()(*bool)
    GetSizeInGiB()(*int32)
    GetVolumeType()(*string)
    SetAvailabilityZone(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCspRegion(value *string)()
    SetCspTags(value []V1beta1CspVolumesGetResponse_items_cspInfoMember1_cspTagsable)()
    SetIops(value *int32)()
    SetIsEncrypted(value *bool)()
    SetSizeInGiB(value *int32)()
    SetVolumeType(value *string)()
}
