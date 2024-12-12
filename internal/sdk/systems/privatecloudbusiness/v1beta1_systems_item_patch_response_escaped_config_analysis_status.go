package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemPatchResponse_configAnalysisStatus status of last run of configuration analysis job.
type V1beta1SystemsItemPatchResponse_configAnalysisStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Time when configuration analysis job was created.
    createdAt *string
    // Configuration analysis job was created by.
    createdBy *string
    // Time when configuration analysis job was next scheduled.
    nextScheduledRunTime *string
    // Time when configuration analysis job was started.
    ruleRunStartTime *string
    // Count of failed checks.
    totalFailed *float64
    // Count of passed checks.
    totalPassed *float64
    // Count of warning checks.
    totalWarning *float64
}
// NewV1beta1SystemsItemPatchResponse_configAnalysisStatus instantiates a new V1beta1SystemsItemPatchResponse_configAnalysisStatus and sets the default values.
func NewV1beta1SystemsItemPatchResponse_configAnalysisStatus()(*V1beta1SystemsItemPatchResponse_configAnalysisStatus) {
    m := &V1beta1SystemsItemPatchResponse_configAnalysisStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemPatchResponse_configAnalysisStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemPatchResponse_configAnalysisStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemPatchResponse_configAnalysisStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the createdAt property value. Time when configuration analysis job was created.
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) GetCreatedAt()(*string) {
    return m.createdAt
}
// GetCreatedBy gets the createdBy property value. Configuration analysis job was created by.
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["nextScheduledRunTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextScheduledRunTime(val)
        }
        return nil
    }
    res["ruleRunStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleRunStartTime(val)
        }
        return nil
    }
    res["totalFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFailed(val)
        }
        return nil
    }
    res["totalPassed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPassed(val)
        }
        return nil
    }
    res["totalWarning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalWarning(val)
        }
        return nil
    }
    return res
}
// GetNextScheduledRunTime gets the nextScheduledRunTime property value. Time when configuration analysis job was next scheduled.
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) GetNextScheduledRunTime()(*string) {
    return m.nextScheduledRunTime
}
// GetRuleRunStartTime gets the ruleRunStartTime property value. Time when configuration analysis job was started.
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) GetRuleRunStartTime()(*string) {
    return m.ruleRunStartTime
}
// GetTotalFailed gets the totalFailed property value. Count of failed checks.
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) GetTotalFailed()(*float64) {
    return m.totalFailed
}
// GetTotalPassed gets the totalPassed property value. Count of passed checks.
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) GetTotalPassed()(*float64) {
    return m.totalPassed
}
// GetTotalWarning gets the totalWarning property value. Count of warning checks.
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) GetTotalWarning()(*float64) {
    return m.totalWarning
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("nextScheduledRunTime", m.GetNextScheduledRunTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleRunStartTime", m.GetRuleRunStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("totalFailed", m.GetTotalFailed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("totalPassed", m.GetTotalPassed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("totalWarning", m.GetTotalWarning())
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
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the createdAt property value. Time when configuration analysis job was created.
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) SetCreatedAt(value *string)() {
    m.createdAt = value
}
// SetCreatedBy sets the createdBy property value. Configuration analysis job was created by.
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetNextScheduledRunTime sets the nextScheduledRunTime property value. Time when configuration analysis job was next scheduled.
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) SetNextScheduledRunTime(value *string)() {
    m.nextScheduledRunTime = value
}
// SetRuleRunStartTime sets the ruleRunStartTime property value. Time when configuration analysis job was started.
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) SetRuleRunStartTime(value *string)() {
    m.ruleRunStartTime = value
}
// SetTotalFailed sets the totalFailed property value. Count of failed checks.
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) SetTotalFailed(value *float64)() {
    m.totalFailed = value
}
// SetTotalPassed sets the totalPassed property value. Count of passed checks.
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) SetTotalPassed(value *float64)() {
    m.totalPassed = value
}
// SetTotalWarning sets the totalWarning property value. Count of warning checks.
func (m *V1beta1SystemsItemPatchResponse_configAnalysisStatus) SetTotalWarning(value *float64)() {
    m.totalWarning = value
}
type V1beta1SystemsItemPatchResponse_configAnalysisStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetCreatedBy()(*string)
    GetNextScheduledRunTime()(*string)
    GetRuleRunStartTime()(*string)
    GetTotalFailed()(*float64)
    GetTotalPassed()(*float64)
    GetTotalWarning()(*float64)
    SetCreatedAt(value *string)()
    SetCreatedBy(value *string)()
    SetNextScheduledRunTime(value *string)()
    SetRuleRunStartTime(value *string)()
    SetTotalFailed(value *float64)()
    SetTotalPassed(value *float64)()
    SetTotalWarning(value *float64)()
}
