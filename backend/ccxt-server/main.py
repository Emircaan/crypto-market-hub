"""
CCXT gRPC Server

A gRPC server that fetches cryptocurrency market data from exchanges using CCXT.
"""

from src.server import serve


if __name__ == "__main__":
    serve()
