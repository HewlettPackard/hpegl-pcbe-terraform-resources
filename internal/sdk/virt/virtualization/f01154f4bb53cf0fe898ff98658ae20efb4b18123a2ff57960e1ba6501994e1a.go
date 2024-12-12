package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo references to the datacenter that house this virtual machine.
type V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // UUID string uniquely identifier of the datacenter.
    id *string
    // VMware provided moref for the datacenter.
    moref *string
    // VMware provided name for the datacenter.
    name *string
}
// NewV1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo instantiates a new V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo and sets the default values.
func NewV1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo()(*V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) {
    m := &V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["moref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMoref(val)
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
// GetId gets the id property value. UUID string uniquely identifier of the datacenter.
// returns a *string when successful
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) GetId()(*string) {
    return m.id
}
// GetMoref gets the moref property value. VMware provided moref for the datacenter.
// returns a *string when successful
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) GetMoref()(*string) {
    return m.moref
}
// GetName gets the name property value. VMware provided name for the datacenter.
// returns a *string when successful
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("moref", m.GetMoref())
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
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. UUID string uniquely identifier of the datacenter.
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) SetId(value *string)() {
    m.id = value
}
// SetMoref sets the moref property value. VMware provided moref for the datacenter.
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) SetMoref(value *string)() {
    m.moref = value
}
// SetName sets the name property value. VMware provided name for the datacenter.
func (m *V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfo) SetName(value *string)() {
    m.name = value
}
type V1beta1HypervisorClustersItemClusterGetResponse_appInfo_vmware_datacenterInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetMoref()(*string)
    GetName()(*string)
    SetId(value *string)()
    SetMoref(value *string)()
    SetName(value *string)()
}
