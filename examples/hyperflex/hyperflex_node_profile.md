### Resource Creation

```hcl
resource "intersight_hyperflex_node_profile" "hyperflex_node_profile1" {
  description = "This is hyperflex_node_profile1"
  assigned_server {
    moid = intersight_hyperflex_cluster_profile.hyperflex_cluster_profile1.moid
  }
  cluster_profile {
    moid = intersight_hyperflex_cluster_profile.hyperflex_cluster_profile1.moid
  }
  name = "ucsback-10G-3nodehx-cluster-NP1"
}
```