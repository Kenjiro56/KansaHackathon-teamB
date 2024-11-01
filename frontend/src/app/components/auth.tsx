'use client';
import { useCallback, useState, FormEvent } from 'react';
import LoginInput from './input';
import Image from 'next/image';

const Auth = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');






  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  let flg = false;
  let message = '';

  const handleChange = (event: FormEvent<HTMLInputElement>) => {
    const { id, value } = event.currentTarget;
    setFormData({
      ...formData,
      [id]: value,
    });
  };

  const handleSubmit = async (e: FormEvent) =>{
    e.preventDefault();

  try {
    const response = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    });

    if (!response.ok) {
      // レスポンスがエラーの場合
      throw new Error('Login failed');
    }

    const jsondata = await response.json();

    if ("token" in jsondata) {
      localStorage.setItem('token', jsondata.token);
      alert('ログイン成功');
    } else {
      alert('トークンが含まれていません');
    }
  } catch (error) {
    alert('ログイン失敗');
  }
  };



  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="flex space-x-8 p-8 bg-white shadow-lg rounded-lg">
        <div>
          <Image
            src="/sample.png"
            alt = "description"
            width={256}
            height={256}
            className="rounded-lg bg-gray-300"
          />
        </div>
        <div className="flex flex-col space-y-4">
          <form onSubmit={handleSubmit}>
            <LoginInput
              id = "email"
              onChange = {handleChange}
              type = "email"
              label = "Email"
              value = {formData.email}
            />
            <LoginInput
              id = "password"
              onChange = {handleChange}
              type = "password"
              label = "password"
              value = {formData.password}
            />
            <button
              className="w-full py-2 text-white bg-teal-400 rounded-full hover:bg-teal-500"
            >
              ログイン
            </button>
          </form>
            <p
            className="mt-2 text-center text-sm text-gray-500 cursor-pointer"
            onClick={() => console.log('新規登録')}
            >
            新規会員登録はこちら
            </p>


        </div>
      </div>
    </div>
  );
}

export default Auth;
