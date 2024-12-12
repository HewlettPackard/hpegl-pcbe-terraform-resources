package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo the network settings for the instance type.
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The index of the default network card, starting at 0.
    defaultNetworkCardIndex *int32
    // Indicates whether Elastic Fabric Adapter (EFA) is supported.
    efaSupported *bool
    // Indicates whether Elastic Network Adapter (ENA) is supported.
    enaSupport *string
    // The maximum number of IPv4 addresses per network interface.
    ipv4AddressesPerInterface *int32
    // The maximum number of IPv6 addresses per network interface.
    ipv6AddressesPerInterface *int32
    // Indicates whether IPv6 is supported.
    ipv6Supported *bool
    // The maximum number of physical network cards that can be allocated to the instance.
    maximumNetworkCards *int32
    // The maximum number of network interfaces for the instance type.
    maximumNetworkInterfaces *int32
    // The network cards for the instance type.
    networkCards V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsable
    // The network performance.
    networkPerformance *string
}
// NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo instantiates a new V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo and sets the default values.
func NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo()(*V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) {
    m := &V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultNetworkCardIndex gets the defaultNetworkCardIndex property value. The index of the default network card, starting at 0.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetDefaultNetworkCardIndex()(*int32) {
    return m.defaultNetworkCardIndex
}
// GetEfaSupported gets the efaSupported property value. Indicates whether Elastic Fabric Adapter (EFA) is supported.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetEfaSupported()(*bool) {
    return m.efaSupported
}
// GetEnaSupport gets the enaSupport property value. Indicates whether Elastic Network Adapter (ENA) is supported.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetEnaSupport()(*string) {
    return m.enaSupport
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultNetworkCardIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultNetworkCardIndex(val)
        }
        return nil
    }
    res["efaSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEfaSupported(val)
        }
        return nil
    }
    res["enaSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnaSupport(val)
        }
        return nil
    }
    res["ipv4AddressesPerInterface"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpv4AddressesPerInterface(val)
        }
        return nil
    }
    res["ipv6AddressesPerInterface"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpv6AddressesPerInterface(val)
        }
        return nil
    }
    res["ipv6Supported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpv6Supported(val)
        }
        return nil
    }
    res["maximumNetworkCards"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumNetworkCards(val)
        }
        return nil
    }
    res["maximumNetworkInterfaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumNetworkInterfaces(val)
        }
        return nil
    }
    res["networkCards"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkCards(val.(V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsable))
        }
        return nil
    }
    res["networkPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkPerformance(val)
        }
        return nil
    }
    return res
}
// GetIpv4AddressesPerInterface gets the ipv4AddressesPerInterface property value. The maximum number of IPv4 addresses per network interface.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetIpv4AddressesPerInterface()(*int32) {
    return m.ipv4AddressesPerInterface
}
// GetIpv6AddressesPerInterface gets the ipv6AddressesPerInterface property value. The maximum number of IPv6 addresses per network interface.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetIpv6AddressesPerInterface()(*int32) {
    return m.ipv6AddressesPerInterface
}
// GetIpv6Supported gets the ipv6Supported property value. Indicates whether IPv6 is supported.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetIpv6Supported()(*bool) {
    return m.ipv6Supported
}
// GetMaximumNetworkCards gets the maximumNetworkCards property value. The maximum number of physical network cards that can be allocated to the instance.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetMaximumNetworkCards()(*int32) {
    return m.maximumNetworkCards
}
// GetMaximumNetworkInterfaces gets the maximumNetworkInterfaces property value. The maximum number of network interfaces for the instance type.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetMaximumNetworkInterfaces()(*int32) {
    return m.maximumNetworkInterfaces
}
// GetNetworkCards gets the networkCards property value. The network cards for the instance type.
// returns a V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsable when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetNetworkCards()(V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsable) {
    return m.networkCards
}
// GetNetworkPerformance gets the networkPerformance property value. The network performance.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) GetNetworkPerformance()(*string) {
    return m.networkPerformance
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("defaultNetworkCardIndex", m.GetDefaultNetworkCardIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("efaSupported", m.GetEfaSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("enaSupport", m.GetEnaSupport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ipv4AddressesPerInterface", m.GetIpv4AddressesPerInterface())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ipv6AddressesPerInterface", m.GetIpv6AddressesPerInterface())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ipv6Supported", m.GetIpv6Supported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maximumNetworkCards", m.GetMaximumNetworkCards())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maximumNetworkInterfaces", m.GetMaximumNetworkInterfaces())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("networkCards", m.GetNetworkCards())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("networkPerformance", m.GetNetworkPerformance())
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
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultNetworkCardIndex sets the defaultNetworkCardIndex property value. The index of the default network card, starting at 0.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetDefaultNetworkCardIndex(value *int32)() {
    m.defaultNetworkCardIndex = value
}
// SetEfaSupported sets the efaSupported property value. Indicates whether Elastic Fabric Adapter (EFA) is supported.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetEfaSupported(value *bool)() {
    m.efaSupported = value
}
// SetEnaSupport sets the enaSupport property value. Indicates whether Elastic Network Adapter (ENA) is supported.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetEnaSupport(value *string)() {
    m.enaSupport = value
}
// SetIpv4AddressesPerInterface sets the ipv4AddressesPerInterface property value. The maximum number of IPv4 addresses per network interface.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetIpv4AddressesPerInterface(value *int32)() {
    m.ipv4AddressesPerInterface = value
}
// SetIpv6AddressesPerInterface sets the ipv6AddressesPerInterface property value. The maximum number of IPv6 addresses per network interface.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetIpv6AddressesPerInterface(value *int32)() {
    m.ipv6AddressesPerInterface = value
}
// SetIpv6Supported sets the ipv6Supported property value. Indicates whether IPv6 is supported.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetIpv6Supported(value *bool)() {
    m.ipv6Supported = value
}
// SetMaximumNetworkCards sets the maximumNetworkCards property value. The maximum number of physical network cards that can be allocated to the instance.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetMaximumNetworkCards(value *int32)() {
    m.maximumNetworkCards = value
}
// SetMaximumNetworkInterfaces sets the maximumNetworkInterfaces property value. The maximum number of network interfaces for the instance type.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetMaximumNetworkInterfaces(value *int32)() {
    m.maximumNetworkInterfaces = value
}
// SetNetworkCards sets the networkCards property value. The network cards for the instance type.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetNetworkCards(value V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsable)() {
    m.networkCards = value
}
// SetNetworkPerformance sets the networkPerformance property value. The network performance.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo) SetNetworkPerformance(value *string)() {
    m.networkPerformance = value
}
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultNetworkCardIndex()(*int32)
    GetEfaSupported()(*bool)
    GetEnaSupport()(*string)
    GetIpv4AddressesPerInterface()(*int32)
    GetIpv6AddressesPerInterface()(*int32)
    GetIpv6Supported()(*bool)
    GetMaximumNetworkCards()(*int32)
    GetMaximumNetworkInterfaces()(*int32)
    GetNetworkCards()(V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsable)
    GetNetworkPerformance()(*string)
    SetDefaultNetworkCardIndex(value *int32)()
    SetEfaSupported(value *bool)()
    SetEnaSupport(value *string)()
    SetIpv4AddressesPerInterface(value *int32)()
    SetIpv6AddressesPerInterface(value *int32)()
    SetIpv6Supported(value *bool)()
    SetMaximumNetworkCards(value *int32)()
    SetMaximumNetworkInterfaces(value *int32)()
    SetNetworkCards(value V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsable)()
    SetNetworkPerformance(value *string)()
}
