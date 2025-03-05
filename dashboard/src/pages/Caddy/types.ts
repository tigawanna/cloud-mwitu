export interface ConfigItem {
  path?: string[];
  value?: string[];
}

export interface RecursiveConfigProps {
  configData: ConfigItem[];
  domain?: string;
}
