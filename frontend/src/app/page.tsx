export default async function Home() {
  const response = await fetchMessage();

  return (
    <>
      {JSON.stringify(response, null, 2)}
    </>
  );
}

export async function fetchMessage() {
  const hoge = "test";
  try {
    console.log("データ取得中です...");
    const response = await fetch("http://backend:8080/test");
    const data = await response.json();
    return data;
  } catch (error) {
    console.log("データ取得エラー");
    return { error: "データ取得エラー" };
  }
}
