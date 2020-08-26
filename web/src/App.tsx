import React, { useState } from "react";
import { initInstances } from "./api";
import Header from "./components/layouts/Header";
import MainLayout from "./components/layouts/MainLayout";

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
    <React.Fragment>
      <MainLayout />
    </React.Fragment>
  );
};

export default App;
