import { QuantityType } from "./unit";

export type Consumption = {
  amount: number;
  unit: string;
};

export type StorageItem = {
  storageItemId?: number;
  title?: string;
  baselineAmount?: number;
  currentAmount?: number;
  categoryId?: number;
  storagePlaceId?: number;
  quantityType?: QuantityType;
  unit?: string;
  expirationDate?: string;
  consumptions?: Consumption[];
};

export type StorageItemListResponse = {
  data: StorageItem[];
  count: number;
};

export type NewStorageItem = {
  categoryId: number;
  title: string;
  storagePlaceId: number;
  amount: number;
  unit: string;
  expirationDate?: Date;
};

export type ConsumptionRequest = Consumption;
