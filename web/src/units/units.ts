import { QuantityType, Unit } from "@api/unit";
import { QSelectProps } from "quasar";

export const createUnitOptions = (
  units: Unit[] = []
): QSelectProps["options"] => {
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
