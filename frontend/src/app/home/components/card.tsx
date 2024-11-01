import React from 'react'

type CardProps = {
  title: string;
};

const Card: React.FC<CardProps> = ({ title, onClick }) => {
  return (
    <div className="w-64 p-6 bg-gray-100 rounded-3xl shadow-lg relative">
      <div className="absolute top-0 left-0 w-full h-10 bg-gray-300 rounded-t-3xl flex">
        <div className="h-full w-1/2 bg-teal-300 rounded-tl-3xl"></div>
      </div>

      <div className="mt-10 mb-4 flex flex-col items-center">
        <div className="bg-white rounded-full py-4 px-6 mb-6 shadow-md w-full text-center">
          <p className="text-gray-700 font-semibold ellipsis">{title}</p>
        </div>

        <button className="bg-teal-300 text-gray-700 font-semibold rounded-full py-3 px-8 shadow-md" onClick={onClick}>
          くわしくみる
        </button>
      </div>
    </div>
  )
}

export default Card
