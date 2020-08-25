import React, { useState } from "react";
import { initInstances } from "./api";
import Header from "./components/layouts/Header";

const App: React.FC = () => {
  const [count, setCount] = useState<number>(0);
  const [region, setRegion] = useState<string>("AMS3");

  const handleChange = (e: any) => {
    switch (e.target.name) {
      case "count":
        setCount(parseInt(e.target.value));
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
    <div>
      <div className="container">
        <Header />
        <div className="container">
          <div className="row">
            <div className="col-sm">Form 1</div>
            <div className="col-sm">Form 2</div>
          </div>
          Node list
        </div>
      </div>
    </div>
  );
};

export default App;
