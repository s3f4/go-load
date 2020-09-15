import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import SpinUpForm from "../components/forms/SpinUpForm";

interface Props {}

const Instances: React.FC<Props> = (props: Props) => {

  return (
    <React.Fragment>
    <MainLayout content={<SpinUpForm />} />
    </React.Fragment>
  );
};

export default Instances;
