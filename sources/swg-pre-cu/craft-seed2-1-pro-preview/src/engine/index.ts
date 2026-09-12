// Engine barrel export — the single import site for the UI layer.
// When porting to UE5, this entire directory becomes a C++ module; the React
// layer is discarded.
export * from "./types";
export * from "./rng";
export * from "./resourceGen";
export * from "./crafting";
export { RESOURCE_CLASSES, getClass, isSubtypeOf, getRootClass } from "./data/resourceTree";
export { SCHEMATICS, SCHEMATIC_LIST } from "./schematics";
