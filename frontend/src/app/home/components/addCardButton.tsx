'use client'
// components/AddCardButton.tsx
import React from 'react';

type AddCardButtonProps = {
  onClick?: () => void; // オプショナルな onClick プロパティ
};

const AddCardButton: React.FC<AddCardButtonProps> = ({ onClick }) => {
  return (
    <div onClick={onClick} className="flex justify-center items-center bg-gray-100 rounded-3xl h-32 w-64 border border-gray-300 cursor-pointer">
      <span className="text-4xl text-gray-400">+</span>
    </div>
  );
};

export default AddCardButton;
