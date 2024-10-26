'use client'
import React, { useState } from 'react';

const ApiButtons: React.FC = () => {

    const [inputData, setInputData] = useState('');

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
      setInputData(event.target.value);
    };

    // データを参照するときのハンドラー
    const handleSelectClick = async () => {
      try {
        const response = await fetch('http://localhost:8080/select', { method: 'GET' });
        const data = await response.json();
        console.log(data); // レスポンスデータを確認
      } catch (error) {
          console.error('Select API error:', error);
      }
    };

    // 何かデータを差し込むときのボタンのクリックハンドラー
    const handleInsertClick = async () => {
        try {
          const response = await fetch('http://localhost:8080/insert', {
              method: 'POST',
              headers: { 'Content-Type': 'application/json' },
              body: JSON.stringify({ message: inputData }), // 入力されたデータを送信
          });
          const data = await response.json();
          console.log(data); // レスポンスデータを確認
      } catch (error) {
          console.error('Insert API error:', error);
      }
    };

    return (
      <>
        <input
            type="text"
            value={inputData}
            onChange={handleInputChange}
            placeholder="データを入力"
        /><br/>
        <button onClick={handleInsertClick}>Insert</button><br/>
        <button onClick={handleSelectClick}>Select</button>
      </>
    );
};

export default ApiButtons;
