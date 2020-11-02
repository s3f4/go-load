import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import Message from "../components/basic/Message";

const NotFound: React.FC = () => {
  return (
    <React.Fragment>
      <MainLayout
        content={() => {
          return <Message message={"Page not found."} type="error" />;
        }}
      />
    </React.Fragment>
  );
};

export default NotFound;
