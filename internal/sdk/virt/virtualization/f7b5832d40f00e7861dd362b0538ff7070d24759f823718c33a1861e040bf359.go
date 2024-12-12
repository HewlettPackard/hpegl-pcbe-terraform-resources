package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo regions in which the Azure VM size is deployed.
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A region in which the Azure VM size is deployed.
    location *string
    // Details of the zone.
    zoneDetails []string
    // A list of availability zones denoting where the resource needs to come from.
    zones []string
}
// NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo instantiates a new V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo and sets the default values.
func NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo()(*V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) {
    m := &V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["zoneDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetZoneDetails(res)
        }
        return nil
    }
    res["zones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetZones(res)
        }
        return nil
    }
    return res
}
// GetLocation gets the location property value. A region in which the Azure VM size is deployed.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) GetLocation()(*string) {
    return m.location
}
// GetZoneDetails gets the zoneDetails property value. Details of the zone.
// returns a []string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) GetZoneDetails()([]string) {
    return m.zoneDetails
}
// GetZones gets the zones property value. A list of availability zones denoting where the resource needs to come from.
// returns a []string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) GetZones()([]string) {
    return m.zones
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    if m.GetZoneDetails() != nil {
        err := writer.WriteCollectionOfStringValues("zoneDetails", m.GetZoneDetails())
        if err != nil {
            return err
        }
    }
    if m.GetZones() != nil {
        err := writer.WriteCollectionOfStringValues("zones", m.GetZones())
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
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocation sets the location property value. A region in which the Azure VM size is deployed.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) SetLocation(value *string)() {
    m.location = value
}
// SetZoneDetails sets the zoneDetails property value. Details of the zone.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) SetZoneDetails(value []string)() {
    m.zoneDetails = value
}
// SetZones sets the zones property value. A list of availability zones denoting where the resource needs to come from.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfo) SetZones(value []string)() {
    m.zones = value
}
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_locationInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocation()(*string)
    GetZoneDetails()([]string)
    GetZones()([]string)
    SetLocation(value *string)()
    SetZoneDetails(value []string)()
    SetZones(value []string)()
}
