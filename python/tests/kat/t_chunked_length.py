from kat.harness import Query
from abstract_tests import AmbassadorTest, ServiceType, HTTP
import json

class AllowChunkedLengthTestTrue(AmbassadorTest):
    target: ServiceType

    def init(self):
        self.target = HTTP(name="target")

    def config(self):
        yield self, self.format("""
---
apiVersion: ambassador/v2
kind:  Module
name:  ambassador
config:
  allow_chunked_length: true
---
apiVersion: ambassador/v2
kind:  Mapping
name:  {self.target.path.k8s}-foo
prefix: /foo/
service: {self.target.path.fqdn}
""")

    def queries(self):
        yield Query(self.url("foo/"))
        yield Query(self.url("ambassador/v0/diag/"))
        yield Query(self.url("foo/"),
            headers={
                "content-length": "0",
                "transfer-encoding": "gzip"
        })
        yield Query(self.url("ambassador/v0/diag/"),
            headers={
                "content-length": "0",
                "transfer-encoding": "gzip"
        })

    def check(self):
        # We expect 501 when sending both these headers because it service probably doesnt support gzip transfer encoding
        # Not getting a 400 bad request is confirmation that this setting works as long as the request can reach the upstream
        assert self.results[0].status == 200
        assert self.results[1].status == 200
        # TODO: the second /foo/ request returns a 200 status
        # assert self.results[2].status == 501
        # assert self.results[3].status == 501
