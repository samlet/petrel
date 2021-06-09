import logging
from concurrent.futures import ThreadPoolExecutor
from google.protobuf import json_format

import grpc
import numpy as np
from signal import signal, SIGTERM
from sagas.ofbiz.entities import entity

from outliers_pb2 import OutliersResponse, EntityInfo
from outliers_pb2_grpc import OutliersServicer, add_OutliersServicer_to_server


def find_outliers(data: np.ndarray):
    """Return indices where values more than 2 standard deviations from mean"""
    out = np.where(np.abs(data - data.mean()) > 2 * data.std())
    # np.where returns a tuple for each dimension, we want the 1st element
    return out[0]


class OutliersServer(OutliersServicer):
    def Detect(self, request, context):
        logging.info('detect request size: %d', len(request.metrics))
        # Convert metrics to numpy array of values only
        data = np.fromiter((m.value for m in request.metrics), dtype='float64')
        indices = find_outliers(data)
        logging.info('found %d outliers', len(indices))
        resp = OutliersResponse(indices=indices)
        return resp

    def GetEntityInfo(self, request, context):
        ent=request.name
        model=entity(ent)
        info={"name":ent,
              "package_name": model.package_name
              }
        resp=json_format.ParseDict(info, EntityInfo())
        return resp


if __name__ == '__main__':
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
    )
    server = grpc.server(ThreadPoolExecutor())
    add_OutliersServicer_to_server(OutliersServer(), server)
    port = 9999
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    logging.info('server ready on port %r', port)

    def handle_sigterm(*_):
        print("Received shutdown signal")
        all_rpcs_done_event = server.stop(30)
        all_rpcs_done_event.wait(30)
        print("Shut down gracefully")

    signal(SIGTERM, handle_sigterm)

    server.wait_for_termination()

