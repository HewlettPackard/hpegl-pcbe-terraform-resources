package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemHypervisorGetResponse represents a single instance of a hypervisor manager (Vmware vCenter)
type V1beta1HypervisorManagersItemHypervisorGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Hypervisor specific information.
    appInfo V1beta1HypervisorManagersItemHypervisorGetResponse_appInfoable
    // The hypervisor manager build details
    buildVersion *string
    // Time in UTC at which the object was created
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier.
    customerId *string
    // Data Orchestrator specific information.
    dataOrchestratorInfo V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoable
    // Data Services Connectors specific information.
    dataServicesConnectorsInfo []V1beta1HypervisorManagersItemHypervisorGetResponse_dataServicesConnectorsInfoable
    // A brief description of the hypervisor manager.
    description *string
    // User defined name for the hypervisor manager.
    displayName *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // UUID string uniquely identifying the hypervisor manager.
    id *string
    // Time in UTC at which the object was last refreshed.
    lastRefreshed *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name as reported by the hypervisor manager.
    name *string
    // An IP address or hostname or FQDN to address the hypervisor manager
    networkAddress *string
    // The hypervisor manager release version.
    releaseVersion *string
    // The 'self' reference for this resource.
    resourceUri *string
    // List of services this object belongs to.  This list can be used to filter specific services in the UI.
    services []string
    // Brief reason for the current state of the hypervisor manager
    stateReason *string
    // The type of resource.
    typeEscaped *string
    // A hypervisor manager provided durable UID. In case of VMware it will be instanceUUID of the vCenter
    uid *string
    // Time in UTC at which the object was last updated
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the user used to access the hypervisor server. Mutually exclusive with credentialInfo.
    username *string
}
// NewV1beta1HypervisorManagersItemHypervisorGetResponse instantiates a new V1beta1HypervisorManagersItemHypervisorGetResponse and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorGetResponse()(*V1beta1HypervisorManagersItemHypervisorGetResponse) {
    m := &V1beta1HypervisorManagersItemHypervisorGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorManagersItemHypervisorGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemHypervisorGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemHypervisorGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppInfo gets the appInfo property value. Hypervisor specific information.
