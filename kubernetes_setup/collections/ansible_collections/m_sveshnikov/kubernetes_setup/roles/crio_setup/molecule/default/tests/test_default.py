def test_crio_is_installed(host):
    crio = host.package("cri-o")
    assert crio.is_installed


def test_crio_running_and_enabled(host):
    crio = host.service("crio")
    assert crio.is_running
    assert crio.is_enabled
