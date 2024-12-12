package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesPostRequestBody_vmConfig defines the virtual machine configurations
type V1beta1VirtualMachinesPostRequestBody_vmConfig struct {
    // Accept EULA by default or not
    acceptEula *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of the image
    annotation *string
    // Locale to use for parsing OVF descriptor
    locale *string
    // Name of the virtual machine to be deployed
    name *string
    // Number of virtual machines to be created.
    numberOfVms *int32
    // Power on/off the virtual machine
    powerOn *bool
    // Properties which can be set to a virtual machine during deployment
    propertyConfig []V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigable
}
// NewV1beta1VirtualMachinesPostRequestBody_vmConfig instantiates a new V1beta1VirtualMachinesPostRequestBody_vmConfig and sets the default values.
func NewV1beta1VirtualMachinesPostRequestBody_vmConfig()(*V1beta1VirtualMachinesPostRequestBody_vmConfig) {
    m := &V1beta1VirtualMachinesPostRequestBody_vmConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesPostRequestBody_vmConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesPostRequestBody_vmConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesPostRequestBody_vmConfig(), nil
}
// GetAcceptEula gets the acceptEula property value. Accept EULA by default or not
// returns a *bool when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) GetAcceptEula()(*bool) {
    return m.acceptEula
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnnotation gets the annotation property value. The description of the image
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) GetAnnotation()(*string) {
    return m.annotation
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptEula"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptEula(val)
        }
        return nil
    }
    res["annotation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnnotation(val)
        }
        return nil
    }
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
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
    res["numberOfVms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfVms(val)
        }
        return nil
    }
    res["powerOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPowerOn(val)
        }
        return nil
    }
    res["propertyConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigable)
                }
            }
            m.SetPropertyConfig(res)
        }
        return nil
    }
    return res
}
// GetLocale gets the locale property value. Locale to use for parsing OVF descriptor
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) GetLocale()(*string) {
    return m.locale
}
// GetName gets the name property value. Name of the virtual machine to be deployed
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) GetName()(*string) {
    return m.name
}
// GetNumberOfVms gets the numberOfVms property value. Number of virtual machines to be created.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) GetNumberOfVms()(*int32) {
    return m.numberOfVms
}
// GetPowerOn gets the powerOn property value. Power on/off the virtual machine
// returns a *bool when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) GetPowerOn()(*bool) {
    return m.powerOn
}
// GetPropertyConfig gets the propertyConfig property value. Properties which can be set to a virtual machine during deployment
// returns a []V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigable when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) GetPropertyConfig()([]V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigable) {
    return m.propertyConfig
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acceptEula", m.GetAcceptEula())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("annotation", m.GetAnnotation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
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
        err := writer.WriteInt32Value("numberOfVms", m.GetNumberOfVms())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("powerOn", m.GetPowerOn())
        if err != nil {
            return err
        }
    }
    if m.GetPropertyConfig() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPropertyConfig()))
        for i, v := range m.GetPropertyConfig() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("propertyConfig", cast)
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
// SetAcceptEula sets the acceptEula property value. Accept EULA by default or not
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) SetAcceptEula(value *bool)() {
    m.acceptEula = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnnotation sets the annotation property value. The description of the image
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) SetAnnotation(value *string)() {
    m.annotation = value
}
// SetLocale sets the locale property value. Locale to use for parsing OVF descriptor
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) SetLocale(value *string)() {
    m.locale = value
}
// SetName sets the name property value. Name of the virtual machine to be deployed
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) SetName(value *string)() {
    m.name = value
}
// SetNumberOfVms sets the numberOfVms property value. Number of virtual machines to be created.
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) SetNumberOfVms(value *int32)() {
    m.numberOfVms = value
}
// SetPowerOn sets the powerOn property value. Power on/off the virtual machine
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) SetPowerOn(value *bool)() {
    m.powerOn = value
}
// SetPropertyConfig sets the propertyConfig property value. Properties which can be set to a virtual machine during deployment
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig) SetPropertyConfig(value []V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigable)() {
    m.propertyConfig = value
}
type V1beta1VirtualMachinesPostRequestBody_vmConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptEula()(*bool)
    GetAnnotation()(*string)
    GetLocale()(*string)
    GetName()(*string)
    GetNumberOfVms()(*int32)
    GetPowerOn()(*bool)
    GetPropertyConfig()([]V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigable)
    SetAcceptEula(value *bool)()
    SetAnnotation(value *string)()
    SetLocale(value *string)()
    SetName(value *string)()
    SetNumberOfVms(value *int32)()
    SetPowerOn(value *bool)()
    SetPropertyConfig(value []V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigable)()
}
