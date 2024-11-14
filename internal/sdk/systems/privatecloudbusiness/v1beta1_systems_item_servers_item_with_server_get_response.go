package privatecloudbusiness

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersItemWithServerGetResponse details of the Server with select information.
type V1beta1SystemsItemServersItemWithServerGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // HPE Cloud module configuration.
    cloudModuleConfig V1beta1SystemsItemServersItemWithServerGetResponse_cloudModuleConfigable
    // Server Compute Usage Information.
    computeUsage V1beta1SystemsItemServersItemWithServerGetResponse_computeUsageable
    // The consumedBy property
    consumedBy V1beta1SystemsItemServersItemWithServerGetResponse_consumedByable
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // Gpus in the server.
    gpus []V1beta1SystemsItemServersItemWithServerGetResponse_gpusable
    // id, URI to reference the hypervisor host. Soon to be deprecated.
    hypervisorHost V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostable
    // Secret information
    hypervisorHostRootUserPasswordSecret V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostRootUserPasswordSecretable
    // An identifier for the resource, usually a UUID.
    id *string
    // Secret information
    iloAdminUserPasswordSecret V1beta1SystemsItemServersItemWithServerGetResponse_iloAdminUserPasswordSecretable
    // firmware version of iLO in the server.
    iloFirmwareVersion *string
    // ILO network information.
    iloNetworkInfo V1beta1SystemsItemServersItemWithServerGetResponse_iloNetworkInfoable
    // state of the iLO in the server.
    iloState *string
    // status of the iLO in the server.
    iloStatus *string
    // iLO user name and authentication credential secret. Property iloAdminUserPasswordSecret will be deprecated.Use this property instead.
    iloUserCredential V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredentialable
    // iLO LED indication.
    indicatorLedStatus *string
    // Information about link local IP addresses.
    linkLocalInfo V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoable
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
    operatingSystemInfo V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoable
    // Operating system user name and authentication credential secret. Property hypervisorHostRootUserPasswordSecret will be deprecated.Use this property instead.
    osUserCredential V1beta1SystemsItemServersItemWithServerGetResponse_osUserCredentialable
    // Number of processors in the server.
    processorCount *string
    // Model of the processors in the server.
    processorModel *string
    // The self reference for this resource.
    resourceUri *string
    // Server's serial number.
    serialNumber *string
    // List of server network interfaces
    serverNetworkInterfaces []V1beta1SystemsItemServersItemWithServerGetResponse_serverNetworkInterfacesable
    // Unique identifier of the system, usually a UUID.
    systemId *string
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1SystemsItemServersItemWithServerGetResponse instantiates a new V1beta1SystemsItemServersItemWithServerGetResponse and sets the default values.
func NewV1beta1SystemsItemServersItemWithServerGetResponse()(*V1beta1SystemsItemServersItemWithServerGetResponse) {
    m := &V1beta1SystemsItemServersItemWithServerGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersItemWithServerGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersItemWithServerGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersItemWithServerGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCloudModuleConfig gets the cloudModuleConfig property value. HPE Cloud module configuration.
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_cloudModuleConfigable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetCloudModuleConfig()(V1beta1SystemsItemServersItemWithServerGetResponse_cloudModuleConfigable) {
    return m.cloudModuleConfig
}
// GetComputeUsage gets the computeUsage property value. Server Compute Usage Information.
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_computeUsageable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetComputeUsage()(V1beta1SystemsItemServersItemWithServerGetResponse_computeUsageable) {
    return m.computeUsage
}
// GetConsumedBy gets the consumedBy property value. The consumedBy property
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_consumedByable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetConsumedBy()(V1beta1SystemsItemServersItemWithServerGetResponse_consumedByable) {
    return m.consumedBy
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudModuleConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_cloudModuleConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudModuleConfig(val.(V1beta1SystemsItemServersItemWithServerGetResponse_cloudModuleConfigable))
        }
        return nil
    }
    res["computeUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_computeUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputeUsage(val.(V1beta1SystemsItemServersItemWithServerGetResponse_computeUsageable))
        }
        return nil
    }
    res["consumedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_consumedByFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsumedBy(val.(V1beta1SystemsItemServersItemWithServerGetResponse_consumedByable))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemServersItemWithServerGetResponse_gpusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemServersItemWithServerGetResponse_gpusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemServersItemWithServerGetResponse_gpusable)
                }
            }
            m.SetGpus(res)
        }
        return nil
    }
    res["hypervisorHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorHost(val.(V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostable))
        }
        return nil
    }
    res["hypervisorHostRootUserPasswordSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostRootUserPasswordSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorHostRootUserPasswordSecret(val.(V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostRootUserPasswordSecretable))
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
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_iloAdminUserPasswordSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloAdminUserPasswordSecret(val.(V1beta1SystemsItemServersItemWithServerGetResponse_iloAdminUserPasswordSecretable))
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
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_iloNetworkInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloNetworkInfo(val.(V1beta1SystemsItemServersItemWithServerGetResponse_iloNetworkInfoable))
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
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloUserCredential(val.(V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredentialable))
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
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkLocalInfo(val.(V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoable))
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
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemInfo(val.(V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoable))
        }
        return nil
    }
    res["osUserCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersItemWithServerGetResponse_osUserCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsUserCredential(val.(V1beta1SystemsItemServersItemWithServerGetResponse_osUserCredentialable))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemServersItemWithServerGetResponse_serverNetworkInterfacesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemServersItemWithServerGetResponse_serverNetworkInterfacesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemServersItemWithServerGetResponse_serverNetworkInterfacesable)
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
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetGpus gets the gpus property value. Gpus in the server.
// returns a []V1beta1SystemsItemServersItemWithServerGetResponse_gpusable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetGpus()([]V1beta1SystemsItemServersItemWithServerGetResponse_gpusable) {
    return m.gpus
}
// GetHypervisorHost gets the hypervisorHost property value. id, URI to reference the hypervisor host. Soon to be deprecated.
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetHypervisorHost()(V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostable) {
    return m.hypervisorHost
}
// GetHypervisorHostRootUserPasswordSecret gets the hypervisorHostRootUserPasswordSecret property value. Secret information
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostRootUserPasswordSecretable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetHypervisorHostRootUserPasswordSecret()(V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostRootUserPasswordSecretable) {
    return m.hypervisorHostRootUserPasswordSecret
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetId()(*string) {
    return m.id
}
// GetIloAdminUserPasswordSecret gets the iloAdminUserPasswordSecret property value. Secret information
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_iloAdminUserPasswordSecretable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetIloAdminUserPasswordSecret()(V1beta1SystemsItemServersItemWithServerGetResponse_iloAdminUserPasswordSecretable) {
    return m.iloAdminUserPasswordSecret
}
// GetIloFirmwareVersion gets the iloFirmwareVersion property value. firmware version of iLO in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetIloFirmwareVersion()(*string) {
    return m.iloFirmwareVersion
}
// GetIloNetworkInfo gets the iloNetworkInfo property value. ILO network information.
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_iloNetworkInfoable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetIloNetworkInfo()(V1beta1SystemsItemServersItemWithServerGetResponse_iloNetworkInfoable) {
    return m.iloNetworkInfo
}
// GetIloState gets the iloState property value. state of the iLO in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetIloState()(*string) {
    return m.iloState
}
// GetIloStatus gets the iloStatus property value. status of the iLO in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetIloStatus()(*string) {
    return m.iloStatus
}
// GetIloUserCredential gets the iloUserCredential property value. iLO user name and authentication credential secret. Property iloAdminUserPasswordSecret will be deprecated.Use this property instead.
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredentialable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetIloUserCredential()(V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredentialable) {
    return m.iloUserCredential
}
// GetIndicatorLedStatus gets the indicatorLedStatus property value. iLO LED indication.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetIndicatorLedStatus()(*string) {
    return m.indicatorLedStatus
}
// GetLinkLocalInfo gets the linkLocalInfo property value. Information about link local IP addresses.
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetLinkLocalInfo()(V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoable) {
    return m.linkLocalInfo
}
// GetMemoryGib gets the memoryGib property value. Memory of the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetMemoryGib()(*string) {
    return m.memoryGib
}
// GetModel gets the model property value. The model of the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetModel()(*string) {
    return m.model
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetName()(*string) {
    return m.name
}
// GetNcmVersion gets the ncmVersion property value. Version of NCM installed on the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetNcmVersion()(*string) {
    return m.ncmVersion
}
// GetOnPremAgentId gets the onPremAgentId property value. An identifier of the on prem agent for the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetOnPremAgentId()(*string) {
    return m.onPremAgentId
}
// GetOperatingSystemInfo gets the operatingSystemInfo property value. Information about the OS on the server.
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetOperatingSystemInfo()(V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoable) {
    return m.operatingSystemInfo
}
// GetOsUserCredential gets the osUserCredential property value. Operating system user name and authentication credential secret. Property hypervisorHostRootUserPasswordSecret will be deprecated.Use this property instead.
// returns a V1beta1SystemsItemServersItemWithServerGetResponse_osUserCredentialable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetOsUserCredential()(V1beta1SystemsItemServersItemWithServerGetResponse_osUserCredentialable) {
    return m.osUserCredential
}
// GetProcessorCount gets the processorCount property value. Number of processors in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetProcessorCount()(*string) {
    return m.processorCount
}
// GetProcessorModel gets the processorModel property value. Model of the processors in the server.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetProcessorModel()(*string) {
    return m.processorModel
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSerialNumber gets the serialNumber property value. Server's serial number.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetServerNetworkInterfaces gets the serverNetworkInterfaces property value. List of server network interfaces
// returns a []V1beta1SystemsItemServersItemWithServerGetResponse_serverNetworkInterfacesable when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetServerNetworkInterfaces()([]V1beta1SystemsItemServersItemWithServerGetResponse_serverNetworkInterfacesable) {
    return m.serverNetworkInterfaces
}
// GetSystemId gets the systemId property value. Unique identifier of the system, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetSystemId()(*string) {
    return m.systemId
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCloudModuleConfig sets the cloudModuleConfig property value. HPE Cloud module configuration.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetCloudModuleConfig(value V1beta1SystemsItemServersItemWithServerGetResponse_cloudModuleConfigable)() {
    m.cloudModuleConfig = value
}
// SetComputeUsage sets the computeUsage property value. Server Compute Usage Information.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetComputeUsage(value V1beta1SystemsItemServersItemWithServerGetResponse_computeUsageable)() {
    m.computeUsage = value
}
// SetConsumedBy sets the consumedBy property value. The consumedBy property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetConsumedBy(value V1beta1SystemsItemServersItemWithServerGetResponse_consumedByable)() {
    m.consumedBy = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetGpus sets the gpus property value. Gpus in the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetGpus(value []V1beta1SystemsItemServersItemWithServerGetResponse_gpusable)() {
    m.gpus = value
}
// SetHypervisorHost sets the hypervisorHost property value. id, URI to reference the hypervisor host. Soon to be deprecated.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetHypervisorHost(value V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostable)() {
    m.hypervisorHost = value
}
// SetHypervisorHostRootUserPasswordSecret sets the hypervisorHostRootUserPasswordSecret property value. Secret information
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetHypervisorHostRootUserPasswordSecret(value V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostRootUserPasswordSecretable)() {
    m.hypervisorHostRootUserPasswordSecret = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetId(value *string)() {
    m.id = value
}
// SetIloAdminUserPasswordSecret sets the iloAdminUserPasswordSecret property value. Secret information
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetIloAdminUserPasswordSecret(value V1beta1SystemsItemServersItemWithServerGetResponse_iloAdminUserPasswordSecretable)() {
    m.iloAdminUserPasswordSecret = value
}
// SetIloFirmwareVersion sets the iloFirmwareVersion property value. firmware version of iLO in the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetIloFirmwareVersion(value *string)() {
    m.iloFirmwareVersion = value
}
// SetIloNetworkInfo sets the iloNetworkInfo property value. ILO network information.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetIloNetworkInfo(value V1beta1SystemsItemServersItemWithServerGetResponse_iloNetworkInfoable)() {
    m.iloNetworkInfo = value
}
// SetIloState sets the iloState property value. state of the iLO in the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetIloState(value *string)() {
    m.iloState = value
}
// SetIloStatus sets the iloStatus property value. status of the iLO in the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetIloStatus(value *string)() {
    m.iloStatus = value
}
// SetIloUserCredential sets the iloUserCredential property value. iLO user name and authentication credential secret. Property iloAdminUserPasswordSecret will be deprecated.Use this property instead.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetIloUserCredential(value V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredentialable)() {
    m.iloUserCredential = value
}
// SetIndicatorLedStatus sets the indicatorLedStatus property value. iLO LED indication.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetIndicatorLedStatus(value *string)() {
    m.indicatorLedStatus = value
}
// SetLinkLocalInfo sets the linkLocalInfo property value. Information about link local IP addresses.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetLinkLocalInfo(value V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoable)() {
    m.linkLocalInfo = value
}
// SetMemoryGib sets the memoryGib property value. Memory of the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetMemoryGib(value *string)() {
    m.memoryGib = value
}
// SetModel sets the model property value. The model of the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetModel(value *string)() {
    m.model = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetName(value *string)() {
    m.name = value
}
// SetNcmVersion sets the ncmVersion property value. Version of NCM installed on the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetNcmVersion(value *string)() {
    m.ncmVersion = value
}
// SetOnPremAgentId sets the onPremAgentId property value. An identifier of the on prem agent for the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetOnPremAgentId(value *string)() {
    m.onPremAgentId = value
}
// SetOperatingSystemInfo sets the operatingSystemInfo property value. Information about the OS on the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetOperatingSystemInfo(value V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoable)() {
    m.operatingSystemInfo = value
}
// SetOsUserCredential sets the osUserCredential property value. Operating system user name and authentication credential secret. Property hypervisorHostRootUserPasswordSecret will be deprecated.Use this property instead.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetOsUserCredential(value V1beta1SystemsItemServersItemWithServerGetResponse_osUserCredentialable)() {
    m.osUserCredential = value
}
// SetProcessorCount sets the processorCount property value. Number of processors in the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetProcessorCount(value *string)() {
    m.processorCount = value
}
// SetProcessorModel sets the processorModel property value. Model of the processors in the server.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetProcessorModel(value *string)() {
    m.processorModel = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSerialNumber sets the serialNumber property value. Server's serial number.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetServerNetworkInterfaces sets the serverNetworkInterfaces property value. List of server network interfaces
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetServerNetworkInterfaces(value []V1beta1SystemsItemServersItemWithServerGetResponse_serverNetworkInterfacesable)() {
    m.serverNetworkInterfaces = value
}
// SetSystemId sets the systemId property value. Unique identifier of the system, usually a UUID.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetSystemId(value *string)() {
    m.systemId = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1SystemsItemServersItemWithServerGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudModuleConfig()(V1beta1SystemsItemServersItemWithServerGetResponse_cloudModuleConfigable)
    GetComputeUsage()(V1beta1SystemsItemServersItemWithServerGetResponse_computeUsageable)
    GetConsumedBy()(V1beta1SystemsItemServersItemWithServerGetResponse_consumedByable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetGpus()([]V1beta1SystemsItemServersItemWithServerGetResponse_gpusable)
    GetHypervisorHost()(V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostable)
    GetHypervisorHostRootUserPasswordSecret()(V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostRootUserPasswordSecretable)
    GetId()(*string)
    GetIloAdminUserPasswordSecret()(V1beta1SystemsItemServersItemWithServerGetResponse_iloAdminUserPasswordSecretable)
    GetIloFirmwareVersion()(*string)
    GetIloNetworkInfo()(V1beta1SystemsItemServersItemWithServerGetResponse_iloNetworkInfoable)
    GetIloState()(*string)
    GetIloStatus()(*string)
    GetIloUserCredential()(V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredentialable)
    GetIndicatorLedStatus()(*string)
    GetLinkLocalInfo()(V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoable)
    GetMemoryGib()(*string)
    GetModel()(*string)
    GetName()(*string)
    GetNcmVersion()(*string)
    GetOnPremAgentId()(*string)
    GetOperatingSystemInfo()(V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoable)
    GetOsUserCredential()(V1beta1SystemsItemServersItemWithServerGetResponse_osUserCredentialable)
    GetProcessorCount()(*string)
    GetProcessorModel()(*string)
    GetResourceUri()(*string)
    GetSerialNumber()(*string)
    GetServerNetworkInterfaces()([]V1beta1SystemsItemServersItemWithServerGetResponse_serverNetworkInterfacesable)
    GetSystemId()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCloudModuleConfig(value V1beta1SystemsItemServersItemWithServerGetResponse_cloudModuleConfigable)()
    SetComputeUsage(value V1beta1SystemsItemServersItemWithServerGetResponse_computeUsageable)()
    SetConsumedBy(value V1beta1SystemsItemServersItemWithServerGetResponse_consumedByable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetGpus(value []V1beta1SystemsItemServersItemWithServerGetResponse_gpusable)()
    SetHypervisorHost(value V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostable)()
    SetHypervisorHostRootUserPasswordSecret(value V1beta1SystemsItemServersItemWithServerGetResponse_hypervisorHostRootUserPasswordSecretable)()
    SetId(value *string)()
    SetIloAdminUserPasswordSecret(value V1beta1SystemsItemServersItemWithServerGetResponse_iloAdminUserPasswordSecretable)()
    SetIloFirmwareVersion(value *string)()
    SetIloNetworkInfo(value V1beta1SystemsItemServersItemWithServerGetResponse_iloNetworkInfoable)()
    SetIloState(value *string)()
    SetIloStatus(value *string)()
    SetIloUserCredential(value V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredentialable)()
    SetIndicatorLedStatus(value *string)()
    SetLinkLocalInfo(value V1beta1SystemsItemServersItemWithServerGetResponse_linkLocalInfoable)()
    SetMemoryGib(value *string)()
    SetModel(value *string)()
    SetName(value *string)()
    SetNcmVersion(value *string)()
    SetOnPremAgentId(value *string)()
    SetOperatingSystemInfo(value V1beta1SystemsItemServersItemWithServerGetResponse_operatingSystemInfoable)()
    SetOsUserCredential(value V1beta1SystemsItemServersItemWithServerGetResponse_osUserCredentialable)()
    SetProcessorCount(value *string)()
    SetProcessorModel(value *string)()
    SetResourceUri(value *string)()
    SetSerialNumber(value *string)()
    SetServerNetworkInterfaces(value []V1beta1SystemsItemServersItemWithServerGetResponse_serverNetworkInterfacesable)()
    SetSystemId(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
