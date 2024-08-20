package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesPostRequestBody deploys one or more virtual machines using specified template and storage provisioning policy.
type V1beta1VirtualMachinesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies where to deploy the virtual machine
    destination V1beta1VirtualMachinesPostRequestBody_destinationable
    // Specifies the hypervisor image information using which the virtual machine is deployed
    imageSource V1beta1VirtualMachinesPostRequestBody_imageSourceable
    // Specifies name and the target network to use for deployment
    networkConfig V1beta1VirtualMachinesPostRequestBody_networkConfigable
    // Unique identifier of a protection policy
    protectionPolicyId *string
    // Specifies the storage configurations for a virtual machine
    storageConfig V1beta1VirtualMachinesPostRequestBody_storageConfigable
    // Defines the virtual machine configurations
    vmConfig V1beta1VirtualMachinesPostRequestBody_vmConfigable
}
// NewV1beta1VirtualMachinesPostRequestBody instantiates a new V1beta1VirtualMachinesPostRequestBody and sets the default values.
func NewV1beta1VirtualMachinesPostRequestBody()(*V1beta1VirtualMachinesPostRequestBody) {
    m := &V1beta1VirtualMachinesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDestination gets the destination property value. Specifies where to deploy the virtual machine
// returns a V1beta1VirtualMachinesPostRequestBody_destinationable when successful
func (m *V1beta1VirtualMachinesPostRequestBody) GetDestination()(V1beta1VirtualMachinesPostRequestBody_destinationable) {
    return m.destination
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesPostRequestBody_destinationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestination(val.(V1beta1VirtualMachinesPostRequestBody_destinationable))
        }
        return nil
    }
    res["imageSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesPostRequestBody_imageSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageSource(val.(V1beta1VirtualMachinesPostRequestBody_imageSourceable))
        }
        return nil
    }
    res["networkConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesPostRequestBody_networkConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkConfig(val.(V1beta1VirtualMachinesPostRequestBody_networkConfigable))
        }
        return nil
    }
    res["protectionPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionPolicyId(val)
        }
        return nil
    }
    res["storageConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesPostRequestBody_storageConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageConfig(val.(V1beta1VirtualMachinesPostRequestBody_storageConfigable))
        }
        return nil
    }
    res["vmConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesPostRequestBody_vmConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmConfig(val.(V1beta1VirtualMachinesPostRequestBody_vmConfigable))
        }
        return nil
    }
    return res
}
// GetImageSource gets the imageSource property value. Specifies the hypervisor image information using which the virtual machine is deployed
// returns a V1beta1VirtualMachinesPostRequestBody_imageSourceable when successful
func (m *V1beta1VirtualMachinesPostRequestBody) GetImageSource()(V1beta1VirtualMachinesPostRequestBody_imageSourceable) {
    return m.imageSource
}
// GetNetworkConfig gets the networkConfig property value. Specifies name and the target network to use for deployment
// returns a V1beta1VirtualMachinesPostRequestBody_networkConfigable when successful
func (m *V1beta1VirtualMachinesPostRequestBody) GetNetworkConfig()(V1beta1VirtualMachinesPostRequestBody_networkConfigable) {
    return m.networkConfig
}
// GetProtectionPolicyId gets the protectionPolicyId property value. Unique identifier of a protection policy
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody) GetProtectionPolicyId()(*string) {
    return m.protectionPolicyId
}
// GetStorageConfig gets the storageConfig property value. Specifies the storage configurations for a virtual machine
// returns a V1beta1VirtualMachinesPostRequestBody_storageConfigable when successful
func (m *V1beta1VirtualMachinesPostRequestBody) GetStorageConfig()(V1beta1VirtualMachinesPostRequestBody_storageConfigable) {
    return m.storageConfig
}
// GetVmConfig gets the vmConfig property value. Defines the virtual machine configurations
// returns a V1beta1VirtualMachinesPostRequestBody_vmConfigable when successful
func (m *V1beta1VirtualMachinesPostRequestBody) GetVmConfig()(V1beta1VirtualMachinesPostRequestBody_vmConfigable) {
    return m.vmConfig
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("destination", m.GetDestination())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("imageSource", m.GetImageSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("networkConfig", m.GetNetworkConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("protectionPolicyId", m.GetProtectionPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storageConfig", m.GetStorageConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vmConfig", m.GetVmConfig())
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
func (m *V1beta1VirtualMachinesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDestination sets the destination property value. Specifies where to deploy the virtual machine
func (m *V1beta1VirtualMachinesPostRequestBody) SetDestination(value V1beta1VirtualMachinesPostRequestBody_destinationable)() {
    m.destination = value
}
// SetImageSource sets the imageSource property value. Specifies the hypervisor image information using which the virtual machine is deployed
func (m *V1beta1VirtualMachinesPostRequestBody) SetImageSource(value V1beta1VirtualMachinesPostRequestBody_imageSourceable)() {
    m.imageSource = value
}
// SetNetworkConfig sets the networkConfig property value. Specifies name and the target network to use for deployment
func (m *V1beta1VirtualMachinesPostRequestBody) SetNetworkConfig(value V1beta1VirtualMachinesPostRequestBody_networkConfigable)() {
    m.networkConfig = value
}
// SetProtectionPolicyId sets the protectionPolicyId property value. Unique identifier of a protection policy
func (m *V1beta1VirtualMachinesPostRequestBody) SetProtectionPolicyId(value *string)() {
    m.protectionPolicyId = value
}
// SetStorageConfig sets the storageConfig property value. Specifies the storage configurations for a virtual machine
func (m *V1beta1VirtualMachinesPostRequestBody) SetStorageConfig(value V1beta1VirtualMachinesPostRequestBody_storageConfigable)() {
    m.storageConfig = value
}
// SetVmConfig sets the vmConfig property value. Defines the virtual machine configurations
func (m *V1beta1VirtualMachinesPostRequestBody) SetVmConfig(value V1beta1VirtualMachinesPostRequestBody_vmConfigable)() {
    m.vmConfig = value
}
type V1beta1VirtualMachinesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestination()(V1beta1VirtualMachinesPostRequestBody_destinationable)
    GetImageSource()(V1beta1VirtualMachinesPostRequestBody_imageSourceable)
    GetNetworkConfig()(V1beta1VirtualMachinesPostRequestBody_networkConfigable)
    GetProtectionPolicyId()(*string)
    GetStorageConfig()(V1beta1VirtualMachinesPostRequestBody_storageConfigable)
    GetVmConfig()(V1beta1VirtualMachinesPostRequestBody_vmConfigable)
    SetDestination(value V1beta1VirtualMachinesPostRequestBody_destinationable)()
    SetImageSource(value V1beta1VirtualMachinesPostRequestBody_imageSourceable)()
    SetNetworkConfig(value V1beta1VirtualMachinesPostRequestBody_networkConfigable)()
    SetProtectionPolicyId(value *string)()
    SetStorageConfig(value V1beta1VirtualMachinesPostRequestBody_storageConfigable)()
    SetVmConfig(value V1beta1VirtualMachinesPostRequestBody_vmConfigable)()
}
