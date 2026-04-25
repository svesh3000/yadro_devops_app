# CRI-O Setup role

## Short description

Role configures the repository for the specified version, installs the cri-o package, and starts the service.

## Structure

├── defaults
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

## Requirements

1) hosts must be Debian or Ubuntu
2) `python3-debian` package must be installed on host
3) Host must use `systemd`

## Dependencies

none.

## Vars

| Variable       | Type   | Default Value|
|----------------|--------|--------------|
| `crio_version` | string | `"1.32"`     |

## Example

```
---
- name: Setup cri-o
  hosts: kuber
  become: true
  roles:
    - m_sveshnikov.kubernetes_setup.crio_setup
```
