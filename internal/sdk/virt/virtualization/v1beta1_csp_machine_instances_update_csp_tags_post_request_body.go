package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A list of CSP tags to be added or updated.
    cspTagsAdded []V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody_cspTagsAddedable
    // A list of tag keys to be removed, if existing.
    cspTagsRemoved []string
}
// NewV1beta1CspMachineInstancesUpdateCspTagsPostRequestBody instantiates a new V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody and sets the default values.
func NewV1beta1CspMachineInstancesUpdateCspTagsPostRequestBody()(*V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody) {
    m := &V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesUpdateCspTagsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesUpdateCspTagsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesUpdateCspTagsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCspTagsAdded gets the cspTagsAdded property value. A list of CSP tags to be added or updated.
// returns a []V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody_cspTagsAddedable when successful
func (m *V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody) GetCspTagsAdded()([]V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody_cspTagsAddedable) {
    return m.cspTagsAdded
}
// GetCspTagsRemoved gets the cspTagsRemoved property value. A list of tag keys to be removed, if existing.
// returns a []string when successful
func (m *V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody) GetCspTagsRemoved()([]string) {
    return m.cspTagsRemoved
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cspTagsAdded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineInstancesUpdateCspTagsPostRequestBody_cspTagsAddedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody_cspTagsAddedable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody_cspTagsAddedable)
                }
            }
            m.SetCspTagsAdded(res)
        }
        return nil
    }
    res["cspTagsRemoved"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCspTagsRemoved(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCspTagsAdded() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCspTagsAdded()))
        for i, v := range m.GetCspTagsAdded() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("cspTagsAdded", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCspTagsRemoved() != nil {
        err := writer.WriteCollectionOfStringValues("cspTagsRemoved", m.GetCspTagsRemoved())
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
func (m *V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCspTagsAdded sets the cspTagsAdded property value. A list of CSP tags to be added or updated.
func (m *V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody) SetCspTagsAdded(value []V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody_cspTagsAddedable)() {
    m.cspTagsAdded = value
}
// SetCspTagsRemoved sets the cspTagsRemoved property value. A list of tag keys to be removed, if existing.
func (m *V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody) SetCspTagsRemoved(value []string)() {
    m.cspTagsRemoved = value
}
type V1beta1CspMachineInstancesUpdateCspTagsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCspTagsAdded()([]V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody_cspTagsAddedable)
    GetCspTagsRemoved()([]string)
    SetCspTagsAdded(value []V1beta1CspMachineInstancesUpdateCspTagsPostRequestBody_cspTagsAddedable)()
    SetCspTagsRemoved(value []string)()
}
