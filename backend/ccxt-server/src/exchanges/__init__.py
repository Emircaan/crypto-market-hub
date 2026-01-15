from .base import BaseExchange, Credentials, TickerData
from .binance import BinanceExchange
from .kraken import KrakenExchange

__all__ = [
    "BaseExchange",
    "Credentials", 
    "TickerData",
    "BinanceExchange",
    "KrakenExchange",
]


def get_exchange(name: str, credentials: Credentials = None) -> BaseExchange:
    """Factory function to get an exchange instance by name."""
    exchanges = {
        "binance": BinanceExchange,
        "kraken": KrakenExchange,
    }
    
    if name.lower() not in exchanges:
        raise ValueError(f"Unsupported exchange: {name}. Supported: {list(exchanges.keys())}")
    
    return exchanges[name.lower()](credentials)
