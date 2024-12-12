package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineImagesGetResponse_items_cspInfoMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cloud machine image architecture (i386 | x86_64 | arm64 ).
    architecture *string
    // Automatic Os Upgrade Properties.
    automaticOsUpgradeProperties V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_automaticOsUpgradePropertiesable
    // The aztags property
    aztags *string
    // Images of data disk.
    dataDiskImages []string
    // The description of the cloud machine image.
    description *string
    // The disallowed property
    disallowed V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedable
    // The extended location of the Virtual Machine.
    extendedLocation *string
    // The features property
    features []V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_featuresable
    // The hypervisor generation of the Virtual Machine.
    hypervGeneration *string
    // Azure VM image identifier.
    id *string
    // Image deprecation status properties on the image.
    imageDeprecationStatusInfo V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_imageDeprecationStatusInfoable
    // Image location.
    location *string
    // Name of the image.
    name *string
    // The os disk image information.
    osDiskImage V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_osDiskImageable
    // The plan property
    plan *string
    // Image version info.
    versionInfo V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoable
}
// NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember2 instantiates a new V1beta1CspMachineImagesGetResponse_items_cspInfoMember2 and sets the default values.
func NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember2()(*V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) {
    m := &V1beta1CspMachineImagesGetResponse_items_cspInfoMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchitecture gets the architecture property value. The cloud machine image architecture (i386 | x86_64 | arm64 ).
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetArchitecture()(*string) {
    return m.architecture
}
// GetAutomaticOsUpgradeProperties gets the automaticOsUpgradeProperties property value. Automatic Os Upgrade Properties.
// returns a V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_automaticOsUpgradePropertiesable when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetAutomaticOsUpgradeProperties()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_automaticOsUpgradePropertiesable) {
    return m.automaticOsUpgradeProperties
}
// GetAztags gets the aztags property value. The aztags property
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetAztags()(*string) {
    return m.aztags
}
// GetDataDiskImages gets the dataDiskImages property value. Images of data disk.
// returns a []string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetDataDiskImages()([]string) {
    return m.dataDiskImages
}
// GetDescription gets the description property value. The description of the cloud machine image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetDescription()(*string) {
    return m.description
}
// GetDisallowed gets the disallowed property value. The disallowed property
// returns a V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedable when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetDisallowed()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedable) {
    return m.disallowed
}
// GetExtendedLocation gets the extendedLocation property value. The extended location of the Virtual Machine.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetExtendedLocation()(*string) {
    return m.extendedLocation
}
// GetFeatures gets the features property value. The features property
// returns a []V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_featuresable when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetFeatures()([]V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_featuresable) {
    return m.features
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["architecture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchitecture(val)
        }
        return nil
    }
    res["automaticOsUpgradeProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_automaticOsUpgradePropertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticOsUpgradeProperties(val.(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_automaticOsUpgradePropertiesable))
        }
        return nil
    }
    res["aztags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAztags(val)
        }
        return nil
    }
    res["dataDiskImages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDataDiskImages(res)
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
    res["disallowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisallowed(val.(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedable))
        }
        return nil
    }
    res["extendedLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtendedLocation(val)
        }
        return nil
    }
    res["features"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_featuresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_featuresable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_featuresable)
                }
            }
            m.SetFeatures(res)
        }
        return nil
    }
    res["hypervGeneration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervGeneration(val)
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
    res["imageDeprecationStatusInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_imageDeprecationStatusInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDeprecationStatusInfo(val.(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_imageDeprecationStatusInfoable))
        }
        return nil
    }
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
    res["osDiskImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_osDiskImageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDiskImage(val.(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_osDiskImageable))
        }
        return nil
    }
    res["plan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlan(val)
        }
        return nil
    }
    res["versionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionInfo(val.(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoable))
        }
        return nil
    }
    return res
}
// GetHypervGeneration gets the hypervGeneration property value. The hypervisor generation of the Virtual Machine.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetHypervGeneration()(*string) {
    return m.hypervGeneration
}
// GetId gets the id property value. Azure VM image identifier.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetId()(*string) {
    return m.id
}
// GetImageDeprecationStatusInfo gets the imageDeprecationStatusInfo property value. Image deprecation status properties on the image.
// returns a V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_imageDeprecationStatusInfoable when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetImageDeprecationStatusInfo()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_imageDeprecationStatusInfoable) {
    return m.imageDeprecationStatusInfo
}
// GetLocation gets the location property value. Image location.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetLocation()(*string) {
    return m.location
}
// GetName gets the name property value. Name of the image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetName()(*string) {
    return m.name
}
// GetOsDiskImage gets the osDiskImage property value. The os disk image information.
// returns a V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_osDiskImageable when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetOsDiskImage()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_osDiskImageable) {
    return m.osDiskImage
}
// GetPlan gets the plan property value. The plan property
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetPlan()(*string) {
    return m.plan
}
// GetVersionInfo gets the versionInfo property value. Image version info.
// returns a V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoable when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) GetVersionInfo()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoable) {
    return m.versionInfo
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("architecture", m.GetArchitecture())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("automaticOsUpgradeProperties", m.GetAutomaticOsUpgradeProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("aztags", m.GetAztags())
        if err != nil {
            return err
        }
    }
    if m.GetDataDiskImages() != nil {
        err := writer.WriteCollectionOfStringValues("dataDiskImages", m.GetDataDiskImages())
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
        err := writer.WriteObjectValue("disallowed", m.GetDisallowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extendedLocation", m.GetExtendedLocation())
        if err != nil {
            return err
        }
    }
    if m.GetFeatures() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFeatures()))
        for i, v := range m.GetFeatures() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("features", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hypervGeneration", m.GetHypervGeneration())
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
        err := writer.WriteObjectValue("imageDeprecationStatusInfo", m.GetImageDeprecationStatusInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
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
        err := writer.WriteObjectValue("osDiskImage", m.GetOsDiskImage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("plan", m.GetPlan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("versionInfo", m.GetVersionInfo())
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
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchitecture sets the architecture property value. The cloud machine image architecture (i386 | x86_64 | arm64 ).
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetArchitecture(value *string)() {
    m.architecture = value
}
// SetAutomaticOsUpgradeProperties sets the automaticOsUpgradeProperties property value. Automatic Os Upgrade Properties.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetAutomaticOsUpgradeProperties(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_automaticOsUpgradePropertiesable)() {
    m.automaticOsUpgradeProperties = value
}
// SetAztags sets the aztags property value. The aztags property
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetAztags(value *string)() {
    m.aztags = value
}
// SetDataDiskImages sets the dataDiskImages property value. Images of data disk.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetDataDiskImages(value []string)() {
    m.dataDiskImages = value
}
// SetDescription sets the description property value. The description of the cloud machine image.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetDescription(value *string)() {
    m.description = value
}
// SetDisallowed sets the disallowed property value. The disallowed property
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetDisallowed(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedable)() {
    m.disallowed = value
}
// SetExtendedLocation sets the extendedLocation property value. The extended location of the Virtual Machine.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetExtendedLocation(value *string)() {
    m.extendedLocation = value
}
// SetFeatures sets the features property value. The features property
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetFeatures(value []V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_featuresable)() {
    m.features = value
}
// SetHypervGeneration sets the hypervGeneration property value. The hypervisor generation of the Virtual Machine.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetHypervGeneration(value *string)() {
    m.hypervGeneration = value
}
// SetId sets the id property value. Azure VM image identifier.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetId(value *string)() {
    m.id = value
}
// SetImageDeprecationStatusInfo sets the imageDeprecationStatusInfo property value. Image deprecation status properties on the image.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetImageDeprecationStatusInfo(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_imageDeprecationStatusInfoable)() {
    m.imageDeprecationStatusInfo = value
}
// SetLocation sets the location property value. Image location.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetLocation(value *string)() {
    m.location = value
}
// SetName sets the name property value. Name of the image.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetName(value *string)() {
    m.name = value
}
// SetOsDiskImage sets the osDiskImage property value. The os disk image information.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetOsDiskImage(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_osDiskImageable)() {
    m.osDiskImage = value
}
// SetPlan sets the plan property value. The plan property
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetPlan(value *string)() {
    m.plan = value
}
// SetVersionInfo sets the versionInfo property value. Image version info.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2) SetVersionInfo(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoable)() {
    m.versionInfo = value
}
type V1beta1CspMachineImagesGetResponse_items_cspInfoMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchitecture()(*string)
    GetAutomaticOsUpgradeProperties()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_automaticOsUpgradePropertiesable)
    GetAztags()(*string)
    GetDataDiskImages()([]string)
    GetDescription()(*string)
    GetDisallowed()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedable)
    GetExtendedLocation()(*string)
    GetFeatures()([]V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_featuresable)
    GetHypervGeneration()(*string)
    GetId()(*string)
    GetImageDeprecationStatusInfo()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_imageDeprecationStatusInfoable)
    GetLocation()(*string)
    GetName()(*string)
    GetOsDiskImage()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_osDiskImageable)
    GetPlan()(*string)
    GetVersionInfo()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoable)
    SetArchitecture(value *string)()
    SetAutomaticOsUpgradeProperties(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_automaticOsUpgradePropertiesable)()
    SetAztags(value *string)()
    SetDataDiskImages(value []string)()
    SetDescription(value *string)()
    SetDisallowed(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedable)()
    SetExtendedLocation(value *string)()
    SetFeatures(value []V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_featuresable)()
    SetHypervGeneration(value *string)()
    SetId(value *string)()
    SetImageDeprecationStatusInfo(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_imageDeprecationStatusInfoable)()
    SetLocation(value *string)()
    SetName(value *string)()
    SetOsDiskImage(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_osDiskImageable)()
    SetPlan(value *string)()
    SetVersionInfo(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoable)()
}
