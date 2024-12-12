package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials hypervisor server credentials.
type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Password used to access the Hypervisor Manager
    password *string
    // Name of the user used to access the Hypervisor Manager
    username *string
}
// NewV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials instantiates a new V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials()(*V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials) {
    m := &V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. Password used to access the Hypervisor Manager
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials) GetPassword()(*string) {
    return m.password
}
// GetUsername gets the username property value. Name of the user used to access the Hypervisor Manager
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("username", m.GetUsername())
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
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPassword sets the password property value. Password used to access the Hypervisor Manager
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials) SetPassword(value *string)() {
    m.password = value
}
// SetUsername sets the username property value. Name of the user used to access the Hypervisor Manager
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentials) SetUsername(value *string)() {
    m.username = value
}
type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPassword()(*string)
    GetUsername()(*string)
    SetPassword(value *string)()
    SetUsername(value *string)()
}
