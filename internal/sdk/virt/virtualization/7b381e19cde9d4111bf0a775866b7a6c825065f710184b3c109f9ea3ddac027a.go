package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo the Amazon EBS settings for the instance type.
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The optimized EBS performance for the instance type.
    ebsOptimizedInfo V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo_ebsOptimizedInfoable
    // Indicates whether the instance type is Amazon EBS-optimized.
    ebsOptimizedSupport *string
    // Indicates whether Amazon EBS encryption is supported.
    encryptionSupport *string
    // Indicates whether non-volatile memory express (NVMe) is supported.
    nvmeSupport *string
}
// NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo instantiates a new V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo and sets the default values.
func NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo()(*V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) {
    m := &V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEbsOptimizedInfo gets the ebsOptimizedInfo property value. The optimized EBS performance for the instance type.
// returns a V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo_ebsOptimizedInfoable when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) GetEbsOptimizedInfo()(V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo_ebsOptimizedInfoable) {
    return m.ebsOptimizedInfo
}
// GetEbsOptimizedSupport gets the ebsOptimizedSupport property value. Indicates whether the instance type is Amazon EBS-optimized.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) GetEbsOptimizedSupport()(*string) {
    return m.ebsOptimizedSupport
}
// GetEncryptionSupport gets the encryptionSupport property value. Indicates whether Amazon EBS encryption is supported.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) GetEncryptionSupport()(*string) {
    return m.encryptionSupport
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ebsOptimizedInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo_ebsOptimizedInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEbsOptimizedInfo(val.(V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo_ebsOptimizedInfoable))
        }
        return nil
    }
    res["ebsOptimizedSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEbsOptimizedSupport(val)
        }
        return nil
    }
    res["encryptionSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionSupport(val)
        }
        return nil
    }
    res["nvmeSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNvmeSupport(val)
        }
        return nil
    }
    return res
}
// GetNvmeSupport gets the nvmeSupport property value. Indicates whether non-volatile memory express (NVMe) is supported.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) GetNvmeSupport()(*string) {
    return m.nvmeSupport
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("ebsOptimizedInfo", m.GetEbsOptimizedInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ebsOptimizedSupport", m.GetEbsOptimizedSupport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("encryptionSupport", m.GetEncryptionSupport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("nvmeSupport", m.GetNvmeSupport())
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
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEbsOptimizedInfo sets the ebsOptimizedInfo property value. The optimized EBS performance for the instance type.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) SetEbsOptimizedInfo(value V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo_ebsOptimizedInfoable)() {
    m.ebsOptimizedInfo = value
}
// SetEbsOptimizedSupport sets the ebsOptimizedSupport property value. Indicates whether the instance type is Amazon EBS-optimized.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) SetEbsOptimizedSupport(value *string)() {
    m.ebsOptimizedSupport = value
}
// SetEncryptionSupport sets the encryptionSupport property value. Indicates whether Amazon EBS encryption is supported.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) SetEncryptionSupport(value *string)() {
    m.encryptionSupport = value
}
// SetNvmeSupport sets the nvmeSupport property value. Indicates whether non-volatile memory express (NVMe) is supported.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo) SetNvmeSupport(value *string)() {
    m.nvmeSupport = value
}
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEbsOptimizedInfo()(V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo_ebsOptimizedInfoable)
    GetEbsOptimizedSupport()(*string)
    GetEncryptionSupport()(*string)
    GetNvmeSupport()(*string)
    SetEbsOptimizedInfo(value V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_ebsInfo_ebsOptimizedInfoable)()
    SetEbsOptimizedSupport(value *string)()
    SetEncryptionSupport(value *string)()
    SetNvmeSupport(value *string)()
}
