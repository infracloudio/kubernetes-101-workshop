## Kubernetes!

- Container orchestration platform

- Portable across clouds/on prem data centers

- Why Platform?
  - CRDs (Custom Resource Definition) and Controller - extensible
  - Build new applications which behave as if they are part of core Kubernetes (But they are not)
  - Many of apps are built this way - Helm, Fission
  - We will touch on this again once we know Kubernetes built in objects better.

- Not a PaaS (But you can build a PaaS on it, in fact FaaS on Kubernetes are somewhat that)
 
- Abstracts the underlying infrastructure. You say run my app - but you need not know where it is running!

- Represents underlying infra as a single entity

- Is pluggable - for example 
  - For networking common networking interface (CNI)
  - CSI for storage