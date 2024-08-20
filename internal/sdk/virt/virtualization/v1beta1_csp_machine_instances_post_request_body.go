package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstancesPostRequestBody struct {
    // ID of the csp-account resource representing the cloud provider account.
    accountId *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cspMachineInstanceInfo property
    cspMachineInstanceInfo V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoable
    // Name of User's KeyPair, required for CSP Machine Instance login
    keyPairName *string
    // Name of CSP machine instance provided by user
    name *string
}
// V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo composed type wrapper for classes V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able, V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able
type V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo struct {
    // Composed type representation for type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able
    v1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1 V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able
    // Composed type representation for type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able
    v1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2 V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able
}
// NewV1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo instantiates a new V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo and sets the default values.
func NewV1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo()(*V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo) {
    m := &V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo{
    }
    return m
}
// CreateV1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewV1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo) GetIsComposedType()(bool) {
    return true
}
// GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember1 gets the V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1 property value. Composed type representation for type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able
// returns a V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo) GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember1()(V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able) {
    return m.v1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1
}
// GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember2 gets the V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2 property value. Composed type representation for type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able
// returns a V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo) GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember2()(V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able) {
    return m.v1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember1() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember1())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember2() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember1 sets the V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1 property value. Composed type representation for type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able
func (m *V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo) SetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember1(value V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able)() {
    m.v1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1 = value
}
// SetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember2 sets the V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2 property value. Composed type representation for type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able
func (m *V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfo) SetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember2(value V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able)() {
    m.v1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2 = value
}
type V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember1()(V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able)
    GetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember2()(V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able)
    SetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember1(value V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able)()
    SetV1beta1CspMachineInstancesPostRequestBodyCspMachineInstanceInfoMember2(value V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able)()
}
// NewV1beta1CspMachineInstancesPostRequestBody instantiates a new V1beta1CspMachineInstancesPostRequestBody and sets the default values.
func NewV1beta1CspMachineInstancesPostRequestBody()(*V1beta1CspMachineInstancesPostRequestBody) {
    m := &V1beta1CspMachineInstancesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesPostRequestBody(), nil
}
// GetAccountId gets the accountId property value. ID of the csp-account resource representing the cloud provider account.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody) GetAccountId()(*string) {
    return m.accountId
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCspMachineInstanceInfo gets the cspMachineInstanceInfo property value. The cspMachineInstanceInfo property
// returns a V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoable when successful
func (m *V1beta1CspMachineInstancesPostRequestBody) GetCspMachineInstanceInfo()(V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoable) {
    return m.cspMachineInstanceInfo
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["cspMachineInstanceInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspMachineInstanceInfo(val.(V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoable))
        }
        return nil
    }
    res["keyPairName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyPairName(val)
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
    return res
}
// GetKeyPairName gets the keyPairName property value. Name of User's KeyPair, required for CSP Machine Instance login
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody) GetKeyPairName()(*string) {
    return m.keyPairName
}
// GetName gets the name property value. Name of CSP machine instance provided by user
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accountId", m.GetAccountId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cspMachineInstanceInfo", m.GetCspMachineInstanceInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("keyPairName", m.GetKeyPairName())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountId sets the accountId property value. ID of the csp-account resource representing the cloud provider account.
func (m *V1beta1CspMachineInstancesPostRequestBody) SetAccountId(value *string)() {
    m.accountId = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspMachineInstancesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCspMachineInstanceInfo sets the cspMachineInstanceInfo property value. The cspMachineInstanceInfo property
func (m *V1beta1CspMachineInstancesPostRequestBody) SetCspMachineInstanceInfo(value V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoable)() {
    m.cspMachineInstanceInfo = value
}
// SetKeyPairName sets the keyPairName property value. Name of User's KeyPair, required for CSP Machine Instance login
func (m *V1beta1CspMachineInstancesPostRequestBody) SetKeyPairName(value *string)() {
    m.keyPairName = value
}
// SetName sets the name property value. Name of CSP machine instance provided by user
func (m *V1beta1CspMachineInstancesPostRequestBody) SetName(value *string)() {
    m.name = value
}
type V1beta1CspMachineInstancesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountId()(*string)
    GetCspMachineInstanceInfo()(V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoable)
    GetKeyPairName()(*string)
    GetName()(*string)
    SetAccountId(value *string)()
    SetCspMachineInstanceInfo(value V1beta1CspMachineInstancesPostRequestBody_CspMachineInstancesPostRequestBody_cspMachineInstanceInfoable)()
    SetKeyPairName(value *string)()
    SetName(value *string)()
}
