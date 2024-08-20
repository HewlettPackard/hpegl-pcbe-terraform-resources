package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemAddHypervisorClusterPostRequestBody add Hypervisor Cluster Request
type V1beta1SystemsItemAddHypervisorClusterPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies if vSphere Distributed Switch (VDS) should be configured for the hypervisor clusterbeing added.
    configureVds *bool
    // Name of the hypervisor cluster to be added.
    hypervisorClusterName *string
    // vSphere datacenter name where this hypervisor cluster is to be added.
    vsphereDatacenterName *string
}
// NewV1beta1SystemsItemAddHypervisorClusterPostRequestBody instantiates a new V1beta1SystemsItemAddHypervisorClusterPostRequestBody and sets the default values.
func NewV1beta1SystemsItemAddHypervisorClusterPostRequestBody()(*V1beta1SystemsItemAddHypervisorClusterPostRequestBody) {
    m := &V1beta1SystemsItemAddHypervisorClusterPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemAddHypervisorClusterPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemAddHypervisorClusterPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemAddHypervisorClusterPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfigureVds gets the configureVds property value. Specifies if vSphere Distributed Switch (VDS) should be configured for the hypervisor clusterbeing added.
// returns a *bool when successful
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) GetConfigureVds()(*bool) {
    return m.configureVds
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configureVds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigureVds(val)
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
    res["vsphereDatacenterName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVsphereDatacenterName(val)
        }
        return nil
    }
    return res
}
// GetHypervisorClusterName gets the hypervisorClusterName property value. Name of the hypervisor cluster to be added.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) GetHypervisorClusterName()(*string) {
    return m.hypervisorClusterName
}
// GetVsphereDatacenterName gets the vsphereDatacenterName property value. vSphere datacenter name where this hypervisor cluster is to be added.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) GetVsphereDatacenterName()(*string) {
    return m.vsphereDatacenterName
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("configureVds", m.GetConfigureVds())
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
        err := writer.WriteStringValue("vsphereDatacenterName", m.GetVsphereDatacenterName())
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
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfigureVds sets the configureVds property value. Specifies if vSphere Distributed Switch (VDS) should be configured for the hypervisor clusterbeing added.
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) SetConfigureVds(value *bool)() {
    m.configureVds = value
}
// SetHypervisorClusterName sets the hypervisorClusterName property value. Name of the hypervisor cluster to be added.
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) SetHypervisorClusterName(value *string)() {
    m.hypervisorClusterName = value
}
// SetVsphereDatacenterName sets the vsphereDatacenterName property value. vSphere datacenter name where this hypervisor cluster is to be added.
func (m *V1beta1SystemsItemAddHypervisorClusterPostRequestBody) SetVsphereDatacenterName(value *string)() {
    m.vsphereDatacenterName = value
}
type V1beta1SystemsItemAddHypervisorClusterPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigureVds()(*bool)
    GetHypervisorClusterName()(*string)
    GetVsphereDatacenterName()(*string)
    SetConfigureVds(value *bool)()
    SetHypervisorClusterName(value *string)()
    SetVsphereDatacenterName(value *string)()
}