// returns a V1beta1HypervisorManagersItemHypervisorGetResponse_appInfoable when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetAppInfo()(V1beta1HypervisorManagersItemHypervisorGetResponse_appInfoable) {
    return m.appInfo
}
// GetBuildVersion gets the buildVersion property value. The hypervisor manager build details
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetBuildVersion()(*string) {
    return m.buildVersion
}
// GetCreatedAt gets the createdAt property value. Time in UTC at which the object was created
// returns a *Time when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDataOrchestratorInfo gets the dataOrchestratorInfo property value. Data Orchestrator specific information.
// returns a V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoable when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetDataOrchestratorInfo()(V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoable) {
    return m.dataOrchestratorInfo
}
// GetDataServicesConnectorsInfo gets the dataServicesConnectorsInfo property value. Data Services Connectors specific information.
// returns a []V1beta1HypervisorManagersItemHypervisorGetResponse_dataServicesConnectorsInfoable when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetDataServicesConnectorsInfo()([]V1beta1HypervisorManagersItemHypervisorGetResponse_dataServicesConnectorsInfoable) {
    return m.dataServicesConnectorsInfo
}
// GetDescription gets the description property value. A brief description of the hypervisor manager.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. User defined name for the hypervisor manager.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemHypervisorGetResponse_appInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppInfo(val.(V1beta1HypervisorManagersItemHypervisorGetResponse_appInfoable))
        }
        return nil
    }
    res["buildVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildVersion(val)
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
    res["dataOrchestratorInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataOrchestratorInfo(val.(V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoable))
        }
        return nil
    }
    res["dataServicesConnectorsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorManagersItemHypervisorGetResponse_dataServicesConnectorsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorManagersItemHypervisorGetResponse_dataServicesConnectorsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorManagersItemHypervisorGetResponse_dataServicesConnectorsInfoable)
                }
            }
            m.SetDataServicesConnectorsInfo(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["lastRefreshed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshed(val)
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
    res["networkAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkAddress(val)
        }
        return nil
    }
    res["releaseVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseVersion(val)
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
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetServices(res)
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
    res["uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
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
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetId gets the id property value. UUID string uniquely identifying the hypervisor manager.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetId()(*string) {
    return m.id
}
// GetLastRefreshed gets the lastRefreshed property value. Time in UTC at which the object was last refreshed.
// returns a *Time when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetLastRefreshed()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshed
}
// GetName gets the name property value. Name as reported by the hypervisor manager.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetName()(*string) {
    return m.name
}
// GetNetworkAddress gets the networkAddress property value. An IP address or hostname or FQDN to address the hypervisor manager
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetNetworkAddress()(*string) {
    return m.networkAddress
}
// GetReleaseVersion gets the releaseVersion property value. The hypervisor manager release version.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetReleaseVersion()(*string) {
    return m.releaseVersion
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
// returns a []string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetServices()([]string) {
    return m.services
}
// GetStateReason gets the stateReason property value. Brief reason for the current state of the hypervisor manager
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetStateReason()(*string) {
    return m.stateReason
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUid gets the uid property value. A hypervisor manager provided durable UID. In case of VMware it will be instanceUUID of the vCenter
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetUid()(*string) {
    return m.uid
}
// GetUpdatedAt gets the updatedAt property value. Time in UTC at which the object was last updated
// returns a *Time when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetUsername gets the username property value. Name of the user used to access the hypervisor server. Mutually exclusive with credentialInfo.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appInfo", m.GetAppInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("buildVersion", m.GetBuildVersion())
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
        err := writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dataOrchestratorInfo", m.GetDataOrchestratorInfo())
        if err != nil {
            return err
        }
    }
    if m.GetDataServicesConnectorsInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDataServicesConnectorsInfo()))
        for i, v := range m.GetDataServicesConnectorsInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("dataServicesConnectorsInfo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastRefreshed", m.GetLastRefreshed())
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
        err := writer.WriteStringValue("networkAddress", m.GetNetworkAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("releaseVersion", m.GetReleaseVersion())
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
    if m.GetServices() != nil {
        err := writer.WriteCollectionOfStringValues("services", m.GetServices())
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
        err := writer.WriteStringValue("uid", m.GetUid())
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
        err := writer.WriteStringValue("username", m.GetUsername())
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
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppInfo sets the appInfo property value. Hypervisor specific information.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetAppInfo(value V1beta1HypervisorManagersItemHypervisorGetResponse_appInfoable)() {
    m.appInfo = value
}
// SetBuildVersion sets the buildVersion property value. The hypervisor manager build details
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetBuildVersion(value *string)() {
    m.buildVersion = value
}
// SetCreatedAt sets the createdAt property value. Time in UTC at which the object was created
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDataOrchestratorInfo sets the dataOrchestratorInfo property value. Data Orchestrator specific information.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetDataOrchestratorInfo(value V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoable)() {
    m.dataOrchestratorInfo = value
}
// SetDataServicesConnectorsInfo sets the dataServicesConnectorsInfo property value. Data Services Connectors specific information.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetDataServicesConnectorsInfo(value []V1beta1HypervisorManagersItemHypervisorGetResponse_dataServicesConnectorsInfoable)() {
    m.dataServicesConnectorsInfo = value
}
// SetDescription sets the description property value. A brief description of the hypervisor manager.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. User defined name for the hypervisor manager.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetId sets the id property value. UUID string uniquely identifying the hypervisor manager.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetId(value *string)() {
    m.id = value
}
// SetLastRefreshed sets the lastRefreshed property value. Time in UTC at which the object was last refreshed.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetLastRefreshed(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshed = value
}
// SetName sets the name property value. Name as reported by the hypervisor manager.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetName(value *string)() {
    m.name = value
}
// SetNetworkAddress sets the networkAddress property value. An IP address or hostname or FQDN to address the hypervisor manager
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetNetworkAddress(value *string)() {
    m.networkAddress = value
}
// SetReleaseVersion sets the releaseVersion property value. The hypervisor manager release version.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetReleaseVersion(value *string)() {
    m.releaseVersion = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetServices(value []string)() {
    m.services = value
}
// SetStateReason sets the stateReason property value. Brief reason for the current state of the hypervisor manager
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetStateReason(value *string)() {
    m.stateReason = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUid sets the uid property value. A hypervisor manager provided durable UID. In case of VMware it will be instanceUUID of the vCenter
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetUid(value *string)() {
    m.uid = value
}
// SetUpdatedAt sets the updatedAt property value. Time in UTC at which the object was last updated
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetUsername sets the username property value. Name of the user used to access the hypervisor server. Mutually exclusive with credentialInfo.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse) SetUsername(value *string)() {
    m.username = value
}
type V1beta1HypervisorManagersItemHypervisorGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppInfo()(V1beta1HypervisorManagersItemHypervisorGetResponse_appInfoable)
    GetBuildVersion()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDataOrchestratorInfo()(V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoable)
    GetDataServicesConnectorsInfo()([]V1beta1HypervisorManagersItemHypervisorGetResponse_dataServicesConnectorsInfoable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetGeneration()(*int64)
    GetId()(*string)
    GetLastRefreshed()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetNetworkAddress()(*string)
    GetReleaseVersion()(*string)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetStateReason()(*string)
    GetTypeEscaped()(*string)
    GetUid()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUsername()(*string)
    SetAppInfo(value V1beta1HypervisorManagersItemHypervisorGetResponse_appInfoable)()
    SetBuildVersion(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDataOrchestratorInfo(value V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoable)()
    SetDataServicesConnectorsInfo(value []V1beta1HypervisorManagersItemHypervisorGetResponse_dataServicesConnectorsInfoable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetGeneration(value *int64)()
    SetId(value *string)()
    SetLastRefreshed(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetNetworkAddress(value *string)()
    SetReleaseVersion(value *string)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetStateReason(value *string)()
    SetTypeEscaped(value *string)()
    SetUid(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUsername(value *string)()
}
