package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersGetResponse_items_hypervisorHost id, URI to reference the hypervisor host.
type V1beta1SystemsItemServersGetResponse_items_hypervisorHost struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hypervisorClusterId property
    hypervisorClusterId *string
    // The hypervisorClusterName property
    hypervisorClusterName *string
    // The hypervisorHostIp property
    hypervisorHostIp *string
    // The id property
    id *string
    // Is the hypervisor host in maintenance mode?
    maintenanceMode *bool
    // The name property
    name *string
    // The resourceUri property
    resourceUri *string
    // The type property
    typeEscaped *string
}
// NewV1beta1SystemsItemServersGetResponse_items_hypervisorHost instantiates a new V1beta1SystemsItemServersGetResponse_items_hypervisorHost and sets the default values.
func NewV1beta1SystemsItemServersGetResponse_items_hypervisorHost()(*V1beta1SystemsItemServersGetResponse_items_hypervisorHost) {
    m := &V1beta1SystemsItemServersGetResponse_items_hypervisorHost{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersGetResponse_items_hypervisorHostFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersGetResponse_items_hypervisorHostFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersGetResponse_items_hypervisorHost(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hypervisorClusterId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorClusterId(val)
        }
        return nil
    }
    res["hypervisorClusterName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorClusterName(val)
        }
        return nil
    }
    res["hypervisorHostIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorHostIp(val)
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
    res["maintenanceMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintenanceMode(val)
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
// GetHypervisorClusterId gets the hypervisorClusterId property value. The hypervisorClusterId property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetHypervisorClusterId()(*string) {
    return m.hypervisorClusterId
}
// GetHypervisorClusterName gets the hypervisorClusterName property value. The hypervisorClusterName property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetHypervisorClusterName()(*string) {
    return m.hypervisorClusterName
}
// GetHypervisorHostIp gets the hypervisorHostIp property value. The hypervisorHostIp property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetHypervisorHostIp()(*string) {
    return m.hypervisorHostIp
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetId()(*string) {
    return m.id
}
// GetMaintenanceMode gets the maintenanceMode property value. Is the hypervisor host in maintenance mode?
// returns a *bool when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetMaintenanceMode()(*bool) {
    return m.maintenanceMode
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The resourceUri property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hypervisorClusterId", m.GetHypervisorClusterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hypervisorClusterName", m.GetHypervisorClusterName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hypervisorHostIp", m.GetHypervisorHostIp())
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
        err := writer.WriteBoolValue("maintenanceMode", m.GetMaintenanceMode())
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
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHypervisorClusterId sets the hypervisorClusterId property value. The hypervisorClusterId property
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) SetHypervisorClusterId(value *string)() {
    m.hypervisorClusterId = value
}
// SetHypervisorClusterName sets the hypervisorClusterName property value. The hypervisorClusterName property
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) SetHypervisorClusterName(value *string)() {
    m.hypervisorClusterName = value
}
// SetHypervisorHostIp sets the hypervisorHostIp property value. The hypervisorHostIp property
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) SetHypervisorHostIp(value *string)() {
    m.hypervisorHostIp = value
}
// SetId sets the id property value. The id property
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) SetId(value *string)() {
    m.id = value
}
// SetMaintenanceMode sets the maintenanceMode property value. Is the hypervisor host in maintenance mode?
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) SetMaintenanceMode(value *bool)() {
    m.maintenanceMode = value
}
// SetName sets the name property value. The name property
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The resourceUri property
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *V1beta1SystemsItemServersGetResponse_items_hypervisorHost) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1SystemsItemServersGetResponse_items_hypervisorHostable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHypervisorClusterId()(*string)
    GetHypervisorClusterName()(*string)
    GetHypervisorHostIp()(*string)
    GetId()(*string)
    GetMaintenanceMode()(*bool)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetHypervisorClusterId(value *string)()
    SetHypervisorClusterName(value *string)()
    SetHypervisorHostIp(value *string)()
    SetId(value *string)()
    SetMaintenanceMode(value *bool)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
