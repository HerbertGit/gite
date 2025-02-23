export async function fetchContet() {
  const res = await fetch("http://127.0.0.1:8080/api/test");
  return res.json();
}
