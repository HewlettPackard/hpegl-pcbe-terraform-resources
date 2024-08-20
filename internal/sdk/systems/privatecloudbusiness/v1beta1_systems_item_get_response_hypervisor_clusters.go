package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemGetResponse_hypervisorClusters hypervisor Cluster Information of the system.
type V1beta1SystemsItemGetResponse_hypervisorClusters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Address of Hypervisor Cluster Manager.
    hypervisorManagerAddress *string
    // Secret information
    hypervisorManagerAdminPasswordSecret V1beta1SystemsItemGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecretable
    // Unique Identifier of Hypervisor Cluster Manager.
    hypervisorManagerId *string
    // Name of Hypervisor Cluster Manager.
    hypervisorManagerName *string
    // Unique Identifier of the Hypervisor Cluster, usually a UUID.
    id *string
    // Name of the Hypervisor Cluster
    name *string
    // Resource URI of the Hypervisor Cluster.
    resourceUri *string
}
// NewV1beta1SystemsItemGetResponse_hypervisorClusters instantiates a new V1beta1SystemsItemGetResponse_hypervisorClusters and sets the default values.
func NewV1beta1SystemsItemGetResponse_hypervisorClusters()(*V1beta1SystemsItemGetResponse_hypervisorClusters) {
    m := &V1beta1SystemsItemGetResponse_hypervisorClusters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemGetResponse_hypervisorClustersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemGetResponse_hypervisorClustersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemGetResponse_hypervisorClusters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerAdminPasswordSecret(val.(V1beta1SystemsItemGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecretable))
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
    return res
}
// GetHypervisorManagerAddress gets the hypervisorManagerAddress property value. Address of Hypervisor Cluster Manager.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) GetHypervisorManagerAddress()(*string) {
    return m.hypervisorManagerAddress
}
// GetHypervisorManagerAdminPasswordSecret gets the hypervisorManagerAdminPasswordSecret property value. Secret information
// returns a V1beta1SystemsItemGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecretable when successful
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) GetHypervisorManagerAdminPasswordSecret()(V1beta1SystemsItemGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecretable) {
    return m.hypervisorManagerAdminPasswordSecret
}
// GetHypervisorManagerId gets the hypervisorManagerId property value. Unique Identifier of Hypervisor Cluster Manager.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) GetHypervisorManagerId()(*string) {
    return m.hypervisorManagerId
}
// GetHypervisorManagerName gets the hypervisorManagerName property value. Name of Hypervisor Cluster Manager.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) GetHypervisorManagerName()(*string) {
    return m.hypervisorManagerName
}
// GetId gets the id property value. Unique Identifier of the Hypervisor Cluster, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the Hypervisor Cluster
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. Resource URI of the Hypervisor Cluster.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) GetResourceUri()(*string) {
    return m.resourceUri
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHypervisorManagerAddress sets the hypervisorManagerAddress property value. Address of Hypervisor Cluster Manager.
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) SetHypervisorManagerAddress(value *string)() {
    m.hypervisorManagerAddress = value
}
// SetHypervisorManagerAdminPasswordSecret sets the hypervisorManagerAdminPasswordSecret property value. Secret information
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) SetHypervisorManagerAdminPasswordSecret(value V1beta1SystemsItemGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecretable)() {
    m.hypervisorManagerAdminPasswordSecret = value
}
// SetHypervisorManagerId sets the hypervisorManagerId property value. Unique Identifier of Hypervisor Cluster Manager.
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) SetHypervisorManagerId(value *string)() {
    m.hypervisorManagerId = value
}
// SetHypervisorManagerName sets the hypervisorManagerName property value. Name of Hypervisor Cluster Manager.
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) SetHypervisorManagerName(value *string)() {
    m.hypervisorManagerName = value
}
// SetId sets the id property value. Unique Identifier of the Hypervisor Cluster, usually a UUID.
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the Hypervisor Cluster
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. Resource URI of the Hypervisor Cluster.
func (m *V1beta1SystemsItemGetResponse_hypervisorClusters) SetResourceUri(value *string)() {
    m.resourceUri = value
}
type V1beta1SystemsItemGetResponse_hypervisorClustersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHypervisorManagerAddress()(*string)
    GetHypervisorManagerAdminPasswordSecret()(V1beta1SystemsItemGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecretable)
    GetHypervisorManagerId()(*string)
    GetHypervisorManagerName()(*string)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    SetHypervisorManagerAddress(value *string)()
    SetHypervisorManagerAdminPasswordSecret(value V1beta1SystemsItemGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecretable)()
    SetHypervisorManagerId(value *string)()
    SetHypervisorManagerName(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
}
