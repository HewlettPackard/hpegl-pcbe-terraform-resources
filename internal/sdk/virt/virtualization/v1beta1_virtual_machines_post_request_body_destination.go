package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesPostRequestBody_destination specifies where to deploy the virtual machine
type V1beta1VirtualMachinesPostRequestBody_destination struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The UUID of the hypervisor cluster where the virtual machine can be deployed
    clusterId *string
    // The UUID of the hypervisor folder where the virtual machine can be deployed
    folderId *string
    // The UUID of the hypervisor host where the virtual machine can be deployed
    hostId *string
    // The UUID of the hypervisor resource pool where the virtual machine can be deployed
    resourcePoolId *string
}
// NewV1beta1VirtualMachinesPostRequestBody_destination instantiates a new V1beta1VirtualMachinesPostRequestBody_destination and sets the default values.
func NewV1beta1VirtualMachinesPostRequestBody_destination()(*V1beta1VirtualMachinesPostRequestBody_destination) {
    m := &V1beta1VirtualMachinesPostRequestBody_destination{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesPostRequestBody_destinationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesPostRequestBody_destinationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesPostRequestBody_destination(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesPostRequestBody_destination) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClusterId gets the clusterId property value. The UUID of the hypervisor cluster where the virtual machine can be deployed
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_destination) GetClusterId()(*string) {
    return m.clusterId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesPostRequestBody_destination) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clusterId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterId(val)
        }
        return nil
    }
    res["folderId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderId(val)
        }
        return nil
    }
    res["hostId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostId(val)
        }
        return nil
    }
    res["resourcePoolId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourcePoolId(val)
        }
        return nil
    }
    return res
}
// GetFolderId gets the folderId property value. The UUID of the hypervisor folder where the virtual machine can be deployed
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_destination) GetFolderId()(*string) {
    return m.folderId
}
// GetHostId gets the hostId property value. The UUID of the hypervisor host where the virtual machine can be deployed
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_destination) GetHostId()(*string) {
    return m.hostId
}
// GetResourcePoolId gets the resourcePoolId property value. The UUID of the hypervisor resource pool where the virtual machine can be deployed
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_destination) GetResourcePoolId()(*string) {
    return m.resourcePoolId
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesPostRequestBody_destination) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clusterId", m.GetClusterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("folderId", m.GetFolderId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostId", m.GetHostId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourcePoolId", m.GetResourcePoolId())
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
func (m *V1beta1VirtualMachinesPostRequestBody_destination) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClusterId sets the clusterId property value. The UUID of the hypervisor cluster where the virtual machine can be deployed
func (m *V1beta1VirtualMachinesPostRequestBody_destination) SetClusterId(value *string)() {
    m.clusterId = value
}
// SetFolderId sets the folderId property value. The UUID of the hypervisor folder where the virtual machine can be deployed
func (m *V1beta1VirtualMachinesPostRequestBody_destination) SetFolderId(value *string)() {
    m.folderId = value
}
// SetHostId sets the hostId property value. The UUID of the hypervisor host where the virtual machine can be deployed
func (m *V1beta1VirtualMachinesPostRequestBody_destination) SetHostId(value *string)() {
    m.hostId = value
}
// SetResourcePoolId sets the resourcePoolId property value. The UUID of the hypervisor resource pool where the virtual machine can be deployed
func (m *V1beta1VirtualMachinesPostRequestBody_destination) SetResourcePoolId(value *string)() {
    m.resourcePoolId = value
}
type V1beta1VirtualMachinesPostRequestBody_destinationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClusterId()(*string)
    GetFolderId()(*string)
    GetHostId()(*string)
    GetResourcePoolId()(*string)
    SetClusterId(value *string)()
    SetFolderId(value *string)()
    SetHostId(value *string)()
    SetResourcePoolId(value *string)()
}
