import { makeReq } from "../api";
export interface Service {
  ID: string;
  CreatedAt: Date;
  UpdatedAt: Date;
  Spec: Spec;
}

export interface Spec {
  Name: string;
  Labels: any;
  EndpointSpec: EndpointSpec;
  Mode: Mode;
}

export interface EndpointSpec {
  Mode: string;
  Ports: Port[];
}

export interface Port {
  Protocol: string;
  TargetPort: number;
  PublishedPort: number;
  PublishMode: "ingress";
}

export interface Mode {
  Replicated: Replicated;
}

export interface Replicated {
  Replicas: string;
}

export const list = async () => {
  return await makeReq("/services");
};

export const stop = async (service: Service) => {
  return await makeReq("/services", "POST", service);
};
