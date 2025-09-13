import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'タスク管理アプリ',
  description: 'シンプルなタスク管理アプリケーション',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="ja">
      <body className="bg-gray-100 min-h-screen">
        <div className="container mx-auto px-4 py-8">
          {children}
        </div>
      </body>
    </html>
  )
}