# Kuber Common role

## Short description

Role configures the repository for the specified version of Kubernetes.

## Structure

```
├── meta
│   └── main.yml
├── molecule
│   └── default
│       ├── converge.yml
│       ├── molecule.yml
│       ├── prepare.yml
│       └── tests
│           └── test_default.py
└── tasks
    └── main.yml
```

## Requirements

1) hosts must be Debian or Ubuntu
2) `python3-debian` package must be installed on host

## Dependencies

none.

## Vars

| Variable             | Type   | Default Value|
|----------------------|--------|--------------|
| `kubernetes_version` | string | `"v1.32"`    |

## Example

```
---
- name: Setup kuber
  hosts: kuber
  become: true
  roles:
    - m_sveshnikov.kubernetes_setup.kubelet_common
```
