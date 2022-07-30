import { QuantityType } from "./unit";

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
};

export type NewStorageItem = {
  categoryId: number;
  title: string;
  storagePlaceId: number;
  amount: number;
  unit: string;
  expirationDate?: Date;
};
