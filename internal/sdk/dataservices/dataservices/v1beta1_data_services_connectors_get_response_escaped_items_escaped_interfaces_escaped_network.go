package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DataServicesConnectorsGetResponse_items_interfaces_network struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // IP address or FQDN of the default gateway.
    defaultGateway *string
    // IP address or FQDN of the Data Services Connector.
    hostname *string
    // List of configured DNS servers configured on the Data Services Connector.
    nameServers []V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nameServersable
    // List of network interfaces configured on the Data Services Connector.
    nic []V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicable
    // The proxy property
    proxy V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyable
    // List of search domains.
    searchDomains []string
}
// NewV1beta1DataServicesConnectorsGetResponse_items_interfaces_network instantiates a new V1beta1DataServicesConnectorsGetResponse_items_interfaces_network and sets the default values.
func NewV1beta1DataServicesConnectorsGetResponse_items_interfaces_network()(*V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) {
    m := &V1beta1DataServicesConnectorsGetResponse_items_interfaces_network{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_networkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_networkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsGetResponse_items_interfaces_network(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultGateway gets the defaultGateway property value. IP address or FQDN of the default gateway.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) GetDefaultGateway()(*string) {
    return m.defaultGateway
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultGateway"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultGateway(val)
        }
        return nil
    }
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
    res["nameServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nameServersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nameServersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nameServersable)
                }
            }
            m.SetNameServers(res)
        }
        return nil
    }
    res["nic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicable)
                }
            }
            m.SetNic(res)
        }
        return nil
    }
    res["proxy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxy(val.(V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyable))
        }
        return nil
    }
    res["searchDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSearchDomains(res)
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. IP address or FQDN of the Data Services Connector.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) GetHostname()(*string) {
    return m.hostname
}
// GetNameServers gets the nameServers property value. List of configured DNS servers configured on the Data Services Connector.
// returns a []V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nameServersable when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) GetNameServers()([]V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nameServersable) {
    return m.nameServers
}
// GetNic gets the nic property value. List of network interfaces configured on the Data Services Connector.
// returns a []V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicable when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) GetNic()([]V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicable) {
    return m.nic
}
// GetProxy gets the proxy property value. The proxy property
// returns a V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyable when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) GetProxy()(V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyable) {
    return m.proxy
}
// GetSearchDomains gets the searchDomains property value. List of search domains.
// returns a []string when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) GetSearchDomains()([]string) {
    return m.searchDomains
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultGateway", m.GetDefaultGateway())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    if m.GetNameServers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNameServers()))
        for i, v := range m.GetNameServers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("nameServers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNic() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNic()))
        for i, v := range m.GetNic() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("nic", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("proxy", m.GetProxy())
        if err != nil {
            return err
        }
    }
    if m.GetSearchDomains() != nil {
        err := writer.WriteCollectionOfStringValues("searchDomains", m.GetSearchDomains())
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
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultGateway sets the defaultGateway property value. IP address or FQDN of the default gateway.
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) SetDefaultGateway(value *string)() {
    m.defaultGateway = value
}
// SetHostname sets the hostname property value. IP address or FQDN of the Data Services Connector.
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) SetHostname(value *string)() {
    m.hostname = value
}
// SetNameServers sets the nameServers property value. List of configured DNS servers configured on the Data Services Connector.
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) SetNameServers(value []V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nameServersable)() {
    m.nameServers = value
}
// SetNic sets the nic property value. List of network interfaces configured on the Data Services Connector.
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) SetNic(value []V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicable)() {
    m.nic = value
}
// SetProxy sets the proxy property value. The proxy property
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) SetProxy(value V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyable)() {
    m.proxy = value
}
// SetSearchDomains sets the searchDomains property value. List of search domains.
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network) SetSearchDomains(value []string)() {
    m.searchDomains = value
}
type V1beta1DataServicesConnectorsGetResponse_items_interfaces_networkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultGateway()(*string)
    GetHostname()(*string)
    GetNameServers()([]V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nameServersable)
    GetNic()([]V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicable)
    GetProxy()(V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyable)
    GetSearchDomains()([]string)
    SetDefaultGateway(value *string)()
    SetHostname(value *string)()
    SetNameServers(value []V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nameServersable)()
    SetNic(value []V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicable)()
    SetProxy(value V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyable)()
    SetSearchDomains(value []string)()
}
