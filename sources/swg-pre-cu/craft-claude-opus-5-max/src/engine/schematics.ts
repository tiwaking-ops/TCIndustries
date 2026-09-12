/**
 * Schematic domain.
 *
 * The critical structure is ExperimentalProperty.groups: a property is a
 * WEIGHTED SUM OVER (slot, stat) pairs. This is the single data shape that
 * makes SWG crafting work, and it is why a nominal/OO model is the wrong
 * tool: the recipe is a sparse matrix, not an object graph.
 *
 *     quality(prop) = Σ_g  w_g * ( Σ_s w_s * stat(resource[g.slot], s) / Σ_s w_s ) / Σ_g w_g
 *
 * Everything downstream (assembly, experimentation, final item) is a scalar
 * transform of that one number per property, in [0,1].
 */

import type { StatKey } from './resources';

export interface SchematicSlot {
  id: string;
  label: string;
  /** any class in this subtree satisfies the slot */
  classId: string;
  units: number;
}

export interface StatWeight {
  stat: StatKey;
  weight: number;
}

export interface PropertyGroup {
  slotId: string;
  weight: number;
  stats: StatWeight[];
}

export interface ExperimentalProperty {
  id: string;
  label: string;
  unit: string;
  min: number;
  max: number;
  /** true => higher quality yields a LOWER number (mass, encumbrance) */
  invert?: boolean;
  decimals?: number;
  groups: PropertyGroup[];
}

/** One experimentation line can drive several properties. Points are scarce. */
export interface ExperimentLine {
  id: string;
  label: string;
  properties: string[];
}

export interface Schematic {
  id: string;
  name: string;
  category: string;
  /** drives assembly difficulty and the experimentation point budget */
  complexity: number;
  xpType: string;
  slots: SchematicSlot[];
  properties: ExperimentalProperty[];
  lines: ExperimentLine[];
}

