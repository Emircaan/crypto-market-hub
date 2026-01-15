'use client';

import { useState, useEffect } from 'react';

// Types
interface Ticker {
  exchange: string;
  symbol: string;
  price: number;
  change_percent: number;
  volume: number;
}

interface ApiResponse {
  success: boolean;
  data: Ticker[];
}

export default function Home() {
  const [binanceTickers, setBinanceTickers] = useState<Ticker[]>([]);
  const [krakenTickers, setKrakenTickers] = useState<Ticker[]>([]);
  const [loading, setLoading] = useState(true);

  const symbols = 'BTC/USDT,ETH/USDT,SOL/USDT,XRP/USDT,ADA/USDT,DOGE/USDT,AVAX/USDT,DOT/USDT,MATIC/USDT,LINK/USDT';

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const binanceRes = await fetch(`http://localhost:3000/api/v1/tickers/binance?symbols=${symbols}`);
        const binanceData: ApiResponse = await binanceRes.json();

        const krakenRes = await fetch(`http://localhost:3000/api/v1/tickers/kraken?symbols=${symbols}`);
        const krakenData: ApiResponse = await krakenRes.json();

        if (binanceData.success) setBinanceTickers(binanceData.data || []);
        if (krakenData.success) setKrakenTickers(krakenData.data || []);
      } catch (error) {
        console.error('Error fetching data:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();

    const interval = setInterval(fetchData, 30000);
    return () => clearInterval(interval);
  }, []);

  return (
    <main className="w-full max-w-7xl p-8">
      <h1 className="text-4xl font-bold mb-12 text-center text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-purple-500 tracking-tighter uppercase font-mono border-b border-gray-800 pb-8">
        Crypto Market Data
      </h1>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-12">
        {/* Binance Section */}
        <section className="bg-gray-900 border border-gray-800 rounded-lg overflow-hidden shadow-2xl">
          <div className="bg-gray-800/50 p-4 border-b border-gray-700 flex justify-between items-center">
            <h2 className="text-xl font-mono font-bold text-yellow-500">BINANCE</h2>
            <span className="text-xs text-gray-500 font-mono">LIVE FEED</span>
          </div>
          <div className="overflow-x-auto">
            <table className="w-full text-left font-mono text-sm">
              <thead className="bg-gray-900 text-gray-400 border-b border-gray-800">
                <tr>
                  <th className="p-4">SYMBOL</th>
                  <th className="p-4 text-right">PRICE</th>
                  <th className="p-4 text-right">24H %</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-800">
                {loading && binanceTickers.length === 0 ? (
                  <tr><td colSpan={3} className="p-4 text-center text-gray-500">Loading stream...</td></tr>
                ) : (
                  binanceTickers.map((t, i) => (
                    <tr key={i} className="hover:bg-gray-800/30 transition-colors">
                      <td className="p-4 font-bold text-gray-300">{t.symbol}</td>
                      <td className="p-4 text-right text-white">${t.price.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 4 })}</td>
                      <td className={`p-4 text-right font-bold ${t.change_percent >= 0 ? 'text-green-500' : 'text-red-500'}`}>
                        {t.change_percent > 0 ? '+' : ''}{t.change_percent.toFixed(2)}%
                      </td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          </div>
        </section>

        {/* Kraken Section */}
        <section className="bg-gray-900 border border-gray-800 rounded-lg overflow-hidden shadow-2xl">
          <div className="bg-gray-800/50 p-4 border-b border-gray-700 flex justify-between items-center">
            <h2 className="text-xl font-mono font-bold text-purple-500">KRAKEN</h2>
            <span className="text-xs text-gray-500 font-mono">LIVE FEED</span>
          </div>
          <div className="overflow-x-auto">
            <table className="w-full text-left font-mono text-sm">
              <thead className="bg-gray-900 text-gray-400 border-b border-gray-800">
                <tr>
                  <th className="p-4">SYMBOL</th>
                  <th className="p-4 text-right">PRICE</th>
                  <th className="p-4 text-right">24H %</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-800">
                {loading && krakenTickers.length === 0 ? (
                  <tr><td colSpan={3} className="p-4 text-center text-gray-500">Loading stream...</td></tr>
                ) : (
                  krakenTickers.map((t, i) => (
                    <tr key={i} className="hover:bg-gray-800/30 transition-colors">
                      <td className="p-4 font-bold text-gray-300">{t.symbol}</td>
                      <td className="p-4 text-right text-white">${t.price.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 4 })}</td>
                      <td className={`p-4 text-right font-bold ${t.change_percent >= 0 ? 'text-green-500' : 'text-red-500'}`}>
                        {t.change_percent > 0 ? '+' : ''}{t.change_percent.toFixed(2)}%
                      </td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          </div>
        </section>
      </div>
    </main>
  );
}
