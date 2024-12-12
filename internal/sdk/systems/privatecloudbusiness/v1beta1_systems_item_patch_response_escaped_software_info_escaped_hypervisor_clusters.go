package privatecloudbusiness

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters hypervisor cluster with details of each software component.
type V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of software catalogs available for update on the hypervisor cluster.
    availableSoftwareCatalogs []V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsable
    // Software Catalog with detailed versions of various software components like HPE Storage Software, hypervisor software, HPE Storage Connection Manager and Server firmware. If the version of this catalog is set to 'Unavailable', it means one or more software component versions are not available. If the version of this catalog is set to 'Non-Compliant', it means the current set of software component versions does not match any of the supported software catalog versions.
    currentSoftwareCatalog V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogable
    // Details of the current software update status
    currentUpdateStatus V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatusable
    // Date on which the last software precheck was run.
    lastPrecheckRunOn *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Version of the software catalog to which last software precheck was run.
    lastPrecheckRunVersion *string
    // Date on which the last software update was performed on this hypervisor cluster.
    lastUpdatedOn *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of software catalogs that are unavailable for update on the hypervisor cluster.
    unavailableSoftwareCatalogs []V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_unavailableSoftwareCatalogsable
}
// NewV1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters instantiates a new V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters and sets the default values.
func NewV1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters()(*V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) {
    m := &V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClustersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClustersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableSoftwareCatalogs gets the availableSoftwareCatalogs property value. List of software catalogs available for update on the hypervisor cluster.
// returns a []V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) GetAvailableSoftwareCatalogs()([]V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsable) {
    return m.availableSoftwareCatalogs
}
// GetCurrentSoftwareCatalog gets the currentSoftwareCatalog property value. Software Catalog with detailed versions of various software components like HPE Storage Software, hypervisor software, HPE Storage Connection Manager and Server firmware. If the version of this catalog is set to 'Unavailable', it means one or more software component versions are not available. If the version of this catalog is set to 'Non-Compliant', it means the current set of software component versions does not match any of the supported software catalog versions.
// returns a V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) GetCurrentSoftwareCatalog()(V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogable) {
    return m.currentSoftwareCatalog
}
// GetCurrentUpdateStatus gets the currentUpdateStatus property value. Details of the current software update status
// returns a V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatusable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) GetCurrentUpdateStatus()(V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatusable) {
    return m.currentUpdateStatus
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availableSoftwareCatalogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsable)
                }
            }
            m.SetAvailableSoftwareCatalogs(res)
        }
        return nil
    }
    res["currentSoftwareCatalog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentSoftwareCatalog(val.(V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogable))
        }
        return nil
    }
    res["currentUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUpdateStatus(val.(V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatusable))
        }
        return nil
    }
    res["lastPrecheckRunOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastPrecheckRunOn(val)
        }
        return nil
    }
    res["lastPrecheckRunVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastPrecheckRunVersion(val)
        }
        return nil
    }
    res["lastUpdatedOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedOn(val)
        }
        return nil
    }
    res["unavailableSoftwareCatalogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_unavailableSoftwareCatalogsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_unavailableSoftwareCatalogsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_unavailableSoftwareCatalogsable)
                }
            }
            m.SetUnavailableSoftwareCatalogs(res)
        }
        return nil
    }
    return res
}
// GetLastPrecheckRunOn gets the lastPrecheckRunOn property value. Date on which the last software precheck was run.
// returns a *Time when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) GetLastPrecheckRunOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastPrecheckRunOn
}
// GetLastPrecheckRunVersion gets the lastPrecheckRunVersion property value. Version of the software catalog to which last software precheck was run.
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) GetLastPrecheckRunVersion()(*string) {
    return m.lastPrecheckRunVersion
}
// GetLastUpdatedOn gets the lastUpdatedOn property value. Date on which the last software update was performed on this hypervisor cluster.
// returns a *Time when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) GetLastUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedOn
}
// GetUnavailableSoftwareCatalogs gets the unavailableSoftwareCatalogs property value. List of software catalogs that are unavailable for update on the hypervisor cluster.
// returns a []V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_unavailableSoftwareCatalogsable when successful
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) GetUnavailableSoftwareCatalogs()([]V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_unavailableSoftwareCatalogsable) {
    return m.unavailableSoftwareCatalogs
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAvailableSoftwareCatalogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAvailableSoftwareCatalogs()))
        for i, v := range m.GetAvailableSoftwareCatalogs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("availableSoftwareCatalogs", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("currentSoftwareCatalog", m.GetCurrentSoftwareCatalog())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("currentUpdateStatus", m.GetCurrentUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastPrecheckRunOn", m.GetLastPrecheckRunOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastPrecheckRunVersion", m.GetLastPrecheckRunVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUpdatedOn", m.GetLastUpdatedOn())
        if err != nil {
            return err
        }
    }
    if m.GetUnavailableSoftwareCatalogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUnavailableSoftwareCatalogs()))
        for i, v := range m.GetUnavailableSoftwareCatalogs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("unavailableSoftwareCatalogs", cast)
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
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableSoftwareCatalogs sets the availableSoftwareCatalogs property value. List of software catalogs available for update on the hypervisor cluster.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) SetAvailableSoftwareCatalogs(value []V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsable)() {
    m.availableSoftwareCatalogs = value
}
// SetCurrentSoftwareCatalog sets the currentSoftwareCatalog property value. Software Catalog with detailed versions of various software components like HPE Storage Software, hypervisor software, HPE Storage Connection Manager and Server firmware. If the version of this catalog is set to 'Unavailable', it means one or more software component versions are not available. If the version of this catalog is set to 'Non-Compliant', it means the current set of software component versions does not match any of the supported software catalog versions.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) SetCurrentSoftwareCatalog(value V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogable)() {
    m.currentSoftwareCatalog = value
}
// SetCurrentUpdateStatus sets the currentUpdateStatus property value. Details of the current software update status
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) SetCurrentUpdateStatus(value V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatusable)() {
    m.currentUpdateStatus = value
}
// SetLastPrecheckRunOn sets the lastPrecheckRunOn property value. Date on which the last software precheck was run.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) SetLastPrecheckRunOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastPrecheckRunOn = value
}
// SetLastPrecheckRunVersion sets the lastPrecheckRunVersion property value. Version of the software catalog to which last software precheck was run.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) SetLastPrecheckRunVersion(value *string)() {
    m.lastPrecheckRunVersion = value
}
// SetLastUpdatedOn sets the lastUpdatedOn property value. Date on which the last software update was performed on this hypervisor cluster.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) SetLastUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedOn = value
}
// SetUnavailableSoftwareCatalogs sets the unavailableSoftwareCatalogs property value. List of software catalogs that are unavailable for update on the hypervisor cluster.
func (m *V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters) SetUnavailableSoftwareCatalogs(value []V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_unavailableSoftwareCatalogsable)() {
    m.unavailableSoftwareCatalogs = value
}
type V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClustersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableSoftwareCatalogs()([]V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsable)
    GetCurrentSoftwareCatalog()(V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogable)
    GetCurrentUpdateStatus()(V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatusable)
    GetLastPrecheckRunOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastPrecheckRunVersion()(*string)
    GetLastUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUnavailableSoftwareCatalogs()([]V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_unavailableSoftwareCatalogsable)
    SetAvailableSoftwareCatalogs(value []V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsable)()
    SetCurrentSoftwareCatalog(value V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogable)()
    SetCurrentUpdateStatus(value V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatusable)()
    SetLastPrecheckRunOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastPrecheckRunVersion(value *string)()
    SetLastUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUnavailableSoftwareCatalogs(value []V1beta1SystemsItemPatchResponse_softwareInfo_hypervisorClusters_unavailableSoftwareCatalogsable)()
}
