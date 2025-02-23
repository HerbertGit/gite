import { useEffect, useState } from "react";
import "./App.css";
import { fetchContet } from "./fetcher";

export function DataComponent() {
  const [data, setData] = useState<undefined | Record<string, string>>(
    undefined
  );

  useEffect(() => {
    fetchContet().then((data) => {
      setData(data);
    });
  }, []);

  return (
    <>
      <p>{data ? JSON.stringify(data) : null}</p>
    </>
  );
}
