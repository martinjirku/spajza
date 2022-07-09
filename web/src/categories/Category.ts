import { QuantityType, Unit } from "@api/unit";
import { QSelectProps } from "quasar";
import { InferType, number, object, string } from "yup";

export const createUnits = (units: Unit[] = []): QSelectProps["options"] => {
  return units.map((a) => {
    return {
      label: `${[a.name]} (${a.symbol})`,
      value: a.name,
      icon: ((key: QuantityType) => {
        if (key === "mass") return "scale";
        if (key === "volume") return "takeout_dining";
        if (key === "time") return "timer";
        if (key === "count") return "tag";
        if (key === "length") return "straighten";
        if (key === "temperature") return "thermostat";
        return "square_foot";
      })(a.quantity),
    };
  });
};

export const schema = object({
  title: string().max(250).required(),
  id: number().optional(),
  path: string().max(250).optional(),
  defaultUnit: string().max(250).required(),
});

// @ts-ignore
window.schema = schema;

export type CategoryFormState = InferType<typeof schema>;
