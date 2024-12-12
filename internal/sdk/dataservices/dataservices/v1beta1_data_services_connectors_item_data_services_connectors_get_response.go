package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse detailed information of Data Services Collector
type V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of available software versions for upgrade
    availableVersions []string
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // This details the current date and time of the Data Services Connector, how the current date and time is determined, as well as the user specified timezone.
    dateTime V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_dateTimeable
    // Device model of Data Services Connector.
    deviceModel *string
    // User-defined name given to Data Services Connector.
    displayName *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // List of hypervisor managers configured on the Data Services Connector.
    hypervisorManagers []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_hypervisorManagersable
    // Unique ID identifying a Data Services Connector.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Indicates whether Infosight is enabled or disabled.
    infosightEnabled *bool
    // The interfaces property
    interfaces V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_interfacesable
    // UTC time when the Data Services Connector was last updated.
    lastUpdateCheckTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of Data Services Connector.
    name *string
    // The ntp property
    ntp V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpable
    // Data Services Connector persona.
    persona *string
    // UTC time when Data Services Connector was last powered on.
    poweredOnAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Product name of Data Services Connector.
    productName *string
    // This provides details of the RDA for Data Services Connector.
    rdaInfo V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoable
    // RDA Station ID identifying a Data Services Connector.
    remoteAccessStationId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The self reference for this resource.
    resourceUri *string
    // Data Services Connector serial number.
    serialNumber *string
    // List of services with name and version configured on the Data Services Connector.
    serviceVersions []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsable
    // Data Services Connector software version.
    softwareVersion *string
    // Reason for the Data Services Connector being in the current state. This will be empty when the Data Services Connector is in an OK state.
    stateReason *string
    // Serial number of the Alletra MP Storage System.
    storageSystemSerialNumber *string
    // Total RAM configured for the Data Services Connector in GiB.
    totalMemoryInGiB *int32
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Returns the number of seconds since Data Services Connector was powered on.
    upTimeInSeconds *float64
    // Number of virtual CPUs configured for the Data Services Connector.
    vCpu *int32
}
// NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse instantiates a new V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse and sets the default values.
func NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse()(*V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) {
    m := &V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableVersions gets the availableVersions property value. List of available software versions for upgrade
// returns a []string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetAvailableVersions()([]string) {
    return m.availableVersions
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDateTime gets the dateTime property value. This details the current date and time of the Data Services Connector, how the current date and time is determined, as well as the user specified timezone.
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_dateTimeable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetDateTime()(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_dateTimeable) {
    return m.dateTime
}
// GetDeviceModel gets the deviceModel property value. Device model of Data Services Connector.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetDeviceModel()(*string) {
    return m.deviceModel
}
// GetDisplayName gets the displayName property value. User-defined name given to Data Services Connector.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availableVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAvailableVersions(res)
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
    res["dateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_dateTimeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTime(val.(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_dateTimeable))
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["hypervisorManagers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_hypervisorManagersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_hypervisorManagersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_hypervisorManagersable)
                }
            }
            m.SetHypervisorManagers(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["infosightEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfosightEnabled(val)
        }
        return nil
    }
    res["interfaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_interfacesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInterfaces(val.(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_interfacesable))
        }
        return nil
    }
    res["lastUpdateCheckTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateCheckTime(val)
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
    res["ntp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNtp(val.(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpable))
        }
        return nil
    }
    res["persona"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersona(val)
        }
        return nil
    }
    res["poweredOnAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPoweredOnAt(val)
        }
        return nil
    }
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["rdaInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRdaInfo(val.(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoable))
        }
        return nil
    }
    res["remoteAccessStationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAccessStationId(val)
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
    res["serviceVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsable)
                }
            }
            m.SetServiceVersions(res)
        }
        return nil
    }
    res["softwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareVersion(val)
        }
        return nil
    }
    res["stateReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateReason(val)
        }
        return nil
    }
    res["storageSystemSerialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageSystemSerialNumber(val)
        }
        return nil
    }
    res["totalMemoryInGiB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalMemoryInGiB(val)
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
    res["upTimeInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpTimeInSeconds(val)
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
    res["vCpu"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVCpu(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetHypervisorManagers gets the hypervisorManagers property value. List of hypervisor managers configured on the Data Services Connector.
// returns a []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_hypervisorManagersable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetHypervisorManagers()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_hypervisorManagersable) {
    return m.hypervisorManagers
}
// GetId gets the id property value. Unique ID identifying a Data Services Connector.
// returns a *UUID when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetInfosightEnabled gets the infosightEnabled property value. Indicates whether Infosight is enabled or disabled.
// returns a *bool when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetInfosightEnabled()(*bool) {
    return m.infosightEnabled
}
// GetInterfaces gets the interfaces property value. The interfaces property
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_interfacesable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetInterfaces()(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_interfacesable) {
    return m.interfaces
}
// GetLastUpdateCheckTime gets the lastUpdateCheckTime property value. UTC time when the Data Services Connector was last updated.
// returns a *Time when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetLastUpdateCheckTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdateCheckTime
}
// GetName gets the name property value. Name of Data Services Connector.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetName()(*string) {
    return m.name
}
// GetNtp gets the ntp property value. The ntp property
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetNtp()(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpable) {
    return m.ntp
}
// GetPersona gets the persona property value. Data Services Connector persona.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetPersona()(*string) {
    return m.persona
}
// GetPoweredOnAt gets the poweredOnAt property value. UTC time when Data Services Connector was last powered on.
// returns a *Time when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetPoweredOnAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.poweredOnAt
}
// GetProductName gets the productName property value. Product name of Data Services Connector.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetProductName()(*string) {
    return m.productName
}
// GetRdaInfo gets the rdaInfo property value. This provides details of the RDA for Data Services Connector.
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetRdaInfo()(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoable) {
    return m.rdaInfo
}
// GetRemoteAccessStationId gets the remoteAccessStationId property value. RDA Station ID identifying a Data Services Connector.
// returns a *UUID when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetRemoteAccessStationId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.remoteAccessStationId
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSerialNumber gets the serialNumber property value. Data Services Connector serial number.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetServiceVersions gets the serviceVersions property value. List of services with name and version configured on the Data Services Connector.
// returns a []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetServiceVersions()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsable) {
    return m.serviceVersions
}
// GetSoftwareVersion gets the softwareVersion property value. Data Services Connector software version.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetSoftwareVersion()(*string) {
    return m.softwareVersion
}
// GetStateReason gets the stateReason property value. Reason for the Data Services Connector being in the current state. This will be empty when the Data Services Connector is in an OK state.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetStateReason()(*string) {
    return m.stateReason
}
// GetStorageSystemSerialNumber gets the storageSystemSerialNumber property value. Serial number of the Alletra MP Storage System.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetStorageSystemSerialNumber()(*string) {
    return m.storageSystemSerialNumber
}
// GetTotalMemoryInGiB gets the totalMemoryInGiB property value. Total RAM configured for the Data Services Connector in GiB.
// returns a *int32 when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetTotalMemoryInGiB()(*int32) {
    return m.totalMemoryInGiB
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetUpTimeInSeconds gets the upTimeInSeconds property value. Returns the number of seconds since Data Services Connector was powered on.
// returns a *float64 when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetUpTimeInSeconds()(*float64) {
    return m.upTimeInSeconds
}
// GetVCpu gets the vCpu property value. Number of virtual CPUs configured for the Data Services Connector.
// returns a *int32 when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) GetVCpu()(*int32) {
    return m.vCpu
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAvailableVersions() != nil {
        err := writer.WriteCollectionOfStringValues("availableVersions", m.GetAvailableVersions())
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
        err := writer.WriteObjectValue("dateTime", m.GetDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetHypervisorManagers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHypervisorManagers()))
        for i, v := range m.GetHypervisorManagers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("hypervisorManagers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("infosightEnabled", m.GetInfosightEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("interfaces", m.GetInterfaces())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUpdateCheckTime", m.GetLastUpdateCheckTime())
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
        err := writer.WriteObjectValue("ntp", m.GetNtp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("persona", m.GetPersona())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("poweredOnAt", m.GetPoweredOnAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rdaInfo", m.GetRdaInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("remoteAccessStationId", m.GetRemoteAccessStationId())
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
    if m.GetServiceVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServiceVersions()))
        for i, v := range m.GetServiceVersions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("serviceVersions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("softwareVersion", m.GetSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stateReason", m.GetStateReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("storageSystemSerialNumber", m.GetStorageSystemSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalMemoryInGiB", m.GetTotalMemoryInGiB())
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
        err := writer.WriteFloat64Value("upTimeInSeconds", m.GetUpTimeInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("vCpu", m.GetVCpu())
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
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableVersions sets the availableVersions property value. List of available software versions for upgrade
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetAvailableVersions(value []string)() {
    m.availableVersions = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDateTime sets the dateTime property value. This details the current date and time of the Data Services Connector, how the current date and time is determined, as well as the user specified timezone.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetDateTime(value V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_dateTimeable)() {
    m.dateTime = value
}
// SetDeviceModel sets the deviceModel property value. Device model of Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// SetDisplayName sets the displayName property value. User-defined name given to Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHypervisorManagers sets the hypervisorManagers property value. List of hypervisor managers configured on the Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetHypervisorManagers(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_hypervisorManagersable)() {
    m.hypervisorManagers = value
}
// SetId sets the id property value. Unique ID identifying a Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetInfosightEnabled sets the infosightEnabled property value. Indicates whether Infosight is enabled or disabled.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetInfosightEnabled(value *bool)() {
    m.infosightEnabled = value
}
// SetInterfaces sets the interfaces property value. The interfaces property
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetInterfaces(value V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_interfacesable)() {
    m.interfaces = value
}
// SetLastUpdateCheckTime sets the lastUpdateCheckTime property value. UTC time when the Data Services Connector was last updated.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetLastUpdateCheckTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateCheckTime = value
}
// SetName sets the name property value. Name of Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetName(value *string)() {
    m.name = value
}
// SetNtp sets the ntp property value. The ntp property
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetNtp(value V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpable)() {
    m.ntp = value
}
// SetPersona sets the persona property value. Data Services Connector persona.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetPersona(value *string)() {
    m.persona = value
}
// SetPoweredOnAt sets the poweredOnAt property value. UTC time when Data Services Connector was last powered on.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetPoweredOnAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.poweredOnAt = value
}
// SetProductName sets the productName property value. Product name of Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetProductName(value *string)() {
    m.productName = value
}
// SetRdaInfo sets the rdaInfo property value. This provides details of the RDA for Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetRdaInfo(value V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoable)() {
    m.rdaInfo = value
}
// SetRemoteAccessStationId sets the remoteAccessStationId property value. RDA Station ID identifying a Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetRemoteAccessStationId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.remoteAccessStationId = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSerialNumber sets the serialNumber property value. Data Services Connector serial number.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetServiceVersions sets the serviceVersions property value. List of services with name and version configured on the Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetServiceVersions(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsable)() {
    m.serviceVersions = value
}
// SetSoftwareVersion sets the softwareVersion property value. Data Services Connector software version.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetSoftwareVersion(value *string)() {
    m.softwareVersion = value
}
// SetStateReason sets the stateReason property value. Reason for the Data Services Connector being in the current state. This will be empty when the Data Services Connector is in an OK state.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetStateReason(value *string)() {
    m.stateReason = value
}
// SetStorageSystemSerialNumber sets the storageSystemSerialNumber property value. Serial number of the Alletra MP Storage System.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetStorageSystemSerialNumber(value *string)() {
    m.storageSystemSerialNumber = value
}
// SetTotalMemoryInGiB sets the totalMemoryInGiB property value. Total RAM configured for the Data Services Connector in GiB.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetTotalMemoryInGiB(value *int32)() {
    m.totalMemoryInGiB = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetUpTimeInSeconds sets the upTimeInSeconds property value. Returns the number of seconds since Data Services Connector was powered on.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetUpTimeInSeconds(value *float64)() {
    m.upTimeInSeconds = value
}
// SetVCpu sets the vCpu property value. Number of virtual CPUs configured for the Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse) SetVCpu(value *int32)() {
    m.vCpu = value
}
type V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableVersions()([]string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDateTime()(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_dateTimeable)
    GetDeviceModel()(*string)
    GetDisplayName()(*string)
    GetGeneration()(*int64)
    GetHypervisorManagers()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_hypervisorManagersable)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetInfosightEnabled()(*bool)
    GetInterfaces()(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_interfacesable)
    GetLastUpdateCheckTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetNtp()(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpable)
    GetPersona()(*string)
    GetPoweredOnAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProductName()(*string)
    GetRdaInfo()(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoable)
    GetRemoteAccessStationId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetResourceUri()(*string)
    GetSerialNumber()(*string)
    GetServiceVersions()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsable)
    GetSoftwareVersion()(*string)
    GetStateReason()(*string)
    GetStorageSystemSerialNumber()(*string)
    GetTotalMemoryInGiB()(*int32)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUpTimeInSeconds()(*float64)
    GetVCpu()(*int32)
    SetAvailableVersions(value []string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDateTime(value V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_dateTimeable)()
    SetDeviceModel(value *string)()
    SetDisplayName(value *string)()
    SetGeneration(value *int64)()
    SetHypervisorManagers(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_hypervisorManagersable)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetInfosightEnabled(value *bool)()
    SetInterfaces(value V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_interfacesable)()
    SetLastUpdateCheckTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetNtp(value V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_ntpable)()
    SetPersona(value *string)()
    SetPoweredOnAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProductName(value *string)()
    SetRdaInfo(value V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoable)()
    SetRemoteAccessStationId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetResourceUri(value *string)()
    SetSerialNumber(value *string)()
    SetServiceVersions(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_serviceVersionsable)()
    SetSoftwareVersion(value *string)()
    SetStateReason(value *string)()
    SetStorageSystemSerialNumber(value *string)()
    SetTotalMemoryInGiB(value *int32)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUpTimeInSeconds(value *float64)()
    SetVCpu(value *int32)()
}
