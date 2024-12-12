package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorHostsGetResponse_items_parentInfo parent of this host. It could be a cluster or folder in case of a VMware ESXi Host. For a Hyper-V host it will be cluster or host group.
type V1beta1HypervisorHostsGetResponse_items_parentInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A user-friendly name that identifies the parent of this host.
    displayName *string
    // UUID uniquely identifying the parent of this host.
    id *string
    // Name of the parent of this host.
    name *string
    // The URI reference for this resource.
    resourceUri *string
}
// NewV1beta1HypervisorHostsGetResponse_items_parentInfo instantiates a new V1beta1HypervisorHostsGetResponse_items_parentInfo and sets the default values.
func NewV1beta1HypervisorHostsGetResponse_items_parentInfo()(*V1beta1HypervisorHostsGetResponse_items_parentInfo) {
    m := &V1beta1HypervisorHostsGetResponse_items_parentInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsGetResponse_items_parentInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsGetResponse_items_parentInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsGetResponse_items_parentInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the parent of this host.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
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
    return res
}
// GetId gets the id property value. UUID uniquely identifying the parent of this host.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the parent of this host.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The URI reference for this resource.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
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
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the parent of this host.
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. UUID uniquely identifying the parent of this host.
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the parent of this host.
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The URI reference for this resource.
func (m *V1beta1HypervisorHostsGetResponse_items_parentInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
type V1beta1HypervisorHostsGetResponse_items_parentInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
}
