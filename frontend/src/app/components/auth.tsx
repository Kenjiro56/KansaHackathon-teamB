'use client';
import { useCallback, useState } from 'react';
import LoginInput from './loginInput';
import Image from 'next/image';

const Auth = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');


  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="flex space-x-8 p-8 bg-white shadow-md rounded-md">
        <div>
          <Image
            src="/sample.png"
            alt = "description"
            objectFit="cover"
            width={256}
            height={256}
          />
        </div>
        <div>
            <>
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
              className="w-full py-2 mt-4 text-white bg-green-300 rounded-md hover:bg-green-400"
            >ログイン</button>
            <button
              onClick = {() => console.log('新規登録')}
            >
              新規会員登録はこちら
            </button>
            </>
        </div>
      </div>
    </div>
  );
}

export default Auth;
