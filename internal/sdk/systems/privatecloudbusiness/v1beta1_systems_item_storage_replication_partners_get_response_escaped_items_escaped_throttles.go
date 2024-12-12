package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The days property
    days *float64
    // The thrAtTime property
    thrAtTime *float64
    // The thrBandwidthLimitKbps property
    thrBandwidthLimitKbps *float64
    // The thrUntilTime property
    thrUntilTime *float64
}
// NewV1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles instantiates a new V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles and sets the default values.
func NewV1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles()(*V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) {
    m := &V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDays gets the days property value. The days property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) GetDays()(*float64) {
    return m.days
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["days"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDays(val)
        }
        return nil
    }
    res["thrAtTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThrAtTime(val)
        }
        return nil
    }
    res["thrBandwidthLimitKbps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThrBandwidthLimitKbps(val)
        }
        return nil
    }
    res["thrUntilTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThrUntilTime(val)
        }
        return nil
    }
    return res
}
// GetThrAtTime gets the thrAtTime property value. The thrAtTime property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) GetThrAtTime()(*float64) {
    return m.thrAtTime
}
// GetThrBandwidthLimitKbps gets the thrBandwidthLimitKbps property value. The thrBandwidthLimitKbps property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) GetThrBandwidthLimitKbps()(*float64) {
    return m.thrBandwidthLimitKbps
}
// GetThrUntilTime gets the thrUntilTime property value. The thrUntilTime property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) GetThrUntilTime()(*float64) {
    return m.thrUntilTime
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("days", m.GetDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("thrAtTime", m.GetThrAtTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("thrBandwidthLimitKbps", m.GetThrBandwidthLimitKbps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("thrUntilTime", m.GetThrUntilTime())
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
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDays sets the days property value. The days property
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) SetDays(value *float64)() {
    m.days = value
}
// SetThrAtTime sets the thrAtTime property value. The thrAtTime property
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) SetThrAtTime(value *float64)() {
    m.thrAtTime = value
}
// SetThrBandwidthLimitKbps sets the thrBandwidthLimitKbps property value. The thrBandwidthLimitKbps property
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) SetThrBandwidthLimitKbps(value *float64)() {
    m.thrBandwidthLimitKbps = value
}
// SetThrUntilTime sets the thrUntilTime property value. The thrUntilTime property
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttles) SetThrUntilTime(value *float64)() {
    m.thrUntilTime = value
}
type V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDays()(*float64)
    GetThrAtTime()(*float64)
    GetThrBandwidthLimitKbps()(*float64)
    GetThrUntilTime()(*float64)
    SetDays(value *float64)()
    SetThrAtTime(value *float64)()
    SetThrBandwidthLimitKbps(value *float64)()
    SetThrUntilTime(value *float64)()
}
