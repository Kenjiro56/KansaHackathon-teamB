// src/components/Header.tsx
import React from 'react';
import Link from 'next/link';
// import Image from 'next/image';
// import LogoIcon from '@/assets/logo.svg';
import LogoIcon from '@/assets/logo.svg'
import DropDown from '@/assets/dropDown.svg'


const Header: React.FC = () => {
  return (
    <header className="bg-white text-gray p-2 min-h-16">
      <nav className="flex justify-between items-center">
        <h1 className="text-lg font-bold">
          <Link href="/">
            <LogoIcon width={270} height={50} />
          </Link>
        </h1>
        <div className="space-x-4 flex mr-10 items-center">
          <h2 className='text-gray-600 font-bold'>ユーザー名</h2>
          <DropDown />
        </div>
      </nav>
    </header>
  );
};

export default Header;
