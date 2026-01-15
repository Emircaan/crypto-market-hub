from abc import ABC, abstractmethod
from typing import List, Optional
from dataclasses import dataclass
import ccxt


@dataclass
class TickerData:
    """Ticker data from an exchange."""
    exchange: str
    symbol: str
    price: float
    volume: float
    high: float
    low: float
    change_percent: float
    timestamp: int


@dataclass
class Credentials:
    """Exchange API credentials."""
    api_key: str = ""
    api_secret: str = ""
    passphrase: str = ""


class BaseExchange(ABC):
    """Abstract base class for exchange implementations."""
    
    def __init__(self, credentials: Optional[Credentials] = None):
        self.credentials = credentials
        self._client: Optional[ccxt.Exchange] = None
    
    @property
    @abstractmethod
    def name(self) -> str:
        """Exchange name identifier."""
        pass
    
    @abstractmethod
    def _create_client(self) -> ccxt.Exchange:
        """Create the CCXT exchange client."""
        pass
    
    @property
    def client(self) -> ccxt.Exchange:
        """Get or create the CCXT client."""
        if self._client is None:
            self._client = self._create_client()
        return self._client
    
    def fetch_ticker(self, symbol: str) -> TickerData:
        """Fetch ticker data for a single symbol."""
        try:
            ticker = self.client.fetch_ticker(symbol)
            
            return TickerData(
                exchange=self.name,
                symbol=symbol,
                price=ticker.get("last", 0.0) or 0.0,
                volume=ticker.get("quoteVolume", 0.0) or 0.0,
                high=ticker.get("high", 0.0) or 0.0,
                low=ticker.get("low", 0.0) or 0.0,
                change_percent=ticker.get("percentage", 0.0) or 0.0,
                timestamp=ticker.get("timestamp", 0) or 0,
            )
        except Exception as e:
            raise Exception(f"Failed to fetch ticker for {symbol} on {self.name}: {e}")
    
    def fetch_tickers(self, symbols: List[str]) -> List[TickerData]:
        """Fetch ticker data for multiple symbols."""
        results = []
        errors = []
        
        for symbol in symbols:
            try:
                ticker = self.fetch_ticker(symbol)
                results.append(ticker)
            except Exception as e:
                errors.append(str(e))
        
        if errors and not results:
            raise Exception(f"All fetches failed: {'; '.join(errors)}")
        
        return results
