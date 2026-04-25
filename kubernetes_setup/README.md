# Kubernetes Setup Playbook

## Short description

Playbook installs cri-o, kubelet and kubeadm

## Structure

```
├── ansible.cfg
├── collections
│   ├── ansible_collections
│   │   └── m_sveshnikov
│   │       └── kubernetes_setup
│   │           ├── galaxy.yml
│   │           └── roles
│   │               ├── crio_setup
│   │               │   ├── README.md
│   │               │   ├── defaults
│   │               │   │   └── main.yml
│   │               │   ├── molecule
│   │               │   │   └── default
│   │               │   │       ├── converge.yml
│   │               │   │       ├── molecule.yml
│   │               │   │       ├── prepare.yml
│   │               │   │       └── tests
│   │               │   │           └── test_default.py
│   │               │   └── tasks
│   │               │       └── main.yml
│   │               ├── kubeadm_setup
│   │               │   ├── README.md
│   │               │   ├── meta
│   │               │   │   └── main.yml
│   │               │   ├── molecule
│   │               │   │   └── default
│   │               │   │       ├── converge.yml
│   │               │   │       ├── molecule.yml
│   │               │   │       ├── prepare.yml
│   │               │   │       └── tests
│   │               │   │           └── test_default.py
│   │               │   └── tasks
│   │               │       └── main.yml
│   │               ├── kubelet_setup
│   │               │   ├── README.md
│   │               │   ├── meta
│   │               │   │   └── main.yml
│   │               │   ├── molecule
│   │               │   │   └── default
│   │               │   │       ├── converge.yml
│   │               │   │       ├── molecule.yml
│   │               │   │       ├── prepare.yml
│   │               │   │       └── tests
│   │               │   │           └── test_default.py
│   │               │   └── tasks
│   │               │       └── main.yml
│   │               └── kuber_common
│   │                   ├── README.md
│   │                   ├── defaults
│   │                   │   └── main.yml
│   │                   └── tasks
│   │                       └── main.yml
│   └── requirements.yml
├── inventory
│   └── hosts.ini
└── site.yml
```

## Requirements

1) hosts must be Debian or Ubuntu
2) `python3-debian` package must be installed on host
3) Host must use `systemd`

## Dependencies

none.

## Vars

| Variable             | Type   | Default Value|
|----------------------|--------|--------------|
| `kubernetes_version` | string | `"v1.32"`    |
| `crio_version`       | string | `"1.32"`     |

## Run:

```bash
  ansible-playbook -i inventory/hosts.ini site.yml
```
