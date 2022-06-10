import { setLocale } from "yup";

setLocale({
  mixed: {
    default: "Zadajte v správnom tvare",
    required: "Toto je povinné pole",
  },
  string: {
    // min: ({ min }) => ({ key: "field_too_short", values: { min }, }),
    // max: ({ max }) => ({ key: "field_too_big", values: { max } }),
    min: ({ min }) => `Minimálne ${min} znaky`,
    max: ({ max }) => `Minimálne ${max} znakov`,
    email: () => "Zadajte email v správnom formáte",
  },
});
