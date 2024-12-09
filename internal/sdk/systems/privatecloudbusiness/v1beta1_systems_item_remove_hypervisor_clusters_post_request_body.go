package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody remove Hypervisor Clusters Request
type V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of Unique Identifiers (usually UUIDs) of Hypervisor Clusters to be removed from the system.
    hypervisor_cluster_ids []string
}
// NewV1beta1SystemsItemRemoveHypervisorClustersPostRequestBody instantiates a new V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody and sets the default values.
func NewV1beta1SystemsItemRemoveHypervisorClustersPostRequestBody()(*V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody) {
    m := &V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemRemoveHypervisorClustersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemRemoveHypervisorClustersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemRemoveHypervisorClustersPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hypervisor_cluster_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetHypervisorClusterIds(res)
        }
        return nil
    }
    return res
}
// GetHypervisorClusterIds gets the hypervisor_cluster_ids property value. List of Unique Identifiers (usually UUIDs) of Hypervisor Clusters to be removed from the system.
// returns a []string when successful
func (m *V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody) GetHypervisorClusterIds()([]string) {
    return m.hypervisor_cluster_ids
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHypervisorClusterIds() != nil {
        err := writer.WriteCollectionOfStringValues("hypervisor_cluster_ids", m.GetHypervisorClusterIds())
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
func (m *V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHypervisorClusterIds sets the hypervisor_cluster_ids property value. List of Unique Identifiers (usually UUIDs) of Hypervisor Clusters to be removed from the system.
func (m *V1beta1SystemsItemRemoveHypervisorClustersPostRequestBody) SetHypervisorClusterIds(value []string)() {
    m.hypervisor_cluster_ids = value
}
type V1beta1SystemsItemRemoveHypervisorClustersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHypervisorClusterIds()([]string)
    SetHypervisorClusterIds(value []string)()
}
