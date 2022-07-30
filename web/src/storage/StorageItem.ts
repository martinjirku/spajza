import { quantities } from "@api/unit";
import { date, number, object, string } from "yup";
export const schema = object({
  categoryId: number()
    .transform((v) => (Number.isNaN(v) ? undefined : v))
    .required(),
  title: string().max(255).required(),
  storagePlaceId: number()
    .transform((v) => (Number.isNaN(v) ? undefined : v))
    .optional(),
  amount: number()
    .transform((v) => (Number.isNaN(v) ? undefined : v))
    .required(),
  unit: string().required(),
  expirationDate: date()
    .nullable()
    .default(undefined)
    .transform((v) => {
      return v === "" ? undefined : v;
    })
    .optional(),
});

export type NewStorageItem = {
  categoryId: number;
  title: string;
  storagePlaceId: number;
  amount: number;
  unit: string;
  expirationDate: Date;
};
