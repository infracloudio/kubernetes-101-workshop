# Scheduling

## Labels and Selectors

- Labels are simple key value pairs that you can assign to resources

```

"environment" : "dev", "environment" : "qa", "environment" : "production"

"tier" : "frontend", "tier" : "backend", "tier" : "cache"

"partition" : "customerA", "partition" : "customerB"

```

- Labels classify things

- Selectors allow you to use the labels to filter things

  - Based on equality
  ```
    environment = production
    tier != frontend
  ```

  - Set based
  ```
    environment in (production, qa)
    tier notin (frontend, backend)
    partition
    !partition
  ```

## Scheduling Decisions


### Node Selection

- Attach a label to a node

- Define a `nodeSelector` for the pod!

### Affinity

Three "levels" of affinities!

- Preferred during scheduling ignored during execution

- Required while scheduling and but ignore during execution

- Required for scheduling and also during execution

#### Node affinity

- A pod can be defined to have affinity for certain nodes

#### Pod Affinity

- Schedule pods together

#### Pod Anti-Affinity

- These pods hate each other

#### Taints & Tolerations

