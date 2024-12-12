package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsGetResponse_items_computeClusters compute cluster information.
type V1beta1SystemsGetResponse_items_computeClusters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Address of hypervisor Manager in case it is a hypervisor cluster.
    hypervisorManagerAddress *string
    // Secret information
    hypervisorManagerAdminPasswordSecret V1beta1SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecretable
    // Unique identifier of hypervisor Manager in case it is a hypervisor cluster.
    hypervisorManagerId *string
    // Name of hypervisor Manager in case it is a hypervisor cluster.
    hypervisorManagerName *string
    // Unique identifier of the compute cluster, usually a UUID.
    id *string
    // Name of the compute cluster.
    name *string
    // Resource URI of the compute cluster.
    resourceUri *string
    // A custom URL for the compute cluster - e.g. to deep link to an external compute cluster management console.
    userUrl *string
}
// NewV1beta1SystemsGetResponse_items_computeClusters instantiates a new V1beta1SystemsGetResponse_items_computeClusters and sets the default values.
func NewV1beta1SystemsGetResponse_items_computeClusters()(*V1beta1SystemsGetResponse_items_computeClusters) {
    m := &V1beta1SystemsGetResponse_items_computeClusters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsGetResponse_items_computeClustersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsGetResponse_items_computeClustersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsGetResponse_items_computeClusters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hypervisorManagerAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerAddress(val)
        }
        return nil
    }
    res["hypervisorManagerAdminPasswordSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerAdminPasswordSecret(val.(V1beta1SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecretable))
        }
        return nil
    }
    res["hypervisorManagerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerId(val)
        }
        return nil
    }
    res["hypervisorManagerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerName(val)
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
    res["userUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserUrl(val)
        }
        return nil
    }
    return res
}
// GetHypervisorManagerAddress gets the hypervisorManagerAddress property value. Address of hypervisor Manager in case it is a hypervisor cluster.
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetHypervisorManagerAddress()(*string) {
    return m.hypervisorManagerAddress
}
// GetHypervisorManagerAdminPasswordSecret gets the hypervisorManagerAdminPasswordSecret property value. Secret information
// returns a V1beta1SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecretable when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetHypervisorManagerAdminPasswordSecret()(V1beta1SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecretable) {
    return m.hypervisorManagerAdminPasswordSecret
}
// GetHypervisorManagerId gets the hypervisorManagerId property value. Unique identifier of hypervisor Manager in case it is a hypervisor cluster.
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetHypervisorManagerId()(*string) {
    return m.hypervisorManagerId
}
// GetHypervisorManagerName gets the hypervisorManagerName property value. Name of hypervisor Manager in case it is a hypervisor cluster.
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetHypervisorManagerName()(*string) {
    return m.hypervisorManagerName
}
// GetId gets the id property value. Unique identifier of the compute cluster, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the compute cluster.
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. Resource URI of the compute cluster.
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetUserUrl gets the userUrl property value. A custom URL for the compute cluster - e.g. to deep link to an external compute cluster management console.
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_computeClusters) GetUserUrl()(*string) {
    return m.userUrl
}
// Serialize serializes information the current object
func (m *V1beta1SystemsGetResponse_items_computeClusters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hypervisorManagerAddress", m.GetHypervisorManagerAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hypervisorManagerAdminPasswordSecret", m.GetHypervisorManagerAdminPasswordSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hypervisorManagerId", m.GetHypervisorManagerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hypervisorManagerName", m.GetHypervisorManagerName())
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
        err := writer.WriteStringValue("userUrl", m.GetUserUrl())
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
func (m *V1beta1SystemsGetResponse_items_computeClusters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHypervisorManagerAddress sets the hypervisorManagerAddress property value. Address of hypervisor Manager in case it is a hypervisor cluster.
func (m *V1beta1SystemsGetResponse_items_computeClusters) SetHypervisorManagerAddress(value *string)() {
    m.hypervisorManagerAddress = value
}
// SetHypervisorManagerAdminPasswordSecret sets the hypervisorManagerAdminPasswordSecret property value. Secret information
func (m *V1beta1SystemsGetResponse_items_computeClusters) SetHypervisorManagerAdminPasswordSecret(value V1beta1SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecretable)() {
    m.hypervisorManagerAdminPasswordSecret = value
}
// SetHypervisorManagerId sets the hypervisorManagerId property value. Unique identifier of hypervisor Manager in case it is a hypervisor cluster.
func (m *V1beta1SystemsGetResponse_items_computeClusters) SetHypervisorManagerId(value *string)() {
    m.hypervisorManagerId = value
}
// SetHypervisorManagerName sets the hypervisorManagerName property value. Name of hypervisor Manager in case it is a hypervisor cluster.
func (m *V1beta1SystemsGetResponse_items_computeClusters) SetHypervisorManagerName(value *string)() {
    m.hypervisorManagerName = value
}
// SetId sets the id property value. Unique identifier of the compute cluster, usually a UUID.
func (m *V1beta1SystemsGetResponse_items_computeClusters) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the compute cluster.
func (m *V1beta1SystemsGetResponse_items_computeClusters) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. Resource URI of the compute cluster.
func (m *V1beta1SystemsGetResponse_items_computeClusters) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetUserUrl sets the userUrl property value. A custom URL for the compute cluster - e.g. to deep link to an external compute cluster management console.
func (m *V1beta1SystemsGetResponse_items_computeClusters) SetUserUrl(value *string)() {
    m.userUrl = value
}
type V1beta1SystemsGetResponse_items_computeClustersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHypervisorManagerAddress()(*string)
    GetHypervisorManagerAdminPasswordSecret()(V1beta1SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecretable)
    GetHypervisorManagerId()(*string)
    GetHypervisorManagerName()(*string)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetUserUrl()(*string)
    SetHypervisorManagerAddress(value *string)()
    SetHypervisorManagerAdminPasswordSecret(value V1beta1SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecretable)()
    SetHypervisorManagerId(value *string)()
    SetHypervisorManagerName(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetUserUrl(value *string)()
}
