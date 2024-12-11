package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorClustersItemClusterGetResponse_appInfo application specific information for this cluster.
type V1beta1HypervisorClustersItemClusterGetResponse_appInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The vmware property
    vmware V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmwareable
}
// NewV1beta1HypervisorClustersItemClusterGetResponse_appInfo instantiates a new V1beta1HypervisorClustersItemClusterGetResponse_appInfo and sets the default values.
func NewV1beta1HypervisorClustersItemClusterGetResponse_appInfo()(*V1beta1HypervisorClustersItemClusterGetResponse_appInfo) {
    m := &V1beta1HypervisorClustersItemClusterGetResponse_appInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorClustersItemClusterGetResponse_appInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorClustersItemClusterGetResponse_appInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorClustersItemClusterGetResponse_appInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["vmware"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmwareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmware(val.(V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmwareable))
        }
        return nil
    }
    return res
}
// GetVmware gets the vmware property value. The vmware property
// returns a V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmwareable when successful
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo) GetVmware()(V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmwareable) {
    return m.vmware
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("vmware", m.GetVmware())
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
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetVmware sets the vmware property value. The vmware property
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo) SetVmware(value V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmwareable)() {
    m.vmware = value
}
type V1beta1HypervisorClustersItemClusterGetResponse_appInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetVmware()(V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmwareable)
    SetVmware(value V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmwareable)()
}