package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemRemoveHypervisorServersPostRequestBody remove Hypervisor Servers Request
type V1beta1SystemsItemRemoveHypervisorServersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of Unique Identifiers (usually UUIDs) of Hypervisor servers to be removed from the system.
    serverIds []string
}
// NewV1beta1SystemsItemRemoveHypervisorServersPostRequestBody instantiates a new V1beta1SystemsItemRemoveHypervisorServersPostRequestBody and sets the default values.
func NewV1beta1SystemsItemRemoveHypervisorServersPostRequestBody()(*V1beta1SystemsItemRemoveHypervisorServersPostRequestBody) {
    m := &V1beta1SystemsItemRemoveHypervisorServersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemRemoveHypervisorServersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemRemoveHypervisorServersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemRemoveHypervisorServersPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemRemoveHypervisorServersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemRemoveHypervisorServersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetServerIds gets the serverIds property value. List of Unique Identifiers (usually UUIDs) of Hypervisor servers to be removed from the system.
// returns a []string when successful
func (m *V1beta1SystemsItemRemoveHypervisorServersPostRequestBody) GetServerIds()([]string) {
    return m.serverIds
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemRemoveHypervisorServersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1SystemsItemRemoveHypervisorServersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetServerIds sets the serverIds property value. List of Unique Identifiers (usually UUIDs) of Hypervisor servers to be removed from the system.
func (m *V1beta1SystemsItemRemoveHypervisorServersPostRequestBody) SetServerIds(value []string)() {
    m.serverIds = value
}
type V1beta1SystemsItemRemoveHypervisorServersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetServerIds()([]string)
    SetServerIds(value []string)()
}
