package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemSoftwareUpdatePostRequestBody request parameters for software update operation on the system.
type V1beta1SystemsItemSoftwareUpdatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Version of the target software catalog for the software precheck/update request.
    catalogVersion *string
    // List of Unique Identifiers (usually UUIDs) of Hypervisor Clusters in the system for which software update operation is to be initiated. If omitted, software update operation will be initiated on all Hypervisor Clusters available in the system.
    hypervisorClusterIds []string
    // Boolean flag to ignore precheck warnings and proceed with software update operation.
    ignorePrecheckWarning *bool
}
// NewV1beta1SystemsItemSoftwareUpdatePostRequestBody instantiates a new V1beta1SystemsItemSoftwareUpdatePostRequestBody and sets the default values.
func NewV1beta1SystemsItemSoftwareUpdatePostRequestBody()(*V1beta1SystemsItemSoftwareUpdatePostRequestBody) {
    m := &V1beta1SystemsItemSoftwareUpdatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemSoftwareUpdatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSoftwareUpdatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSoftwareUpdatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCatalogVersion gets the catalogVersion property value. Version of the target software catalog for the software precheck/update request.
// returns a *string when successful
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) GetCatalogVersion()(*string) {
    return m.catalogVersion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["catalogVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogVersion(val)
        }
        return nil
    }
    res["hypervisorClusterIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetHypervisorClusterIds(res)
        }
        return nil
    }
    res["ignorePrecheckWarning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnorePrecheckWarning(val)
        }
        return nil
    }
    return res
}
// GetHypervisorClusterIds gets the hypervisorClusterIds property value. List of Unique Identifiers (usually UUIDs) of Hypervisor Clusters in the system for which software update operation is to be initiated. If omitted, software update operation will be initiated on all Hypervisor Clusters available in the system.
// returns a []string when successful
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) GetHypervisorClusterIds()([]string) {
    return m.hypervisorClusterIds
}
// GetIgnorePrecheckWarning gets the ignorePrecheckWarning property value. Boolean flag to ignore precheck warnings and proceed with software update operation.
// returns a *bool when successful
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) GetIgnorePrecheckWarning()(*bool) {
    return m.ignorePrecheckWarning
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("catalogVersion", m.GetCatalogVersion())
        if err != nil {
            return err
        }
    }
    if m.GetHypervisorClusterIds() != nil {
        err := writer.WriteCollectionOfStringValues("hypervisorClusterIds", m.GetHypervisorClusterIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignorePrecheckWarning", m.GetIgnorePrecheckWarning())
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
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCatalogVersion sets the catalogVersion property value. Version of the target software catalog for the software precheck/update request.
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) SetCatalogVersion(value *string)() {
    m.catalogVersion = value
}
// SetHypervisorClusterIds sets the hypervisorClusterIds property value. List of Unique Identifiers (usually UUIDs) of Hypervisor Clusters in the system for which software update operation is to be initiated. If omitted, software update operation will be initiated on all Hypervisor Clusters available in the system.
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) SetHypervisorClusterIds(value []string)() {
    m.hypervisorClusterIds = value
}
// SetIgnorePrecheckWarning sets the ignorePrecheckWarning property value. Boolean flag to ignore precheck warnings and proceed with software update operation.
func (m *V1beta1SystemsItemSoftwareUpdatePostRequestBody) SetIgnorePrecheckWarning(value *bool)() {
    m.ignorePrecheckWarning = value
}
type V1beta1SystemsItemSoftwareUpdatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCatalogVersion()(*string)
    GetHypervisorClusterIds()([]string)
    GetIgnorePrecheckWarning()(*bool)
    SetCatalogVersion(value *string)()
    SetHypervisorClusterIds(value []string)()
    SetIgnorePrecheckWarning(value *bool)()
}
