package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemStorageReplicationPartnersGetResponse_items system's storage system replication partner info.
type V1beta1SystemsItemStorageReplicationPartnersGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Remote replication partner hostname.
    hostname *string
    // Replication partner availability.
    isAlive *bool
    // Type of replication partner.
    partnerType *string
    // Indicates whether replication traffic from/to this partner has been halted.
    paused *bool
    // Direction of replication configured with this partner.Possible values: none, downstream, upstream, bi_directional.
    replicationDirection *string
    // Throttles used while replicating from/to this partner.
    throttles []V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesable
    // Number of volume collections that are replicating from/to this partner.
    volumeCollectionCount *float64
    // Number of volumes that are replicating from/to this partner.
    volumeCount *float64
}
// NewV1beta1SystemsItemStorageReplicationPartnersGetResponse_items instantiates a new V1beta1SystemsItemStorageReplicationPartnersGetResponse_items and sets the default values.
func NewV1beta1SystemsItemStorageReplicationPartnersGetResponse_items()(*V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) {
    m := &V1beta1SystemsItemStorageReplicationPartnersGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemStorageReplicationPartnersGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemStorageReplicationPartnersGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemStorageReplicationPartnersGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostname(val)
        }
        return nil
    }
    res["isAlive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAlive(val)
        }
        return nil
    }
    res["partnerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerType(val)
        }
        return nil
    }
    res["paused"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaused(val)
        }
        return nil
    }
    res["replicationDirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplicationDirection(val)
        }
        return nil
    }
    res["throttles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesable)
                }
            }
            m.SetThrottles(res)
        }
        return nil
    }
    res["volumeCollectionCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeCollectionCount(val)
        }
        return nil
    }
    res["volumeCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeCount(val)
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. Remote replication partner hostname.
// returns a *string when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetHostname()(*string) {
    return m.hostname
}
// GetIsAlive gets the isAlive property value. Replication partner availability.
// returns a *bool when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetIsAlive()(*bool) {
    return m.isAlive
}
// GetPartnerType gets the partnerType property value. Type of replication partner.
// returns a *string when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetPartnerType()(*string) {
    return m.partnerType
}
// GetPaused gets the paused property value. Indicates whether replication traffic from/to this partner has been halted.
// returns a *bool when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetPaused()(*bool) {
    return m.paused
}
// GetReplicationDirection gets the replicationDirection property value. Direction of replication configured with this partner.Possible values: none, downstream, upstream, bi_directional.
// returns a *string when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetReplicationDirection()(*string) {
    return m.replicationDirection
}
// GetThrottles gets the throttles property value. Throttles used while replicating from/to this partner.
// returns a []V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesable when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetThrottles()([]V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesable) {
    return m.throttles
}
// GetVolumeCollectionCount gets the volumeCollectionCount property value. Number of volume collections that are replicating from/to this partner.
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetVolumeCollectionCount()(*float64) {
    return m.volumeCollectionCount
}
// GetVolumeCount gets the volumeCount property value. Number of volumes that are replicating from/to this partner.
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) GetVolumeCount()(*float64) {
    return m.volumeCount
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAlive", m.GetIsAlive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("partnerType", m.GetPartnerType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("paused", m.GetPaused())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("replicationDirection", m.GetReplicationDirection())
        if err != nil {
            return err
        }
    }
    if m.GetThrottles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetThrottles()))
        for i, v := range m.GetThrottles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("throttles", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("volumeCollectionCount", m.GetVolumeCollectionCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("volumeCount", m.GetVolumeCount())
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
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHostname sets the hostname property value. Remote replication partner hostname.
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) SetHostname(value *string)() {
    m.hostname = value
}
// SetIsAlive sets the isAlive property value. Replication partner availability.
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) SetIsAlive(value *bool)() {
    m.isAlive = value
}
// SetPartnerType sets the partnerType property value. Type of replication partner.
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) SetPartnerType(value *string)() {
    m.partnerType = value
}
// SetPaused sets the paused property value. Indicates whether replication traffic from/to this partner has been halted.
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) SetPaused(value *bool)() {
    m.paused = value
}
// SetReplicationDirection sets the replicationDirection property value. Direction of replication configured with this partner.Possible values: none, downstream, upstream, bi_directional.
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) SetReplicationDirection(value *string)() {
    m.replicationDirection = value
}
// SetThrottles sets the throttles property value. Throttles used while replicating from/to this partner.
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) SetThrottles(value []V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesable)() {
    m.throttles = value
}
// SetVolumeCollectionCount sets the volumeCollectionCount property value. Number of volume collections that are replicating from/to this partner.
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) SetVolumeCollectionCount(value *float64)() {
    m.volumeCollectionCount = value
}
// SetVolumeCount sets the volumeCount property value. Number of volumes that are replicating from/to this partner.
func (m *V1beta1SystemsItemStorageReplicationPartnersGetResponse_items) SetVolumeCount(value *float64)() {
    m.volumeCount = value
}
type V1beta1SystemsItemStorageReplicationPartnersGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHostname()(*string)
    GetIsAlive()(*bool)
    GetPartnerType()(*string)
    GetPaused()(*bool)
    GetReplicationDirection()(*string)
    GetThrottles()([]V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesable)
    GetVolumeCollectionCount()(*float64)
    GetVolumeCount()(*float64)
    SetHostname(value *string)()
    SetIsAlive(value *bool)()
    SetPartnerType(value *string)()
    SetPaused(value *bool)()
    SetReplicationDirection(value *string)()
    SetThrottles(value []V1beta1SystemsItemStorageReplicationPartnersGetResponse_items_throttlesable)()
    SetVolumeCollectionCount(value *float64)()
    SetVolumeCount(value *float64)()
}
