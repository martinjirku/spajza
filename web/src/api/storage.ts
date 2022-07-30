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
  expiration_date?: string;
};
