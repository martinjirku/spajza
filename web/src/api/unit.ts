const tuple = <T extends string[]>(...args: T) => args;
export const quantities = tuple(
  "mass",
  "length",
  "volume",
  "temperature",
  "time",
  "count"
);

export type QuantityType = typeof quantities[number];
export type Unit = {
  name: string;
  names: string[];
  pluralName: string;
  quantity: QuantityType;
  symbol: string;
  system: string;
};
