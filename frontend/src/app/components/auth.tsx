'use client';
import { useCallback, useState } from 'react';
import LoginInput from './input';
import Image from 'next/image';

const Auth = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');


  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="flex space-x-8 p-8 bg-white shadow-lg rounded-lg">
        <div>
          <Image
            src="/sample.png"
            alt = "description"
            objectFit="cover"
            width={256}
            height={256}
            className="rounded-lg bg-gray-300"
          />
        </div>
        <div className="flex flex-col space-y-4">
            <LoginInput
              id = "email"
              onChange = {(event:any) => setEmail(event.target.value)}
              type = "email"
              label = "Email"
              value = {email}
            />
            <LoginInput
              id = "password"
              onChange = {(event:any) => setPassword(event.target.value)}
              type = "password"
              label = "password"
              value = {password}
            />
            <button
              onClick = {() => {console.log(email);}}
              className="w-full py-2 text-white bg-teal-400 rounded-full hover:bg-teal-500"
            >
              ログイン
            </button>
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
