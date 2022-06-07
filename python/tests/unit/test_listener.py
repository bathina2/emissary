from dataclasses import dataclass
from tests.utils import Compile, logger
from typing import List, Optional

import pytest


def _generateListener(name: str, protocol: Optional[str], protocol_stack:Optional[List[str]]):
    yaml = f"""
apiVersion: getambassador.io/v3alpha1
kind: Listener
metadata:
    name: {name}
    namespace: ambassador
spec:
    port: 8443
    {f"protocolStack: {protocol_stack}" if protocol == None else f"protocol: {protocol}"}
    securityModel: XFP
    hostBinding:
        namespace:
            from: ALL
"""
    return yaml

class TestListener:

    @pytest.mark.compilertest
    def test_socket_protocol(self):
        """ensure that we can identify the listener socket protocol based on the provided protocol and protocolStack"""

        @dataclass
        class TestCase:
            name: str
            protocol: Optional[str]
            protocolStack: Optional[List[str]]
            expectedSocketProtocol: Optional[str]

        testcases = [
                # test with emissary defined protcolStacks via pre-definied protocol enum
                TestCase(name="http_protocol", protocol="HTTP", protocolStack=None, expectedSocketProtocol="TCP"),
                TestCase(name="https_protocol", protocol="HTTPS", protocolStack=None, expectedSocketProtocol="TCP"),
                TestCase(name="httpproxy_protocol", protocol="HTTPPROXY", protocolStack=None, expectedSocketProtocol="TCP"),
                TestCase(name="httpsproxy_protocol", protocol="HTTPSPROXY", protocolStack=None, expectedSocketProtocol="TCP"),
                TestCase(name="tcp_protocol", protocol="TCP", protocolStack=None, expectedSocketProtocol="TCP"),
                TestCase(name="tls_protocol", protocol="TLS", protocolStack=None, expectedSocketProtocol="TCP"),

                # test with custom stacks
                TestCase(name="tcp_stack", protocol=None, protocolStack=["TLS", "HTTP", "TCP"], expectedSocketProtocol="TCP"),
                TestCase(name="udp_stack", protocol=None, protocolStack=["TLS", "HTTP", "UDP"], expectedSocketProtocol="UDP"),
                TestCase(name="invalid_stack", protocol=None, protocolStack=["TLS", "HTTP"], expectedSocketProtocol=None),
                TestCase(name="empty_stack", protocol=None, protocolStack=[], expectedSocketProtocol=None),
            ]
        for case in testcases:
            yaml = _generateListener(case.name, case.protocol, case.protocolStack)

            compiled_ir = Compile(logger, yaml, k8s=True)
            result_ir = compiled_ir['ir']
            listeners = list(result_ir.listeners.values())
            errors = result_ir.aconf.errors

            if case.expectedSocketProtocol == None:
                assert len(errors) == 1
                assert len(listeners) == 0
            else:
                assert len(listeners) == 1
                assert listeners[0].socket_protocol == case.expectedSocketProtocol
    