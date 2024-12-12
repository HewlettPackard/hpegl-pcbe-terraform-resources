package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesPostRequestBody_imageSource specifies the hypervisor image information using which the virtual machine is deployed
type V1beta1VirtualMachinesPostRequestBody_imageSource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The UUID of the hypervisor image using which a virtual machine is deployed
    imageId *string
    // The name of the hypervisor image using which a virtual machine is deployed
    imageName *string
}
// NewV1beta1VirtualMachinesPostRequestBody_imageSource instantiates a new V1beta1VirtualMachinesPostRequestBody_imageSource and sets the default values.
func NewV1beta1VirtualMachinesPostRequestBody_imageSource()(*V1beta1VirtualMachinesPostRequestBody_imageSource) {
    m := &V1beta1VirtualMachinesPostRequestBody_imageSource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesPostRequestBody_imageSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesPostRequestBody_imageSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesPostRequestBody_imageSource(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesPostRequestBody_imageSource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesPostRequestBody_imageSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["imageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageId(val)
        }
        return nil
    }
    res["imageName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageName(val)
        }
        return nil
    }
    return res
}
// GetImageId gets the imageId property value. The UUID of the hypervisor image using which a virtual machine is deployed
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_imageSource) GetImageId()(*string) {
    return m.imageId
}
// GetImageName gets the imageName property value. The name of the hypervisor image using which a virtual machine is deployed
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_imageSource) GetImageName()(*string) {
    return m.imageName
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesPostRequestBody_imageSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("imageId", m.GetImageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("imageName", m.GetImageName())
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
func (m *V1beta1VirtualMachinesPostRequestBody_imageSource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetImageId sets the imageId property value. The UUID of the hypervisor image using which a virtual machine is deployed
func (m *V1beta1VirtualMachinesPostRequestBody_imageSource) SetImageId(value *string)() {
    m.imageId = value
}
// SetImageName sets the imageName property value. The name of the hypervisor image using which a virtual machine is deployed
func (m *V1beta1VirtualMachinesPostRequestBody_imageSource) SetImageName(value *string)() {
    m.imageName = value
}
type V1beta1VirtualMachinesPostRequestBody_imageSourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetImageId()(*string)
    GetImageName()(*string)
    SetImageId(value *string)()
    SetImageName(value *string)()
}
