import { number, object, string } from "yup";

export const schema = object({
  title: string().max(250).required(),
  storagePlaceId: number().optional(),
  code: string().max(50).nullable(),
});
