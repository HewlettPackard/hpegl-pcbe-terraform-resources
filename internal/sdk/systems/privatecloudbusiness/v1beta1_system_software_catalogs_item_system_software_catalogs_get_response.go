package privatecloudbusiness

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse system Software Catalog with all the detailed information of software components, end user license agreement and a list of systems that have an update path to the catalog.
type V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // End User License Agreement for this software catalog
    eula *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // Details of the hypervisor software.
    hypervisor V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorable
    // Details of the hypervisor manager software.
    hypervisorManager V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorManagerable
    // An identifier for the resource, usually a UUID.
    id *string
    // A system specified name for the resource.
    name *string
    // Release date of the software catalog
    releaseDate *string
    // The self reference for this resource.
    resourceUri *string
    // Details of the HPE Server Firmware.
    serverFirmware V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_serverFirmwareable
    // Details of the HPE Storage Connection Manager software.
    storageConnectionManager V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageConnectionManagerable
    // Details of the HPE Storage software.
    storageSoftware V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageSoftwareable
    // List of systems having update path to this software catalog
    systemsWithUpdatePath []V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePathable
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Catalog version
    version *string
}
// NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse instantiates a new V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse and sets the default values.
func NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse()(*V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) {
    m := &V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetEula gets the eula property value. End User License Agreement for this software catalog
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetEula()(*string) {
    return m.eula
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["eula"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEula(val)
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
    res["hypervisor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisor(val.(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorable))
        }
        return nil
    }
    res["hypervisorManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorManagerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManager(val.(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorManagerable))
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
    res["releaseDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDate(val)
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
    res["serverFirmware"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_serverFirmwareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerFirmware(val.(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_serverFirmwareable))
        }
        return nil
    }
    res["storageConnectionManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageConnectionManagerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageConnectionManager(val.(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageConnectionManagerable))
        }
        return nil
    }
    res["storageSoftware"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageSoftwareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageSoftware(val.(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageSoftwareable))
        }
        return nil
    }
    res["systemsWithUpdatePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePathFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePathable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePathable)
                }
            }
            m.SetSystemsWithUpdatePath(res)
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
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetHypervisor gets the hypervisor property value. Details of the hypervisor software.
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorable when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetHypervisor()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorable) {
    return m.hypervisor
}
// GetHypervisorManager gets the hypervisorManager property value. Details of the hypervisor manager software.
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorManagerable when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetHypervisorManager()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorManagerable) {
    return m.hypervisorManager
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetName()(*string) {
    return m.name
}
// GetReleaseDate gets the releaseDate property value. Release date of the software catalog
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetReleaseDate()(*string) {
    return m.releaseDate
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServerFirmware gets the serverFirmware property value. Details of the HPE Server Firmware.
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_serverFirmwareable when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetServerFirmware()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_serverFirmwareable) {
    return m.serverFirmware
}
// GetStorageConnectionManager gets the storageConnectionManager property value. Details of the HPE Storage Connection Manager software.
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageConnectionManagerable when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetStorageConnectionManager()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageConnectionManagerable) {
    return m.storageConnectionManager
}
// GetStorageSoftware gets the storageSoftware property value. Details of the HPE Storage software.
// returns a V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageSoftwareable when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetStorageSoftware()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageSoftwareable) {
    return m.storageSoftware
}
// GetSystemsWithUpdatePath gets the systemsWithUpdatePath property value. List of systems having update path to this software catalog
// returns a []V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePathable when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetSystemsWithUpdatePath()([]V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePathable) {
    return m.systemsWithUpdatePath
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetVersion gets the version property value. Catalog version
// returns a *string when successful
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eula", m.GetEula())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hypervisor", m.GetHypervisor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hypervisorManager", m.GetHypervisorManager())
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
        err := writer.WriteStringValue("releaseDate", m.GetReleaseDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("serverFirmware", m.GetServerFirmware())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storageConnectionManager", m.GetStorageConnectionManager())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storageSoftware", m.GetStorageSoftware())
        if err != nil {
            return err
        }
    }
    if m.GetSystemsWithUpdatePath() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSystemsWithUpdatePath()))
        for i, v := range m.GetSystemsWithUpdatePath() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("systemsWithUpdatePath", cast)
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
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetEula sets the eula property value. End User License Agreement for this software catalog
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetEula(value *string)() {
    m.eula = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHypervisor sets the hypervisor property value. Details of the hypervisor software.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetHypervisor(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorable)() {
    m.hypervisor = value
}
// SetHypervisorManager sets the hypervisorManager property value. Details of the hypervisor manager software.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetHypervisorManager(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorManagerable)() {
    m.hypervisorManager = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetName(value *string)() {
    m.name = value
}
// SetReleaseDate sets the releaseDate property value. Release date of the software catalog
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetReleaseDate(value *string)() {
    m.releaseDate = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServerFirmware sets the serverFirmware property value. Details of the HPE Server Firmware.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetServerFirmware(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_serverFirmwareable)() {
    m.serverFirmware = value
}
// SetStorageConnectionManager sets the storageConnectionManager property value. Details of the HPE Storage Connection Manager software.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetStorageConnectionManager(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageConnectionManagerable)() {
    m.storageConnectionManager = value
}
// SetStorageSoftware sets the storageSoftware property value. Details of the HPE Storage software.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetStorageSoftware(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageSoftwareable)() {
    m.storageSoftware = value
}
// SetSystemsWithUpdatePath sets the systemsWithUpdatePath property value. List of systems having update path to this software catalog
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetSystemsWithUpdatePath(value []V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePathable)() {
    m.systemsWithUpdatePath = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetVersion sets the version property value. Catalog version
func (m *V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse) SetVersion(value *string)() {
    m.version = value
}
type V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetEula()(*string)
    GetGeneration()(*int64)
    GetHypervisor()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorable)
    GetHypervisorManager()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorManagerable)
    GetId()(*string)
    GetName()(*string)
    GetReleaseDate()(*string)
    GetResourceUri()(*string)
    GetServerFirmware()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_serverFirmwareable)
    GetStorageConnectionManager()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageConnectionManagerable)
    GetStorageSoftware()(V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageSoftwareable)
    GetSystemsWithUpdatePath()([]V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePathable)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersion()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetEula(value *string)()
    SetGeneration(value *int64)()
    SetHypervisor(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorable)()
    SetHypervisorManager(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_hypervisorManagerable)()
    SetId(value *string)()
    SetName(value *string)()
    SetReleaseDate(value *string)()
    SetResourceUri(value *string)()
    SetServerFirmware(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_serverFirmwareable)()
    SetStorageConnectionManager(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageConnectionManagerable)()
    SetStorageSoftware(value V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_storageSoftwareable)()
    SetSystemsWithUpdatePath(value []V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse_systemsWithUpdatePathable)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersion(value *string)()
}
