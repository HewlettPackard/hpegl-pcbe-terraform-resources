package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Storage allocated to the RDS database instance
    allocatedStorageInGb *int32
    // Maximum storage allocated to the RDS database instance
    maxAllocatedStorageInGb *int32
    // Indicates if the RDS database instance is encrypted
    storageEncrypted *bool
    // Indicates type of the storage this RDS Database instance is running.
    storageType *string
}
// NewV1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo instantiates a new V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo and sets the default values.
func NewV1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo()(*V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) {
    m := &V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllocatedStorageInGb gets the allocatedStorageInGb property value. Storage allocated to the RDS database instance
// returns a *int32 when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) GetAllocatedStorageInGb()(*int32) {
    return m.allocatedStorageInGb
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allocatedStorageInGb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllocatedStorageInGb(val)
        }
        return nil
    }
    res["maxAllocatedStorageInGb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxAllocatedStorageInGb(val)
        }
        return nil
    }
    res["storageEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageEncrypted(val)
        }
        return nil
    }
    res["storageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageType(val)
        }
        return nil
    }
    return res
}
// GetMaxAllocatedStorageInGb gets the maxAllocatedStorageInGb property value. Maximum storage allocated to the RDS database instance
// returns a *int32 when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) GetMaxAllocatedStorageInGb()(*int32) {
    return m.maxAllocatedStorageInGb
}
// GetStorageEncrypted gets the storageEncrypted property value. Indicates if the RDS database instance is encrypted
// returns a *bool when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) GetStorageEncrypted()(*bool) {
    return m.storageEncrypted
}
// GetStorageType gets the storageType property value. Indicates type of the storage this RDS Database instance is running.
// returns a *string when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) GetStorageType()(*string) {
    return m.storageType
}
// Serialize serializes information the current object
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("allocatedStorageInGb", m.GetAllocatedStorageInGb())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxAllocatedStorageInGb", m.GetMaxAllocatedStorageInGb())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("storageEncrypted", m.GetStorageEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("storageType", m.GetStorageType())
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
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllocatedStorageInGb sets the allocatedStorageInGb property value. Storage allocated to the RDS database instance
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) SetAllocatedStorageInGb(value *int32)() {
    m.allocatedStorageInGb = value
}
// SetMaxAllocatedStorageInGb sets the maxAllocatedStorageInGb property value. Maximum storage allocated to the RDS database instance
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) SetMaxAllocatedStorageInGb(value *int32)() {
    m.maxAllocatedStorageInGb = value
}
// SetStorageEncrypted sets the storageEncrypted property value. Indicates if the RDS database instance is encrypted
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) SetStorageEncrypted(value *bool)() {
    m.storageEncrypted = value
}
// SetStorageType sets the storageType property value. Indicates type of the storage this RDS Database instance is running.
func (m *V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfo) SetStorageType(value *string)() {
    m.storageType = value
}
type V1beta1CspRdsInstancesGetResponse_items_cspInfo_storageInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllocatedStorageInGb()(*int32)
    GetMaxAllocatedStorageInGb()(*int32)
    GetStorageEncrypted()(*bool)
    GetStorageType()(*string)
    SetAllocatedStorageInGb(value *int32)()
    SetMaxAllocatedStorageInGb(value *int32)()
    SetStorageEncrypted(value *bool)()
    SetStorageType(value *string)()
}
