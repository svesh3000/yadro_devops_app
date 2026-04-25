def test_kubeadm_bin_exists(host):
    kubeadm_file = host.file("/usr/bin/kubeadm")
    assert kubeadm_file.exists
    assert kubeadm_file.mode == 0o755


def test_kubeadm_version(host):
    kubeadm_version = host.run("kubeadm version -o short")
    assert kubeadm_version.rc == 0
