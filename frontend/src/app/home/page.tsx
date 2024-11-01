import React from 'react'
import ObjCard from '@/app/home/components/card'

const page = () => {
  const cardData = [
    { id: 1, title: '目標1を達成したい！' },
    { id: 2, title: '目標2を達成したい！' },
    { id: 3, title: '目標3を達成したい！' },
    { id: 4, title: '目標4を達成したい！' },
    { id: 5, title: '目標5を達成したい！' },
    { id: 6, title: '目標6を達成したい！' },
    { id: 7, title: '目標7を達成したい！' },
    // { id: 8, title: '目標8を達成したい！' },
  ];

  return (
    <div className="min-h-screen bg-gray-200 flex flex-col items-center py-8">
      {/* タイトルとフィルタ */}
      <div className="text-center mb-6">
        <h1 className="text-lg font-semibold text-gray-700">すべて　未完了のみ　完了済み</h1>
      </div>

      {/* カードを並べるグリッド */}
      <div className="grid grid-cols-3 gap-6 max-w-5xl">
        {cardData.map((card) => (
          <ObjCard key={card.id} title={card.title} />
        ))}
        {/* 新しい目標を追加するためのボタン */}
        <div className="flex justify-center items-center bg-gray-100 rounded-3xl h-32 w-64 border border-gray-300">
          <span className="text-4xl text-gray-400">+</span>
        </div>
      </div>

      {/* 目標を設定するボタン */}
      <div className="fixed bottom-10 right-10">
        <button className="bg-orange-400 text-white font-semibold py-3 px-6 rounded-full shadow-md">
          ＋ 目標を設定する！
        </button>
      </div>
    </div>
  )
}

export default page
