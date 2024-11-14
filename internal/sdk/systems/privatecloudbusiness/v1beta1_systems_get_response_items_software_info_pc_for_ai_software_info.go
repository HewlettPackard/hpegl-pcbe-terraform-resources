package privatecloudbusiness

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo software details of the Private Cloud AI. Applicable only for Private Cloud AI systems.
type V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of software catalogs available for update on the Private Cloud AI.
    availableSoftwareCatalogs []V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_availableSoftwareCatalogsable
    // Current Version of the software catalog and its constituent software components. Applicable only for Private Cloud AI systems.
    currentSoftwareCatalog V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogable
    // Details of the current software update status
    currentUpdateStatus V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusable
    // Date on which the last software precheck was run.
    lastPrecheckRunOn *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Version of the software catalog to which last software precheck was run.
    lastPrecheckRunVersion *string
    // Date on which the last software update was performed on this Private Cloud AI.
    lastUpdatedOn *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of software catalogs that are unavailable for update on the Private Cloud AI.
    unavailableSoftwareCatalogs []V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsable
}
// NewV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo instantiates a new V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo and sets the default values.
func NewV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo()(*V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) {
    m := &V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableSoftwareCatalogs gets the availableSoftwareCatalogs property value. List of software catalogs available for update on the Private Cloud AI.
// returns a []V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_availableSoftwareCatalogsable when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) GetAvailableSoftwareCatalogs()([]V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_availableSoftwareCatalogsable) {
    return m.availableSoftwareCatalogs
}
// GetCurrentSoftwareCatalog gets the currentSoftwareCatalog property value. Current Version of the software catalog and its constituent software components. Applicable only for Private Cloud AI systems.
// returns a V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogable when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) GetCurrentSoftwareCatalog()(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogable) {
    return m.currentSoftwareCatalog
}
// GetCurrentUpdateStatus gets the currentUpdateStatus property value. Details of the current software update status
// returns a V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusable when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) GetCurrentUpdateStatus()(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusable) {
    return m.currentUpdateStatus
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availableSoftwareCatalogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_availableSoftwareCatalogsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_availableSoftwareCatalogsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_availableSoftwareCatalogsable)
                }
            }
            m.SetAvailableSoftwareCatalogs(res)
        }
        return nil
    }
    res["currentSoftwareCatalog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentSoftwareCatalog(val.(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogable))
        }
        return nil
    }
    res["currentUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUpdateStatus(val.(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusable))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsable)
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
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) GetLastPrecheckRunOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastPrecheckRunOn
}
// GetLastPrecheckRunVersion gets the lastPrecheckRunVersion property value. Version of the software catalog to which last software precheck was run.
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) GetLastPrecheckRunVersion()(*string) {
    return m.lastPrecheckRunVersion
}
// GetLastUpdatedOn gets the lastUpdatedOn property value. Date on which the last software update was performed on this Private Cloud AI.
// returns a *Time when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) GetLastUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedOn
}
// GetUnavailableSoftwareCatalogs gets the unavailableSoftwareCatalogs property value. List of software catalogs that are unavailable for update on the Private Cloud AI.
// returns a []V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsable when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) GetUnavailableSoftwareCatalogs()([]V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsable) {
    return m.unavailableSoftwareCatalogs
}
// Serialize serializes information the current object
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableSoftwareCatalogs sets the availableSoftwareCatalogs property value. List of software catalogs available for update on the Private Cloud AI.
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) SetAvailableSoftwareCatalogs(value []V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_availableSoftwareCatalogsable)() {
    m.availableSoftwareCatalogs = value
}
// SetCurrentSoftwareCatalog sets the currentSoftwareCatalog property value. Current Version of the software catalog and its constituent software components. Applicable only for Private Cloud AI systems.
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) SetCurrentSoftwareCatalog(value V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogable)() {
    m.currentSoftwareCatalog = value
}
// SetCurrentUpdateStatus sets the currentUpdateStatus property value. Details of the current software update status
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) SetCurrentUpdateStatus(value V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusable)() {
    m.currentUpdateStatus = value
}
// SetLastPrecheckRunOn sets the lastPrecheckRunOn property value. Date on which the last software precheck was run.
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) SetLastPrecheckRunOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastPrecheckRunOn = value
}
// SetLastPrecheckRunVersion sets the lastPrecheckRunVersion property value. Version of the software catalog to which last software precheck was run.
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) SetLastPrecheckRunVersion(value *string)() {
    m.lastPrecheckRunVersion = value
}
// SetLastUpdatedOn sets the lastUpdatedOn property value. Date on which the last software update was performed on this Private Cloud AI.
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) SetLastUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedOn = value
}
// SetUnavailableSoftwareCatalogs sets the unavailableSoftwareCatalogs property value. List of software catalogs that are unavailable for update on the Private Cloud AI.
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo) SetUnavailableSoftwareCatalogs(value []V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsable)() {
    m.unavailableSoftwareCatalogs = value
}
type V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableSoftwareCatalogs()([]V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_availableSoftwareCatalogsable)
    GetCurrentSoftwareCatalog()(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogable)
    GetCurrentUpdateStatus()(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusable)
    GetLastPrecheckRunOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastPrecheckRunVersion()(*string)
    GetLastUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUnavailableSoftwareCatalogs()([]V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsable)
    SetAvailableSoftwareCatalogs(value []V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_availableSoftwareCatalogsable)()
    SetCurrentSoftwareCatalog(value V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalogable)()
    SetCurrentUpdateStatus(value V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_currentUpdateStatusable)()
    SetLastPrecheckRunOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastPrecheckRunVersion(value *string)()
    SetLastUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUnavailableSoftwareCatalogs(value []V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsable)()
}
