import React, { useState } from "react";

import "./App.css";
import { initInstances } from "./api";

const App: React.FC = () => {
  const [count, setCount] = useState(0);
  const [region, setRegion] = useState("AMS3");

  const handleChange = (e: any) => {
    switch (e.target.name) {
      case "count":
        setCount(e.target.value);
        break;
      case "region":
        setRegion(e.target.value);
        break;
    }
  };

  const sendRequest = (e: any) => {
    e.preventDefault();
    const instances = { count, region };
    initInstances(instances).then((data) => console.log(data));
  };

  return (
    <div className="App">
      <form>
        <input type="number" name="count" onChange={handleChange} />
        <input type="text" name="region" onChange={handleChange} />
        <button onClick={sendRequest}>Send</button>
      </form>
    </div>
  );
};

export default App;
