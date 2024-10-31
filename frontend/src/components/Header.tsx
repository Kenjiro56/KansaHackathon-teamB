// 'use client';
// import React, {useState} from 'react';
// import Link from 'next/link';
// import LogoIcon from '@/assets/logo.svg'
// import DropDown from '@/assets/dropDown.svg'
// import DropMenu from './DropMenu'



// const Header: React.FC = () => {
//   const [isDropMenuOpen, setIsDropMenuOpen] = useState(false);

//   // メニューをトグルする関数
//   const toggleDropMenu = () => {
//     setIsDropMenuOpen(!isDropMenuOpen);
//   };
//   return (
//     <header className="bg-white text-gray p-2 min-h-16">
//       <nav className="flex justify-between items-center">
//         <h1 className="text-lg font-bold">
//           <Link href="/">
//             <LogoIcon width={270} height={50} />
//           </Link>
//         </h1>
//         <div className="space-x-4 flex mr-10 items-center" onClick={toggleDropMenu}>
//           <h2 className='text-gray-600 font-bold'>ユーザー名</h2>
//           <DropDown />
//         </div>
//         {isDropMenuOpen && ( // メニューが開いているときのみ表示
//           <div className="absolute top-full right-0 mt-2">
//             <DropMenu />
//           </div>
//         )}
//       </nav>
//     </header>
//   );
// };

// export default Header;

'use client';
import React, { useState, useEffect, useRef } from 'react';
import Link from 'next/link';
import LogoIcon from '@/assets/logo.svg';
import DropDown from '@/assets/dropDown.svg';
import DropMenu from './DropMenu';

const Header: React.FC = () => {
  const [isDropMenuOpen, setIsDropMenuOpen] = useState(false);
  const menuRef = useRef<HTMLDivElement>(null);

  // メニューをトグルする関数
  const toggleDropMenu = () => {
    setIsDropMenuOpen((prev) => !prev);
  };

  // メニュー外をクリックしたときに閉じる
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (menuRef.current && !menuRef.current.contains(event.target as Node)) {
        setIsDropMenuOpen(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <header className="bg-white text-gray p-2 min-h-16">
      <nav className="flex justify-between items-center">
        <h1 className="text-lg font-bold">
          <Link href="/">
            <LogoIcon width={270} height={50} />
          </Link>
        </h1>
        <div className="relative" ref={menuRef}>
          <div
            className="space-x-4 flex mr-10 items-center cursor-pointer"
            onClick={toggleDropMenu} // クリック時にメニューの表示/非表示をトグル
          >
            <h2 className="text-gray-600 font-bold cursor-pointer">ユーザー名</h2>
            <DropDown />
          </div>
          {isDropMenuOpen && ( // メニューが開いているときのみ表示
            <div className="absolute top-full right-0 mt-2 transform translate-y-4 mr-7 cursor-pointer">
              <DropMenu />
            </div>
          )}
        </div>
      </nav>
    </header>
  );
};

export default Header;
