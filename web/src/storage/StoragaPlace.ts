import { StoragePlace } from "@api/storagePlace";
import { number, object, string } from "yup";

export const schema = object({
  title: string().max(250).required(),
  storagePlaceId: number().optional(),
  code: string().max(50).nullable(),
});

export const createStoragePlaceOptions = (places: StoragePlace[] = []) => {
  return places.map((p) => {
    return {
      value: p.storagePlaceId,
      label: p.title,
    };
  });
};
