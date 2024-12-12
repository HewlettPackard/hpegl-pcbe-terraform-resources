package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters hypervisor Cluster id and Name
type V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique Identifier of the Hypervisor Cluster, usually a UUID.
    hypervisorClusterId *string
    // Name of the Hypervisor Cluster.
    hypervisorClusterName *string
}
// NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters instantiates a new V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters and sets the default values.
func NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters()(*V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters) {
    m := &V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClustersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClustersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetHypervisorClusterId gets the hypervisorClusterId property value. Unique Identifier of the Hypervisor Cluster, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters) GetHypervisorClusterId()(*string) {
    return m.hypervisorClusterId
}
// GetHypervisorClusterName gets the hypervisorClusterName property value. Name of the Hypervisor Cluster.
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters) GetHypervisorClusterName()(*string) {
    return m.hypervisorClusterName
}
// Serialize serializes information the current object
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHypervisorClusterId sets the hypervisorClusterId property value. Unique Identifier of the Hypervisor Cluster, usually a UUID.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters) SetHypervisorClusterId(value *string)() {
    m.hypervisorClusterId = value
}
// SetHypervisorClusterName sets the hypervisorClusterName property value. Name of the Hypervisor Cluster.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClusters) SetHypervisorClusterName(value *string)() {
    m.hypervisorClusterName = value
}
type V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePath_hypervisorClustersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHypervisorClusterId()(*string)
    GetHypervisorClusterName()(*string)
    SetHypervisorClusterId(value *string)()
    SetHypervisorClusterName(value *string)()
}
