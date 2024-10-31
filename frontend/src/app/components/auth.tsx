import { useCallback, useState } from 'react';
import LoginInput from './loginInput';
import Image from 'next/image';

const Auth = () => {
  // ログインしているか否かの状態保持
  const [isLogin, setIsLogin] = useState(true);

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');


  return (
    <div className="">
      <div className="">
        <div className="">
          <Image
            src=""
            alt = ""
            layout="fill"
            objectFit="cover"
          />
        </div>

        <div className="">
          <LoginInput
            id = "email"
            onChange = {(event:any) => setEmail(event.target.value)}
            type = "email"
            label = "メールアドレス"
            value = {email}
          />
          <LoginInput
            id = "password"
            onChange = {(event:any) => setPassword(event.target.value)}
            type = "password"
            label = "パスワード"
            value = {password}
          />
        </div>
      </div>
    </div>
  );
}

export default Auth;
