import logging
from concurrent import futures

import grpc

from src.generated import crypto_market_pb2_grpc as pb2_grpc
from src.services import CryptoMarketServicer
from src.config import settings


logger = logging.getLogger(__name__)


def create_server(host: str = None, port: int = None) -> grpc.Server:
    host = host or settings.grpc_host
    port = port or settings.grpc_port
    
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    
    pb2_grpc.add_CryptoMarketServiceServicer_to_server(
        CryptoMarketServicer(), 
        server
    )
    
    address = f"{host}:{port}"
    server.add_insecure_port(address)
    
    logger.info(f"gRPC server configured on {address}")
    
    return server


def serve():
    logging.basicConfig(
        level=getattr(logging, settings.log_level),
        format="%(asctime)s - %(name)s - %(levelname)s - %(message)s"
    )
    
    server = create_server()
    server.start()
    
    logger.info(f"CCXT gRPC Server started on {settings.grpc_host}:{settings.grpc_port}")
    logger.info(f"Supported exchanges: {settings.supported_exchanges}")
    
    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        logger.info("Shutting down server...")
        server.stop(grace=5)
        logger.info("Server stopped")

if __name__ == "__main__":
    serve()
