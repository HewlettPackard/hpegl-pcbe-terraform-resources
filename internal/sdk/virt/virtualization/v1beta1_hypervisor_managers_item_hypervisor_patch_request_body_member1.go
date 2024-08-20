package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Hypervisor server credentials.
    credentials V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsable
}
// NewV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 instantiates a new V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1()(*V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1) {
    m := &V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCredentials gets the credentials property value. Hypervisor server credentials.
// returns a V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsable when successful
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1) GetCredentials()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsable) {
    return m.credentials
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["credentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentials(val.(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("credentials", m.GetCredentials())
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
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCredentials sets the credentials property value. Hypervisor server credentials.
func (m *V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1) SetCredentials(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsable)() {
    m.credentials = value
}
type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCredentials()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsable)
    SetCredentials(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1_credentialsable)()
}
