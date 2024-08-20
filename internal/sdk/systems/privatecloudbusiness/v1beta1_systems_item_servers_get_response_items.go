package privatecloudbusiness

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersGetResponse_items details of the Server with select information.
type V1beta1SystemsItemServersGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // Server Health information
    health V1beta1SystemsItemServersGetResponse_items_healthable
    // id, URI to reference the hypervisor host.
    hypervisorHost V1beta1SystemsItemServersGetResponse_items_hypervisorHostable
    // Secret information
    hypervisorHostRootUserPasswordSecret V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable
    // An identifier for the resource, usually a UUID.
    id *string
    // Secret information
    iloAdminUserPasswordSecret V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable
    // firmware version of iLO in the server.
    iloFirmwareVersion *string
    // ILO Network Information.
    iloNetworkInfo V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable
    // state of the iLO in the server.
    iloState *string
    // status of the iLO in the server.
    iloStatus *string
    // iLO LED indication.
    indicatorLedStatus *string
    // Memory of the server.
    memoryGib *string
    // The model of the server.
    model *string
    // A system specified name for the resource.
    name *string
    // Version of NCM installed on the server.
    ncmVersion *string
    // An identifier of the on prem agent for the server.
    onPremAgentId *string
    // Number of processors in the server.
    processorCount *string
    // Model of the processors in the server.
    processorModel *string
    // The self reference for this resource.
    resourceUri *string
    // Server's serial number.
    serialNumber *string
    // List of server network interfaces
    serverNetworkInterfaces []V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable
    // Unique Identifier of the system, usually a UUID.
    systemId *string
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1SystemsItemServersGetResponse_items instantiates a new V1beta1SystemsItemServersGetResponse_items and sets the default values.
func NewV1beta1SystemsItemServersGetResponse_items()(*V1beta1SystemsItemServersGetResponse_items) {
    m := &V1beta1SystemsItemServersGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["health"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_healthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealth(val.(V1beta1SystemsItemServersGetResponse_items_healthable))
        }
        return nil
    }
    res["hypervisorHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_hypervisorHostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorHost(val.(V1beta1SystemsItemServersGetResponse_items_hypervisorHostable))
        }
        return nil
    }
    res["hypervisorHostRootUserPasswordSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorHostRootUserPasswordSecret(val.(V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable))
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
    res["iloAdminUserPasswordSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloAdminUserPasswordSecret(val.(V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable))
        }
        return nil
    }
    res["iloFirmwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloFirmwareVersion(val)
        }
        return nil
    }
    res["iloNetworkInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_iloNetworkInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloNetworkInfo(val.(V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable))
        }
        return nil
    }
    res["iloState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloState(val)
        }
        return nil
    }
    res["iloStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloStatus(val)
        }
        return nil
    }
    res["indicatorLedStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndicatorLedStatus(val)
        }
        return nil
    }
    res["memoryGib"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryGib(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
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
    res["ncmVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNcmVersion(val)
        }
        return nil
    }
    res["onPremAgentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremAgentId(val)
        }
        return nil
    }
    res["processorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorCount(val)
        }
        return nil
    }
    res["processorModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorModel(val)
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
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["serverNetworkInterfaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable)
                }
            }
            m.SetServerNetworkInterfaces(res)
        }
        return nil
    }
    res["systemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemId(val)
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
func (m *V1beta1SystemsItemServersGetResponse_items) GetGeneration()(*int64) {
    return m.generation
}
// GetHealth gets the health property value. Server Health information
// returns a V1beta1SystemsItemServersGetResponse_items_healthable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetHealth()(V1beta1SystemsItemServersGetResponse_items_healthable) {
    return m.health
}
// GetHypervisorHost gets the hypervisorHost property value. id, URI to reference the hypervisor host.
// returns a V1beta1SystemsItemServersGetResponse_items_hypervisorHostable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetHypervisorHost()(V1beta1SystemsItemServersGetResponse_items_hypervisorHostable) {
    return m.hypervisorHost
}
// GetHypervisorHostRootUserPasswordSecret gets the hypervisorHostRootUserPasswordSecret property value. Secret information
// returns a V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetHypervisorHostRootUserPasswordSecret()(V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable) {
    return m.hypervisorHostRootUserPasswordSecret
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetId()(*string) {
    return m.id
}
// GetIloAdminUserPasswordSecret gets the iloAdminUserPasswordSecret property value. Secret information
// returns a V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetIloAdminUserPasswordSecret()(V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable) {
    return m.iloAdminUserPasswordSecret
}
// GetIloFirmwareVersion gets the iloFirmwareVersion property value. firmware version of iLO in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetIloFirmwareVersion()(*string) {
    return m.iloFirmwareVersion
}
// GetIloNetworkInfo gets the iloNetworkInfo property value. ILO Network Information.
// returns a V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetIloNetworkInfo()(V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable) {
    return m.iloNetworkInfo
}
// GetIloState gets the iloState property value. state of the iLO in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetIloState()(*string) {
    return m.iloState
}
// GetIloStatus gets the iloStatus property value. status of the iLO in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetIloStatus()(*string) {
    return m.iloStatus
}
// GetIndicatorLedStatus gets the indicatorLedStatus property value. iLO LED indication.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetIndicatorLedStatus()(*string) {
    return m.indicatorLedStatus
}
// GetMemoryGib gets the memoryGib property value. Memory of the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetMemoryGib()(*string) {
    return m.memoryGib
}
// GetModel gets the model property value. The model of the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetModel()(*string) {
    return m.model
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetName()(*string) {
    return m.name
}
// GetNcmVersion gets the ncmVersion property value. Version of NCM installed on the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetNcmVersion()(*string) {
    return m.ncmVersion
}
// GetOnPremAgentId gets the onPremAgentId property value. An identifier of the on prem agent for the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetOnPremAgentId()(*string) {
    return m.onPremAgentId
}
// GetProcessorCount gets the processorCount property value. Number of processors in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetProcessorCount()(*string) {
    return m.processorCount
}
// GetProcessorModel gets the processorModel property value. Model of the processors in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetProcessorModel()(*string) {
    return m.processorModel
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSerialNumber gets the serialNumber property value. Server's serial number.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetServerNetworkInterfaces gets the serverNetworkInterfaces property value. List of server network interfaces
// returns a []V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetServerNetworkInterfaces()([]V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable) {
    return m.serverNetworkInterfaces
}
// GetSystemId gets the systemId property value. Unique Identifier of the system, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetSystemId()(*string) {
    return m.systemId
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("health", m.GetHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hypervisorHost", m.GetHypervisorHost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hypervisorHostRootUserPasswordSecret", m.GetHypervisorHostRootUserPasswordSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("iloAdminUserPasswordSecret", m.GetIloAdminUserPasswordSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("iloFirmwareVersion", m.GetIloFirmwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("iloNetworkInfo", m.GetIloNetworkInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("iloState", m.GetIloState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("iloStatus", m.GetIloStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("indicatorLedStatus", m.GetIndicatorLedStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("memoryGib", m.GetMemoryGib())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
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
        err := writer.WriteStringValue("ncmVersion", m.GetNcmVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("processorCount", m.GetProcessorCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("processorModel", m.GetProcessorModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    if m.GetServerNetworkInterfaces() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServerNetworkInterfaces()))
        for i, v := range m.GetServerNetworkInterfaces() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("serverNetworkInterfaces", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("systemId", m.GetSystemId())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1SystemsItemServersGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1SystemsItemServersGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1SystemsItemServersGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1SystemsItemServersGetResponse_items) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHealth sets the health property value. Server Health information
func (m *V1beta1SystemsItemServersGetResponse_items) SetHealth(value V1beta1SystemsItemServersGetResponse_items_healthable)() {
    m.health = value
}
// SetHypervisorHost sets the hypervisorHost property value. id, URI to reference the hypervisor host.
func (m *V1beta1SystemsItemServersGetResponse_items) SetHypervisorHost(value V1beta1SystemsItemServersGetResponse_items_hypervisorHostable)() {
    m.hypervisorHost = value
}
// SetHypervisorHostRootUserPasswordSecret sets the hypervisorHostRootUserPasswordSecret property value. Secret information
func (m *V1beta1SystemsItemServersGetResponse_items) SetHypervisorHostRootUserPasswordSecret(value V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable)() {
    m.hypervisorHostRootUserPasswordSecret = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1SystemsItemServersGetResponse_items) SetId(value *string)() {
    m.id = value
}
// SetIloAdminUserPasswordSecret sets the iloAdminUserPasswordSecret property value. Secret information
func (m *V1beta1SystemsItemServersGetResponse_items) SetIloAdminUserPasswordSecret(value V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable)() {
    m.iloAdminUserPasswordSecret = value
}
// SetIloFirmwareVersion sets the iloFirmwareVersion property value. firmware version of iLO in the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetIloFirmwareVersion(value *string)() {
    m.iloFirmwareVersion = value
}
// SetIloNetworkInfo sets the iloNetworkInfo property value. ILO Network Information.
func (m *V1beta1SystemsItemServersGetResponse_items) SetIloNetworkInfo(value V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable)() {
    m.iloNetworkInfo = value
}
// SetIloState sets the iloState property value. state of the iLO in the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetIloState(value *string)() {
    m.iloState = value
}
// SetIloStatus sets the iloStatus property value. status of the iLO in the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetIloStatus(value *string)() {
    m.iloStatus = value
}
// SetIndicatorLedStatus sets the indicatorLedStatus property value. iLO LED indication.
func (m *V1beta1SystemsItemServersGetResponse_items) SetIndicatorLedStatus(value *string)() {
    m.indicatorLedStatus = value
}
// SetMemoryGib sets the memoryGib property value. Memory of the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetMemoryGib(value *string)() {
    m.memoryGib = value
}
// SetModel sets the model property value. The model of the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetModel(value *string)() {
    m.model = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1SystemsItemServersGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetNcmVersion sets the ncmVersion property value. Version of NCM installed on the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetNcmVersion(value *string)() {
    m.ncmVersion = value
}
// SetOnPremAgentId sets the onPremAgentId property value. An identifier of the on prem agent for the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetOnPremAgentId(value *string)() {
    m.onPremAgentId = value
}
// SetProcessorCount sets the processorCount property value. Number of processors in the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetProcessorCount(value *string)() {
    m.processorCount = value
}
// SetProcessorModel sets the processorModel property value. Model of the processors in the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetProcessorModel(value *string)() {
    m.processorModel = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1SystemsItemServersGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSerialNumber sets the serialNumber property value. Server's serial number.
func (m *V1beta1SystemsItemServersGetResponse_items) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetServerNetworkInterfaces sets the serverNetworkInterfaces property value. List of server network interfaces
func (m *V1beta1SystemsItemServersGetResponse_items) SetServerNetworkInterfaces(value []V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable)() {
    m.serverNetworkInterfaces = value
}
// SetSystemId sets the systemId property value. Unique Identifier of the system, usually a UUID.
func (m *V1beta1SystemsItemServersGetResponse_items) SetSystemId(value *string)() {
    m.systemId = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1SystemsItemServersGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1SystemsItemServersGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1SystemsItemServersGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetHealth()(V1beta1SystemsItemServersGetResponse_items_healthable)
    GetHypervisorHost()(V1beta1SystemsItemServersGetResponse_items_hypervisorHostable)
    GetHypervisorHostRootUserPasswordSecret()(V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable)
    GetId()(*string)
    GetIloAdminUserPasswordSecret()(V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable)
    GetIloFirmwareVersion()(*string)
    GetIloNetworkInfo()(V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable)
    GetIloState()(*string)
    GetIloStatus()(*string)
    GetIndicatorLedStatus()(*string)
    GetMemoryGib()(*string)
    GetModel()(*string)
    GetName()(*string)
    GetNcmVersion()(*string)
    GetOnPremAgentId()(*string)
    GetProcessorCount()(*string)
    GetProcessorModel()(*string)
    GetResourceUri()(*string)
    GetSerialNumber()(*string)
    GetServerNetworkInterfaces()([]V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable)
    GetSystemId()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetHealth(value V1beta1SystemsItemServersGetResponse_items_healthable)()
    SetHypervisorHost(value V1beta1SystemsItemServersGetResponse_items_hypervisorHostable)()
    SetHypervisorHostRootUserPasswordSecret(value V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable)()
    SetId(value *string)()
    SetIloAdminUserPasswordSecret(value V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable)()
    SetIloFirmwareVersion(value *string)()
    SetIloNetworkInfo(value V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable)()
    SetIloState(value *string)()
    SetIloStatus(value *string)()
    SetIndicatorLedStatus(value *string)()
    SetMemoryGib(value *string)()
    SetModel(value *string)()
    SetName(value *string)()
    SetNcmVersion(value *string)()
    SetOnPremAgentId(value *string)()
    SetProcessorCount(value *string)()
    SetProcessorModel(value *string)()
    SetResourceUri(value *string)()
    SetSerialNumber(value *string)()
    SetServerNetworkInterfaces(value []V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable)()
    SetSystemId(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
