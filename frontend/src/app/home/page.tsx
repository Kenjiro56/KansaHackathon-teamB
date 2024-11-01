'use client';
import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'; // 修正点
import ObjCard from '@/app/home/components/card'
import AddCard from '@/app/home/components/addCard'
import AddCardButton from '@/app/home/components/addCardButton';

const page = () => {
  const dummyData = [
    { id: 1, title: '目標1を達成したい！' },
    { id: 2, title: '目標2を達成したい！' },
    { id: 3, title: '目標3を達成したい！' },
    { id: 4, title: '目標4を達成したい！' },
    { id: 5, title: '目標5を達成したい！' },
    { id: 6, title: '目標6を達成したい！' },
    { id: 7, title: '目標7を達成したい！' },
    { id: 8, title: '目標8を達成したい！' },
    { id: 8, title: '目標8を達成したい！' },
    { id: 8, title: '目標8を達成したい！' },
    { id: 8, title: '目標8を達成したい！' },
    { id: 8, title: '目標8を達成したい！' },
    { id: 8, title: '目標8を達成したい！' },
    { id: 8, title: '目標8を達成したい！' },
    { id: 8, title: '目標8を達成したい！' },
    { id: 8, title: '目標8を達成したい！' },
  ];

  const router = useRouter();
  const [cardData, setCardData] = useState<{ id: number; title: string }[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/obj/getAll'); // APIエンドポイントに合わせて変更してください
        if (!response.ok) throw new Error('Failed to fetch data');

        const data = await response.json();
        //ここから↓でデータをcardDataにformattingしてる
        const formattedData = data.map((item: any) => ({
          id: item.id,
          title: item.obj_title,
        }));
        setCardData(formattedData);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };
    fetchData(); // 関数を呼び出してデータフェッチを行う
  }, []); // 空の依存配列により、コンポーネントがマウントされたときのみ実行されます

  const handleNavigate = () => {
    router.push('/addTodo');
  };

  return (
    <div className="min-h-screen bg-gray-200 flex flex-col items-center py-8">
      <div className="text-center mb-6">
        <h1 className="text-lg font-semibold text-gray-700">すべて　未完了のみ　完了済み</h1>
      </div>

      {/* カード並べてる */}
      <div className="grid grid-cols-3 gap-6 max-w-5xl">
        {cardData.map((card) => (
          <ObjCard key={card.id} title={card.title} />
        ))}
        {/* 新しい目標を追加するボタン */}
        <AddCardButton onClick={handleNavigate}/>
      </div>

      {/* 目標を設定するボタン */}
        <AddCard onClick={handleNavigate}/>
    </div>
  )
}

export default page
