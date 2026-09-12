// Stat bar — compact visual indicator of a single resource stat (0..1000).
// Color-coded: 0 = dim grey, 500 = yellow, 900+ = green, 950+ = cyan (rare).
import { cn } from "../utils/cn";

const STAT_COLORS: Record<string, string> = {
  OQ: "text-amber-300",
  PE: "text-sky-300",
  UT: "text-emerald-300",
  CR: "text-cyan-300",
  CD: "text-violet-300",
  DR: "text-rose-300",
  HR: "text-orange-300",
  MA: "text-lime-300",
  SR: "text-fuchsia-300",
};

function colorFor(v: number) {
  if (v >= 950) return "bg-cyan-400";
  if (v >= 900) return "bg-emerald-400";
  if (v >= 800) return "bg-lime-400";
  if (v >= 600) return "bg-yellow-400";
  if (v >= 400) return "bg-amber-500";
  if (v > 0)   return "bg-zinc-500";
  return "bg-zinc-800";
}

export function StatPill({ stat, value }: { stat: string; value: number }) {
  if (value === 0) return null;
  return (
    <div className="flex items-center gap-1" title={`${stat}: ${value}`}>
      <span className={cn("font-mono text-[10px] font-bold w-5", STAT_COLORS[stat] ?? "text-zinc-400")}>
        {stat}
      </span>
      <span className="font-mono text-[10px] text-zinc-400 w-8 tabular-nums">{value}</span>
      <div className="h-1.5 w-10 bg-zinc-800 rounded-sm overflow-hidden">
        <div
          className={cn("h-full", colorFor(value))}
          style={{ width: `${(value / 1000) * 100}%` }}
        />
      </div>
    </div>
  );
}
