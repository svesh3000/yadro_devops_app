# Kubectl Setup role

## Short description

Role installs the kubectl package.

## Structure

```
├── meta
│   └── main.yml
└── tasks
    └── main.yml
```

## Requirements

1) hosts must be Debian or Ubuntu
2) `python3-debian` package must be installed on host

## Dependencies

m_sveshnikov.kubernetes_setup.kuber_common

## Vars

| Variable             | Type   | Default Value|
|----------------------|--------|--------------|
| `kubernetes_version` | string | `"v1.32"`    |

## Example

```
---
- name: Setup kubectl
  hosts: kuber
  become: true
  roles:
    - m_sveshnikov.kubernetes_setup.kubectl_setup
```
