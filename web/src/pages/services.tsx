import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import ServicesContent from "../components/contents/ServicesContent";
import { list, Service } from "../api/entity/service";

const Services: React.FC = () => {
  const [services, setServices] = React.useState<Service[]>([]);
  const [loader, setLoader] = React.useState<boolean>(false);

  React.useEffect(() => {
    setLoader(true);
    list()
      .then((response) => {
        setServices(response.data.containers);
        setLoader(false);
      })
      .catch((err) => {
        setLoader(false);
        console.log(err);
      });
    return () => {};
  }, []);

  return (
    <React.Fragment>
      <MainLayout
        content={<ServicesContent loader={loader} services={services} />}
      />
    </React.Fragment>
  );
};

export default Services;
