import { Category } from "./category";
import { Unit } from "./unit";

export const getCategories = () => {
  return fetch("/api/categories")
    .then((resp) => {
      return resp.json();
    })
    .then((data) => data as Category[]);
};

export const getUnits = () => {
  return fetch("api/units")
    .then((resp) => resp.json())
    .then((data) => data as Unit[]);
};
