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

export const createCategory = (category: Category) => {
  return fetch("/api/categories", {
    method: "POST",
    headers: { Accept: "application/json", "Content-Type": "application/json" },
    body: JSON.stringify(category),
  })
    .then((resp) => resp.json())
    .then((data) => data as Category);
};

export const updateCategory = (category: Category) => {
  return fetch(`/api/categories/${category.id}`, {
    method: "POST",
    headers: { Accept: "application/json", "Content-Type": "application/json" },
    body: JSON.stringify(category),
  })
    .then((resp) => resp.json())
    .then((data) => data as Category);
};
