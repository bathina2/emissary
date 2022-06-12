from dataclasses import dataclass
from tests.utils import (Compile, logger, default_listener_manifests,
  default_http3_listener_manifest, default_udp_listener_manifest, default_tcp_listener_manifest)
from typing import List, Optional

import pytest
import logging

def http3_quick_start_manifests():
  return default_listener_manifests() + default_http3_listener_manifest()

class TestIR:
  def test_http3_enabled(self, caplog):
    caplog.set_level(logging.WARNING, logger="ambassador")
  
    @dataclass
    class TestCase:
      name:str
      inputYaml: str
      expected: dict[str,bool]
      expectedLog: Optional[str] = None

    testcases = [
      TestCase("quick-start", default_listener_manifests(), { "tcp-0.0.0.0-8080": False, "tcp-0.0.0.0-8443": False }),
      TestCase("quick-start-with_http3", http3_quick_start_manifests(), { "tcp-0.0.0.0-8080": False, "tcp-0.0.0.0-8443": True, "udp-0.0.0.0-8443": True }),
      TestCase("http3-only", default_http3_listener_manifest(),{ "udp-0.0.0.0-8443": True }),
      TestCase("raw-udp", default_udp_listener_manifest(),{}),
      TestCase("raw-tcp", default_tcp_listener_manifest(),{ "tcp-0.0.0.0-8443": False }),
    ]

    for case in testcases:
      compiled_ir = Compile(logger, case.inputYaml, k8s=True)
      result_ir = compiled_ir['ir']

      listeners = result_ir.listeners

      assert len(case.expected.items()) == len(listeners)

      for key, http3_enabled in case.expected.items():
        listener = listeners.get(key, None)
        assert listener != None
        assert listener.http3_enabled == http3_enabled
      
      if case.expectedLog != None:
        assert case.expectedLog in caplog.text
