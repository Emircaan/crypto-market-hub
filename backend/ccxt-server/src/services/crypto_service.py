import logging
from typing import List

from src.generated import crypto_market_pb2 as pb2
from src.generated import crypto_market_pb2_grpc as pb2_grpc
from src.exchanges import get_exchange, Credentials, TickerData
from src.config import settings


logger = logging.getLogger(__name__)


class CryptoMarketServicer(pb2_grpc.CryptoMarketServiceServicer):
    
    def FetchTickers(self, request, context):
        """Fetch ticker data from the specified exchange."""
        exchange_name = request.exchange.lower()
        symbols = list(request.symbols) if request.symbols else settings.default_symbols
        
        logger.info(f"FetchTickers request: exchange={exchange_name}, symbols={symbols}")
        
        try:
            credentials = None
            if request.HasField("credentials"):
                credentials = Credentials(
                    api_key=request.credentials.api_key,
                    api_secret=request.credentials.api_secret,
                    passphrase=request.credentials.passphrase,
                )
            
            exchange = get_exchange(exchange_name, credentials)
            tickers = exchange.fetch_tickers(symbols)
            
            proto_tickers = [
                pb2.Ticker(
                    exchange=t.exchange,
                    symbol=t.symbol,
                    price=t.price,
                    volume=t.volume,
                    high=t.high,
                    low=t.low,
                    change_percent=t.change_percent,
                    timestamp=t.timestamp,
                )
                for t in tickers
            ]
            
            logger.info(f"Successfully fetched {len(proto_tickers)} tickers from {exchange_name}")
            
            return pb2.FetchTickersResponse(
                success=True,
                error_message="",
                tickers=proto_tickers,
            )
            
        except ValueError as e:
            logger.warning(f"Invalid exchange: {e}")
            return pb2.FetchTickersResponse(
                success=False,
                error_message=str(e),
                tickers=[],
            )
        except Exception as e:
            logger.error(f"Error fetching tickers: {e}")
            return pb2.FetchTickersResponse(
                success=False,
                error_message=str(e),
                tickers=[],
            )
    
    def GetSupportedExchanges(self, request, context):
        """Return list of supported exchanges."""
        logger.info("GetSupportedExchanges request")
        return pb2.SupportedExchangesResponse(
            exchanges=settings.supported_exchanges
        )
    
    def HealthCheck(self, request, context):
        """Health check endpoint."""
        return pb2.HealthCheckResponse(
            healthy=True,
            version="1.0.0"
        )
