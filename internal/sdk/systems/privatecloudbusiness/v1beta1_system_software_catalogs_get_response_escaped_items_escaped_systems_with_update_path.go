package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath system id and Name with a list of hypervisor clusters
type V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hypervisorClusters property
    hypervisorClusters []V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath_hypervisorClustersable
    // Unique Identifier of the system, usually a UUID.
    id *string
    // Name of the system.
    name *string
}
// NewV1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath instantiates a new V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath and sets the default values.
func NewV1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath()(*V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) {
    m := &V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePathFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePathFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hypervisorClusters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath_hypervisorClustersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath_hypervisorClustersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath_hypervisorClustersable)
                }
            }
            m.SetHypervisorClusters(res)
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
    return res
}
// GetHypervisorClusters gets the hypervisorClusters property value. The hypervisorClusters property
// returns a []V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath_hypervisorClustersable when successful
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) GetHypervisorClusters()([]V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath_hypervisorClustersable) {
    return m.hypervisorClusters
}
// GetId gets the id property value. Unique Identifier of the system, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the system.
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHypervisorClusters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHypervisorClusters()))
        for i, v := range m.GetHypervisorClusters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("hypervisorClusters", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHypervisorClusters sets the hypervisorClusters property value. The hypervisorClusters property
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) SetHypervisorClusters(value []V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath_hypervisorClustersable)() {
    m.hypervisorClusters = value
}
// SetId sets the id property value. Unique Identifier of the system, usually a UUID.
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the system.
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath) SetName(value *string)() {
    m.name = value
}
type V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePathable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHypervisorClusters()([]V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath_hypervisorClustersable)
    GetId()(*string)
    GetName()(*string)
    SetHypervisorClusters(value []V1beta1SystemSoftwareCatalogsGetResponse_items_systemsWithUpdatePath_hypervisorClustersable)()
    SetId(value *string)()
    SetName(value *string)()
}
