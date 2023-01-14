import { setLocale } from "yup";

export const initLocale = () => {
  setLocale({
    mixed: {
      default: "Zadajte v správnom tvare",
      required: "Toto je povinné pole",
    },
    string: {
      min: ({ min }) => `Minimálne ${min} znaky`,
      max: ({ max }) => `Minimálne ${max} znakov`,
      email: () => "Zadajte email v správnom formáte",
    },
  });
};

initLocale();
