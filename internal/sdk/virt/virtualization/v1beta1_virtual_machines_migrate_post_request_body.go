package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesMigratePostRequestBody request body for Migrate API
type V1beta1VirtualMachinesMigratePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An identifier for the Hypervisor Manager.
    hypervisorManagerId *string
    // Target Datastore ID to which VMs will be migrated.
    targetDatastoreId *string
    // An identifier for the target Hypervisor cluster.
    targetHypervisorClusterId *string
}
// NewV1beta1VirtualMachinesMigratePostRequestBody instantiates a new V1beta1VirtualMachinesMigratePostRequestBody and sets the default values.
func NewV1beta1VirtualMachinesMigratePostRequestBody()(*V1beta1VirtualMachinesMigratePostRequestBody) {
    m := &V1beta1VirtualMachinesMigratePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesMigratePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesMigratePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesMigratePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesMigratePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesMigratePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["targetDatastoreId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDatastoreId(val)
        }
        return nil
    }
    res["targetHypervisorClusterId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetHypervisorClusterId(val)
        }
        return nil
    }
    return res
}
// GetHypervisorManagerId gets the hypervisorManagerId property value. An identifier for the Hypervisor Manager.
// returns a *string when successful
func (m *V1beta1VirtualMachinesMigratePostRequestBody) GetHypervisorManagerId()(*string) {
    return m.hypervisorManagerId
}
// GetTargetDatastoreId gets the targetDatastoreId property value. Target Datastore ID to which VMs will be migrated.
// returns a *string when successful
func (m *V1beta1VirtualMachinesMigratePostRequestBody) GetTargetDatastoreId()(*string) {
    return m.targetDatastoreId
}
// GetTargetHypervisorClusterId gets the targetHypervisorClusterId property value. An identifier for the target Hypervisor cluster.
// returns a *string when successful
func (m *V1beta1VirtualMachinesMigratePostRequestBody) GetTargetHypervisorClusterId()(*string) {
    return m.targetHypervisorClusterId
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesMigratePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hypervisorManagerId", m.GetHypervisorManagerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetDatastoreId", m.GetTargetDatastoreId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetHypervisorClusterId", m.GetTargetHypervisorClusterId())
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
func (m *V1beta1VirtualMachinesMigratePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHypervisorManagerId sets the hypervisorManagerId property value. An identifier for the Hypervisor Manager.
func (m *V1beta1VirtualMachinesMigratePostRequestBody) SetHypervisorManagerId(value *string)() {
    m.hypervisorManagerId = value
}
// SetTargetDatastoreId sets the targetDatastoreId property value. Target Datastore ID to which VMs will be migrated.
func (m *V1beta1VirtualMachinesMigratePostRequestBody) SetTargetDatastoreId(value *string)() {
    m.targetDatastoreId = value
}
// SetTargetHypervisorClusterId sets the targetHypervisorClusterId property value. An identifier for the target Hypervisor cluster.
func (m *V1beta1VirtualMachinesMigratePostRequestBody) SetTargetHypervisorClusterId(value *string)() {
    m.targetHypervisorClusterId = value
}
type V1beta1VirtualMachinesMigratePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHypervisorManagerId()(*string)
    GetTargetDatastoreId()(*string)
    GetTargetHypervisorClusterId()(*string)
    SetHypervisorManagerId(value *string)()
    SetTargetDatastoreId(value *string)()
    SetTargetHypervisorClusterId(value *string)()
}
