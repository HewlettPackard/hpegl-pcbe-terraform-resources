package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog current Version of the software catalog and its constituent software components. Applicable only for Private Cloud AI systems.
type V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Details of the AI Essentials software.
    aiEssentials V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_aiEssentialsable
    // Details of the Data Services Connector software.
    dataServicesConnector V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorable
    // Details of the hypervisor software.
    hypervisor V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorable
    // Details of the hypervisor manager software.
    hypervisorManager V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorManagerable
    // Reasons if any as to why the current software catalog version cannot be determined
    reasons []string
    // Release date of the software catalog
    releaseDate *string
    // Details of the HPE Server Firmware.
    serverFirmware V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_serverFirmwareable
    // Details of the HPE Storage software.
    storageSoftware V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_storageSoftwareable
    // Catalog version
    version *string
}
// NewV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog instantiates a new V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog and sets the default values.
func NewV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog()(*V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) {
    m := &V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAiEssentials gets the aiEssentials property value. Details of the AI Essentials software.
// returns a V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_aiEssentialsable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetAiEssentials()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_aiEssentialsable) {
    return m.aiEssentials
}
// GetDataServicesConnector gets the dataServicesConnector property value. Details of the Data Services Connector software.
// returns a V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetDataServicesConnector()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorable) {
    return m.dataServicesConnector
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aiEssentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_aiEssentialsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAiEssentials(val.(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_aiEssentialsable))
        }
        return nil
    }
    res["dataServicesConnector"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataServicesConnector(val.(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorable))
        }
        return nil
    }
    res["hypervisor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisor(val.(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorable))
        }
        return nil
    }
    res["hypervisorManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorManagerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManager(val.(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorManagerable))
        }
        return nil
    }
    res["reasons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetReasons(res)
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
    res["serverFirmware"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_serverFirmwareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerFirmware(val.(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_serverFirmwareable))
        }
        return nil
    }
    res["storageSoftware"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_storageSoftwareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageSoftware(val.(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_storageSoftwareable))
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
// GetHypervisor gets the hypervisor property value. Details of the hypervisor software.
// returns a V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetHypervisor()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorable) {
    return m.hypervisor
}
// GetHypervisorManager gets the hypervisorManager property value. Details of the hypervisor manager software.
// returns a V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorManagerable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetHypervisorManager()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorManagerable) {
    return m.hypervisorManager
}
// GetReasons gets the reasons property value. Reasons if any as to why the current software catalog version cannot be determined
// returns a []string when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetReasons()([]string) {
    return m.reasons
}
// GetReleaseDate gets the releaseDate property value. Release date of the software catalog
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetReleaseDate()(*string) {
    return m.releaseDate
}
// GetServerFirmware gets the serverFirmware property value. Details of the HPE Server Firmware.
// returns a V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_serverFirmwareable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetServerFirmware()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_serverFirmwareable) {
    return m.serverFirmware
}
// GetStorageSoftware gets the storageSoftware property value. Details of the HPE Storage software.
// returns a V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_storageSoftwareable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetStorageSoftware()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_storageSoftwareable) {
    return m.storageSoftware
}
// GetVersion gets the version property value. Catalog version
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("aiEssentials", m.GetAiEssentials())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dataServicesConnector", m.GetDataServicesConnector())
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
    if m.GetReasons() != nil {
        err := writer.WriteCollectionOfStringValues("reasons", m.GetReasons())
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
        err := writer.WriteObjectValue("storageSoftware", m.GetStorageSoftware())
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
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAiEssentials sets the aiEssentials property value. Details of the AI Essentials software.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetAiEssentials(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_aiEssentialsable)() {
    m.aiEssentials = value
}
// SetDataServicesConnector sets the dataServicesConnector property value. Details of the Data Services Connector software.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetDataServicesConnector(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorable)() {
    m.dataServicesConnector = value
}
// SetHypervisor sets the hypervisor property value. Details of the hypervisor software.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetHypervisor(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorable)() {
    m.hypervisor = value
}
// SetHypervisorManager sets the hypervisorManager property value. Details of the hypervisor manager software.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetHypervisorManager(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorManagerable)() {
    m.hypervisorManager = value
}
// SetReasons sets the reasons property value. Reasons if any as to why the current software catalog version cannot be determined
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetReasons(value []string)() {
    m.reasons = value
}
// SetReleaseDate sets the releaseDate property value. Release date of the software catalog
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetReleaseDate(value *string)() {
    m.releaseDate = value
}
// SetServerFirmware sets the serverFirmware property value. Details of the HPE Server Firmware.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetServerFirmware(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_serverFirmwareable)() {
    m.serverFirmware = value
}
// SetStorageSoftware sets the storageSoftware property value. Details of the HPE Storage software.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetStorageSoftware(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_storageSoftwareable)() {
    m.storageSoftware = value
}
// SetVersion sets the version property value. Catalog version
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog) SetVersion(value *string)() {
    m.version = value
}
type V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAiEssentials()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_aiEssentialsable)
    GetDataServicesConnector()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorable)
    GetHypervisor()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorable)
    GetHypervisorManager()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorManagerable)
    GetReasons()([]string)
    GetReleaseDate()(*string)
    GetServerFirmware()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_serverFirmwareable)
    GetStorageSoftware()(V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_storageSoftwareable)
    GetVersion()(*string)
    SetAiEssentials(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_aiEssentialsable)()
    SetDataServicesConnector(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorable)()
    SetHypervisor(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorable)()
    SetHypervisorManager(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_hypervisorManagerable)()
    SetReasons(value []string)()
    SetReleaseDate(value *string)()
    SetServerFirmware(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_serverFirmwareable)()
    SetStorageSoftware(value V1beta1SystemsItemPatchResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_storageSoftwareable)()
    SetVersion(value *string)()
}
