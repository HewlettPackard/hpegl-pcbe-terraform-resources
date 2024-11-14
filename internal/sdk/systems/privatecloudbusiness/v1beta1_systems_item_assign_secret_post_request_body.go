package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemAssignSecretPostRequestBody reassign secret Request
type V1beta1SystemsItemAssignSecretPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifier of the hypervisor cluster. If provided along with server IDs or If only the cluster ID is specified and no server IDs are provided, the operation will apply to all servers within the specified cluster. This behavior allows the API to target the entire cluster, depending on the provided parameters.
    hypervisorClusterId *string
    // Indicates whether device updates should be ignored.This is useful when a user creates a new secret with the same password as the device and wants to update the secret details without changing the device's password. By setting this field to true, the user can ensure that the existing password remains unchanged while still managing the secret associated with it.
    ignoreDeviceUpdate *bool
    // Identifier of the secret to be assigned to the servers.
    secretId *string
    // List of server UUIDs. If serverIds are specified, the operation will apply to all mentioned servers. If hypervisorClusterId is also provided then the operation will apply to all servers within the specified cluster irrespective of servers provided in serverIds.
    serverIds []string
}
// NewV1beta1SystemsItemAssignSecretPostRequestBody instantiates a new V1beta1SystemsItemAssignSecretPostRequestBody and sets the default values.
func NewV1beta1SystemsItemAssignSecretPostRequestBody()(*V1beta1SystemsItemAssignSecretPostRequestBody) {
    m := &V1beta1SystemsItemAssignSecretPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemAssignSecretPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemAssignSecretPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemAssignSecretPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["ignoreDeviceUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreDeviceUpdate(val)
        }
        return nil
    }
    res["secretId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretId(val)
        }
        return nil
    }
    res["serverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetServerIds(res)
        }
        return nil
    }
    return res
}
// GetHypervisorClusterId gets the hypervisorClusterId property value. Identifier of the hypervisor cluster. If provided along with server IDs or If only the cluster ID is specified and no server IDs are provided, the operation will apply to all servers within the specified cluster. This behavior allows the API to target the entire cluster, depending on the provided parameters.
// returns a *string when successful
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) GetHypervisorClusterId()(*string) {
    return m.hypervisorClusterId
}
// GetIgnoreDeviceUpdate gets the ignoreDeviceUpdate property value. Indicates whether device updates should be ignored.This is useful when a user creates a new secret with the same password as the device and wants to update the secret details without changing the device's password. By setting this field to true, the user can ensure that the existing password remains unchanged while still managing the secret associated with it.
// returns a *bool when successful
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) GetIgnoreDeviceUpdate()(*bool) {
    return m.ignoreDeviceUpdate
}
// GetSecretId gets the secretId property value. Identifier of the secret to be assigned to the servers.
// returns a *string when successful
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) GetSecretId()(*string) {
    return m.secretId
}
// GetServerIds gets the serverIds property value. List of server UUIDs. If serverIds are specified, the operation will apply to all mentioned servers. If hypervisorClusterId is also provided then the operation will apply to all servers within the specified cluster irrespective of servers provided in serverIds.
// returns a []string when successful
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) GetServerIds()([]string) {
    return m.serverIds
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hypervisorClusterId", m.GetHypervisorClusterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignoreDeviceUpdate", m.GetIgnoreDeviceUpdate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secretId", m.GetSecretId())
        if err != nil {
            return err
        }
    }
    if m.GetServerIds() != nil {
        err := writer.WriteCollectionOfStringValues("serverIds", m.GetServerIds())
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
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHypervisorClusterId sets the hypervisorClusterId property value. Identifier of the hypervisor cluster. If provided along with server IDs or If only the cluster ID is specified and no server IDs are provided, the operation will apply to all servers within the specified cluster. This behavior allows the API to target the entire cluster, depending on the provided parameters.
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) SetHypervisorClusterId(value *string)() {
    m.hypervisorClusterId = value
}
// SetIgnoreDeviceUpdate sets the ignoreDeviceUpdate property value. Indicates whether device updates should be ignored.This is useful when a user creates a new secret with the same password as the device and wants to update the secret details without changing the device's password. By setting this field to true, the user can ensure that the existing password remains unchanged while still managing the secret associated with it.
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) SetIgnoreDeviceUpdate(value *bool)() {
    m.ignoreDeviceUpdate = value
}
// SetSecretId sets the secretId property value. Identifier of the secret to be assigned to the servers.
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) SetSecretId(value *string)() {
    m.secretId = value
}
// SetServerIds sets the serverIds property value. List of server UUIDs. If serverIds are specified, the operation will apply to all mentioned servers. If hypervisorClusterId is also provided then the operation will apply to all servers within the specified cluster irrespective of servers provided in serverIds.
func (m *V1beta1SystemsItemAssignSecretPostRequestBody) SetServerIds(value []string)() {
    m.serverIds = value
}
type V1beta1SystemsItemAssignSecretPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHypervisorClusterId()(*string)
    GetIgnoreDeviceUpdate()(*bool)
    GetSecretId()(*string)
    GetServerIds()([]string)
    SetHypervisorClusterId(value *string)()
    SetIgnoreDeviceUpdate(value *bool)()
    SetSecretId(value *string)()
    SetServerIds(value []string)()
}
