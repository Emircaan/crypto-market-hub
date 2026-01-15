import ccxt
from .base import BaseExchange


class KrakenExchange(BaseExchange):
    """Kraken exchange implementation."""
    
    @property
    def name(self) -> str:
        return "kraken"
    
    def _create_client(self) -> ccxt.Exchange:
        config = {
            "enableRateLimit": True,
        }
        
        if self.credentials and self.credentials.api_key:
            config["apiKey"] = self.credentials.api_key
            config["secret"] = self.credentials.api_secret
        
        return ccxt.kraken(config)
