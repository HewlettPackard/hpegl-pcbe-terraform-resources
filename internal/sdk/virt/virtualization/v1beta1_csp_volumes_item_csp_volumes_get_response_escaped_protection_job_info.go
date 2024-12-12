package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifies the asset that is associated with the protection job, through which the containing assetis indirectly protected.  This property is null if the protection job is associated directly withthe containing asset.  Thus, this information is populated only when the asset type is"backup-recovery/csp-protection-group".
    assetInfo V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_assetInfoable
    // Information about the protection policy that was used to create the protection job
    protectionPolicyInfo V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_protectionPolicyInfoable
    // The self reference for this resource.
    resourceUri *string
    // Protection status information for each of the protection job's schedules
    scheduleInfo []V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_scheduleInfoable
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo instantiates a new V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo and sets the default values.
func NewV1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo()(*V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) {
    m := &V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssetInfo gets the assetInfo property value. Identifies the asset that is associated with the protection job, through which the containing assetis indirectly protected.  This property is null if the protection job is associated directly withthe containing asset.  Thus, this information is populated only when the asset type is"backup-recovery/csp-protection-group".
// returns a V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_assetInfoable when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) GetAssetInfo()(V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_assetInfoable) {
    return m.assetInfo
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assetInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_assetInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetInfo(val.(V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_assetInfoable))
        }
        return nil
    }
    res["protectionPolicyInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_protectionPolicyInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionPolicyInfo(val.(V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_protectionPolicyInfoable))
        }
        return nil
    }
    res["resourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceUri(val)
        }
        return nil
    }
    res["scheduleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_scheduleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_scheduleInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_scheduleInfoable)
                }
            }
            m.SetScheduleInfo(res)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetProtectionPolicyInfo gets the protectionPolicyInfo property value. Information about the protection policy that was used to create the protection job
// returns a V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_protectionPolicyInfoable when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) GetProtectionPolicyInfo()(V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_protectionPolicyInfoable) {
    return m.protectionPolicyInfo
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetScheduleInfo gets the scheduleInfo property value. Protection status information for each of the protection job's schedules
// returns a []V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_scheduleInfoable when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) GetScheduleInfo()([]V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_scheduleInfoable) {
    return m.scheduleInfo
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("assetInfo", m.GetAssetInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("protectionPolicyInfo", m.GetProtectionPolicyInfo())
        if err != nil {
            return err
        }
    }
    if m.GetScheduleInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScheduleInfo()))
        for i, v := range m.GetScheduleInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("scheduleInfo", cast)
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
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssetInfo sets the assetInfo property value. Identifies the asset that is associated with the protection job, through which the containing assetis indirectly protected.  This property is null if the protection job is associated directly withthe containing asset.  Thus, this information is populated only when the asset type is"backup-recovery/csp-protection-group".
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) SetAssetInfo(value V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_assetInfoable)() {
    m.assetInfo = value
}
// SetProtectionPolicyInfo sets the protectionPolicyInfo property value. Information about the protection policy that was used to create the protection job
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) SetProtectionPolicyInfo(value V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_protectionPolicyInfoable)() {
    m.protectionPolicyInfo = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetScheduleInfo sets the scheduleInfo property value. Protection status information for each of the protection job's schedules
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) SetScheduleInfo(value []V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_scheduleInfoable)() {
    m.scheduleInfo = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssetInfo()(V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_assetInfoable)
    GetProtectionPolicyInfo()(V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_protectionPolicyInfoable)
    GetResourceUri()(*string)
    GetScheduleInfo()([]V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_scheduleInfoable)
    GetTypeEscaped()(*string)
    SetAssetInfo(value V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_assetInfoable)()
    SetProtectionPolicyInfo(value V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_protectionPolicyInfoable)()
    SetResourceUri(value *string)()
    SetScheduleInfo(value []V1beta1CspVolumesItemCspVolumesGetResponse_protectionJobInfo_scheduleInfoable)()
    SetTypeEscaped(value *string)()
}
