import { useMemo } from 'react';
import { runSelfTests } from '@/engine/selftest';
import type { Galaxy } from '@/engine/resources';
import { Panel, Tag } from '@/components/ui';
import { cn } from '@/utils/cn';

export function SelfTest({ galaxy }: { galaxy: Galaxy }) {
  const results = useMemo(() => runSelfTests(galaxy), [galaxy]);
  const failed = results.filter((r) => !r.pass).length;

  return (
    <Panel
      title="Runtime Invariants"
      subtitle="re-executed on every galaxy regeneration · a bad tuning edit fails here, not in playtest"
      right={
        <Tag tone={failed ? 'rose' : 'emerald'}>
          {failed ? `${failed} failing` : `${results.length}/${results.length} holding`}
        </Tag>
      }
      bodyClassName="p-0"
    >
      <div className="grid gap-px bg-zinc-900 sm:grid-cols-2 xl:grid-cols-3">
        {results.map((r) => (
          <div key={r.name} className="bg-zinc-950 p-3">
            <div className="flex items-center gap-2">
              <span
                className={cn(
                  'inline-block h-1.5 w-1.5 rounded-full',
                  r.pass ? 'bg-emerald-400' : 'bg-rose-500',
                )}
              />
              <span
                className={cn(
                  'font-mono text-[11px] uppercase tracking-wider',
                  r.pass ? 'text-zinc-300' : 'text-rose-300',
                )}
              >
                {r.name}
              </span>
              <span
                className={cn(
                  'ml-auto font-mono text-[10px]',
                  r.pass ? 'text-emerald-500/80' : 'text-rose-400',
                )}
              >
                {r.pass ? 'PASS' : 'FAIL'}
              </span>
            </div>
            <div className="mt-1 font-mono text-[10px] leading-relaxed text-zinc-600">
              {r.detail}
            </div>
          </div>
        ))}
      </div>
    </Panel>
  );
}
