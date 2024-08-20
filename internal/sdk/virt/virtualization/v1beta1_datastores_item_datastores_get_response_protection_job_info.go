package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo information about the assigned Protection Policy and the Protection Job.
type V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // UUID string uniquely identifying the Protection Job.
    id *string
    // name of the Protection Job.
    name *string
    // Information about the Protection Policy that was used to create the job.
    protectionPolicyInfo V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo_protectionPolicyInfoable
    // Reference to resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo instantiates a new V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo and sets the default values.
func NewV1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo()(*V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) {
    m := &V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["protectionPolicyInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo_protectionPolicyInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionPolicyInfo(val.(V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo_protectionPolicyInfoable))
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
// GetId gets the id property value. UUID string uniquely identifying the Protection Job.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. name of the Protection Job.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) GetName()(*string) {
    return m.name
}
// GetProtectionPolicyInfo gets the protectionPolicyInfo property value. Information about the Protection Policy that was used to create the job.
// returns a V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo_protectionPolicyInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) GetProtectionPolicyInfo()(V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo_protectionPolicyInfoable) {
    return m.protectionPolicyInfo
}
// GetResourceUri gets the resourceUri property value. Reference to resource.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
    {
        err := writer.WriteStringValue("resourceUri", m.GetResourceUri())
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
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. UUID string uniquely identifying the Protection Job.
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. name of the Protection Job.
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) SetName(value *string)() {
    m.name = value
}
// SetProtectionPolicyInfo sets the protectionPolicyInfo property value. Information about the Protection Policy that was used to create the job.
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) SetProtectionPolicyInfo(value V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo_protectionPolicyInfoable)() {
    m.protectionPolicyInfo = value
}
// SetResourceUri sets the resourceUri property value. Reference to resource.
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    GetProtectionPolicyInfo()(V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo_protectionPolicyInfoable)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetId(value *string)()
    SetName(value *string)()
    SetProtectionPolicyInfo(value V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfo_protectionPolicyInfoable)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
