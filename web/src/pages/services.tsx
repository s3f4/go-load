import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import ServicesContent from "../components/contents/ServicesContent";
import { list, Service } from "../api/entity/service";
import { IMessage } from "../components/basic/Message";

const Services: React.FC = () => {
  const [services, setServices] = React.useState<Service[]>([]);
  const [loader, setLoader] = React.useState<boolean>(false);
  const [message, setMessage] = React.useState<IMessage>();

  React.useEffect(() => {
    setLoader(true);
    list()
      .then((response) => {
        setServices(response.data.containers);
        setLoader(false);
      })
      .catch((err) => {
        setLoader(false);
        setMessage({
          type: "error",
          message: err.message,
        });
      });
    return () => {};
  }, []);

  return (
    <React.Fragment>
      <MainLayout
        content={
          <ServicesContent
            loader={loader}
            services={services}
            message={message}
          />
        }
      />
    </React.Fragment>
  );
};

export default Services;
