
resource "intersight_fabric_eth_network_policy" "fabric_eth_network_policy1" {
  name        = "fabric_eth_network_policy1"
  description = "fabric ethernet network policy"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
