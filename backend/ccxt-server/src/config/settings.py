from pydantic_settings import BaseSettings
from typing import List


class Settings(BaseSettings):
    """Application settings loaded from environment variables."""
    
    grpc_host: str = "0.0.0.0"
    grpc_port: int = 50052
    
    supported_exchanges: List[str] = ["binance", "kraken"]
    
    default_symbols: List[str] = [
        "BTC/USDT",
        "ETH/USDT",
        "BNB/USDT",
        "SOL/USDT",
        "XRP/USDT",
        "ADA/USDT",
        "DOGE/USDT",
        "AVAX/USDT",
        "DOT/USDT",
        "LINK/USDT",
    ]
    
    log_level: str = "INFO"
    
    class Config:
        env_file = ".env"
        env_file_encoding = "utf-8"


settings = Settings()
