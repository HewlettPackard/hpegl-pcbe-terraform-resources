package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemGetResponse_storageSystem storage Information of system.
type V1beta1SystemsItemGetResponse_storageSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Compression ratio of the storage system. Not available if storage system type is "Alletra MP File".
    compressionRatio *float64
    // Deduplication ratio of the storage system. Not available if storage system type is "Alletra MP File".
    dedupeRatio *float64
    // Name of the storage system group leader. Only available if storage system type is "Alletra".
    groupLeaderName *string
    // Serial number of the storage system group leader. Only available if storage system type is "Alletra".
    groupLeaderSerialNumber *string
    // Name of the storage system group. Only available if storage system type is "Alletra".
    groupName *string
    // State of the storage system group. Only available if storage system type is "Alletra".
    groupState *string
    // Storage system identifier, usually a UUID.
    id *string
    // Name of the storage system
    name *string
    // Number of cpu cores in the storage system. Only available if storage system type is "Alletra MP Block".
    numberOfCores *float64
    // Number of JBOFs in the storage system. Only available if storage system type is "Alletra MP Block".
    numberOfJbofs *float64
    // Number of nodes in the storage system. Only available if storage system type is "Alletra MP Block".
    numberOfNodes *float64
    // Resource URI of the storage system
    resourceUri *string
    // Serial number of the storage system.
    storageSerial *string
    // Storage system type
    storageType *string
}
// NewV1beta1SystemsItemGetResponse_storageSystem instantiates a new V1beta1SystemsItemGetResponse_storageSystem and sets the default values.
func NewV1beta1SystemsItemGetResponse_storageSystem()(*V1beta1SystemsItemGetResponse_storageSystem) {
    m := &V1beta1SystemsItemGetResponse_storageSystem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemGetResponse_storageSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemGetResponse_storageSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemGetResponse_storageSystem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompressionRatio gets the compressionRatio property value. Compression ratio of the storage system. Not available if storage system type is "Alletra MP File".
// returns a *float64 when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetCompressionRatio()(*float64) {
    return m.compressionRatio
}
// GetDedupeRatio gets the dedupeRatio property value. Deduplication ratio of the storage system. Not available if storage system type is "Alletra MP File".
// returns a *float64 when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetDedupeRatio()(*float64) {
    return m.dedupeRatio
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compressionRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompressionRatio(val)
        }
        return nil
    }
    res["dedupeRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDedupeRatio(val)
        }
        return nil
    }
    res["groupLeaderName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupLeaderName(val)
        }
        return nil
    }
    res["groupLeaderSerialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupLeaderSerialNumber(val)
        }
        return nil
    }
    res["groupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    res["groupState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupState(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["numberOfCores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCores(val)
        }
        return nil
    }
    res["numberOfJbofs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfJbofs(val)
        }
        return nil
    }
    res["numberOfNodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfNodes(val)
        }
        return nil
    }
    res["resourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceUri(val)
        }
        return nil
    }
    res["storageSerial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageSerial(val)
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
// GetGroupLeaderName gets the groupLeaderName property value. Name of the storage system group leader. Only available if storage system type is "Alletra".
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetGroupLeaderName()(*string) {
    return m.groupLeaderName
}
// GetGroupLeaderSerialNumber gets the groupLeaderSerialNumber property value. Serial number of the storage system group leader. Only available if storage system type is "Alletra".
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetGroupLeaderSerialNumber()(*string) {
    return m.groupLeaderSerialNumber
}
// GetGroupName gets the groupName property value. Name of the storage system group. Only available if storage system type is "Alletra".
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetGroupName()(*string) {
    return m.groupName
}
// GetGroupState gets the groupState property value. State of the storage system group. Only available if storage system type is "Alletra".
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetGroupState()(*string) {
    return m.groupState
}
// GetId gets the id property value. Storage system identifier, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the storage system
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetName()(*string) {
    return m.name
}
// GetNumberOfCores gets the numberOfCores property value. Number of cpu cores in the storage system. Only available if storage system type is "Alletra MP Block".
// returns a *float64 when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetNumberOfCores()(*float64) {
    return m.numberOfCores
}
// GetNumberOfJbofs gets the numberOfJbofs property value. Number of JBOFs in the storage system. Only available if storage system type is "Alletra MP Block".
// returns a *float64 when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetNumberOfJbofs()(*float64) {
    return m.numberOfJbofs
}
// GetNumberOfNodes gets the numberOfNodes property value. Number of nodes in the storage system. Only available if storage system type is "Alletra MP Block".
// returns a *float64 when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetNumberOfNodes()(*float64) {
    return m.numberOfNodes
}
// GetResourceUri gets the resourceUri property value. Resource URI of the storage system
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetStorageSerial gets the storageSerial property value. Serial number of the storage system.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetStorageSerial()(*string) {
    return m.storageSerial
}
// GetStorageType gets the storageType property value. Storage system type
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse_storageSystem) GetStorageType()(*string) {
    return m.storageType
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemGetResponse_storageSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("compressionRatio", m.GetCompressionRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("dedupeRatio", m.GetDedupeRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupLeaderName", m.GetGroupLeaderName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupLeaderSerialNumber", m.GetGroupLeaderSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupState", m.GetGroupState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("numberOfCores", m.GetNumberOfCores())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("numberOfJbofs", m.GetNumberOfJbofs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("numberOfNodes", m.GetNumberOfNodes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceUri", m.GetResourceUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("storageSerial", m.GetStorageSerial())
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
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompressionRatio sets the compressionRatio property value. Compression ratio of the storage system. Not available if storage system type is "Alletra MP File".
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetCompressionRatio(value *float64)() {
    m.compressionRatio = value
}
// SetDedupeRatio sets the dedupeRatio property value. Deduplication ratio of the storage system. Not available if storage system type is "Alletra MP File".
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetDedupeRatio(value *float64)() {
    m.dedupeRatio = value
}
// SetGroupLeaderName sets the groupLeaderName property value. Name of the storage system group leader. Only available if storage system type is "Alletra".
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetGroupLeaderName(value *string)() {
    m.groupLeaderName = value
}
// SetGroupLeaderSerialNumber sets the groupLeaderSerialNumber property value. Serial number of the storage system group leader. Only available if storage system type is "Alletra".
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetGroupLeaderSerialNumber(value *string)() {
    m.groupLeaderSerialNumber = value
}
// SetGroupName sets the groupName property value. Name of the storage system group. Only available if storage system type is "Alletra".
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetGroupName(value *string)() {
    m.groupName = value
}
// SetGroupState sets the groupState property value. State of the storage system group. Only available if storage system type is "Alletra".
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetGroupState(value *string)() {
    m.groupState = value
}
// SetId sets the id property value. Storage system identifier, usually a UUID.
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the storage system
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetName(value *string)() {
    m.name = value
}
// SetNumberOfCores sets the numberOfCores property value. Number of cpu cores in the storage system. Only available if storage system type is "Alletra MP Block".
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetNumberOfCores(value *float64)() {
    m.numberOfCores = value
}
// SetNumberOfJbofs sets the numberOfJbofs property value. Number of JBOFs in the storage system. Only available if storage system type is "Alletra MP Block".
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetNumberOfJbofs(value *float64)() {
    m.numberOfJbofs = value
}
// SetNumberOfNodes sets the numberOfNodes property value. Number of nodes in the storage system. Only available if storage system type is "Alletra MP Block".
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetNumberOfNodes(value *float64)() {
    m.numberOfNodes = value
}
// SetResourceUri sets the resourceUri property value. Resource URI of the storage system
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetStorageSerial sets the storageSerial property value. Serial number of the storage system.
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetStorageSerial(value *string)() {
    m.storageSerial = value
}
// SetStorageType sets the storageType property value. Storage system type
func (m *V1beta1SystemsItemGetResponse_storageSystem) SetStorageType(value *string)() {
    m.storageType = value
}
type V1beta1SystemsItemGetResponse_storageSystemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompressionRatio()(*float64)
    GetDedupeRatio()(*float64)
    GetGroupLeaderName()(*string)
    GetGroupLeaderSerialNumber()(*string)
    GetGroupName()(*string)
    GetGroupState()(*string)
    GetId()(*string)
    GetName()(*string)
    GetNumberOfCores()(*float64)
    GetNumberOfJbofs()(*float64)
    GetNumberOfNodes()(*float64)
    GetResourceUri()(*string)
    GetStorageSerial()(*string)
    GetStorageType()(*string)
    SetCompressionRatio(value *float64)()
    SetDedupeRatio(value *float64)()
    SetGroupLeaderName(value *string)()
    SetGroupLeaderSerialNumber(value *string)()
    SetGroupName(value *string)()
    SetGroupState(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetNumberOfCores(value *float64)()
    SetNumberOfJbofs(value *float64)()
    SetNumberOfNodes(value *float64)()
    SetResourceUri(value *string)()
    SetStorageSerial(value *string)()
    SetStorageType(value *string)()
}
