export interface RunConfig {
  requestCount: number;
  goroutineCount: number;
  url: string;
  transportConfig: TransportConfig,
}

export interface TransportConfig {
  DisableKeepAlives: boolean;
}