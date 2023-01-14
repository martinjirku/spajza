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
  ean: string()
    .nullable()
    .max(20)
    .default(undefined)
    .transform((v) => {
      return v === "" ? undefined : v;
    })
    .optional(),
  expirationDate: date()
    .nullable()
    .default(undefined)
    .transform((v) => {
      return v === "" ? undefined : v;
    })
    .optional(),
});
