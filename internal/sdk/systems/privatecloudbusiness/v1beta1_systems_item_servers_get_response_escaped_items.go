package privatecloudbusiness

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersGetResponse_items details of the Server with select information.
type V1beta1SystemsItemServersGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // HPE Cloud module configuration.
    cloudModuleConfig V1beta1SystemsItemServersGetResponse_items_cloudModuleConfigable
    // Server Compute Usage Information.
    computeUsage V1beta1SystemsItemServersGetResponse_items_computeUsageable
    // The consumedBy property
    consumedBy V1beta1SystemsItemServersGetResponse_items_consumedByable
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // Gpus in the server.
    gpus []V1beta1SystemsItemServersGetResponse_items_gpusable
    // id, URI to reference the hypervisor host. Soon to be deprecated.
    hypervisorHost V1beta1SystemsItemServersGetResponse_items_hypervisorHostable
    // Secret information
    hypervisorHostRootUserPasswordSecret V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable
    // An identifier for the resource, usually a UUID.
    id *string
    // Secret information
    iloAdminUserPasswordSecret V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable
    // firmware version of iLO in the server.
    iloFirmwareVersion *string
    // ILO network information.
    iloNetworkInfo V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable
    // state of the iLO in the server.
    iloState *string
    // status of the iLO in the server.
    iloStatus *string
    // iLO user name and authentication credential secret. Property iloAdminUserPasswordSecret will be deprecated.Use this property instead.
    iloUserCredential V1beta1SystemsItemServersGetResponse_items_iloUserCredentialable
    // iLO LED indication.
    indicatorLedStatus *string
    // Information about link local IP addresses.
    linkLocalInfo V1beta1SystemsItemServersGetResponse_items_linkLocalInfoable
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
    // Information about the OS on the server.
    operatingSystemInfo V1beta1SystemsItemServersGetResponse_items_operatingSystemInfoable
    // Operating system user name and authentication credential secret. Property hypervisorHostRootUserPasswordSecret will be deprecated.Use this property instead.
    osUserCredential V1beta1SystemsItemServersGetResponse_items_osUserCredentialable
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
    // Unique identifier of the system, usually a UUID.
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
// GetCloudModuleConfig gets the cloudModuleConfig property value. HPE Cloud module configuration.
// returns a V1beta1SystemsItemServersGetResponse_items_cloudModuleConfigable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetCloudModuleConfig()(V1beta1SystemsItemServersGetResponse_items_cloudModuleConfigable) {
    return m.cloudModuleConfig
}
// GetComputeUsage gets the computeUsage property value. Server Compute Usage Information.
// returns a V1beta1SystemsItemServersGetResponse_items_computeUsageable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetComputeUsage()(V1beta1SystemsItemServersGetResponse_items_computeUsageable) {
    return m.computeUsage
}
// GetConsumedBy gets the consumedBy property value. The consumedBy property
// returns a V1beta1SystemsItemServersGetResponse_items_consumedByable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetConsumedBy()(V1beta1SystemsItemServersGetResponse_items_consumedByable) {
    return m.consumedBy
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
    res["cloudModuleConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_cloudModuleConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudModuleConfig(val.(V1beta1SystemsItemServersGetResponse_items_cloudModuleConfigable))
        }
        return nil
    }
    res["computeUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_computeUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputeUsage(val.(V1beta1SystemsItemServersGetResponse_items_computeUsageable))
        }
        return nil
    }
    res["consumedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_consumedByFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsumedBy(val.(V1beta1SystemsItemServersGetResponse_items_consumedByable))
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
    res["gpus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemServersGetResponse_items_gpusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemServersGetResponse_items_gpusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemServersGetResponse_items_gpusable)
                }
            }
            m.SetGpus(res)
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
    res["iloUserCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_iloUserCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloUserCredential(val.(V1beta1SystemsItemServersGetResponse_items_iloUserCredentialable))
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
    res["linkLocalInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_linkLocalInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkLocalInfo(val.(V1beta1SystemsItemServersGetResponse_items_linkLocalInfoable))
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
    res["operatingSystemInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_operatingSystemInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemInfo(val.(V1beta1SystemsItemServersGetResponse_items_operatingSystemInfoable))
        }
        return nil
    }
    res["osUserCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_osUserCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsUserCredential(val.(V1beta1SystemsItemServersGetResponse_items_osUserCredentialable))
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
// GetGpus gets the gpus property value. Gpus in the server.
// returns a []V1beta1SystemsItemServersGetResponse_items_gpusable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetGpus()([]V1beta1SystemsItemServersGetResponse_items_gpusable) {
    return m.gpus
}
// GetHypervisorHost gets the hypervisorHost property value. id, URI to reference the hypervisor host. Soon to be deprecated.
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
// GetIloNetworkInfo gets the iloNetworkInfo property value. ILO network information.
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
// GetIloUserCredential gets the iloUserCredential property value. iLO user name and authentication credential secret. Property iloAdminUserPasswordSecret will be deprecated.Use this property instead.
// returns a V1beta1SystemsItemServersGetResponse_items_iloUserCredentialable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetIloUserCredential()(V1beta1SystemsItemServersGetResponse_items_iloUserCredentialable) {
    return m.iloUserCredential
}
// GetIndicatorLedStatus gets the indicatorLedStatus property value. iLO LED indication.
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetIndicatorLedStatus()(*string) {
    return m.indicatorLedStatus
}
// GetLinkLocalInfo gets the linkLocalInfo property value. Information about link local IP addresses.
// returns a V1beta1SystemsItemServersGetResponse_items_linkLocalInfoable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetLinkLocalInfo()(V1beta1SystemsItemServersGetResponse_items_linkLocalInfoable) {
    return m.linkLocalInfo
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
// GetOperatingSystemInfo gets the operatingSystemInfo property value. Information about the OS on the server.
// returns a V1beta1SystemsItemServersGetResponse_items_operatingSystemInfoable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetOperatingSystemInfo()(V1beta1SystemsItemServersGetResponse_items_operatingSystemInfoable) {
    return m.operatingSystemInfo
}
// GetOsUserCredential gets the osUserCredential property value. Operating system user name and authentication credential secret. Property hypervisorHostRootUserPasswordSecret will be deprecated.Use this property instead.
// returns a V1beta1SystemsItemServersGetResponse_items_osUserCredentialable when successful
func (m *V1beta1SystemsItemServersGetResponse_items) GetOsUserCredential()(V1beta1SystemsItemServersGetResponse_items_osUserCredentialable) {
    return m.osUserCredential
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
// GetSystemId gets the systemId property value. Unique identifier of the system, usually a UUID.
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
        err := writer.WriteObjectValue("cloudModuleConfig", m.GetCloudModuleConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("computeUsage", m.GetComputeUsage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("consumedBy", m.GetConsumedBy())
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
    if m.GetGpus() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGpus()))
        for i, v := range m.GetGpus() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("gpus", cast)
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
        err := writer.WriteObjectValue("iloUserCredential", m.GetIloUserCredential())
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
        err := writer.WriteObjectValue("linkLocalInfo", m.GetLinkLocalInfo())
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
        err := writer.WriteObjectValue("operatingSystemInfo", m.GetOperatingSystemInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("osUserCredential", m.GetOsUserCredential())
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
// SetCloudModuleConfig sets the cloudModuleConfig property value. HPE Cloud module configuration.
func (m *V1beta1SystemsItemServersGetResponse_items) SetCloudModuleConfig(value V1beta1SystemsItemServersGetResponse_items_cloudModuleConfigable)() {
    m.cloudModuleConfig = value
}
// SetComputeUsage sets the computeUsage property value. Server Compute Usage Information.
func (m *V1beta1SystemsItemServersGetResponse_items) SetComputeUsage(value V1beta1SystemsItemServersGetResponse_items_computeUsageable)() {
    m.computeUsage = value
}
// SetConsumedBy sets the consumedBy property value. The consumedBy property
func (m *V1beta1SystemsItemServersGetResponse_items) SetConsumedBy(value V1beta1SystemsItemServersGetResponse_items_consumedByable)() {
    m.consumedBy = value
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
// SetGpus sets the gpus property value. Gpus in the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetGpus(value []V1beta1SystemsItemServersGetResponse_items_gpusable)() {
    m.gpus = value
}
// SetHypervisorHost sets the hypervisorHost property value. id, URI to reference the hypervisor host. Soon to be deprecated.
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
// SetIloNetworkInfo sets the iloNetworkInfo property value. ILO network information.
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
// SetIloUserCredential sets the iloUserCredential property value. iLO user name and authentication credential secret. Property iloAdminUserPasswordSecret will be deprecated.Use this property instead.
func (m *V1beta1SystemsItemServersGetResponse_items) SetIloUserCredential(value V1beta1SystemsItemServersGetResponse_items_iloUserCredentialable)() {
    m.iloUserCredential = value
}
// SetIndicatorLedStatus sets the indicatorLedStatus property value. iLO LED indication.
func (m *V1beta1SystemsItemServersGetResponse_items) SetIndicatorLedStatus(value *string)() {
    m.indicatorLedStatus = value
}
// SetLinkLocalInfo sets the linkLocalInfo property value. Information about link local IP addresses.
func (m *V1beta1SystemsItemServersGetResponse_items) SetLinkLocalInfo(value V1beta1SystemsItemServersGetResponse_items_linkLocalInfoable)() {
    m.linkLocalInfo = value
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
// SetOperatingSystemInfo sets the operatingSystemInfo property value. Information about the OS on the server.
func (m *V1beta1SystemsItemServersGetResponse_items) SetOperatingSystemInfo(value V1beta1SystemsItemServersGetResponse_items_operatingSystemInfoable)() {
    m.operatingSystemInfo = value
}
// SetOsUserCredential sets the osUserCredential property value. Operating system user name and authentication credential secret. Property hypervisorHostRootUserPasswordSecret will be deprecated.Use this property instead.
func (m *V1beta1SystemsItemServersGetResponse_items) SetOsUserCredential(value V1beta1SystemsItemServersGetResponse_items_osUserCredentialable)() {
    m.osUserCredential = value
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
// SetSystemId sets the systemId property value. Unique identifier of the system, usually a UUID.
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
    GetCloudModuleConfig()(V1beta1SystemsItemServersGetResponse_items_cloudModuleConfigable)
    GetComputeUsage()(V1beta1SystemsItemServersGetResponse_items_computeUsageable)
    GetConsumedBy()(V1beta1SystemsItemServersGetResponse_items_consumedByable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetGpus()([]V1beta1SystemsItemServersGetResponse_items_gpusable)
    GetHypervisorHost()(V1beta1SystemsItemServersGetResponse_items_hypervisorHostable)
    GetHypervisorHostRootUserPasswordSecret()(V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable)
    GetId()(*string)
    GetIloAdminUserPasswordSecret()(V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable)
    GetIloFirmwareVersion()(*string)
    GetIloNetworkInfo()(V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable)
    GetIloState()(*string)
    GetIloStatus()(*string)
    GetIloUserCredential()(V1beta1SystemsItemServersGetResponse_items_iloUserCredentialable)
    GetIndicatorLedStatus()(*string)
    GetLinkLocalInfo()(V1beta1SystemsItemServersGetResponse_items_linkLocalInfoable)
    GetMemoryGib()(*string)
    GetModel()(*string)
    GetName()(*string)
    GetNcmVersion()(*string)
    GetOnPremAgentId()(*string)
    GetOperatingSystemInfo()(V1beta1SystemsItemServersGetResponse_items_operatingSystemInfoable)
    GetOsUserCredential()(V1beta1SystemsItemServersGetResponse_items_osUserCredentialable)
    GetProcessorCount()(*string)
    GetProcessorModel()(*string)
    GetResourceUri()(*string)
    GetSerialNumber()(*string)
    GetServerNetworkInterfaces()([]V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable)
    GetSystemId()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCloudModuleConfig(value V1beta1SystemsItemServersGetResponse_items_cloudModuleConfigable)()
    SetComputeUsage(value V1beta1SystemsItemServersGetResponse_items_computeUsageable)()
    SetConsumedBy(value V1beta1SystemsItemServersGetResponse_items_consumedByable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetGpus(value []V1beta1SystemsItemServersGetResponse_items_gpusable)()
    SetHypervisorHost(value V1beta1SystemsItemServersGetResponse_items_hypervisorHostable)()
    SetHypervisorHostRootUserPasswordSecret(value V1beta1SystemsItemServersGetResponse_items_hypervisorHostRootUserPasswordSecretable)()
    SetId(value *string)()
    SetIloAdminUserPasswordSecret(value V1beta1SystemsItemServersGetResponse_items_iloAdminUserPasswordSecretable)()
    SetIloFirmwareVersion(value *string)()
    SetIloNetworkInfo(value V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable)()
    SetIloState(value *string)()
    SetIloStatus(value *string)()
    SetIloUserCredential(value V1beta1SystemsItemServersGetResponse_items_iloUserCredentialable)()
    SetIndicatorLedStatus(value *string)()
    SetLinkLocalInfo(value V1beta1SystemsItemServersGetResponse_items_linkLocalInfoable)()
    SetMemoryGib(value *string)()
    SetModel(value *string)()
    SetName(value *string)()
    SetNcmVersion(value *string)()
    SetOnPremAgentId(value *string)()
    SetOperatingSystemInfo(value V1beta1SystemsItemServersGetResponse_items_operatingSystemInfoable)()
    SetOsUserCredential(value V1beta1SystemsItemServersGetResponse_items_osUserCredentialable)()
    SetProcessorCount(value *string)()
    SetProcessorModel(value *string)()
    SetResourceUri(value *string)()
    SetSerialNumber(value *string)()
    SetServerNetworkInterfaces(value []V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable)()
    SetSystemId(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
