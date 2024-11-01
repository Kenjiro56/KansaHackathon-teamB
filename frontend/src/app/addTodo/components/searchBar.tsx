'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation'; // 修正点

const SearchBar: React.FC = () => {
  const [inputValue, setInputValue] = useState('');
  const router = useRouter();

  // 入力が変更されたときに呼び出される関数
  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };

  // SVGがクリックされたときに呼び出される関数
  const handleSendClick = async () => {
    console.log('Input Value:', inputValue);
    // ここでinputValueを使って他の処理を実行できます
    try {
      const response = await fetch("http://localhost:8080/obj/add", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ // 送信するデータをここで設定
          user_id: 1,
          obj_title: inputValue,
        }),
      });

      if (response.ok) {
        const data = await response.json();
        console.log("Response from server:", data);
        router.push('/home');
      } else {
        console.error("Failed to send data:", response.statusText);
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <div className="flex items-center bg-white rounded-3xl shadow-md w-full max-w-6xl p-2 border border-gray-300">
      <input
        type="text"
        placeholder="〜を達成したい！！"
        className="flex-grow px-4 py-2 text-gray-400 text-base focus:outline-none font-bold"
        value={inputValue}
        onChange={handleInputChange}
      />
      <svg className='mr-5 cursor-pointer' onClick={handleSendClick} width="50" height="50" viewBox="0 0 54 54" fill="none" xmlns="http://www.w3.org/2000/svg">
        <rect x="0.663818" y="0.332031" width="53.3363" height="53.3363" rx="16" fill="#02F7C8"/>
        <path d="M28.9158 40.2865C28.962 40.4017 29.0424 40.5 29.1461 40.5682C29.2498 40.6364 29.3719 40.6713 29.496 40.6681C29.62 40.6649 29.7402 40.6238 29.8403 40.5504C29.9403 40.477 30.0155 40.3747 30.0558 40.2573L37.9642 17.1405C38.0031 17.0327 38.0106 16.916 37.9856 16.8042C37.9607 16.6923 37.9044 16.5898 37.8233 16.5088C37.7423 16.4277 37.6398 16.3714 37.5279 16.3465C37.4161 16.3216 37.2994 16.329 37.1916 16.3679L14.0748 24.2763C13.9574 24.3166 13.8551 24.3918 13.7817 24.4919C13.7083 24.5919 13.6672 24.7121 13.664 24.8362C13.6608 24.9602 13.6957 25.0823 13.7639 25.186C13.8321 25.2897 13.9304 25.3701 14.0456 25.4163L23.6938 29.2853C23.9988 29.4074 24.276 29.5901 24.5085 29.8222C24.741 30.0543 24.9241 30.3311 25.0468 30.6358L28.9158 40.2865Z" stroke="#585D72" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M37.8194 16.5142L24.509 29.8233" stroke="#585D72" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
      </svg>
    </div>
  );
};

export default SearchBar;
