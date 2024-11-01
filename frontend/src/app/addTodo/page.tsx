'use client';
// app/setTodo/page.tsx

import React from 'react';
import SearchBar from './components/searchBar';
import { useRouter } from 'next/navigation'; // 修正点

const AddTodoPage: React.FC = () => {
  const router = useRouter();
  const handleBack = () => {
    // router.push('/home');
    router.push('/setTarget'); // 修正点
  }
  return (
    <div className='min-h-screen bg-gray-500 h-screen flex flex-col items-center justify-center space-y-20'>
      <h1 className='md:text-4xl lg:text-5xl font-bold text-white'>なにを達成したい？</h1>
      <SearchBar />
      <h2 className='md:text-4xl lg:text-3xl font-bold text-white' onClick={handleBack}>もどる</h2>
    </div>
  );
};

export default AddTodoPage;
