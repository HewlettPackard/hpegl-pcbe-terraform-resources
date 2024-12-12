package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostRequestBodyMember2 secret Specification
type V1beta1SecretsPostRequestBodyMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Description for the resource
    description *string
    // Groups to be associated with the resource
    groups []string
    // Consumer-defined label for the resource
    label *string
    // Key for resource lifecycle event messages
    lifecycleEventKey *string
    // Name of the resource
    name *string
    // Secret policy identifier
    policy *string
    // Secret properties definition
    secret V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretable
    // Name of the service originating the resource
    service *string
}
// V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret composed type wrapper for classes V1beta1SecretsPostRequestBodyMember2_secretMember10able, V1beta1SecretsPostRequestBodyMember2_secretMember1able, V1beta1SecretsPostRequestBodyMember2_secretMember2able, V1beta1SecretsPostRequestBodyMember2_secretMember3able, V1beta1SecretsPostRequestBodyMember2_secretMember4able, V1beta1SecretsPostRequestBodyMember2_secretMember5able, V1beta1SecretsPostRequestBodyMember2_secretMember6able, V1beta1SecretsPostRequestBodyMember2_secretMember7able, V1beta1SecretsPostRequestBodyMember2_secretMember8able, V1beta1SecretsPostRequestBodyMember2_secretMember9able
type V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret struct {
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember1able
    v1beta1SecretsPostRequestBodyMember2_secretMember1 V1beta1SecretsPostRequestBodyMember2_secretMember1able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember10able
    v1beta1SecretsPostRequestBodyMember2_secretMember10 V1beta1SecretsPostRequestBodyMember2_secretMember10able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember2able
    v1beta1SecretsPostRequestBodyMember2_secretMember2 V1beta1SecretsPostRequestBodyMember2_secretMember2able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember3able
    v1beta1SecretsPostRequestBodyMember2_secretMember3 V1beta1SecretsPostRequestBodyMember2_secretMember3able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember4able
    v1beta1SecretsPostRequestBodyMember2_secretMember4 V1beta1SecretsPostRequestBodyMember2_secretMember4able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember5able
    v1beta1SecretsPostRequestBodyMember2_secretMember5 V1beta1SecretsPostRequestBodyMember2_secretMember5able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember6able
    v1beta1SecretsPostRequestBodyMember2_secretMember6 V1beta1SecretsPostRequestBodyMember2_secretMember6able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember7able
    v1beta1SecretsPostRequestBodyMember2_secretMember7 V1beta1SecretsPostRequestBodyMember2_secretMember7able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember8able
    v1beta1SecretsPostRequestBodyMember2_secretMember8 V1beta1SecretsPostRequestBodyMember2_secretMember8able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember9able
    v1beta1SecretsPostRequestBodyMember2_secretMember9 V1beta1SecretsPostRequestBodyMember2_secretMember9able
}
// NewV1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret instantiates a new V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret()(*V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) {
    m := &V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret{
    }
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewV1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember1able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember1(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember10FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember10able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember10(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember2FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember2able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember2(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember3FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember3able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember3(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember4FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember4able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember4(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember5FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember5able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember5(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember6FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember6able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember6(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember7FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember7able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember7(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember8FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember8able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember8(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_secretMember9FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2_secretMember9able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2SecretMember9(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetIsComposedType()(bool) {
    return true
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember1 gets the V1beta1SecretsPostRequestBodyMember2_secretMember1 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember1able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember1able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember1()(V1beta1SecretsPostRequestBodyMember2_secretMember1able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember1
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember10 gets the V1beta1SecretsPostRequestBodyMember2_secretMember10 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember10able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember10able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember10()(V1beta1SecretsPostRequestBodyMember2_secretMember10able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember10
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember2 gets the V1beta1SecretsPostRequestBodyMember2_secretMember2 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember2able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember2able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember2()(V1beta1SecretsPostRequestBodyMember2_secretMember2able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember2
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember3 gets the V1beta1SecretsPostRequestBodyMember2_secretMember3 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember3able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember3able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember3()(V1beta1SecretsPostRequestBodyMember2_secretMember3able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember3
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember4 gets the V1beta1SecretsPostRequestBodyMember2_secretMember4 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember4able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember4able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember4()(V1beta1SecretsPostRequestBodyMember2_secretMember4able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember4
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember5 gets the V1beta1SecretsPostRequestBodyMember2_secretMember5 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember5able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember5able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember5()(V1beta1SecretsPostRequestBodyMember2_secretMember5able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember5
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember6 gets the V1beta1SecretsPostRequestBodyMember2_secretMember6 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember6able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember6able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember6()(V1beta1SecretsPostRequestBodyMember2_secretMember6able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember6
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember7 gets the V1beta1SecretsPostRequestBodyMember2_secretMember7 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember7able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember7able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember7()(V1beta1SecretsPostRequestBodyMember2_secretMember7able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember7
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember8 gets the V1beta1SecretsPostRequestBodyMember2_secretMember8 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember8able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember8able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember8()(V1beta1SecretsPostRequestBodyMember2_secretMember8able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember8
}
// GetV1beta1SecretsPostRequestBodyMember2SecretMember9 gets the V1beta1SecretsPostRequestBodyMember2_secretMember9 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember9able
// returns a V1beta1SecretsPostRequestBodyMember2_secretMember9able when successful
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) GetV1beta1SecretsPostRequestBodyMember2SecretMember9()(V1beta1SecretsPostRequestBodyMember2_secretMember9able) {
    return m.v1beta1SecretsPostRequestBodyMember2_secretMember9
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember1() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember1())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember10() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember10())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember2() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember2())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember3() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember3())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember4() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember4())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember5() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember5())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember6() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember6())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember7() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember7())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember8() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember8())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2SecretMember9() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2SecretMember9())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember1 sets the V1beta1SecretsPostRequestBodyMember2_secretMember1 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember1able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember1(value V1beta1SecretsPostRequestBodyMember2_secretMember1able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember1 = value
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember10 sets the V1beta1SecretsPostRequestBodyMember2_secretMember10 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember10able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember10(value V1beta1SecretsPostRequestBodyMember2_secretMember10able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember10 = value
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember2 sets the V1beta1SecretsPostRequestBodyMember2_secretMember2 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember2able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember2(value V1beta1SecretsPostRequestBodyMember2_secretMember2able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember2 = value
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember3 sets the V1beta1SecretsPostRequestBodyMember2_secretMember3 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember3able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember3(value V1beta1SecretsPostRequestBodyMember2_secretMember3able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember3 = value
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember4 sets the V1beta1SecretsPostRequestBodyMember2_secretMember4 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember4able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember4(value V1beta1SecretsPostRequestBodyMember2_secretMember4able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember4 = value
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember5 sets the V1beta1SecretsPostRequestBodyMember2_secretMember5 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember5able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember5(value V1beta1SecretsPostRequestBodyMember2_secretMember5able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember5 = value
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember6 sets the V1beta1SecretsPostRequestBodyMember2_secretMember6 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember6able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember6(value V1beta1SecretsPostRequestBodyMember2_secretMember6able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember6 = value
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember7 sets the V1beta1SecretsPostRequestBodyMember2_secretMember7 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember7able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember7(value V1beta1SecretsPostRequestBodyMember2_secretMember7able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember7 = value
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember8 sets the V1beta1SecretsPostRequestBodyMember2_secretMember8 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember8able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember8(value V1beta1SecretsPostRequestBodyMember2_secretMember8able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember8 = value
}
// SetV1beta1SecretsPostRequestBodyMember2SecretMember9 sets the V1beta1SecretsPostRequestBodyMember2_secretMember9 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2_secretMember9able
func (m *V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secret) SetV1beta1SecretsPostRequestBodyMember2SecretMember9(value V1beta1SecretsPostRequestBodyMember2_secretMember9able)() {
    m.v1beta1SecretsPostRequestBodyMember2_secretMember9 = value
}
type V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetV1beta1SecretsPostRequestBodyMember2SecretMember1()(V1beta1SecretsPostRequestBodyMember2_secretMember1able)
    GetV1beta1SecretsPostRequestBodyMember2SecretMember10()(V1beta1SecretsPostRequestBodyMember2_secretMember10able)
    GetV1beta1SecretsPostRequestBodyMember2SecretMember2()(V1beta1SecretsPostRequestBodyMember2_secretMember2able)
    GetV1beta1SecretsPostRequestBodyMember2SecretMember3()(V1beta1SecretsPostRequestBodyMember2_secretMember3able)
    GetV1beta1SecretsPostRequestBodyMember2SecretMember4()(V1beta1SecretsPostRequestBodyMember2_secretMember4able)
    GetV1beta1SecretsPostRequestBodyMember2SecretMember5()(V1beta1SecretsPostRequestBodyMember2_secretMember5able)
    GetV1beta1SecretsPostRequestBodyMember2SecretMember6()(V1beta1SecretsPostRequestBodyMember2_secretMember6able)
    GetV1beta1SecretsPostRequestBodyMember2SecretMember7()(V1beta1SecretsPostRequestBodyMember2_secretMember7able)
    GetV1beta1SecretsPostRequestBodyMember2SecretMember8()(V1beta1SecretsPostRequestBodyMember2_secretMember8able)
    GetV1beta1SecretsPostRequestBodyMember2SecretMember9()(V1beta1SecretsPostRequestBodyMember2_secretMember9able)
    SetV1beta1SecretsPostRequestBodyMember2SecretMember1(value V1beta1SecretsPostRequestBodyMember2_secretMember1able)()
    SetV1beta1SecretsPostRequestBodyMember2SecretMember10(value V1beta1SecretsPostRequestBodyMember2_secretMember10able)()
    SetV1beta1SecretsPostRequestBodyMember2SecretMember2(value V1beta1SecretsPostRequestBodyMember2_secretMember2able)()
    SetV1beta1SecretsPostRequestBodyMember2SecretMember3(value V1beta1SecretsPostRequestBodyMember2_secretMember3able)()
    SetV1beta1SecretsPostRequestBodyMember2SecretMember4(value V1beta1SecretsPostRequestBodyMember2_secretMember4able)()
    SetV1beta1SecretsPostRequestBodyMember2SecretMember5(value V1beta1SecretsPostRequestBodyMember2_secretMember5able)()
    SetV1beta1SecretsPostRequestBodyMember2SecretMember6(value V1beta1SecretsPostRequestBodyMember2_secretMember6able)()
    SetV1beta1SecretsPostRequestBodyMember2SecretMember7(value V1beta1SecretsPostRequestBodyMember2_secretMember7able)()
    SetV1beta1SecretsPostRequestBodyMember2SecretMember8(value V1beta1SecretsPostRequestBodyMember2_secretMember8able)()
    SetV1beta1SecretsPostRequestBodyMember2SecretMember9(value V1beta1SecretsPostRequestBodyMember2_secretMember9able)()
}
// NewV1beta1SecretsPostRequestBodyMember2 instantiates a new V1beta1SecretsPostRequestBodyMember2 and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember2()(*V1beta1SecretsPostRequestBodyMember2) {
    m := &V1beta1SecretsPostRequestBodyMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostRequestBodyMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Description for the resource
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetGroups(res)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["lifecycleEventKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifecycleEventKey(val)
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
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val)
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val.(V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretable))
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Groups to be associated with the resource
// returns a []string when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetGroups()([]string) {
    return m.groups
}
// GetLabel gets the label property value. Consumer-defined label for the resource
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetLabel()(*string) {
    return m.label
}
// GetLifecycleEventKey gets the lifecycleEventKey property value. Key for resource lifecycle event messages
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetLifecycleEventKey()(*string) {
    return m.lifecycleEventKey
}
// GetName gets the name property value. Name of the resource
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetName()(*string) {
    return m.name
}
// GetPolicy gets the policy property value. Secret policy identifier
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetPolicy()(*string) {
    return m.policy
}
// GetSecret gets the secret property value. Secret properties definition
// returns a V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretable when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetSecret()(V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretable) {
    return m.secret
}
// GetService gets the service property value. Name of the service originating the resource
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember2) GetService()(*string) {
    return m.service
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        err := writer.WriteCollectionOfStringValues("groups", m.GetGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lifecycleEventKey", m.GetLifecycleEventKey())
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
        err := writer.WriteStringValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret", m.GetSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("service", m.GetService())
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
func (m *V1beta1SecretsPostRequestBodyMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Description for the resource
func (m *V1beta1SecretsPostRequestBodyMember2) SetDescription(value *string)() {
    m.description = value
}
// SetGroups sets the groups property value. Groups to be associated with the resource
func (m *V1beta1SecretsPostRequestBodyMember2) SetGroups(value []string)() {
    m.groups = value
}
// SetLabel sets the label property value. Consumer-defined label for the resource
func (m *V1beta1SecretsPostRequestBodyMember2) SetLabel(value *string)() {
    m.label = value
}
// SetLifecycleEventKey sets the lifecycleEventKey property value. Key for resource lifecycle event messages
func (m *V1beta1SecretsPostRequestBodyMember2) SetLifecycleEventKey(value *string)() {
    m.lifecycleEventKey = value
}
// SetName sets the name property value. Name of the resource
func (m *V1beta1SecretsPostRequestBodyMember2) SetName(value *string)() {
    m.name = value
}
// SetPolicy sets the policy property value. Secret policy identifier
func (m *V1beta1SecretsPostRequestBodyMember2) SetPolicy(value *string)() {
    m.policy = value
}
// SetSecret sets the secret property value. Secret properties definition
func (m *V1beta1SecretsPostRequestBodyMember2) SetSecret(value V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretable)() {
    m.secret = value
}
// SetService sets the service property value. Name of the service originating the resource
func (m *V1beta1SecretsPostRequestBodyMember2) SetService(value *string)() {
    m.service = value
}
type V1beta1SecretsPostRequestBodyMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetGroups()([]string)
    GetLabel()(*string)
    GetLifecycleEventKey()(*string)
    GetName()(*string)
    GetPolicy()(*string)
    GetSecret()(V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretable)
    GetService()(*string)
    SetDescription(value *string)()
    SetGroups(value []string)()
    SetLabel(value *string)()
    SetLifecycleEventKey(value *string)()
    SetName(value *string)()
    SetPolicy(value *string)()
    SetSecret(value V1beta1SecretsPostRequestBodyMember2_SecretsPostRequestBodyMember2_secretable)()
    SetService(value *string)()
}
