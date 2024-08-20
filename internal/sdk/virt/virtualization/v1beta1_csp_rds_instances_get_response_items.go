package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspRdsInstancesGetResponse_items details of a CSP RDS machine instance
type V1beta1CspRdsInstancesGetResponse_items struct {
    // Information about the cloud provider account where the resource is located.
    accountInfo V1beta1CspRdsInstancesGetResponse_items_accountInfoable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The backupInfo property
    backupInfo []V1beta1CspRdsInstancesGetResponse_items_backupInfoable
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // RDS machine instance properties whose values are defined by, and synchronized with, a cloud service provider.
    cspInfo V1beta1CspRdsInstancesGetResponse_items_cspInfoable
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // An identifier for the resource, usually a UUID.
    id *string
    // Metadata information related to possible backups supported for given instance
    metadata V1beta1CspRdsInstancesGetResponse_items_metadataable
    // A system specified name for the resource.
    name *string
    // Protection jobs associated directly with this RDS Instance.
    protectionJobInfo []V1beta1CspRdsInstancesGetResponse_items_protectionJobInfoable
    // Property definitions for Refresh inventory Assets
    refreshInfo V1beta1CspRdsInstancesGetResponse_items_refreshInfoable
    // The self reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1CspRdsInstancesGetResponse_items instantiates a new V1beta1CspRdsInstancesGetResponse_items and sets the default values.
func NewV1beta1CspRdsInstancesGetResponse_items()(*V1beta1CspRdsInstancesGetResponse_items) {
    m := &V1beta1CspRdsInstancesGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspRdsInstancesGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspRdsInstancesGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspRdsInstancesGetResponse_items(), nil
}
// GetAccountInfo gets the accountInfo property value. Information about the cloud provider account where the resource is located.
// returns a V1beta1CspRdsInstancesGetResponse_items_accountInfoable when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetAccountInfo()(V1beta1CspRdsInstancesGetResponse_items_accountInfoable) {
    return m.accountInfo
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBackupInfo gets the backupInfo property value. The backupInfo property
// returns a []V1beta1CspRdsInstancesGetResponse_items_backupInfoable when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetBackupInfo()([]V1beta1CspRdsInstancesGetResponse_items_backupInfoable) {
    return m.backupInfo
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCspInfo gets the cspInfo property value. RDS machine instance properties whose values are defined by, and synchronized with, a cloud service provider.
// returns a V1beta1CspRdsInstancesGetResponse_items_cspInfoable when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetCspInfo()(V1beta1CspRdsInstancesGetResponse_items_cspInfoable) {
    return m.cspInfo
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspRdsInstancesGetResponse_items_accountInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountInfo(val.(V1beta1CspRdsInstancesGetResponse_items_accountInfoable))
        }
        return nil
    }
    res["backupInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspRdsInstancesGetResponse_items_backupInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspRdsInstancesGetResponse_items_backupInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspRdsInstancesGetResponse_items_backupInfoable)
                }
            }
            m.SetBackupInfo(res)
        }
        return nil
    }
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["cspInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspRdsInstancesGetResponse_items_cspInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspInfo(val.(V1beta1CspRdsInstancesGetResponse_items_cspInfoable))
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["generation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
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
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspRdsInstancesGetResponse_items_metadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val.(V1beta1CspRdsInstancesGetResponse_items_metadataable))
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
    res["protectionJobInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspRdsInstancesGetResponse_items_protectionJobInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspRdsInstancesGetResponse_items_protectionJobInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspRdsInstancesGetResponse_items_protectionJobInfoable)
                }
            }
            m.SetProtectionJobInfo(res)
        }
        return nil
    }
    res["refreshInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspRdsInstancesGetResponse_items_refreshInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRefreshInfo(val.(V1beta1CspRdsInstancesGetResponse_items_refreshInfoable))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["updatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetGeneration()(*int64) {
    return m.generation
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetId()(*string) {
    return m.id
}
// GetMetadata gets the metadata property value. Metadata information related to possible backups supported for given instance
// returns a V1beta1CspRdsInstancesGetResponse_items_metadataable when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetMetadata()(V1beta1CspRdsInstancesGetResponse_items_metadataable) {
    return m.metadata
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetName()(*string) {
    return m.name
}
// GetProtectionJobInfo gets the protectionJobInfo property value. Protection jobs associated directly with this RDS Instance.
// returns a []V1beta1CspRdsInstancesGetResponse_items_protectionJobInfoable when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetProtectionJobInfo()([]V1beta1CspRdsInstancesGetResponse_items_protectionJobInfoable) {
    return m.protectionJobInfo
}
// GetRefreshInfo gets the refreshInfo property value. Property definitions for Refresh inventory Assets
// returns a V1beta1CspRdsInstancesGetResponse_items_refreshInfoable when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetRefreshInfo()(V1beta1CspRdsInstancesGetResponse_items_refreshInfoable) {
    return m.refreshInfo
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1CspRdsInstancesGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1CspRdsInstancesGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBackupInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBackupInfo()))
        for i, v := range m.GetBackupInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("backupInfo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cspInfo", m.GetCspInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("metadata", m.GetMetadata())
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
    if m.GetProtectionJobInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProtectionJobInfo()))
        for i, v := range m.GetProtectionJobInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("protectionJobInfo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("refreshInfo", m.GetRefreshInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedAt", m.GetUpdatedAt())
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
// SetAccountInfo sets the accountInfo property value. Information about the cloud provider account where the resource is located.
func (m *V1beta1CspRdsInstancesGetResponse_items) SetAccountInfo(value V1beta1CspRdsInstancesGetResponse_items_accountInfoable)() {
    m.accountInfo = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspRdsInstancesGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBackupInfo sets the backupInfo property value. The backupInfo property
func (m *V1beta1CspRdsInstancesGetResponse_items) SetBackupInfo(value []V1beta1CspRdsInstancesGetResponse_items_backupInfoable)() {
    m.backupInfo = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1CspRdsInstancesGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCspInfo sets the cspInfo property value. RDS machine instance properties whose values are defined by, and synchronized with, a cloud service provider.
func (m *V1beta1CspRdsInstancesGetResponse_items) SetCspInfo(value V1beta1CspRdsInstancesGetResponse_items_cspInfoable)() {
    m.cspInfo = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1CspRdsInstancesGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1CspRdsInstancesGetResponse_items) SetGeneration(value *int64)() {
    m.generation = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1CspRdsInstancesGetResponse_items) SetId(value *string)() {
    m.id = value
}
// SetMetadata sets the metadata property value. Metadata information related to possible backups supported for given instance
func (m *V1beta1CspRdsInstancesGetResponse_items) SetMetadata(value V1beta1CspRdsInstancesGetResponse_items_metadataable)() {
    m.metadata = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1CspRdsInstancesGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetProtectionJobInfo sets the protectionJobInfo property value. Protection jobs associated directly with this RDS Instance.
func (m *V1beta1CspRdsInstancesGetResponse_items) SetProtectionJobInfo(value []V1beta1CspRdsInstancesGetResponse_items_protectionJobInfoable)() {
    m.protectionJobInfo = value
}
// SetRefreshInfo sets the refreshInfo property value. Property definitions for Refresh inventory Assets
func (m *V1beta1CspRdsInstancesGetResponse_items) SetRefreshInfo(value V1beta1CspRdsInstancesGetResponse_items_refreshInfoable)() {
    m.refreshInfo = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1CspRdsInstancesGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1CspRdsInstancesGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1CspRdsInstancesGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1CspRdsInstancesGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountInfo()(V1beta1CspRdsInstancesGetResponse_items_accountInfoable)
    GetBackupInfo()([]V1beta1CspRdsInstancesGetResponse_items_backupInfoable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCspInfo()(V1beta1CspRdsInstancesGetResponse_items_cspInfoable)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetId()(*string)
    GetMetadata()(V1beta1CspRdsInstancesGetResponse_items_metadataable)
    GetName()(*string)
    GetProtectionJobInfo()([]V1beta1CspRdsInstancesGetResponse_items_protectionJobInfoable)
    GetRefreshInfo()(V1beta1CspRdsInstancesGetResponse_items_refreshInfoable)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccountInfo(value V1beta1CspRdsInstancesGetResponse_items_accountInfoable)()
    SetBackupInfo(value []V1beta1CspRdsInstancesGetResponse_items_backupInfoable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCspInfo(value V1beta1CspRdsInstancesGetResponse_items_cspInfoable)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetId(value *string)()
    SetMetadata(value V1beta1CspRdsInstancesGetResponse_items_metadataable)()
    SetName(value *string)()
    SetProtectionJobInfo(value []V1beta1CspRdsInstancesGetResponse_items_protectionJobInfoable)()
    SetRefreshInfo(value V1beta1CspRdsInstancesGetResponse_items_refreshInfoable)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
