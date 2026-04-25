def test_kubelet_is_installed(host):
    kubelet = host.package("kubelet")
    assert kubelet.is_installed


def test_kubelet_running_and_enabled(host):
    kubelet = host.service("kubelet")
    assert kubelet.exists
