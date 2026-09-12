import type { ReactNode } from 'react';
import { cn } from '@/utils/cn';

export function Panel({
  title,
  subtitle,
  right,
  children,
  className,
  bodyClassName,
}: {
  title?: ReactNode;
  subtitle?: ReactNode;
  right?: ReactNode;
  children: ReactNode;
  className?: string;
  bodyClassName?: string;
}) {
  return (
    <div
      className={cn(
        'rounded-lg border border-zinc-800 bg-zinc-950/70 shadow-[0_0_0_1px_rgba(0,0,0,0.4)] backdrop-blur',
        className,
      )}
    >
      {(title || right) && (
        <div className="flex items-start justify-between gap-4 border-b border-zinc-800 px-4 py-2.5">
          <div className="min-w-0">
            <div className="font-mono text-[11px] font-semibold uppercase tracking-[0.16em] text-zinc-300">
              {title}
            </div>
            {subtitle && <div className="mt-0.5 text-[11px] text-zinc-500">{subtitle}</div>}
          </div>
          {right && <div className="shrink-0">{right}</div>}
        </div>
      )}
      <div className={cn('p-4', bodyClassName)}>{children}</div>
    </div>
  );
}

export function SectionHeader({
  index,
  kicker,
  title,
  lede,
}: {
  index: string;
  kicker?: string;
  title: string;
  lede?: ReactNode;
}) {
  return (
    <div className="mb-8 border-l-2 border-amber-500/70 pl-5">
      <div className="flex items-center gap-3 font-mono text-[11px] uppercase tracking-[0.22em] text-amber-500/90">
        <span>{index}</span>
        {kicker && <span className="text-zinc-600">/ {kicker}</span>}
      </div>
      <h2 className="mt-2 text-2xl font-semibold tracking-tight text-zinc-100 sm:text-3xl">{title}</h2>
      {lede && <p className="mt-3 max-w-3xl text-[15px] leading-relaxed text-zinc-400">{lede}</p>}
    </div>
  );
}

export function Tag({
  children,
  tone = 'zinc',
  className,
}: {
  children: ReactNode;
  tone?: 'zinc' | 'amber' | 'cyan' | 'emerald' | 'rose' | 'violet';
  className?: string;
}) {
  const tones: Record<string, string> = {
    zinc: 'border-zinc-700 bg-zinc-800/60 text-zinc-300',
    amber: 'border-amber-600/50 bg-amber-500/10 text-amber-300',
    cyan: 'border-cyan-600/50 bg-cyan-500/10 text-cyan-300',
    emerald: 'border-emerald-600/50 bg-emerald-500/10 text-emerald-300',
    rose: 'border-rose-600/50 bg-rose-500/10 text-rose-300',
    violet: 'border-violet-600/50 bg-violet-500/10 text-violet-300',
  };
  return (
    <span
      className={cn(
        'inline-flex items-center gap-1 rounded border px-1.5 py-0.5 font-mono text-[10px] uppercase tracking-wider',
        tones[tone],
        className,
      )}
    >
      {children}
    </span>
  );
}

const TOKENS =
  /(\/\/[^\n]*|\/\*[\s\S]*?\*\/)|('(?:[^'\\]|\\.)*'|"(?:[^"\\]|\\.)*"|`(?:[^`\\]|\\.)*`)|\b(export|import|from|interface|type|const|let|readonly|function|return|new|class|extends|implements|if|else|for|of|in|switch|case|as|await|async|null|true|false|void|never|number|string|boolean)\b|\b([A-Z][A-Za-z0-9_]*)\b|\b(\d+(?:\.\d+)?)\b/g;

export function CodeBlock({ code, caption }: { code: string; caption?: string }) {
  const out: ReactNode[] = [];
  let last = 0;
  let m: RegExpExecArray | null;
  let k = 0;
  TOKENS.lastIndex = 0;
  while ((m = TOKENS.exec(code)) !== null) {
    if (m.index > last) out.push(code.slice(last, m.index));
    const [full, comment, str, kw, ty, num] = m;
    const cls = comment
      ? 'text-zinc-600 italic'
      : str
        ? 'text-emerald-400/90'
        : kw
          ? 'text-violet-400'
          : ty
            ? 'text-cyan-300'
            : num
              ? 'text-amber-300'
              : '';
    out.push(
      <span key={k++} className={cls}>
        {full}
      </span>,
    );
    last = m.index + full.length;
  }
  out.push(code.slice(last));

  return (
    <figure className="overflow-hidden rounded-lg border border-zinc-800 bg-black/60">
      {caption && (
        <figcaption className="border-b border-zinc-800 bg-zinc-900/60 px-3 py-1.5 font-mono text-[10px] uppercase tracking-[0.16em] text-zinc-500">
          {caption}
        </figcaption>
      )}
      <pre className="overflow-x-auto px-4 py-3.5 font-mono text-[12px] leading-[1.65] text-zinc-300">
        <code>{out}</code>
      </pre>
    </figure>
  );
}

/** Heat ramp for "% of class cap". SWG players read resources by color. */
export function heat(pct: number): { text: string; bg: string; bar: string } {
  if (pct >= 0.92) return { text: 'text-fuchsia-300', bg: 'bg-fuchsia-500/15', bar: 'bg-fuchsia-400' };
  if (pct >= 0.78) return { text: 'text-amber-300', bg: 'bg-amber-500/15', bar: 'bg-amber-400' };
  if (pct >= 0.6) return { text: 'text-emerald-300', bg: 'bg-emerald-500/15', bar: 'bg-emerald-400' };
  if (pct >= 0.35) return { text: 'text-cyan-300', bg: 'bg-cyan-500/10', bar: 'bg-cyan-400' };
  return { text: 'text-zinc-500', bg: 'bg-zinc-800/40', bar: 'bg-zinc-600' };
}

export function Bullet({ label, children }: { label: string; children: ReactNode }) {
  return (
    <li className="grid grid-cols-[auto_1fr] gap-x-3 py-2.5">
      <span className="mt-0.5 font-mono text-[11px] text-amber-500/80">▸</span>
      <span className="text-[14px] leading-relaxed text-zinc-400">
        <strong className="font-semibold text-zinc-100">{label}</strong>{' '}
        <span className="text-zinc-600">—</span> {children}
      </span>
    </li>
  );
}
