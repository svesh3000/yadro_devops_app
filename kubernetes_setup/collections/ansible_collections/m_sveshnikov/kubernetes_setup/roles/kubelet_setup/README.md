# Kubelet Setup role

## Short description

Role installs the kubelet package, and starts the service.

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
3) Host must use `systemd`

## Dependencies

m_sveshnikov.kubernetes_setup.kuber_common

## Vars

| Variable             | Type   | Default Value|
|----------------------|--------|--------------|
| `kubernetes_version` | string | `"v1.32"`    |

## Example

```
---
- name: Setup kubelet
  hosts: kuber
  become: true
  roles:
    - m_sveshnikov.kubernetes_setup.kubelet_setup
```
