package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemAddHypervisorServersPostRequestBody add Hypervisor Servers Request
type V1beta1SystemsItemAddHypervisorServersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Type of the device to which hypervisor servers are being added
    deviceType *string
    // ID corresponding to the ESXi password. Applies only to Alletra dHCI.
    esxRootCredentialId *string
    // ID of the hypervisor cluster to which hypervisor servers are being added
    hypervisorClusterId *string
    // ID corresponding to the password for the server iLO account. Applies only to Alletra dHCI.
    iloAdminCredentialId *string
    // Server networks. Applies only to Alletra dHCI.
    serverNetwork []V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable
}
// NewV1beta1SystemsItemAddHypervisorServersPostRequestBody instantiates a new V1beta1SystemsItemAddHypervisorServersPostRequestBody and sets the default values.
func NewV1beta1SystemsItemAddHypervisorServersPostRequestBody()(*V1beta1SystemsItemAddHypervisorServersPostRequestBody) {
    m := &V1beta1SystemsItemAddHypervisorServersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemAddHypervisorServersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemAddHypervisorServersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemAddHypervisorServersPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceType gets the deviceType property value. Type of the device to which hypervisor servers are being added
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) GetDeviceType()(*string) {
    return m.deviceType
}
// GetEsxRootCredentialId gets the esxRootCredentialId property value. ID corresponding to the ESXi password. Applies only to Alletra dHCI.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) GetEsxRootCredentialId()(*string) {
    return m.esxRootCredentialId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val)
        }
        return nil
    }
    res["esxRootCredentialId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEsxRootCredentialId(val)
        }
        return nil
    }
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
    res["iloAdminCredentialId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloAdminCredentialId(val)
        }
        return nil
    }
    res["serverNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable)
                }
            }
            m.SetServerNetwork(res)
        }
        return nil
    }
    return res
}
// GetHypervisorClusterId gets the hypervisorClusterId property value. ID of the hypervisor cluster to which hypervisor servers are being added
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) GetHypervisorClusterId()(*string) {
    return m.hypervisorClusterId
}
// GetIloAdminCredentialId gets the iloAdminCredentialId property value. ID corresponding to the password for the server iLO account. Applies only to Alletra dHCI.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) GetIloAdminCredentialId()(*string) {
    return m.iloAdminCredentialId
}
// GetServerNetwork gets the serverNetwork property value. Server networks. Applies only to Alletra dHCI.
// returns a []V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) GetServerNetwork()([]V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable) {
    return m.serverNetwork
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceType", m.GetDeviceType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("esxRootCredentialId", m.GetEsxRootCredentialId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hypervisorClusterId", m.GetHypervisorClusterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("iloAdminCredentialId", m.GetIloAdminCredentialId())
        if err != nil {
            return err
        }
    }
    if m.GetServerNetwork() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServerNetwork()))
        for i, v := range m.GetServerNetwork() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("serverNetwork", cast)
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
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceType sets the deviceType property value. Type of the device to which hypervisor servers are being added
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) SetDeviceType(value *string)() {
    m.deviceType = value
}
// SetEsxRootCredentialId sets the esxRootCredentialId property value. ID corresponding to the ESXi password. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) SetEsxRootCredentialId(value *string)() {
    m.esxRootCredentialId = value
}
// SetHypervisorClusterId sets the hypervisorClusterId property value. ID of the hypervisor cluster to which hypervisor servers are being added
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) SetHypervisorClusterId(value *string)() {
    m.hypervisorClusterId = value
}
// SetIloAdminCredentialId sets the iloAdminCredentialId property value. ID corresponding to the password for the server iLO account. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) SetIloAdminCredentialId(value *string)() {
    m.iloAdminCredentialId = value
}
// SetServerNetwork sets the serverNetwork property value. Server networks. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody) SetServerNetwork(value []V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable)() {
    m.serverNetwork = value
}
type V1beta1SystemsItemAddHypervisorServersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceType()(*string)
    GetEsxRootCredentialId()(*string)
    GetHypervisorClusterId()(*string)
    GetIloAdminCredentialId()(*string)
    GetServerNetwork()([]V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable)
    SetDeviceType(value *string)()
    SetEsxRootCredentialId(value *string)()
    SetHypervisorClusterId(value *string)()
    SetIloAdminCredentialId(value *string)()
    SetServerNetwork(value []V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable)()
}
