import { Category } from "@api/category";
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
  path: string().max(250).nullable(),
  defaultUnit: string().max(250).required(),
});

export type CategoryFormState = InferType<typeof schema>;

export type ParentOption = { value: string; label: string };
export const createParentOptions = (
  categories: Category[] = [],
  indexedParents: Record<string, Category> = {}
) => {
  return (
    categories?.map((c) => {
      let label = [
        ...c.path
          .split(".")
          .filter((i) => i !== "")
          .map(Number),
        c.id,
      ]
        .map((key) => indexedParents[key]?.title)
        .join(" > ");

      return {
        value: c.path ? [c.path, c.id].join(".") : c.id.toString(),
        label,
      };
    }) ?? []
  );
};
