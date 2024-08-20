package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo rDS machine instance properties whose values are defined by, and synchronized with, a cloud service provider.
type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Cloud provider availability zone
    availabilityZone *string
    // RDS Database instance Class
    class *string
    // The reference Id for the asset with cloud service provider
    cspReference *string
    // Cloud provider region
    cspRegion *string
    // The cspTags property
    cspTags []V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_cspTagsable
    // RDS Engine name
    engine *string
    // RDS Database engine version
    engineVersion *string
    // RDS database instance identifier
    identifier *string
    // Optional keyId if the machine instance is encrypted
    kmsKeyId *string
    // Indicated if the RDS database instance is deployed in Multi-AZ deployment
    multiAz *bool
    // Network properties for the machine instance
    networkInfo V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoable
    // The optionGroupNames property
    optionGroupNames []string
    // The parameterGroupNames property
    parameterGroupNames []string
    // If multi-az deployment enabled for this RDS instance, it indicates the availability zone for the secondary RDS instance
    secondaryAvailabilityZone *string
    // RDS machine instance status in cloud service provider
    status *string
    // Storage information of the RDS Database instance
    storageInfo V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_storageInfoable
}
// NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo instantiates a new V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo and sets the default values.
func NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo()(*V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) {
    m := &V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailabilityZone gets the availabilityZone property value. Cloud provider availability zone
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetAvailabilityZone()(*string) {
    return m.availabilityZone
}
// GetClass gets the class property value. RDS Database instance Class
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetClass()(*string) {
    return m.class
}
// GetCspReference gets the cspReference property value. The reference Id for the asset with cloud service provider
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetCspReference()(*string) {
    return m.cspReference
}
// GetCspRegion gets the cspRegion property value. Cloud provider region
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetCspRegion()(*string) {
    return m.cspRegion
}
// GetCspTags gets the cspTags property value. The cspTags property
// returns a []V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_cspTagsable when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetCspTags()([]V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_cspTagsable) {
    return m.cspTags
}
// GetEngine gets the engine property value. RDS Engine name
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetEngine()(*string) {
    return m.engine
}
// GetEngineVersion gets the engineVersion property value. RDS Database engine version
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetEngineVersion()(*string) {
    return m.engineVersion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availabilityZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityZone(val)
        }
        return nil
    }
    res["class"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClass(val)
        }
        return nil
    }
    res["cspReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspReference(val)
        }
        return nil
    }
    res["cspRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspRegion(val)
        }
        return nil
    }
    res["cspTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_cspTagsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_cspTagsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_cspTagsable)
                }
            }
            m.SetCspTags(res)
        }
        return nil
    }
    res["engine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEngine(val)
        }
        return nil
    }
    res["engineVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEngineVersion(val)
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["kmsKeyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKmsKeyId(val)
        }
        return nil
    }
    res["multiAz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultiAz(val)
        }
        return nil
    }
    res["networkInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkInfo(val.(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoable))
        }
        return nil
    }
    res["optionGroupNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetOptionGroupNames(res)
        }
        return nil
    }
    res["parameterGroupNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetParameterGroupNames(res)
        }
        return nil
    }
    res["secondaryAvailabilityZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryAvailabilityZone(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["storageInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_storageInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageInfo(val.(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_storageInfoable))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. RDS database instance identifier
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetIdentifier()(*string) {
    return m.identifier
}
// GetKmsKeyId gets the kmsKeyId property value. Optional keyId if the machine instance is encrypted
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetKmsKeyId()(*string) {
    return m.kmsKeyId
}
// GetMultiAz gets the multiAz property value. Indicated if the RDS database instance is deployed in Multi-AZ deployment
// returns a *bool when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetMultiAz()(*bool) {
    return m.multiAz
}
// GetNetworkInfo gets the networkInfo property value. Network properties for the machine instance
// returns a V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoable when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetNetworkInfo()(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoable) {
    return m.networkInfo
}
// GetOptionGroupNames gets the optionGroupNames property value. The optionGroupNames property
// returns a []string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetOptionGroupNames()([]string) {
    return m.optionGroupNames
}
// GetParameterGroupNames gets the parameterGroupNames property value. The parameterGroupNames property
// returns a []string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetParameterGroupNames()([]string) {
    return m.parameterGroupNames
}
// GetSecondaryAvailabilityZone gets the secondaryAvailabilityZone property value. If multi-az deployment enabled for this RDS instance, it indicates the availability zone for the secondary RDS instance
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetSecondaryAvailabilityZone()(*string) {
    return m.secondaryAvailabilityZone
}
// GetStatus gets the status property value. RDS machine instance status in cloud service provider
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetStatus()(*string) {
    return m.status
}
// GetStorageInfo gets the storageInfo property value. Storage information of the RDS Database instance
// returns a V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_storageInfoable when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) GetStorageInfo()(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_storageInfoable) {
    return m.storageInfo
}
// Serialize serializes information the current object
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("class", m.GetClass())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cspReference", m.GetCspReference())
        if err != nil {
            return err
        }
    }
    if m.GetCspTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCspTags()))
        for i, v := range m.GetCspTags() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("cspTags", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("engineVersion", m.GetEngineVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("kmsKeyId", m.GetKmsKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("multiAz", m.GetMultiAz())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("networkInfo", m.GetNetworkInfo())
        if err != nil {
            return err
        }
    }
    if m.GetOptionGroupNames() != nil {
        err := writer.WriteCollectionOfStringValues("optionGroupNames", m.GetOptionGroupNames())
        if err != nil {
            return err
        }
    }
    if m.GetParameterGroupNames() != nil {
        err := writer.WriteCollectionOfStringValues("parameterGroupNames", m.GetParameterGroupNames())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secondaryAvailabilityZone", m.GetSecondaryAvailabilityZone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storageInfo", m.GetStorageInfo())
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
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailabilityZone sets the availabilityZone property value. Cloud provider availability zone
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetAvailabilityZone(value *string)() {
    m.availabilityZone = value
}
// SetClass sets the class property value. RDS Database instance Class
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetClass(value *string)() {
    m.class = value
}
// SetCspReference sets the cspReference property value. The reference Id for the asset with cloud service provider
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetCspReference(value *string)() {
    m.cspReference = value
}
// SetCspRegion sets the cspRegion property value. Cloud provider region
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetCspRegion(value *string)() {
    m.cspRegion = value
}
// SetCspTags sets the cspTags property value. The cspTags property
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetCspTags(value []V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_cspTagsable)() {
    m.cspTags = value
}
// SetEngine sets the engine property value. RDS Engine name
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetEngine(value *string)() {
    m.engine = value
}
// SetEngineVersion sets the engineVersion property value. RDS Database engine version
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetEngineVersion(value *string)() {
    m.engineVersion = value
}
// SetIdentifier sets the identifier property value. RDS database instance identifier
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetIdentifier(value *string)() {
    m.identifier = value
}
// SetKmsKeyId sets the kmsKeyId property value. Optional keyId if the machine instance is encrypted
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetKmsKeyId(value *string)() {
    m.kmsKeyId = value
}
// SetMultiAz sets the multiAz property value. Indicated if the RDS database instance is deployed in Multi-AZ deployment
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetMultiAz(value *bool)() {
    m.multiAz = value
}
// SetNetworkInfo sets the networkInfo property value. Network properties for the machine instance
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetNetworkInfo(value V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoable)() {
    m.networkInfo = value
}
// SetOptionGroupNames sets the optionGroupNames property value. The optionGroupNames property
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetOptionGroupNames(value []string)() {
    m.optionGroupNames = value
}
// SetParameterGroupNames sets the parameterGroupNames property value. The parameterGroupNames property
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetParameterGroupNames(value []string)() {
    m.parameterGroupNames = value
}
// SetSecondaryAvailabilityZone sets the secondaryAvailabilityZone property value. If multi-az deployment enabled for this RDS instance, it indicates the availability zone for the secondary RDS instance
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetSecondaryAvailabilityZone(value *string)() {
    m.secondaryAvailabilityZone = value
}
// SetStatus sets the status property value. RDS machine instance status in cloud service provider
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetStatus(value *string)() {
    m.status = value
}
// SetStorageInfo sets the storageInfo property value. Storage information of the RDS Database instance
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo) SetStorageInfo(value V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_storageInfoable)() {
    m.storageInfo = value
}
type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityZone()(*string)
    GetClass()(*string)
    GetCspReference()(*string)
    GetCspRegion()(*string)
    GetCspTags()([]V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_cspTagsable)
    GetEngine()(*string)
    GetEngineVersion()(*string)
    GetIdentifier()(*string)
    GetKmsKeyId()(*string)
    GetMultiAz()(*bool)
    GetNetworkInfo()(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoable)
    GetOptionGroupNames()([]string)
    GetParameterGroupNames()([]string)
    GetSecondaryAvailabilityZone()(*string)
    GetStatus()(*string)
    GetStorageInfo()(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_storageInfoable)
    SetAvailabilityZone(value *string)()
    SetClass(value *string)()
    SetCspReference(value *string)()
    SetCspRegion(value *string)()
    SetCspTags(value []V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_cspTagsable)()
    SetEngine(value *string)()
    SetEngineVersion(value *string)()
    SetIdentifier(value *string)()
    SetKmsKeyId(value *string)()
    SetMultiAz(value *bool)()
    SetNetworkInfo(value V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoable)()
    SetOptionGroupNames(value []string)()
    SetParameterGroupNames(value []string)()
    SetSecondaryAvailabilityZone(value *string)()
    SetStatus(value *string)()
    SetStorageInfo(value V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_storageInfoable)()
}
