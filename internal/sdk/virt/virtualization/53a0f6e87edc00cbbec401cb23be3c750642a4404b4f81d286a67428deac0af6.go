package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Numeric identifier assigned to this schedule.
    scheduleId *int32
    // The time at which the status was last updated
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo instantiates a new V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo and sets the default values.
func NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo()(*V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo) {
    m := &V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["scheduleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleId(val)
        }
        return nil
    }
    res["updatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetScheduleId gets the scheduleId property value. Numeric identifier assigned to this schedule.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo) GetScheduleId()(*int32) {
    return m.scheduleId
}
// GetUpdatedAt gets the updatedAt property value. The time at which the status was last updated
// returns a *Time when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("scheduleId", m.GetScheduleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedAt", m.GetUpdatedAt())
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
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetScheduleId sets the scheduleId property value. Numeric identifier assigned to this schedule.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo) SetScheduleId(value *int32)() {
    m.scheduleId = value
}
// SetUpdatedAt sets the updatedAt property value. The time at which the status was last updated
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfo) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_protectionJobInfo_scheduleInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetScheduleId()(*int32)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetScheduleId(value *int32)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
