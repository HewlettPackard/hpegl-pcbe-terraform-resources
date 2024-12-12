package privatecloudbusiness

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsGetResponse_items_softwareInfo system software information.
type V1beta1SystemsGetResponse_items_softwareInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Details of the current software update status
    currentUpdateStatus V1beta1SystemsGetResponse_items_softwareInfo_currentUpdateStatusable
    // List of hypervisor clusters in the system with their software details.
    hypervisorClusters []V1beta1SystemsGetResponse_items_softwareInfo_hypervisorClustersable
    // Date on which the last software update was performed on this system.
    lastUpdatedOn *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Software details of the Private Cloud AI. Applicable only for Private Cloud AI systems.
    pcForAiSoftwareInfo V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoable
    // If a software update precheck is completed recently, time until which that precheck is valid for software update to be initiated.
    precheckValidUntil *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1SystemsGetResponse_items_softwareInfo instantiates a new V1beta1SystemsGetResponse_items_softwareInfo and sets the default values.
func NewV1beta1SystemsGetResponse_items_softwareInfo()(*V1beta1SystemsGetResponse_items_softwareInfo) {
    m := &V1beta1SystemsGetResponse_items_softwareInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsGetResponse_items_softwareInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsGetResponse_items_softwareInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsGetResponse_items_softwareInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentUpdateStatus gets the currentUpdateStatus property value. Details of the current software update status
// returns a V1beta1SystemsGetResponse_items_softwareInfo_currentUpdateStatusable when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo) GetCurrentUpdateStatus()(V1beta1SystemsGetResponse_items_softwareInfo_currentUpdateStatusable) {
    return m.currentUpdateStatus
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsGetResponse_items_softwareInfo_currentUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUpdateStatus(val.(V1beta1SystemsGetResponse_items_softwareInfo_currentUpdateStatusable))
        }
        return nil
    }
    res["hypervisorClusters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsGetResponse_items_softwareInfo_hypervisorClustersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsGetResponse_items_softwareInfo_hypervisorClustersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsGetResponse_items_softwareInfo_hypervisorClustersable)
                }
            }
            m.SetHypervisorClusters(res)
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
    res["pcForAiSoftwareInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPcForAiSoftwareInfo(val.(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoable))
        }
        return nil
    }
    res["precheckValidUntil"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrecheckValidUntil(val)
        }
        return nil
    }
    return res
}
// GetHypervisorClusters gets the hypervisorClusters property value. List of hypervisor clusters in the system with their software details.
// returns a []V1beta1SystemsGetResponse_items_softwareInfo_hypervisorClustersable when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo) GetHypervisorClusters()([]V1beta1SystemsGetResponse_items_softwareInfo_hypervisorClustersable) {
    return m.hypervisorClusters
}
// GetLastUpdatedOn gets the lastUpdatedOn property value. Date on which the last software update was performed on this system.
// returns a *Time when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo) GetLastUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedOn
}
// GetPcForAiSoftwareInfo gets the pcForAiSoftwareInfo property value. Software details of the Private Cloud AI. Applicable only for Private Cloud AI systems.
// returns a V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoable when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo) GetPcForAiSoftwareInfo()(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoable) {
    return m.pcForAiSoftwareInfo
}
// GetPrecheckValidUntil gets the precheckValidUntil property value. If a software update precheck is completed recently, time until which that precheck is valid for software update to be initiated.
// returns a *Time when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo) GetPrecheckValidUntil()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.precheckValidUntil
}
// Serialize serializes information the current object
func (m *V1beta1SystemsGetResponse_items_softwareInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("currentUpdateStatus", m.GetCurrentUpdateStatus())
        if err != nil {
            return err
        }
    }
    if m.GetHypervisorClusters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHypervisorClusters()))
        for i, v := range m.GetHypervisorClusters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("hypervisorClusters", cast)
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
    {
        err := writer.WriteObjectValue("pcForAiSoftwareInfo", m.GetPcForAiSoftwareInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("precheckValidUntil", m.GetPrecheckValidUntil())
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
func (m *V1beta1SystemsGetResponse_items_softwareInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentUpdateStatus sets the currentUpdateStatus property value. Details of the current software update status
func (m *V1beta1SystemsGetResponse_items_softwareInfo) SetCurrentUpdateStatus(value V1beta1SystemsGetResponse_items_softwareInfo_currentUpdateStatusable)() {
    m.currentUpdateStatus = value
}
// SetHypervisorClusters sets the hypervisorClusters property value. List of hypervisor clusters in the system with their software details.
func (m *V1beta1SystemsGetResponse_items_softwareInfo) SetHypervisorClusters(value []V1beta1SystemsGetResponse_items_softwareInfo_hypervisorClustersable)() {
    m.hypervisorClusters = value
}
// SetLastUpdatedOn sets the lastUpdatedOn property value. Date on which the last software update was performed on this system.
func (m *V1beta1SystemsGetResponse_items_softwareInfo) SetLastUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedOn = value
}
// SetPcForAiSoftwareInfo sets the pcForAiSoftwareInfo property value. Software details of the Private Cloud AI. Applicable only for Private Cloud AI systems.
func (m *V1beta1SystemsGetResponse_items_softwareInfo) SetPcForAiSoftwareInfo(value V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoable)() {
    m.pcForAiSoftwareInfo = value
}
// SetPrecheckValidUntil sets the precheckValidUntil property value. If a software update precheck is completed recently, time until which that precheck is valid for software update to be initiated.
func (m *V1beta1SystemsGetResponse_items_softwareInfo) SetPrecheckValidUntil(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.precheckValidUntil = value
}
type V1beta1SystemsGetResponse_items_softwareInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentUpdateStatus()(V1beta1SystemsGetResponse_items_softwareInfo_currentUpdateStatusable)
    GetHypervisorClusters()([]V1beta1SystemsGetResponse_items_softwareInfo_hypervisorClustersable)
    GetLastUpdatedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPcForAiSoftwareInfo()(V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoable)
    GetPrecheckValidUntil()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCurrentUpdateStatus(value V1beta1SystemsGetResponse_items_softwareInfo_currentUpdateStatusable)()
    SetHypervisorClusters(value []V1beta1SystemsGetResponse_items_softwareInfo_hypervisorClustersable)()
    SetLastUpdatedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPcForAiSoftwareInfo(value V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfoable)()
    SetPrecheckValidUntil(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
