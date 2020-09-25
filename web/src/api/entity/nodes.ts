export interface Node {
  ID: string;
  CreatedAt: Date;
  Spec: Spec;
  Status: Status;
  Description: Description;
}

interface Spec {
  Availability: string;
  Labels: string[];
  Role: string;
}

interface Status {
  State: string;
  Addr: string;
}

interface Description {
  Hostname: string;
  Platform: Platform;
  Resources: Resources;
}

interface Platform {
  Architecture: string;
  OS: string;
}

interface Resources {
  MemoryBytes: number;
  NanoCPUs: number;
}
