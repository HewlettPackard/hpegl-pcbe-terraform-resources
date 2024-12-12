package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus details of the current software update status
type V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Software Catalog version to which the update operation is in progress, if any.
    catalogVersion *string
    // Uniform Resource Identifier (URI) of the parent software update task (asynchronous operation).
    parentTaskUri *string
    // Percentage of the software update progress.
    percentage *int32
    // Uniform Resource Identifier (URI) of the software update task (asynchronous operation).
    taskUri *string
}
// NewV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus instantiates a new V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus and sets the default values.
func NewV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus()(*V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) {
    m := &V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCatalogVersion gets the catalogVersion property value. Software Catalog version to which the update operation is in progress, if any.
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) GetCatalogVersion()(*string) {
    return m.catalogVersion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["catalogVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogVersion(val)
        }
        return nil
    }
    res["parentTaskUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentTaskUri(val)
        }
        return nil
    }
    res["percentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentage(val)
        }
        return nil
    }
    res["taskUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaskUri(val)
        }
        return nil
    }
    return res
}
// GetParentTaskUri gets the parentTaskUri property value. Uniform Resource Identifier (URI) of the parent software update task (asynchronous operation).
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) GetParentTaskUri()(*string) {
    return m.parentTaskUri
}
// GetPercentage gets the percentage property value. Percentage of the software update progress.
// returns a *int32 when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) GetPercentage()(*int32) {
    return m.percentage
}
// GetTaskUri gets the taskUri property value. Uniform Resource Identifier (URI) of the software update task (asynchronous operation).
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) GetTaskUri()(*string) {
    return m.taskUri
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("catalogVersion", m.GetCatalogVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("parentTaskUri", m.GetParentTaskUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("percentage", m.GetPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("taskUri", m.GetTaskUri())
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
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCatalogVersion sets the catalogVersion property value. Software Catalog version to which the update operation is in progress, if any.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) SetCatalogVersion(value *string)() {
    m.catalogVersion = value
}
// SetParentTaskUri sets the parentTaskUri property value. Uniform Resource Identifier (URI) of the parent software update task (asynchronous operation).
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) SetParentTaskUri(value *string)() {
    m.parentTaskUri = value
}
// SetPercentage sets the percentage property value. Percentage of the software update progress.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) SetPercentage(value *int32)() {
    m.percentage = value
}
// SetTaskUri sets the taskUri property value. Uniform Resource Identifier (URI) of the software update task (asynchronous operation).
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatus) SetTaskUri(value *string)() {
    m.taskUri = value
}
type V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCatalogVersion()(*string)
    GetParentTaskUri()(*string)
    GetPercentage()(*int32)
    GetTaskUri()(*string)
    SetCatalogVersion(value *string)()
    SetParentTaskUri(value *string)()
    SetPercentage(value *int32)()
    SetTaskUri(value *string)()
}
