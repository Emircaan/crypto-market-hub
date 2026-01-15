import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'Crypto Exchange',
  description: 'Real-time crypto market data',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className="antialiased min-h-screen bg-black text-gray-100 flex flex-col items-center">
        {children}
      </body>
    </html>
  )
}
