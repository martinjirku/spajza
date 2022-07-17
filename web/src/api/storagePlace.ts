export type StoragePlace = {
  storagePlaceId: number;
  title?: string;
  code: string;
};

export type CreateStoragePlace = {
  title?: string;
  code: string;
};

export const isStoragePlaace = (
  s: StoragePlace | CreateStoragePlace
): s is StoragePlace => {
  return (s as StoragePlace).storagePlaceId !== undefined;
};
