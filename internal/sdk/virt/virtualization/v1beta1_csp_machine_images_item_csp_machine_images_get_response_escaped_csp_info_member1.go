package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cloud machine image architecture (i386 | x86_64 | arm64 ).
    architecture *string
    // Any block device mapping entries. Describes a block device mapping, which defines the EBS volumes  and instance store volumes to attach to an instance at launch.
    blockDeviceMappings []V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1_blockDeviceMappingsable
    // The description of the cloud machine image.
    description *string
    // Specifies whether enhanced networking with ENA is enabled.
    enaSupport *bool
    // The hypervisor type of the image.
    hypervisor *string
    // Resource identifier assigned by the cloud provider
    id *string
    // The type of image.
    imageType *string
    // The location of the AMI.
    location *string
    // The Amazon Web Services account alias (for example, amazon , self ) or  the Amazon Web Services account ID of the AMI owner.
    ownerAlias *string
    // The ID of the Amazon Web Services account that owns the image.
    ownerId *string
    // The platform details associated with the billing code of the AMI.
    platformDetails *string
    // Indicates whether the image has public launch permissions. The value is true if this image has  public launch permissions or false if it has only implicit and explicit launch permissions.
    public *bool
    // Cloud provider region.
    region *string
    // The device name of the root device volume.
    rootDeviceName *string
    // The type of root device used by the AMI. The AMI can use an Amazon EBS volume  or an instance store volume.
    rootDeviceType *string
    // Specifies whether enhanced networking with the Intel 82599 Virtual Function  interface is enabled.
    sriovNetSupport *string
    // The current state of the AMI. If the state is available, the image is  successfully registered and can be used to launch an instance.
    state *string
    // The operation of the Amazon EC2 instance and the billing code that is associated with the AMI.  usageOperation corresponds to the lineitem/Operation column on your Amazon Web Services Cost and  Usage Report and in the Amazon Web Services Price List API .
    usageOperation *string
    // The type of virtualization of the AMI.
    virtualizationType *string
}
// NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1 instantiates a new V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1 and sets the default values.
func NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1()(*V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) {
    m := &V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchitecture gets the architecture property value. The cloud machine image architecture (i386 | x86_64 | arm64 ).
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetArchitecture()(*string) {
    return m.architecture
}
// GetBlockDeviceMappings gets the blockDeviceMappings property value. Any block device mapping entries. Describes a block device mapping, which defines the EBS volumes  and instance store volumes to attach to an instance at launch.
// returns a []V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1_blockDeviceMappingsable when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetBlockDeviceMappings()([]V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1_blockDeviceMappingsable) {
    return m.blockDeviceMappings
}
// GetDescription gets the description property value. The description of the cloud machine image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetDescription()(*string) {
    return m.description
}
// GetEnaSupport gets the enaSupport property value. Specifies whether enhanced networking with ENA is enabled.
// returns a *bool when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetEnaSupport()(*bool) {
    return m.enaSupport
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["blockDeviceMappings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1_blockDeviceMappingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1_blockDeviceMappingsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1_blockDeviceMappingsable)
                }
            }
            m.SetBlockDeviceMappings(res)
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
    res["enaSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnaSupport(val)
        }
        return nil
    }
    res["hypervisor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisor(val)
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
    res["imageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageType(val)
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
    res["ownerAlias"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerAlias(val)
        }
        return nil
    }
    res["ownerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerId(val)
        }
        return nil
    }
    res["platformDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformDetails(val)
        }
        return nil
    }
    res["public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublic(val)
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["rootDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootDeviceName(val)
        }
        return nil
    }
    res["rootDeviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootDeviceType(val)
        }
        return nil
    }
    res["sriovNetSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSriovNetSupport(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["usageOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageOperation(val)
        }
        return nil
    }
    res["virtualizationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualizationType(val)
        }
        return nil
    }
    return res
}
// GetHypervisor gets the hypervisor property value. The hypervisor type of the image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetHypervisor()(*string) {
    return m.hypervisor
}
// GetId gets the id property value. Resource identifier assigned by the cloud provider
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetId()(*string) {
    return m.id
}
// GetImageType gets the imageType property value. The type of image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetImageType()(*string) {
    return m.imageType
}
// GetLocation gets the location property value. The location of the AMI.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetLocation()(*string) {
    return m.location
}
// GetOwnerAlias gets the ownerAlias property value. The Amazon Web Services account alias (for example, amazon , self ) or  the Amazon Web Services account ID of the AMI owner.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetOwnerAlias()(*string) {
    return m.ownerAlias
}
// GetOwnerId gets the ownerId property value. The ID of the Amazon Web Services account that owns the image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetOwnerId()(*string) {
    return m.ownerId
}
// GetPlatformDetails gets the platformDetails property value. The platform details associated with the billing code of the AMI.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetPlatformDetails()(*string) {
    return m.platformDetails
}
// GetPublic gets the public property value. Indicates whether the image has public launch permissions. The value is true if this image has  public launch permissions or false if it has only implicit and explicit launch permissions.
// returns a *bool when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetPublic()(*bool) {
    return m.public
}
// GetRegion gets the region property value. Cloud provider region.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetRegion()(*string) {
    return m.region
}
// GetRootDeviceName gets the rootDeviceName property value. The device name of the root device volume.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetRootDeviceName()(*string) {
    return m.rootDeviceName
}
// GetRootDeviceType gets the rootDeviceType property value. The type of root device used by the AMI. The AMI can use an Amazon EBS volume  or an instance store volume.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetRootDeviceType()(*string) {
    return m.rootDeviceType
}
// GetSriovNetSupport gets the sriovNetSupport property value. Specifies whether enhanced networking with the Intel 82599 Virtual Function  interface is enabled.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetSriovNetSupport()(*string) {
    return m.sriovNetSupport
}
// GetState gets the state property value. The current state of the AMI. If the state is available, the image is  successfully registered and can be used to launch an instance.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetState()(*string) {
    return m.state
}
// GetUsageOperation gets the usageOperation property value. The operation of the Amazon EC2 instance and the billing code that is associated with the AMI.  usageOperation corresponds to the lineitem/Operation column on your Amazon Web Services Cost and  Usage Report and in the Amazon Web Services Price List API .
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetUsageOperation()(*string) {
    return m.usageOperation
}
// GetVirtualizationType gets the virtualizationType property value. The type of virtualization of the AMI.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) GetVirtualizationType()(*string) {
    return m.virtualizationType
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("architecture", m.GetArchitecture())
        if err != nil {
            return err
        }
    }
    if m.GetBlockDeviceMappings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBlockDeviceMappings()))
        for i, v := range m.GetBlockDeviceMappings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("blockDeviceMappings", cast)
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
        err := writer.WriteBoolValue("enaSupport", m.GetEnaSupport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hypervisor", m.GetHypervisor())
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
        err := writer.WriteStringValue("imageType", m.GetImageType())
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
        err := writer.WriteStringValue("ownerAlias", m.GetOwnerAlias())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ownerId", m.GetOwnerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("platformDetails", m.GetPlatformDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("public", m.GetPublic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rootDeviceName", m.GetRootDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rootDeviceType", m.GetRootDeviceType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sriovNetSupport", m.GetSriovNetSupport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("usageOperation", m.GetUsageOperation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("virtualizationType", m.GetVirtualizationType())
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
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchitecture sets the architecture property value. The cloud machine image architecture (i386 | x86_64 | arm64 ).
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetArchitecture(value *string)() {
    m.architecture = value
}
// SetBlockDeviceMappings sets the blockDeviceMappings property value. Any block device mapping entries. Describes a block device mapping, which defines the EBS volumes  and instance store volumes to attach to an instance at launch.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetBlockDeviceMappings(value []V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1_blockDeviceMappingsable)() {
    m.blockDeviceMappings = value
}
// SetDescription sets the description property value. The description of the cloud machine image.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetDescription(value *string)() {
    m.description = value
}
// SetEnaSupport sets the enaSupport property value. Specifies whether enhanced networking with ENA is enabled.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetEnaSupport(value *bool)() {
    m.enaSupport = value
}
// SetHypervisor sets the hypervisor property value. The hypervisor type of the image.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetHypervisor(value *string)() {
    m.hypervisor = value
}
// SetId sets the id property value. Resource identifier assigned by the cloud provider
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetId(value *string)() {
    m.id = value
}
// SetImageType sets the imageType property value. The type of image.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetImageType(value *string)() {
    m.imageType = value
}
// SetLocation sets the location property value. The location of the AMI.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetLocation(value *string)() {
    m.location = value
}
// SetOwnerAlias sets the ownerAlias property value. The Amazon Web Services account alias (for example, amazon , self ) or  the Amazon Web Services account ID of the AMI owner.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetOwnerAlias(value *string)() {
    m.ownerAlias = value
}
// SetOwnerId sets the ownerId property value. The ID of the Amazon Web Services account that owns the image.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetOwnerId(value *string)() {
    m.ownerId = value
}
// SetPlatformDetails sets the platformDetails property value. The platform details associated with the billing code of the AMI.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetPlatformDetails(value *string)() {
    m.platformDetails = value
}
// SetPublic sets the public property value. Indicates whether the image has public launch permissions. The value is true if this image has  public launch permissions or false if it has only implicit and explicit launch permissions.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetPublic(value *bool)() {
    m.public = value
}
// SetRegion sets the region property value. Cloud provider region.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetRegion(value *string)() {
    m.region = value
}
// SetRootDeviceName sets the rootDeviceName property value. The device name of the root device volume.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetRootDeviceName(value *string)() {
    m.rootDeviceName = value
}
// SetRootDeviceType sets the rootDeviceType property value. The type of root device used by the AMI. The AMI can use an Amazon EBS volume  or an instance store volume.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetRootDeviceType(value *string)() {
    m.rootDeviceType = value
}
// SetSriovNetSupport sets the sriovNetSupport property value. Specifies whether enhanced networking with the Intel 82599 Virtual Function  interface is enabled.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetSriovNetSupport(value *string)() {
    m.sriovNetSupport = value
}
// SetState sets the state property value. The current state of the AMI. If the state is available, the image is  successfully registered and can be used to launch an instance.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetState(value *string)() {
    m.state = value
}
// SetUsageOperation sets the usageOperation property value. The operation of the Amazon EC2 instance and the billing code that is associated with the AMI.  usageOperation corresponds to the lineitem/Operation column on your Amazon Web Services Cost and  Usage Report and in the Amazon Web Services Price List API .
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetUsageOperation(value *string)() {
    m.usageOperation = value
}
// SetVirtualizationType sets the virtualizationType property value. The type of virtualization of the AMI.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1) SetVirtualizationType(value *string)() {
    m.virtualizationType = value
}
type V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchitecture()(*string)
    GetBlockDeviceMappings()([]V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1_blockDeviceMappingsable)
    GetDescription()(*string)
    GetEnaSupport()(*bool)
    GetHypervisor()(*string)
    GetId()(*string)
    GetImageType()(*string)
    GetLocation()(*string)
    GetOwnerAlias()(*string)
    GetOwnerId()(*string)
    GetPlatformDetails()(*string)
    GetPublic()(*bool)
    GetRegion()(*string)
    GetRootDeviceName()(*string)
    GetRootDeviceType()(*string)
    GetSriovNetSupport()(*string)
    GetState()(*string)
    GetUsageOperation()(*string)
    GetVirtualizationType()(*string)
    SetArchitecture(value *string)()
    SetBlockDeviceMappings(value []V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember1_blockDeviceMappingsable)()
    SetDescription(value *string)()
    SetEnaSupport(value *bool)()
    SetHypervisor(value *string)()
    SetId(value *string)()
    SetImageType(value *string)()
    SetLocation(value *string)()
    SetOwnerAlias(value *string)()
    SetOwnerId(value *string)()
    SetPlatformDetails(value *string)()
    SetPublic(value *bool)()
    SetRegion(value *string)()
    SetRootDeviceName(value *string)()
    SetRootDeviceType(value *string)()
    SetSriovNetSupport(value *string)()
    SetState(value *string)()
    SetUsageOperation(value *string)()
    SetVirtualizationType(value *string)()
}
