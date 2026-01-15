from typing import Optional
import ccxt
from .base import BaseExchange, Credentials


class BinanceExchange(BaseExchange):
    """Binance exchange implementation."""
    
    @property
    def name(self) -> str:
        return "binance"
    
    def _create_client(self) -> ccxt.Exchange:
        config = {
            "enableRateLimit": True,
        }
        
        if self.credentials and self.credentials.api_key:
            config["apiKey"] = self.credentials.api_key
            config["secret"] = self.credentials.api_secret
        
        return ccxt.binance(config)
