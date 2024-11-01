'use client'
// components/AddCard.tsx
import React from 'react';

type AddCardProps = {
  onClick?: () => void; // オプショナルな onClick プロパティ
};

const AddCard: React.FC<AddCardProps> = ({ onClick }) => {
  return (
    <div className="fixed bottom-10 right-10">
      <button onClick={onClick} className="bg-orange-400 text-white font-semibold py-3 px-6 rounded-full shadow-md">
        ＋ 目標を設定する！
      </button>
    </div>
  );
};

export default AddCard;
