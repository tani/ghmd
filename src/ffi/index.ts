import { cc as compile } from "#ffi";

export interface NativeSymbol {
  args?: readonly (string | number)[];
  returns?: string | number;
}

export interface CCOptions {
  source: string | URL;
  include?: string | readonly string[];
  library?: string | readonly string[];
  define?: Record<string, string>;
  flags?: string | readonly string[];
  symbols: Record<string, NativeSymbol>;
}

export interface CCLibrary {
  symbols: Record<string, (...args: any[]) => any>;
  close(): void;
}

// package.json resolves #ffi for each runtime without loading the other backend.
export const cc = compile as (options: CCOptions) => CCLibrary;
