import { useEffect, useState } from "react";

function App() {
  const [message, setMessage] = useState("");
  const [error, setError] = useState(null);

  useEffect(() => {
    console.log("Fetching /api/hello...");
    fetch("/api/hello")
        .then((res) => {
          if (!res.ok) throw new Error("API error");
          return res.json();
        })
        .then((data) => setMessage(data.message))
        .catch((err) => {
          console.error("Fetch error:", err);
          setError("Failed to fetch message");
        });
  }, []);

  return (
      <div>
        <h1>{message || error}</h1>
      </div>
  );
}

export default App;
