package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Capabilities that are enabled one the VM.
    capabilities []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_capabilitiesable
    // The family property
    family *string
    // The locationInfo property
    locationInfo []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_locationInfoable
    // A region in which the Azure VM size is deployed.
    locations []string
    // Name of the Azure VM size.
    name *string
    // Type of resource.
    resourceType *string
    // The restrictions because of which Vm size cannot be used.
    restrictions []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_restrictionsable
    // VM size name.
    size *string
    // Tier of the disk
    tier *string
}
// NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2 instantiates a new V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2 and sets the default values.
func NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2()(*V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) {
    m := &V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapabilities gets the capabilities property value. Capabilities that are enabled one the VM.
// returns a []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_capabilitiesable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetCapabilities()([]V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_capabilitiesable) {
    return m.capabilities
}
// GetFamily gets the family property value. The family property
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetFamily()(*string) {
    return m.family
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_capabilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_capabilitiesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_capabilitiesable)
                }
            }
            m.SetCapabilities(res)
        }
        return nil
    }
    res["family"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFamily(val)
        }
        return nil
    }
    res["locationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_locationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_locationInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_locationInfoable)
                }
            }
            m.SetLocationInfo(res)
        }
        return nil
    }
    res["locations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetLocations(res)
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
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val)
        }
        return nil
    }
    res["restrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_restrictionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_restrictionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_restrictionsable)
                }
            }
            m.SetRestrictions(res)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["tier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTier(val)
        }
        return nil
    }
    return res
}
// GetLocationInfo gets the locationInfo property value. The locationInfo property
// returns a []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_locationInfoable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetLocationInfo()([]V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_locationInfoable) {
    return m.locationInfo
}
// GetLocations gets the locations property value. A region in which the Azure VM size is deployed.
// returns a []string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetLocations()([]string) {
    return m.locations
}
// GetName gets the name property value. Name of the Azure VM size.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetName()(*string) {
    return m.name
}
// GetResourceType gets the resourceType property value. Type of resource.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetResourceType()(*string) {
    return m.resourceType
}
// GetRestrictions gets the restrictions property value. The restrictions because of which Vm size cannot be used.
// returns a []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_restrictionsable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetRestrictions()([]V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_restrictionsable) {
    return m.restrictions
}
// GetSize gets the size property value. VM size name.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetSize()(*string) {
    return m.size
}
// GetTier gets the tier property value. Tier of the disk
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) GetTier()(*string) {
    return m.tier
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCapabilities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCapabilities()))
        for i, v := range m.GetCapabilities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("capabilities", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("family", m.GetFamily())
        if err != nil {
            return err
        }
    }
    if m.GetLocationInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocationInfo()))
        for i, v := range m.GetLocationInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("locationInfo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocations() != nil {
        err := writer.WriteCollectionOfStringValues("locations", m.GetLocations())
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
        err := writer.WriteStringValue("resourceType", m.GetResourceType())
        if err != nil {
            return err
        }
    }
    if m.GetRestrictions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRestrictions()))
        for i, v := range m.GetRestrictions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("restrictions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tier", m.GetTier())
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
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapabilities sets the capabilities property value. Capabilities that are enabled one the VM.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetCapabilities(value []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_capabilitiesable)() {
    m.capabilities = value
}
// SetFamily sets the family property value. The family property
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetFamily(value *string)() {
    m.family = value
}
// SetLocationInfo sets the locationInfo property value. The locationInfo property
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetLocationInfo(value []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_locationInfoable)() {
    m.locationInfo = value
}
// SetLocations sets the locations property value. A region in which the Azure VM size is deployed.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetLocations(value []string)() {
    m.locations = value
}
// SetName sets the name property value. Name of the Azure VM size.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetName(value *string)() {
    m.name = value
}
// SetResourceType sets the resourceType property value. Type of resource.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetResourceType(value *string)() {
    m.resourceType = value
}
// SetRestrictions sets the restrictions property value. The restrictions because of which Vm size cannot be used.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetRestrictions(value []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_restrictionsable)() {
    m.restrictions = value
}
// SetSize sets the size property value. VM size name.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetSize(value *string)() {
    m.size = value
}
// SetTier sets the tier property value. Tier of the disk
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2) SetTier(value *string)() {
    m.tier = value
}
type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapabilities()([]V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_capabilitiesable)
    GetFamily()(*string)
    GetLocationInfo()([]V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_locationInfoable)
    GetLocations()([]string)
    GetName()(*string)
    GetResourceType()(*string)
    GetRestrictions()([]V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_restrictionsable)
    GetSize()(*string)
    GetTier()(*string)
    SetCapabilities(value []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_capabilitiesable)()
    SetFamily(value *string)()
    SetLocationInfo(value []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_locationInfoable)()
    SetLocations(value []string)()
    SetName(value *string)()
    SetResourceType(value *string)()
    SetRestrictions(value []V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2_restrictionsable)()
    SetSize(value *string)()
    SetTier(value *string)()
}