export const SCHEMATICS: readonly Schematic[] = [
  {
    id: 'blaster_power_handler',
    name: 'Blaster Power Handler, Small',
    category: 'Weapon Component',
    complexity: 12,
    xpType: 'Weaponsmithing',
    slots: [
      { id: 'chassis', label: 'Conductive Housing', classId: 'nonferrous', units: 20 },
      { id: 'cell', label: 'Energy Substrate', classId: 'chemical', units: 12 },
      { id: 'shield', label: 'Shielding Plate', classId: 'ferrous', units: 8 },
    ],
    properties: [
      {
        id: 'throughput', label: 'Power Throughput', unit: 'MW', min: 8, max: 34, decimals: 1,
        groups: [
          { slotId: 'chassis', weight: 2, stats: [{ stat: 'CD', weight: 3 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'cell', weight: 1, stats: [{ stat: 'PE', weight: 2 }, { stat: 'OQ', weight: 1 }] },
        ],
      },
      {
        id: 'integrity', label: 'Thermal Integrity', unit: 'pts', min: 120, max: 640,
        groups: [
          { slotId: 'shield', weight: 2, stats: [{ stat: 'HR', weight: 2 }, { stat: 'SR', weight: 1 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'chassis', weight: 1, stats: [{ stat: 'DR', weight: 1 }, { stat: 'OQ', weight: 1 }] },
        ],
      },
      {
        id: 'mass', label: 'Component Mass', unit: 'kg', min: 1.2, max: 6.4, invert: true, decimals: 2,
        groups: [
          { slotId: 'chassis', weight: 1, stats: [{ stat: 'MA', weight: 2 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'shield', weight: 1, stats: [{ stat: 'MA', weight: 1 }] },
        ],
      },
    ],
    lines: [
      { id: 'exp_power', label: 'Experimental Power', properties: ['throughput'] },
      { id: 'exp_durability', label: 'Experimental Durability', properties: ['integrity', 'mass'] },
    ],
  },

  {
    id: 'shield_generator',
    name: 'Personal Shield Generator, Adv.',
    category: 'Wearable',
    complexity: 22,
    xpType: 'Engineering',
    slots: [
      { id: 'frame', label: 'Structural Frame', classId: 'steel', units: 40 },
      { id: 'core', label: 'Reactant Core', classId: 'radioactive', units: 25 },
      { id: 'buffer', label: 'Inert Buffer', classId: 'inert_gas', units: 15 },
      { id: 'wiring', label: 'Signal Wiring', classId: 'copper', units: 10 },
    ],
    properties: [
      {
        id: 'capacity', label: 'Shield Capacity', unit: 'pts', min: 400, max: 3200,
        groups: [
          { slotId: 'core', weight: 3, stats: [{ stat: 'PE', weight: 3 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'buffer', weight: 1, stats: [{ stat: 'DR', weight: 1 }, { stat: 'OQ', weight: 1 }] },
        ],
      },
      {
        id: 'recharge', label: 'Recharge Rate', unit: 'pts/s', min: 4, max: 42, decimals: 1,
        groups: [
          { slotId: 'wiring', weight: 2, stats: [{ stat: 'CD', weight: 3 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'core', weight: 1, stats: [{ stat: 'CD', weight: 1 }, { stat: 'PE', weight: 1 }] },
        ],
      },
      {
        id: 'hitpoints', label: 'Chassis Hit Points', unit: 'hp', min: 300, max: 1400,
        groups: [
          { slotId: 'frame', weight: 3, stats: [{ stat: 'SR', weight: 2 }, { stat: 'DR', weight: 1 }, { stat: 'OQ', weight: 1 }] },
        ],
      },
      {
        id: 'encumbrance', label: 'Health Encumbrance', unit: 'enc', min: 15, max: 95, invert: true,
        groups: [
          { slotId: 'frame', weight: 2, stats: [{ stat: 'MA', weight: 2 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'buffer', weight: 1, stats: [{ stat: 'OQ', weight: 1 }] },
        ],
      },
    ],
    lines: [
      { id: 'exp_effectiveness', label: 'Experimental Effectiveness', properties: ['capacity', 'recharge'] },
      { id: 'exp_durability', label: 'Experimental Durability', properties: ['hitpoints'] },
      { id: 'exp_quality', label: 'Experimental Quality', properties: ['encumbrance'] },
    ],
  },

  {
    id: 'composite_armor_segment',
    name: 'Composite Armor Segment',
    category: 'Armor',
    complexity: 18,
    xpType: 'Armorsmithing',
    slots: [
      { id: 'plate', label: 'Ablative Plate', classId: 'ferrous', units: 45 },
      { id: 'weave', label: 'Hide Weave', classId: 'hide', units: 30 },
      { id: 'binder', label: 'Polymer Binder', classId: 'petrochem_solid', units: 20 },
      { id: 'lattice', label: 'Crystal Lattice', classId: 'mineral', units: 25 },
    ],
    properties: [
      {
        id: 'kinetic', label: 'Kinetic Effectiveness', unit: '%', min: 12, max: 62, decimals: 1,
        groups: [
          { slotId: 'plate', weight: 2, stats: [{ stat: 'SR', weight: 2 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'weave', weight: 2, stats: [{ stat: 'SR', weight: 1 }, { stat: 'MA', weight: 1 }, { stat: 'OQ', weight: 1 }] },
        ],
      },
      {
        id: 'energy', label: 'Energy Effectiveness', unit: '%', min: 8, max: 58, decimals: 1,
        groups: [
          { slotId: 'lattice', weight: 2, stats: [{ stat: 'HR', weight: 2 }, { stat: 'CD', weight: 1 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'binder', weight: 1, stats: [{ stat: 'PE', weight: 1 }, { stat: 'OQ', weight: 1 }] },
        ],
      },
      {
        id: 'condition', label: 'Condition', unit: 'hp', min: 8000, max: 42000,
        groups: [
          { slotId: 'binder', weight: 2, stats: [{ stat: 'DR', weight: 2 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'plate', weight: 1, stats: [{ stat: 'DR', weight: 1 }, { stat: 'OQ', weight: 1 }] },
        ],
      },
      {
        id: 'enc', label: 'Encumbrance', unit: 'enc', min: 8, max: 64, invert: true,
        groups: [
          { slotId: 'plate', weight: 2, stats: [{ stat: 'MA', weight: 2 }, { stat: 'OQ', weight: 1 }] },
          { slotId: 'weave', weight: 1, stats: [{ stat: 'MA', weight: 1 }, { stat: 'UT', weight: 1 }] },
        ],
      },
    ],
    lines: [
      { id: 'exp_armor', label: 'Experimental Effectiveness', properties: ['kinetic', 'energy'] },
      { id: 'exp_condition', label: 'Experimental Durability', properties: ['condition'] },
      { id: 'exp_enc', label: 'Experimental Quality', properties: ['enc'] },
    ],
  },
];

export const SCHEMATIC_INDEX: ReadonlyMap<string, Schematic> = new Map(
  SCHEMATICS.map((s) => [s.id, s]),
);
