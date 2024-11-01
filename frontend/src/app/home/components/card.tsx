import React from 'react'

const card = () => {
  return (
    <div className="w-80 p-6 bg-gray-100 rounded-3xl shadow-lg relative">
      {/* トップバー */}
      <div className="absolute top-0 left-0 w-full h-10 bg-gray-300 rounded-t-3xl flex">
        <div className="h-full w-1/2 bg-teal-300 rounded-tl-3xl"></div>
      </div>

      {/* メインコンテンツ */}
      <div className="mt-10 mb-4 flex flex-col items-center">
        <div className="bg-white rounded-full py-4 px-6 mb-6 shadow-md w-full text-center">
          <p className="text-gray-700 font-semibold">~~を達成したい！</p>
        </div>

        <button className="bg-teal-300 text-gray-700 font-semibold rounded-full py-3 px-8 shadow-md">
          くわしくみる
        </button>
      </div>
    </div>
  )
}

export default card
