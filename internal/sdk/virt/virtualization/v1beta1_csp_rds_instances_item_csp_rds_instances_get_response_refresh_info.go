package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo property definitions for Refresh inventory Assets
type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Time of the most recent attempt to refresh the inventory with the cloud service provider
    lastRefreshedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo instantiates a new V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo and sets the default values.
func NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo()(*V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo) {
    m := &V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lastRefreshedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedAt(val)
        }
        return nil
    }
    return res
}
// GetLastRefreshedAt gets the lastRefreshedAt property value. Time of the most recent attempt to refresh the inventory with the cloud service provider
// returns a *Time when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo) GetLastRefreshedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshedAt
}
// Serialize serializes information the current object
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastRefreshedAt", m.GetLastRefreshedAt())
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
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLastRefreshedAt sets the lastRefreshedAt property value. Time of the most recent attempt to refresh the inventory with the cloud service provider
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfo) SetLastRefreshedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedAt = value
}
type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_refreshInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastRefreshedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetLastRefreshedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
